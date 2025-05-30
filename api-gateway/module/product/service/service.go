package service

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/product/v1"

	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	IProductService interface {
		FindProductById(ctx context.Context, requestId string, req *productRpc.FindProductByIdRequest) (*productRpc.Product, error)
		FindProductsWithPagination(ctx context.Context, requestId string, req *productRpc.FindProductsWithPaginationRequest) (*productRpc.FindProductsWithPaginationResponse, error)
		Close() error
	}

	productService struct {
		conn   *grpc.ClientConn
		svc    productRpc.ProductServiceClient
		cb     pkg.ICircuitBreaker
		logger pkg.IZapLogger
	}
)

func (s *productService) Close() error {
	if err := s.conn.Close(); err != nil {
		s.logger.Error(fmt.Sprintf("error while closing product service"))
		return err
	}
	return nil
}

func NewProductService(cb pkg.ICircuitBreaker, logger pkg.IZapLogger) IProductService {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development

	conn, err := grpc.NewClient(config.Get().ProductServiceURL, insecureCredentials)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a user client: %v", err))
		return nil
	}

	return &productService{
		conn:   conn,
		svc:    productRpc.NewProductServiceClient(conn),
		cb:     cb,
		logger: logger,
	}
}
