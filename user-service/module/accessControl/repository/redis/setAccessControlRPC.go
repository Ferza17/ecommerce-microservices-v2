package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.uber.org/zap"
)

func (r *accessControlRedisRepository) SetAccessControlRPC(ctx context.Context, requestId string, role string, fullMethodName string) error {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "AccessControlRedisRepository.SetAccessControlRPC")
	defer span.End()

	key := fmt.Sprintf(accessControlRPCPrefixKey, config.Get().UserServiceServiceName, role, fullMethodName)
	if err := r.redisInfrastructure.
		GetClient().
		SetEX(ctx, key, true, accessControlTTL).
		Err(); err != nil {
		r.logger.Error("AccessControlRedisRepository.SetAccessControlRPC", zap.Error(err))
		span.RecordError(err)
		return err
	}
	return nil
}
