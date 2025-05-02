package utils

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ConvertToProtoTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func ConvertPbProductStateToBsonProductState(input *pb.ProductState) *bson.ProductState {
	if input == nil {
		return nil
	}

	state := &bson.ProductState{
		ID:          input.Id,
		Name:        input.Name,
		Price:       input.Price,
		Stock:       input.Stock,
		Description: input.Description,
		Image:       input.Image,
		Uom:         input.Uom,
	}

	if input.CreatedAt != nil {
		t := input.CreatedAt.AsTime().UTC()
		state.CreatedAt = &t
	}
	if input.UpdatedAt != nil {
		t := input.UpdatedAt.AsTime().UTC()
		state.UpdatedAt = &t
	}
	if input.DiscardedAt != nil {
		t := input.DiscardedAt.AsTime().UTC()
		state.DiscardedAt = &t
	}

	return state
}
