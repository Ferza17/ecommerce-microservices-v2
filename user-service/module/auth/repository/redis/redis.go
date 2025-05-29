package redis

import (
	"context"
	redisInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/redis"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type (
	IAuthRedisRepository interface {
		SetOtp(ctx context.Context, requestId string, otp string, value string) (err error)
		GetOtp(ctx context.Context, requestId string, otp string) (*string, error)
	}
	authRedisRepository struct {
		redisInfrastructure     redisInfrastructure.IRedisInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewAuthRedisRepository(
	redisInfrastructure redisInfrastructure.IRedisInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) IAuthRedisRepository {
	return &authRedisRepository{
		redisInfrastructure:     redisInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
