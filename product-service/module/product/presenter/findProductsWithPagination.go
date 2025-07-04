package presenter

import (
	"context"
	"fmt"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"go.uber.org/zap"
)

func (p *ProductPresenter) FindProductsWithPagination(ctx context.Context, req *productRpc.FindProductsWithPaginationRequest) (*productRpc.FindProductsWithPaginationResponse, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductPresenter.FindProductsWithPagination")
	defer span.End()

	requestId := pkgContext.GetRequestIDFromContext(ctx)
	if err := p.userService.AuthUserVerifyAccessControl(ctx, requestId); err != nil {
		p.logger.Error("ProductPresenter.CreateProduct", zap.String("requestID", requestId), zap.Error(err))
		return nil, err
	}

	resp, err := p.productUseCase.FindProductsWithPagination(ctx, requestId, req)
	if err != nil {
		p.logger.Error(fmt.Sprintf("error finding products: %v", err))
		return nil, err
	}

	return resp, nil
}
