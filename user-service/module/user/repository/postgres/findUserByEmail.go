package postgres

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *userPostgresqlRepository) FindUserByEmail(ctx context.Context, requestId string, email string, tx *gorm.DB) (*orm.User, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.FindUserByEmailWithTransaction")
	defer span.End()
	user := new(orm.User)
	if err := tx.WithContext(ctx).
		Where("email = ?", email).
		Preload("Role").
		Preload("Role.AccessControls").
		First(user).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error finding user by email and password: %v", requestId, err))
		return nil, err
	}
	return user, nil
}
