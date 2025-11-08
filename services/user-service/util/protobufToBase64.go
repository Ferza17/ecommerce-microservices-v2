package util

import (
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/proto"
)

// ProtobufToBase64URL converts a protobuf message to a URL-safe base64-encoded string
func ProtobufToBase64URL(msg proto.Message) (string, error) {
	protoBytes, err := proto.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal protobuf: %w", err)
	}

	base64String := base64.URLEncoding.EncodeToString(protoBytes)
	return base64String, nil
}

// Base64URLToProtobuf converts a URL-safe base64-encoded string back to protobuf
func Base64URLToProtobuf(base64String string, msg proto.Message) error {
	protoBytes, err := base64.URLEncoding.DecodeString(base64String)
	if err != nil {
		return fmt.Errorf("failed to decode base64 URL: %w", err)
	}

	if err := proto.Unmarshal(protoBytes, msg); err != nil {
		return fmt.Errorf("failed to unmarshal protobuf: %w", err)
	}

	return nil
}
