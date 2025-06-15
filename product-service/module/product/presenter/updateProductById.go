package presenter

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"
	"time"
)

func (p *ProductGrpcPresenter) UpdateProductById(ctx context.Context, req *productRpc.UpdateProductByIdRequest) (*productRpc.Product, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()
	ctxTimeout, span := p.telemetryInfrastructure.Tracer(ctxTimeout, "Presenter.UpdateProductById")
	defer span.End()
	return nil, nil
}
