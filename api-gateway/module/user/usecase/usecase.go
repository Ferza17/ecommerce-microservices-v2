package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	userService "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/user/service"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IUserUseCase interface {
		CreateUser(ctx context.Context, requestId string, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
		UpdateUserById(ctx context.Context, requestId string, req *pb.UpdateUserByIdRequest) (*pb.UpdateUserByIdResponse, error)
	}
	UserUseCase struct {
		userService             userService.IUserService
		rabbitMQ                rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewUserUseCase(
	userService userService.IUserService,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) IUserUseCase {
	return &UserUseCase{
		userService:             userService,
		rabbitMQ:                rabbitMQ,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
