package presenter

import (
	"context"
	"time"

	pbEvent "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/event"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
)

func (p *ProductPresenter) UpdateProductById(ctx context.Context, req *productRpc.UpdateProductByIdRequest) (*empty.Empty, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductPresenter.UpdateProductById")
	defer span.End()

	if err := req.Validate(); err != nil {
		p.logger.Error("ProductPresenter.UpdateProductById", zap.String("requestID", pkgContext.GetRequestIDFromContext(ctx)), zap.Error(err))
		return nil, err
	}

	res, err := p.productUseCase.UpdateProductById(ctx, pkgContext.GetRequestIDFromContext(ctx), req)
	if err != nil {
		p.logger.Error("ProductPresenter.CreateProduct", zap.String("requestID", pkgContext.GetRequestIDFromContext(ctx)), zap.Error(err))
		return nil, err
	}

	go func() {
		time.Sleep(5 * time.Second) // ensure event is created
		if err = p.productUseCase.ConfirmUpdateProductById(context.WithoutCancel(ctx), pkgContext.GetRequestIDFromContext(ctx), &pbEvent.ReserveEvent{
			SagaId:        pkgContext.GetRequestIDFromContext(ctx),
			AggregateType: "products",
		}); err != nil {
			p.logger.Error("ProductPresenter.CreateProduct", zap.String("requestID", pkgContext.GetRequestIDFromContext(ctx)), zap.Error(err))
			return
		}
	}()

	return res, nil
}
