package service

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/user/v1"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	IAuthService interface {
		FindUserByToken(ctx context.Context, requestId string, req *userRpc.FindUserByTokenRequest) (*userRpc.User, error)
		UserVerifyOtp(ctx context.Context, requestId string, req *userRpc.UserVerifyOtpRequest) (*userRpc.UserVerifyOtpResponse, error)
		Close() error
	}

	authService struct {
		conn   *grpc.ClientConn
		svc    userRpc.AuthServiceClient
		cb     pkg.ICircuitBreaker
		logger pkg.IZapLogger
	}
)

func (s *authService) Close() error {
	if err := s.conn.Close(); err != nil {
		s.logger.Error(fmt.Sprintf("error while closing product service"))
		return err
	}
	return nil
}

func NewAuthService(cb pkg.ICircuitBreaker, logger pkg.IZapLogger) IAuthService {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development

	conn, err := grpc.NewClient(config.Get().UserServiceURL, insecureCredentials)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a user client: %v", err))
		return nil
	}

	return &authService{
		conn:   conn,
		svc:    userRpc.NewAuthServiceClient(conn),
		cb:     cb,
		logger: logger,
	}
}
