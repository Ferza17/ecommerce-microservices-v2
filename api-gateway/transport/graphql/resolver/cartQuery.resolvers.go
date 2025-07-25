package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.73

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/graph/gen"
	gen1 "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/commerce/v1"
)

// FindCartItemsWithPagination is the resolver for the findCartItemsWithPagination field.
func (r *queryResolver) FindCartItemsWithPagination(ctx context.Context, input *gen1.FindCartItemsWithPaginationRequest) (*gen1.FindCartItemsWithPaginationResponse, error) {
	ctx, span := r.TelemetryInfrastructure.Tracer(ctx, "Resolver.FindCartItemsWithPagination")
	defer span.End()
	userId := ctx.Value(enum.ContextKeyUserID.String()).(string)
	if input == nil {
		input = &gen1.FindCartItemsWithPaginationRequest{}
	}
	input.UserId = userId
	return r.CartUseCase.FindCartItemsWithPagination(ctx, ctx.Value(enum.XRequestIDHeader.String()).(string), input)
}

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
