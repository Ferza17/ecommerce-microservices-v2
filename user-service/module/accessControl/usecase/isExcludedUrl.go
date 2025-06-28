package usecase

import (
	"context"
)

func (u *accessControlUseCase) IsExcludedUrl(ctx context.Context, requestId string, url string) (bool, error) {
	// Get On Redis
	isExcluded, _ := u.accessControlRedisRepository.GetAccessControlExcluded(ctx, requestId, url)
	if isExcluded {
		return true, nil
	}

	// Check On Postgres
	tx := u.postgresSQL.GormDB.Begin()
	if _, err := u.accessControlPostgresqlRepository.FindAccessControlExcludedByFullMethodName(ctx, requestId, url, tx); err != nil {
		tx.Rollback()
		return false, err
	}

	// Set On Redis
	if err := u.accessControlRedisRepository.SetAccessControlExcluded(ctx, requestId, url); err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil
}
