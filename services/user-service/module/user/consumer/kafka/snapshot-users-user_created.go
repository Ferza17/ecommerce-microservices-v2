package kafka

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"google.golang.org/protobuf/proto"
)

func (c *userConsumer) SnapshotUsersUserCreated(ctx context.Context, message *kafka.Message) error {
	var (
		request   pb.AuthUserRegisterRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
	)

	if err := proto.Unmarshal(message.Value, &request); err != nil {
		c.logger.Info(fmt.Sprintf("proto.Unmarshal: %v", err))
		return err
	}

	return nil
}
