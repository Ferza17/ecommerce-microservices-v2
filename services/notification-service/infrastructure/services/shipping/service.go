package shipping

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/shipping"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	IShippingService interface {
		// Shipping Provider
		GetShippingProviderById(ctx context.Context, requestId string, request *pb.GetShippingProviderByIdRequest) (*pb.GetShippingProviderByIdResponse, error)
	}

	shippingService struct {
		logger              logger.IZapLogger
		shippingSvc         pb.ShippingServiceClient
		shippingProviderSvc pb.ShippingProviderServiceClient
	}
)

var Set = wire.NewSet(NewShippingService)

func NewShippingService(logger logger.IZapLogger) IShippingService {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", config.Get().ConfigServiceShipping.RpcHost, config.Get().ConfigServiceShipping.RpcPort), insecureCredentials)
	if err != nil {
		logger.Error("ShippingService.NewShippingService", zap.Error(err))
		return nil
	}
	return &shippingService{
		logger:              logger,
		shippingSvc:         pb.NewShippingServiceClient(conn),
		shippingProviderSvc: pb.NewShippingProviderServiceClient(conn),
	}
}
