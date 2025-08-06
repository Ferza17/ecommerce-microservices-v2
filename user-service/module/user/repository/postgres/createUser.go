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

func (r *userPostgresqlRepository) CreateUser(ctx context.Context, requestId string, user *orm.User, tx *gorm.DB) (*orm.User, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "UserPostgresRepository.CreateUser")
	defer span.End()
	if err := tx.WithContext(ctx).
		Create(user).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating user: %v", requestId, err))
		return nil, err
	}

	if err := r.temporal.SignalWorkflow(ctx, requestId, "UserPostgresqlRepository.CreateUser", user); err != nil {
		r.logger.Error("UserPostgresqlRepository.CreateUser - Failed to signal workflow",
			zap.String("requestId", requestId),
			zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return user, nil
}
