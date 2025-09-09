package presenter

import (
	"context"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"go.uber.org/zap"
)

func (p *paymentProviderPresenter) FindPaymentProviders(ctx context.Context, req *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentPresenter.FindPaymentProviders")
	defer span.End()
	requestId := pkgContext.GetRequestIDFromContext(ctx)

	if err := req.Validate(); err != nil {
		p.logger.Error("PaymentPresenter.FindPaymentProviders", zap.String("requestID", requestId), zap.Error(err))
		return nil, err
	}

	// Call the use case's FindPaymentProviders method
	response, err := p.paymentProviderUseCase.FindPaymentProviders(ctx, requestId, req)
	if err != nil {
		// Log the error with the requestID
		p.logger.Error(fmt.Sprintf("Failed to find payment providers. RequestID: %s, Error: %v", requestId, err))
		return nil, err
	}

	// Log success
	return response, nil
}
