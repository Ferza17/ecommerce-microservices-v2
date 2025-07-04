package presenter

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
)

func (p *ProductPresenter) UpdateProductById(ctx context.Context, req *productRpc.UpdateProductByIdRequest) (*empty.Empty, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductPresenter.UpdateProductById")
	defer span.End()
	requestId := pkgContext.GetRequestIDFromContext(ctx)

	if err := p.userService.AuthUserVerifyAccessControl(ctx, requestId); err != nil {
		p.logger.Error("ProductPresenter.UpdateProductById", zap.String("requestID", requestId), zap.Error(err))
		return nil, err
	}

	if err := req.Validate(); err != nil {
		p.logger.Error("ProductPresenter.UpdateProductById", zap.String("requestID", requestId), zap.Error(err))
		return nil, err
	}

	return nil, nil
}
