package consumer

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *notificationEmailConsumer) SnapshotNotificationsEmailOtpUserRegister(ctx context.Context, message *kafka.Message) error {
	var (
		request   pb.SendOtpEmailNotificationRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
	)

	if err := proto.Unmarshal(message.Value, &request); err != nil {
		c.logger.Info(fmt.Sprintf("proto.Unmarshal: %v", err))
		return err
	}

	if err := c.notificationUseCase.SendNotificationEmailOTP(ctx, requestId, &request); err != nil {
		c.logger.Error(fmt.Sprintf("failed to send email OTP : %v", zap.Error(err)))
		return err
	}

	return nil
}
