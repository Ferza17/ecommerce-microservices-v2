package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/utils"
)

func (u *ProductUseCase) FindProductById(ctx context.Context, req *pb.FindProductByIdRequest) (*pb.Product, error) {
	tx := u.ProductPgsqlRepository.OpenTransactionWithContext(ctx)

	fetchProduct, err := u.ProductPgsqlRepository.FindProductById(ctx, req.GetId(), tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &pb.Product{
		Id:          fetchProduct.ID,
		Name:        fetchProduct.Name,
		Description: fetchProduct.Description,
		Uom:         fetchProduct.Uom,
		Image:       fetchProduct.Image,
		Price:       fetchProduct.Price,
		Stock:       fetchProduct.Stock,
		DiscardedAt: utils.ConvertToProtoTimestamp(fetchProduct.DiscardedAt),
		CreatedAt:   utils.ConvertToProtoTimestamp(fetchProduct.CreatedAt),
		UpdatedAt:   utils.ConvertToProtoTimestamp(fetchProduct.UpdatedAt),
	}, nil
}
