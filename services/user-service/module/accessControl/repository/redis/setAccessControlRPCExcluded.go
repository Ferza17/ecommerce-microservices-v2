package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.uber.org/zap"
)

func (r *accessControlRedisRepository) SetAccessControlRPCExcluded(ctx context.Context, requestId string, fullMethodName string) error {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "AccessControlRedisRepository.SetAccessControlRPCExcluded")
	defer span.End()

	key := fmt.Sprintf(accessControlRPCExcludedPrefixKey, config.Get().UserServiceServiceName, fullMethodName)
	if err := r.redisInfrastructure.
		GetClient().
		SetEX(ctx, key, true, accessControlExcludedTTL).
		Err(); err != nil {
		r.logger.Error("AccessControlRedisRepository.SetAccessControlRPCExcluded", zap.Error(err))
		span.RecordError(err)
		return err
	}
	return nil
}
