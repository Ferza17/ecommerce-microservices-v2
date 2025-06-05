//go:build wireinject
// +build wireinject

package repository

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// ProvidePaymentProviderRepository initializes the PaymentProviderRepository with its dependencies.
func ProvidePaymentProviderRepository() IPaymentProviderRepository {
	wire.Build(
		postgresql.ProvidePostgreSQLInfrastructure, // Provides IPostgreSQLInfrastructure
		telemetry.ProvideTelemetry,                 // Provides ITelemetryInfrastructure
		logger.ProvideLogger,                       // Provides IZapLogger
		NewPaymentProviderRepository,               // Provides IPaymentProviderRepository
	)
	return nil
}
