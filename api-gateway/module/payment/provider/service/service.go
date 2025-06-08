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
	IPaymentProviderService interface {
		FindPaymentProviders(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error)
		Close() error
	}

	paymentProviderService struct {
		conn   *grpc.ClientConn
		svc    paymentRpc.PaymentProviderServiceClient
		cb     pkg.ICircuitBreaker
		logger pkg.IZapLogger
	}
)

func NewPaymentProviderService(cb pkg.ICircuitBreaker, logger pkg.IZapLogger) IPaymentProviderService {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials()) // For Local Development

	conn, err := grpc.Dial(config.Get().PaymentServiceURL, insecureCredentials)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create payment provider service client: %v", err))
		return nil
	}

	return &paymentProviderService{
		conn:   conn,
		svc:    paymentRpc.NewPaymentProviderServiceClient(conn),
		cb:     cb,
		logger: logger,
	}
}

func (p *paymentProviderService) Close() error {
	if err := p.conn.Close(); err != nil {
		p.logger.Error(fmt.Sprintf("Error while closing payment provider service connection: %v", err))
		return err
	}
	return nil
}
