package presenter

import (
	"context"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"go.uber.org/zap"
)

func (p *paymentPresenter) CreatePayment(ctx context.Context, request *paymentRpc.CreatePaymentRequest) (*paymentRpc.CreatePaymentResponse, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentPresenter.CreatePayment")
	defer span.End()
	requestId := pkgContext.GetRequestIDFromContext(ctx)

	if err := request.Validate(); err != nil {
		p.logger.Error("PaymentPresenter.CreatePayment", zap.String("requestID", requestId), zap.Error(err))
		return nil, err
	}

	// Call the use case's CreatePayment method
	response, err := p.paymentUseCase.CreatePayment(ctx, requestId, request)
	if err != nil {
		p.logger.Error(fmt.Sprintf("Failed to create payment. RequestId: %s, Error: %v", requestId, err))
		return nil, err
	}

	return response, nil
}
