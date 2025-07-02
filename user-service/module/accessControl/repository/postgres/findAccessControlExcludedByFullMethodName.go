package postgres

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *accessControlPostgresSQLRepository) FindAccessControlExcludedByFullMethodName(ctx context.Context, requestId string, fullMethodName string, tx *gorm.DB) (*orm.AccessControlExcluded, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "AccessControlPostgresRepository.FindAccessControlExcludedByFullMethodName")
	defer span.End()

	excluded := new(orm.AccessControlExcluded)
	if err := tx.WithContext(ctx).
		Where("full_method_name = ?", fullMethodName).
		First(excluded).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error finding role by name: %v", requestId, err))
		return nil, err
	}

	return excluded, nil
}
