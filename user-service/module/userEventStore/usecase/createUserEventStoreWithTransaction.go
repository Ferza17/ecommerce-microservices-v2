package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func (u *userEventStoreUseCase) CreateUserEventStoreWithTransaction(ctx context.Context, requestId string, req *pb.CreateUserEventStoreRequest) (*pb.CreateUserEventStoreResponse, error) {
	var (
		result string
		now    = time.Now().UTC()
	)

	session, err := u.userEventStoreRepository.StartSession(ctx)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		result, err = u.userEventStoreRepository.CreateUserEventStoreWithSession(ctx, requestId, &bson.Event{
			SagaID:        requestId,
			Entity:        req.Entity,
			EntityID:      req.EntityId,
			EventType:     req.EventType,
			Status:        req.Status,
			Payload:       utils.ConvertPbUserStateToBsonUserState(req.Payload),
			PreviousState: utils.ConvertPbUserStateToBsonUserState(req.PreviousState),
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

	return &pb.CreateUserEventStoreResponse{
		Id: result,
	}, nil
}
