package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
)

func (p *ProductGrpcPresenter) FindProductsWithPagination(ctx context.Context, req *pb.FindProductsWithPaginationRequest) (*pb.FindProductsWithPaginationResponse, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}
	
	return nil, nil
}
