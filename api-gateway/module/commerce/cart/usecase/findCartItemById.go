package usecase

import (
	"context"
	"fmt"
	commerceRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/commerce/v1"
)

func (u *CartUseCase) FindCartItemById(ctx context.Context, requestId string, req *commerceRpc.FindCartItemByIdRequest) (*commerceRpc.CartItem, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindCartItemById")
	defer span.End()

	cartItem, err := u.commerceCartService.FindCartItemById(ctx, requestId, req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding products with pagination: %v", err))
		return nil, err
	}

	return cartItem, nil
}
