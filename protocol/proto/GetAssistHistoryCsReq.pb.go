// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetAssistHistoryCsReq.proto

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

type GetAssistHistoryCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAssistHistoryCsReq) Reset() {
	*x = GetAssistHistoryCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetAssistHistoryCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssistHistoryCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssistHistoryCsReq) ProtoMessage() {}

func (x *GetAssistHistoryCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetAssistHistoryCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssistHistoryCsReq.ProtoReflect.Descriptor instead.
func (*GetAssistHistoryCsReq) Descriptor() ([]byte, []int) {
	return file_GetAssistHistoryCsReq_proto_rawDescGZIP(), []int{0}
}

var File_GetAssistHistoryCsReq_proto protoreflect.FileDescriptor

var file_GetAssistHistoryCsReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x43, 0x73, 0x52, 0x65, 0x71, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetAssistHistoryCsReq_proto_rawDescOnce sync.Once
	file_GetAssistHistoryCsReq_proto_rawDescData = file_GetAssistHistoryCsReq_proto_rawDesc
)

func file_GetAssistHistoryCsReq_proto_rawDescGZIP() []byte {
	file_GetAssistHistoryCsReq_proto_rawDescOnce.Do(func() {
		file_GetAssistHistoryCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetAssistHistoryCsReq_proto_rawDescData)
	})
	return file_GetAssistHistoryCsReq_proto_rawDescData
}

var file_GetAssistHistoryCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetAssistHistoryCsReq_proto_goTypes = []interface{}{
	(*GetAssistHistoryCsReq)(nil), // 0: GetAssistHistoryCsReq
}
var file_GetAssistHistoryCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetAssistHistoryCsReq_proto_init() }
func file_GetAssistHistoryCsReq_proto_init() {
	if File_GetAssistHistoryCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetAssistHistoryCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssistHistoryCsReq); i {
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
			RawDescriptor: file_GetAssistHistoryCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetAssistHistoryCsReq_proto_goTypes,
		DependencyIndexes: file_GetAssistHistoryCsReq_proto_depIdxs,
		MessageInfos:      file_GetAssistHistoryCsReq_proto_msgTypes,
	}.Build()
	File_GetAssistHistoryCsReq_proto = out.File
	file_GetAssistHistoryCsReq_proto_rawDesc = nil
	file_GetAssistHistoryCsReq_proto_goTypes = nil
	file_GetAssistHistoryCsReq_proto_depIdxs = nil
}
