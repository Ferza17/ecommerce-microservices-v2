package postgres

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (r *rolePostgresSQLRepository) UpdateRoleById(ctx context.Context, requestId string, role *orm.Role, tx *gorm.DB) (*orm.Role, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "RolePostgresRepository.UpdateRoleById")
	defer span.End()
	if err := tx.WithContext(ctx).
		Save(role).
		Where("id = ?", role.ID).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error updating access control: %v", requestId, err))
		return nil, err
	}
	if err := r.temporal.SignalWorkflow(ctx, requestId, "RolePostgresSQLRepository.UpdateRoleById", role); err != nil {
		r.logger.Error("RolePostgresSQLRepository.UpdateRoleById - Failed to signal workflow",
			zap.String("requestId", requestId),
			zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return role, nil
}
