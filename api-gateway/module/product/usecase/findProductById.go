package usecase

import (
	"context"
	"fmt"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/product/v1"
)

func (u *ProductUseCase) FindProductById(ctx context.Context, requestId string, req *productRpc.FindProductByIdRequest) (*productRpc.Product, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindProductById")
	defer span.End()

	product, err := u.productService.FindProductById(ctx, requestId, req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding product by id: %v", err))
		return nil, err
	}

	return product, nil
}
