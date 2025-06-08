package service

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/payment/v1"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	IPaymentService interface {
		FindPaymentById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.Payment, error)
		FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error)
		Close() error
	}
	paymentService struct {
		conn   *grpc.ClientConn
		svc    paymentRpc.PaymentServiceClient
		cb     pkg.ICircuitBreaker
		logger pkg.IZapLogger
	}
)

func (s *paymentService) Close() error {
	if err := s.conn.Close(); err != nil {
		s.logger.Error(fmt.Sprintf("error while closing payment service: %v", err))
		return err
	}
	return nil
}

func NewPaymentService(cb pkg.ICircuitBreaker, logger pkg.IZapLogger) IPaymentService {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development

	conn, err := grpc.Dial(config.Get().PaymentServiceURL, insecureCredentials)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create payment service client: %v", err))
		return nil
	}

	return &paymentService{
		conn:   conn,
		svc:    paymentRpc.NewPaymentServiceClient(conn),
		cb:     cb,
		logger: logger,
	}
}
