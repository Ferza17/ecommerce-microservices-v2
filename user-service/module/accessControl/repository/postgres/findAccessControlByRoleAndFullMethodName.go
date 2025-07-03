package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *accessControlPostgresSQLRepository) FindAccessControlByRoleAndFullMethodName(ctx context.Context, requestId string, role, fullMethodName string, tx *gorm.DB) (*orm.AccessControl, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "AccessControlPostgresRepository.FindRoleByName")
	defer span.End()
	acl := new(orm.AccessControl)
	if err := tx.WithContext(ctx).
		Table("access_controls").
		Select("access_controls.*").
		Joins("JOIN roles r ON r.id = access_controls.role_id").
		Where("access_controls.full_method_name = ?", fullMethodName).
		Where("r.role = ?", role).
		First(acl).
		Error; err != nil {
		r.logger.Error("AccessControlUseCase.FindAccessControlByRoleAndFullMethodName", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}
	return acl, nil
}
