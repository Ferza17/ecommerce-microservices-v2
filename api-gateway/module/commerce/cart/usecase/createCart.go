package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *CartUseCase) CreateCart(ctx context.Context, requestId string, req *pb.CreateCartItemRequest) (*pb.CreateCartItemResponse, error) {
	var (
		err        error = nil
		eventStore       = &pb.EventStore{
			RequestId:     requestId,
			Service:       enum.ProductService.String(),
			EventType:     enum.PRODUCT_CREATED.String(),
			Status:        enum.PENDING.String(),
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
		event = enum.CART_CREATED.String()
	)

	defer func(err error, eventStore *pb.EventStore) {
		if err != nil {
			eventStore.Status = enum.FAILED.String()
		}

		eventStoreMessage, err := proto.Marshal(eventStore)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
			return
		}

		if err = u.rabbitMQ.Publish(ctx, requestId, enum.EventExchange, enum.EVENT_CREATED, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}
	}(err, eventStore)

	payload, err := util.ConvertStructToProtoStruct(req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error converting struct to proto struct: %s", err.Error()))
		return nil, err
	}
	eventStore.Payload = payload

	//message, err := proto.Marshal(req)
	//if err != nil {
	//	u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
	//	return nil, err
	//}

	msg := map[string]interface{}{
		"pattern": enum.CART_CREATED,
		"data":    req,
	}

	message, err := json.Marshal(msg)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		return nil, err
	}

	if err = u.rabbitMQ.Publish(ctx, requestId, enum.CommerceExchange, enum.Queue(event), message); err != nil {
		u.logger.Error(fmt.Sprintf("error publishing message to rabbitmq: %s", err.Error()))
		return nil, err
	}

	return &pb.CreateCartItemResponse{}, nil
}
