package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (p *ProductGrpcPresenter) CreateProduct(ctx context.Context, req *productRpc.CreateProductRequest) (*productRpc.CreateProductResponse, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()

	ctxTimeout, span := p.telemetryInfrastructure.Tracer(ctxTimeout, "Presenter.CreateProduct")
	defer span.End()
	md, ok := metadata.FromIncomingContext(ctxTimeout)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "metadata not found")
	}
	requestID := ""
	if values := md.Get(enum.XRequestIDHeader.String()); len(values) > 0 {
		requestID = values[0]
	}

	res, err := p.productUseCase.CreateProduct(ctxTimeout, requestID, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
