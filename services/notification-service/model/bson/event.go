package bson

import (
	"time"
)

// Event represents a stored event in MongoDB.
type Event struct {
	ID            string            `bson:"_id,omitempty" json:"_id"`
	AggregateID   string            `bson:"aggregate_id" json:"aggregate_id"`
	AggregateType string            `bson:"aggregate_type" json:"aggregate_type"`
	EventType     string            `bson:"event_type" json:"event_type"`
	Version       int32             `bson:"version" json:"version"`
	Timestamp     time.Time         `bson:"timestamp" json:"timestamp"`
	SagaID        string            `bson:"saga_id" json:"saga_id"`
	Metadata      map[string]string `bson:"metadata" json:"metadata"`
	Payload       []byte            `bson:"payload" json:"payload"` // stored as BinData in Mongo
}

func (Event) CollectionName() string {
	return "notification_event_stores"
}
