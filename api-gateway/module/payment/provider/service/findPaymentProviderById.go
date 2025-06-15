package service

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/payment/v1"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/util"
	"github.com/sony/gobreaker"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (p *paymentProviderService) FindPaymentProviderById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProviderByIdRequest) (*paymentRpc.Provider, error) {
	// Execute functionality within the circuit breaker
	resp, err := p.cb.Execute(func() (interface{}, error) {
		md := metadata.New(map[string]string{enum.XRequestIDHeader.String(): requestId})
		otel.GetTextMapPropagator().Inject(ctx, &util.MetadataHeaderCarrier{md})
		ctx = metadata.NewOutgoingContext(ctx, md)

		resp, err := p.svc.FindPaymentProviderById(ctx, request)
		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				if st.Code() == codes.NotFound {
					return nil, gobreaker.ErrOpenState
				}
			}
			return nil, err
		}
		return resp, nil
	})

	if err != nil {
		p.logger.Error(fmt.Sprintf("Error while finding payment providers: %v | RequestID: %s", err, requestId))
		return nil, err
	}
	return resp.(*paymentRpc.Provider), nil
}
