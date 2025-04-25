package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
)

func (p *ProductGrpcPresenter) FindProductById(ctx context.Context, req *pb.FindProductByIdRequest) (*pb.Product, error) {
	return nil, nil
}
