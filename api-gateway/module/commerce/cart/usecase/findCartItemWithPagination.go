package usecase

import (
	"context"
	"fmt"
	commerceRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/commerce/v1"
)

func (u *CartUseCase) FindCartItemsWithPagination(ctx context.Context, requestId string, req *commerceRpc.FindCartItemsWithPaginationRequest) (*commerceRpc.FindCartItemsWithPaginationResponse, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindCartItemsWithPagination")
	defer span.End()

	cartItems, err := u.commerceCartService.FindCartItemsWithPagination(ctx, requestId, req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding cartItem with pagination: %v", err))
		return nil, err
	}

	return cartItems, nil
}
