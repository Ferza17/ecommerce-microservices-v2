package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	authRedisRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type (
	IAuthUseCase interface {
		UserLoginByEmailAndPassword(ctx context.Context, requestId string, req *pb.UserLoginByEmailAndPasswordRequest) (*pb.UserLoginByEmailAndPasswordResponse, error)
	}
	authUseCase struct {
		userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository
		authRedisRepository      authRedisRepository.IAuthRedisRepository
		rabbitmqInfrastructure   rabbitmqInfrastructure.IRabbitMQInfrastructure
		logger                   pkg.IZapLogger
	}
)

func NewAuthUseCase(
	userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository,
	authRedisRepository authRedisRepository.IAuthRedisRepository,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	logger pkg.IZapLogger) IAuthUseCase {
	return &authUseCase{
		userPostgresqlRepository: userPostgresqlRepository,
		authRedisRepository:      authRedisRepository,
		rabbitmqInfrastructure:   rabbitmqInfrastructure,
		logger:                   logger,
	}
}
