package postgres

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *rolePostgresSQLRepository) UpdateRoleById(ctx context.Context, requestId string, role *orm.Role, tx *gorm.DB) (*orm.Role, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.UpdateRoleById")
	defer span.End()
	if err := tx.WithContext(ctx).
		Save(role).
		Where("id = ?", role.ID).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error updating access control: %v", requestId, err))
		return nil, err
	}
	return role, nil
}
