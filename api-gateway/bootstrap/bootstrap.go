package bootstrap

import (
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	rpcClientInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/service"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	authPresenter "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/presenter"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/usecase"
	newCartUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/commerce/cart/usecase"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/product/usecase"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type Bootstrap struct {
	Logger pkg.IZapLogger

	RabbitMQInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
	RpcClientInfrastructure rpcClientInfrastructure.IService
	TelemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure

	AuthPresenter authPresenter.IAuthPresenter

	// For Injection to GraphQL Resolver
	UserUseCase    userUseCase.IUserUseCase
	ProductUseCase productUseCase.IProductUseCase
	CartUseCase    newCartUseCase.ICartUseCase
	AuthUseCase    authUseCase.IAuthUseCase
}

func NewBootstrap() *Bootstrap {
	//pkg
	logger := pkg.NewZapLogger()

	// Infrastructure
	newTelemetryInfrastructure := telemetryInfrastructure.NewTelemetry(logger)
	newRabbitMQInfrastructure := rabbitmqInfrastructure.NewRabbitMQInfrastructure(logger, newTelemetryInfrastructure)
	newRpcClientInfrastructure := rpcClientInfrastructure.NewRpcClient(logger)

	// UseCase
	newUserUseCase := userUseCase.NewUserUseCase(newRpcClientInfrastructure, newRabbitMQInfrastructure, newTelemetryInfrastructure, logger)
	newProductUseCase := productUseCase.NewProductUseCase(newRpcClientInfrastructure, newRabbitMQInfrastructure, newTelemetryInfrastructure, logger)
	newCartUseCase := newCartUseCase.NewCartUseCase(newRpcClientInfrastructure, newRabbitMQInfrastructure, newTelemetryInfrastructure, logger)
	newAuthUseCase := authUseCase.NewAuthUseCase(newRabbitMQInfrastructure, newTelemetryInfrastructure, newRpcClientInfrastructure, logger)

	// Presenter (Only for REST public API)
	newAuthPresenter := authPresenter.NewAuthPresenter(newAuthUseCase, newUserUseCase, newTelemetryInfrastructure, logger)

	return &Bootstrap{
		Logger:                  logger,
		RabbitMQInfrastructure:  newRabbitMQInfrastructure,
		RpcClientInfrastructure: newRpcClientInfrastructure,
		TelemetryInfrastructure: newTelemetryInfrastructure,
		UserUseCase:             newUserUseCase,
		ProductUseCase:          newProductUseCase,
		CartUseCase:             newCartUseCase,
		AuthUseCase:             newAuthUseCase,
		AuthPresenter:           newAuthPresenter,
	}
}
