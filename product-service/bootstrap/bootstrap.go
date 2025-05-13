package bootstrap

import (
	elasticsearchInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/elasticsearch"
	postgreSQLInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgresql"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	productConsumer "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/consumer"
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
	TelemetryInfrastructure     telemetryInfrastructure.ITelemetryInfrastructure

	ProductPostgresSQLRepository   productPostgresqlRepository.IProductPostgresqlRepository
	ProductElasticsearchRepository productElasticsearchRepository.IProductElasticsearchRepository
	ProductUseCase                 productUseCase.IProductUseCase
	ProductConsumer                productConsumer.IProductConsumer
}

func NewBootstrap() *Bootstrap {
	logger := pkg.NewZapLogger()

	// Infrastructure
	newTelemetryInfrastructure := telemetryInfrastructure.NewTelemetry(logger)
	newRabbitMQInfrastructure := rabbitmqInfrastructure.NewRabbitMQInfrastructure(newTelemetryInfrastructure, logger)
	newPostgreSQLInfrastructure := postgreSQLInfrastructure.NewPostgresqlInfrastructure(newTelemetryInfrastructure, logger)
	newElasticsearchInfrastructure := elasticsearchInfrastructure.NewElasticsearchInfrastructure(newTelemetryInfrastructure, logger)

	// Repository
	newProductPostgresqlRepository := productPostgresqlRepository.NewProductPostgresqlRepository(newPostgreSQLInfrastructure, newTelemetryInfrastructure, logger)
	NewProductElasticsearchRepository := productElasticsearchRepository.NewProductElasticsearchRepository(newElasticsearchInfrastructure, newTelemetryInfrastructure, logger)
	// UseCase
	newProductUseCase := productUseCase.NewProductUseCase(newProductPostgresqlRepository, newRabbitMQInfrastructure, NewProductElasticsearchRepository, newTelemetryInfrastructure, logger)
	// Consumer
	productConsumer := productConsumer.NewProductConsumer(newRabbitMQInfrastructure, newProductUseCase, newTelemetryInfrastructure, logger)

	return &Bootstrap{
		Logger:                   logger,
		RabbitMQInfrastructure:   newRabbitMQInfrastructure,
		PostgreSQLInfrastructure: newPostgreSQLInfrastructure,
		TelemetryInfrastructure:  newTelemetryInfrastructure,

		// Modules Dependency
		ProductPostgresSQLRepository:   newProductPostgresqlRepository,
		ProductElasticsearchRepository: NewProductElasticsearchRepository,
		ProductUseCase:                 newProductUseCase,
		ProductConsumer:                productConsumer,
	}
}
