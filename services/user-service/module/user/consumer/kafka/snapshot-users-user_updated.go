package kafka

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pbUser "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

//TODO
// 1. Handle Snapshot
// 2. Handle Confirm
// 3. Handle Compensate

func (c *userConsumer) SnapshotUsersUserUpdated(ctx context.Context, message *kafka.Message) error {
	var (
		req pbUser.UpdateUserByIdRequest
		err error
	)
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "UserConsumer.FindUserByEmail")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	if err = proto.Unmarshal(message.Value, &req); err != nil {
		c.logger.Error("SnapshotUsersUserUpdated", zap.Error(err))
		return err
	}

	if _, err = c.userUseCase.UpdateUserById(ctx, pkgContext.GetRequestIDFromContext(ctx), &req); err != nil {
		c.logger.Error("SnapshotUsersUserUpdated", zap.Error(err))
		return err
	}

	return nil

}
