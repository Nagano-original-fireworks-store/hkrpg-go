// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GEPGLOBCIKM.proto

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

type GEPGLOBCIKM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BNHHDJKDEBI uint32 `protobuf:"varint,2,opt,name=BNHHDJKDEBI,proto3" json:"BNHHDJKDEBI,omitempty"`
}

func (x *GEPGLOBCIKM) Reset() {
	*x = GEPGLOBCIKM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GEPGLOBCIKM_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GEPGLOBCIKM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GEPGLOBCIKM) ProtoMessage() {}

func (x *GEPGLOBCIKM) ProtoReflect() protoreflect.Message {
	mi := &file_GEPGLOBCIKM_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GEPGLOBCIKM.ProtoReflect.Descriptor instead.
func (*GEPGLOBCIKM) Descriptor() ([]byte, []int) {
	return file_GEPGLOBCIKM_proto_rawDescGZIP(), []int{0}
}

func (x *GEPGLOBCIKM) GetBNHHDJKDEBI() uint32 {
	if x != nil {
		return x.BNHHDJKDEBI
	}
	return 0
}

var File_GEPGLOBCIKM_proto protoreflect.FileDescriptor

var file_GEPGLOBCIKM_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x45, 0x50, 0x47, 0x4c, 0x4f, 0x42, 0x43, 0x49, 0x4b, 0x4d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0b, 0x47, 0x45, 0x50, 0x47, 0x4c, 0x4f, 0x42, 0x43, 0x49,
	0x4b, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x4e, 0x48, 0x48, 0x44, 0x4a, 0x4b, 0x44, 0x45, 0x42,
	0x49, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x4e, 0x48, 0x48, 0x44, 0x4a, 0x4b,
	0x44, 0x45, 0x42, 0x49, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GEPGLOBCIKM_proto_rawDescOnce sync.Once
	file_GEPGLOBCIKM_proto_rawDescData = file_GEPGLOBCIKM_proto_rawDesc
)

func file_GEPGLOBCIKM_proto_rawDescGZIP() []byte {
	file_GEPGLOBCIKM_proto_rawDescOnce.Do(func() {
		file_GEPGLOBCIKM_proto_rawDescData = protoimpl.X.CompressGZIP(file_GEPGLOBCIKM_proto_rawDescData)
	})
	return file_GEPGLOBCIKM_proto_rawDescData
}

var file_GEPGLOBCIKM_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GEPGLOBCIKM_proto_goTypes = []interface{}{
	(*GEPGLOBCIKM)(nil), // 0: GEPGLOBCIKM
}
var file_GEPGLOBCIKM_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GEPGLOBCIKM_proto_init() }
func file_GEPGLOBCIKM_proto_init() {
	if File_GEPGLOBCIKM_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GEPGLOBCIKM_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GEPGLOBCIKM); i {
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
			RawDescriptor: file_GEPGLOBCIKM_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GEPGLOBCIKM_proto_goTypes,
		DependencyIndexes: file_GEPGLOBCIKM_proto_depIdxs,
		MessageInfos:      file_GEPGLOBCIKM_proto_msgTypes,
	}.Build()
	File_GEPGLOBCIKM_proto = out.File
	file_GEPGLOBCIKM_proto_rawDesc = nil
	file_GEPGLOBCIKM_proto_goTypes = nil
	file_GEPGLOBCIKM_proto_depIdxs = nil
}
