package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *userPostgresqlRepository) FindUserByEmailAndPasswordWithTransaction(ctx context.Context, requestId string, email string, password string, tx *gorm.DB) (*orm.User, error) {
	//TODO implement me
	panic("implement me")
}
