// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: v1/user/service.proto

package user

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_v1_user_service_proto protoreflect.FileDescriptor

var file_v1_user_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xc3, 0x06, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xe2, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x9a, 0x01, 0x92, 0x41, 0x40,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20,
	0x69, 0x64, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00,
	0x92, 0xb5, 0x18, 0x35, 0x12, 0x06, 0x82, 0x2b, 0x01, 0x00, 0x81, 0x2b, 0x1a, 0x1b, 0x0a, 0x13,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x2a, 0x0e, 0x0a, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a,
	0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xc3, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x7c, 0x92, 0x41, 0x36, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x19, 0x67, 0x65, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x77, 0x69, 0x74, 0x68,
	0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x69, 0x64, 0x62, 0x0c, 0x0a, 0x0a,
	0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x92, 0xb5, 0x18, 0x24, 0x12, 0x06,
	0x82, 0x2b, 0x01, 0x00, 0x81, 0x2b, 0x1a, 0x1a, 0x0a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x03, 0x67,
	0x65, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x88, 0x03,
	0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x27, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x6e, 0x64, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x96, 0x02, 0x92, 0x41, 0xb3, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x28, 0x66, 0x69, 0x6e, 0x64, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x62,
	0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x20, 0x61, 0x6e, 0x64, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x6c, 0x69,
	0x74, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x6c, 0x79, 0x20,
	0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x68, 0x74, 0x74, 0x70, 0x2c, 0x20, 0x69,
	0x74, 0x73, 0x20, 0x6f, 0x6e, 0x6c, 0x79, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x72, 0x70, 0x63, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x77, 0x69, 0x6c, 0x6c,
	0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x68, 0x74, 0x74, 0x70,
	0x20, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x62, 0x0c, 0x0a, 0x0a, 0x0a,
	0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x92, 0xb5, 0x18, 0x32, 0x12, 0x06, 0x82,
	0x2b, 0x01, 0x00, 0x81, 0x2b, 0x1a, 0x28, 0x0a, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x7d, 0x2f,
	0x7b, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x7d, 0x12, 0x03, 0x67, 0x65, 0x74, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x7d, 0x2f, 0x7b, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x7d, 0x32, 0xdb, 0x0d, 0x0a, 0x0b, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc6, 0x01, 0x0a, 0x10, 0x41, 0x75, 0x74,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x7b, 0x92, 0x41, 0x34, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x92, 0xb5, 0x18,
	0x22, 0x08, 0x01, 0x1a, 0x1e, 0x0a, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x81, 0x02, 0x0a, 0x1f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x97, 0x01, 0x92, 0x41,
	0x56, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x75, 0x73, 0x65, 0x72, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x92, 0xb5, 0x18, 0x1f, 0x08, 0x01, 0x1a, 0x1b, 0x0a,
	0x13, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0xdd, 0x01, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x74, 0x70, 0x12, 0x1e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x86, 0x01, 0x92,
	0x41, 0x30, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x21, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x74, 0x70,
	0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x6f,
	0x74, 0x70, 0x92, 0xb5, 0x18, 0x2c, 0x08, 0x01, 0x12, 0x06, 0x82, 0x2b, 0x01, 0x00, 0x81, 0x2b,
	0x1a, 0x20, 0x0a, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x6f, 0x74, 0x70, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6f, 0x74, 0x70, 0x2f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x8b, 0x02, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa8, 0x01, 0x92, 0x41, 0x5c, 0x0a, 0x0b,
	0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x6c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x1a, 0x29, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x0c, 0x0a, 0x0a,
	0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x92, 0xb5, 0x18, 0x26, 0x12, 0x06,
	0x82, 0x2b, 0x01, 0x00, 0x81, 0x2b, 0x1a, 0x1c, 0x0a, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x04,
	0x70, 0x6f, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x12, 0x85, 0x02, 0x0a, 0x1b, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x12, 0x28, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x90, 0x01, 0x92, 0x41, 0x3f, 0x0a, 0x0b,
	0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x92, 0xb5, 0x18,
	0x2a, 0x12, 0x06, 0x82, 0x2b, 0x01, 0x00, 0x81, 0x2b, 0x1a, 0x20, 0x0a, 0x18, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x63, 0x6c, 0x2f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x61, 0x63, 0x6c, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0xa4, 0x02, 0x0a, 0x1b,
	0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x49, 0x73, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x12, 0x28, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x49, 0x73, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x73,
	0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xaf, 0x01, 0x92, 0x41, 0x58, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x49, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x52, 0x50, 0x43, 0x20,
	0x46, 0x75, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x6f,
	0x72, 0x20, 0x48, 0x54, 0x54, 0x50, 0x20, 0x55, 0x52, 0x4c, 0x20, 0x26, 0x20, 0x48, 0x54, 0x54,
	0x50, 0x20, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x92, 0xb5,
	0x18, 0x2e, 0x08, 0x01, 0x12, 0x06, 0x82, 0x2b, 0x01, 0x00, 0x81, 0x2b, 0x1a, 0x22, 0x0a, 0x1a,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x63,
	0x6c, 0x2f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x63, 0x6c, 0x2f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x64, 0x12, 0xe1, 0x01, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x79, 0x92, 0x41, 0x2f,
	0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x67,
	0x65, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x74, 0x70, 0x20, 0x77, 0x69, 0x74, 0x68,
	0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x92,
	0xb5, 0x18, 0x26, 0x12, 0x06, 0x82, 0x2b, 0x01, 0x00, 0x81, 0x2b, 0x1a, 0x1c, 0x0a, 0x15, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x7b, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x7d, 0x12, 0x03, 0x67, 0x65, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12,
	0x15, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x7b,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x7d, 0x42, 0xbb, 0x02, 0x92, 0x41, 0xed, 0x01, 0x12, 0x4a, 0x0a,
	0x10, 0x55, 0x53, 0x45, 0x52, 0x20, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x20, 0x41, 0x50,
	0x49, 0x12, 0x2f, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x55, 0x53, 0x45, 0x52, 0x2c,
	0x20, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x2c,
	0x20, 0x61, 0x6e, 0x64, 0x20, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x2a, 0x02, 0x02, 0x01, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x5a, 0x69, 0x0a, 0x67, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x5d, 0x08,
	0x02, 0x12, 0x48, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x20,
	0x28, 0x65, 0x2e, 0x67, 0x2e, 0x2c, 0x20, 0x4a, 0x57, 0x54, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x20,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3a, 0x20, 0x60, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20,
	0x59, 0x4f, 0x55, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x60, 0x1a, 0x0d, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a,
	0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55,
	0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_v1_user_service_proto_goTypes = []interface{}{
	(*UpdateUserByIdRequest)(nil),                  // 0: user.UpdateUserByIdRequest
	(*FindUserByIdRequest)(nil),                    // 1: user.FindUserByIdRequest
	(*FindUserByEmailAndPasswordRequest)(nil),      // 2: user.FindUserByEmailAndPasswordRequest
	(*AuthUserRegisterRequest)(nil),                // 3: user.AuthUserRegisterRequest
	(*AuthUserLoginByEmailAndPasswordRequest)(nil), // 4: user.AuthUserLoginByEmailAndPasswordRequest
	(*AuthUserVerifyOtpRequest)(nil),               // 5: user.AuthUserVerifyOtpRequest
	(*AuthUserLogoutByTokenRequest)(nil),           // 6: user.AuthUserLogoutByTokenRequest
	(*AuthUserVerifyAccessControlRequest)(nil),     // 7: user.AuthUserVerifyAccessControlRequest
	(*AuthServiceVerifyIsExcludedRequest)(nil),     // 8: user.AuthServiceVerifyIsExcludedRequest
	(*AuthUserFindUserByTokenRequest)(nil),         // 9: user.AuthUserFindUserByTokenRequest
	(*emptypb.Empty)(nil),                          // 10: google.protobuf.Empty
	(*FindUserByIdResponse)(nil),                   // 11: user.FindUserByIdResponse
	(*FindUserByEmailAndPasswordResponse)(nil),     // 12: user.FindUserByEmailAndPasswordResponse
	(*AuthUserVerifyOtpResponse)(nil),              // 13: user.AuthUserVerifyOtpResponse
	(*AuthUserLogoutByTokenResponse)(nil),          // 14: user.AuthUserLogoutByTokenResponse
	(*AuthUserVerifyAccessControlResponse)(nil),    // 15: user.AuthUserVerifyAccessControlResponse
	(*AuthServiceVerifyIsExcludedResponse)(nil),    // 16: user.AuthServiceVerifyIsExcludedResponse
	(*AuthUserFindUserByTokenResponse)(nil),        // 17: user.AuthUserFindUserByTokenResponse
}
var file_v1_user_service_proto_depIdxs = []int32{
	0,  // 0: user.UserService.UpdateUserById:input_type -> user.UpdateUserByIdRequest
	1,  // 1: user.UserService.FindUserById:input_type -> user.FindUserByIdRequest
	2,  // 2: user.UserService.FindUserByEmailAndPassword:input_type -> user.FindUserByEmailAndPasswordRequest
	3,  // 3: user.AuthService.AuthUserRegister:input_type -> user.AuthUserRegisterRequest
	4,  // 4: user.AuthService.AuthUserLoginByEmailAndPassword:input_type -> user.AuthUserLoginByEmailAndPasswordRequest
	5,  // 5: user.AuthService.AuthUserVerifyOtp:input_type -> user.AuthUserVerifyOtpRequest
	6,  // 6: user.AuthService.AuthUserLogoutByToken:input_type -> user.AuthUserLogoutByTokenRequest
	7,  // 7: user.AuthService.AuthUserVerifyAccessControl:input_type -> user.AuthUserVerifyAccessControlRequest
	8,  // 8: user.AuthService.AuthServiceVerifyIsExcluded:input_type -> user.AuthServiceVerifyIsExcludedRequest
	9,  // 9: user.AuthService.AuthUserFindUserByToken:input_type -> user.AuthUserFindUserByTokenRequest
	10, // 10: user.UserService.UpdateUserById:output_type -> google.protobuf.Empty
	11, // 11: user.UserService.FindUserById:output_type -> user.FindUserByIdResponse
	12, // 12: user.UserService.FindUserByEmailAndPassword:output_type -> user.FindUserByEmailAndPasswordResponse
	10, // 13: user.AuthService.AuthUserRegister:output_type -> google.protobuf.Empty
	10, // 14: user.AuthService.AuthUserLoginByEmailAndPassword:output_type -> google.protobuf.Empty
	13, // 15: user.AuthService.AuthUserVerifyOtp:output_type -> user.AuthUserVerifyOtpResponse
	14, // 16: user.AuthService.AuthUserLogoutByToken:output_type -> user.AuthUserLogoutByTokenResponse
	15, // 17: user.AuthService.AuthUserVerifyAccessControl:output_type -> user.AuthUserVerifyAccessControlResponse
	16, // 18: user.AuthService.AuthServiceVerifyIsExcluded:output_type -> user.AuthServiceVerifyIsExcludedResponse
	17, // 19: user.AuthService.AuthUserFindUserByToken:output_type -> user.AuthUserFindUserByTokenResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_v1_user_service_proto_init() }
func file_v1_user_service_proto_init() {
	if File_v1_user_service_proto != nil {
		return
	}
	file_v1_user_model_proto_init()
	file_v1_user_request_proto_init()
	file_v1_user_response_proto_init()
	file_v1_user_option_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_v1_user_service_proto_goTypes,
		DependencyIndexes: file_v1_user_service_proto_depIdxs,
	}.Build()
	File_v1_user_service_proto = out.File
	file_v1_user_service_proto_rawDesc = nil
	file_v1_user_service_proto_goTypes = nil
	file_v1_user_service_proto_depIdxs = nil
}
