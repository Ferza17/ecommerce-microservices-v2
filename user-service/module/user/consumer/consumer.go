package consumer

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
)

type (
	IUserConsumer interface {
		UserCreated(ctx context.Context, d *amqp091.Delivery) error
		UserUpdated(ctx context.Context, d *amqp091.Delivery) error
	}
	userConsumer struct {
		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		userUseCase             userUseCase.IUserUseCase
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewUserConsumer)

func NewUserConsumer(
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	userUseCase userUseCase.IUserUseCase,
	logger logger.IZapLogger) IUserConsumer {
	return &userConsumer{
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		userUseCase:             userUseCase,
		logger:                  logger,
	}
}
