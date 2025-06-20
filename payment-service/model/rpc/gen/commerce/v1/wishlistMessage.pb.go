// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: commerce/v1/wishlistMessage.proto

package gen

import (
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

type WishlistItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId string `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *WishlistItem) Reset() {
	*x = WishlistItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WishlistItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishlistItem) ProtoMessage() {}

func (x *WishlistItem) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishlistItem.ProtoReflect.Descriptor instead.
func (*WishlistItem) Descriptor() ([]byte, []int) {
	return file_commerce_v1_wishlistMessage_proto_rawDescGZIP(), []int{0}
}

func (x *WishlistItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WishlistItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *WishlistItem) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateWishlistItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateWishlistItemRequest) Reset() {
	*x = CreateWishlistItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWishlistItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWishlistItemRequest) ProtoMessage() {}

func (x *CreateWishlistItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWishlistItemRequest.ProtoReflect.Descriptor instead.
func (*CreateWishlistItemRequest) Descriptor() ([]byte, []int) {
	return file_commerce_v1_wishlistMessage_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWishlistItemRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CreateWishlistItemRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateWishlistItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateWishlistItemResponse) Reset() {
	*x = CreateWishlistItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWishlistItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWishlistItemResponse) ProtoMessage() {}

func (x *CreateWishlistItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWishlistItemResponse.ProtoReflect.Descriptor instead.
func (*CreateWishlistItemResponse) Descriptor() ([]byte, []int) {
	return file_commerce_v1_wishlistMessage_proto_rawDescGZIP(), []int{2}
}

func (x *CreateWishlistItemResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindWishlistItemWithPaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ProductIds []string `protobuf:"bytes,2,rep,name=productIds,proto3" json:"productIds,omitempty"`
	Page       int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit      int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *FindWishlistItemWithPaginationRequest) Reset() {
	*x = FindWishlistItemWithPaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindWishlistItemWithPaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindWishlistItemWithPaginationRequest) ProtoMessage() {}

func (x *FindWishlistItemWithPaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindWishlistItemWithPaginationRequest.ProtoReflect.Descriptor instead.
func (*FindWishlistItemWithPaginationRequest) Descriptor() ([]byte, []int) {
	return file_commerce_v1_wishlistMessage_proto_rawDescGZIP(), []int{3}
}

func (x *FindWishlistItemWithPaginationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FindWishlistItemWithPaginationRequest) GetProductIds() []string {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

func (x *FindWishlistItemWithPaginationRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindWishlistItemWithPaginationRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type FindWishlistItemWithPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*WishlistItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Page  int32           `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32           `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *FindWishlistItemWithPaginationResponse) Reset() {
	*x = FindWishlistItemWithPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindWishlistItemWithPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindWishlistItemWithPaginationResponse) ProtoMessage() {}

func (x *FindWishlistItemWithPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindWishlistItemWithPaginationResponse.ProtoReflect.Descriptor instead.
func (*FindWishlistItemWithPaginationResponse) Descriptor() ([]byte, []int) {
	return file_commerce_v1_wishlistMessage_proto_rawDescGZIP(), []int{4}
}

func (x *FindWishlistItemWithPaginationResponse) GetItems() []*WishlistItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FindWishlistItemWithPaginationResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindWishlistItemWithPaginationResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type DeleteWishlistItemByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWishlistItemByIdRequest) Reset() {
	*x = DeleteWishlistItemByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWishlistItemByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWishlistItemByIdRequest) ProtoMessage() {}

func (x *DeleteWishlistItemByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWishlistItemByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteWishlistItemByIdRequest) Descriptor() ([]byte, []int) {
	return file_commerce_v1_wishlistMessage_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteWishlistItemByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteWishlistItemByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DeleteWishlistItemByIdResponse) Reset() {
	*x = DeleteWishlistItemByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWishlistItemByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWishlistItemByIdResponse) ProtoMessage() {}

func (x *DeleteWishlistItemByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWishlistItemByIdResponse.ProtoReflect.Descriptor instead.
func (*DeleteWishlistItemByIdResponse) Descriptor() ([]byte, []int) {
	return file_commerce_v1_wishlistMessage_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteWishlistItemByIdResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_commerce_v1_wishlistMessage_proto protoreflect.FileDescriptor

var file_commerce_v1_wishlistMessage_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x69,
	0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x31,
	0x22, 0x54, 0x0a, 0x0c, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x1a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x25, 0x46, 0x69, 0x6e, 0x64,
	0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x57, 0x69, 0x74, 0x68,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x26, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x69, 0x73, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x73, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x2f, 0x0a, 0x1d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x1e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commerce_v1_wishlistMessage_proto_rawDescOnce sync.Once
	file_commerce_v1_wishlistMessage_proto_rawDescData = file_commerce_v1_wishlistMessage_proto_rawDesc
)

func file_commerce_v1_wishlistMessage_proto_rawDescGZIP() []byte {
	file_commerce_v1_wishlistMessage_proto_rawDescOnce.Do(func() {
		file_commerce_v1_wishlistMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_commerce_v1_wishlistMessage_proto_rawDescData)
	})
	return file_commerce_v1_wishlistMessage_proto_rawDescData
}

var file_commerce_v1_wishlistMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_commerce_v1_wishlistMessage_proto_goTypes = []interface{}{
	(*WishlistItem)(nil),                           // 0: commerce_v1.WishlistItem
	(*CreateWishlistItemRequest)(nil),              // 1: commerce_v1.CreateWishlistItemRequest
	(*CreateWishlistItemResponse)(nil),             // 2: commerce_v1.CreateWishlistItemResponse
	(*FindWishlistItemWithPaginationRequest)(nil),  // 3: commerce_v1.FindWishlistItemWithPaginationRequest
	(*FindWishlistItemWithPaginationResponse)(nil), // 4: commerce_v1.FindWishlistItemWithPaginationResponse
	(*DeleteWishlistItemByIdRequest)(nil),          // 5: commerce_v1.DeleteWishlistItemByIdRequest
	(*DeleteWishlistItemByIdResponse)(nil),         // 6: commerce_v1.DeleteWishlistItemByIdResponse
}
var file_commerce_v1_wishlistMessage_proto_depIdxs = []int32{
	0, // 0: commerce_v1.FindWishlistItemWithPaginationResponse.items:type_name -> commerce_v1.WishlistItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_commerce_v1_wishlistMessage_proto_init() }
func file_commerce_v1_wishlistMessage_proto_init() {
	if File_commerce_v1_wishlistMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commerce_v1_wishlistMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WishlistItem); i {
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
		file_commerce_v1_wishlistMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWishlistItemRequest); i {
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
		file_commerce_v1_wishlistMessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWishlistItemResponse); i {
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
		file_commerce_v1_wishlistMessage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindWishlistItemWithPaginationRequest); i {
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
		file_commerce_v1_wishlistMessage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindWishlistItemWithPaginationResponse); i {
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
		file_commerce_v1_wishlistMessage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWishlistItemByIdRequest); i {
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
		file_commerce_v1_wishlistMessage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWishlistItemByIdResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commerce_v1_wishlistMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commerce_v1_wishlistMessage_proto_goTypes,
		DependencyIndexes: file_commerce_v1_wishlistMessage_proto_depIdxs,
		MessageInfos:      file_commerce_v1_wishlistMessage_proto_msgTypes,
	}.Build()
	File_commerce_v1_wishlistMessage_proto = out.File
	file_commerce_v1_wishlistMessage_proto_rawDesc = nil
	file_commerce_v1_wishlistMessage_proto_goTypes = nil
	file_commerce_v1_wishlistMessage_proto_depIdxs = nil
}
