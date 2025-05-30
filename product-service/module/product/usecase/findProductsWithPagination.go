package usecase

import (
	"context"
	"fmt"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *productUseCase) FindProductsWithPagination(ctx context.Context, requestId string, req *productRpc.FindProductsWithPaginationRequest) (*productRpc.FindProductsWithPaginationResponse, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindProductsWithPagination")
	defer span.End()

	fetchedProducts, total, err := u.productElasticsearchRepository.FindProductsWithPagination(ctx, requestId, req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding products: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	var products []*productRpc.Product
	for _, product := range fetchedProducts {
		products = append(products, &productRpc.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Uom:         product.Uom,
			Image:       product.Image,
			Price:       product.Price,
			Stock:       product.Stock,
		})
		if product.CreatedAt != nil {
			products[len(products)-1].CreatedAt = timestamppb.New(*product.CreatedAt)
		}
		if product.UpdatedAt != nil {
			products[len(products)-1].UpdatedAt = timestamppb.New(*product.UpdatedAt)
		}
		if product.DiscardedAt != nil {
			products[len(products)-1].DiscardedAt = timestamppb.New(*product.DiscardedAt)
		}
	}

	response := &productRpc.FindProductsWithPaginationResponse{
		Data:  products,
		Total: int32(total),
		Page:  req.GetPage(),
		Limit: req.GetLimit(),
	}

	return response, nil
}
