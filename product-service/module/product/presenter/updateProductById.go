package presenter

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"github.com/golang/protobuf/ptypes/empty"
	"time"
)

func (p *ProductPresenter) UpdateProductById(ctx context.Context, req *productRpc.UpdateProductByIdRequest) (*empty.Empty, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()
	ctxTimeout, span := p.telemetryInfrastructure.Tracer(ctxTimeout, "Presenter.UpdateProductById")
	defer span.End()
	return nil, nil
}
