package bootstrap

import (
	mongoDBInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/mongodb"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	rpcClientInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rpcclient"
	gatewayEventStoreRepository "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/gatewayEventStore/repository/mongodb"
	gatewayEventStoreUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/gatewayEventStore/usecase"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/product/usecase"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type Bootstrap struct {
	Logger pkg.IZapLogger

	MongoDBInfrastructure   mongoDBInfrastructure.IMongoDBInfrastructure
	RabbitMQInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
	RpcClientInfrastructure rpcClientInfrastructure.IRpcClient

	UserUseCase    userUseCase.IUserUseCase
	ProductUseCase productUseCase.IProductUseCase
}

func NewBootstrap() *Bootstrap {
	//pkg
	logger := pkg.NewZapLogger()

	// Infrastructure
	newMongoDBInfrastructure := mongoDBInfrastructure.NewMongoDBInfrastructure(logger)
	newRabbitMQInfrastructure := rabbitmqInfrastructure.NewRabbitMQInfrastructure(logger)
	newRpcClientInfrastructure := rpcClientInfrastructure.NewRpcClient(logger)

	// Repository
	newGatewayEventStoreRepository := gatewayEventStoreRepository.NewGatewayEventStoreRepository(newMongoDBInfrastructure, logger)

	// UseCase
	newGatewayEventStoreUseCase := gatewayEventStoreUseCase.NewGatewayEventStoreUseCase(newGatewayEventStoreRepository, logger)
	newUserUseCase := userUseCase.NewUserUseCase(newRpcClientInfrastructure, newRabbitMQInfrastructure, newGatewayEventStoreUseCase, logger)
	NewProductUseCase := productUseCase.NewProductUseCase(newRpcClientInfrastructure, newRabbitMQInfrastructure, newGatewayEventStoreUseCase, logger)

	return &Bootstrap{
		Logger:                  logger,
		MongoDBInfrastructure:   newMongoDBInfrastructure,
		RabbitMQInfrastructure:  newRabbitMQInfrastructure,
		RpcClientInfrastructure: newRpcClientInfrastructure,
		UserUseCase:             newUserUseCase,
		ProductUseCase:          NewProductUseCase,
	}
}
