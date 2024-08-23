// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PVEBattleResultScRsp.proto

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

type PVEBattleResultScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DropData          *ItemList       `protobuf:"bytes,6,opt,name=drop_data,json=dropData,proto3" json:"drop_data,omitempty"`
	EventId           uint32          `protobuf:"varint,3,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	BattleAvatarList  []*BattleAvatar `protobuf:"bytes,11,rep,name=battle_avatar_list,json=battleAvatarList,proto3" json:"battle_avatar_list,omitempty"`
	Unk2              *ItemList       `protobuf:"bytes,1822,opt,name=unk2,proto3" json:"unk2,omitempty"`
	MultipleDropData  *ItemList       `protobuf:"bytes,14,opt,name=multiple_drop_data,json=multipleDropData,proto3" json:"multiple_drop_data,omitempty"`
	CheckIdentical    bool            `protobuf:"varint,1,opt,name=check_identical,json=checkIdentical,proto3" json:"check_identical,omitempty"`
	IEOHEALBOKF       uint32          `protobuf:"varint,9,opt,name=IEOHEALBOKF,proto3" json:"IEOHEALBOKF,omitempty"`
	Retcode           uint32          `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	BinVersion        string          `protobuf:"bytes,12,opt,name=bin_version,json=binVersion,proto3" json:"bin_version,omitempty"` // 15
	Unk3              *ItemList       `protobuf:"bytes,5,opt,name=unk3,proto3" json:"unk3,omitempty"`
	BattleId          uint32          `protobuf:"varint,10,opt,name=battle_id,json=battleId,proto3" json:"battle_id,omitempty"`
	MismatchTurnCount uint32          `protobuf:"varint,2,opt,name=mismatch_turn_count,json=mismatchTurnCount,proto3" json:"mismatch_turn_count,omitempty"`
	ResVersion        string          `protobuf:"bytes,15,opt,name=res_version,json=resVersion,proto3" json:"res_version,omitempty"` // 12
	EndStatus         BattleEndStatus `protobuf:"varint,7,opt,name=end_status,json=endStatus,proto3,enum=BattleEndStatus" json:"end_status,omitempty"`
	IAFPGFMHPCJ       uint32          `protobuf:"varint,13,opt,name=IAFPGFMHPCJ,proto3" json:"IAFPGFMHPCJ,omitempty"`
	StageId           uint32          `protobuf:"varint,4,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
}

func (x *PVEBattleResultScRsp) Reset() {
	*x = PVEBattleResultScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PVEBattleResultScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PVEBattleResultScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PVEBattleResultScRsp) ProtoMessage() {}

func (x *PVEBattleResultScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PVEBattleResultScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PVEBattleResultScRsp.ProtoReflect.Descriptor instead.
func (*PVEBattleResultScRsp) Descriptor() ([]byte, []int) {
	return file_PVEBattleResultScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PVEBattleResultScRsp) GetDropData() *ItemList {
	if x != nil {
		return x.DropData
	}
	return nil
}

func (x *PVEBattleResultScRsp) GetEventId() uint32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *PVEBattleResultScRsp) GetBattleAvatarList() []*BattleAvatar {
	if x != nil {
		return x.BattleAvatarList
	}
	return nil
}

func (x *PVEBattleResultScRsp) GetUnk2() *ItemList {
	if x != nil {
		return x.Unk2
	}
	return nil
}

func (x *PVEBattleResultScRsp) GetMultipleDropData() *ItemList {
	if x != nil {
		return x.MultipleDropData
	}
	return nil
}

func (x *PVEBattleResultScRsp) GetCheckIdentical() bool {
	if x != nil {
		return x.CheckIdentical
	}
	return false
}

func (x *PVEBattleResultScRsp) GetIEOHEALBOKF() uint32 {
	if x != nil {
		return x.IEOHEALBOKF
	}
	return 0
}

func (x *PVEBattleResultScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *PVEBattleResultScRsp) GetBinVersion() string {
	if x != nil {
		return x.BinVersion
	}
	return ""
}

func (x *PVEBattleResultScRsp) GetUnk3() *ItemList {
	if x != nil {
		return x.Unk3
	}
	return nil
}

func (x *PVEBattleResultScRsp) GetBattleId() uint32 {
	if x != nil {
		return x.BattleId
	}
	return 0
}

func (x *PVEBattleResultScRsp) GetMismatchTurnCount() uint32 {
	if x != nil {
		return x.MismatchTurnCount
	}
	return 0
}

func (x *PVEBattleResultScRsp) GetResVersion() string {
	if x != nil {
		return x.ResVersion
	}
	return ""
}

func (x *PVEBattleResultScRsp) GetEndStatus() BattleEndStatus {
	if x != nil {
		return x.EndStatus
	}
	return BattleEndStatus_BATTLE_END_NONE
}

func (x *PVEBattleResultScRsp) GetIAFPGFMHPCJ() uint32 {
	if x != nil {
		return x.IAFPGFMHPCJ
	}
	return 0
}

func (x *PVEBattleResultScRsp) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

var File_PVEBattleResultScRsp_proto protoreflect.FileDescriptor

var file_PVEBattleResultScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x50, 0x56, 0x45, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x04, 0x0a, 0x14, 0x50, 0x56, 0x45, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x12, 0x26, 0x0a, 0x09, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08,
	0x64, 0x72, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x12, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x10,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x32, 0x18, 0x9e, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x32,
	0x12, 0x37, 0x0a, 0x12, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x64, 0x72, 0x6f,
	0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x10, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x65, 0x44, 0x72, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x45, 0x4f, 0x48, 0x45, 0x41, 0x4c, 0x42, 0x4f, 0x4b,
	0x46, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x45, 0x4f, 0x48, 0x45, 0x41, 0x4c,
	0x42, 0x4f, 0x4b, 0x46, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x33, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x33, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6d,
	0x69, 0x73, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6d, 0x69, 0x73, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x54, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x0a,
	0x65, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x49, 0x41, 0x46, 0x50, 0x47, 0x46, 0x4d, 0x48, 0x50, 0x43, 0x4a, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x49, 0x41, 0x46, 0x50, 0x47, 0x46, 0x4d, 0x48, 0x50, 0x43, 0x4a, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_PVEBattleResultScRsp_proto_rawDescOnce sync.Once
	file_PVEBattleResultScRsp_proto_rawDescData = file_PVEBattleResultScRsp_proto_rawDesc
)

func file_PVEBattleResultScRsp_proto_rawDescGZIP() []byte {
	file_PVEBattleResultScRsp_proto_rawDescOnce.Do(func() {
		file_PVEBattleResultScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PVEBattleResultScRsp_proto_rawDescData)
	})
	return file_PVEBattleResultScRsp_proto_rawDescData
}

var file_PVEBattleResultScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PVEBattleResultScRsp_proto_goTypes = []interface{}{
	(*PVEBattleResultScRsp)(nil), // 0: PVEBattleResultScRsp
	(*ItemList)(nil),             // 1: ItemList
	(*BattleAvatar)(nil),         // 2: BattleAvatar
	(BattleEndStatus)(0),         // 3: BattleEndStatus
}
var file_PVEBattleResultScRsp_proto_depIdxs = []int32{
	1, // 0: PVEBattleResultScRsp.drop_data:type_name -> ItemList
	2, // 1: PVEBattleResultScRsp.battle_avatar_list:type_name -> BattleAvatar
	1, // 2: PVEBattleResultScRsp.unk2:type_name -> ItemList
	1, // 3: PVEBattleResultScRsp.multiple_drop_data:type_name -> ItemList
	1, // 4: PVEBattleResultScRsp.unk3:type_name -> ItemList
	3, // 5: PVEBattleResultScRsp.end_status:type_name -> BattleEndStatus
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_PVEBattleResultScRsp_proto_init() }
func file_PVEBattleResultScRsp_proto_init() {
	if File_PVEBattleResultScRsp_proto != nil {
		return
	}
	file_ItemList_proto_init()
	file_BattleEndStatus_proto_init()
	file_BattleAvatar_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PVEBattleResultScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PVEBattleResultScRsp); i {
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
			RawDescriptor: file_PVEBattleResultScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PVEBattleResultScRsp_proto_goTypes,
		DependencyIndexes: file_PVEBattleResultScRsp_proto_depIdxs,
		MessageInfos:      file_PVEBattleResultScRsp_proto_msgTypes,
	}.Build()
	File_PVEBattleResultScRsp_proto = out.File
	file_PVEBattleResultScRsp_proto_rawDesc = nil
	file_PVEBattleResultScRsp_proto_goTypes = nil
	file_PVEBattleResultScRsp_proto_depIdxs = nil
}
