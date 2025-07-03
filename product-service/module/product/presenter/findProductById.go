package presenter

import (
	"context"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *ProductPresenter) FindProductById(ctx context.Context, req *productRpc.FindProductByIdRequest) (*productRpc.Product, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthPresenter.AuthUserRegister")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	if err := p.userService.AuthUserVerifyAccessControl(ctx, requestID); err != nil {
		p.logger.Error("Presenter.CreateProduct", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}

	if err := req.Validate(); err != nil {
		p.logger.Error("ProductPresenter.FindProductById", zap.String("requestID", requestID), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res, err := p.productUseCase.FindProductById(ctx, requestID, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
