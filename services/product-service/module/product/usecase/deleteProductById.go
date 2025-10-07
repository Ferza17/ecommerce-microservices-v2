package usecase

import (
	"context"

	pbEvent "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/event"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"github.com/golang/protobuf/ptypes/empty"
)

func (u *productUseCase) DeleteProductById(ctx context.Context, requestId string, req *productRpc.DeleteProductByIdRequest) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u *productUseCase) ConfirmDeleteProductById(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error {
	//TODO implement me
	panic("implement me")
}

func (u *productUseCase) CompensateDeleteProductById(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error {
	//TODO implement me
	panic("implement me")
}
