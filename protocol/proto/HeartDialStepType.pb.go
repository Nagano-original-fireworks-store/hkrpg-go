// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HeartDialStepType.proto

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

type HeartDialStepType int32

const (
	HeartDialStepType_HEART_DIAL_STEP_TYPE_MISSING HeartDialStepType = 0
	HeartDialStepType_HEART_DIAL_STEP_TYPE_FULL    HeartDialStepType = 1
	HeartDialStepType_HEART_DIAL_STEP_TYPE_LOCK    HeartDialStepType = 2
	HeartDialStepType_HEART_DIAL_STEP_TYPE_UNLOCK  HeartDialStepType = 3
	HeartDialStepType_HEART_DIAL_STEP_TYPE_NORMAL  HeartDialStepType = 4
	HeartDialStepType_HEART_DIAL_STEP_TYPE_CONTROL HeartDialStepType = 5
)

// Enum value maps for HeartDialStepType.
var (
	HeartDialStepType_name = map[int32]string{
		0: "HEART_DIAL_STEP_TYPE_MISSING",
		1: "HEART_DIAL_STEP_TYPE_FULL",
		2: "HEART_DIAL_STEP_TYPE_LOCK",
		3: "HEART_DIAL_STEP_TYPE_UNLOCK",
		4: "HEART_DIAL_STEP_TYPE_NORMAL",
		5: "HEART_DIAL_STEP_TYPE_CONTROL",
	}
	HeartDialStepType_value = map[string]int32{
		"HEART_DIAL_STEP_TYPE_MISSING": 0,
		"HEART_DIAL_STEP_TYPE_FULL":    1,
		"HEART_DIAL_STEP_TYPE_LOCK":    2,
		"HEART_DIAL_STEP_TYPE_UNLOCK":  3,
		"HEART_DIAL_STEP_TYPE_NORMAL":  4,
		"HEART_DIAL_STEP_TYPE_CONTROL": 5,
	}
)

func (x HeartDialStepType) Enum() *HeartDialStepType {
	p := new(HeartDialStepType)
	*p = x
	return p
}

func (x HeartDialStepType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HeartDialStepType) Descriptor() protoreflect.EnumDescriptor {
	return file_HeartDialStepType_proto_enumTypes[0].Descriptor()
}

func (HeartDialStepType) Type() protoreflect.EnumType {
	return &file_HeartDialStepType_proto_enumTypes[0]
}

func (x HeartDialStepType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HeartDialStepType.Descriptor instead.
func (HeartDialStepType) EnumDescriptor() ([]byte, []int) {
	return file_HeartDialStepType_proto_rawDescGZIP(), []int{0}
}

var File_HeartDialStepType_proto protoreflect.FileDescriptor

var file_HeartDialStepType_proto_rawDesc = []byte{
	0x0a, 0x17, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xd7, 0x01, 0x0a, 0x11, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x1c, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x54,
	0x45, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10,
	0x00, 0x12, 0x1d, 0x0a, 0x19, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x5f,
	0x53, 0x54, 0x45, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x01,
	0x12, 0x1d, 0x0a, 0x19, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x5f, 0x53,
	0x54, 0x45, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x02, 0x12,
	0x1f, 0x0a, 0x1b, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x54,
	0x45, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x03,
	0x12, 0x1f, 0x0a, 0x1b, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x5f, 0x53,
	0x54, 0x45, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10,
	0x04, 0x12, 0x20, 0x0a, 0x1c, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x5f,
	0x53, 0x54, 0x45, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f,
	0x4c, 0x10, 0x05, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HeartDialStepType_proto_rawDescOnce sync.Once
	file_HeartDialStepType_proto_rawDescData = file_HeartDialStepType_proto_rawDesc
)

func file_HeartDialStepType_proto_rawDescGZIP() []byte {
	file_HeartDialStepType_proto_rawDescOnce.Do(func() {
		file_HeartDialStepType_proto_rawDescData = protoimpl.X.CompressGZIP(file_HeartDialStepType_proto_rawDescData)
	})
	return file_HeartDialStepType_proto_rawDescData
}

var file_HeartDialStepType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_HeartDialStepType_proto_goTypes = []interface{}{
	(HeartDialStepType)(0), // 0: HeartDialStepType
}
var file_HeartDialStepType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HeartDialStepType_proto_init() }
func file_HeartDialStepType_proto_init() {
	if File_HeartDialStepType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HeartDialStepType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HeartDialStepType_proto_goTypes,
		DependencyIndexes: file_HeartDialStepType_proto_depIdxs,
		EnumInfos:         file_HeartDialStepType_proto_enumTypes,
	}.Build()
	File_HeartDialStepType_proto = out.File
	file_HeartDialStepType_proto_rawDesc = nil
	file_HeartDialStepType_proto_goTypes = nil
	file_HeartDialStepType_proto_depIdxs = nil
}
