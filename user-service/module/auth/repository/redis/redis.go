package redis

import (
	"context"
	redisInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/redis"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/temporal"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IAuthRedisRepository interface {
		SetOtp(ctx context.Context, requestId string, otp string, value string) (err error)
		GetOtp(ctx context.Context, requestId string, otp string) (*string, error)
	}
	authRedisRepository struct {
		redisInfrastructure     redisInfrastructure.IRedisInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		temporal                temporal.ITemporalInfrastructure
		logger                  logger.IZapLogger
	}
)

const ()

var (
	Set         = wire.NewSet(NewAuthRedisRepository)
	RedisKeyOtp = "user:otp:%s:value:user_id"
)

func NewAuthRedisRepository(
	redisInfrastructure redisInfrastructure.IRedisInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	temporal temporal.ITemporalInfrastructure,
	logger logger.IZapLogger) IAuthRedisRepository {
	c := &authRedisRepository{
		redisInfrastructure:     redisInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		temporal:                temporal,
		logger:                  logger,
	}
	c.temporal = temporal.
		RegisterActivity(c.SetOtp).
		RegisterActivity(c.GetOtp)
	return c
}
