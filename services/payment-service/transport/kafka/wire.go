//go:build wireinject
// +build wireinject

package kafka

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	productService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/product"
	shippingService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/shipping"
	userService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/user"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	eventMongodbRepository "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/event/repository/mongodb"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/event/usecase"
	paymentConsumer "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/consumer"
	paymentRepository "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/repository"
	paymentUseCase "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	paymentProviderRepository "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/repository"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

func Provide() *Transport {
	wire.Build(
		// Infrastructure layer
		logger.Set,
		postgresql.Set,
		telemetry.Set,
		kafka.Set,
		mongodb.Set,
		userService.Set,
		shippingService.Set,
		productService.Set,

		// Repository layer
		paymentRepository.Set,
		paymentProviderRepository.Set,
		eventMongodbRepository.Set,

		// Use case layer
		paymentUseCase.Set,
		eventUseCase.Set,

		// Presenter layer
		paymentConsumer.Set,

		// gRPC Server
		Set,
	)
	return nil
}
