package usecase

import (
	"context"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/payment/v1"
)

func (u *paymentUseCase) FindPaymentById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.Payment, error) {
	// Tracing the method execution
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "PaymentUseCase.FindPaymentById")
	defer span.End()

	// Using the payment service to find payment by ID
	payment, err := u.paymentSvc.FindPaymentById(ctx, requestId, request)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding payment by ID: %v", err))
		return nil, err
	}

	return payment, nil
}
