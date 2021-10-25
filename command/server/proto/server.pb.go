// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: command/server/proto/server.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PprofRequest_Type int32

const (
	PprofRequest_LOOKUP PprofRequest_Type = 0
	PprofRequest_CPU    PprofRequest_Type = 1
	PprofRequest_TRACE  PprofRequest_Type = 2
)

// Enum value maps for PprofRequest_Type.
var (
	PprofRequest_Type_name = map[int32]string{
		0: "LOOKUP",
		1: "CPU",
		2: "TRACE",
	}
	PprofRequest_Type_value = map[string]int32{
		"LOOKUP": 0,
		"CPU":    1,
		"TRACE":  2,
	}
)

func (x PprofRequest_Type) Enum() *PprofRequest_Type {
	p := new(PprofRequest_Type)
	*p = x
	return p
}

func (x PprofRequest_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PprofRequest_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_command_server_proto_server_proto_enumTypes[0].Descriptor()
}

func (PprofRequest_Type) Type() protoreflect.EnumType {
	return &file_command_server_proto_server_proto_enumTypes[0]
}

func (x PprofRequest_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PprofRequest_Type.Descriptor instead.
func (PprofRequest_Type) EnumDescriptor() ([]byte, []int) {
	return file_command_server_proto_server_proto_rawDescGZIP(), []int{0, 0}
}

type PprofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    PprofRequest_Type `protobuf:"varint,1,opt,name=type,proto3,enum=proto.PprofRequest_Type" json:"type,omitempty"`
	Profile string            `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	Seconds int64             `protobuf:"varint,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
}

func (x *PprofRequest) Reset() {
	*x = PprofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_server_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PprofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PprofRequest) ProtoMessage() {}

func (x *PprofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_server_proto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PprofRequest.ProtoReflect.Descriptor instead.
func (*PprofRequest) Descriptor() ([]byte, []int) {
	return file_command_server_proto_server_proto_rawDescGZIP(), []int{0}
}

func (x *PprofRequest) GetType() PprofRequest_Type {
	if x != nil {
		return x.Type
	}
	return PprofRequest_LOOKUP
}

func (x *PprofRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *PprofRequest) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

type PprofResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string            `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PprofResponse) Reset() {
	*x = PprofResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_server_proto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PprofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PprofResponse) ProtoMessage() {}

func (x *PprofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_command_server_proto_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PprofResponse.ProtoReflect.Descriptor instead.
func (*PprofResponse) Descriptor() ([]byte, []int) {
	return file_command_server_proto_server_proto_rawDescGZIP(), []int{1}
}

func (x *PprofResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *PprofResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

var File_command_server_proto_server_proto protoreflect.FileDescriptor

var file_command_server_proto_server_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0c, 0x50,
	0x70, 0x72, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x26, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x4f, 0x4b, 0x55, 0x50, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52,
	0x41, 0x43, 0x45, 0x10, 0x02, 0x22, 0xa2, 0x01, 0x0a, 0x0d, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x3b, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x70, 0x72, 0x6f, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a,
	0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x39, 0x0a, 0x03, 0x42, 0x6f,
	0x72, 0x12, 0x32, 0x0a, 0x05, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_server_proto_server_proto_rawDescOnce sync.Once
	file_command_server_proto_server_proto_rawDescData = file_command_server_proto_server_proto_rawDesc
)

func file_command_server_proto_server_proto_rawDescGZIP() []byte {
	file_command_server_proto_server_proto_rawDescOnce.Do(func() {
		file_command_server_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_server_proto_server_proto_rawDescData)
	})
	return file_command_server_proto_server_proto_rawDescData
}

var file_command_server_proto_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_command_server_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_command_server_proto_server_proto_goTypes = []interface{}{
	(PprofRequest_Type)(0), // 0: proto.PprofRequest.Type
	(*PprofRequest)(nil),   // 1: proto.PprofRequest
	(*PprofResponse)(nil),  // 2: proto.PprofResponse
	nil,                    // 3: proto.PprofResponse.HeadersEntry
}
var file_command_server_proto_server_proto_depIdxs = []int32{
	0, // 0: proto.PprofRequest.type:type_name -> proto.PprofRequest.Type
	3, // 1: proto.PprofResponse.headers:type_name -> proto.PprofResponse.HeadersEntry
	1, // 2: proto.Bor.Pprof:input_type -> proto.PprofRequest
	2, // 3: proto.Bor.Pprof:output_type -> proto.PprofResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_command_server_proto_server_proto_init() }
func file_command_server_proto_server_proto_init() {
	if File_command_server_proto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_command_server_proto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PprofRequest); i {
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
		file_command_server_proto_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PprofResponse); i {
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
			RawDescriptor: file_command_server_proto_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_command_server_proto_server_proto_goTypes,
		DependencyIndexes: file_command_server_proto_server_proto_depIdxs,
		EnumInfos:         file_command_server_proto_server_proto_enumTypes,
		MessageInfos:      file_command_server_proto_server_proto_msgTypes,
	}.Build()
	File_command_server_proto_server_proto = out.File
	file_command_server_proto_server_proto_rawDesc = nil
	file_command_server_proto_server_proto_goTypes = nil
	file_command_server_proto_server_proto_depIdxs = nil
}