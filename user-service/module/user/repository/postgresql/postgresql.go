package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"gorm.io/gorm"
)

type (
	IUserPostgresqlRepository interface {
		CreateUserWithTransaction(ctx context.Context, requestId string, req *orm.User, tx *gorm.DB) (string, error)

		FindUserById(ctx context.Context, requestId string, id string) (*orm.User, error)
		FindUserByIdWithTransaction(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.User, error)
		FindUserByEmailWithTransaction(ctx context.Context, requestId string, email string, tx *gorm.DB) (*orm.User, error)

		UpdateUserByIdWithTransaction(ctx context.Context, requestId string, req *pb.UpdateUserByIdRequest, tx *gorm.DB) (string, error)

		// OpenTransactionWithContext
		OpenTransactionWithContext(ctx context.Context) *gorm.DB
	}

	userPostgresqlRepository struct {
		postgresSQLInfrastructure postgresql.IPostgreSQLInfrastructure
		logger                    pkg.IZapLogger
	}
)

const userTable = "users"

func NewUserPostgresqlRepository(connector postgresql.IPostgreSQLInfrastructure, logger pkg.IZapLogger) IUserPostgresqlRepository {
	return &userPostgresqlRepository{
		postgresSQLInfrastructure: connector,
		logger:                    logger,
	}
}

func (r *userPostgresqlRepository) OpenTransactionWithContext(ctx context.Context) *gorm.DB {
	return r.postgresSQLInfrastructure.GormDB().WithContext(ctx).Begin()
}
