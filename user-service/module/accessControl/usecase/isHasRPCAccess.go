package usecase

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *accessControlUseCase) IsHasRPCAccess(ctx context.Context, requestId string, role string, fullMethodName string) (bool, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AccessControlUseCase.IsHasRPCAccess")
	defer span.End()
	// Get On Redis First
	isHasAccess, _ := u.accessControlRedisRepository.GetAccessControlRPC(ctx, requestId, role, fullMethodName)
	if isHasAccess {
		return true, nil
	}

	// Check On Postgres
	tx := u.postgresSQL.GormDB.Begin()
	if _, err := u.accessControlPostgresqlRepository.FindAccessControlByRoleAndFullMethodName(ctx, requestId, role, fullMethodName, tx); err != nil {
		tx.Rollback()
		u.logger.Error("AccessControlUseCase.IsHasRPCAccess", zap.String("requestId", requestId), zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, status.Error(codes.NotFound, "acl Not Found")
		}
		return false, status.Error(codes.Internal, "internal server error")
	}

	// Set On Redis
	if err := u.accessControlRedisRepository.SetAccessControlRPC(ctx, requestId, role, fullMethodName); err != nil {
		tx.Rollback()
		u.logger.Error("AccessControlUseCase.IsHasRPCAccess", zap.String("requestId", requestId), zap.String("role", role), zap.Error(err))
		return false, status.Error(codes.Internal, "internal server error")
	}
	tx.Commit()
	return true, nil
}
