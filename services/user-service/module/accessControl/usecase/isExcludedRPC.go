package usecase

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (u *accessControlUseCase) IsExcludedRPC(ctx context.Context, requestId string, fullMethodName string) (bool, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AccessControlUseCase.IsExcludedRPC")
	defer span.End()
	// Get On Redis
	isExcluded, _ := u.accessControlRedisRepository.GetAccessControlRPCExcluded(ctx, requestId, fullMethodName)
	if isExcluded {
		return true, nil
	}

	// Check On Postgres
	tx := u.postgresSQL.GormDB().Begin()
	if _, err := u.accessControlPostgresqlRepository.FindAccessControlExcludedByFullMethodName(ctx, requestId, fullMethodName, tx); err != nil {
		tx.Rollback()
		u.logger.Error("AccessControlUseCase.IsExcludedRPC", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	// Set On Redis
	if err := u.accessControlRedisRepository.SetAccessControlRPCExcluded(ctx, requestId, fullMethodName); err != nil {
		u.logger.Error("AccessControlUseCase.IsExcludedRPC", zap.Error(err))
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil
}
