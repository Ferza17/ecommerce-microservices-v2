package redis

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	redisClient "github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

type (
	IRedisInfrastructure interface {
		GetClient() *redisClient.Client
		Close() error
	}

	redisInfrastructure struct {
		client *redisClient.Client
		logger logger.IZapLogger
	}
)

var Set = wire.NewSet(NewRedisInfrastructure)

func NewRedisInfrastructure(logger logger.IZapLogger) IRedisInfrastructure {
	client := redisClient.NewClient(&redisClient.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Get().RedisHost, config.Get().RedisPort),
		Password: config.Get().RedisPassword,
		DB:       config.Get().RedisDB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		logger.Error(fmt.Sprintf("failed to connect to redis: %v", err))
	}

	return &redisInfrastructure{
		client: client,
		logger: logger,
	}
}

func (r *redisInfrastructure) GetClient() *redisClient.Client {
	return r.client
}

func (r *redisInfrastructure) Close() error {
	if err := r.client.Close(); err != nil {
		r.logger.Error(fmt.Sprintf("failed to close redis client: %v", err))
		return err
	}
	return nil
}
