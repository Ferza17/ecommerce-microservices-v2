package redis

import (
	"context"
	"fmt"
)

func (r authRedisRepository) GetOtp(ctx context.Context, requestId string, otp string) (*string, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "AuthRedisRepository.GetOtp")
	defer span.End()
	key := fmt.Sprintf(RedisKeyOtp, otp)
	result, err := r.redisInfrastructure.GetClient().Get(ctx, key).Result()
	if err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error getting otp: %v", requestId, err))
		span.RecordError(err)
		return nil, err
	}

	return &result, nil
}
