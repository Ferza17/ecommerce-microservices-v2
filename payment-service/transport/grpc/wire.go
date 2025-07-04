//go:build wireinject
// +build wireinject

package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
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

// ProvideGrpcServer wires all dependencies for IGrpcServer
func ProvideGrpcServer() IGrpcServer {
	wire.Build(
		// Infrastructure layer
		logger.Set,
		postgresql.Set,
		telemetry.Set,
		rabbitmq.Set,
		userService.Set,

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
