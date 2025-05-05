package bootstrap

import (
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	rpcClientInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/service"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/product/usecase"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type Bootstrap struct {
	Logger pkg.IZapLogger

	RabbitMQInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
	RpcClientInfrastructure rpcClientInfrastructure.IService

	// For Injection to GraphQL Resolver
	UserUseCase    userUseCase.IUserUseCase
	ProductUseCase productUseCase.IProductUseCase
}

func NewBootstrap() *Bootstrap {
	//pkg
	logger := pkg.NewZapLogger()

	// Infrastructure
	newRabbitMQInfrastructure := rabbitmqInfrastructure.NewRabbitMQInfrastructure(logger)
	newRpcClientInfrastructure := rpcClientInfrastructure.NewRpcClient(logger)

	// UseCase
	newUserUseCase := userUseCase.NewUserUseCase(newRpcClientInfrastructure, newRabbitMQInfrastructure, logger)
	newProductUseCase := productUseCase.NewProductUseCase(newRpcClientInfrastructure, newRabbitMQInfrastructure, logger)

	return &Bootstrap{
		Logger:                  logger,
		RabbitMQInfrastructure:  newRabbitMQInfrastructure,
		RpcClientInfrastructure: newRpcClientInfrastructure,
		UserUseCase:             newUserUseCase,
		ProductUseCase:          newProductUseCase,
	}
}
