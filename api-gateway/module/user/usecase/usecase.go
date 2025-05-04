package usecase

import (
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	rpcClientInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/service"
	gatewayEventStoreUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/gatewayEventStore/usecase"

	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IUserUseCase interface {
	}
	UserUseCase struct {
		rpcClient                rpcClientInfrastructure.Service
		rabbitMQ                 rabbitmqInfrastructure.IRabbitMQInfrastructure
		gatewayEventStoreUseCase gatewayEventStoreUseCase.IGatewayEventStoreUseCase
		logger                   pkg.IZapLogger
	}
)

func NewUserUseCase(
	rpcClient rpcClientInfrastructure.Service,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	gatewayEventStoreUseCase gatewayEventStoreUseCase.IGatewayEventStoreUseCase,
	logger pkg.IZapLogger,
) IUserUseCase {
	return &UserUseCase{
		rpcClient:                rpcClient,
		rabbitMQ:                 rabbitMQ,
		gatewayEventStoreUseCase: gatewayEventStoreUseCase,
		logger:                   logger,
	}
}
