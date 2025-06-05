package usecase

import (
	"context"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
)

func (u *paymentUseCase) CreatePayment(ctx context.Context, requestId string, request *paymentRpc.CreatePaymentRequest) error {
	//TODO implement me
	panic("implement me")
}
