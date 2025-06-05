package presenter

import (
	"context"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
)

func (p *paymentPresenter) FindPaymentByUserIdAndStatus(ctx context.Context, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error) {
	return nil, nil
}
