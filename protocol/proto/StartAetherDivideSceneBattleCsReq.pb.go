// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: StartAetherDivideSceneBattleCsReq.proto

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

type StartAetherDivideSceneBattleCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttackedByEntityId        uint32                     `protobuf:"varint,12,opt,name=attacked_by_entity_id,json=attackedByEntityId,proto3" json:"attacked_by_entity_id,omitempty"`
	SkillIndex                uint32                     `protobuf:"varint,9,opt,name=skill_index,json=skillIndex,proto3" json:"skill_index,omitempty"`
	CastEntityId              uint32                     `protobuf:"varint,10,opt,name=cast_entity_id,json=castEntityId,proto3" json:"cast_entity_id,omitempty"`
	AssistMonsterEntityIdList []uint32                   `protobuf:"varint,1,rep,packed,name=assist_monster_entity_id_list,json=assistMonsterEntityIdList,proto3" json:"assist_monster_entity_id_list,omitempty"`
	AssistMonsterEntityInfo   []*AssistMonsterEntityInfo `protobuf:"bytes,15,rep,name=assist_monster_entity_info,json=assistMonsterEntityInfo,proto3" json:"assist_monster_entity_info,omitempty"`
}

func (x *StartAetherDivideSceneBattleCsReq) Reset() {
	*x = StartAetherDivideSceneBattleCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StartAetherDivideSceneBattleCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartAetherDivideSceneBattleCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartAetherDivideSceneBattleCsReq) ProtoMessage() {}

func (x *StartAetherDivideSceneBattleCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_StartAetherDivideSceneBattleCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartAetherDivideSceneBattleCsReq.ProtoReflect.Descriptor instead.
func (*StartAetherDivideSceneBattleCsReq) Descriptor() ([]byte, []int) {
	return file_StartAetherDivideSceneBattleCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *StartAetherDivideSceneBattleCsReq) GetAttackedByEntityId() uint32 {
	if x != nil {
		return x.AttackedByEntityId
	}
	return 0
}

func (x *StartAetherDivideSceneBattleCsReq) GetSkillIndex() uint32 {
	if x != nil {
		return x.SkillIndex
	}
	return 0
}

func (x *StartAetherDivideSceneBattleCsReq) GetCastEntityId() uint32 {
	if x != nil {
		return x.CastEntityId
	}
	return 0
}

func (x *StartAetherDivideSceneBattleCsReq) GetAssistMonsterEntityIdList() []uint32 {
	if x != nil {
		return x.AssistMonsterEntityIdList
	}
	return nil
}

func (x *StartAetherDivideSceneBattleCsReq) GetAssistMonsterEntityInfo() []*AssistMonsterEntityInfo {
	if x != nil {
		return x.AssistMonsterEntityInfo
	}
	return nil
}

var File_StartAetherDivideSceneBattleCsReq_proto protoreflect.FileDescriptor

var file_StartAetherDivideSceneBattleCsReq_proto_rawDesc = []byte{
	0x0a, 0x27, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x44, 0x69, 0x76,
	0x69, 0x64, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x73,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x41, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x21, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x31,
	0x0a, 0x15, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x61, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x1d, 0x61, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x19, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x1a, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x17, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_StartAetherDivideSceneBattleCsReq_proto_rawDescOnce sync.Once
	file_StartAetherDivideSceneBattleCsReq_proto_rawDescData = file_StartAetherDivideSceneBattleCsReq_proto_rawDesc
)

func file_StartAetherDivideSceneBattleCsReq_proto_rawDescGZIP() []byte {
	file_StartAetherDivideSceneBattleCsReq_proto_rawDescOnce.Do(func() {
		file_StartAetherDivideSceneBattleCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_StartAetherDivideSceneBattleCsReq_proto_rawDescData)
	})
	return file_StartAetherDivideSceneBattleCsReq_proto_rawDescData
}

var file_StartAetherDivideSceneBattleCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StartAetherDivideSceneBattleCsReq_proto_goTypes = []interface{}{
	(*StartAetherDivideSceneBattleCsReq)(nil), // 0: StartAetherDivideSceneBattleCsReq
	(*AssistMonsterEntityInfo)(nil),           // 1: AssistMonsterEntityInfo
}
var file_StartAetherDivideSceneBattleCsReq_proto_depIdxs = []int32{
	1, // 0: StartAetherDivideSceneBattleCsReq.assist_monster_entity_info:type_name -> AssistMonsterEntityInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_StartAetherDivideSceneBattleCsReq_proto_init() }
func file_StartAetherDivideSceneBattleCsReq_proto_init() {
	if File_StartAetherDivideSceneBattleCsReq_proto != nil {
		return
	}
	file_AssistMonsterEntityInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_StartAetherDivideSceneBattleCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartAetherDivideSceneBattleCsReq); i {
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
			RawDescriptor: file_StartAetherDivideSceneBattleCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StartAetherDivideSceneBattleCsReq_proto_goTypes,
		DependencyIndexes: file_StartAetherDivideSceneBattleCsReq_proto_depIdxs,
		MessageInfos:      file_StartAetherDivideSceneBattleCsReq_proto_msgTypes,
	}.Build()
	File_StartAetherDivideSceneBattleCsReq_proto = out.File
	file_StartAetherDivideSceneBattleCsReq_proto_rawDesc = nil
	file_StartAetherDivideSceneBattleCsReq_proto_goTypes = nil
	file_StartAetherDivideSceneBattleCsReq_proto_depIdxs = nil
}
