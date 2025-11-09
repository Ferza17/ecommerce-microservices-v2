package kafka

import (
	"context"
	"fmt"

	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	pbUser "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
)

func (c *userConsumer) SnapshotUsersUserCreated(ctx context.Context, message *pbEvent.EventEnvelope) error {
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

	if err = util.JSONToProto(message.Payload, &request); err != nil {
		c.logger.Info(fmt.Sprintf("util.JSONToProto: %v", err))
		return err
	}

	if _, err = c.userUseCase.CreateUser(ctx, pkgContext.GetRequestIDFromContext(ctx), &request); err != nil {
		c.logger.Info(fmt.Sprintf("userConsumer.CreateUser: %v", err))
		return err
	}

	return nil
}
