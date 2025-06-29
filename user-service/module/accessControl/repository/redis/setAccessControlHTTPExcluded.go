package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.uber.org/zap"
)

func (r *accessControlRedisRepository) SetAccessControlHTTPExcluded(ctx context.Context, requestId string, method, url string) error {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "AuthRedisRepository.SetAccessControlHTTPExcluded")
	defer span.End()

	key := fmt.Sprintf(accessControlHTTPExcludedPrefixKey, config.Get().ServiceName, method, url)
	if err := r.redisInfrastructure.
		GetClient().
		SetEX(ctx, key, true, accessControlExcludedTTL).
		Err(); err != nil {
		r.logger.Error("AuthRedisRepository.SetAccessControlHTTPExcluded", zap.Error(err))
		span.RecordError(err)
		return err
	}
	return nil
}
