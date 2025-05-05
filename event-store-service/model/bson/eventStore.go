package bson

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Event struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RequestID     string             `bson:"request_id"`
	Service       string             `bson:"service"`
	EventType     string             `bson:"event_type"`
	Status        string             `bson:"status"`
	Payload       *map[string]any    `bson:"payload"`
	PreviousState *map[string]any    `bson:"previous_state,omitempty"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}
