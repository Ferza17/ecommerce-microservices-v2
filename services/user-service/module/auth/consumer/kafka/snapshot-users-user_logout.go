package kafka

import (
	"context"
	"fmt"

	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
)

func (c *authConsumer) SnapshotUsersUserLogout(ctx context.Context, message *pbEvent.EventEnvelope) error {
	c.logger.Info("SnapshotUsersUserLogout")
	c.logger.Info(fmt.Sprintf("message: %v", message))
	return nil
}
