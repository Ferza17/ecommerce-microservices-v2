package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"gorm.io/gorm"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *productUseCase) FindProductsWithPagination(ctx context.Context, requestId string, req *productRpc.FindProductsWithPaginationRequest) (*productRpc.FindProductsWithPaginationResponse, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()

	ctxTimeout, span := u.telemetryInfrastructure.StartSpanFromContext(ctxTimeout, "UseCase.FindProductsWithPagination")
	defer span.End()

	fetchedProducts, total, err := u.productElasticsearchRepository.FindProductsWithPagination(ctxTimeout, requestId, req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding products: %v", requestId, err))
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "product not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(fetchedProducts) == 0 {
		return nil, status.Error(codes.NotFound, "product not found")
	}

	return &productRpc.FindProductsWithPaginationResponse{
		Status:  "success",
		Message: "FindProductsWithPagination",
		Data: &productRpc.FindProductsWithPaginationResponse_FindProductsWithPaginationResponseData{
			Data:  orm.ProductsToProto(fetchedProducts),
			Total: int32(total),
			Page:  req.GetPage(),
			Limit: req.GetLimit(),
		},
	}, nil
}
