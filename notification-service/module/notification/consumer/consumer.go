package consumer

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	notificationUseCase "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/notification/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg"
)

type (
	INotificationConsumer interface {
		NotificationLoginCreated(ctx context.Context) error
		NotificationUserCreated(ctx context.Context) error
	}

	notificationConsumer struct {
		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		notificationUseCase     notificationUseCase.INotificationUseCase
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewNotificationConsumer(
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	notificationUseCase notificationUseCase.INotificationUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) INotificationConsumer {
	return &notificationConsumer{
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		notificationUseCase:     notificationUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
