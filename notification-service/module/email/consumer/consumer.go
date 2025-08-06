package consumer

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/temporal"
	notificationUseCase "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/usecase"
	emailWorkflow "github.com/ferza17/ecommerce-microservices-v2/notification-service/module/email/workflow"
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
		temporal                temporal.ITemporalInfrastructure
		emailWorkflow           emailWorkflow.IEmailWorkflow
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewNotificationConsumer)

func NewNotificationConsumer(
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	notificationUseCase notificationUseCase.INotificationEmailUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	temporal temporal.ITemporalInfrastructure,
	emailWorkflow emailWorkflow.IEmailWorkflow,
	logger logger.IZapLogger) INotificationEmailConsumer {
	return &notificationEmailConsumer{
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		notificationUseCase:     notificationUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		temporal:                temporal,
		emailWorkflow:           emailWorkflow,
		logger:                  logger,
	}
}
