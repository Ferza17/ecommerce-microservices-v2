//go:build wireinject
// +build wireinject

package kafka

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/elasticsearch"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	eventMongoDBRepository "github.com/ferza17/ecommerce-microservices-v2/product-service/module/event/repository/mongodb"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/event/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/consumer"
	productEsRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/elasticsearch"
	productPgRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgres"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
)

func Provide() *Transport {
	wire.Build(
		logger.Set,

		// Infra
		elasticsearch.Set,
		postgres.Set,
		telemetry.Set,
		kafka.Set,
		mongodb.Set,

		productEsRepo.Set,
		productPgRepo.Set,
		eventMongoDBRepository.Set,

		productUseCase.Set,
		eventUseCase.Set,

		consumer.Set,
		Set,
	)

	return nil
}
