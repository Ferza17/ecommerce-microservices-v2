package usecase

import (
	"context"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"

	authRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	eventMongoDBRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/event/repository/mongodb"
	eventUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/event/usecase"
	rolePostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IUserUseCase interface {
		// COMMAND
		CreateUser(ctx context.Context, requestId string, req *pb.AuthUserRegisterRequest) (*pb.AuthUserRegisterResponse, error)
		ConfirmCreateUser(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error
		CompensateCreateUser(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error

		UpdateUserById(ctx context.Context, requestId string, req *pb.UpdateUserByIdRequest) (*pb.UpdateUserByIdResponse, error)

		// QUERY
		FindUserById(ctx context.Context, requestId string, req *pb.FindUserByIdRequest) (*pb.FindUserByIdResponse, error)
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
		eventUseCase              eventUseCase.IEventUseCase
		eventMongoDBRepository    eventMongoDBRepository.IEventMongoRepository
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
	eventUseCase eventUseCase.IEventUseCase,
	eventMongoDBRepository eventMongoDBRepository.IEventMongoRepository,
	logger logger.IZapLogger) IUserUseCase {
	return &userUseCase{
		userPostgresqlRepository:  userPostgresqlRepository,
		rolePostgresqlRepository:  rolePostgresqlRepository,
		kafkaInfrastructure:       kafkaInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		authRedisRepository:       authRedisRepository,
		eventUseCase:              eventUseCase,
		eventMongoDBRepository:    eventMongoDBRepository,
		logger:                    logger,
	}
}
