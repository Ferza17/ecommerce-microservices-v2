package presenter

import (
	"context"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
)

func (p *paymentProviderPresenter) FindPaymentProviders(ctx context.Context, request *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error) {
	//TODO implement me
	panic("implement me")
}
