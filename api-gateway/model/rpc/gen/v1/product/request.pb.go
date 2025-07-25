// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: v1/product/request.proto

package product

import (
	_ "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/external/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FindProductsWithPaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids   []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Name  []string `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	Page  int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *FindProductsWithPaginationRequest) Reset() {
	*x = FindProductsWithPaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_product_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindProductsWithPaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindProductsWithPaginationRequest) ProtoMessage() {}

func (x *FindProductsWithPaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_product_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindProductsWithPaginationRequest.ProtoReflect.Descriptor instead.
func (*FindProductsWithPaginationRequest) Descriptor() ([]byte, []int) {
	return file_v1_product_request_proto_rawDescGZIP(), []int{0}
}

func (x *FindProductsWithPaginationRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *FindProductsWithPaginationRequest) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *FindProductsWithPaginationRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindProductsWithPaginationRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type FindProductByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindProductByIdRequest) Reset() {
	*x = FindProductByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_product_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindProductByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindProductByIdRequest) ProtoMessage() {}

func (x *FindProductByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_product_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindProductByIdRequest.ProtoReflect.Descriptor instead.
func (*FindProductByIdRequest) Descriptor() ([]byte, []int) {
	return file_v1_product_request_proto_rawDescGZIP(), []int{1}
}

func (x *FindProductByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Uom         string  `protobuf:"bytes,3,opt,name=uom,proto3" json:"uom,omitempty"`
	Image       string  `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Price       float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Stock       int32   `protobuf:"varint,6,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_product_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_product_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_v1_product_request_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductRequest) GetUom() string {
	if x != nil {
		return x.Uom
	}
	return ""
}

func (x *CreateProductRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type UpdateProductByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        *string  `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description *string  `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Uom         *string  `protobuf:"bytes,4,opt,name=uom,proto3,oneof" json:"uom,omitempty"`
	Image       *string  `protobuf:"bytes,5,opt,name=image,proto3,oneof" json:"image,omitempty"`
	Price       *float64 `protobuf:"fixed64,6,opt,name=price,proto3,oneof" json:"price,omitempty"`
	Stock       *int32   `protobuf:"varint,7,opt,name=stock,proto3,oneof" json:"stock,omitempty"`
}

func (x *UpdateProductByIdRequest) Reset() {
	*x = UpdateProductByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_product_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductByIdRequest) ProtoMessage() {}

func (x *UpdateProductByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_product_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductByIdRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductByIdRequest) Descriptor() ([]byte, []int) {
	return file_v1_product_request_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProductByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductByIdRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateProductByIdRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateProductByIdRequest) GetUom() string {
	if x != nil && x.Uom != nil {
		return *x.Uom
	}
	return ""
}

func (x *UpdateProductByIdRequest) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *UpdateProductByIdRequest) GetPrice() float64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *UpdateProductByIdRequest) GetStock() int32 {
	if x != nil && x.Stock != nil {
		return *x.Stock
	}
	return 0
}

type DeleteProductByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteProductByIdRequest) Reset() {
	*x = DeleteProductByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_product_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductByIdRequest) ProtoMessage() {}

func (x *DeleteProductByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_product_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductByIdRequest) Descriptor() ([]byte, []int) {
	return file_v1_product_request_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteProductByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_v1_product_request_proto protoreflect.FileDescriptor

var file_v1_product_request_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x1a, 0x20, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x21, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x08, 0x00, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x00, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x31, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x6f, 0x6d,
	0x12, 0x1d, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x24, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0e,
	0xfa, 0x42, 0x0b, 0x12, 0x09, 0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x22, 0xd7, 0x02, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x03,
	0x75, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x48, 0x02, 0x52, 0x03, 0x75, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x48, 0x03, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x29, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x42,
	0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x48,
	0x04, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a,
	0x02, 0x20, 0x00, 0x48, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x6f, 0x6d,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x33,
	0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x42, 0x59, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x42, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x48, 0x02, 0x50, 0x01, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0xca, 0x02, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0xe2, 0x02,
	0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_product_request_proto_rawDescOnce sync.Once
	file_v1_product_request_proto_rawDescData = file_v1_product_request_proto_rawDesc
)

func file_v1_product_request_proto_rawDescGZIP() []byte {
	file_v1_product_request_proto_rawDescOnce.Do(func() {
		file_v1_product_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_product_request_proto_rawDescData)
	})
	return file_v1_product_request_proto_rawDescData
}

var file_v1_product_request_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_product_request_proto_goTypes = []interface{}{
	(*FindProductsWithPaginationRequest)(nil), // 0: product.FindProductsWithPaginationRequest
	(*FindProductByIdRequest)(nil),            // 1: product.FindProductByIdRequest
	(*CreateProductRequest)(nil),              // 2: product.CreateProductRequest
	(*UpdateProductByIdRequest)(nil),          // 3: product.UpdateProductByIdRequest
	(*DeleteProductByIdRequest)(nil),          // 4: product.DeleteProductByIdRequest
}
var file_v1_product_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_product_request_proto_init() }
func file_v1_product_request_proto_init() {
	if File_v1_product_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_product_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindProductsWithPaginationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_product_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindProductByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_product_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_product_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_product_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_v1_product_request_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_product_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_product_request_proto_goTypes,
		DependencyIndexes: file_v1_product_request_proto_depIdxs,
		MessageInfos:      file_v1_product_request_proto_msgTypes,
	}.Build()
	File_v1_product_request_proto = out.File
	file_v1_product_request_proto_rawDesc = nil
	file_v1_product_request_proto_goTypes = nil
	file_v1_product_request_proto_depIdxs = nil
}
