package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	authService "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/service"
	userService "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/user/service"

	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

type (
	IAuthUseCase interface {
		UserLoginByEmailAndPassword(ctx context.Context, requestId string, input *pb.UserLoginByEmailAndPasswordRequest) (*pb.UserLoginByEmailAndPasswordResponse, error)
		UserLogoutByToken(ctx context.Context, requestId string, input *pb.UserLogoutByTokenRequest) (*pb.UserLogoutByTokenResponse, error)
	}
	authUseCase struct {
		authService             authService.IAuthService
		userService             userService.IUserService
		rabbitMQ                rabbitmqInfrastructure.IRabbitMQInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewAuthUseCase(
	authService authService.IAuthService,
	userService userService.IUserService,
	rabbitMQ rabbitmqInfrastructure.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) IAuthUseCase {
	return &authUseCase{
		authService:             authService,
		userService:             userService,
		rabbitMQ:                rabbitMQ,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
