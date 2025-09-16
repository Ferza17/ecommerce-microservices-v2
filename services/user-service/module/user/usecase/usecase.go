package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	userPgSinkRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/kafkaSink"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"

	authRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	rolePostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgres"
)

type (
	IUserUseCase interface {
		FindUserById(ctx context.Context, requestId string, req *pb.FindUserByIdRequest) (*pb.FindUserByIdResponse, error)
		UpdateUserById(ctx context.Context, requestId string, req *pb.UpdateUserByIdRequest) (*pb.UpdateUserByIdResponse, error)
		FindUserByEmailAndPassword(ctx context.Context, requestId string, req *pb.FindUserByEmailAndPasswordRequest) (*pb.FindUserByEmailAndPasswordResponse, error)
		FindUserByEmail(ctx context.Context, requestId string, req *pb.FindUserByEmailRequest) (*pb.FindUserByEmailResponse, error)
	}

	userUseCase struct {
		userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository
		userPgSinkRepository     userPgSinkRepository.IUserKafkaSink
		rolePostgresqlRepository rolePostgresqlRepository.IRolePostgresqlRepository

		rabbitmqInfrastructure    rabbitmqInfrastructure.IRabbitMQInfrastructure
		postgresSQLInfrastructure postgres.IPostgresSQL
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		authRedisRepository       authRedisRepository.IAuthRedisRepository
		logger                    logger.IZapLogger
	}
)

var Set = wire.NewSet(NewUserUseCase)

func NewUserUseCase(
	userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository,
	userPgSinkRepository userPgSinkRepository.IUserKafkaSink,
	rolePostgresqlRepository rolePostgresqlRepository.IRolePostgresqlRepository,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	authRedisRepository authRedisRepository.IAuthRedisRepository,
	postgresSQLInfrastructure postgres.IPostgresSQL,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) IUserUseCase {
	return &userUseCase{
		userPostgresqlRepository:  userPostgresqlRepository,
		userPgSinkRepository:      userPgSinkRepository,
		rolePostgresqlRepository:  rolePostgresqlRepository,
		rabbitmqInfrastructure:    rabbitmqInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		authRedisRepository:       authRedisRepository,
		logger:                    logger,
	}
}
