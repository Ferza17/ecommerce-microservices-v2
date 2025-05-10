package postgresql

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
)

func (r *userPostgresqlRepository) FindUserById(ctx context.Context, requestId string, id string) (*orm.User, error) {
	user := new(orm.User)
	if err := r.postgresSQLInfrastructure.GormDB().WithContext(ctx).
		Table(userTable).
		Where("id = ?", id).
		First(user).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error finding user by id: %v", requestId, err))
		return nil, err
	}
	return user, nil
}
