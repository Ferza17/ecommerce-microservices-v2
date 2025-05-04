package rpcclient

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	IRpcClient interface {
		Close() error
	}

	RpcClient struct {
		User    pb.UserServiceClient
		Product pb.ProductServiceClient

		productClient *grpc.ClientConn
		userClient    *grpc.ClientConn

		logger pkg.IZapLogger
	}
)

func NewRpcClient(logger pkg.IZapLogger) IRpcClient {

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

	return &RpcClient{
		Product:       pb.NewProductServiceClient(productClient),
		productClient: productClient,
		User:          pb.NewUserServiceClient(userClient),
		userClient:    userClient,
		logger:        logger,
	}
}

func (c *RpcClient) Close() error {
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
