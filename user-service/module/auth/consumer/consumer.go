package consumer

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
)

type (
	IAuthConsumer interface {
		UserLogin(ctx context.Context, d *amqp091.Delivery) error
	}
	authConsumer struct {
		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		authUseCase             authUseCase.IAuthUseCase
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewAuthConsumer)

func NewAuthConsumer(
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	authUseCase authUseCase.IAuthUseCase,
	logger logger.IZapLogger,

) IAuthConsumer {
	c := &authConsumer{
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		authUseCase:             authUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
	return c
}
