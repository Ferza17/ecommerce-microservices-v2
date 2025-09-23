//go:build wireinject
// +build wireinject

package http

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	productService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/product"
	shippingService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/shipping"

	userService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/user"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	paymentPresenter "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/presenter"
	paymentRepository "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/repository"
	paymentUseCase "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	paymentProviderPresenter "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/presenter"
	paymentProviderRepository "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/repository"
	paymentProviderUseCase "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/usecase"
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
		userService.Set,
		shippingService.Set,
		productService.Set,

		// Repository layer
		paymentRepository.Set,
		paymentProviderRepository.Set,

		// Use case layer
		paymentUseCase.Set,
		paymentProviderUseCase.Set,

		// Presenter layer
		paymentPresenter.Set,
		paymentProviderPresenter.Set,

		// gRPC Server
		Set,
	)
	return nil
}
