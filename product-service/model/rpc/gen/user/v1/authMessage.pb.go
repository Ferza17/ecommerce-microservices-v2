// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: user/v1/authMessage.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type UserLoginByEmailAndPasswordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserLoginByEmailAndPasswordRequest) Reset() {
	*x = UserLoginByEmailAndPasswordRequest{}
	mi := &file_user_v1_authMessage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLoginByEmailAndPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginByEmailAndPasswordRequest) ProtoMessage() {}

func (x *UserLoginByEmailAndPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_authMessage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginByEmailAndPasswordRequest.ProtoReflect.Descriptor instead.
func (*UserLoginByEmailAndPasswordRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_authMessage_proto_rawDescGZIP(), []int{0}
}

func (x *UserLoginByEmailAndPasswordRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserLoginByEmailAndPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLogoutByTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserLogoutByTokenRequest) Reset() {
	*x = UserLogoutByTokenRequest{}
	mi := &file_user_v1_authMessage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLogoutByTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLogoutByTokenRequest) ProtoMessage() {}

func (x *UserLogoutByTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_authMessage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLogoutByTokenRequest.ProtoReflect.Descriptor instead.
func (*UserLogoutByTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_authMessage_proto_rawDescGZIP(), []int{1}
}

func (x *UserLogoutByTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserLogoutByTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserLogoutByTokenResponse) Reset() {
	*x = UserLogoutByTokenResponse{}
	mi := &file_user_v1_authMessage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLogoutByTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLogoutByTokenResponse) ProtoMessage() {}

func (x *UserLogoutByTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_authMessage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLogoutByTokenResponse.ProtoReflect.Descriptor instead.
func (*UserLogoutByTokenResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_authMessage_proto_rawDescGZIP(), []int{2}
}

type FindUserByTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindUserByTokenRequest) Reset() {
	*x = FindUserByTokenRequest{}
	mi := &file_user_v1_authMessage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindUserByTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByTokenRequest) ProtoMessage() {}

func (x *FindUserByTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_authMessage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByTokenRequest.ProtoReflect.Descriptor instead.
func (*FindUserByTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_authMessage_proto_rawDescGZIP(), []int{3}
}

func (x *FindUserByTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserVerifyOtpRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Otp           string                 `protobuf:"bytes,1,opt,name=otp,proto3" json:"otp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserVerifyOtpRequest) Reset() {
	*x = UserVerifyOtpRequest{}
	mi := &file_user_v1_authMessage_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserVerifyOtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVerifyOtpRequest) ProtoMessage() {}

func (x *UserVerifyOtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_authMessage_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVerifyOtpRequest.ProtoReflect.Descriptor instead.
func (*UserVerifyOtpRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_authMessage_proto_rawDescGZIP(), []int{4}
}

func (x *UserVerifyOtpRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type UserVerifyOtpResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserVerifyOtpResponse) Reset() {
	*x = UserVerifyOtpResponse{}
	mi := &file_user_v1_authMessage_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserVerifyOtpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVerifyOtpResponse) ProtoMessage() {}

func (x *UserVerifyOtpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_authMessage_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVerifyOtpResponse.ProtoReflect.Descriptor instead.
func (*UserVerifyOtpResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_authMessage_proto_rawDescGZIP(), []int{5}
}

func (x *UserVerifyOtpResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *UserVerifyOtpResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_user_v1_authMessage_proto protoreflect.FileDescriptor

const file_user_v1_authMessage_proto_rawDesc = "" +
	"\n" +
	"\x19user/v1/authMessage.proto\x12\auser_v1\x1a\x1fgoogle/protobuf/timestamp.proto\"V\n" +
	"\"UserLoginByEmailAndPasswordRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"0\n" +
	"\x18UserLogoutByTokenRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"\x1b\n" +
	"\x19UserLogoutByTokenResponse\".\n" +
	"\x16FindUserByTokenRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"(\n" +
	"\x14UserVerifyOtpRequest\x12\x10\n" +
	"\x03otp\x18\x01 \x01(\tR\x03otp\"]\n" +
	"\x15UserVerifyOtpResponse\x12 \n" +
	"\vaccessToken\x18\x01 \x01(\tR\vaccessToken\x12\"\n" +
	"\frefreshToken\x18\x02 \x01(\tR\frefreshTokenb\x06proto3"

var (
	file_user_v1_authMessage_proto_rawDescOnce sync.Once
	file_user_v1_authMessage_proto_rawDescData []byte
)

func file_user_v1_authMessage_proto_rawDescGZIP() []byte {
	file_user_v1_authMessage_proto_rawDescOnce.Do(func() {
		file_user_v1_authMessage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_v1_authMessage_proto_rawDesc), len(file_user_v1_authMessage_proto_rawDesc)))
	})
	return file_user_v1_authMessage_proto_rawDescData
}

var file_user_v1_authMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_v1_authMessage_proto_goTypes = []any{
	(*UserLoginByEmailAndPasswordRequest)(nil), // 0: user_v1.UserLoginByEmailAndPasswordRequest
	(*UserLogoutByTokenRequest)(nil),           // 1: user_v1.UserLogoutByTokenRequest
	(*UserLogoutByTokenResponse)(nil),          // 2: user_v1.UserLogoutByTokenResponse
	(*FindUserByTokenRequest)(nil),             // 3: user_v1.FindUserByTokenRequest
	(*UserVerifyOtpRequest)(nil),               // 4: user_v1.UserVerifyOtpRequest
	(*UserVerifyOtpResponse)(nil),              // 5: user_v1.UserVerifyOtpResponse
}
var file_user_v1_authMessage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_v1_authMessage_proto_init() }
func file_user_v1_authMessage_proto_init() {
	if File_user_v1_authMessage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_v1_authMessage_proto_rawDesc), len(file_user_v1_authMessage_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_v1_authMessage_proto_goTypes,
		DependencyIndexes: file_user_v1_authMessage_proto_depIdxs,
		MessageInfos:      file_user_v1_authMessage_proto_msgTypes,
	}.Build()
	File_user_v1_authMessage_proto = out.File
	file_user_v1_authMessage_proto_goTypes = nil
	file_user_v1_authMessage_proto_depIdxs = nil
}
