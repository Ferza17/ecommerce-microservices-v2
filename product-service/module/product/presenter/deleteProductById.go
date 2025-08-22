package presenter

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
)

func (p *ProductPresenter) DeleteProductById(ctx context.Context, req *productRpc.DeleteProductByIdRequest) (*empty.Empty, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "Presenter.DeleteProductById")
	defer span.End()

	if err := req.Validate(); err != nil {
		p.logger.Error("ProductPresenter.CreateProduct", zap.String("requestID", pkgContext.GetRequestIDFromContext(ctx)), zap.Error(err))
		return nil, err
	}

	return nil, nil
}
