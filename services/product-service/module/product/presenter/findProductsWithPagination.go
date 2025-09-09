package presenter

import (
	"context"
	"fmt"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
)

func (p *ProductPresenter) FindProductsWithPagination(ctx context.Context, req *productRpc.FindProductsWithPaginationRequest) (*productRpc.FindProductsWithPaginationResponse, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductPresenter.FindProductsWithPagination")
	defer span.End()
	resp, err := p.productUseCase.FindProductsWithPagination(ctx, pkgContext.GetRequestIDFromContext(ctx), req)
	if err != nil {
		p.logger.Error(fmt.Sprintf("error finding products: %v", err))
		return nil, err
	}

	return resp, nil
}
