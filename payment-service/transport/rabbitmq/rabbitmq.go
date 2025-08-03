package rabbitmq

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	paymentConsumer "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/consumer"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/worker"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"log"
)

type (
	IRabbitMQServer interface {
		Serve(ctx context.Context) error
	}

	rabbitMQServer struct {
		rabbitMQ                rabbitmq.IRabbitMQInfrastructure
		workerPool              *pkgWorker.WorkerPool
		paymentConsumer         paymentConsumer.IPaymentConsumer
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

// NewRabbitMQServer creates and returns a new instance of RabbitMQServer with all dependencies.
func NewRabbitMQServer(
	rabbitMQ rabbitmq.IRabbitMQInfrastructure,
	paymentConsumer paymentConsumer.IPaymentConsumer,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IRabbitMQServer {
	return &rabbitMQServer{
		workerPool: pkgWorker.NewWorkerPoolTaskQueue(
			"RabbitMQ Consumer", 9, 1000),
		rabbitMQ:                rabbitMQ,
		paymentConsumer:         paymentConsumer,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}

// Set is a Wire provider set for rabbitMQServer dependencies.
var Set = wire.NewSet(
	NewRabbitMQServer,
)

func (srv *rabbitMQServer) Serve(ctx context.Context) error {
	srv.workerPool.Start()

	queues := []struct {
		Queue    string
		Exchange string
		Topic    string
	}{

		{
			Queue:    config.Get().QueuePaymentOrderCreated,
			Exchange: config.Get().ExchangePaymentDirect,
			Topic:    amqp091.ExchangeDirect,
		},
		{
			Queue:    config.Get().QueuePaymentOrderDelayedCancelled,
			Exchange: config.Get().ExchangePaymentDelayed,
			Topic:    "x-delayed-message",
		},
	}

	for _, queue := range queues {
		if err := srv.rabbitMQ.SetupQueue(
			queue.Exchange,
			queue.Topic,
			queue.Queue,
		); err != nil {
			srv.logger.Error("failed to setup queue", zap.Error(err))
			return err
		}

		go func(queue string) {

			deliveries, err := srv.rabbitMQ.Consume(ctx, queue)
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
						newCtx, _ = context.WithTimeout(ctx, 20)
					)

					for key, value := range d.Headers {
						if key == pkgContext.CtxKeyRequestID {
							requestId = value.(string)
							newCtx = pkgContext.SetRequestIDToContext(ctx, requestId)
						}

						if key == pkgContext.CtxKeyAuthorization {
							newCtx = pkgContext.SetTokenAuthorizationToContext(ctx, value.(string))
						}
					}

					newCtx, span := srv.telemetryInfrastructure.StartSpanFromRabbitMQHeader(newCtx, d.Headers, "RabbitMQTransport")
					span.SetAttributes(attribute.String("messaging.destination", queue))
					span.SetAttributes(attribute.String(pkgContext.CtxKeyRequestID, requestId))

					task := pkgWorker.TaskQueue{
						QueueName: queue,
						Ctx:       newCtx,
						Delivery:  &d,
					}

					// REGISTER HANDLER
					switch queue {
					case config.Get().QueuePaymentOrderCreated:
						task.Handler = func(ctx context.Context, d *amqp091.Delivery) error {
							return srv.paymentConsumer.PaymentOrderCreated(ctx, d)
						}
					case config.Get().QueuePaymentOrderDelayedCancelled:
						task.Handler = func(ctx context.Context, d *amqp091.Delivery) error {
							return srv.paymentConsumer.PaymentOrderDelayedCancelled(ctx, d)
						}
					default:
						log.Fatalf("invalid queue %s", queue)
					}

					task.Ctx = newCtx
					srv.workerPool.AddTaskQueue(task)
					span.End()

				case <-ctx.Done():
					log.Printf("Context cancelled, stopping consumer for queue %s", queue)
					return
				}
			}

		}(queue.Queue)
	}

	<-ctx.Done()
	srv.workerPool.Stop()
	return nil
}
