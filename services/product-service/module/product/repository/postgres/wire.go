//go:build wireinject
// +build wireinject

package postgres

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgres"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
)

func ProvideProductPostgresSQLRepository() IProductPostgresqlRepository {
	wire.Build(
		postgres.Set,
		telemetryInfrastructure.Set,
		logger.Set,
		Set,
	)

	return nil
}
