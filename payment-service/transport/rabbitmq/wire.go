//go:build wireinject
// +build wireinject

package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	rabbitmqInfra "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/consumer"
	paymentRepository "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/repository"
	paymentUseCase "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
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

		// Repository layer
		paymentRepository.Set,

		// Use case layer
		paymentUseCase.Set,

		// Presenter layer
		consumer.Set,

		// gRPC Server
		Set,
	)
	return nil
}
