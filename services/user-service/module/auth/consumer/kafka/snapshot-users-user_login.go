package kafka

import (
	"context"
	"fmt"

	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
)

func (c *authConsumer) SnapshotUsersUserLogin(ctx context.Context, message *pbEvent.EventEnvelope) error {
	var (
		request   pb.AuthUserLoginByEmailAndPasswordRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
	)

	if err := util.Base64URLToProtobuf(message.Payload, &request); err != nil {
		c.logger.Info(fmt.Sprintf("util.Base64URLToProtobuf: %v", err))
		return err
	}

	if _, err := c.authUseCase.AuthUserLoginByEmailAndPassword(ctx, requestId, &request); err != nil {
		c.logger.Info(fmt.Sprintf("authUseCase.AuthUserLoginByEmailAndPassword: %v", err))
		return err
	}

	return nil
}
