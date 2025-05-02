package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (u *productUseCase) CreateProduct(ctx context.Context, requestId string, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	tx := u.productPgsqlRepository.OpenTransactionWithContext(ctx)
	now := time.Now().UTC()
	createProductEventStoreReq := &pb.CreateProductEventStoreRequest{
		SagaId:    requestId,
		Entity:    "user",
		EventType: enum.PRODUCT_CREATED.String(),
		Status:    enum.SUCCESS.String(),
		Payload: &pb.ProductState{
			Name:        &req.Name,
			Price:       &req.Price,
			Stock:       &req.Stock,
			Description: &req.Description,
			Image:       &req.Image,
			Uom:         &req.Uom,
			CreatedAt:   timestamppb.New(now),
			UpdatedAt:   timestamppb.New(now),
			DiscardedAt: nil,
		},
	}

	result, err := u.productPgsqlRepository.CreateProduct(ctx, &orm.Product{
		ID:          uuid.NewString(),
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
		defer tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error creating product: %v", requestId, err))
		createProductEventStoreReq.Status = enum.FAILED.String()
		if _, err = u.productEventStoreUseCase.CreateProductEventStore(ctx, requestId, createProductEventStoreReq); err != nil {
			u.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
		}
		return nil, err
	}

	createProductEventStoreReq.EntityId = result
	createProductEventStoreReq.Payload.Id = &result
	if _, err = u.productEventStoreUseCase.CreateProductEventStore(ctx, requestId, createProductEventStoreReq); err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
		return nil, err
	}

	tx.Commit()
	return &pb.CreateProductResponse{
		Id: result,
	}, nil
}
