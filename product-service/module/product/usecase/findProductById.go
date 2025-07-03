package usecase

import (
	"context"
	"fmt"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *productUseCase) FindProductById(ctx context.Context, requestId string, req *productRpc.FindProductByIdRequest) (*productRpc.Product, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductUseCase.FindProductById")
	defer span.End()

	fetchProduct, err := u.productElasticsearchRepository.FindProductById(ctx, requestId, req.GetId())
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding product by id: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return fetchProduct.ToProto(), nil
}
