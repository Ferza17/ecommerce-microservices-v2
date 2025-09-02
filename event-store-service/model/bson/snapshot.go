package bson

import (
	"time"

	pb "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/v1/event"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Snapshot struct {
	AggregateID   string            `bson:"aggregate_id" json:"aggregate_id"`
	AggregateType string            `bson:"aggregate_type" json:"aggregate_type"`
	Version       int64             `bson:"version" json:"version"`
	State         []byte            `bson:"state" json:"state"`
	TakenAt       time.Time         `bson:"taken_at" json:"taken_at"`
	Metadata      map[string]string `bson:"metadata" json:"metadata"`
}

func (s *Snapshot) ToProto() *pb.Snapshot {
	return &pb.Snapshot{
		AggregateId:   s.AggregateID,
		AggregateType: s.AggregateType,
		Version:       s.Version,
		State:         s.State,
		TakenAt:       timestamppb.New(s.TakenAt),
		Metadata:      s.Metadata,
	}
}

func SnapshotFromProto(pb *pb.Snapshot) *Snapshot {
	if pb == nil {
		return nil
	}
	return &Snapshot{
		AggregateID:   pb.GetAggregateId(),
		AggregateType: pb.GetAggregateType(),
		Version:       pb.GetVersion(),
		State:         pb.GetState(),
		TakenAt:       pb.GetTakenAt().AsTime(),
		Metadata:      pb.GetMetadata(),
	}
}
