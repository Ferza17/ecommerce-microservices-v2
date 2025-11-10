package consumer

import (
	"context"
	"fmt"

	pbEvent "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/event"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/util"
	"go.uber.org/zap"
)

func (c *notificationEmailConsumer) SnapshotNotificationsEmailPaymentOrderCreated(ctx context.Context, message *pbEvent.EventEnvelope) error {
	var (
		request   pb.SendEmailPaymentOrderCreateRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
	)

	if err := util.JSONToProto(message.Payload, &request); err != nil {
		c.logger.Info(fmt.Sprintf("util.JSONToProto: %v", err))
		return err
	}

	if err := c.notificationUseCase.SendNotificationEmailPaymentOrderCreated(ctx, requestId, &request); err != nil {
		c.logger.Error(fmt.Sprintf("failed to send email OTP : %v", zap.Error(err)))
		return err
	}

	return nil
}
