// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneEntityInfo.proto

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

type SceneEntityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId   uint32               `protobuf:"varint,12,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	GroupId    uint32               `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	InstId     uint32               `protobuf:"varint,1,opt,name=inst_id,json=instId,proto3" json:"inst_id,omitempty"`
	Motion     *MotionInfo          `protobuf:"bytes,5,opt,name=motion,proto3" json:"motion,omitempty"`
	Actor      *SceneActorInfo      `protobuf:"bytes,8,opt,name=actor,proto3" json:"actor,omitempty"`
	NpcMonster *SceneNpcMonsterInfo `protobuf:"bytes,10,opt,name=npc_monster,json=npcMonster,proto3" json:"npc_monster,omitempty"`
	Npc        *SceneNpcInfo        `protobuf:"bytes,3,opt,name=npc,proto3" json:"npc,omitempty"`
	Prop       *ScenePropInfo       `protobuf:"bytes,13,opt,name=prop,proto3" json:"prop,omitempty"`
	SummonUnit *SceneSummonUnitInfo `protobuf:"bytes,14,opt,name=summon_unit,json=summonUnit,proto3" json:"summon_unit,omitempty"`
}

func (x *SceneEntityInfo) Reset() {
	*x = SceneEntityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneEntityInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneEntityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneEntityInfo) ProtoMessage() {}

func (x *SceneEntityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneEntityInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneEntityInfo.ProtoReflect.Descriptor instead.
func (*SceneEntityInfo) Descriptor() ([]byte, []int) {
	return file_SceneEntityInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneEntityInfo) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *SceneEntityInfo) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *SceneEntityInfo) GetInstId() uint32 {
	if x != nil {
		return x.InstId
	}
	return 0
}

func (x *SceneEntityInfo) GetMotion() *MotionInfo {
	if x != nil {
		return x.Motion
	}
	return nil
}

func (x *SceneEntityInfo) GetActor() *SceneActorInfo {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *SceneEntityInfo) GetNpcMonster() *SceneNpcMonsterInfo {
	if x != nil {
		return x.NpcMonster
	}
	return nil
}

func (x *SceneEntityInfo) GetNpc() *SceneNpcInfo {
	if x != nil {
		return x.Npc
	}
	return nil
}

func (x *SceneEntityInfo) GetProp() *ScenePropInfo {
	if x != nil {
		return x.Prop
	}
	return nil
}

func (x *SceneEntityInfo) GetSummonUnit() *SceneSummonUnitInfo {
	if x != nil {
		return x.SummonUnit
	}
	return nil
}

var File_SceneEntityInfo_proto protoreflect.FileDescriptor

var file_SceneEntityInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x4e, 0x70, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x50, 0x72, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1,
	0x02, 0x0a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6e,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x6e, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x35, 0x0a, 0x0b, 0x6e, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6e, 0x70, 0x63, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x03, 0x6e, 0x70, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x03, 0x6e, 0x70, 0x63, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x72, 0x6f, 0x70, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x70, 0x72, 0x6f, 0x70, 0x12, 0x35, 0x0a, 0x0b, 0x73,
	0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x55, 0x6e,
	0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x55, 0x6e,
	0x69, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneEntityInfo_proto_rawDescOnce sync.Once
	file_SceneEntityInfo_proto_rawDescData = file_SceneEntityInfo_proto_rawDesc
)

func file_SceneEntityInfo_proto_rawDescGZIP() []byte {
	file_SceneEntityInfo_proto_rawDescOnce.Do(func() {
		file_SceneEntityInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneEntityInfo_proto_rawDescData)
	})
	return file_SceneEntityInfo_proto_rawDescData
}

var file_SceneEntityInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneEntityInfo_proto_goTypes = []interface{}{
	(*SceneEntityInfo)(nil),     // 0: SceneEntityInfo
	(*MotionInfo)(nil),          // 1: MotionInfo
	(*SceneActorInfo)(nil),      // 2: SceneActorInfo
	(*SceneNpcMonsterInfo)(nil), // 3: SceneNpcMonsterInfo
	(*SceneNpcInfo)(nil),        // 4: SceneNpcInfo
	(*ScenePropInfo)(nil),       // 5: ScenePropInfo
	(*SceneSummonUnitInfo)(nil), // 6: SceneSummonUnitInfo
}
var file_SceneEntityInfo_proto_depIdxs = []int32{
	1, // 0: SceneEntityInfo.motion:type_name -> MotionInfo
	2, // 1: SceneEntityInfo.actor:type_name -> SceneActorInfo
	3, // 2: SceneEntityInfo.npc_monster:type_name -> SceneNpcMonsterInfo
	4, // 3: SceneEntityInfo.npc:type_name -> SceneNpcInfo
	5, // 4: SceneEntityInfo.prop:type_name -> ScenePropInfo
	6, // 5: SceneEntityInfo.summon_unit:type_name -> SceneSummonUnitInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_SceneEntityInfo_proto_init() }
func file_SceneEntityInfo_proto_init() {
	if File_SceneEntityInfo_proto != nil {
		return
	}
	file_MotionInfo_proto_init()
	file_SceneActorInfo_proto_init()
	file_SceneNpcMonsterInfo_proto_init()
	file_SceneNpcInfo_proto_init()
	file_SceneSummonUnitInfo_proto_init()
	file_ScenePropInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneEntityInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneEntityInfo); i {
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
			RawDescriptor: file_SceneEntityInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneEntityInfo_proto_goTypes,
		DependencyIndexes: file_SceneEntityInfo_proto_depIdxs,
		MessageInfos:      file_SceneEntityInfo_proto_msgTypes,
	}.Build()
	File_SceneEntityInfo_proto = out.File
	file_SceneEntityInfo_proto_rawDesc = nil
	file_SceneEntityInfo_proto_goTypes = nil
	file_SceneEntityInfo_proto_depIdxs = nil
}
