package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/utils"
	"time"
)

func (u *userEventStoreUseCase) CreateUserEventStore(ctx context.Context, requestId string, req *pb.CreateUserEventStoreRequest) (*pb.CreateUserEventStoreResponse, error) {
	var (
		now = time.Now().UTC()
	)

	result, err := u.userEventStoreRepository.CreateUserEventStore(ctx, requestId, &bson.Event{
		SagaID:        requestId,
		Entity:        "users",
		EntityID:      req.EntityId,
		EventType:     req.EventType,
		Status:        req.Status,
		Payload:       utils.ConvertPbUserStateToBsonUserState(req.Payload),
		PreviousState: utils.ConvertPbUserStateToBsonUserState(req.PreviousState),
		CreatedAt:     now,
		UpdatedAt:     now,
	})
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
	}

	return &pb.CreateUserEventStoreResponse{
		Id: result,
	}, err
}
