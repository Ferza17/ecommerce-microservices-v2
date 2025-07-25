package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.uber.org/zap"
	"strconv"
)

func (r *accessControlRedisRepository) GetAccessControlRPCExcluded(ctx context.Context, requestId string, fullMethodName string) (bool, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "AccessControlRedisRepository.GetAccessControlRPCExcluded")
	defer span.End()

	key := fmt.Sprintf(accessControlRPCExcludedPrefixKey, config.Get().UserServiceServiceName, fullMethodName)
	val, err := r.redisInfrastructure.GetClient().Get(ctx, key).Result()
	if err != nil {
		r.logger.Error("AccessControlRedisRepository.GetAccessControlRPCExcluded", zap.Error(err))
		span.RecordError(err)
		return false, err
	}

	result, err := strconv.ParseBool(val)
	if err != nil {
		r.logger.Error("AccessControlRedisRepository.GetAccessControlRPCExcluded", zap.Error(err))
		return false, err
	}

	return result, nil
}
