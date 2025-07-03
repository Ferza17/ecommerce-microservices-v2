package presenter

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (p *paymentPresenter) FindPaymentById(ctx context.Context, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.Payment, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "metadata not found")
	}
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.FindPaymentById")
	defer span.End()

	requestID := ""
	if values := md.Get(enum.XRequestIDHeader.String()); len(values) > 0 {
		requestID = values[0]
	}

	// Call the use case's FindPaymentById method
	payment, err := p.paymentUseCase.FindPaymentById(ctx, requestID, request)
	if err != nil {
		// Log the error and return it
		p.logger.Error(fmt.Sprintf("Failed to find payment. RequestId: %s, Error: %v", requestID, err))
		return nil, err
	}

	return payment, nil
}
