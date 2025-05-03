package bson

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Event struct {
		ID            primitive.ObjectID      `bson:"_id,omitempty" json:"id"`
		SagaID        string                  `bson:"saga_id" json:"saga_id"`
		Entity        string                  `bson:"entity" json:"entity"`
		EntityID      string                  `bson:"entity_id" json:"entity_id"`
		EventType     string                  `bson:"event_type" json:"event_type"`
		Status        string                  `bson:"status" json:"status"` // pending, completed, failed, rolled_back
		Payload       *map[string]interface{} `bson:"payload" json:"payload"`
		PreviousState *map[string]interface{} `bson:"previous_state" json:"previous_state"`
		CreatedAt     time.Time               `bson:"created_at" json:"created_at"`
		UpdatedAt     time.Time               `bson:"updated_at" json:"updated_at"`
	}
)
