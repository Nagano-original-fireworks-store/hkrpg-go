// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: InteractTreasureDungeonGridScRsp.proto

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

type InteractTreasureDungeonGridScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FECNFEBKCDI *TreasureDungeonLevel `protobuf:"bytes,10,opt,name=FECNFEBKCDI,proto3" json:"FECNFEBKCDI,omitempty"`
	Retcode     uint32                `protobuf:"varint,11,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *InteractTreasureDungeonGridScRsp) Reset() {
	*x = InteractTreasureDungeonGridScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InteractTreasureDungeonGridScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractTreasureDungeonGridScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractTreasureDungeonGridScRsp) ProtoMessage() {}

func (x *InteractTreasureDungeonGridScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_InteractTreasureDungeonGridScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractTreasureDungeonGridScRsp.ProtoReflect.Descriptor instead.
func (*InteractTreasureDungeonGridScRsp) Descriptor() ([]byte, []int) {
	return file_InteractTreasureDungeonGridScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *InteractTreasureDungeonGridScRsp) GetFECNFEBKCDI() *TreasureDungeonLevel {
	if x != nil {
		return x.FECNFEBKCDI
	}
	return nil
}

func (x *InteractTreasureDungeonGridScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_InteractTreasureDungeonGridScRsp_proto protoreflect.FileDescriptor

var file_InteractTreasureDungeonGridScRsp_proto_rawDesc = []byte{
	0x0a, 0x26, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x47, 0x72, 0x69, 0x64, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x47,
	0x72, 0x69, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x0b, 0x46, 0x45, 0x43, 0x4e,
	0x46, 0x45, 0x42, 0x4b, 0x43, 0x44, 0x49, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x0b, 0x46, 0x45, 0x43, 0x4e, 0x46, 0x45, 0x42, 0x4b, 0x43, 0x44,
	0x49, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InteractTreasureDungeonGridScRsp_proto_rawDescOnce sync.Once
	file_InteractTreasureDungeonGridScRsp_proto_rawDescData = file_InteractTreasureDungeonGridScRsp_proto_rawDesc
)

func file_InteractTreasureDungeonGridScRsp_proto_rawDescGZIP() []byte {
	file_InteractTreasureDungeonGridScRsp_proto_rawDescOnce.Do(func() {
		file_InteractTreasureDungeonGridScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_InteractTreasureDungeonGridScRsp_proto_rawDescData)
	})
	return file_InteractTreasureDungeonGridScRsp_proto_rawDescData
}

var file_InteractTreasureDungeonGridScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InteractTreasureDungeonGridScRsp_proto_goTypes = []interface{}{
	(*InteractTreasureDungeonGridScRsp)(nil), // 0: InteractTreasureDungeonGridScRsp
	(*TreasureDungeonLevel)(nil),             // 1: TreasureDungeonLevel
}
var file_InteractTreasureDungeonGridScRsp_proto_depIdxs = []int32{
	1, // 0: InteractTreasureDungeonGridScRsp.FECNFEBKCDI:type_name -> TreasureDungeonLevel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_InteractTreasureDungeonGridScRsp_proto_init() }
func file_InteractTreasureDungeonGridScRsp_proto_init() {
	if File_InteractTreasureDungeonGridScRsp_proto != nil {
		return
	}
	file_TreasureDungeonLevel_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InteractTreasureDungeonGridScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InteractTreasureDungeonGridScRsp); i {
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
			RawDescriptor: file_InteractTreasureDungeonGridScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InteractTreasureDungeonGridScRsp_proto_goTypes,
		DependencyIndexes: file_InteractTreasureDungeonGridScRsp_proto_depIdxs,
		MessageInfos:      file_InteractTreasureDungeonGridScRsp_proto_msgTypes,
	}.Build()
	File_InteractTreasureDungeonGridScRsp_proto = out.File
	file_InteractTreasureDungeonGridScRsp_proto_rawDesc = nil
	file_InteractTreasureDungeonGridScRsp_proto_goTypes = nil
	file_InteractTreasureDungeonGridScRsp_proto_depIdxs = nil
}
