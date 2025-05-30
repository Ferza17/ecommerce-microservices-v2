package usecase

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"
)

func (u *productUseCase) DeleteProductById(ctx context.Context, requestId string, req *productRpc.DeleteProductByIdRequest) (*productRpc.DeleteProductByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}
