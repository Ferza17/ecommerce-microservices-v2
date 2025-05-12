package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/util"
)

func (u *productUseCase) FindProductById(ctx context.Context, requestId string, req *pb.FindProductByIdRequest) (*pb.Product, error) {
	tx := u.productPgsqlRepository.OpenTransactionWithContext(ctx)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.UserLoginByEmailAndPassword")
	defer span.End()

	fetchProduct, err := u.productPgsqlRepository.FindProductById(ctx, req.GetId(), tx)
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
		DiscardedAt: util.ConvertToProtoTimestamp(fetchProduct.DiscardedAt),
		CreatedAt:   util.ConvertToProtoTimestamp(fetchProduct.CreatedAt),
		UpdatedAt:   util.ConvertToProtoTimestamp(fetchProduct.UpdatedAt),
	}, nil
}
