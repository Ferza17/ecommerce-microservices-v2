package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (p *ProductGrpcPresenter) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "metadata not found")
	}
	requestID := ""
	if values := md.Get(enum.XRequestIDHeader.String()); len(values) > 0 {
		requestID = values[0]
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := p.ProductUseCase.CreateProduct(ctx, requestID, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
