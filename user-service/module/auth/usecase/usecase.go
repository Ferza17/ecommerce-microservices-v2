package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/user/v1"

	authRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type (
	IAuthUseCase interface {
		UserLoginByEmailAndPassword(ctx context.Context, requestId string, req *userRpc.UserLoginByEmailAndPasswordRequest) error
		FindUserByToken(ctx context.Context, requestId string, req *userRpc.FindUserByTokenRequest) (*userRpc.User, error)
		UserVerifyOtp(ctx context.Context, requestId string, req *userRpc.UserVerifyOtpRequest) (*userRpc.UserVerifyOtpResponse, error)
	}
	authUseCase struct {
		userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository
		authRedisRepository      authRedisRepository.IAuthRedisRepository
		rabbitmqInfrastructure   rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure  telemetryInfrastructure.ITelemetryInfrastructure
		logger                   pkg.IZapLogger
	}
)

func NewAuthUseCase(
	userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository,
	authRedisRepository authRedisRepository.IAuthRedisRepository,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) IAuthUseCase {
	return &authUseCase{
		userPostgresqlRepository: userPostgresqlRepository,
		authRedisRepository:      authRedisRepository,
		rabbitmqInfrastructure:   rabbitmqInfrastructure,
		telemetryInfrastructure:  telemetryInfrastructure,
		logger:                   logger,
	}
}
