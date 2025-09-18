//go:build wireinject
// +build wireinject

package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/elasticsearch"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	userService "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/service/user"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/presenter"
	productEsRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/elasticsearch"
	productPgRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgres"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
)

func ProvideGrpcTransport() *GrpcTransport {
	wire.Build(
		logger.Set,

		// Infra
		elasticsearch.Set,
		postgres.Set,
		rabbitmq.Set,
		telemetry.Set,
		kafka.Set,

		// Services
		userService.Set,

		productUseCase.Set,
		productPgRepo.Set,
		productEsRepo.Set,

		presenter.Set,

		Set,
	)

	return nil
}
