package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UserUseCase) CreateUser(ctx context.Context, requestId string, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var (
		err error = nil
	)

	eventStore := &pb.CreateGatewayEventStoreRequest{
		SagaId:        requestId,
		Entity:        "api-gateway",
		EntityId:      "",
		EventType:     "",
		Status:        enum.SUCCESS.String(),
		PreviousState: nil,
		CreatedAt:     timestamppb.Now(),
		UpdatedAt:     timestamppb.Now(),
	}

	defer func(err error, eventStore *pb.CreateGatewayEventStoreRequest) {
		eventStore.Status = enum.SUCCESS.String()
		if err != nil {
			eventStore.Status = enum.FAILED.String()
		}
		if _, err = u.gatewayEventStoreUseCase.CreateGatewayEventStore(ctx, requestId, eventStore); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
		}
	}(err, eventStore)

	payload, err := util.ConvertStructToProtoStruct(req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error converting struct to proto struct: %s", err.Error()))
		return nil, err
	}
	eventStore.Payload = payload

	message, err := proto.Marshal(req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		return nil, err
	}

	if err = u.rabbitMQ.Publish(ctx, requestId, enum.UserExchange, enum.USER_CREATED, message); err != nil {
		u.logger.Error(fmt.Sprintf("error publishing message to rabbitmq: %s", err.Error()))
		return nil, err
	}

	return &pb.CreateUserResponse{}, nil
}
