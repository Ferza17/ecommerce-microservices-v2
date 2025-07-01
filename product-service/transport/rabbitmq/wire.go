//go:build wireinject
// +build wireinject

package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/elasticsearch"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	productConsumer "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/consumer"
	productEsRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/elasticsearch"
	productPgRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgres"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
)

func ProvideRabbitMQTransport() *RabbitMQTransport {
	wire.Build(
		logger.Set,

		// Infra
		elasticsearch.Set,
		postgres.Set,
		rabbitmq.Set,
		telemetry.Set,

		productUseCase.Set,
		productPgRepo.Set,
		productEsRepo.Set,

		productConsumer.Set,

		Set,
	)

	return nil
}
