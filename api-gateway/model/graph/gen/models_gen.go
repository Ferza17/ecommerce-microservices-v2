// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

import (
	"time"
)

type DeleteCartItemRequest struct {
	ProductID *string `json:"productId,omitempty"`
	UserID    *string `json:"userId,omitempty"`
}

type DeleteCartItemResponse struct {
	UserID *string `json:"userId,omitempty"`
}

type Mutation struct {
}

type PaymentItemInput struct {
	ID          string     `json:"id"`
	ProductID   string     `json:"productId"`
	Amount      float64    `json:"amount"`
	Qty         int        `json:"qty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	DiscardedAt *time.Time `json:"discardedAt,omitempty"`
}

type Query struct {
}
