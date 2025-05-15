package presenter

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (p *ProductGrpcPresenter) FindProductsWithPagination(ctx context.Context, req *pb.FindProductsWithPaginationRequest) (*pb.FindProductsWithPaginationResponse, error) {
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.FindProductsWithPagination")
	defer span.End()
	if err := req.Validate(); err != nil {
		p.logger.Error(fmt.Sprintf("error validating request: %v", err))
		return nil, err
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		p.logger.Error(fmt.Sprintf("metadata not found"))
		return nil, status.Error(codes.InvalidArgument, "metadata not found")
	}

	requestID := ""
	if values := md.Get(enum.XRequestIDHeader.String()); len(values) > 0 {
		requestID = values[0]
	}

	resp, err := p.productUseCase.FindProductsWithPagination(ctx, requestID, req)
	if err != nil {
		p.logger.Error(fmt.Sprintf("error finding products: %v", err))
		return nil, err
	}

	return resp, nil
}
