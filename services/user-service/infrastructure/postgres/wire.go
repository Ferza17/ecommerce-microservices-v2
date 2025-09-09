//go:build wireinject
// +build wireinject

package postgres

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

func ProvidePostgresInfrastructure() IPostgresSQL {
	wire.Build(
		logger.Set,
		Set,
	)
	return nil
}
