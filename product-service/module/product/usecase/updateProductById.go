package usecase

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"
)

func (u *productUseCase) UpdateProductById(ctx context.Context, requestId string, req *productRpc.UpdateProductByIdRequest) (*productRpc.Product, error) {
	//TODO implement me
	panic("implement me")
}
