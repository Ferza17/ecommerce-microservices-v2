//go:build wireinject
// +build wireinject

package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	rabbitmqInfra "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	productService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/product"
	shippingService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/shipping"
	userService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/user"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/consumer"
	paymentRepository "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/repository"
	paymentUseCase "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	paymentProviderRepository "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/repository"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// ProvideGrpcServer wires all dependencies for IGrpcServer
func ProvideGrpcServer() IRabbitMQServer {
	wire.Build(
		// Infrastructure layer
		logger.Set,
		postgresql.Set,
		telemetry.Set,
		rabbitmqInfra.Set,
		shippingService.Set,
		userService.Set,
		productService.Set,

		// Repository layer
		paymentRepository.Set,
		paymentProviderRepository.Set,

		// Use case layer
		paymentUseCase.Set,

		// Presenter layer
		consumer.Set,

		// gRPC Server
		Set,
	)
	return nil
}
