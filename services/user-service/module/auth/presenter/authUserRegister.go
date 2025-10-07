package presenter

import (
	"context"
	"time"

	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *AuthPresenter) AuthUserRegister(ctx context.Context, req *pb.AuthUserRegisterRequest) (*pb.AuthUserRegisterResponse, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthPresenter.AuthUserRegister")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	if err := req.Validate(); err != nil {
		p.logger.Error("AuthPresenter.AuthUserRegister", zap.String("requestID", requestID), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := p.userUseCase.CreateUser(ctx, requestID, req)
	if err != nil {
		p.logger.Error("AuthPresenter.AuthUserRegister", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}

	// Fire Forget
	go func() {
		select {
		case <-time.After(1000 * time.Millisecond):
			if err = p.userUseCase.ConfirmCreateUser(context.WithoutCancel(ctx), requestID, &pbEvent.ReserveEvent{
				SagaId:        requestID,
				AggregateType: "users",
			}); err != nil {
				p.logger.Error("AuthPresenter.AuthUserRegister", zap.String("requestID", requestID), zap.Error(err))
				return
			}
			return
		}
	}()

	return resp, nil
}
