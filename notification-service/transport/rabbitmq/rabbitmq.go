package rabbitmq

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/temporal"
	notificationEmailConsumer "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/consumer"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/worker"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"log"
	"strings"
)

type (
	RabbitMQTransport struct {
		workerPool                *pkgWorker.WorkerPool
		logger                    logger.IZapLogger
		notificationEmailConsumer notificationEmailConsumer.INotificationEmailConsumer
		rabbitmq                  rabbitmq.IRabbitMQInfrastructure
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		temporal                  temporal.ITemporalInfrastructure
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	logger logger.IZapLogger,
	notificationEmailConsumer notificationEmailConsumer.INotificationEmailConsumer,
	rabbitmq rabbitmq.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	temporal temporal.ITemporalInfrastructure,
) *RabbitMQTransport {
	return &RabbitMQTransport{
		workerPool: pkgWorker.NewWorkerPoolTaskQueue(
			"RabbitMQ Consumer", 9, 1000),
		logger:                    logger,
		notificationEmailConsumer: notificationEmailConsumer,
		rabbitmq:                  rabbitmq,
		telemetryInfrastructure:   telemetryInfrastructure,
		temporal:                  temporal,
	}
}

func (srv *RabbitMQTransport) Serve(ctx context.Context) error {
	srv.workerPool.Start()
	queues := []struct {
		Queue    string
		Exchange string
		Topic    string
	}{
		{
			Queue:    config.Get().QueueNotificationEmailOtpCreated,
			Exchange: config.Get().ExchangeNotification,
			Topic:    amqp091.ExchangeDirect,
		},
		{
			Queue:    config.Get().QueueNotificationEmailPaymentOrderCreated,
			Exchange: config.Get().ExchangeNotification,
			Topic:    amqp091.ExchangeDirect,
		},
	}

	for _, queue := range queues {
		if err := srv.rabbitmq.SetupQueue(
			queue.Exchange,
			queue.Topic,
			queue.Queue,
		); err != nil {
			srv.logger.Error("failed to setup queue", zap.Error(err))
			return err
		}

		go func(queue string) {
			deliveries, err := srv.rabbitmq.Consume(ctx, queue)
			if err != nil {
				srv.logger.Error("failed to consume queue", zap.Error(err))
				return
			}

			for {
				select {
				case d, ok := <-deliveries:
					if !ok {
						log.Printf("Consumer channel closed for queue %s", queue)
						return
					}

					var (
						requestId string
					)

					for key, value := range d.Headers {
						if strings.ToLower(key) == strings.ToLower(pkgContext.CtxKeyRequestID) {
							requestId = value.(string)
							ctx = pkgContext.SetRequestIDToContext(ctx, requestId)
						}

						if strings.ToLower(key) == strings.ToLower(pkgContext.CtxKeyAuthorization) {
							ctx = pkgContext.SetTokenAuthorizationToContext(ctx, value.(string))
						}
					}

					ctx, span := srv.telemetryInfrastructure.StartSpanFromRabbitMQHeader(ctx, d.Headers, "RabbitMQTransport")
					span.SetAttributes(attribute.String("messaging.destination", queue))
					span.SetAttributes(attribute.String(pkgContext.CtxKeyRequestID, requestId))

					task := pkgWorker.TaskQueue{
						QueueName: queue,
						Ctx:       context.WithoutCancel(ctx),
						Delivery:  &d,
					}

					// REGISTER HANDLER
					switch queue {
					case config.Get().QueueNotificationEmailOtpCreated:
						task.Handler = func(newCtx context.Context, d *amqp091.Delivery) error {
							return srv.notificationEmailConsumer.NotificationEmailOTP(context.WithoutCancel(ctx), d)
						}
					case config.Get().QueueNotificationEmailPaymentOrderCreated:
						task.Handler = func(newCtx context.Context, d *amqp091.Delivery) error {
							return srv.notificationEmailConsumer.NotificationEmailPaymentOrderCreated(context.WithoutCancel(ctx), d)
						}
					default:
						log.Fatalf("invalid queue %s", queue)
					}

					srv.workerPool.AddTaskQueue(task)
					span.End()

				case <-ctx.Done():
					log.Printf("Context cancelled, stopping consumer for queue %s", queue)
					return
				}
			}

		}(queue.Queue)
	}

	// Setup Temporal
	go func() {
		if err := srv.temporal.Start(); err != nil {
			srv.logger.Error("failed to start temporal server", zap.Error(err))
			return
		}
	}()

	<-ctx.Done()
	srv.workerPool.Stop()
	return nil
}
