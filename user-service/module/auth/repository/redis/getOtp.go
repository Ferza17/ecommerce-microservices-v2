package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
)

func (r authRedisRepository) GetOtp(ctx context.Context, requestId string, otp string) (*string, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.GetOtp")
	defer span.End()
	key := fmt.Sprintf(enum.REDIS_KEY_OTP.String(), otp)
	result, err := r.redisInfrastructure.GetClient().Get(ctx, key).Result()
	if err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error getting otp: %v", requestId, err))
		span.RecordError(err)
		return nil, err
	}

	return &result, nil
}
