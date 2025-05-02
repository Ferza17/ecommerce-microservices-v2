package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
)

func (p *ProductGrpcPresenter) UpdateProductById(ctx context.Context, req *pb.UpdateProductByIdRequest) (*pb.Product, error) {
	return nil, nil
}
