package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
)

func (u *ProductUseCase) FindProductById(ctx context.Context, requestId string, req *pb.FindProductByIdRequest) (*pb.Product, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindProductById")
	defer span.End()

	product, err := u.productService.FindProductById(ctx, requestId, req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding product by id: %v", err))
		return nil, err
	}

	return product, nil
}
