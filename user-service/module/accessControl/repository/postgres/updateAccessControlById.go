package postgres

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *accessControlPostgresSQLRepository) UpdateAccessControlById(ctx context.Context, requestId string, accessControl *orm.AccessControl, tx *gorm.DB) (*orm.AccessControl, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "AccessControlPostgresRepository.UpdateAccessControlById")
	defer span.End()
	if err := tx.WithContext(ctx).
		Save(accessControl).
		Where("id = ?", accessControl.ID).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error updating access control: %v", requestId, err))
		return nil, err
	}
	return accessControl, nil
}
