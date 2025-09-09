package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.uber.org/zap"
	"strconv"
)

func (r *accessControlRedisRepository) GetAccessControlRPC(ctx context.Context, requestId string, role string, fullMethodName string) (bool, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "AccessControlRedisRepository.GetAccessControlRPC")
	defer span.End()

	key := fmt.Sprintf(accessControlRPCPrefixKey, config.Get().UserServiceServiceName, role, fullMethodName)
	val, err := r.redisInfrastructure.GetClient().Get(ctx, key).Result()
	if err != nil {
		r.logger.Error("AccessControlRedisRepository.GetAccessControlRPC", zap.String("requestId", requestId), zap.Error(err))
		return false, err
	}

	result, err := strconv.ParseBool(val)
	if err != nil {
		r.logger.Error("AccessControlRedisRepository.GetAccessControlRPC", zap.String("requestId", requestId), zap.Error(err))
		return false, err
	}

	return result, nil
}
