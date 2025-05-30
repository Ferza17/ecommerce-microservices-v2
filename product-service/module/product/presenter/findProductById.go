package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (p *ProductGrpcPresenter) FindProductById(ctx context.Context, req *productRpc.FindProductByIdRequest) (*productRpc.Product, error) {
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.FindProductById")
	defer span.End()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "metadata not found")
	}
	requestID := ""
	if values := md.Get(enum.XRequestIDHeader.String()); len(values) > 0 {
		requestID = values[0]
	}

	res, err := p.productUseCase.FindProductById(ctx, requestID, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
