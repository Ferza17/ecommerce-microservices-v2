package util

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtoToJSON converts a protobuf message to a JSON string
func ProtoToJSON(msg proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		Multiline:       false,
		Indent:          "",
		EmitUnpopulated: true, // Include fields with default values
		UseProtoNames:   true, // Use proto field names (snake_case) instead of lowerCamelCase
	}

	jsonBytes, err := marshaler.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal proto to JSON: %w", err)
	}

	return string(jsonBytes), nil
}

// JSONToProto converts a JSON string to a protobuf message
func JSONToProto(jsonStr string, msg proto.Message) error {
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true, // Ignore unknown fields in JSON
	}

	err := unmarshaler.Unmarshal([]byte(jsonStr), msg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON to proto: %w", err)
	}

	return nil
}
