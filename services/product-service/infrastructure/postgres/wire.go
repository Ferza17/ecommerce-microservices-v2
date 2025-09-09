//go:build wireinject
// +build wireinject

package postgres

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
)

func ProvidePostgresSQL() *PostgresSQL {
	wire.Build(
		logger.Set,
		Set,
	)
	return nil
}
