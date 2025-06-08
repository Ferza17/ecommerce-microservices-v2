package presenter

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (p *paymentPresenter) FindPaymentByUserIdAndStatus(ctx context.Context, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "metadata not found")
	}
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.FindPaymentByUserIdAndStatus")
	defer span.End()

	requestId := ""
	if values := md.Get(enum.XRequestIDHeader.String()); len(values) > 0 {
		requestId = values[0]
	}

	// Call the use case's FindPaymentByUserIdAndStatus method
	payment, err := p.paymentUseCase.FindPaymentByUserIdAndStatus(ctx, requestId, request)
	if err != nil {
		// Log the error and return it
		p.logger.Error(fmt.Sprintf("Failed to find payment by user ID and status. RequestId: %s, Error: %v", requestId, err))
		return nil, err
	}

	return payment, nil

}
