// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HNGDNBIAAMC.proto

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

type HNGDNBIAAMC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OHOEGOAECGK uint32 `protobuf:"varint,12,opt,name=OHOEGOAECGK,proto3" json:"OHOEGOAECGK,omitempty"`
	FBPPKGCGILB bool   `protobuf:"varint,7,opt,name=FBPPKGCGILB,proto3" json:"FBPPKGCGILB,omitempty"`
}

func (x *HNGDNBIAAMC) Reset() {
	*x = HNGDNBIAAMC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HNGDNBIAAMC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HNGDNBIAAMC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HNGDNBIAAMC) ProtoMessage() {}

func (x *HNGDNBIAAMC) ProtoReflect() protoreflect.Message {
	mi := &file_HNGDNBIAAMC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HNGDNBIAAMC.ProtoReflect.Descriptor instead.
func (*HNGDNBIAAMC) Descriptor() ([]byte, []int) {
	return file_HNGDNBIAAMC_proto_rawDescGZIP(), []int{0}
}

func (x *HNGDNBIAAMC) GetOHOEGOAECGK() uint32 {
	if x != nil {
		return x.OHOEGOAECGK
	}
	return 0
}

func (x *HNGDNBIAAMC) GetFBPPKGCGILB() bool {
	if x != nil {
		return x.FBPPKGCGILB
	}
	return false
}

var File_HNGDNBIAAMC_proto protoreflect.FileDescriptor

var file_HNGDNBIAAMC_proto_rawDesc = []byte{
	0x0a, 0x11, 0x48, 0x4e, 0x47, 0x44, 0x4e, 0x42, 0x49, 0x41, 0x41, 0x4d, 0x43, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0b, 0x48, 0x4e, 0x47, 0x44, 0x4e, 0x42, 0x49, 0x41, 0x41,
	0x4d, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x48, 0x4f, 0x45, 0x47, 0x4f, 0x41, 0x45, 0x43, 0x47,
	0x4b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x48, 0x4f, 0x45, 0x47, 0x4f, 0x41,
	0x45, 0x43, 0x47, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x42, 0x50, 0x50, 0x4b, 0x47, 0x43, 0x47,
	0x49, 0x4c, 0x42, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x46, 0x42, 0x50, 0x50, 0x4b,
	0x47, 0x43, 0x47, 0x49, 0x4c, 0x42, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HNGDNBIAAMC_proto_rawDescOnce sync.Once
	file_HNGDNBIAAMC_proto_rawDescData = file_HNGDNBIAAMC_proto_rawDesc
)

func file_HNGDNBIAAMC_proto_rawDescGZIP() []byte {
	file_HNGDNBIAAMC_proto_rawDescOnce.Do(func() {
		file_HNGDNBIAAMC_proto_rawDescData = protoimpl.X.CompressGZIP(file_HNGDNBIAAMC_proto_rawDescData)
	})
	return file_HNGDNBIAAMC_proto_rawDescData
}

var file_HNGDNBIAAMC_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HNGDNBIAAMC_proto_goTypes = []interface{}{
	(*HNGDNBIAAMC)(nil), // 0: HNGDNBIAAMC
}
var file_HNGDNBIAAMC_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HNGDNBIAAMC_proto_init() }
func file_HNGDNBIAAMC_proto_init() {
	if File_HNGDNBIAAMC_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HNGDNBIAAMC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HNGDNBIAAMC); i {
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
			RawDescriptor: file_HNGDNBIAAMC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HNGDNBIAAMC_proto_goTypes,
		DependencyIndexes: file_HNGDNBIAAMC_proto_depIdxs,
		MessageInfos:      file_HNGDNBIAAMC_proto_msgTypes,
	}.Build()
	File_HNGDNBIAAMC_proto = out.File
	file_HNGDNBIAAMC_proto_rawDesc = nil
	file_HNGDNBIAAMC_proto_goTypes = nil
	file_HNGDNBIAAMC_proto_depIdxs = nil
}
