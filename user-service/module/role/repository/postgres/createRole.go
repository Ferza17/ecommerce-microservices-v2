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

func (r *rolePostgresSQLRepository) CreateRole(ctx context.Context, requestId string, role *orm.Role, tx *gorm.DB) (*orm.Role, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "RolePostgresRepository.CreateRole")
	defer span.End()
	if err := tx.WithContext(ctx).
		Create(role).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating access control: %v", requestId, err))
		return nil, err
	}
	if err := r.temporal.SignalWorkflow(ctx, requestId, "RolePostgresSQLRepository.CreateRole", role); err != nil {
		r.logger.Error("RolePostgresSQLRepository.CreateRole - Failed to signal workflow",
			zap.String("requestId", requestId),
			zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return role, nil
}
