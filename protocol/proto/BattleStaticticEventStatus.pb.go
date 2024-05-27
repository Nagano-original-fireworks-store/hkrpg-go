// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BattleStaticticEventStatus.proto

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

type BattleStaticticEventStatus int32

const (
	BattleStaticticEventStatus_BATTLE_STATICTIC_EVENT_NONE                                       BattleStaticticEventStatus = 0
	BattleStaticticEventStatus_BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_ADD_EXPLORE               BattleStaticticEventStatus = 1
	BattleStaticticEventStatus_BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_OPEN_GRID                 BattleStaticticEventStatus = 2
	BattleStaticticEventStatus_BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_PICKUP_ITEM               BattleStaticticEventStatus = 3
	BattleStaticticEventStatus_BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_USE_BUFF                  BattleStaticticEventStatus = 4
	BattleStaticticEventStatus_BATTLE_STATICTIC_EVENT_TELEVISION_ACTIVITY_UPDATE_MAZE_BUFF_LAYER BattleStaticticEventStatus = 5
)

// Enum value maps for BattleStaticticEventStatus.
var (
	BattleStaticticEventStatus_name = map[int32]string{
		0: "BATTLE_STATICTIC_EVENT_NONE",
		1: "BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_ADD_EXPLORE",
		2: "BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_OPEN_GRID",
		3: "BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_PICKUP_ITEM",
		4: "BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_USE_BUFF",
		5: "BATTLE_STATICTIC_EVENT_TELEVISION_ACTIVITY_UPDATE_MAZE_BUFF_LAYER",
	}
	BattleStaticticEventStatus_value = map[string]int32{
		"BATTLE_STATICTIC_EVENT_NONE":                                       0,
		"BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_ADD_EXPLORE":               1,
		"BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_OPEN_GRID":                 2,
		"BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_PICKUP_ITEM":               3,
		"BATTLE_STATICTIC_EVENT_TREASURE_DUNGEON_USE_BUFF":                  4,
		"BATTLE_STATICTIC_EVENT_TELEVISION_ACTIVITY_UPDATE_MAZE_BUFF_LAYER": 5,
	}
)

func (x BattleStaticticEventStatus) Enum() *BattleStaticticEventStatus {
	p := new(BattleStaticticEventStatus)
	*p = x
	return p
}

func (x BattleStaticticEventStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BattleStaticticEventStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_BattleStaticticEventStatus_proto_enumTypes[0].Descriptor()
}

func (BattleStaticticEventStatus) Type() protoreflect.EnumType {
	return &file_BattleStaticticEventStatus_proto_enumTypes[0]
}

func (x BattleStaticticEventStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BattleStaticticEventStatus.Descriptor instead.
func (BattleStaticticEventStatus) EnumDescriptor() ([]byte, []int) {
	return file_BattleStaticticEventStatus_proto_rawDescGZIP(), []int{0}
}

var File_BattleStaticticEventStatus_proto protoreflect.FileDescriptor

var file_BattleStaticticEventStatus_proto_rawDesc = []byte{
	0x0a, 0x20, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x74, 0x69,
	0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0xe3, 0x02, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x74, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1f, 0x0a, 0x1b, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x49, 0x43, 0x54, 0x49, 0x43, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x37, 0x0a, 0x33, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x49, 0x43, 0x54, 0x49, 0x43, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x52, 0x45,
	0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x41, 0x44,
	0x44, 0x5f, 0x45, 0x58, 0x50, 0x4c, 0x4f, 0x52, 0x45, 0x10, 0x01, 0x12, 0x35, 0x0a, 0x31, 0x42,
	0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x43, 0x54, 0x49, 0x43, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44,
	0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x47, 0x52, 0x49, 0x44,
	0x10, 0x02, 0x12, 0x37, 0x0a, 0x33, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x49, 0x43, 0x54, 0x49, 0x43, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x52, 0x45,
	0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x50, 0x49,
	0x43, 0x4b, 0x55, 0x50, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x10, 0x03, 0x12, 0x34, 0x0a, 0x30, 0x42,
	0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x43, 0x54, 0x49, 0x43, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44,
	0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x10,
	0x04, 0x12, 0x45, 0x0a, 0x41, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x49, 0x43, 0x54, 0x49, 0x43, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x45, 0x4c, 0x45,
	0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x41, 0x5a, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46,
	0x5f, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x10, 0x05, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattleStaticticEventStatus_proto_rawDescOnce sync.Once
	file_BattleStaticticEventStatus_proto_rawDescData = file_BattleStaticticEventStatus_proto_rawDesc
)

func file_BattleStaticticEventStatus_proto_rawDescGZIP() []byte {
	file_BattleStaticticEventStatus_proto_rawDescOnce.Do(func() {
		file_BattleStaticticEventStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattleStaticticEventStatus_proto_rawDescData)
	})
	return file_BattleStaticticEventStatus_proto_rawDescData
}

var file_BattleStaticticEventStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BattleStaticticEventStatus_proto_goTypes = []interface{}{
	(BattleStaticticEventStatus)(0), // 0: BattleStaticticEventStatus
}
var file_BattleStaticticEventStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BattleStaticticEventStatus_proto_init() }
func file_BattleStaticticEventStatus_proto_init() {
	if File_BattleStaticticEventStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_BattleStaticticEventStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattleStaticticEventStatus_proto_goTypes,
		DependencyIndexes: file_BattleStaticticEventStatus_proto_depIdxs,
		EnumInfos:         file_BattleStaticticEventStatus_proto_enumTypes,
	}.Build()
	File_BattleStaticticEventStatus_proto = out.File
	file_BattleStaticticEventStatus_proto_rawDesc = nil
	file_BattleStaticticEventStatus_proto_goTypes = nil
	file_BattleStaticticEventStatus_proto_depIdxs = nil
}
