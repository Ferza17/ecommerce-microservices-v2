package bootstrap

import (
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	authPresenter "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/presenter"
	authServiceInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/service"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/usecase"
	newCartUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/commerce/cart/usecase"
	productServiceInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/product/service"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/product/usecase"
	userServiceInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/user/service"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type Bootstrap struct {
	Logger pkg.IZapLogger

	RabbitMQInfrastructure       rabbitmqInfrastructure.IRabbitMQInfrastructure
	TelemetryInfrastructure      telemetryInfrastructure.ITelemetryInfrastructure
	UserServiceInfrastructure    userServiceInfrastructure.IUserService
	AuthServiceInfrastructure    authServiceInfrastructure.IAuthService
	ProductServiceInfrastructure productServiceInfrastructure.IProductService

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

	// Initialize GRPC Service Client
	newUserServiceInfrastructure := userServiceInfrastructure.NewUserService(pkg.NewCircuitBreaker(config.Get().UserServiceURL, logger), logger)
	newAuthServiceInfrastructure := authServiceInfrastructure.NewAuthService(pkg.NewCircuitBreaker(config.Get().UserServiceURL, logger), logger)
	newProductServiceInfrastructure := productServiceInfrastructure.NewProductService(pkg.NewCircuitBreaker(config.Get().ProductServiceName, logger), logger)

	// UseCase
	newUserUseCase := userUseCase.NewUserUseCase(newUserServiceInfrastructure, newRabbitMQInfrastructure, newTelemetryInfrastructure, logger)
	newProductUseCase := productUseCase.NewProductUseCase(newProductServiceInfrastructure, newRabbitMQInfrastructure, newTelemetryInfrastructure, logger)
	newCartUseCase := newCartUseCase.NewCartUseCase(newRabbitMQInfrastructure, newTelemetryInfrastructure, logger)
	newAuthUseCase := authUseCase.NewAuthUseCase(newAuthServiceInfrastructure, newUserServiceInfrastructure, newRabbitMQInfrastructure, newTelemetryInfrastructure, logger)

	// Presenter (Only for REST public API)
	newAuthPresenter := authPresenter.NewAuthPresenter(newAuthUseCase, newUserUseCase, newTelemetryInfrastructure, logger)

	return &Bootstrap{
		Logger:                       logger,
		RabbitMQInfrastructure:       newRabbitMQInfrastructure,
		UserServiceInfrastructure:    newUserServiceInfrastructure,
		AuthServiceInfrastructure:    newAuthServiceInfrastructure,
		ProductServiceInfrastructure: newProductServiceInfrastructure,
		TelemetryInfrastructure:      newTelemetryInfrastructure,
		UserUseCase:                  newUserUseCase,
		ProductUseCase:               newProductUseCase,
		CartUseCase:                  newCartUseCase,
		AuthUseCase:                  newAuthUseCase,
		AuthPresenter:                newAuthPresenter,
	}
}
