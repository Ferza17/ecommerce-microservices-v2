package presenter

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (p *ProductGrpcPresenter) FindProductsWithPagination(ctx context.Context, req *productRpc.FindProductsWithPaginationRequest) (*productRpc.FindProductsWithPaginationResponse, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()

	ctxTimeout, span := p.telemetryInfrastructure.Tracer(ctxTimeout, "Presenter.FindProductsWithPagination")
	defer span.End()

	md, ok := metadata.FromIncomingContext(ctxTimeout)
	if !ok {
		p.logger.Error(fmt.Sprintf("metadata not found"))
		return nil, status.Error(codes.InvalidArgument, "metadata not found")
	}

	requestID := ""
	if values := md.Get(enum.XRequestIDHeader.String()); len(values) > 0 {
		requestID = values[0]
	}

	resp, err := p.productUseCase.FindProductsWithPagination(ctxTimeout, requestID, req)
	if err != nil {
		p.logger.Error(fmt.Sprintf("error finding products: %v", err))
		return nil, err
	}

	return resp, nil
}
