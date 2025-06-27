package postgres

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *accessControlPostgresSQLRepository) FindRoleByRoleIdAndFullMethodName(ctx context.Context, requestId string, roleId, fullMethodName string, tx *gorm.DB) (*orm.AccessControl, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.FindRoleByName")
	defer span.End()
	role := new(orm.AccessControl)
	if err := tx.WithContext(ctx).
		Where("role_id = ?", roleId).
		Where("full_method_name = ?", fullMethodName).
		First(role).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error finding role by name: %v", requestId, err))
		return nil, err
	}
	return role, nil
}
