// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/elasticsearch"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/consumer"
	elasticsearch2 "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/elasticsearch"
	postgres2 "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
)

// Injectors from wire.go:

func ProvideRabbitMQTransport() *RabbitMQTransport {
	iZapLogger := logger.NewZapLogger()
	iTelemetryInfrastructure := telemetry.NewTelemetry(iZapLogger)
	iRabbitMQInfrastructure := rabbitmq.NewRabbitMQInfrastructure(iTelemetryInfrastructure, iZapLogger)
	postgresSQL := postgres.NewPostgresqlInfrastructure(iZapLogger)
	iProductPostgresqlRepository := postgres2.NewProductPostgresqlRepository(postgresSQL, iTelemetryInfrastructure, iZapLogger)
	iElasticsearchInfrastructure := elasticsearch.NewElasticsearchInfrastructure(iTelemetryInfrastructure, iZapLogger)
	iProductElasticsearchRepository := elasticsearch2.NewProductElasticsearchRepository(iElasticsearchInfrastructure, iTelemetryInfrastructure, iZapLogger)
	iProductUseCase := usecase.NewProductUseCase(postgresSQL, iProductPostgresqlRepository, iRabbitMQInfrastructure, iProductElasticsearchRepository, iTelemetryInfrastructure, iZapLogger)
	iProductConsumer := consumer.NewProductConsumer(iRabbitMQInfrastructure, iProductUseCase, iTelemetryInfrastructure, iZapLogger)
	rabbitMQTransport := NewServer(iZapLogger, iProductConsumer, iTelemetryInfrastructure, iRabbitMQInfrastructure)
	return rabbitMQTransport
}
