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

func (p *paymentProviderPresenter) FindPaymentProviders(ctx context.Context, request *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error) {
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.FindPaymentProviders")
	defer span.End()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		p.logger.Error(fmt.Sprintf("metadata not found"))
		return nil, status.Error(codes.InvalidArgument, "metadata not found")
	}

	requestID := ""
	if values := md.Get(enum.XRequestIDHeader.String()); len(values) > 0 {
		requestID = values[0]
	}

	// Call the use case's FindPaymentProviders method
	response, err := p.paymentProviderUseCase.FindPaymentProviders(ctx, requestID, request)
	if err != nil {
		// Log the error with the requestID
		p.logger.Error(fmt.Sprintf("Failed to find payment providers. RequestID: %s, Error: %v", requestID, err))
		return nil, err
	}

	// Log success
	return response, nil
}
