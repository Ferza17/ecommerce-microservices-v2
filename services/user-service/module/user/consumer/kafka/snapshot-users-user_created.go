package kafka

import (
	"context"
	"fmt"

	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pbUser "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"

	"google.golang.org/protobuf/proto"
)

func (c *userConsumer) SnapshotUsersUserCreated(ctx context.Context, message *kafka.Message) error {
	var (
		request pbUser.AuthUserRegisterRequest
		err     error
	)

	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "UserConsumer.FindUserByEmail")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	if err = proto.Unmarshal(message.Value, &request); err != nil {
		c.logger.Info(fmt.Sprintf("proto.Unmarshal: %v", err))
		return err
	}

	if _, err = c.userUseCase.CreateUser(ctx, pkgContext.GetRequestIDFromContext(ctx), &request); err != nil {
		c.logger.Info(fmt.Sprintf("userConsumer.CreateUser: %v", err))
		return err
	}

	return nil
}

func (c *userConsumer) ConfirmSnapshotUsersUserCreated(ctx context.Context, message *kafka.Message) error {
	var (
		request pbEvent.ReserveEvent
		err     error
	)

	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "UserConsumer.FindUserByEmail")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	if err = proto.Unmarshal(message.Value, &request); err != nil {
		c.logger.Info(fmt.Sprintf("proto.Unmarshal: %v", err))
		return err
	}

	if err = c.userUseCase.ConfirmCreateUser(ctx, pkgContext.GetRequestIDFromContext(ctx), &request); err != nil {
		c.logger.Info(fmt.Sprintf("userConsumer.ConfirmSnapshotUsersUserCreated: %v", err))
		return err
	}

	return nil
}

func (c *userConsumer) CompensateSnapshotUsersUserCreated(ctx context.Context, message *kafka.Message) error {
	var (
		request pbEvent.ReserveEvent
		err     error
	)

	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "UserConsumer.FindUserByEmail")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	if err = proto.Unmarshal(message.Value, &request); err != nil {
		c.logger.Info(fmt.Sprintf("proto.Unmarshal: %v", err))
		return err
	}

	if err = c.userUseCase.CompensateCreateUser(ctx, pkgContext.GetRequestIDFromContext(ctx), &request); err != nil {
		c.logger.Info(fmt.Sprintf("userConsumer.CompensateSnapshotUsersUserCreated: %v", err))
		return err
	}

	return nil
}
