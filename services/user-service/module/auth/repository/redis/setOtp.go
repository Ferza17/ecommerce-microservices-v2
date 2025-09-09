package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
)

func (r authRedisRepository) SetOtp(ctx context.Context, requestId string, otp string, value string) (err error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthRedisRepository.SetOtp")
	defer span.End()
	if err = r.redisInfrastructure.
		GetClient().
		SetEX(ctx, fmt.Sprintf(RedisKeyOtp, otp), value, config.Get().OtpExpirationTime).
		Err(); err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error setting otp: %v", requestId, err))
		span.RecordError(err)
		return err
	}
	return err
}
