package postgres

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *rolePostgresSQLRepository) FindRoleByName(ctx context.Context, requestId string, name string, tx *gorm.DB) (*orm.Role, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "RolePostgresRepository.FindRoleByName")
	defer span.End()
	role := new(orm.Role)
	if err := tx.WithContext(ctx).
		Where("role = ?", name).
		Preload("AccessControls").
		First(role).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error finding role by name: %v", requestId, err))
		return nil, err
	}
	return role, nil
}
