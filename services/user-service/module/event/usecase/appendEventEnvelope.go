package usecase

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func (u *eventUseCase) AppendEventEnvelope(ctx context.Context, request *pbEvent.EventEnvelope, payload proto.Message) error {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "EventUseCase.AppendEvent")
	defer span.End()

	metadata := map[string]string{}
	if pkgContext.GetRequestIDFromContext(ctx) != "" {
		metadata[pkgContext.CtxKeyRequestID] = pkgContext.GetRequestIDFromContext(ctx)
	}
	if pkgContext.GetTokenAuthorizationFromContext(ctx) != "" {
		metadata[pkgContext.CtxKeyAuthorization] = pkgContext.GetTokenAuthorizationFromContext(ctx)
	}
	carrier := u.telemetryInfrastructure.InjectSpanToTextMapPropagator(ctx)
	for k, v := range carrier {
		metadata[k] = v
	}
	request.Metadata = metadata

	payloadBase64String, err := util.ProtoToJSON(payload)
	if err != nil {
		u.logger.Error("eventUseCase.AppendEvent", zap.String("requestId", request.CorrelationId), zap.Error(err))
		return err
	}
	request.Payload = payloadBase64String

	// Send to an Event Store
	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkMongoEvent.EventEnvelopes, request.XId, kafka.PROTOBUF_SCHEMA, request); err != nil {
		u.logger.Error("eventUseCase.AppendEvent", zap.String("requestId", request.CorrelationId), zap.Error(err))
		return status.Error(codes.Internal, "internal server error")
	}
	return nil
}
