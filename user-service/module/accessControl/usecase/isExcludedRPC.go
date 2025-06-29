package usecase

import (
	"context"
)

func (u *accessControlUseCase) IsExcludedRPC(ctx context.Context, requestId string, fullMethodName string) (bool, error) {
	// Get On Redis
	isExcluded, _ := u.accessControlRedisRepository.GetAccessControlRPCExcluded(ctx, requestId, fullMethodName)
	if isExcluded {
		return true, nil
	}

	// Check On Postgres
	tx := u.postgresSQL.GormDB.Begin()
	if _, err := u.accessControlPostgresqlRepository.FindAccessControlExcludedByFullMethodName(ctx, requestId, fullMethodName, tx); err != nil {
		tx.Rollback()
		return false, err
	}

	// Set On Redis
	if err := u.accessControlRedisRepository.SetAccessControlRPCExcluded(ctx, requestId, fullMethodName); err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil
}
