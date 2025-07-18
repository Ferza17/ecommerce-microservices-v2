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

		// ACCESS CONTROL
		FindAccessControlByRoleIdAndFullMethodName(ctx context.Context, requestId string, roleId, fullMethodName string, tx *gorm.DB) (*orm.AccessControl, error)
		FindAccessControlByRoleIdAndHttpMethodAndHttpUrl(ctx context.Context, requestId string, roleId, httpMethod, httpUrl string, tx *gorm.DB) (*orm.AccessControl, error)

		FindAccessControlByRoleAndFullMethodName(ctx context.Context, requestId string, role, fullMethodName string, tx *gorm.DB) (*orm.AccessControl, error)
		FindAccessControlByRoleAndHttpMethodAndHttpUrl(ctx context.Context, requestId string, role, httpMethod, httpUrl string, tx *gorm.DB) (*orm.AccessControl, error)

		CreateAccessControl(ctx context.Context, requestId string, accessControl *orm.AccessControl, tx *gorm.DB) (*orm.AccessControl, error)
		UpdateAccessControlById(ctx context.Context, requestId string, accessControl *orm.AccessControl, tx *gorm.DB) (*orm.AccessControl, error)

		// ACCESS CONTROL EXCLUDED
		FindAccessControlExcludedByFullMethodName(ctx context.Context, requestId string, fullMethodName string, tx *gorm.DB) (*orm.AccessControlExcluded, error)
		FindAccessControlExcludedByHttpUrlAndHttpMethod(ctx context.Context, requestId string, method, url string, tx *gorm.DB) (*orm.AccessControlExcluded, error)
		CreateAccessControlExcluded(ctx context.Context, requestId string, accessControlExcluded *orm.AccessControlExcluded, tx *gorm.DB) (*orm.AccessControlExcluded, error)
	}
	accessControlPostgresSQLRepository struct {
		postgresSQLInfrastructure postgres.IPostgresSQL
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

var Set = wire.NewSet(NewAccessControlPostgresqlRepository)

func NewAccessControlPostgresqlRepository(
	postgresSQLInfrastructure postgres.IPostgresSQL,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IAccessControlPostgresqlRepository {
	return &accessControlPostgresSQLRepository{
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}
