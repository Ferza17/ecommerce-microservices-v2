package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	rpcClientInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/service"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IUserUseCase interface {
		CreateUser(ctx context.Context, requestId string, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	}
	UserUseCase struct {
		rpcClient rpcClientInfrastructure.IService
		rabbitMQ  rabbitmqInfrastructure.IRabbitMQInfrastructure
		logger    pkg.IZapLogger
	}
)

func NewUserUseCase(
	rpcClient rpcClientInfrastructure.IService,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	logger pkg.IZapLogger,
) IUserUseCase {
	return &UserUseCase{
		rpcClient: rpcClient,
		rabbitMQ:  rabbitMQ,
		logger:    logger,
	}
}
