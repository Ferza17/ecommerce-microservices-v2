package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/util"
)

func (u *productUseCase) FindProductById(ctx context.Context, requestId string, req *pb.FindProductByIdRequest) (*pb.Product, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindProductById")
	defer span.End()

	fetchProduct, err := u.productElasticsearchRepository.FindProductById(ctx, requestId, req.GetId())
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding product by id: %v", requestId, err))
		return nil, err
	}

	return &pb.Product{
		Id:          fetchProduct.ID,
		Name:        fetchProduct.Name,
		Description: fetchProduct.Description,
		Uom:         fetchProduct.Uom,
		Image:       fetchProduct.Image,
		Price:       fetchProduct.Price,
		Stock:       fetchProduct.Stock,
		DiscardedAt: util.ConvertToProtoTimestamp(fetchProduct.DiscardedAt),
		CreatedAt:   util.ConvertToProtoTimestamp(fetchProduct.CreatedAt),
		UpdatedAt:   util.ConvertToProtoTimestamp(fetchProduct.UpdatedAt),
	}, nil
}
