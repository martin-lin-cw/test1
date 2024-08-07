// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: test2/v1/test2.proto

package test2pb

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

type Hello2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Hello2Request) Reset() {
	*x = Hello2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test2_v1_test2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello2Request) ProtoMessage() {}

func (x *Hello2Request) ProtoReflect() protoreflect.Message {
	mi := &file_test2_v1_test2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello2Request.ProtoReflect.Descriptor instead.
func (*Hello2Request) Descriptor() ([]byte, []int) {
	return file_test2_v1_test2_proto_rawDescGZIP(), []int{0}
}

type Hello2Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Hello2Response) Reset() {
	*x = Hello2Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test2_v1_test2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello2Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello2Response) ProtoMessage() {}

func (x *Hello2Response) ProtoReflect() protoreflect.Message {
	mi := &file_test2_v1_test2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello2Response.ProtoReflect.Descriptor instead.
func (*Hello2Response) Descriptor() ([]byte, []int) {
	return file_test2_v1_test2_proto_rawDescGZIP(), []int{1}
}

func (x *Hello2Response) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_test2_v1_test2_proto protoreflect.FileDescriptor

var file_test2_v1_test2_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x65, 0x73, 0x74, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x32,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x65, 0x73, 0x74, 0x32, 0x22, 0x0f, 0x0a,
	0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28,
	0x0a, 0x0e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x47, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74,
	0x32, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x32, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x32, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x32,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x7a, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x32, 0x42, 0x0a,
	0x54, 0x65, 0x73, 0x74, 0x32, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2d,
	0x6c, 0x69, 0x6e, 0x2d, 0x63, 0x77, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x32,
	0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x32, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x54, 0x58,
	0x58, 0xaa, 0x02, 0x05, 0x54, 0x65, 0x73, 0x74, 0x32, 0xca, 0x02, 0x05, 0x54, 0x65, 0x73, 0x74,
	0x32, 0xe2, 0x02, 0x11, 0x54, 0x65, 0x73, 0x74, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x54, 0x65, 0x73, 0x74, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test2_v1_test2_proto_rawDescOnce sync.Once
	file_test2_v1_test2_proto_rawDescData = file_test2_v1_test2_proto_rawDesc
)

func file_test2_v1_test2_proto_rawDescGZIP() []byte {
	file_test2_v1_test2_proto_rawDescOnce.Do(func() {
		file_test2_v1_test2_proto_rawDescData = protoimpl.X.CompressGZIP(file_test2_v1_test2_proto_rawDescData)
	})
	return file_test2_v1_test2_proto_rawDescData
}

var file_test2_v1_test2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test2_v1_test2_proto_goTypes = []any{
	(*Hello2Request)(nil),  // 0: test2.Hello2Request
	(*Hello2Response)(nil), // 1: test2.Hello2Response
}
var file_test2_v1_test2_proto_depIdxs = []int32{
	0, // 0: test2.Test2Service.Hello2:input_type -> test2.Hello2Request
	1, // 1: test2.Test2Service.Hello2:output_type -> test2.Hello2Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test2_v1_test2_proto_init() }
func file_test2_v1_test2_proto_init() {
	if File_test2_v1_test2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test2_v1_test2_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Hello2Request); i {
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
		file_test2_v1_test2_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Hello2Response); i {
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
			RawDescriptor: file_test2_v1_test2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test2_v1_test2_proto_goTypes,
		DependencyIndexes: file_test2_v1_test2_proto_depIdxs,
		MessageInfos:      file_test2_v1_test2_proto_msgTypes,
	}.Build()
	File_test2_v1_test2_proto = out.File
	file_test2_v1_test2_proto_rawDesc = nil
	file_test2_v1_test2_proto_goTypes = nil
	file_test2_v1_test2_proto_depIdxs = nil
}
