package usecase

import (
	"context"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/payment/v1"
)

func (u *paymentProviderUseCase) FindPaymentProviders(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error) {
	// Tracing for telemetry
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "ProviderUseCase.FindPaymentProviders")
	defer span.End()

	// Call the paymentProviderSvc to fetch payment providers
	response, err := u.paymentProviderSvc.FindPaymentProviders(ctx, requestId, request)
	if err != nil {
		u.logger.Error(fmt.Sprintf("Error while retrieving payment providers: %v", err))
		return nil, err
	}

	return response, nil
}
