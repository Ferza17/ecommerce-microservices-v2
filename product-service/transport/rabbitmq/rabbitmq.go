package rabbitmq

import (
	"context"
	"log"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	productConsumer "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/consumer"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/worker"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type (
	RabbitMQTransport struct {
		workerPool              *pkgWorker.WorkerPool
		logger                  logger.IZapLogger
		productConsumer         productConsumer.IProductConsumer
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	logger logger.IZapLogger,
	productConsumer productConsumer.IProductConsumer,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
) *RabbitMQTransport {
	return &RabbitMQTransport{
		workerPool: pkgWorker.NewWorkerPoolTaskQueue(
			"RabbitMQ Consumer", 9, 1000),
		logger:                  logger,
		productConsumer:         productConsumer,
		telemetryInfrastructure: telemetryInfrastructure,
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
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
			Queue:    config.Get().QueueProductCreated,
			Exchange: config.Get().ExchangeProduct,
			Topic:    amqp091.ExchangeDirect,
		},
		{
			Queue:    config.Get().QueueProductUpdated,
			Exchange: config.Get().ExchangeProduct,
			Topic:    amqp091.ExchangeDirect,
		},
		{
			Queue:    config.Get().QueueProductDeleted,
			Exchange: config.Get().ExchangeProduct,
			Topic:    amqp091.ExchangeDirect,
		},
	}

	for _, queue := range queues {
		if err := srv.rabbitmqInfrastructure.SetupQueue(
			queue.Exchange,
			queue.Topic,
			queue.Queue,
		); err != nil {
			srv.logger.Error("failed to setup queue", zap.Error(err))
			return err
		}

		go func(queue string) {

			deliveries, err := srv.rabbitmqInfrastructure.Consume(ctx, queue)
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
						requestId         string
						newCtx, cancelCtx = context.WithTimeout(ctx, 20)
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
					case config.Get().QueueProductCreated:
						task.Handler = func(ctx context.Context, d *amqp091.Delivery) error {
							return srv.productConsumer.ProductCreated(ctx, d)
						}
					case config.Get().QueueProductUpdated:
						task.Handler = func(ctx context.Context, d *amqp091.Delivery) error {
							return srv.productConsumer.ProductUpdated(ctx, d)
						}
					case config.Get().QueueProductDeleted:
						task.Handler = func(ctx context.Context, d *amqp091.Delivery) error {
							return srv.productConsumer.ProductDeleted(ctx, d)
						}
					default:
						log.Fatalf("invalid queue %s", queue)
					}

					task.Ctx = newCtx
					srv.workerPool.AddTaskQueue(task)

					cancelCtx()
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
