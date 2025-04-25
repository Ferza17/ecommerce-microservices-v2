package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
)

func (p *ProductGrpcPresenter) FindProductsWithPagination(ctx context.Context, req *pb.FindProductsWithPaginationRequest) (*pb.FindProductsWithPaginationResponse, error) {
	return nil, nil
}
