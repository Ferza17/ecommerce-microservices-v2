package consumer

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	notificationUseCase "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
)

type (
	INotificationEmailConsumer interface {
		NotificationEmailOTP(ctx context.Context, d *amqp091.Delivery) error
		NotificationEmailPaymentOrderCreated(ctx context.Context, d *amqp091.Delivery) error
	}

	notificationEmailConsumer struct {
		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		notificationUseCase     notificationUseCase.INotificationEmailUseCase
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewNotificationConsumer)

func NewNotificationConsumer(
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	notificationUseCase notificationUseCase.INotificationEmailUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) INotificationEmailConsumer {
	return &notificationEmailConsumer{
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		notificationUseCase:     notificationUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
