package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"go.uber.org/zap"
)

func (c *userConsumer) UserCreated(ctx context.Context) error {
	var (
		request   *pb.CreateUserRequest
		requestId string
		ok        bool
	)

	q, err := c.amqpChannel.QueueDeclare(
		enum.USER_CREATED.String(),
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to serve queue : %v", zap.Error(err)))
		return err
	}

	msgs, err := c.amqpChannel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
	messages:
		for d := range msgs {
			if requestId, ok = d.Headers[enum.XRequestId.String()].(string); !ok {
				c.logger.Error("failed to get request id")
				continue messages
			}

			if err = json.Unmarshal(d.Body, &request); err != nil {
				c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
				continue messages
			}

			if _, err = c.userUseCase.CreateUser(ctx, requestId, request); err != nil {
				c.logger.Error(fmt.Sprintf("failed to create user : %v", zap.Error(err)))
				continue messages
			}
		}
	}()

	<-ctx.Done()
	return nil
}
