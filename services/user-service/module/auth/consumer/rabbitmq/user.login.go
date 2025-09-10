package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	eventPb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/metric"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (c *authConsumer) UserLogin(ctx context.Context, d *amqp091.Delivery) error {
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "userConsumer.UserUpdated")
	defer span.End()

	var (
		request   pb.AuthUserLoginByEmailAndPasswordRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
		err       error
	)

	// TESTING
	defer func(err error) {
		var AppendEvent = eventPb.AppendRequest{
			AggregateId:     requestId,
			AggregateType:   config.Get().QueueUserLogin,
			ExpectedVersion: 0,
			Events: []*eventPb.Event{
				{
					Id:            uuid.NewString(),
					AggregateId:   requestId,
					AggregateType: config.Get().QueueUserLogin,
					Version:       0,
					Name:          config.Get().QueueUserLogin,
					OccurredAt:    timestamppb.New(time.Now()),
					Payload:       d.Body,
				},
			},
		}

		if err != nil {
			span.RecordError(err)
			pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueUserLogin, "failed").Inc()
		} else {
			pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueUserLogin, "success").Inc()
			c.logger.Info(fmt.Sprintf("success to consume message : %s", d.Body))
			if err = c.eventUseCase.AppendEvent(ctx, &AppendEvent); err != nil {
				c.logger.Error(fmt.Sprintf("failed to append event : %v", zap.Error(err)))
			}
		}

		span.End()
	}(err)

	switch d.ContentType {
	case enum.XProtobuf.String():
		if err = proto.Unmarshal(d.Body, &request); err != nil {
			c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
			return err
		}
	case enum.JSON.String():
		if err = json.Unmarshal(d.Body, &request); err != nil {
			c.logger.Error(fmt.Sprintf("failed to unmarshal request : %v", zap.Error(err)))
			return err
		}
	default:
		err = fmt.Errorf("invalid content type : %s", d.ContentType)
		c.logger.Error(fmt.Sprintf("failed to get request id"))
		return err
	}

	if _, err = c.authUseCase.AuthUserLoginByEmailAndPassword(ctx, requestId, &request); err != nil {
		c.logger.Error(fmt.Sprintf("failed to request user login : %v", zap.Error(err)))
		return err
	}

	if err = d.Ack(true); err != nil {
		c.logger.Error(fmt.Sprintf("failed to ack delivery message : %v", zap.Error(err)))
		return err
	}
	return nil
}
