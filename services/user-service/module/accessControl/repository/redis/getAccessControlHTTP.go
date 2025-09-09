package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.uber.org/zap"
	"strconv"
)

func (r *accessControlRedisRepository) GetAccessControlHTTP(ctx context.Context, requestId string, role string, method, url string) (bool, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "AccessControlRedisRepository.GetAccessControlHTTP")
	defer span.End()

	key := fmt.Sprintf(accessControlHTTPPrefixKey, config.Get().UserServiceServiceName, role, method, url)
	val, err := r.redisInfrastructure.GetClient().Get(ctx, key).Result()
	if err != nil {
		r.logger.Error("AccessControlRedisRepository.GetAccessControlHTTP", zap.Error(err))
		span.RecordError(err)
		return false, err
	}

	result, err := strconv.ParseBool(val)
	if err != nil {
		r.logger.Error("AccessControlRedisRepository.GetAccessControlHTTP", zap.Error(err))
		return false, err
	}

	return result, nil
}
