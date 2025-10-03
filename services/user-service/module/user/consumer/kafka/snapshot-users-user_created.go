package kafka

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"google.golang.org/protobuf/proto"
)

//TODO
// 1. Handle Snapshot
// 2. Handle Confirm
// 3. Handle Compensate

func (c *userConsumer) SnapshotUsersUserCreated(ctx context.Context, message *kafka.Message) error {
	var (
		request pb.AuthUserRegisterRequest
	)

	if err := proto.Unmarshal(message.Value, &request); err != nil {
		c.logger.Info(fmt.Sprintf("proto.Unmarshal: %v", err))
		return err
	}

	return nil
}

func (c *userConsumer) ConfirmSnapshotUsersUserCreated(ctx context.Context, message *kafka.Message) error {
	var (
		request pb.AuthUserRegisterRequest
	)

	if err := proto.Unmarshal(message.Value, &request); err != nil {
		c.logger.Info(fmt.Sprintf("proto.Unmarshal: %v", err))
		return err
	}

	return nil
}

func (c *userConsumer) CompensateSnapshotUsersUserCreated(ctx context.Context, message *kafka.Message) error {
	var (
		request pb.AuthUserRegisterRequest
	)

	if err := proto.Unmarshal(message.Value, &request); err != nil {
		c.logger.Info(fmt.Sprintf("proto.Unmarshal: %v", err))
		return err
	}

	return nil
}
