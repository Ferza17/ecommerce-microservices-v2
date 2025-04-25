package presenter

import "github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"

type ProductGrpcPresenter struct {
	pb.UnimplementedProductServiceServer
}

func NewProductGrpcPresenter() *ProductGrpcPresenter {
	return &ProductGrpcPresenter{}
}
