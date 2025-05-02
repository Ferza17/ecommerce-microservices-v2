package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *ProductGrpcPresenter) UpdateProductById(ctx context.Context, req *pb.UpdateProductByIdRequest) (*pb.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductById not implemented")
}
