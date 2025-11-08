package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	authKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/consumer/kafka"
	userKafkaConsumer "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/consumer/kafka"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/uuid"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/protobuf/encoding/protojson"
)

type (
	Transport struct {
		kafkaInfrastructure     kafkaInfrastructure.IKafkaInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
		authKafkaConsumer       authKafkaConsumer.IAuthConsumer
		userKafkaConsumer       userKafkaConsumer.IUserConsumer
	}

	handler func(ctx context.Context, message *pbEvent.EventEnvelope) error
)

var Set = wire.NewSet(NewTransport)

func NewTransport(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	authKafkaConsumer authKafkaConsumer.IAuthConsumer,
	userKafkaConsumer userKafkaConsumer.IUserConsumer,
	logger logger.IZapLogger,
) *Transport {
	return &Transport{
		kafkaInfrastructure:     kafkaInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		authKafkaConsumer:       authKafkaConsumer,
		userKafkaConsumer:       userKafkaConsumer,
		logger:                  logger,
	}
}

func (srv *Transport) Serve(mainCtx context.Context) error {
	var (
		pool          = pond.NewPool(10, pond.WithContext(mainCtx), pond.WithQueueSize(1000), pond.WithNonBlocking(true))
		kafkaHandlers = srv.RegisterKafkaHandlers()
	)

	if err := srv.kafkaInfrastructure.SetupTopics([]string{"source.mongo.outbox.event_envelopes"}); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to setup kafka topics: %v", err))
		return err
	}

	run := true
	for run {
		select {
		case <-mainCtx.Done():
			pool.StopAndWait()
			srv.Close()
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
				childCtx = context.WithoutCancel(mainCtx)
			)
			if msg.TopicPartition.Topic != nil {
				var (
					request pbEvent.EventEnvelope
					// First, check if the message is double-encoded (string containing JSON)
					jsonString string
					requestId  = uuid.NewString()
					token      string
					ok         bool
				)
				if err = json.Unmarshal(msg.Value, &jsonString); err == nil {
					// It was double-encoded, use the unescaped string
					if err = protojson.Unmarshal([]byte(jsonString), &request); err != nil {
						srv.logger.Error(fmt.Sprintf("Failed to deserialize after unescape: %v", err))
						return err
					}
				} else {
					// It's normal JSON, unmarshal directly
					if err = protojson.Unmarshal(msg.Value, &request); err != nil {
						srv.logger.Error(fmt.Sprintf("Failed to deserialize: %v", err))
						return err
					}
				}

				if requestId, ok = request.Metadata[pkgContext.CtxKeyRequestID]; ok {
					childCtx = pkgContext.SetRequestIDToContext(childCtx, requestId)
				}

				if token, ok = request.Metadata[pkgContext.CtxKeyAuthorization]; ok {
					childCtx = pkgContext.SetTokenAuthorizationToContext(childCtx, token)
				}

				// TODO: Add tracing headers

				childCtx, span := srv.telemetryInfrastructure.StartSpanFromKafkaHeader(childCtx, msg.Headers, "KafkaTransport")
				span.SetAttributes(attribute.String("messaging.destination", *msg.TopicPartition.Topic))
				span.SetAttributes(attribute.String(pkgContext.CtxKeyRequestID, requestId))

				handler, ok := kafkaHandlers[request.EventType]
				if !ok {
					srv.logger.Error(fmt.Sprintf("invalid topic %s", *msg.TopicPartition.Topic))
					span.End()
					continue
				}

				pool.SubmitErr(func() error {
					if err = handler(childCtx, &request); err != nil {
						srv.logger.Error(fmt.Sprintf("failed to handle message: %v", err))
						span.RecordError(err)
						span.End()
						return err
					}
					if err = srv.kafkaInfrastructure.CommitMessage(msg); err != nil {
						srv.logger.Error(fmt.Sprintf("failed to commit message: %v", err))
						span.RecordError(err)
						span.End()
						return err
					}
					span.End()
					return nil
				})
			}
		}
	}

	return nil
}

func (srv *Transport) RegisterKafkaHandlers() map[string]handler {
	var handlers = map[string]handler{}

	// SNAPSHOT
	handlers[config.Get().BrokerKafkaTopicUsers.UserUserLogin] = srv.authKafkaConsumer.SnapshotUsersUserLogin
	handlers[config.Get().BrokerKafkaTopicUsers.UserUserLogout] = srv.authKafkaConsumer.SnapshotUsersUserLogout

	handlers[config.Get().BrokerKafkaTopicUsers.UserUserCreated] = srv.userKafkaConsumer.SnapshotUsersUserCreated

	//handlers[config.Get().BrokerKafkaTopicUsers.UserUserUpdated] = srv.userKafkaConsumer.SnapshotUsersUserUpdated

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
