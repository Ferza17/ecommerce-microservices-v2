package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.uber.org/zap"
	"strconv"
)

func (r *accessControlRedisRepository) GetAccessControlRPC(ctx context.Context, requestId string, role string, fullMethodName string) (bool, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "AuthRedisRepository.GetAccessControlRPC")
	defer span.End()

	key := fmt.Sprintf(accessControlRPCPrefixKey, config.Get().ServiceName, role, fullMethodName)
	val, err := r.redisInfrastructure.GetClient().Get(ctx, key).Result()
	if err != nil {
		r.logger.Error("AccessControlRedisRepository.GetAccessControlRPC", zap.Error(err))
		span.RecordError(err)
		return false, err
	}

	result, err := strconv.ParseBool(val)
	if err != nil {
		r.logger.Error("AccessControlRedisRepository.GetAccessControlRPC", zap.Error(err))
		return false, err
	}

	return result, nil
}
