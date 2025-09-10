package rabbitmq

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/event/usecase"
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
		eventUseCase            eventUseCase.IEventUseCase
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewAuthConsumer)

func NewAuthConsumer(
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	authUseCase authUseCase.IAuthUseCase,
	eventUseCase eventUseCase.IEventUseCase,
	logger logger.IZapLogger,

) IAuthConsumer {
	return &authConsumer{
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		authUseCase:             authUseCase,
		eventUseCase:            eventUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
