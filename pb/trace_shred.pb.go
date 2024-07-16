// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: trace_shred.proto

package jito_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TraceShred struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// source region, one of: https://jito-labs.gitbook.io/mev/systems/connecting/mainnet
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// timestamp of creation
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// monotonically increases, resets upon service restart
	SeqNum uint32 `protobuf:"varint,3,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
}

func (x *TraceShred) Reset() {
	*x = TraceShred{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trace_shred_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceShred) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceShred) ProtoMessage() {}

func (x *TraceShred) ProtoReflect() protoreflect.Message {
	mi := &file_trace_shred_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceShred.ProtoReflect.Descriptor instead.
func (*TraceShred) Descriptor() ([]byte, []int) {
	return file_trace_shred_proto_rawDescGZIP(), []int{0}
}

func (x *TraceShred) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *TraceShred) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TraceShred) GetSeqNum() uint32 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

var File_trace_shred_proto protoreflect.FileDescriptor

var file_trace_shred_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x68, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x68, 0x72, 0x65, 0x64,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x78, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x68, 0x72, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x42, 0x0e, 0x5a, 0x0c, 0x6a,
	0x69, 0x74, 0x6f, 0x2f, 0x6a, 0x69, 0x74, 0x6f, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_trace_shred_proto_rawDescOnce sync.Once
	file_trace_shred_proto_rawDescData = file_trace_shred_proto_rawDesc
)

func file_trace_shred_proto_rawDescGZIP() []byte {
	file_trace_shred_proto_rawDescOnce.Do(func() {
		file_trace_shred_proto_rawDescData = protoimpl.X.CompressGZIP(file_trace_shred_proto_rawDescData)
	})
	return file_trace_shred_proto_rawDescData
}

var file_trace_shred_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_trace_shred_proto_goTypes = []interface{}{
	(*TraceShred)(nil),            // 0: trace_shred.TraceShred
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_trace_shred_proto_depIdxs = []int32{
	1, // 0: trace_shred.TraceShred.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trace_shred_proto_init() }
func file_trace_shred_proto_init() {
	if File_trace_shred_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trace_shred_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceShred); i {
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
			RawDescriptor: file_trace_shred_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trace_shred_proto_goTypes,
		DependencyIndexes: file_trace_shred_proto_depIdxs,
		MessageInfos:      file_trace_shred_proto_msgTypes,
	}.Build()
	File_trace_shred_proto = out.File
	file_trace_shred_proto_rawDesc = nil
	file_trace_shred_proto_goTypes = nil
	file_trace_shred_proto_depIdxs = nil
}
