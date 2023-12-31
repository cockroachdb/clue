// Code generated with goa v3.9.1, DO NOT EDIT.
//
// test protocol buffer definition
//
// Command:
// $ goa gen goa.design/clue/internal/testsvc/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: goadesign_goagen_test.proto

package testpb

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

type GrpcMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String operand
	S string `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	// Int operand
	I int32 `protobuf:"zigzag32,2,opt,name=i,proto3" json:"i,omitempty"`
}

func (x *GrpcMethodRequest) Reset() {
	*x = GrpcMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcMethodRequest) ProtoMessage() {}

func (x *GrpcMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcMethodRequest.ProtoReflect.Descriptor instead.
func (*GrpcMethodRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_test_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcMethodRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *GrpcMethodRequest) GetI() int32 {
	if x != nil {
		return x.I
	}
	return 0
}

type GrpcMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String operand
	S string `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	// Int operand
	I int32 `protobuf:"zigzag32,2,opt,name=i,proto3" json:"i,omitempty"`
}

func (x *GrpcMethodResponse) Reset() {
	*x = GrpcMethodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcMethodResponse) ProtoMessage() {}

func (x *GrpcMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcMethodResponse.ProtoReflect.Descriptor instead.
func (*GrpcMethodResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_test_proto_rawDescGZIP(), []int{1}
}

func (x *GrpcMethodResponse) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *GrpcMethodResponse) GetI() int32 {
	if x != nil {
		return x.I
	}
	return 0
}

type GrpcStreamStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String operand
	S string `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	// Int operand
	I int32 `protobuf:"zigzag32,2,opt,name=i,proto3" json:"i,omitempty"`
}

func (x *GrpcStreamStreamingRequest) Reset() {
	*x = GrpcStreamStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcStreamStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcStreamStreamingRequest) ProtoMessage() {}

func (x *GrpcStreamStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcStreamStreamingRequest.ProtoReflect.Descriptor instead.
func (*GrpcStreamStreamingRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_test_proto_rawDescGZIP(), []int{2}
}

func (x *GrpcStreamStreamingRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *GrpcStreamStreamingRequest) GetI() int32 {
	if x != nil {
		return x.I
	}
	return 0
}

type GrpcStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String operand
	S string `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	// Int operand
	I int32 `protobuf:"zigzag32,2,opt,name=i,proto3" json:"i,omitempty"`
}

func (x *GrpcStreamResponse) Reset() {
	*x = GrpcStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcStreamResponse) ProtoMessage() {}

func (x *GrpcStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcStreamResponse.ProtoReflect.Descriptor instead.
func (*GrpcStreamResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_test_proto_rawDescGZIP(), []int{3}
}

func (x *GrpcStreamResponse) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *GrpcStreamResponse) GetI() int32 {
	if x != nil {
		return x.I
	}
	return 0
}

var File_goadesign_goagen_test_proto protoreflect.FileDescriptor

var file_goadesign_goagen_test_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74,
	0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x11, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x01, 0x69, 0x22, 0x30, 0x0a, 0x12, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x69, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x01, 0x69, 0x22, 0x38, 0x0a, 0x1a, 0x47, 0x72, 0x70, 0x63, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x01, 0x69,
	0x22, 0x30, 0x0a, 0x12, 0x47, 0x72, 0x70, 0x63, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x01, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x01, 0x69, 0x32, 0x95, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x47,
	0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0a,
	0x47, 0x72, 0x70, 0x63, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_test_proto_rawDescOnce sync.Once
	file_goadesign_goagen_test_proto_rawDescData = file_goadesign_goagen_test_proto_rawDesc
)

func file_goadesign_goagen_test_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_test_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_test_proto_rawDescData)
	})
	return file_goadesign_goagen_test_proto_rawDescData
}

var file_goadesign_goagen_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_goadesign_goagen_test_proto_goTypes = []interface{}{
	(*GrpcMethodRequest)(nil),          // 0: test.GrpcMethodRequest
	(*GrpcMethodResponse)(nil),         // 1: test.GrpcMethodResponse
	(*GrpcStreamStreamingRequest)(nil), // 2: test.GrpcStreamStreamingRequest
	(*GrpcStreamResponse)(nil),         // 3: test.GrpcStreamResponse
}
var file_goadesign_goagen_test_proto_depIdxs = []int32{
	0, // 0: test.Test.GrpcMethod:input_type -> test.GrpcMethodRequest
	2, // 1: test.Test.GrpcStream:input_type -> test.GrpcStreamStreamingRequest
	1, // 2: test.Test.GrpcMethod:output_type -> test.GrpcMethodResponse
	3, // 3: test.Test.GrpcStream:output_type -> test.GrpcStreamResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_test_proto_init() }
func file_goadesign_goagen_test_proto_init() {
	if File_goadesign_goagen_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcMethodRequest); i {
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
		file_goadesign_goagen_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcMethodResponse); i {
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
		file_goadesign_goagen_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcStreamStreamingRequest); i {
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
		file_goadesign_goagen_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcStreamResponse); i {
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
			RawDescriptor: file_goadesign_goagen_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_test_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_test_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_test_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_test_proto = out.File
	file_goadesign_goagen_test_proto_rawDesc = nil
	file_goadesign_goagen_test_proto_goTypes = nil
	file_goadesign_goagen_test_proto_depIdxs = nil
}
