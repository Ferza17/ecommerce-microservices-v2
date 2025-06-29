package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	accessControlPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/postgres"
	accessControlRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/redis"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IAccessControlUseCase interface {
		IsExcludedRPC(ctx context.Context, requestId string, url string) (bool, error)
		IsExcludedHTTP(ctx context.Context, requestId string, method, url string) (bool, error)
	}

	accessControlUseCase struct {
		accessControlPostgresqlRepository accessControlPostgresqlRepository.IAccessControlPostgresqlRepository
		accessControlRedisRepository      accessControlRedisRepository.IAccessControlRedisRepository
		telemetryInfrastructure           telemetryInfrastructure.ITelemetryInfrastructure
		postgresSQL                       *postgres.PostgresSQL
		logger                            logger.IZapLogger
	}
)

var Set = wire.NewSet(NewAccessControlUseCase)

func NewAccessControlUseCase(
	accessControlPostgresqlRepository accessControlPostgresqlRepository.IAccessControlPostgresqlRepository,
	accessControlRedisRepository accessControlRedisRepository.IAccessControlRedisRepository,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	postgresSQL *postgres.PostgresSQL,
	logger logger.IZapLogger,
) IAccessControlUseCase {
	return &accessControlUseCase{
		accessControlPostgresqlRepository: accessControlPostgresqlRepository,
		accessControlRedisRepository:      accessControlRedisRepository,
		telemetryInfrastructure:           telemetryInfrastructure,
		postgresSQL:                       postgresSQL,
		logger:                            logger,
	}
}
