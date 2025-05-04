package usecase

import (
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	rpcClientInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rpcclient"
	gatewayEventStoreUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/gatewayEventStore/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IProductUseCase interface{}
	ProductUseCase  struct {
		rpcClient                rpcClientInfrastructure.IRpcClient
		rabbitMQ                 rabbitmqInfrastructure.IRabbitMQInfrastructure
		gatewayEventStoreUseCase gatewayEventStoreUseCase.IGatewayEventStoreUseCase
		logger                   pkg.IZapLogger
	}
)

func NewProductUseCase(
	rpcClient rpcClientInfrastructure.IRpcClient,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	gatewayEventStoreUseCase gatewayEventStoreUseCase.IGatewayEventStoreUseCase,
	logger pkg.IZapLogger,
) IProductUseCase {
	return &ProductUseCase{
		rpcClient:                rpcClient,
		rabbitMQ:                 rabbitMQ,
		gatewayEventStoreUseCase: gatewayEventStoreUseCase,
		logger:                   logger,
	}
}
