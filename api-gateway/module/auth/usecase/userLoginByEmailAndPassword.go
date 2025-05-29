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

func (u *authUseCase) UserLoginByEmailAndPassword(ctx context.Context, requestId string, req *pb.UserLoginByEmailAndPasswordRequest) error {
	var (
		err        error = nil
		eventStore       = &pb.EventStore{
			RequestId:     requestId,
			Service:       enum.UserService.String(),
			EventType:     enum.USER_LOGIN.String(),
			Status:        enum.PENDING.String(),
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.UserLoginByEmailAndPassword")
	defer span.End()

	// Validation If User Exists
	user, err := u.userService.FindUserByEmailAndPassword(ctx, requestId, &pb.FindUserByEmailAndPasswordRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding user by email and password: %s", err.Error()))
		return err
	}
	if user == nil {
		return fmt.Errorf("user not found")
	}

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
		return err
	}
	eventStore.Payload = payload

	message, err := proto.Marshal(req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		return err
	}

	if err = u.rabbitMQ.Publish(ctx, requestId, enum.UserExchange, enum.USER_LOGIN, message); err != nil {
		u.logger.Error(fmt.Sprintf("error publishing message to rabbitmq: %s", err.Error()))
		return err
	}

	return nil
}
