package payment

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/payment"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	IPaymentService interface {
		FindPaymentById(ctx context.Context, requestId string, req *pb.FindPaymentByIdRequest) (*pb.FindPaymentByIdResponse, error)
	}

	paymentService struct {
		logger     logger.IZapLogger
		paymentSvc pb.PaymentServiceClient
	}
)

var Set = wire.NewSet(NewPaymentService)

func NewPaymentService(
	logger logger.IZapLogger,
) IPaymentService {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", config.Get().PaymentServiceRpcHost, config.Get().PaymentServiceRpcPort), insecureCredentials)
	if err != nil {
		logger.Error("ShippingService.NewShippingService", zap.Error(err))
		return nil
	}
	return &paymentService{
		logger:     logger,
		paymentSvc: pb.NewPaymentServiceClient(conn),
	}
}
