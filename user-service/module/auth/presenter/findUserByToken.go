package presenter

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/user/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (p *AuthPresenter) FindUserByToken(ctx context.Context, req *userRpc.FindUserByTokenRequest) (*userRpc.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "metadata not found")
	}
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.FindUserByToken")
	defer span.End()

	requestID := ""
	if values := md.Get(enum.XRequestIDHeader.String()); len(values) > 0 {
		requestID = values[0]
	}

	resp, err := p.authUseCase.FindUserByToken(ctx, requestID, req)
	if err != nil {
		span.RecordError(err)
		p.logger.Error(fmt.Sprintf("requestId : %s , error finding user by token: %v", requestID, err))
		return nil, err
	}

	return resp, nil
}
