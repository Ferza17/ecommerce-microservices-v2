package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (p *UserPresenter) FindUserByEmailAndPassword(ctx context.Context, req *userRpc.FindUserByEmailAndPasswordRequest) (*userRpc.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "metadata not found")
	}
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.FindUserByEmailAndPassword")
	defer span.End()

	requestID := ""
	if values := md.Get(enum.XRequestIDHeader.String()); len(values) > 0 {
		requestID = values[0]
	}

	res, err := p.userUseCase.FindUserByEmailAndPassword(ctx, requestID, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
