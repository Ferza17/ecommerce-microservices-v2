// File: wire.go

//go:build wireinject
// +build wireinject

package usecase

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/repository"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// ProvidePaymentProvider initializes the paymentProvider with its dependencies.
func ProvidePaymentProviderUseCase() IPaymentProviderUseCase {
	wire.Build(
		NewPaymentProviderUseCase,                   // Provides IPaymentProvider
		repository.ProvidePaymentProviderRepository, // Provides IPaymentProviderRepository from Repository
		rabbitmq.ProvideRabbitMQInfrastructure,      // Provides IRabbitMQInfrastructure
		telemetry.ProvideTelemetry,                  // Provides ITelemetryInfrastructure
		logger.ProvideLogger,                        // Provides IZapLogger
	)

	return nil
}
