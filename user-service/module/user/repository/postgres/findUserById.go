package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *userPostgresqlRepository) FindUserById(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.User, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "UserPostgresRepository.FindUserById")
	defer span.End()
	user := new(orm.User)
	if err := tx.WithContext(ctx).
		Where("id = ?", id).
		Preload("Role").
		Preload("Role.AccessControls").
		First(user).
		Error; err != nil {
		r.logger.Error("UserPostgresRepository.FindUserById", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}
	return user, nil
}
