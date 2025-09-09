package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *accessControlPostgresSQLRepository) FindAccessControlExcludedByHttpUrlAndHttpMethod(ctx context.Context, requestId string, method, url string, tx *gorm.DB) (*orm.AccessControlExcluded, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "AccessControlPostgresRepository.FindAccessControlExcludedByHttpUrlAndHttpMethod")
	defer span.End()

	excluded := new(orm.AccessControlExcluded)
	if err := tx.WithContext(ctx).
		Where("http_url = ?", url).
		Where("http_method = ?", method).
		First(excluded).
		Error; err != nil {
		r.logger.Error("AccessControlPostgresRepository.FindAccessControlExcludedByHttpUrlAndHttpMethod", zap.Error(err))
		return nil, err
	}

	return excluded, nil
}
