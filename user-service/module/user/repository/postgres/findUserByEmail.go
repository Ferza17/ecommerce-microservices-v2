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

func (r *userPostgresqlRepository) FindUserByEmail(ctx context.Context, requestId string, email string, tx *gorm.DB) (*orm.User, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "UserPostgresRepository.FindUserByEmail")
	defer span.End()
	user := new(orm.User)
	if err := tx.WithContext(ctx).
		Where("email = ?", email).
		Preload("Role").
		Preload("Role.AccessControls").
		First(user).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error finding user by email and password: %v", requestId, err))
		return nil, err
	}
	if err := r.temporal.SignalWorkflow(ctx, requestId, "UserPostgresqlRepository.FindUserByEmail", user); err != nil {
		r.logger.Error("UserPostgresqlRepository.FindUserByEmail - Failed to signal workflow",
			zap.String("requestId", requestId),
			zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return user, nil
}
