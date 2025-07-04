package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
)

func (p *paymentProviderUseCase) FindPaymentProviders(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "ProviderUseCase.FindPaymentProviders")
	defer span.End()

	tx := p.postgresql.GormDB.Begin()
	// Call the repository's FindPaymentProviders method
	providers, err := p.paymentProviderRepository.FindPaymentProviders(ctx, requestId, request, tx)
	if err != nil {
		tx.Rollback()
		p.logger.Error(fmt.Sprintf("Failed to retrieve payment providers, requestId: %s, error: %v", requestId, err))
		return nil, err
	}

	tx.Commit()
	return &paymentRpc.FindPaymentProvidersResponse{
		Providers: orm.ProvidersToProto(providers),
	}, nil
}
