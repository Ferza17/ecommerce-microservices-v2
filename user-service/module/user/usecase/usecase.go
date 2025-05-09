package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type (
	IUserUseCase interface {
		CreateUser(ctx context.Context, requestId string, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
		FindUserById(ctx context.Context, requestId string, req *pb.FindUserByIdRequest) (*pb.User, error)
		FindUserByEmailAndPassword(ctx context.Context, requestId string, req *pb.FindUserByEmailAndPasswordRequest) (*pb.User, error)
		UpdateUserById(ctx context.Context, requestId string, req *pb.UpdateUserByIdRequest) (*pb.UpdateUserByIdResponse, error)
	}

	userUseCase struct {
		userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository
		rabbitmqInfrastructure   rabbitmqInfrastructure.IRabbitMQInfrastructure
		logger                   pkg.IZapLogger
	}
)

func NewUserUseCase(userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository, rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure, logger pkg.IZapLogger) IUserUseCase {
	return &userUseCase{
		userPostgresqlRepository: userPostgresqlRepository,
		rabbitmqInfrastructure:   rabbitmqInfrastructure,
		logger:                   logger,
	}
}
