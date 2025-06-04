// File: wire.go

//go:build wireinject
// +build wireinject

package consumer

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// ProvidePaymentConsumer initializes the paymentConsumer using Wire.
func ProvidePaymentConsumer() IPaymentConsumer {
	wire.Build(
		rabbitmq.ProvideRabbitMQInfrastructure, // Provides IRabbitMQInfrastructure
		telemetry.ProvideTelemetry,             // Provides ITelemetryInfrastructure
		usecase.ProvidePaymentUseCase,          // Provides IPaymentUseCase
		logger.ProvideLogger,                   // Provides IZapLogger
		NewPaymentConsumer,                     // Combines all dependencies into IPaymentConsumer
	)
	return nil
}
