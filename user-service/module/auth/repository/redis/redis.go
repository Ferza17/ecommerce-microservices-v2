package redis

import (
	redisInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/redis"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type (
	IAuthRedisRepository interface{}
	authRedisRepository  struct {
		redisInfrastructure redisInfrastructure.IRedisInfrastructure
		logger              pkg.IZapLogger
	}
)

func NewAuthRedisRepository(redisInfrastructure redisInfrastructure.IRedisInfrastructure, logger pkg.IZapLogger) IAuthRedisRepository {
	return &authRedisRepository{
		redisInfrastructure: redisInfrastructure,
		logger:              logger,
	}
}
