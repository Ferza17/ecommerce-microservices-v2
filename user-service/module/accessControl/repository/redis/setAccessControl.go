package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
)

func (r *accessControlRedisRepository) SetAccessControl(ctx context.Context, requestId string, role string, fullMethodName string) error {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.SetAccessControl")
	defer span.End()

	key := fmt.Sprintf(accessControlPrefixKey, config.Get().ServiceName, role, fullMethodName)
	if err := r.redisInfrastructure.
		GetClient().
		SetEX(ctx, key, true, accessControlTTL).
		Err(); err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error setting otp: %v", requestId, err))
		span.RecordError(err)
		return err
	}
	return nil
}
