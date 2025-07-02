package postgres

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *rolePostgresSQLRepository) CreateRole(ctx context.Context, requestId string, role *orm.Role, tx *gorm.DB) (*orm.Role, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "RolePostgresRepository.CreateRole")
	defer span.End()
	if err := tx.WithContext(ctx).
		Create(role).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating access control: %v", requestId, err))
		return nil, err
	}
	return role, nil
}
