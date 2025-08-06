package consumer

import "C"
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

func (c *authConsumer) UserLogin(ctx context.Context, d *amqp091.Delivery) error {
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "userConsumer.UserUpdated")
	defer span.End()

	var (
		request   pb.AuthUserLoginByEmailAndPasswordRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
		err       error
	)

	defer func(err error) {
		if err != nil {
			span.RecordError(err)
			pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueUserLogin, "failed").Inc()
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

	//workflowID := fmt.Sprintf("user-login-%s", requestId)
	//workflowRun, err := c.temporal.StartWorkflow(ctx, workflowID, c.UserLoginWorkflow, ctx, requestId, &request)
	//if err != nil {
	//	c.logger.Error(fmt.Sprintf("failed to start workflow : %v", zap.Error(err)))
	//}

	//if _, err = c.authUseCase.AuthUserLoginByEmailAndPassword(ctx, requestId, &request); err != nil {
	//	c.logger.Error(fmt.Sprintf("failed to request user login : %v", zap.Error(err)))
	//	return err
	//}

	//c.logger.Info("UserLoginWorkflow started",
	//	zap.String("requestId", requestId),
	//	zap.String("workflowId", workflowID),
	//	zap.String("runId", workflowRun.GetRunID()))

	if err = d.Ack(true); err != nil {
		c.logger.Error(fmt.Sprintf("failed to ack delivery message : %v", zap.Error(err)))
		return err
	}
	return nil
}
