package consumer

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	notificationUseCase "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg"
)

type (
	INotificationEmailConsumer interface {
		NotificationEmailOTP(ctx context.Context) error
		NotificationEmailPaymentOrderCreated(ctx context.Context) error
	}

	notificationEmailConsumer struct {
		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		notificationUseCase     notificationUseCase.INotificationEmailUseCase
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewNotificationConsumer(
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	notificationUseCase notificationUseCase.INotificationEmailUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) INotificationEmailConsumer {
	return &notificationEmailConsumer{
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		notificationUseCase:     notificationUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
