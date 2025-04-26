package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"github.com/google/uuid"
	"time"
)

func (u *ProductUseCase) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	tx := u.ProductPgsqlRepository.OpenTransactionWithContext(ctx)
	now := time.Now().UTC()

	result, err := u.ProductPgsqlRepository.CreateProduct(ctx, &pb.ProductORM{
		Id:          uuid.NewString(),
		Name:        req.GetName(),
		Price:       req.GetPrice(),
		Stock:       req.GetStock(),
		Description: req.GetDescription(),
		Image:       req.GetImage(),
		Uom:         req.GetUom(),
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &pb.CreateProductResponse{
		Id: result,
	}, nil
}
