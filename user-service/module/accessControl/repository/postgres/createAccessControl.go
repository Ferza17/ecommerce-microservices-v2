package postgres

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *accessControlPostgresSQLRepository) CreateAccessControl(ctx context.Context, requestId string, accessControl *orm.AccessControl, tx *gorm.DB) (*orm.AccessControl, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.CreateAccessControlWithTransaction")
	defer span.End()
	if err := tx.WithContext(ctx).
		Create(accessControl).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating access control: %v", requestId, err))
		return nil, err
	}
	return accessControl, nil
}
