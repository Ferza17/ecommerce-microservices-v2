package bootstrap

import (
	postgreSQLInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgresql"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	productPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgresql"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

type Bootstrap struct {
	Logger pkg.IZapLogger

	RabbitMQInfrastructure   rabbitmqInfrastructure.IRabbitMQInfrastructure
	PostgreSQLInfrastructure postgreSQLInfrastructure.IPostgreSQLInfrastructure

	ProductRepository productPostgresqlRepository.IProductPostgresqlRepository
	ProductUseCase    productUseCase.IProductUseCase
}

func NewBootstrap() *Bootstrap {
	logger := pkg.NewZapLogger()

	// Infrastructure
	newRabbitMQInfrastructure := rabbitmqInfrastructure.NewRabbitMQInfrastructure(logger)
	newPostgreSQLInfrastructure := postgreSQLInfrastructure.NewPostgresqlInfrastructure(logger)

	// Repository
	newProductPostgresqlRepository := productPostgresqlRepository.NewProductPostgresqlRepository(newPostgreSQLInfrastructure, logger)

	// UseCase
	newProductUseCase := productUseCase.NewProductUseCase(newProductPostgresqlRepository, newRabbitMQInfrastructure, logger)

	return &Bootstrap{
		Logger:                   logger,
		RabbitMQInfrastructure:   newRabbitMQInfrastructure,
		PostgreSQLInfrastructure: newPostgreSQLInfrastructure,

		// Modules Dependency
		ProductRepository: newProductPostgresqlRepository,
		ProductUseCase:    newProductUseCase,
	}
}
