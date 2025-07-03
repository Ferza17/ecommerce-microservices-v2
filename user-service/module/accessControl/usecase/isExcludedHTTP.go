package usecase

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

func (u *accessControlUseCase) IsExcludedHTTP(ctx context.Context, requestId string, method, url string) (bool, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AccessControlUseCase.IsExcludedHTTP")
	defer span.End()

	// Hardcoded for swagger api docs
	ok := strings.Contains(url, "/docs/")
	if ok {
		return true, nil
	}

	// Hardcoded for metrics prometheus
	ok = strings.Contains(url, "/metrics")
	if ok {
		return true, nil
	}

	// Get On Redis
	isExcluded, err := u.accessControlRedisRepository.GetAccessControlHTTPExcluded(ctx, requestId, method, url)
	if err != nil {
		u.logger.Error("AccessControlUseCase.IsExcludedHTTP", zap.Error(err))
	}

	if isExcluded {
		return true, nil
	}

	// Check On Postgres
	tx := u.postgresSQL.GormDB.Begin()
	if _, err = u.accessControlPostgresqlRepository.FindAccessControlExcludedByHttpUrlAndHttpMethod(ctx, requestId, method, url, tx); err != nil {
		tx.Rollback()
		u.logger.Error("AccessControlUseCase.IsExcludedHTTP", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	// Set On Redis
	if err = u.accessControlRedisRepository.SetAccessControlHTTPExcluded(ctx, requestId, method, url); err != nil {
		tx.Rollback()
		u.logger.Error("AccessControlUseCase.IsExcludedHTTP", zap.Error(err))
		return false, err
	}

	tx.Commit()
	return true, nil
}
