package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.uber.org/zap"
	"strconv"
)

func (r *accessControlRedisRepository) GetAccessControlHTTPExcluded(ctx context.Context, requestId string, method, url string) (bool, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "AccessControlRedisRepository.GetAccessControlHTTPExcluded")
	defer span.End()

	key := fmt.Sprintf(accessControlHTTPExcludedPrefixKey, config.Get().ServiceName, method, url)
	val, err := r.redisInfrastructure.GetClient().Get(ctx, key).Result()
	if err != nil {
		r.logger.Error("AccessControlRedisRepository.GetAccessControlHTTPExcluded", zap.Error(err), zap.Error(errors.New(fmt.Sprintf("requestId : %s , error getting HTTP Access Control Excluded: %v", requestId, err))))
		return false, err
	}

	result, err := strconv.ParseBool(val)
	if err != nil {
		r.logger.Error("AccessControlRedisRepository.GetAccessControlHTTPExcluded", zap.Error(err))
		return false, err
	}

	return result, nil
}
