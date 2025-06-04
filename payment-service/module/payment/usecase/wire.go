// File: wire.go

//go:build wireinject
// +build wireinject

package usecase

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/repository"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// ProvidePaymentUseCase wires dependencies for IPaymentUseCase.
func ProvidePaymentUseCase() IPaymentUseCase {
	wire.Build(
		NewPaymentUseCase,
		repository.ProvidePaymentRepository,
		rabbitmq.ProvideRabbitMQInfrastructure,
		telemetry.ProvideTelemetry,
		logger.ProvideLogger,
	)
	return nil // Wire will generate the concrete implementation.
}
