package rabbitmq

import (
	"context"
	"log"

	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/worker"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type (
	Transport struct {
		workerPool              *worker.WorkerPool
		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
	}
)

var Set = wire.NewSet(NewTransport)

func NewTransport(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
) *Transport {
	return &Transport{
		workerPool: worker.NewWorkerPoolTaskQueue(
			"RabbitMQ Consumer", 9, 1000),
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
	}
}

func (s *Transport) Serve(ctx context.Context) error {
	s.workerPool.Start()
	queues := []struct {
		Queue    string
		Exchange string
		Topic    string
	}{}

	for _, queue := range queues {
		if err := s.rabbitmqInfrastructure.SetupQueue(
			queue.Exchange,
			queue.Topic,
			queue.Queue,
		); err != nil {
			s.logger.Error("failed to setup queue", zap.Error(err))
			return err
		}

		go func(queue string) {
			deliveries, err := s.rabbitmqInfrastructure.Consume(ctx, queue)
			if err != nil {
				s.logger.Error("failed to consume queue", zap.Error(err))
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

					newCtx, span := s.telemetryInfrastructure.StartSpanFromRabbitMQHeader(newCtx, d.Headers, "RabbitMQTransport")
					span.SetAttributes(attribute.String("messaging.destination", queue))
					span.SetAttributes(attribute.String(pkgContext.CtxKeyRequestID, requestId))

					task := worker.TaskQueue{
						QueueName: queue,
						Ctx:       newCtx,
						Delivery:  &d,
					}

					// REGISTER HANDLER
					switch queue {
					default:
						log.Fatalf("invalid queue %s", queue)
					}

					task.Ctx = newCtx
					s.workerPool.AddTaskQueue(task)

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
	s.workerPool.Stop()
	return nil
}
