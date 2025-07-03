package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
)

func (p *paymentProviderUseCase) FindPaymentProviders(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error) {
	// Start tracing the use case
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "UseCase.FindPaymentProviders")
	defer span.End()

	// Call the repository's FindPaymentProviders method
	providers, err := p.paymentProviderRepository.FindPaymentProviders(ctx, requestId, request)
	if err != nil {
		// Log the error and return it
		p.logger.Error(fmt.Sprintf("Failed to retrieve payment providers, requestId: %s, error: %v", requestId, err))
		return nil, err
	}

	// Log success
	return &paymentRpc.FindPaymentProvidersResponse{
		Providers: orm.ProvidersToProto(providers),
	}, nil
}
