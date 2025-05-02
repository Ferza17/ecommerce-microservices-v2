package bson

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Event struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SagaID        string             `bson:"saga_id" json:"saga_id"`
	Entity        string             `bson:"entity" json:"entity"`
	EntityID      string             `bson:"entity_id" json:"entity_id"`
	EventType     string             `bson:"event_type" json:"event_type"`
	Status        string             `bson:"status" json:"status"` // pending, completed, failed, rolled_back
	Payload       *UserState         `bson:"payload" json:"payload"`
	PreviousState *UserState         `bson:"previous_state" json:"previous_state"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}

type UserState struct {
	ID          *string    `json:"id" bson:"id"`
	Name        *string    `json:"name" bson:"name"`
	Email       *string    `json:"email" bson:"email"`
	Password    *string    `json:"password" bson:"password"`
	CreatedAt   *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" bson:"updated_at"`
	DiscardedAt *time.Time `json:"discarded_at" bson:"discarded_at"`
}
