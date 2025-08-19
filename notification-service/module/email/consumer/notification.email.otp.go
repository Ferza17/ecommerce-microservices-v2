package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/context"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/metric"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *notificationEmailConsumer) NotificationEmailOTP(ctx context.Context, d *amqp091.Delivery) error {
	var (
		request   pb.SendOtpEmailNotificationRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
		err       error
	)

	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "EmailConsumer.NotificationEmailOTP")
	defer func(err error) {
		if err != nil {
			span.RecordError(err)
			pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueNotificationEmailOtpCreated, "failed").Inc()
			if err = d.Nack(true, true); err != nil {
				c.logger.Error(fmt.Sprintf("failed to nack : %v", zap.Error(err)))
				return
			}
			//TODO: Publish Event Notification Email Otp Failed
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

	if err = c.notificationUseCase.SendNotificationEmailOTP(ctx, requestId, &request); err != nil {
		c.logger.Error(fmt.Sprintf("failed to send email OTP : %v", zap.Error(err)))
		return err
	}

	if err = d.Ack(true); err != nil {
		c.logger.Error(fmt.Sprintf("failed to ack delivery message : %v", zap.Error(err)))
		return err
	}

	return nil
}
