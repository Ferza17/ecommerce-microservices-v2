package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	accessControlPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/postgres"
	accessControlRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/repository/redis"
	rolePostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	"github.com/google/wire"

	authRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
)

type (
	IAuthUseCase interface {
		AuthUserRegister(ctx context.Context, requestId string, req *pb.AuthUserRegisterRequest) (*pb.AuthUserRegisterResponse, error)
		AuthUserFindUserByToken(ctx context.Context, requestId string, req *pb.AuthUserFindUserByTokenRequest) (*pb.AuthUserFindUserByTokenResponse, error)
		AuthUserLoginByEmailAndPassword(ctx context.Context, requestId string, req *pb.AuthUserLoginByEmailAndPasswordRequest) (*pb.AuthUserLoginByEmailAndPasswordResponse, error)
		AuthUserVerifyOtp(ctx context.Context, requestId string, req *pb.AuthUserVerifyOtpRequest) (*pb.AuthUserVerifyOtpResponse, error)
		AuthUserLogoutByToken(ctx context.Context, requestId string, req *pb.AuthUserLogoutByTokenRequest) (*pb.AuthUserLogoutByTokenResponse, error)
		AuthUserVerifyAccessControl(ctx context.Context, requestId string, req *pb.AuthUserVerifyAccessControlRequest) (*pb.AuthUserVerifyAccessControlResponse, error)
	}
	authUseCase struct {
		userPostgresqlRepository          userPostgresqlRepository.IUserPostgresqlRepository
		rolePostgresqlRepository          rolePostgresqlRepository.IRolePostgresqlRepository
		accessControlPostgresqlRepository accessControlPostgresqlRepository.IAccessControlPostgresqlRepository

		authRedisRepository          authRedisRepository.IAuthRedisRepository
		accessControlRedisRepository accessControlRedisRepository.IAccessControlRedisRepository

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
	accessControlPostgresqlRepository accessControlPostgresqlRepository.IAccessControlPostgresqlRepository,
	authRedisRepository authRedisRepository.IAuthRedisRepository,
	accessControlRedisRepository accessControlRedisRepository.IAccessControlRedisRepository,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	postgresSQL *postgres.PostgresSQL,
	logger logger.IZapLogger,
) IAuthUseCase {
	return &authUseCase{
		userPostgresqlRepository:          userPostgresqlRepository,
		rolePostgresqlRepository:          rolePostgresqlRepository,
		accessControlPostgresqlRepository: accessControlPostgresqlRepository,
		authRedisRepository:               authRedisRepository,
		accessControlRedisRepository:      accessControlRedisRepository,
		rabbitmqInfrastructure:            rabbitmqInfrastructure,
		telemetryInfrastructure:           telemetryInfrastructure,
		postgresSQL:                       postgresSQL,
		logger:                            logger,
	}
}
