package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *accessControlPostgresSQLRepository) FindAccessControlByRoleIdAndHttpMethodAndHttpUrl(ctx context.Context, requestId string, roleId, httpMethod, httpUrl string, tx *gorm.DB) (*orm.AccessControl, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "AccessControlPostgresRepository.FindAccessControlByRoleAndHttpMethodAndHttpUrl")
	defer span.End()
	acl := new(orm.AccessControl)
	if err := tx.WithContext(ctx).
		Table("access_controls ac").
		Where("ac.http_url = ?", httpUrl).
		Where("ac.http_method = ?", httpMethod).
		Where("r.role = ?", roleId).
		First(acl).
		Error; err != nil {
		r.logger.Error("AccessControlPostgresRepository.FindAccessControlByRoleIdAndHttpMethodAndHttpUrl", zap.Error(err))
		return nil, err
	}
	return acl, nil
}
