package consumer

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type (
	IUserConsumer interface {
		UserCreated(ctx context.Context) error
		UserUpdated(ctx context.Context) error
	}
	userConsumer struct {
		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		userUseCase             userUseCase.IUserUseCase
		logger                  pkg.IZapLogger
	}
)

func NewUserConsumer(
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	userUseCase userUseCase.IUserUseCase,
	logger pkg.IZapLogger) IUserConsumer {
	return &userConsumer{
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		userUseCase:             userUseCase,
		logger:                  logger,
	}
}
