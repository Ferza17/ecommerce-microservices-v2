package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
)

func (u *ProductUseCase) FindProductsWithPagination(ctx context.Context, requestId string, req *pb.FindProductsWithPaginationRequest) (*pb.FindProductsWithPaginationResponse, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindProductsWithPagination")
	defer span.End()

	products, err := u.productService.FindProductsWithPagination(ctx, requestId, req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding products with pagination: %v", err))
		return nil, err
	}

	return products, nil
}
