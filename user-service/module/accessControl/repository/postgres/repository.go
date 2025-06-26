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
	IAccessControlPostgresqlRepository interface {
		CreateAccessControl(ctx context.Context, requestId string, accessControl *orm.AccessControl, tx *gorm.DB) (*orm.AccessControl, error)
		UpdateAccessControlById(ctx context.Context, requestId string, accessControl *orm.AccessControl, tx *gorm.DB) (*orm.AccessControl, error)
	}
	accessControlPostgresSQLRepository struct {
		postgresSQLInfrastructure postgres.PostgresSQL
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

var Set = wire.NewSet(NewAccessControlPostgresqlRepository)

func NewAccessControlPostgresqlRepository(
	postgresSQLInfrastructure postgres.PostgresSQL,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IAccessControlPostgresqlRepository {
	return &accessControlPostgresSQLRepository{
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}
