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
	Service interface {
		Close() error
		GetProductService() pb.ProductServiceClient
		GetUserService() pb.UserServiceClient
	}

	service struct {
		user    pb.UserServiceClient
		product pb.ProductServiceClient

		// For Closing purpose
		productClient *grpc.ClientConn
		userClient    *grpc.ClientConn

		logger pkg.IZapLogger
	}
)

func NewRpcClient(logger pkg.IZapLogger) Service {

	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development

	productClient, err := grpc.NewClient(config.Get().ProductServiceURL, insecureCredentials)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a product client: %v", err))
		return nil
	}

	userClient, err := grpc.NewClient(config.Get().UserServiceURL, insecureCredentials)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a user client: %v", err))
		return nil
	}

	return &service{
		product:       pb.NewProductServiceClient(productClient),
		productClient: productClient,
		user:          pb.NewUserServiceClient(userClient),
		userClient:    userClient,
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
