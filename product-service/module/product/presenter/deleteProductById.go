package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *ProductGrpcPresenter) DeleteProductById(ctx context.Context, req *pb.DeleteProductByIdRequest) (*pb.DeleteProductByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductById not implemented")
}
