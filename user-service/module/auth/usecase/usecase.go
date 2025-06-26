package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/google/wire"

	authRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
)

type (
	IAuthUseCase interface {
		UserLoginByEmailAndPassword(ctx context.Context, requestId string, req *userRpc.AuthLoginByEmailAndPasswordRequest) error
		FindUserByToken(ctx context.Context, requestId string, req *userRpc.AuthFindUserByTokenRequest) (*userRpc.User, error)
		UserVerifyOtp(ctx context.Context, requestId string, req *userRpc.AuthVerifyOtpRequest) (*userRpc.AuthVerifyOtpResponse, error)
	}
	authUseCase struct {
		userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository
		authRedisRepository      authRedisRepository.IAuthRedisRepository
		rabbitmqInfrastructure   rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure  telemetryInfrastructure.ITelemetryInfrastructure
		postgresSQL              *postgres.PostgresSQL
		logger                   logger.IZapLogger
	}
)

var Set = wire.NewSet(NewAuthUseCase)

func NewAuthUseCase(
	userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository,
	authRedisRepository authRedisRepository.IAuthRedisRepository,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	postgresSQL *postgres.PostgresSQL,
	logger logger.IZapLogger,
) IAuthUseCase {
	return &authUseCase{
		userPostgresqlRepository: userPostgresqlRepository,
		authRedisRepository:      authRedisRepository,
		rabbitmqInfrastructure:   rabbitmqInfrastructure,
		telemetryInfrastructure:  telemetryInfrastructure,
		postgresSQL:              postgresSQL,
		logger:                   logger,
	}
}
