package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
)

func (u *productUseCase) FindProductsWithPagination(ctx context.Context, requestId string, req *pb.FindProductsWithPaginationRequest) (*pb.FindProductsWithPaginationResponse, error) {
	//TODO implement me
	panic("implement me")
}
