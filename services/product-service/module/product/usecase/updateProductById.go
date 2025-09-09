package usecase

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"github.com/golang/protobuf/ptypes/empty"
)

func (u *productUseCase) UpdateProductById(ctx context.Context, requestId string, req *productRpc.UpdateProductByIdRequest) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}
