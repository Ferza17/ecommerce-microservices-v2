package bootstrap

import (
	elasticsearchInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/elasticsearch"
	postgreSQLInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgresql"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	productElasticsearchRepository "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/elasticsearch"
	productPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgresql"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

type Bootstrap struct {
	Logger pkg.IZapLogger

	RabbitMQInfrastructure      rabbitmqInfrastructure.IRabbitMQInfrastructure
	PostgreSQLInfrastructure    postgreSQLInfrastructure.IPostgreSQLInfrastructure
	ElasticsearchInfrastructure elasticsearchInfrastructure.IElasticsearchInfrastructure

	ProductPostgresSQLRepository   productPostgresqlRepository.IProductPostgresqlRepository
	ProductElasticsearchRepository productElasticsearchRepository.IProductElasticsearchRepository
	ProductUseCase                 productUseCase.IProductUseCase
}

func NewBootstrap() *Bootstrap {
	logger := pkg.NewZapLogger()

	// Infrastructure
	newRabbitMQInfrastructure := rabbitmqInfrastructure.NewRabbitMQInfrastructure(logger)
	newPostgreSQLInfrastructure := postgreSQLInfrastructure.NewPostgresqlInfrastructure(logger)
	newElasticsearchInfrastructure := elasticsearchInfrastructure.NewElasticsearchInfrastructure(logger)

	// Repository
	newProductPostgresqlRepository := productPostgresqlRepository.NewProductPostgresqlRepository(newPostgreSQLInfrastructure, logger)
	NewProductElasticsearchRepository := productElasticsearchRepository.NewProductElasticsearchRepository(newElasticsearchInfrastructure, logger)
	// UseCase
	newProductUseCase := productUseCase.NewProductUseCase(newProductPostgresqlRepository, newRabbitMQInfrastructure, logger)

	return &Bootstrap{
		Logger:                   logger,
		RabbitMQInfrastructure:   newRabbitMQInfrastructure,
		PostgreSQLInfrastructure: newPostgreSQLInfrastructure,

		// Modules Dependency
		ProductPostgresSQLRepository:   newProductPostgresqlRepository,
		ProductElasticsearchRepository: NewProductElasticsearchRepository,
		ProductUseCase:                 newProductUseCase,
	}
}
