package presenter

import (
	"context"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
)

func (p *UserPresenter) FindUserByEmailAndPassword(ctx context.Context, req *userRpc.FindUserByEmailAndPasswordRequest) (*userRpc.FindUserByEmailAndPasswordResponse, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "UserPresenter.FindUserByEmailAndPassword")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	res, err := p.userUseCase.FindUserByEmailAndPassword(ctx, requestID, req)
	if err != nil {
		p.logger.Error("UserPresenter.FindUserByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}
	return res, nil
}
