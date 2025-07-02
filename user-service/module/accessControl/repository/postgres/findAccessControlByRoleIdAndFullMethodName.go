package postgres

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *accessControlPostgresSQLRepository) FindAccessControlByRoleIdAndFullMethodName(ctx context.Context, requestId string, roleId, fullMethodName string, tx *gorm.DB) (*orm.AccessControl, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "AccessControlPostgresRepository.FindRoleByName")
	defer span.End()
	acl := new(orm.AccessControl)
	if err := tx.WithContext(ctx).
		Table("access_controls ac").
		Where("ac.full_method_name = ?", fullMethodName).
		Where("ac.role_id = ?", roleId).
		First(acl).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error finding role by name: %v", requestId, err))
		return nil, err
	}
	return acl, nil
}
