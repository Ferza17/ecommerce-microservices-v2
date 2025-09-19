package kafka

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	authKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/consumer/kafka"
	userKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer/kafka"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/worker"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/attribute"
)

type (
	Server struct {
		kafkaInfrastructure     kafkaInfrastructure.IKafkaInfrastructure
		workerPool              *pkgWorker.WorkerPool
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
		authKafkaConsumer       authKafkaConsumer.IAuthConsumer
		userKafkaConsumer       userKafkaConsumer.IUserConsumer
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	authKafkaConsumer authKafkaConsumer.IAuthConsumer,
	userKafkaConsumer userKafkaConsumer.IUserConsumer,
	logger logger.IZapLogger,
) *Server {
	return &Server{
		kafkaInfrastructure:     kafkaInfrastructure,
		workerPool:              pkgWorker.NewWorkerPoolKafkaTaskQueue("Kafka Consumer", 9, 1000),
		telemetryInfrastructure: telemetryInfrastructure,
		authKafkaConsumer:       authKafkaConsumer,
		userKafkaConsumer:       userKafkaConsumer,
		logger:                  logger,
	}
}

func (srv *Server) Serve(mainCtx context.Context) error {
	srv.workerPool.Start()
	topics := []string{
		config.Get().BrokerKafkaTopicUsers.UserUserLogin,
		config.Get().BrokerKafkaTopicUsers.UserUserLogout,
		config.Get().BrokerKafkaTopicUsers.UserUserCreated,
		config.Get().BrokerKafkaTopicUsers.UserUserUpdated,
	}

	if err := srv.kafkaInfrastructure.SetupTopics(topics); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to setup kafka topics: %v", err))
		return err
	}

	run := true
	for run {
		select {
		case <-mainCtx.Done():
			srv.workerPool.Stop()
			run = false
		default:
			msg, err := srv.kafkaInfrastructure.ReadMessage(time.Second * 2)
			if err != nil {
				srv.logger.Error(fmt.Sprintf("failed to read message: %v", err))
				return err
			}

			if msg == nil {
				continue
			}

			var (
				requestId   string
				childCtx, _ = context.WithTimeout(mainCtx, 20*time.Second)
			)

			for _, header := range msg.Headers {
				if strings.ToLower(header.Key) == strings.ToLower(pkgContext.CtxKeyRequestID) {
					requestId = string(header.Value)
					childCtx = pkgContext.SetRequestIDToContext(childCtx, requestId)
				}

				if strings.ToLower(header.Key) == strings.ToLower(pkgContext.CtxKeyAuthorization) {
					childCtx = pkgContext.SetTokenAuthorizationToContext(childCtx, string(header.Value))
				}
			}
			childCtx, span := srv.telemetryInfrastructure.StartSpanFromKafkaHeader(childCtx, msg.Headers, "KafkaTransport")

			task := pkgWorker.KafkaTaskQueue{
				Message: msg,
				Ctx:     childCtx,
			}

			if msg.TopicPartition.Topic != nil {
				span.SetAttributes(attribute.String("messaging.destination", *msg.TopicPartition.Topic))
				span.SetAttributes(attribute.String(pkgContext.CtxKeyRequestID, requestId))

				switch *msg.TopicPartition.Topic {
				case config.Get().BrokerKafkaTopicUsers.UserUserLogin:
					task.Handler = func(ctx context.Context, message *kafka.Message) error {
						return srv.authKafkaConsumer.SnapshotUsersUserLogin(childCtx, message)
					}
				case config.Get().BrokerKafkaTopicUsers.UserUserLogout:
					task.Handler = func(ctx context.Context, message *kafka.Message) error {
						return srv.authKafkaConsumer.SnapshotUsersUserLogout(childCtx, message)
					}
				case config.Get().BrokerKafkaTopicUsers.UserUserCreated:
					task.Handler = func(ctx context.Context, message *kafka.Message) error {
						return srv.userKafkaConsumer.SnapshotUsersUserCreated(childCtx, message)
					}
				case config.Get().BrokerKafkaTopicUsers.UserUserUpdated:
					task.Handler = func(ctx context.Context, message *kafka.Message) error {
						return srv.userKafkaConsumer.SnapshotUsersUserUpdated(childCtx, message)
					}
				default:
					srv.logger.Error(fmt.Sprintf("invalid topic %s", *msg.TopicPartition.Topic))
					continue
				}
			}

			srv.workerPool.AddKafkaTaskQueue(task)
			span.End()
		}
	}

	return nil
}

func (srv *Server) Close() {
	if err := srv.kafkaInfrastructure.Close(); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to close kafka infrastructure: %v", err))
	}
	if err := srv.telemetryInfrastructure.Close(); err != nil {
		srv.logger.Error(fmt.Sprintf("error closing telemetry on kafka consumer"))
	}
}
