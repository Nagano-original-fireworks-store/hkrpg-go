// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RankUpAvatarCsReq.proto

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

type RankUpAvatarCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank     uint32        `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	CostData *ItemCostData `protobuf:"bytes,2,opt,name=cost_data,json=costData,proto3" json:"cost_data,omitempty"`
	AvatarId uint32        `protobuf:"varint,12,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
}

func (x *RankUpAvatarCsReq) Reset() {
	*x = RankUpAvatarCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RankUpAvatarCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankUpAvatarCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankUpAvatarCsReq) ProtoMessage() {}

func (x *RankUpAvatarCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_RankUpAvatarCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankUpAvatarCsReq.ProtoReflect.Descriptor instead.
func (*RankUpAvatarCsReq) Descriptor() ([]byte, []int) {
	return file_RankUpAvatarCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *RankUpAvatarCsReq) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *RankUpAvatarCsReq) GetCostData() *ItemCostData {
	if x != nil {
		return x.CostData
	}
	return nil
}

func (x *RankUpAvatarCsReq) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

var File_RankUpAvatarCsReq_proto protoreflect.FileDescriptor

var file_RankUpAvatarCsReq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x52, 0x61, 0x6e, 0x6b, 0x55, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x73,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x49, 0x74, 0x65, 0x6d, 0x43,
	0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a,
	0x11, 0x52, 0x61, 0x6e, 0x6b, 0x55, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x2a, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RankUpAvatarCsReq_proto_rawDescOnce sync.Once
	file_RankUpAvatarCsReq_proto_rawDescData = file_RankUpAvatarCsReq_proto_rawDesc
)

func file_RankUpAvatarCsReq_proto_rawDescGZIP() []byte {
	file_RankUpAvatarCsReq_proto_rawDescOnce.Do(func() {
		file_RankUpAvatarCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_RankUpAvatarCsReq_proto_rawDescData)
	})
	return file_RankUpAvatarCsReq_proto_rawDescData
}

var file_RankUpAvatarCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RankUpAvatarCsReq_proto_goTypes = []interface{}{
	(*RankUpAvatarCsReq)(nil), // 0: RankUpAvatarCsReq
	(*ItemCostData)(nil),      // 1: ItemCostData
}
var file_RankUpAvatarCsReq_proto_depIdxs = []int32{
	1, // 0: RankUpAvatarCsReq.cost_data:type_name -> ItemCostData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RankUpAvatarCsReq_proto_init() }
func file_RankUpAvatarCsReq_proto_init() {
	if File_RankUpAvatarCsReq_proto != nil {
		return
	}
	file_ItemCostData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RankUpAvatarCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankUpAvatarCsReq); i {
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
			RawDescriptor: file_RankUpAvatarCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RankUpAvatarCsReq_proto_goTypes,
		DependencyIndexes: file_RankUpAvatarCsReq_proto_depIdxs,
		MessageInfos:      file_RankUpAvatarCsReq_proto_msgTypes,
	}.Build()
	File_RankUpAvatarCsReq_proto = out.File
	file_RankUpAvatarCsReq_proto_rawDesc = nil
	file_RankUpAvatarCsReq_proto_goTypes = nil
	file_RankUpAvatarCsReq_proto_depIdxs = nil
}
