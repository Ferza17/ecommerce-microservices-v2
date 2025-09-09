package rabbitmq

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	authConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/consumer"
	userConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/worker"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"log"
	"strings"
	"time"
)

type (
	Server struct {
		workerPool              *pkgWorker.WorkerPool
		amqpInfrastructure      rabbitmq.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
		userConsumer            userConsumer.IUserConsumer
		authConsumer            authConsumer.IAuthConsumer
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	amqpInfrastructure rabbitmq.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
	userConsumer userConsumer.IUserConsumer,
	authConsumer authConsumer.IAuthConsumer,
) *Server {
	return &Server{
		workerPool: pkgWorker.NewWorkerPoolTaskQueue(
			"RabbitMQ Consumer", 9, 1000),
		amqpInfrastructure:      amqpInfrastructure,
		logger:                  logger,
		userConsumer:            userConsumer,
		authConsumer:            authConsumer,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}

func (srv *Server) Serve(ctx context.Context) error {
	srv.workerPool.Start()

	queues := []struct {
		Queue    string
		Exchange string
		Topic    string
	}{
		// AUTH QUEUE
		{
			Queue:    config.Get().QueueUserLogin,
			Exchange: config.Get().ExchangeUser,
			Topic:    amqp091.ExchangeDirect,
		},
		// USER QUEUE
		{
			Queue:    config.Get().QueueUserCreated,
			Exchange: config.Get().ExchangeUser,
			Topic:    amqp091.ExchangeDirect,
		},
		{
			Queue:    config.Get().QueueUserUpdated,
			Exchange: config.Get().ExchangeUser,
			Topic:    amqp091.ExchangeDirect,
		},
	}

	for _, queue := range queues {
		if err := srv.amqpInfrastructure.SetupQueue(
			queue.Exchange,
			queue.Topic,
			queue.Queue,
		); err != nil {
			srv.logger.Error("failed to setup queue", zap.Error(err))
			return err
		}
		go func(queue string) {

			deliveries, err := srv.amqpInfrastructure.Consume(ctx, queue)
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
						newCtx, _ = context.WithTimeout(ctx, 20*time.Second)
					)

					for key, value := range d.Headers {
						if strings.ToLower(key) == strings.ToLower(pkgContext.CtxKeyRequestID) {
							requestId = value.(string)
							newCtx = pkgContext.SetRequestIDToContext(newCtx, requestId)
						}

						if strings.ToLower(key) == strings.ToLower(pkgContext.CtxKeyAuthorization) {
							newCtx = pkgContext.SetTokenAuthorizationToContext(newCtx, value.(string))
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
					case config.Get().QueueUserLogin:
						task.Handler = func(ctx context.Context, d *amqp091.Delivery) error {
							return srv.authConsumer.UserLogin(newCtx, d)
						}
					case config.Get().QueueUserCreated:
						task.Handler = func(ctx context.Context, d *amqp091.Delivery) error {
							return srv.userConsumer.UserCreated(newCtx, d)
						}
					case config.Get().QueueUserUpdated:
						task.Handler = func(ctx context.Context, d *amqp091.Delivery) error {
							return srv.userConsumer.UserCreated(newCtx, d)
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
