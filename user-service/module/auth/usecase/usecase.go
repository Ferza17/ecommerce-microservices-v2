package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	accessControlUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/usecase"
	rolePostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/emptypb"

	authRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
)

type (
	IAuthUseCase interface {
		// COMMAND
		AuthUserRegister(ctx context.Context, requestId string, req *pb.AuthUserRegisterRequest) (*emptypb.Empty, error)
		AuthUserLoginByEmailAndPassword(ctx context.Context, requestId string, req *pb.AuthUserLoginByEmailAndPasswordRequest) (*pb.AuthUserLoginByEmailAndPasswordResponse, error)
		AuthUserVerifyOtp(ctx context.Context, requestId string, req *pb.AuthUserVerifyOtpRequest) (*pb.AuthUserVerifyOtpResponse, error)
		AuthUserLogoutByToken(ctx context.Context, requestId string, req *pb.AuthUserLogoutByTokenRequest) (*pb.AuthUserLogoutByTokenResponse, error)
		AuthUserVerifyAccessControl(ctx context.Context, requestId string, req *pb.AuthUserVerifyAccessControlRequest) (*pb.AuthUserVerifyAccessControlResponse, error)

		// QUERY
		AuthUserFindUserByToken(ctx context.Context, requestId string, req *pb.AuthUserFindUserByTokenRequest) (*pb.AuthUserFindUserByTokenResponse, error)
	}
	authUseCase struct {
		userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository
		rolePostgresqlRepository rolePostgresqlRepository.IRolePostgresqlRepository

		authRedisRepository authRedisRepository.IAuthRedisRepository

		accessControlUseCase accessControlUseCase.IAccessControlUseCase

		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		postgresSQL             *postgres.PostgresSQL
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewAuthUseCase)

func NewAuthUseCase(
	userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository,
	rolePostgresqlRepository rolePostgresqlRepository.IRolePostgresqlRepository,
	authRedisRepository authRedisRepository.IAuthRedisRepository,
	accessControlUseCase accessControlUseCase.IAccessControlUseCase,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	postgresSQL *postgres.PostgresSQL,
	logger logger.IZapLogger,
) IAuthUseCase {
	return &authUseCase{
		userPostgresqlRepository: userPostgresqlRepository,
		rolePostgresqlRepository: rolePostgresqlRepository,
		authRedisRepository:      authRedisRepository,
		accessControlUseCase:     accessControlUseCase,
		rabbitmqInfrastructure:   rabbitmqInfrastructure,
		telemetryInfrastructure:  telemetryInfrastructure,
		postgresSQL:              postgresSQL,
		logger:                   logger,
	}
}
