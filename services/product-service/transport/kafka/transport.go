package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/kafka"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/event"
	productConsumer "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/consumer"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/uuid"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/protobuf/encoding/protojson"
)

type (
	Transport struct {
		kafkaInfrastructure     kafkaInfrastructure.IKafkaInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		productConsumer         productConsumer.IProductConsumer
		logger                  logger.IZapLogger
		topics                  []string
	}

	handler func(ctx context.Context, message *pbEvent.EventEnvelope) error
)

var Set = wire.NewSet(
	NewTransport,
)

func NewTransport(
	productConsumer productConsumer.IProductConsumer,
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) *Transport {
	return &Transport{
		productConsumer:         productConsumer,
		kafkaInfrastructure:     kafkaInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
		topics: []string{
			config.Get().BrokerKafkaTopicProducts.ProductCreated,
			config.Get().BrokerKafkaTopicProducts.ProductUpdated,
			config.Get().BrokerKafkaTopicProducts.ProductDeleted,
		},
	}
}

func (srv *Transport) Serve(mainCtx context.Context) error {
	var (
		pool          = pond.NewPool(10, pond.WithContext(mainCtx), pond.WithQueueSize(1000), pond.WithNonBlocking(true))
		kafkaHandlers = srv.RegisterKafkaHandlers()
	)

	if err := srv.kafkaInfrastructure.SetupTopics([]string{config.Get().BrokerKafkaTopicConnectorSinkMongoEvent.SourceConnectorEventEnvelopes}); err != nil {
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

			if msg.TopicPartition.Topic != nil && *msg.TopicPartition.Topic == config.Get().BrokerKafkaTopicConnectorSinkMongoEvent.SourceConnectorEventEnvelopes {
				var (
					childCtx   = context.WithoutCancel(mainCtx)
					request    pbEvent.EventEnvelope
					jsonString string
					requestId  = uuid.NewString()
					token      string
					ok         bool
				)
				if err = json.Unmarshal(msg.Value, &jsonString); err == nil {
					// It was double-encoded, use the unescaped string
					if err = protojson.Unmarshal([]byte(jsonString), &request); err != nil {
						srv.logger.Error(fmt.Sprintf("Failed to deserialize after unescape: %v", err))
						continue
					}
				} else {
					// It's normal JSON, unmarshal directly
					if err = protojson.Unmarshal(msg.Value, &request); err != nil {
						srv.logger.Error(fmt.Sprintf("Failed to deserialize: %v", err))
						continue
					}
				}

				childCtx = pkgContext.SetCausationIdToContext(childCtx, request.XId)
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
					srv.logger.Error(fmt.Sprintf("unregistered event type  %s", request.EventType))
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
					span.End()
					return nil
				})
				continue
			}

			srv.logger.Error(fmt.Sprintf("failed to handle message: %v , message should be inserted into outbox", err))
		}
	}

	return nil
}

func (srv *Transport) RegisterKafkaHandlers() map[string]handler {
	var handlers = map[string]handler{}

	handlers[config.Get().BrokerKafkaTopicProducts.ProductCreated] = srv.productConsumer.SnapshotProductsProductCreated
	handlers[config.Get().BrokerKafkaTopicProducts.ProductUpdated] = srv.productConsumer.SnapshotProductsProductUpdated
	handlers[config.Get().BrokerKafkaTopicProducts.ProductDeleted] = srv.productConsumer.SnapshotProductsProductDeleted

	return handlers
}

func (srv *Transport) Close() {
	if err := srv.kafkaInfrastructure.Close(); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to close kafka infrastructure: %v", err))
	}
}
