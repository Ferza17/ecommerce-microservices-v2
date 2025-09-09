//go:build wireinject
// +build wireinject

package postgres

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

func ProvideRoleRepository() IRolePostgresqlRepository {
	wire.Build(
		postgres.Set,
		telemetryInfrastructure.Set,
		logger.Set,
		Set,
	)

	return nil
}
