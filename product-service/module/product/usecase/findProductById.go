package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
)

func (u *ProductUseCase) FindProductById(ctx context.Context, req *pb.FindProductByIdRequest) (*pb.Product, error) {
	tx := u.ProductPgsqlRepository.OpenTransactionWithContext(ctx)

	fetchProduct, err := u.ProductPgsqlRepository.FindProductById(ctx, req.GetId(), tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	product, err := fetchProduct.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &product, nil
}
