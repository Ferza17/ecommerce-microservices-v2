package presenter

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"
)

func (p *ProductGrpcPresenter) DeleteProductById(ctx context.Context, req *productRpc.DeleteProductByIdRequest) (*productRpc.DeleteProductByIdResponse, error) {
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.DeleteProductById")
	defer span.End()
	return nil, nil
}
