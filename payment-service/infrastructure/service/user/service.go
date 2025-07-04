package user

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	IUserService interface {
		AuthServiceVerifyIsExcluded(ctx context.Context, requestId string, in *pb.AuthServiceVerifyIsExcludedRequest) (*pb.AuthServiceVerifyIsExcludedResponse, error)
		AuthUserVerifyAccessControl(ctx context.Context, requestId string) error
	}

	userService struct {
		logger  logger.IZapLogger
		authSvc pb.AuthServiceClient
		userSvc pb.UserServiceClient
	}
)

var Set = wire.NewSet(NewUserService)

func NewUserService(logger logger.IZapLogger) IUserService {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", config.Get().UserServiceRpcHost, config.Get().UserServiceRpcPort), insecureCredentials)
	if err != nil {
		logger.Error("UserService.NewUserService", zap.Error(err))
		return nil
	}

	return &userService{
		logger:  logger,
		authSvc: pb.NewAuthServiceClient(conn),
		userSvc: pb.NewUserServiceClient(conn),
	}
}
