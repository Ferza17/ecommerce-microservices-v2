package service

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	IService interface {
		Close() error
		GetProductService() pb.ProductServiceClient
		GetUserService() pb.UserServiceClient
		GetAuthService() pb.AuthServiceClient
	}

	service struct {
		user    pb.UserServiceClient
		auth    pb.AuthServiceClient
		product pb.ProductServiceClient

		// For Closing purpose
		productClient *grpc.ClientConn
		authClient    *grpc.ClientConn
		userClient    *grpc.ClientConn

		logger pkg.IZapLogger
	}
)

func (c *service) GetAuthService() pb.AuthServiceClient {
	return c.auth
}

func NewRpcClient(logger pkg.IZapLogger) IService {

	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development

	productServiceUrl := config.Get().ProductServiceURL
	productClient, err := grpc.NewClient(productServiceUrl, insecureCredentials)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a product client: %v", err))
		return nil
	}

	userAndAuthClient, err := grpc.NewClient(config.Get().UserServiceURL, insecureCredentials)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a user client: %v", err))
		return nil
	}

	return &service{
		product:       pb.NewProductServiceClient(productClient),
		productClient: productClient,
		user:          pb.NewUserServiceClient(userAndAuthClient),
		userClient:    userAndAuthClient,
		auth:          pb.NewAuthServiceClient(userAndAuthClient),
		logger:        logger,
	}
}

func (c *service) Close() error {
	if err := c.productClient.Close(); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close a product client: %v", err))
		return err
	}

	if err := c.userClient.Close(); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close a user client: %v", err))
		return err
	}

	return nil
}

func (c *service) GetProductService() pb.ProductServiceClient {
	return c.product
}

func (c *service) GetUserService() pb.UserServiceClient {
	return c.user
}
