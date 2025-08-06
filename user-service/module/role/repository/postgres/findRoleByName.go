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

func (r *rolePostgresSQLRepository) FindRoleByName(ctx context.Context, requestId string, name string, tx *gorm.DB) (*orm.Role, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "RolePostgresRepository.FindRoleByName")
	defer span.End()
	role := new(orm.Role)
	if err := tx.WithContext(ctx).
		Where("role = ?", name).
		Preload("AccessControls").
		First(role).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error finding role by name: %v", requestId, err))
		return nil, err
	}
	if err := r.temporal.SignalWorkflow(ctx, requestId, "RolePostgresSQLRepository.FindRoleByName", role); err != nil {
		r.logger.Error("RolePostgresSQLRepository.FindRoleByName - Failed to signal workflow",
			zap.String("requestId", requestId),
			zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return role, nil
}
