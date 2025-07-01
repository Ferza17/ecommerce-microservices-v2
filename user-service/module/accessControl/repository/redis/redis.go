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
		// ACCESS CONTROL RPC
		SetAccessControlRPC(ctx context.Context, requestId string, role string, fullMethodName string) error
		GetAccessControlRPC(ctx context.Context, requestId string, role string, fullMethodName string) (bool, error)

		// ACCESS CONTROL HTTP
		SetAccessControlHTTP(ctx context.Context, requestId string, role string, method, url string) error
		GetAccessControlHTTP(ctx context.Context, requestId string, role string, method, url string) (bool, error)

		// ACCESS EXCLUDED RPC
		SetAccessControlRPCExcluded(ctx context.Context, requestId string, fullMethodName string) error
		GetAccessControlRPCExcluded(ctx context.Context, requestId string, fullMethodName string) (bool, error)

		// ACCESS EXCLUDED HTTP
		SetAccessControlHTTPExcluded(ctx context.Context, requestId string, method, url string) error
		GetAccessControlHTTPExcluded(ctx context.Context, requestId string, method, url string) (bool, error)
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

	accessControlTTL = 15 * time.Minute
	//prefix key "<service-name>:access_control:rpc:<role>:<full_method_name>"
	accessControlRPCPrefixKey = "%s:access_control:rpc:%s:%s"
	//prefix key "<service-name>:access_control:http:<role>:<http_method>:<http_url>"
	accessControlHTTPPrefixKey = "%s:access_control:http:%s:%s:%s"

	accessControlExcludedTTL = 5 * time.Minute
	//prefix key "<service-name>:rpc:access_control_excluded_method:<full_method_name>"
	accessControlRPCExcludedPrefixKey = "%s:rpc:access_control_excluded_method:%s"

	//prefix key "<service-name>:http:<http_method>:access_control_excluded_method:<full_method_name>"
	accessControlHTTPExcludedPrefixKey = "%s:http:%s:access_control_excluded_method:%s"
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
