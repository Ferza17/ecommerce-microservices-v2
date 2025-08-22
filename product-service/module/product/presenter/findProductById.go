package presenter

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *ProductPresenter) FindProductById(ctx context.Context, req *productRpc.FindProductByIdRequest) (*productRpc.Product, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductPresenter.FindProductById")
	defer span.End()
	if err := req.Validate(); err != nil {
		p.logger.Error("ProductPresenter.FindProductById", zap.String("requestID", pkgContext.GetRequestIDFromContext(ctx)), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res, err := p.productUseCase.FindProductById(ctx, pkgContext.GetRequestIDFromContext(ctx), req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
