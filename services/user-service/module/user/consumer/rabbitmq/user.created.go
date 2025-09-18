package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/metric"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *userConsumer) UserCreated(ctx context.Context, d *amqp091.Delivery) error {
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "userConsumer.UserCreated")
	defer span.End()

	var (
		request   pb.AuthUserRegisterRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
	)

	switch d.ContentType {
	case enum.XProtobuf.String():
		if err := proto.Unmarshal(d.Body, &request); err != nil {
			span.RecordError(err)
			c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
			pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueUserCreated, "failed").Inc()
			return err
		}
	case enum.JSON.String():
		if err := json.Unmarshal(d.Body, &request); err != nil {
			span.RecordError(err)
			pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueUserCreated, "failed").Inc()
			c.logger.Error(fmt.Sprintf("failed to unmarshal request : %v", zap.Error(err)))
			return err
		}
	default:
		err := fmt.Errorf("invalid content type : %s", d.ContentType)
		span.RecordError(err)
		pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueUserCreated, "failed").Inc()
		c.logger.Error(fmt.Sprintf("failed to get request id"))
		return err
	}

	// TODO: Fix This Handler
	//if _, err := c.userUseCase.UpdateUserById(ctx, requestId, request); err != nil {
	//	span.RecordError(err)
	//	pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueUserCreated, "failed").Inc()
	//	c.logger.Error(fmt.Sprintf("failed to create user : %v", zap.Error(err)))
	//	return err
	//}

	if err := d.Ack(true); err != nil {
		return err
	}

	return nil
}
