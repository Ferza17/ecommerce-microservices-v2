package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *eventUseCase) AppendEvent(ctx context.Context, request *pb.Event) error {
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

	if err := u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkMongoEvent.User, request.XId, kafka.PROTOBUF_SCHEMA, request); err != nil {
		u.logger.Error("eventUseCase.AppendEvent", zap.String("requestId", request.SagaId), zap.Error(err))
		return status.Error(codes.Internal, "internal server error")
	}

	return nil
}
