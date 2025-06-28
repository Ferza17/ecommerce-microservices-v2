package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"strconv"
)

func (r *accessControlRedisRepository) GetAccessControl(ctx context.Context, requestId string, role string, fullMethodName string) (bool, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.GetAccessControl")
	defer span.End()

	key := fmt.Sprintf(accessControlPrefixKey, config.Get().ServiceName, role, fullMethodName)
	val, err := r.redisInfrastructure.GetClient().Get(ctx, key).Result()
	if err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error getting otp: %v", requestId, err))
		span.RecordError(err)
		return false, err
	}

	result, err := strconv.ParseBool(val)
	if err != nil {
		span.RecordError(err)
		return false, err
	}

	return result, nil
}
