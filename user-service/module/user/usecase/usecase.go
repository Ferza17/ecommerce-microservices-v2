package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"

	authRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	rolePostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgres"
)

type (
	IUserUseCase interface {
		FindUserById(ctx context.Context, requestId string, req *userRpc.FindUserByIdRequest) (*userRpc.User, error)
		UpdateUserById(ctx context.Context, requestId string, req *userRpc.UpdateUserByIdRequest) (*userRpc.UpdateUserByIdResponse, error)
		FindUserByEmailAndPassword(context.Context, string, *userRpc.FindUserByEmailAndPasswordRequest) (*userRpc.User, error)
	}

	userUseCase struct {
		userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository
		rolePostgresqlRepository rolePostgresqlRepository.IRolePostgresqlRepository

		rabbitmqInfrastructure    rabbitmqInfrastructure.IRabbitMQInfrastructure
		postgresSQLInfrastructure *postgres.PostgresSQL
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		authRedisRepository       authRedisRepository.IAuthRedisRepository
		logger                    logger.IZapLogger
	}
)

var Set = wire.NewSet(NewUserUseCase)

func NewUserUseCase(
	userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository,
	rolePostgresqlRepository rolePostgresqlRepository.IRolePostgresqlRepository,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	authRedisRepository authRedisRepository.IAuthRedisRepository,
	postgresSQLInfrastructure *postgres.PostgresSQL,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) IUserUseCase {
	return &userUseCase{
		userPostgresqlRepository:  userPostgresqlRepository,
		rolePostgresqlRepository:  rolePostgresqlRepository,
		rabbitmqInfrastructure:    rabbitmqInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		authRedisRepository:       authRedisRepository,
		logger:                    logger,
	}
}
