package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func (u *productEventStoreUseCase) CreateProductEventStoreWithTransaction(ctx context.Context, requestId string, req *pb.CreateProductEventStoreRequest) (*pb.CreateProductEventStoreResponse, error) {
	var (
		result string
		now    = time.Now().UTC()
	)

	session, err := u.productEventStoreRepository.StartSession()
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		result, err = u.productEventStoreRepository.CreateProductEventStoreWithSession(ctx, requestId, &bson.Event{
			SagaID:        requestId,
			Entity:        "products",
			EntityID:      req.EntityId,
			EventType:     req.EventType,
			Status:        req.Status,
			Payload:       utils.ConvertPbProductStateToBsonProductState(req.Payload),
			PreviousState: utils.ConvertPbProductStateToBsonProductState(req.PreviousState),
			CreatedAt:     now,
			UpdatedAt:     now,
		}, session)
		if err != nil {
			u.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStoreUseCase event: %v", requestId, err))
			return nil, err
		}
		return nil, nil
	}

	if _, err = session.WithTransaction(ctx, callback); err != nil {
		log.Fatalf("Transaction failed: %v", err)
	}

	return &pb.CreateProductEventStoreResponse{
		Id: result,
	}, nil
}
