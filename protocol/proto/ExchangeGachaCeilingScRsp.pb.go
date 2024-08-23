// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ExchangeGachaCeilingScRsp.proto

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

type ExchangeGachaCeilingScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarId         uint32        `protobuf:"varint,1,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	TransferItemList *ItemList     `protobuf:"bytes,8,opt,name=transfer_item_list,json=transferItemList,proto3" json:"transfer_item_list,omitempty"`
	GachaType        uint32        `protobuf:"varint,5,opt,name=gacha_type,json=gachaType,proto3" json:"gacha_type,omitempty"`
	Retcode          uint32        `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	GachaCeiling     *GachaCeiling `protobuf:"bytes,9,opt,name=gacha_ceiling,json=gachaCeiling,proto3" json:"gacha_ceiling,omitempty"`
}

func (x *ExchangeGachaCeilingScRsp) Reset() {
	*x = ExchangeGachaCeilingScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExchangeGachaCeilingScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeGachaCeilingScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeGachaCeilingScRsp) ProtoMessage() {}

func (x *ExchangeGachaCeilingScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ExchangeGachaCeilingScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeGachaCeilingScRsp.ProtoReflect.Descriptor instead.
func (*ExchangeGachaCeilingScRsp) Descriptor() ([]byte, []int) {
	return file_ExchangeGachaCeilingScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeGachaCeilingScRsp) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *ExchangeGachaCeilingScRsp) GetTransferItemList() *ItemList {
	if x != nil {
		return x.TransferItemList
	}
	return nil
}

func (x *ExchangeGachaCeilingScRsp) GetGachaType() uint32 {
	if x != nil {
		return x.GachaType
	}
	return 0
}

func (x *ExchangeGachaCeilingScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ExchangeGachaCeilingScRsp) GetGachaCeiling() *GachaCeiling {
	if x != nil {
		return x.GachaCeiling
	}
	return nil
}

var File_ExchangeGachaCeilingScRsp_proto protoreflect.FileDescriptor

var file_ExchangeGachaCeilingScRsp_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x47, 0x61, 0x63, 0x68, 0x61, 0x43,
	0x65, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x47, 0x61, 0x63, 0x68, 0x61, 0x43, 0x65, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x19, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x47, 0x61, 0x63, 0x68, 0x61, 0x43, 0x65, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64,
	0x12, 0x37, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x63,
	0x68, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67,
	0x61, 0x63, 0x68, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x63, 0x65, 0x69, 0x6c,
	0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x47, 0x61, 0x63, 0x68,
	0x61, 0x43, 0x65, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x67, 0x61, 0x63, 0x68, 0x61, 0x43,
	0x65, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ExchangeGachaCeilingScRsp_proto_rawDescOnce sync.Once
	file_ExchangeGachaCeilingScRsp_proto_rawDescData = file_ExchangeGachaCeilingScRsp_proto_rawDesc
)

func file_ExchangeGachaCeilingScRsp_proto_rawDescGZIP() []byte {
	file_ExchangeGachaCeilingScRsp_proto_rawDescOnce.Do(func() {
		file_ExchangeGachaCeilingScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExchangeGachaCeilingScRsp_proto_rawDescData)
	})
	return file_ExchangeGachaCeilingScRsp_proto_rawDescData
}

var file_ExchangeGachaCeilingScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ExchangeGachaCeilingScRsp_proto_goTypes = []interface{}{
	(*ExchangeGachaCeilingScRsp)(nil), // 0: ExchangeGachaCeilingScRsp
	(*ItemList)(nil),                  // 1: ItemList
	(*GachaCeiling)(nil),              // 2: GachaCeiling
}
var file_ExchangeGachaCeilingScRsp_proto_depIdxs = []int32{
	1, // 0: ExchangeGachaCeilingScRsp.transfer_item_list:type_name -> ItemList
	2, // 1: ExchangeGachaCeilingScRsp.gacha_ceiling:type_name -> GachaCeiling
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ExchangeGachaCeilingScRsp_proto_init() }
func file_ExchangeGachaCeilingScRsp_proto_init() {
	if File_ExchangeGachaCeilingScRsp_proto != nil {
		return
	}
	file_GachaCeiling_proto_init()
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ExchangeGachaCeilingScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeGachaCeilingScRsp); i {
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
			RawDescriptor: file_ExchangeGachaCeilingScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExchangeGachaCeilingScRsp_proto_goTypes,
		DependencyIndexes: file_ExchangeGachaCeilingScRsp_proto_depIdxs,
		MessageInfos:      file_ExchangeGachaCeilingScRsp_proto_msgTypes,
	}.Build()
	File_ExchangeGachaCeilingScRsp_proto = out.File
	file_ExchangeGachaCeilingScRsp_proto_rawDesc = nil
	file_ExchangeGachaCeilingScRsp_proto_goTypes = nil
	file_ExchangeGachaCeilingScRsp_proto_depIdxs = nil
}
