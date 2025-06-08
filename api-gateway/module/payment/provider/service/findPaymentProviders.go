package service

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/payment/v1"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/util"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/metadata"
)

func (p *paymentProviderService) FindPaymentProviders(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error) {
	// Execute functionality within the circuit breaker
	resp, err := p.cb.Execute(func() (interface{}, error) {
		md := metadata.New(map[string]string{enum.XRequestIDHeader.String(): requestId})
		otel.GetTextMapPropagator().Inject(ctx, &util.MetadataHeaderCarrier{md})
		ctx = metadata.NewOutgoingContext(ctx, md)

		return p.svc.FindPaymentProviders(ctx, request)
	})

	if err != nil {
		p.logger.Error(fmt.Sprintf("Error while finding payment providers: %v | RequestID: %s", err, requestId))
		return nil, err
	}
	return resp.(*paymentRpc.FindPaymentProvidersResponse), nil
}
