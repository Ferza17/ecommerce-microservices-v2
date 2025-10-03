package kafka

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/kafka"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	notificationEmailConsumer "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/consumer"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/worker"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/attribute"
)

type (
	Transport struct {
		kafkaInfrastructure       kafkaInfrastructure.IKafkaInfrastructure
		workerPool                *pkgWorker.WorkerPool
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		logger                    logger.IZapLogger
		notificationEmailConsumer notificationEmailConsumer.INotificationEmailConsumer
	}
)

var Set = wire.NewSet(NewTransport)

func NewTransport(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
	notificationEmailConsumer notificationEmailConsumer.INotificationEmailConsumer,
) *Transport {
	return &Transport{
		kafkaInfrastructure:       kafkaInfrastructure,
		workerPool:                pkgWorker.NewWorkerPoolKafkaTaskQueue("kafka-consumer", 10, 1000),
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
		notificationEmailConsumer: notificationEmailConsumer,
	}
}

func (srv *Transport) Serve(mainCtx context.Context) error {
	srv.workerPool.Start()

	topics := []string{
		config.Get().BrokerKafkaTopicNotifications.EmailPaymentOrderCreated,
		config.Get().BrokerKafkaTopicNotifications.EmailOtpCreated,
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
				case config.Get().BrokerKafkaTopicNotifications.EmailOtpCreated:
					task.Handler = func(ctx context.Context, message *kafka.Message) error {
						return srv.notificationEmailConsumer.SnapshotNotificationsEmailOtpCreated(childCtx, message)
					}
				case config.Get().BrokerKafkaTopicNotifications.EmailPaymentOrderCreated:
					task.Handler = func(ctx context.Context, message *kafka.Message) error {
						return srv.notificationEmailConsumer.SnapshotNotificationsEmailPaymentOrderCreated(childCtx, message)
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

func (srv *Transport) Close() {
	if err := srv.kafkaInfrastructure.Close(); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to close kafka infrastructure: %v", err))
	}
	if err := srv.telemetryInfrastructure.Close(); err != nil {
		srv.logger.Error(fmt.Sprintf("error closing telemetry on kafka consumer"))
	}
}
