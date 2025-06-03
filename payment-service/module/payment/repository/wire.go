// File: wire.go

//go:build wireinject
// +build wireinject

package repository

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// ProvidePaymentRepository wires dependencies for IPaymentRepository.
func ProvidePaymentRepository() IPaymentRepository {
	wire.Build(
		NewPaymentRepository,
		postgresql.ProvidePostgreSQLInfrastructure,
		telemetry.ProvideTelemetry,
		logger.NewZapLogger,
	)
	return nil // Wire will generate the concrete implementation.
}
