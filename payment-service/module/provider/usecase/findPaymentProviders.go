package usecase

import (
	"context"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
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

	// Convert the list of ORM providers to gRPC-compatible responses
	var providerResponses []*paymentRpc.Provider
	for _, provider := range providers {
		providerResponses = append(providerResponses, &paymentRpc.Provider{
			Id:   provider.ID,
			Name: provider.Name,
		})
	}

	// Create a gRPC response containing the providers
	response := &paymentRpc.FindPaymentProvidersResponse{
		Data: providerResponses,
	}

	// Log success
	return response, nil
}
