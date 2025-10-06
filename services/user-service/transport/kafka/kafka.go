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
	eventKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/event/consumer"
	roleKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/consumer/kafka"
	userKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer/kafka"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/worker"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/attribute"
)

type (
	Transport struct {
		kafkaInfrastructure     kafkaInfrastructure.IKafkaInfrastructure
		workerPool              *pkgWorker.WorkerPool
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
		authKafkaConsumer       authKafkaConsumer.IAuthConsumer
		userKafkaConsumer       userKafkaConsumer.IUserConsumer
		roleKafkaConsumer       roleKafkaConsumer.IRoleConsumer
		eventConsumer           eventKafkaConsumer.IEventConsumer
	}

	handler func(ctx context.Context, message *kafka.Message) error
)

var Set = wire.NewSet(NewTransport)

func NewTransport(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	authKafkaConsumer authKafkaConsumer.IAuthConsumer,
	userKafkaConsumer userKafkaConsumer.IUserConsumer,
	roleKafkaConsumer roleKafkaConsumer.IRoleConsumer,
	eventConsumer eventKafkaConsumer.IEventConsumer,
	logger logger.IZapLogger,
) *Transport {
	return &Transport{
		kafkaInfrastructure:     kafkaInfrastructure,
		workerPool:              pkgWorker.NewWorkerPoolKafkaTaskQueue("Kafka Consumer", 9, 1000),
		telemetryInfrastructure: telemetryInfrastructure,
		authKafkaConsumer:       authKafkaConsumer,
		userKafkaConsumer:       userKafkaConsumer,
		roleKafkaConsumer:       roleKafkaConsumer,
		eventConsumer:           eventConsumer,
		logger:                  logger,
	}
}

func (srv *Transport) Serve(mainCtx context.Context) error {
	srv.workerPool.Start()

	var (
		topics        []string
		kafkaHandlers = srv.RegisterKafkaHandlers()
	)

	for s, _ := range kafkaHandlers {
		topics = append(topics, s)
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
				requestId string
				childCtx  = context.WithoutCancel(mainCtx)
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

				h, ok := kafkaHandlers[*msg.TopicPartition.Topic]
				if !ok {
					srv.logger.Error(fmt.Sprintf("invalid topic %s", *msg.TopicPartition.Topic))
					span.End()
					continue
				}
				task.Handler = h
			}

			srv.workerPool.AddKafkaTaskQueue(task)
			span.End()
		}
	}

	return nil
}

func (srv *Transport) RegisterKafkaHandlers() map[string]handler {
	var handlers = map[string]handler{}

	// SNAPSHOT
	handlers[config.Get().BrokerKafkaTopicUsers.UserUserLogin] = srv.authKafkaConsumer.SnapshotUsersUserLogin
	handlers[config.Get().BrokerKafkaTopicUsers.ConfirmUserUserLogin] = srv.authKafkaConsumer.ConfirmSnapshotUsersUserLogin
	handlers[config.Get().BrokerKafkaTopicUsers.CompensateUserUserLogin] = srv.authKafkaConsumer.CompensateSnapshotUsersUserLogin

	handlers[config.Get().BrokerKafkaTopicUsers.UserUserLogout] = srv.authKafkaConsumer.SnapshotUsersUserLogout
	handlers[config.Get().BrokerKafkaTopicUsers.ConfirmUserUserLogout] = srv.authKafkaConsumer.ConfirmSnapshotUsersUserLogout
	handlers[config.Get().BrokerKafkaTopicUsers.CompensateUserUserLogout] = srv.authKafkaConsumer.CompensateSnapshotUsersUserLogout

	handlers[config.Get().BrokerKafkaTopicUsers.UserUserCreated] = srv.userKafkaConsumer.SnapshotUsersUserCreated
	handlers[config.Get().BrokerKafkaTopicUsers.ConfirmUserUserCreated] = srv.userKafkaConsumer.ConfirmSnapshotUsersUserCreated
	handlers[config.Get().BrokerKafkaTopicUsers.CompensateUserUserCreated] = srv.userKafkaConsumer.CompensateSnapshotUsersUserCreated

	handlers[config.Get().BrokerKafkaTopicUsers.UserUserUpdated] = srv.userKafkaConsumer.SnapshotUsersUserUpdated

	// DLQ
	handlers[config.Get().BrokerKafkaTopicConnectorSinkPgUser.DlqUsers] = srv.userKafkaConsumer.DlqSinkPgUsersUsers
	handlers[config.Get().BrokerKafkaTopicConnectorSinkPgUser.DlqRoles] = srv.roleKafkaConsumer.DlqSinkPgUsersRoles

	handlers[config.Get().BrokerKafkaTopicConnectorSinkMongoEvent.DlqUser] = srv.eventConsumer.DlqSinkMongoEventsUserEventStores

	return handlers
}

func (srv *Transport) Close() {
	if err := srv.kafkaInfrastructure.Close(); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to close kafka infrastructure: %v", err))
	}
	if err := srv.telemetryInfrastructure.Close(); err != nil {
		srv.logger.Error(fmt.Sprintf("error closing telemetry on kafka consumer"))
	}
}
