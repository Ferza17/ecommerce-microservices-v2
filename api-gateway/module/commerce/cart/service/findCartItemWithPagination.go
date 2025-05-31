package service

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	commerceRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/commerce/v1"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/util"
	"github.com/sony/gobreaker"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *commerceCartService) FindCartItemsWithPagination(ctx context.Context, requestId string, req *commerceRpc.FindCartItemsWithPaginationRequest) (*commerceRpc.FindCartItemsWithPaginationResponse, error) {
	result, err := s.cb.Execute(func() (interface{}, error) {
		md := metadata.New(map[string]string{enum.XRequestIDHeader.String(): requestId})
		otel.GetTextMapPropagator().Inject(ctx, &util.MetadataHeaderCarrier{md})
		ctx = metadata.NewOutgoingContext(ctx, md)

		resp, err := s.svc.FindCartItemsWithPagination(ctx, req)
		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				if st.Code() == codes.NotFound {
					return nil, gobreaker.ErrOpenState
				}
			}
			return nil, err
		}
		return resp, err
	})

	if err != nil {
		s.logger.Info(fmt.Sprintf("Error Breaker %v", err))
		return nil, err
	}

	return result.(*commerceRpc.FindCartItemsWithPaginationResponse), nil
}
