package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type (
	IRolePostgresqlRepository interface {
		FindRoleByName(ctx context.Context, requestId string, name string, tx *gorm.DB) (*orm.Role, error)

		CreateRole(ctx context.Context, requestId string, accessControl *orm.Role, tx *gorm.DB) (*orm.Role, error)
		UpdateRoleById(ctx context.Context, requestId string, accessControl *orm.Role, tx *gorm.DB) (*orm.Role, error)
	}
	rolePostgresSQLRepository struct {
		postgresSQLInfrastructure postgres.IPostgresSQL
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

var Set = wire.NewSet(NewRolePostgresqlRepository)

func NewRolePostgresqlRepository(
	postgresSQLInfrastructure postgres.IPostgresSQL,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IRolePostgresqlRepository {
	return &rolePostgresSQLRepository{
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}
