package consumer

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/temporal"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	authWorkflow "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/workflow"
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
		authWorkflow            authWorkflow.IAuthWorkflow
		temporal                temporal.ITemporalInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewAuthConsumer)

func NewAuthConsumer(
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	authUseCase authUseCase.IAuthUseCase,
	authWorkflow authWorkflow.IAuthWorkflow,
	temporal temporal.ITemporalInfrastructure,
	logger logger.IZapLogger,

) IAuthConsumer {
	c := &authConsumer{
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		authUseCase:             authUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		temporal:                temporal,
		logger:                  logger,
		authWorkflow:            authWorkflow,
	}
	return c
}
