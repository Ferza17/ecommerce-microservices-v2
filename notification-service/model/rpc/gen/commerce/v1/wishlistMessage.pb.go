// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: commerce/v1/wishlistMessage.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WishlistItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId     string                 `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WishlistItem) Reset() {
	*x = WishlistItem{}
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WishlistItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishlistItem) ProtoMessage() {}

func (x *WishlistItem) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[0]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWishlistItemRequest) Reset() {
	*x = CreateWishlistItemRequest{}
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWishlistItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWishlistItemRequest) ProtoMessage() {}

func (x *CreateWishlistItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[1]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWishlistItemResponse) Reset() {
	*x = CreateWishlistItemResponse{}
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWishlistItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWishlistItemResponse) ProtoMessage() {}

func (x *CreateWishlistItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[2]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ProductIds    []string               `protobuf:"bytes,2,rep,name=productIds,proto3" json:"productIds,omitempty"`
	Page          int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindWishlistItemWithPaginationRequest) Reset() {
	*x = FindWishlistItemWithPaginationRequest{}
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindWishlistItemWithPaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindWishlistItemWithPaginationRequest) ProtoMessage() {}

func (x *FindWishlistItemWithPaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[3]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*WishlistItem        `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindWishlistItemWithPaginationResponse) Reset() {
	*x = FindWishlistItemWithPaginationResponse{}
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindWishlistItemWithPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindWishlistItemWithPaginationResponse) ProtoMessage() {}

func (x *FindWishlistItemWithPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[4]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWishlistItemByIdRequest) Reset() {
	*x = DeleteWishlistItemByIdRequest{}
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWishlistItemByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWishlistItemByIdRequest) ProtoMessage() {}

func (x *DeleteWishlistItemByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[5]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWishlistItemByIdResponse) Reset() {
	*x = DeleteWishlistItemByIdResponse{}
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWishlistItemByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWishlistItemByIdResponse) ProtoMessage() {}

func (x *DeleteWishlistItemByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commerce_v1_wishlistMessage_proto_msgTypes[6]
	if x != nil {
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

const file_commerce_v1_wishlistMessage_proto_rawDesc = "" +
	"\n" +
	"!commerce/v1/wishlistMessage.proto\x12\vcommerce_v1\"T\n" +
	"\fWishlistItem\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1c\n" +
	"\tproductId\x18\x02 \x01(\tR\tproductId\x12\x16\n" +
	"\x06userId\x18\x03 \x01(\tR\x06userId\"Q\n" +
	"\x19CreateWishlistItemRequest\x12\x1c\n" +
	"\tproductId\x18\x01 \x01(\tR\tproductId\x12\x16\n" +
	"\x06userId\x18\x02 \x01(\tR\x06userId\",\n" +
	"\x1aCreateWishlistItemResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\x89\x01\n" +
	"%FindWishlistItemWithPaginationRequest\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\tR\x06userId\x12\x1e\n" +
	"\n" +
	"productIds\x18\x02 \x03(\tR\n" +
	"productIds\x12\x12\n" +
	"\x04page\x18\x03 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x04 \x01(\x05R\x05limit\"\x83\x01\n" +
	"&FindWishlistItemWithPaginationResponse\x12/\n" +
	"\x05items\x18\x01 \x03(\v2\x19.commerce_v1.WishlistItemR\x05items\x12\x12\n" +
	"\x04page\x18\x02 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\x05R\x05limit\"/\n" +
	"\x1dDeleteWishlistItemByIdRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"8\n" +
	"\x1eDeleteWishlistItemByIdResponse\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\tR\x06userIdb\x06proto3"

var (
	file_commerce_v1_wishlistMessage_proto_rawDescOnce sync.Once
	file_commerce_v1_wishlistMessage_proto_rawDescData []byte
)

func file_commerce_v1_wishlistMessage_proto_rawDescGZIP() []byte {
	file_commerce_v1_wishlistMessage_proto_rawDescOnce.Do(func() {
		file_commerce_v1_wishlistMessage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_commerce_v1_wishlistMessage_proto_rawDesc), len(file_commerce_v1_wishlistMessage_proto_rawDesc)))
	})
	return file_commerce_v1_wishlistMessage_proto_rawDescData
}

var file_commerce_v1_wishlistMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_commerce_v1_wishlistMessage_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_commerce_v1_wishlistMessage_proto_rawDesc), len(file_commerce_v1_wishlistMessage_proto_rawDesc)),
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
	file_commerce_v1_wishlistMessage_proto_goTypes = nil
	file_commerce_v1_wishlistMessage_proto_depIdxs = nil
}
