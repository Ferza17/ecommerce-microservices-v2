package usecase

import (
	"context"
	"errors"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (p *paymentProviderUseCase) FindPaymentProviderById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProviderByIdRequest) (*paymentRpc.Provider, error) {
	// Fetch the provider by ID from the repository
	provider, err := p.paymentProviderRepository.FindPaymentProviderById(ctx, requestId, request.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			p.logger.Error(fmt.Sprintf("Provider not found for ID: %s", request.GetId()))
			return nil, status.Error(codes.NotFound, "provider not found")
		}

		// Log the error and return a meaningful message
		p.logger.Error(fmt.Sprintf("Error fetching provider by ID: %s | Error: %v", request.GetId(), err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	// Handle the case when no provider is found
	if provider == nil {
		p.logger.Error(fmt.Sprintf("Provider not found for ID: %s", request.GetId()))
		return nil, status.Error(codes.NotFound, "provider not found")
	}

	// Return the response
	return provider.ToProto(), nil
}
