package usecase

import (
	"context"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/payment/v1"
)

func (u *paymentUseCase) FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error) {
	// Tracing the method execution
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "PaymentUseCase.FindPaymentByUserIdAndStatus")
	defer span.End()

	// Using the payment service to find payment by UserId and Status
	payment, err := u.paymentSvc.FindPaymentByUserIdAndStatus(ctx, requestId, request)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding payment by user id and status: %v", err))
		return nil, err
	}

	return payment, nil
}
