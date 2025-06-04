package presenter

import (
	"context"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
)

func (p *paymentPresenter) FindPaymentById(ctx context.Context, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.Payment, error) {
	//TODO implement me
	panic("implement me")
}
