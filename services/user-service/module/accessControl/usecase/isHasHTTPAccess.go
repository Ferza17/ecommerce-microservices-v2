package usecase

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *accessControlUseCase) IsHasHTTPAccess(ctx context.Context, requestId string, role string, httpMethod string, httpUrl string) (bool, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AccessControlUseCase.IsHasHTTPAccess")
	defer span.End()
	// Get On Redis First
	isHasAccess, _ := u.accessControlRedisRepository.GetAccessControlHTTP(ctx, requestId, role, httpMethod, httpUrl)
	if isHasAccess {
		return true, nil
	}

	// Check On Postgres
	tx := u.postgresSQL.GormDB().Begin()
	if _, err := u.accessControlPostgresqlRepository.FindAccessControlByRoleAndHttpMethodAndHttpUrl(ctx, requestId, role, httpMethod, httpUrl, tx); err != nil {
		tx.Rollback()
		u.logger.Error("AccessControlUseCase.IsHasHTTPAccess", zap.String("requestId", requestId), zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, status.Error(codes.NotFound, "acl Not Found")
		}
		return false, status.Error(codes.Internal, "internal server error")
	}

	// Set On Redis
	if err := u.accessControlRedisRepository.SetAccessControlHTTP(ctx, requestId, role, httpMethod, httpUrl); err != nil {
		tx.Rollback()
		u.logger.Error("AccessControlUseCase.IsHasHTTPAccess", zap.String("requestId", requestId), zap.Error(err))
		return false, status.Error(codes.Internal, "internal server error")
	}
	tx.Commit()
	return true, nil
}
