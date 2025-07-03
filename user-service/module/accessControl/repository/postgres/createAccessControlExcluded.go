package postgres

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *accessControlPostgresSQLRepository) CreateAccessControlExcluded(ctx context.Context, requestId string, accessControlExcluded *orm.AccessControlExcluded, tx *gorm.DB) (*orm.AccessControlExcluded, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "AccessControlPostgresRepository.CreateAccessControlExcluded")
	defer span.End()
	if err := tx.WithContext(ctx).
		Create(accessControlExcluded).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating access control: %v", requestId, err))
		return nil, err
	}
	return accessControlExcluded, nil
}
