package service

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	IUserService interface {
		FindUserByEmailAndPassword(ctx context.Context, requestId string, req *pb.FindUserByEmailAndPasswordRequest) (*pb.User, error)
		Close() error
	}

	userService struct {
		conn   *grpc.ClientConn
		svc    pb.UserServiceClient
		cb     pkg.ICircuitBreaker
		logger pkg.IZapLogger
	}
)

func (s *userService) Close() error {
	if err := s.conn.Close(); err != nil {
		s.logger.Error(fmt.Sprintf("error while closing product service"))
		return err
	}
	return nil
}

func NewUserService(cb pkg.ICircuitBreaker, logger pkg.IZapLogger) IUserService {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development

	conn, err := grpc.NewClient(config.Get().UserServiceURL, insecureCredentials)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a user client: %v", err))
		return nil
	}

	return &userService{
		conn:   conn,
		svc:    pb.NewUserServiceClient(conn),
		cb:     cb,
		logger: logger,
	}
}
