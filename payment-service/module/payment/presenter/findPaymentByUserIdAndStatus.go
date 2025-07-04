package presenter

import (
	"context"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"go.uber.org/zap"
)

func (p *paymentPresenter) FindPaymentByUserIdAndStatus(ctx context.Context, req *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentPresenter.FindPaymentByUserIdAndStatus")
	defer span.End()
	requestId := pkgContext.GetRequestIDFromContext(ctx)

	if err := p.userService.AuthUserVerifyAccessControl(ctx, requestId); err != nil {
		p.logger.Error("PaymentPresenter.FindPaymentByUserIdAndStatus", zap.String("requestID", requestId), zap.Error(err))
		return nil, err
	}

	if err := req.Validate(); err != nil {
		p.logger.Error("PaymentPresenter.FindPaymentByUserIdAndStatus", zap.String("requestID", requestId), zap.Error(err))
		return nil, err
	}

	// Call the use case's FindPaymentByUserIdAndStatus method
	payment, err := p.paymentUseCase.FindPaymentByUserIdAndStatus(ctx, requestId, req)
	if err != nil {
		p.logger.Error(fmt.Sprintf("Failed to find payment by user ID and status. RequestId: %s, Error: %v", requestId, err))
		return nil, err
	}

	return payment, nil

}
