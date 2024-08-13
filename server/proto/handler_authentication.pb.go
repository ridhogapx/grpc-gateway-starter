// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/handler_authentication.proto

package proto

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

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email          string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username       string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Password       string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	ProfilePicture string `protobuf:"bytes,5,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	AllowNsfw      bool   `protobuf:"varint,6,opt,name=allow_nsfw,json=allowNsfw,proto3" json:"allow_nsfw,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_proto_handler_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignUpRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignUpRequest) GetProfilePicture() string {
	if x != nil {
		return x.ProfilePicture
	}
	return ""
}

func (x *SignUpRequest) GetAllowNsfw() bool {
	if x != nil {
		return x.AllowNsfw
	}
	return false
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string                  `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *SignUpResponse_TSignUp `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_proto_handler_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SignUpResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SignUpResponse) GetData() *SignUpResponse_TSignUp {
	if x != nil {
		return x.Data
	}
	return nil
}

type SignUpResponse_TSignUp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *SignUpResponse_TSignUp) Reset() {
	*x = SignUpResponse_TSignUp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse_TSignUp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse_TSignUp) ProtoMessage() {}

func (x *SignUpResponse_TSignUp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse_TSignUp.ProtoReflect.Descriptor instead.
func (*SignUpResponse_TSignUp) Descriptor() ([]byte, []int) {
	return file_proto_handler_authentication_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SignUpResponse_TSignUp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SignUpResponse_TSignUp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SignUpResponse_TSignUp) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_proto_handler_authentication_proto protoreflect.FileDescriptor

var file_proto_handler_authentication_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x22, 0xb9, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6e, 0x73, 0x66, 0x77, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4e, 0x73, 0x66, 0x77,
	0x22, 0xd2, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x54, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x56, 0x0a,
	0x07, 0x54, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_handler_authentication_proto_rawDescOnce sync.Once
	file_proto_handler_authentication_proto_rawDescData = file_proto_handler_authentication_proto_rawDesc
)

func file_proto_handler_authentication_proto_rawDescGZIP() []byte {
	file_proto_handler_authentication_proto_rawDescOnce.Do(func() {
		file_proto_handler_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_handler_authentication_proto_rawDescData)
	})
	return file_proto_handler_authentication_proto_rawDescData
}

var file_proto_handler_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_handler_authentication_proto_goTypes = []interface{}{
	(*SignUpRequest)(nil),          // 0: api_service.v1.SignUpRequest
	(*SignUpResponse)(nil),         // 1: api_service.v1.SignUpResponse
	(*SignUpResponse_TSignUp)(nil), // 2: api_service.v1.SignUpResponse.TSignUp
}
var file_proto_handler_authentication_proto_depIdxs = []int32{
	2, // 0: api_service.v1.SignUpResponse.data:type_name -> api_service.v1.SignUpResponse.TSignUp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_handler_authentication_proto_init() }
func file_proto_handler_authentication_proto_init() {
	if File_proto_handler_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_handler_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
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
		file_proto_handler_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
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
		file_proto_handler_authentication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse_TSignUp); i {
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
			RawDescriptor: file_proto_handler_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_handler_authentication_proto_goTypes,
		DependencyIndexes: file_proto_handler_authentication_proto_depIdxs,
		MessageInfos:      file_proto_handler_authentication_proto_msgTypes,
	}.Build()
	File_proto_handler_authentication_proto = out.File
	file_proto_handler_authentication_proto_rawDesc = nil
	file_proto_handler_authentication_proto_goTypes = nil
	file_proto_handler_authentication_proto_depIdxs = nil
}