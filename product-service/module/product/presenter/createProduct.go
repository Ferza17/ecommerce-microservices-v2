package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
)

func (p *ProductGrpcPresenter) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := p.ProductUseCase.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
