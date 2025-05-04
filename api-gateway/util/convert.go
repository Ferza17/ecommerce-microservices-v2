package util

import (
	"encoding/json"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ConvertProtoTimestampToTime(ts *timestamppb.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	t := ts.AsTime()
	return &t
}

func ConvertTimeToProtoTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func ConvertStructToProtoStruct(input interface{}) (*structpb.Struct, error) {
	if input == nil {
		return nil, nil
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var mapData map[string]interface{}
	if err = json.Unmarshal(jsonData, &mapData); err != nil {
		return nil, err
	}

	return structpb.NewStruct(mapData)
}
