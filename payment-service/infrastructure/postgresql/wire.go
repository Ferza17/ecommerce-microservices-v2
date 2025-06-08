// File: wire.go

//go:build wireinject
// +build wireinject

package postgresql

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

// ProvidePostgreSQLInfrastructure wires dependencies for IPostgreSQLInfrastructure.
func ProvidePostgreSQLInfrastructure() IPostgreSQLInfrastructure {
	wire.Build(
		logger.Set,
		Set,
	)
	return nil
}
