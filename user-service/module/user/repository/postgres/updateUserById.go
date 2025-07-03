package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
	"log"
)

func (r *userPostgresqlRepository) UpdateUserById(ctx context.Context, requestId string, req *orm.User, tx *gorm.DB) (*orm.User, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "UserPostgresRepository.UpdateUserById")
	defer span.End()

	err := tx.WithContext(ctx).Save(req).Error
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}

	return req, nil
}
