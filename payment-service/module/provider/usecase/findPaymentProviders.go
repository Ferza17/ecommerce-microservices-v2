package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		var (
			pr          paymentRpc.ProviderMethod
			discardedAt *timestamppb.Timestamp
		)

		if pr, err = enum.ProviderMethodToProto(provider.Method); err != nil {
			p.logger.Error(fmt.Sprintf("Failed to convert provider method: %v", err))
			return nil, err
		}

		if provider.DiscardedAt != nil {
			discardedAt = timestamppb.New(*provider.DiscardedAt)
		}

		providerResponses = append(providerResponses, &paymentRpc.Provider{
			Id:          provider.ID,
			Name:        provider.Name,
			Method:      pr,
			CreatedAt:   timestamppb.New(provider.CreatedAt),
			UpdatedAt:   timestamppb.New(provider.UpdatedAt),
			DiscardedAt: discardedAt,
		})
	}

	// Create a gRPC response containing the providers
	response := &paymentRpc.FindPaymentProvidersResponse{
		Data: providerResponses,
	}

	// Log success
	return response, nil
}
