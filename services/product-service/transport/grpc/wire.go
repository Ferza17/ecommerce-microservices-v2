//go:build wireinject
// +build wireinject

package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/elasticsearch"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgres"
	userService "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/service/user"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	eventMongoDBRepository "github.com/ferza17/ecommerce-microservices-v2/product-service/module/event/repository/mongodb"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/event/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/presenter"
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

		// Services
		userService.Set,

		productUseCase.Set,
		eventUseCase.Set,
		productPgRepo.Set,
		productEsRepo.Set,
		eventMongoDBRepository.Set,

		presenter.Set,

		Set,
	)

	return nil
}
