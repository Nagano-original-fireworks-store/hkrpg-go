// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetUpdatedArchiveDataScRsp.proto

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

type GetUpdatedArchiveDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArchiveInfo *ArchiveData `protobuf:"bytes,14,opt,name=archive_info,json=archiveInfo,proto3" json:"archive_info,omitempty"`
	Retcode     uint32       `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetUpdatedArchiveDataScRsp) Reset() {
	*x = GetUpdatedArchiveDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetUpdatedArchiveDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpdatedArchiveDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdatedArchiveDataScRsp) ProtoMessage() {}

func (x *GetUpdatedArchiveDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetUpdatedArchiveDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdatedArchiveDataScRsp.ProtoReflect.Descriptor instead.
func (*GetUpdatedArchiveDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetUpdatedArchiveDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetUpdatedArchiveDataScRsp) GetArchiveInfo() *ArchiveData {
	if x != nil {
		return x.ArchiveInfo
	}
	return nil
}

func (x *GetUpdatedArchiveDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetUpdatedArchiveDataScRsp_proto protoreflect.FileDescriptor

var file_GetUpdatedArchiveDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GetUpdatedArchiveDataScRsp_proto_rawDescOnce sync.Once
	file_GetUpdatedArchiveDataScRsp_proto_rawDescData = file_GetUpdatedArchiveDataScRsp_proto_rawDesc
)

func file_GetUpdatedArchiveDataScRsp_proto_rawDescGZIP() []byte {
	file_GetUpdatedArchiveDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetUpdatedArchiveDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetUpdatedArchiveDataScRsp_proto_rawDescData)
	})
	return file_GetUpdatedArchiveDataScRsp_proto_rawDescData
}

var file_GetUpdatedArchiveDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetUpdatedArchiveDataScRsp_proto_goTypes = []interface{}{
	(*GetUpdatedArchiveDataScRsp)(nil), // 0: GetUpdatedArchiveDataScRsp
	(*ArchiveData)(nil),                // 1: ArchiveData
}
var file_GetUpdatedArchiveDataScRsp_proto_depIdxs = []int32{
	1, // 0: GetUpdatedArchiveDataScRsp.archive_info:type_name -> ArchiveData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetUpdatedArchiveDataScRsp_proto_init() }
func file_GetUpdatedArchiveDataScRsp_proto_init() {
	if File_GetUpdatedArchiveDataScRsp_proto != nil {
		return
	}
	file_ArchiveData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetUpdatedArchiveDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpdatedArchiveDataScRsp); i {
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
			RawDescriptor: file_GetUpdatedArchiveDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetUpdatedArchiveDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetUpdatedArchiveDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetUpdatedArchiveDataScRsp_proto_msgTypes,
	}.Build()
	File_GetUpdatedArchiveDataScRsp_proto = out.File
	file_GetUpdatedArchiveDataScRsp_proto_rawDesc = nil
	file_GetUpdatedArchiveDataScRsp_proto_goTypes = nil
	file_GetUpdatedArchiveDataScRsp_proto_depIdxs = nil
}
