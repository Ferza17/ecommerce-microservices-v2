// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: notification/v1/notificationMessage.proto

package gen

import (
	v1 "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/payment/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotificationTypeEnum int32

const (
	NotificationTypeEnum_NOTIFICATION_EMAIL_USER_OTP              NotificationTypeEnum = 0
	NotificationTypeEnum_NOTIFICATION_EMAIL_PAYMENT_ORDER_CREATED NotificationTypeEnum = 1
)

// Enum value maps for NotificationTypeEnum.
var (
	NotificationTypeEnum_name = map[int32]string{
		0: "NOTIFICATION_EMAIL_USER_OTP",
		1: "NOTIFICATION_EMAIL_PAYMENT_ORDER_CREATED",
	}
	NotificationTypeEnum_value = map[string]int32{
		"NOTIFICATION_EMAIL_USER_OTP":              0,
		"NOTIFICATION_EMAIL_PAYMENT_ORDER_CREATED": 1,
	}
)

func (x NotificationTypeEnum) Enum() *NotificationTypeEnum {
	p := new(NotificationTypeEnum)
	*p = x
	return p
}

func (x NotificationTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_v1_notificationMessage_proto_enumTypes[0].Descriptor()
}

func (NotificationTypeEnum) Type() protoreflect.EnumType {
	return &file_notification_v1_notificationMessage_proto_enumTypes[0]
}

func (x NotificationTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationTypeEnum.Descriptor instead.
func (NotificationTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_notification_v1_notificationMessage_proto_rawDescGZIP(), []int{0}
}

type NotificationTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type         string           `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Template     string           `protobuf:"bytes,3,opt,name=template,proto3" json:"template,omitempty"`
	TemplateVars *structpb.Struct `protobuf:"bytes,4,opt,name=templateVars,proto3" json:"templateVars,omitempty"`
}

func (x *NotificationTemplate) Reset() {
	*x = NotificationTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_v1_notificationMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTemplate) ProtoMessage() {}

func (x *NotificationTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notificationMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationTemplate.ProtoReflect.Descriptor instead.
func (*NotificationTemplate) Descriptor() ([]byte, []int) {
	return file_notification_v1_notificationMessage_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationTemplate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NotificationTemplate) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NotificationTemplate) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

func (x *NotificationTemplate) GetTemplateVars() *structpb.Struct {
	if x != nil {
		return x.TemplateVars
	}
	return nil
}

type SendOtpEmailNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Otp              string               `protobuf:"bytes,1,opt,name=otp,proto3" json:"otp,omitempty"`
	Email            string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	NotificationType NotificationTypeEnum `protobuf:"varint,3,opt,name=notificationType,proto3,enum=notification_v1.NotificationTypeEnum" json:"notificationType,omitempty"`
}

func (x *SendOtpEmailNotificationRequest) Reset() {
	*x = SendOtpEmailNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_v1_notificationMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOtpEmailNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOtpEmailNotificationRequest) ProtoMessage() {}

func (x *SendOtpEmailNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notificationMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOtpEmailNotificationRequest.ProtoReflect.Descriptor instead.
func (*SendOtpEmailNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notification_v1_notificationMessage_proto_rawDescGZIP(), []int{1}
}

func (x *SendOtpEmailNotificationRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *SendOtpEmailNotificationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendOtpEmailNotificationRequest) GetNotificationType() NotificationTypeEnum {
	if x != nil {
		return x.NotificationType
	}
	return NotificationTypeEnum_NOTIFICATION_EMAIL_USER_OTP
}

type SendEmailPaymentOrderCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email            string               `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Payment          *v1.Payment          `protobuf:"bytes,2,opt,name=payment,proto3" json:"payment,omitempty"`
	NotificationType NotificationTypeEnum `protobuf:"varint,3,opt,name=notificationType,proto3,enum=notification_v1.NotificationTypeEnum" json:"notificationType,omitempty"`
}

func (x *SendEmailPaymentOrderCreateRequest) Reset() {
	*x = SendEmailPaymentOrderCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_v1_notificationMessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailPaymentOrderCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailPaymentOrderCreateRequest) ProtoMessage() {}

func (x *SendEmailPaymentOrderCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notificationMessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailPaymentOrderCreateRequest.ProtoReflect.Descriptor instead.
func (*SendEmailPaymentOrderCreateRequest) Descriptor() ([]byte, []int) {
	return file_notification_v1_notificationMessage_proto_rawDescGZIP(), []int{2}
}

func (x *SendEmailPaymentOrderCreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendEmailPaymentOrderCreateRequest) GetPayment() *v1.Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *SendEmailPaymentOrderCreateRequest) GetNotificationType() NotificationTypeEnum {
	if x != nil {
		return x.NotificationType
	}
	return NotificationTypeEnum_NOTIFICATION_EMAIL_USER_OTP
}

var File_notification_v1_notificationMessage_proto protoreflect.FileDescriptor

var file_notification_v1_notificationMessage_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a,
	0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x1f, 0x53,
	0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x51, 0x0a, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x25, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x22, 0x53, 0x65,
	0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x51, 0x0a, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x25, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x65, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x1f, 0x0a, 0x1b, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4f, 0x54, 0x50, 0x10,
	0x00, 0x12, 0x2c, 0x0a, 0x28, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_v1_notificationMessage_proto_rawDescOnce sync.Once
	file_notification_v1_notificationMessage_proto_rawDescData = file_notification_v1_notificationMessage_proto_rawDesc
)

func file_notification_v1_notificationMessage_proto_rawDescGZIP() []byte {
	file_notification_v1_notificationMessage_proto_rawDescOnce.Do(func() {
		file_notification_v1_notificationMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_v1_notificationMessage_proto_rawDescData)
	})
	return file_notification_v1_notificationMessage_proto_rawDescData
}

var file_notification_v1_notificationMessage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notification_v1_notificationMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_notification_v1_notificationMessage_proto_goTypes = []interface{}{
	(NotificationTypeEnum)(0),                  // 0: notification_v1.NotificationTypeEnum
	(*NotificationTemplate)(nil),               // 1: notification_v1.NotificationTemplate
	(*SendOtpEmailNotificationRequest)(nil),    // 2: notification_v1.SendOtpEmailNotificationRequest
	(*SendEmailPaymentOrderCreateRequest)(nil), // 3: notification_v1.SendEmailPaymentOrderCreateRequest
	(*structpb.Struct)(nil),                    // 4: google.protobuf.Struct
	(*v1.Payment)(nil),                         // 5: payment_v1.Payment
}
var file_notification_v1_notificationMessage_proto_depIdxs = []int32{
	4, // 0: notification_v1.NotificationTemplate.templateVars:type_name -> google.protobuf.Struct
	0, // 1: notification_v1.SendOtpEmailNotificationRequest.notificationType:type_name -> notification_v1.NotificationTypeEnum
	5, // 2: notification_v1.SendEmailPaymentOrderCreateRequest.payment:type_name -> payment_v1.Payment
	0, // 3: notification_v1.SendEmailPaymentOrderCreateRequest.notificationType:type_name -> notification_v1.NotificationTypeEnum
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_notification_v1_notificationMessage_proto_init() }
func file_notification_v1_notificationMessage_proto_init() {
	if File_notification_v1_notificationMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_v1_notificationMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationTemplate); i {
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
		file_notification_v1_notificationMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOtpEmailNotificationRequest); i {
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
		file_notification_v1_notificationMessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailPaymentOrderCreateRequest); i {
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
			RawDescriptor: file_notification_v1_notificationMessage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notification_v1_notificationMessage_proto_goTypes,
		DependencyIndexes: file_notification_v1_notificationMessage_proto_depIdxs,
		EnumInfos:         file_notification_v1_notificationMessage_proto_enumTypes,
		MessageInfos:      file_notification_v1_notificationMessage_proto_msgTypes,
	}.Build()
	File_notification_v1_notificationMessage_proto = out.File
	file_notification_v1_notificationMessage_proto_rawDesc = nil
	file_notification_v1_notificationMessage_proto_goTypes = nil
	file_notification_v1_notificationMessage_proto_depIdxs = nil
}
