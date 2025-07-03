package presenter

import (
	"context"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
)

func (p *UserPresenter) FindUserById(ctx context.Context, req *userRpc.FindUserByIdRequest) (*userRpc.User, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "UserPresenter.FindUserById")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	fullMethodName, err := pkgContext.GetFullMethodNameFromContext(ctx)
	if err != nil {
		p.logger.Error("UserPresenter.FindUserByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}

	// Access Control Authorization
	acl, err := p.authUseCase.AuthUserVerifyAccessControl(
		ctx,
		pkgContext.GetRequestIDFromContext(ctx),
		&userRpc.AuthUserVerifyAccessControlRequest{
			Token:          pkgContext.GetTokenAuthorizationFromContext(ctx),
			FullMethodName: &fullMethodName,
		},
	)
	if err != nil {
		p.logger.Error("UserPresenter.FindUserByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, nil
	}

	if !acl.IsValid {
		p.logger.Error("UserPresenter.FindUserByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, nil
	}

	res, err := p.userUseCase.FindUserById(ctx, requestID, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
