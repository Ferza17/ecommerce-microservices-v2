package product

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/product"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	IProductService interface {
		FindProductsWithPagination(ctx context.Context, requestId string, in *pb.FindProductsWithPaginationRequest) (*pb.FindProductsWithPaginationResponse, error)
	}

	productService struct {
		logger     logger.IZapLogger
		productSvc pb.ProductServiceClient
	}
)

var Set = wire.NewSet(NewProductService)

func NewProductService(logger logger.IZapLogger) IProductService {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", config.Get().ProductServiceRpcHost, config.Get().ProductServiceRpcPort), insecureCredentials)
	if err != nil {
		logger.Error("ProductService.NewProductService", zap.Error(err))
		return nil
	}
	return &productService{
		logger:     logger,
		productSvc: pb.NewProductServiceClient(conn),
	}
}
