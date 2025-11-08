package util

import (
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func ProtobufToBase64(msg proto.Message) (string, error) {
	// Marshal protobuf to bytes
	protoBytes, err := proto.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal protobuf: %w", err)
	}

	// Encode to base64
	base64Str := base64.StdEncoding.EncodeToString(protoBytes)
	return base64Str, nil
}

func Base64URLToProtobuf(base64String string, msg proto.Message) error {
	protoBytes, err := base64.URLEncoding.DecodeString(base64String)
	if err != nil {
		return fmt.Errorf("failed to decode base64 URL: %w", err)
	}

	if err = proto.Unmarshal(protoBytes, msg); err != nil {
		return fmt.Errorf("failed to unmarshal protobuf: %w", err)
	}

	return nil
}
