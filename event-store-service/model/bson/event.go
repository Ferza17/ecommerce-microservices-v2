package bson

import (
	"time"

	pb "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/v1/event"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Event struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AggregateID   string             `bson:"aggregate_id" json:"aggregate_id"`
	AggregateType string             `bson:"aggregate_type" json:"aggregate_type"`
	Version       int64              `bson:"version" json:"version"`
	Name          string             `bson:"name" json:"name"`
	OccurredAt    time.Time          `bson:"occurred_at" json:"occurred_at"`
	Payload       []byte             `bson:"payload" json:"payload"`
	Metadata      map[string]string  `bson:"metadata" json:"metadata"`
}

func (e *Event) ToProto() *pb.Event {
	return &pb.Event{
		Id:            e.ID.String(),
		AggregateId:   e.AggregateID,
		AggregateType: e.AggregateType,
		Version:       e.Version,
		Name:          e.Name,
		OccurredAt:    timestamppb.New(e.OccurredAt),
		Payload:       e.Payload,
		Metadata:      e.Metadata,
	}
}

func EventFromProto(pb *pb.Event) *Event {
	if pb == nil {
		return nil
	}
	return &Event{
		ID:            primitive.NewObjectID(),
		AggregateID:   pb.GetAggregateId(),
		AggregateType: pb.GetAggregateType(),
		Version:       pb.GetVersion(),
		Name:          pb.GetName(),
		OccurredAt:    pb.GetOccurredAt().AsTime(),
		Payload:       pb.GetPayload(),
		Metadata:      pb.GetMetadata(),
	}
}
