package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (u *productUseCase) CreateProduct(ctx context.Context, requestId string, req *productRpc.CreateProductRequest) (*empty.Empty, error) {
	var (
		err error
		tx  = u.postgres.GormDB.Begin()
		now = time.Now().UTC()
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "ProductUseCase.CreateProduct")
	defer span.End()

	_, err = u.productPgsqlRepository.CreateProduct(ctx, &orm.Product{
		ID:          uuid.NewString(),
		Name:        req.GetName(),
		Price:       req.GetPrice(),
		Stock:       int64(req.Stock),
		Description: req.GetDescription(),
		Image:       req.GetImage(),
		Uom:         req.GetUom(),
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}, tx)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error creating product: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// TODO: Insert Via CONNECTOR SINK

	tx.Commit()
	return nil, nil
}
