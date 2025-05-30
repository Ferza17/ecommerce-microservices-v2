package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgresql"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/user/v1"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"gorm.io/gorm"
)

type (
	IUserPostgresqlRepository interface {
		CreateUserWithTransaction(ctx context.Context, requestId string, req *orm.User, tx *gorm.DB) (string, error)

		FindUserById(ctx context.Context, requestId string, id string) (*orm.User, error)
		FindUserByIdWithTransaction(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.User, error)
		FindUserByEmailWithTransaction(ctx context.Context, requestId string, email string, tx *gorm.DB) (*orm.User, error)

		UpdateUserByIdWithTransaction(ctx context.Context, requestId string, req *userRpc.UpdateUserByIdRequest, tx *gorm.DB) (string, error)

		// OpenTransactionWithContext
		OpenTransactionWithContext(ctx context.Context) *gorm.DB
	}

	userPostgresqlRepository struct {
		postgresSQLInfrastructure postgresql.IPostgreSQLInfrastructure
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		logger                    pkg.IZapLogger
	}
)

const userTable = "users"

func NewUserPostgresqlRepository(
	connector postgresql.IPostgreSQLInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) IUserPostgresqlRepository {
	return &userPostgresqlRepository{
		postgresSQLInfrastructure: connector,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}

func (r *userPostgresqlRepository) OpenTransactionWithContext(ctx context.Context) *gorm.DB {
	return r.postgresSQLInfrastructure.GormDB().WithContext(ctx).Begin()
}
