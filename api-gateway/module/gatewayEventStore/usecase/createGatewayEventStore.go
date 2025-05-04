package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"time"
)

func (u *gatewayEventStoreUseCase) CreateGatewayEventStore(ctx context.Context, requestId string, req *pb.CreateGatewayEventStoreRequest) (*pb.CreateGatewayEventStoreResponse, error) {
	var (
		now     = time.Now().UTC()
		payload = req.Payload.AsMap()
	)

	result, err := u.gatewayEventStoreRepository.CreateGatewayEventStore(ctx, requestId, &bson.Event{
		SagaID:    requestId,
		Entity:    req.Entity,
		EntityID:  req.EntityId,
		EventType: req.EventType,
		Status:    req.Status,
		Payload:   &payload,
		//Payload:       utils.ConvertPbUserStateToBsonUserState(req.Payload),
		//PreviousState: utils.ConvertPbUserStateToBsonUserState(req.PreviousState),
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
	}

	return &pb.CreateGatewayEventStoreResponse{
		Id: result,
	}, err
}
