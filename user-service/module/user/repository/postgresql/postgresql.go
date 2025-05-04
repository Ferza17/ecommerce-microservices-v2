package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"gorm.io/gorm"
)

type (
	IUserPostgresqlRepository interface {
		CreateUserWithTransaction(ctx context.Context, requestId string, req *orm.User, tx *gorm.DB) (string, error)

		FindUserByIdWithTransaction(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.User, error)
		FindUserByEmailAndPasswordWithTransaction(ctx context.Context, requestId string, email string, password string, tx *gorm.DB) (*orm.User, error)

		// OpenTransactionWithContext
		OpenTransactionWithContext(ctx context.Context) *gorm.DB
	}

	userPostgresqlRepository struct {
		connector *infrastructure.PostgresqlConnector
		logger    pkg.IZapLogger
	}
)

const userTable = "users"

func NewUserPostgresqlRepository(connector *infrastructure.PostgresqlConnector, logger pkg.IZapLogger) IUserPostgresqlRepository {
	return &userPostgresqlRepository{
		connector: connector,
		logger:    logger,
	}
}

func (r *userPostgresqlRepository) OpenTransactionWithContext(ctx context.Context) *gorm.DB {
	return r.connector.GormDB.WithContext(ctx).Begin()
}
