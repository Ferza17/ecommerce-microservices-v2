package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
)

func (r *accessControlRedisRepository) SetAccessControlExcluded(ctx context.Context, requestId string, fullMethodName string) error {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.SetAccessControlExcluded")
	defer span.End()

	key := fmt.Sprintf(accessControlExcludedPrefixKey, config.Get().ServiceName, fullMethodName)
	if err := r.redisInfrastructure.
		GetClient().
		SetEX(ctx, key, true, accessControlExcludedTTL).
		Err(); err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error setting otp: %v", requestId, err))
		span.RecordError(err)
		return err
	}
	return nil
}
