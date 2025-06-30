package presenter

import (
	"context"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
)

func (p *UserPresenter) FindUserById(ctx context.Context, req *userRpc.FindUserByIdRequest) (*userRpc.User, error) {
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "UserPresenter.FindUserById")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	res, err := p.userUseCase.FindUserById(ctx, requestID, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
