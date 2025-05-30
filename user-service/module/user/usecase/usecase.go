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
	IUserUseCase interface {
		CreateUser(ctx context.Context, requestId string, req *userRpc.CreateUserRequest) (*userRpc.CreateUserResponse, error)
		FindUserById(ctx context.Context, requestId string, req *userRpc.FindUserByIdRequest) (*userRpc.User, error)
		UpdateUserById(ctx context.Context, requestId string, req *userRpc.UpdateUserByIdRequest) (*userRpc.UpdateUserByIdResponse, error)
		FindUserByEmailAndPassword(context.Context, string, *userRpc.FindUserByEmailAndPasswordRequest) (*userRpc.User, error)
	}

	userUseCase struct {
		userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository
		rabbitmqInfrastructure   rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure  telemetryInfrastructure.ITelemetryInfrastructure
		authRedisRepository      authRedisRepository.IAuthRedisRepository
		logger                   pkg.IZapLogger
	}
)

func NewUserUseCase(
	userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	authRedisRepository authRedisRepository.IAuthRedisRepository,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) IUserUseCase {
	return &userUseCase{
		userPostgresqlRepository: userPostgresqlRepository,
		rabbitmqInfrastructure:   rabbitmqInfrastructure,
		telemetryInfrastructure:  telemetryInfrastructure,
		authRedisRepository:      authRedisRepository,
		logger:                   logger,
	}
}
