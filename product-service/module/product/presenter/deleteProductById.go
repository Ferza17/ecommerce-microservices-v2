package presenter

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"
	"time"
)

func (p *ProductGrpcPresenter) DeleteProductById(ctx context.Context, req *productRpc.DeleteProductByIdRequest) (*productRpc.DeleteProductByIdResponse, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()
	ctxTimeout, span := p.telemetryInfrastructure.Tracer(ctxTimeout, "Presenter.DeleteProductById")
	defer span.End()
	return nil, nil
}
