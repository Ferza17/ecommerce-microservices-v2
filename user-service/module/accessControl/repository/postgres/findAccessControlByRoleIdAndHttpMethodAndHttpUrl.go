package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *accessControlPostgresSQLRepository) FindAccessControlByRoleIdAndHttpMethodAndHttpUrl(ctx context.Context, requestId string, roleId, httpMethod, httpUrl string, tx *gorm.DB) (*orm.AccessControl, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "AccessControlRepository.FindAccessControlByRoleIdAndHttpMethodAndHttpUrl")
	defer span.End()
	role := new(orm.AccessControl)
	if err := tx.WithContext(ctx).
		Where("role_id = ?", roleId).
		Where("http_url = ?", httpUrl).
		Where("http_method = ?", httpMethod).
		First(role).
		Error; err != nil {
		r.logger.Error("AccessControlRepository.FindAccessControlByRoleIdAndHttpMethodAndHttpUrl", zap.Error(err))
		return nil, err
	}
	return role, nil
}
