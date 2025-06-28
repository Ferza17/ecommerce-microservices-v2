package redis

import (
	"context"
	redisInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/redis"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
	"time"
)

type (
	IAccessControlRedisRepository interface {
		// ACCESS CONTROL
		SetAccessControl(ctx context.Context, requestId string, role string, fullMethodName string) error
		GetAccessControl(ctx context.Context, requestId string, role string, fullMethodName string) (bool, error)

		// ACCESS EXCLUDED
		SetAccessControlExcluded(ctx context.Context, requestId string, fullMethodName string) error
		GetAccessControlExcluded(ctx context.Context, requestId string, fullMethodName string) (bool, error)
	}

	accessControlRedisRepository struct {
		redisInfrastructure     redisInfrastructure.IRedisInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

var (
	// Set for Wire Dependency Injection
	Set = wire.NewSet(NewAccessControlRedisRepository)

	//prefix key "<service-name>:access_control:<role>:<full_method_name>"
	accessControlPrefixKey = "%s:access_control:%s:%s"
	accessControlTTL       = 15 * time.Minute

	//prefix key "<service-name>:excluded_method:<full_method_name>"
	accessControlExcludedPrefixKey = "%s:access_control_excluded_method:%s"
	accessControlExcludedTTL       = 5 * time.Minute
)

func NewAccessControlRedisRepository(
	redisInfrastructure redisInfrastructure.IRedisInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) IAccessControlRedisRepository {
	return &accessControlRedisRepository{
		redisInfrastructure:     redisInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
