package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.uber.org/zap"
)

func (r *accessControlRedisRepository) SetAccessControlHTTP(ctx context.Context, requestId string, role string, method, url string) error {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "AuthRedisRepository.SetAccessControlHTTP")
	defer span.End()

	key := fmt.Sprintf(accessControlHTTPPrefixKey, config.Get().ServiceName, role, method, url)
	if err := r.redisInfrastructure.
		GetClient().
		SetEX(ctx, key, true, accessControlTTL).
		Err(); err != nil {
		r.logger.Error("AccessControlRedisRepository.SetAccessControlHTTP", zap.Error(err))
		span.RecordError(err)
		return err
	}
	return nil
}
