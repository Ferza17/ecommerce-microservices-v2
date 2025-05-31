package service

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	commerceRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/commerce/v1"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	ICommerceCartService interface {
		FindCartItemById(ctx context.Context, requestId string, req *commerceRpc.FindCartItemByIdRequest) (*commerceRpc.CartItem, error)
		FindCartItemsWithPagination(ctx context.Context, requestId string, req *commerceRpc.FindCartItemsWithPaginationRequest) (*commerceRpc.FindCartItemsWithPaginationResponse, error)
		Close() error
	}
	commerceCartService struct {
		conn   *grpc.ClientConn
		svc    commerceRpc.CartServiceClient
		cb     pkg.ICircuitBreaker
		logger pkg.IZapLogger
	}
)

func (s *commerceCartService) Close() error {
	if err := s.conn.Close(); err != nil {
		s.logger.Error(fmt.Sprintf("error while closing product service"))
		return err
	}
	return nil
}

func NewCommerceCartService(cb pkg.ICircuitBreaker, logger pkg.IZapLogger) ICommerceCartService {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development

	conn, err := grpc.NewClient(config.Get().CommerceServiceURL, insecureCredentials)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a user client: %v", err))
		return nil
	}

	return &commerceCartService{
		conn:   conn,
		svc:    commerceRpc.NewCartServiceClient(conn),
		cb:     cb,
		logger: logger,
	}
}
