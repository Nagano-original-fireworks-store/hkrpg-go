// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AddAvatarSrcState.proto

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

type AddAvatarSrcState int32

const (
	AddAvatarSrcState_ADD_AVATAR_SRC_NONE  AddAvatarSrcState = 0
	AddAvatarSrcState_ADD_AVATAR_SRC_GACHA AddAvatarSrcState = 1
	AddAvatarSrcState_ADD_AVATAR_SRC_ROGUE AddAvatarSrcState = 2
)

// Enum value maps for AddAvatarSrcState.
var (
	AddAvatarSrcState_name = map[int32]string{
		0: "ADD_AVATAR_SRC_NONE",
		1: "ADD_AVATAR_SRC_GACHA",
		2: "ADD_AVATAR_SRC_ROGUE",
	}
	AddAvatarSrcState_value = map[string]int32{
		"ADD_AVATAR_SRC_NONE":  0,
		"ADD_AVATAR_SRC_GACHA": 1,
		"ADD_AVATAR_SRC_ROGUE": 2,
	}
)

func (x AddAvatarSrcState) Enum() *AddAvatarSrcState {
	p := new(AddAvatarSrcState)
	*p = x
	return p
}

func (x AddAvatarSrcState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddAvatarSrcState) Descriptor() protoreflect.EnumDescriptor {
	return file_AddAvatarSrcState_proto_enumTypes[0].Descriptor()
}

func (AddAvatarSrcState) Type() protoreflect.EnumType {
	return &file_AddAvatarSrcState_proto_enumTypes[0]
}

func (x AddAvatarSrcState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddAvatarSrcState.Descriptor instead.
func (AddAvatarSrcState) EnumDescriptor() ([]byte, []int) {
	return file_AddAvatarSrcState_proto_rawDescGZIP(), []int{0}
}

var File_AddAvatarSrcState_proto protoreflect.FileDescriptor

var file_AddAvatarSrcState_proto_rawDesc = []byte{
	0x0a, 0x17, 0x41, 0x64, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x72, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x60, 0x0a, 0x11, 0x41, 0x64, 0x64,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x72, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17,
	0x0a, 0x13, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x53, 0x52, 0x43,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x44, 0x44, 0x5f, 0x41,
	0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x53, 0x52, 0x43, 0x5f, 0x47, 0x41, 0x43, 0x48, 0x41, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f,
	0x53, 0x52, 0x43, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_AddAvatarSrcState_proto_rawDescOnce sync.Once
	file_AddAvatarSrcState_proto_rawDescData = file_AddAvatarSrcState_proto_rawDesc
)

func file_AddAvatarSrcState_proto_rawDescGZIP() []byte {
	file_AddAvatarSrcState_proto_rawDescOnce.Do(func() {
		file_AddAvatarSrcState_proto_rawDescData = protoimpl.X.CompressGZIP(file_AddAvatarSrcState_proto_rawDescData)
	})
	return file_AddAvatarSrcState_proto_rawDescData
}

var file_AddAvatarSrcState_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AddAvatarSrcState_proto_goTypes = []interface{}{
	(AddAvatarSrcState)(0), // 0: AddAvatarSrcState
}
var file_AddAvatarSrcState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AddAvatarSrcState_proto_init() }
func file_AddAvatarSrcState_proto_init() {
	if File_AddAvatarSrcState_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AddAvatarSrcState_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AddAvatarSrcState_proto_goTypes,
		DependencyIndexes: file_AddAvatarSrcState_proto_depIdxs,
		EnumInfos:         file_AddAvatarSrcState_proto_enumTypes,
	}.Build()
	File_AddAvatarSrcState_proto = out.File
	file_AddAvatarSrcState_proto_rawDesc = nil
	file_AddAvatarSrcState_proto_goTypes = nil
	file_AddAvatarSrcState_proto_depIdxs = nil
}
