package payment

import (
	"context"
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/payment"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *paymentService) FindPaymentById(ctx context.Context, requestId string, req *pb.FindPaymentByIdRequest) (*pb.FindPaymentByIdResponse, error) {
	md := metadata.New(map[string]string{
		pkgContext.CtxKeyRequestID:     requestId,
		pkgContext.CtxKeyAuthorization: fmt.Sprintf("Bearer %s", pkgContext.GetTokenAuthorizationFromContext(ctx)),
	})
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	for key, value := range carrier {
		md.Set(key, value)
	}

	resp, err := s.paymentSvc.FindPaymentById(metadata.NewOutgoingContext(ctx, md), req, grpc.Header(&md))
	if err != nil {
		s.logger.Error("PaymentService.FindPaymentById", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}

	if resp.Data == nil {
		s.logger.Error("PaymentService.FindPaymentById", zap.String("requestId", requestId))
		return nil, status.Error(codes.NotFound, "Payment not found")
	}

	if resp.Data.Payment == nil {
		s.logger.Error("PaymentService.FindPaymentById", zap.String("requestId", requestId))
		return nil, status.Error(codes.NotFound, "Payment not found")
	}

	if resp.Data.PaymentItems == nil || len(resp.Data.PaymentItems) == 0 {
		s.logger.Error("PaymentService.FindPaymentById", zap.String("requestId", requestId))
		return nil, status.Error(codes.NotFound, "Payment item not found")
	}

	if resp.Data.Provider == nil {
		s.logger.Error("PaymentService.FindPaymentById", zap.String("requestId", requestId))
		return nil, status.Error(codes.NotFound, "Payment provider not found")
	}

	return resp, nil
}
