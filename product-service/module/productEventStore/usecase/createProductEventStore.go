package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/utils"
	"time"
)

func (u *productEventStoreUseCase) CreateProductEventStore(ctx context.Context, requestId string, req *pb.CreateProductEventStoreRequest) (*pb.CreateProductEventStoreResponse, error) {
	var (
		now = time.Now().UTC()
	)

	result, err := u.productEventStoreRepository.CreateProductEventStore(ctx, requestId, &bson.Event{
		SagaID:        requestId,
		Entity:        "products",
		EntityID:      req.EntityId,
		EventType:     req.EventType,
		Status:        req.Status,
		Payload:       utils.ConvertPbProductStateToBsonProductState(req.Payload),
		PreviousState: utils.ConvertPbProductStateToBsonProductState(req.PreviousState),
		CreatedAt:     now,
		UpdatedAt:     now,
	})
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
	}

	return &pb.CreateProductEventStoreResponse{
		Id: result,
	}, err
}
