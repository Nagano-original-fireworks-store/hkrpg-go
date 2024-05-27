// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SkillExtraTag.proto

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

type SkillExtraTag int32

const (
	SkillExtraTag_SCENE_CAST_SKILL_NONE                       SkillExtraTag = 0
	SkillExtraTag_SCENE_CAST_SKILL_PROJECTILE_HIT             SkillExtraTag = 1
	SkillExtraTag_SCENE_CAST_SKILL_PROJECTILE_LIFETIME_FINISH SkillExtraTag = 2
)

// Enum value maps for SkillExtraTag.
var (
	SkillExtraTag_name = map[int32]string{
		0: "SCENE_CAST_SKILL_NONE",
		1: "SCENE_CAST_SKILL_PROJECTILE_HIT",
		2: "SCENE_CAST_SKILL_PROJECTILE_LIFETIME_FINISH",
	}
	SkillExtraTag_value = map[string]int32{
		"SCENE_CAST_SKILL_NONE":                       0,
		"SCENE_CAST_SKILL_PROJECTILE_HIT":             1,
		"SCENE_CAST_SKILL_PROJECTILE_LIFETIME_FINISH": 2,
	}
)

func (x SkillExtraTag) Enum() *SkillExtraTag {
	p := new(SkillExtraTag)
	*p = x
	return p
}

func (x SkillExtraTag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SkillExtraTag) Descriptor() protoreflect.EnumDescriptor {
	return file_SkillExtraTag_proto_enumTypes[0].Descriptor()
}

func (SkillExtraTag) Type() protoreflect.EnumType {
	return &file_SkillExtraTag_proto_enumTypes[0]
}

func (x SkillExtraTag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SkillExtraTag.Descriptor instead.
func (SkillExtraTag) EnumDescriptor() ([]byte, []int) {
	return file_SkillExtraTag_proto_rawDescGZIP(), []int{0}
}

var File_SkillExtraTag_proto protoreflect.FileDescriptor

var file_SkillExtraTag_proto_rawDesc = []byte{
	0x0a, 0x13, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x54, 0x61, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x80, 0x01, 0x0a, 0x0d, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x54, 0x61, 0x67, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x43, 0x45, 0x4e, 0x45,
	0x5f, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x43, 0x41, 0x53, 0x54,
	0x5f, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x49, 0x4c,
	0x45, 0x5f, 0x48, 0x49, 0x54, 0x10, 0x01, 0x12, 0x2f, 0x0a, 0x2b, 0x53, 0x43, 0x45, 0x4e, 0x45,
	0x5f, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x4a,
	0x45, 0x43, 0x54, 0x49, 0x4c, 0x45, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x5f,
	0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SkillExtraTag_proto_rawDescOnce sync.Once
	file_SkillExtraTag_proto_rawDescData = file_SkillExtraTag_proto_rawDesc
)

func file_SkillExtraTag_proto_rawDescGZIP() []byte {
	file_SkillExtraTag_proto_rawDescOnce.Do(func() {
		file_SkillExtraTag_proto_rawDescData = protoimpl.X.CompressGZIP(file_SkillExtraTag_proto_rawDescData)
	})
	return file_SkillExtraTag_proto_rawDescData
}

var file_SkillExtraTag_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SkillExtraTag_proto_goTypes = []interface{}{
	(SkillExtraTag)(0), // 0: SkillExtraTag
}
var file_SkillExtraTag_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SkillExtraTag_proto_init() }
func file_SkillExtraTag_proto_init() {
	if File_SkillExtraTag_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SkillExtraTag_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SkillExtraTag_proto_goTypes,
		DependencyIndexes: file_SkillExtraTag_proto_depIdxs,
		EnumInfos:         file_SkillExtraTag_proto_enumTypes,
	}.Build()
	File_SkillExtraTag_proto = out.File
	file_SkillExtraTag_proto_rawDesc = nil
	file_SkillExtraTag_proto_goTypes = nil
	file_SkillExtraTag_proto_depIdxs = nil
}
