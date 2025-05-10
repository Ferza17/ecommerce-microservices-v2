package consumer

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type (
	IAuthConsumer interface {
		UserLogin(ctx context.Context) error
	}
	authConsumer struct {
		rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure
		authUseCase            authUseCase.IAuthUseCase
		logger                 pkg.IZapLogger
	}
)

func NewAuthConsumer(
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	authUseCase authUseCase.IAuthUseCase,
	logger pkg.IZapLogger,

) IAuthConsumer {
	return &authConsumer{
		rabbitmqInfrastructure: rabbitmqInfrastructure,
		authUseCase:            authUseCase,
		logger:                 logger,
	}
}
