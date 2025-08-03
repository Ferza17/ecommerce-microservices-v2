package presenter

import (
	"context"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"go.uber.org/zap"
)

func (p *paymentPresenter) FindPaymentById(ctx context.Context, req *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.FindPaymentByIdResponse, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentPresenter.FindPaymentById")
	defer span.End()
	requestId := pkgContext.GetRequestIDFromContext(ctx)

	if err := p.userService.AuthUserVerifyAccessControl(ctx, requestId); err != nil {
		p.logger.Error("PaymentPresenter.FindPaymentById", zap.String("requestID", requestId), zap.Error(err))
		return nil, err
	}

	if err := req.Validate(); err != nil {
		p.logger.Error("PaymentPresenter.FindPaymentById", zap.String("requestID", requestId), zap.Error(err))
		return nil, err
	}

	// Call the use case's FindPaymentById method
	payment, err := p.paymentUseCase.FindPaymentById(ctx, requestId, req)
	if err != nil {
		p.logger.Error(fmt.Sprintf("Failed to find payment. RequestId: %s, Error: %v", requestId, err))
		return nil, err
	}

	return payment, nil
}
