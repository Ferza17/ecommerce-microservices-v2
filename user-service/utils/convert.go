package utils

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
)

func ConvertPbUserStateToBsonUserState(input *pb.UserState) *bson.UserState {
	if input == nil {
		return nil
	}

	state := &bson.UserState{
		ID:       input.Id,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
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
