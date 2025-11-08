package consumer

import (
	"context"

	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/kafka"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/event"
	notificationUseCase "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
)

type (
	INotificationEmailConsumer interface {
		SnapshotNotificationsEmailOtpUserLogin(ctx context.Context, message *pbEvent.EventEnvelope) error
		SnapshotNotificationsEmailOtpUserRegister(ctx context.Context, message *pbEvent.EventEnvelope) error

		SnapshotNotificationsEmailPaymentOrderCreated(ctx context.Context, message *pbEvent.EventEnvelope) error
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
