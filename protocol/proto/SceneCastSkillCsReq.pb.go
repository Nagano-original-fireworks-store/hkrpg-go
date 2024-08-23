// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneCastSkillCsReq.proto

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

type SceneCastSkillCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LFHDFMJDKMN               uint32                     `protobuf:"varint,3,opt,name=LFHDFMJDKMN,proto3" json:"LFHDFMJDKMN,omitempty"`
	CastEntityId              uint32                     `protobuf:"varint,10,opt,name=cast_entity_id,json=castEntityId,proto3" json:"cast_entity_id,omitempty"`
	AssistMonsterEntityInfo   []*AssistMonsterEntityInfo `protobuf:"bytes,4,rep,name=assist_monster_entity_info,json=assistMonsterEntityInfo,proto3" json:"assist_monster_entity_info,omitempty"`
	DynamicValues             []*ABJCBAOKICE             `protobuf:"bytes,11,rep,name=dynamic_values,json=dynamicValues,proto3" json:"dynamic_values,omitempty"`
	SkillIndex                uint32                     `protobuf:"varint,15,opt,name=skill_index,json=skillIndex,proto3" json:"skill_index,omitempty"`
	TargetMotion              *MotionInfo                `protobuf:"bytes,12,opt,name=target_motion,json=targetMotion,proto3" json:"target_motion,omitempty"`
	HitTargetEntityIdList     []uint32                   `protobuf:"varint,8,rep,packed,name=hit_target_entity_id_list,json=hitTargetEntityIdList,proto3" json:"hit_target_entity_id_list,omitempty"`
	AttackedByEntityId        uint32                     `protobuf:"varint,13,opt,name=attacked_by_entity_id,json=attackedByEntityId,proto3" json:"attacked_by_entity_id,omitempty"`
	SkillExtraTags            []SkillExtraTag            `protobuf:"varint,5,rep,packed,name=skill_extra_tags,json=skillExtraTags,proto3,enum=SkillExtraTag" json:"skill_extra_tags,omitempty"`
	AssistMonsterEntityIdList []uint32                   `protobuf:"varint,2,rep,packed,name=assist_monster_entity_id_list,json=assistMonsterEntityIdList,proto3" json:"assist_monster_entity_id_list,omitempty"`
}

func (x *SceneCastSkillCsReq) Reset() {
	*x = SceneCastSkillCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneCastSkillCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneCastSkillCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneCastSkillCsReq) ProtoMessage() {}

func (x *SceneCastSkillCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SceneCastSkillCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneCastSkillCsReq.ProtoReflect.Descriptor instead.
func (*SceneCastSkillCsReq) Descriptor() ([]byte, []int) {
	return file_SceneCastSkillCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SceneCastSkillCsReq) GetLFHDFMJDKMN() uint32 {
	if x != nil {
		return x.LFHDFMJDKMN
	}
	return 0
}

func (x *SceneCastSkillCsReq) GetCastEntityId() uint32 {
	if x != nil {
		return x.CastEntityId
	}
	return 0
}

func (x *SceneCastSkillCsReq) GetAssistMonsterEntityInfo() []*AssistMonsterEntityInfo {
	if x != nil {
		return x.AssistMonsterEntityInfo
	}
	return nil
}

func (x *SceneCastSkillCsReq) GetDynamicValues() []*ABJCBAOKICE {
	if x != nil {
		return x.DynamicValues
	}
	return nil
}

func (x *SceneCastSkillCsReq) GetSkillIndex() uint32 {
	if x != nil {
		return x.SkillIndex
	}
	return 0
}

func (x *SceneCastSkillCsReq) GetTargetMotion() *MotionInfo {
	if x != nil {
		return x.TargetMotion
	}
	return nil
}

func (x *SceneCastSkillCsReq) GetHitTargetEntityIdList() []uint32 {
	if x != nil {
		return x.HitTargetEntityIdList
	}
	return nil
}

func (x *SceneCastSkillCsReq) GetAttackedByEntityId() uint32 {
	if x != nil {
		return x.AttackedByEntityId
	}
	return 0
}

func (x *SceneCastSkillCsReq) GetSkillExtraTags() []SkillExtraTag {
	if x != nil {
		return x.SkillExtraTags
	}
	return nil
}

func (x *SceneCastSkillCsReq) GetAssistMonsterEntityIdList() []uint32 {
	if x != nil {
		return x.AssistMonsterEntityIdList
	}
	return nil
}

var File_SceneCastSkillCsReq_proto protoreflect.FileDescriptor

var file_SceneCastSkillCsReq_proto_rawDesc = []byte{
	0x0a, 0x19, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x61, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x54, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x41, 0x42, 0x4a, 0x43, 0x42, 0x41, 0x4f, 0x4b, 0x49, 0x43, 0x45, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x04, 0x0a, 0x13, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x61,
	0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b,
	0x4c, 0x46, 0x48, 0x44, 0x46, 0x4d, 0x4a, 0x44, 0x4b, 0x4d, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x4c, 0x46, 0x48, 0x44, 0x46, 0x4d, 0x4a, 0x44, 0x4b, 0x4d, 0x4e, 0x12, 0x24,
	0x0a, 0x0e, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x55, 0x0a, 0x1a, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x6d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x17, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x0e, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x42, 0x4a, 0x43, 0x42, 0x41, 0x4f, 0x4b, 0x49, 0x43,
	0x45, 0x52, 0x0d, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x30, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x19, 0x68, 0x69, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x15, 0x68, 0x69, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x15, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x38, 0x0a, 0x10, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x54, 0x61, 0x67, 0x52, 0x0e, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x54, 0x61, 0x67, 0x73, 0x12, 0x40, 0x0a, 0x1d, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x19, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneCastSkillCsReq_proto_rawDescOnce sync.Once
	file_SceneCastSkillCsReq_proto_rawDescData = file_SceneCastSkillCsReq_proto_rawDesc
)

func file_SceneCastSkillCsReq_proto_rawDescGZIP() []byte {
	file_SceneCastSkillCsReq_proto_rawDescOnce.Do(func() {
		file_SceneCastSkillCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneCastSkillCsReq_proto_rawDescData)
	})
	return file_SceneCastSkillCsReq_proto_rawDescData
}

var file_SceneCastSkillCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneCastSkillCsReq_proto_goTypes = []interface{}{
	(*SceneCastSkillCsReq)(nil),     // 0: SceneCastSkillCsReq
	(*AssistMonsterEntityInfo)(nil), // 1: AssistMonsterEntityInfo
	(*ABJCBAOKICE)(nil),             // 2: ABJCBAOKICE
	(*MotionInfo)(nil),              // 3: MotionInfo
	(SkillExtraTag)(0),              // 4: SkillExtraTag
}
var file_SceneCastSkillCsReq_proto_depIdxs = []int32{
	1, // 0: SceneCastSkillCsReq.assist_monster_entity_info:type_name -> AssistMonsterEntityInfo
	2, // 1: SceneCastSkillCsReq.dynamic_values:type_name -> ABJCBAOKICE
	3, // 2: SceneCastSkillCsReq.target_motion:type_name -> MotionInfo
	4, // 3: SceneCastSkillCsReq.skill_extra_tags:type_name -> SkillExtraTag
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_SceneCastSkillCsReq_proto_init() }
func file_SceneCastSkillCsReq_proto_init() {
	if File_SceneCastSkillCsReq_proto != nil {
		return
	}
	file_MotionInfo_proto_init()
	file_SkillExtraTag_proto_init()
	file_ABJCBAOKICE_proto_init()
	file_AssistMonsterEntityInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneCastSkillCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneCastSkillCsReq); i {
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
			RawDescriptor: file_SceneCastSkillCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneCastSkillCsReq_proto_goTypes,
		DependencyIndexes: file_SceneCastSkillCsReq_proto_depIdxs,
		MessageInfos:      file_SceneCastSkillCsReq_proto_msgTypes,
	}.Build()
	File_SceneCastSkillCsReq_proto = out.File
	file_SceneCastSkillCsReq_proto_rawDesc = nil
	file_SceneCastSkillCsReq_proto_goTypes = nil
	file_SceneCastSkillCsReq_proto_depIdxs = nil
}
