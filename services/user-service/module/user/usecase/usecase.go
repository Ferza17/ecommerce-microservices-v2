package usecase

import (
	"context"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
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
		userPostgresqlRepository  userPostgresqlRepository.IUserPostgresqlRepository
		rolePostgresqlRepository  rolePostgresqlRepository.IRolePostgresqlRepository
		kafkaInfrastructure       kafkaInfrastructure.IKafkaInfrastructure
		postgresSQLInfrastructure postgres.IPostgresSQL
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		authRedisRepository       authRedisRepository.IAuthRedisRepository
		logger                    logger.IZapLogger
	}
)

var Set = wire.NewSet(NewUserUseCase)

func NewUserUseCase(
	userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository,
	rolePostgresqlRepository rolePostgresqlRepository.IRolePostgresqlRepository,
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	authRedisRepository authRedisRepository.IAuthRedisRepository,
	postgresSQLInfrastructure postgres.IPostgresSQL,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) IUserUseCase {
	return &userUseCase{
		userPostgresqlRepository:  userPostgresqlRepository,
		rolePostgresqlRepository:  rolePostgresqlRepository,
		kafkaInfrastructure:       kafkaInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		authRedisRepository:       authRedisRepository,
		logger:                    logger,
	}
}
