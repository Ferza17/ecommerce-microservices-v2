package utils

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ConvertToProtoTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}
