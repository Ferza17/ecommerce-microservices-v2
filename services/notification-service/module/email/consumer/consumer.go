package consumer

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/kafka"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	notificationUseCase "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
)

type (
	INotificationEmailConsumer interface {
		SnapshotNotificationsEmailOtpCreated(ctx context.Context, message *kafka.Message) error
		SnapshotNotificationsEmailPaymentOrderCreated(ctx context.Context, message *kafka.Message) error
	}

	notificationEmailConsumer struct {
		kafkaInfrastructure     kafkaInfrastructure.IKafkaInfrastructure
		notificationUseCase     notificationUseCase.INotificationEmailUseCase
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewNotificationConsumer)

func NewNotificationConsumer(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	notificationUseCase notificationUseCase.INotificationEmailUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) INotificationEmailConsumer {
	return &notificationEmailConsumer{
		kafkaInfrastructure:     kafkaInfrastructure,
		notificationUseCase:     notificationUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
