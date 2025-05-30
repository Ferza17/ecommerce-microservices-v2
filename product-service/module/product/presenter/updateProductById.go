package presenter

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"
)

func (p *ProductGrpcPresenter) UpdateProductById(ctx context.Context, req *productRpc.UpdateProductByIdRequest) (*productRpc.Product, error) {
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.UpdateProductById")
	defer span.End()
	return nil, nil
}
