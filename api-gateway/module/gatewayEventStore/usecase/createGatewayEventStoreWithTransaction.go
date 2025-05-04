package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
)

func (u *gatewayEventStoreUseCase) CreateGatewayEventStoreWithTransaction(ctx context.Context, requestId string, req *pb.CreateGatewayEventStoreRequest) (*pb.CreateGatewayEventStoreResponse, error) {
	//var (
	//	result string
	//	now    = time.Now().UTC()
	//)
	//
	//session, err := u.gatewayEventStoreRepository.C
	//if err != nil {
	//	u.logger.Error(err.Error())
	//	return nil, err
	//}
	//defer session.EndSession(ctx)
	//
	//callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
	//	result, err = u.gatewayEventStoreRepository.CreateUserEventStoreWithSession(ctx, requestId, &bson.Event{
	//		SagaID:        requestId,
	//		Entity:        req.Entity,
	//		EntityID:      req.EntityId,
	//		EventType:     req.EventType,
	//		Status:        req.Status,
	//		//Payload:       utils.ConvertPbUserStateToBsonUserState(req.Payload),
	//		//PreviousState: utils.ConvertPbUserStateToBsonUserState(req.PreviousState),
	//		CreatedAt:     now,
	//		UpdatedAt:     now,
	//	}, session)
	//	if err != nil {
	//		u.logger.Error(fmt.Sprintf("requestId : %s , error creating gatewayEventStoreUseCase event: %v", requestId, err))
	//		return nil, err
	//	}
	//	return nil, nil
	//}
	//
	//if _, err = session.WithTransaction(ctx, callback); err != nil {
	//	log.Fatalf("Transaction failed: %v", err)
	//}
	//
	//return &pb.CreateUserEventStoreResponse{
	//	Id: result,
	//}, nil
	return nil, nil
}
