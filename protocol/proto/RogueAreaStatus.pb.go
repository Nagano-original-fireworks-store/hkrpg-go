// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueAreaStatus.proto

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

type RogueAreaStatus int32

const (
	RogueAreaStatus_ROGUE_AREA_STATUS_LOCK       RogueAreaStatus = 0
	RogueAreaStatus_ROGUE_AREA_STATUS_UNLOCK     RogueAreaStatus = 1
	RogueAreaStatus_ROGUE_AREA_STATUS_FIRST_PASS RogueAreaStatus = 2
	RogueAreaStatus_ROGUE_AREA_STATUS_CLOSE      RogueAreaStatus = 3
)

// Enum value maps for RogueAreaStatus.
var (
	RogueAreaStatus_name = map[int32]string{
		0: "ROGUE_AREA_STATUS_LOCK",
		1: "ROGUE_AREA_STATUS_UNLOCK",
		2: "ROGUE_AREA_STATUS_FIRST_PASS",
		3: "ROGUE_AREA_STATUS_CLOSE",
	}
	RogueAreaStatus_value = map[string]int32{
		"ROGUE_AREA_STATUS_LOCK":       0,
		"ROGUE_AREA_STATUS_UNLOCK":     1,
		"ROGUE_AREA_STATUS_FIRST_PASS": 2,
		"ROGUE_AREA_STATUS_CLOSE":      3,
	}
)

func (x RogueAreaStatus) Enum() *RogueAreaStatus {
	p := new(RogueAreaStatus)
	*p = x
	return p
}

func (x RogueAreaStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RogueAreaStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_RogueAreaStatus_proto_enumTypes[0].Descriptor()
}

func (RogueAreaStatus) Type() protoreflect.EnumType {
	return &file_RogueAreaStatus_proto_enumTypes[0]
}

func (x RogueAreaStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RogueAreaStatus.Descriptor instead.
func (RogueAreaStatus) EnumDescriptor() ([]byte, []int) {
	return file_RogueAreaStatus_proto_rawDescGZIP(), []int{0}
}

var File_RogueAreaStatus_proto protoreflect.FileDescriptor

var file_RogueAreaStatus_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x72, 0x65, 0x61, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x8a, 0x01, 0x0a, 0x0f, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x41, 0x72, 0x65, 0x61, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x52,
	0x4f, 0x47, 0x55, 0x45, 0x5f, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x4f, 0x47, 0x55, 0x45,
	0x5f, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4c,
	0x4f, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x41,
	0x52, 0x45, 0x41, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54,
	0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x4f, 0x47, 0x55, 0x45,
	0x5f, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4c, 0x4f,
	0x53, 0x45, 0x10, 0x03, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueAreaStatus_proto_rawDescOnce sync.Once
	file_RogueAreaStatus_proto_rawDescData = file_RogueAreaStatus_proto_rawDesc
)

func file_RogueAreaStatus_proto_rawDescGZIP() []byte {
	file_RogueAreaStatus_proto_rawDescOnce.Do(func() {
		file_RogueAreaStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueAreaStatus_proto_rawDescData)
	})
	return file_RogueAreaStatus_proto_rawDescData
}

var file_RogueAreaStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RogueAreaStatus_proto_goTypes = []interface{}{
	(RogueAreaStatus)(0), // 0: RogueAreaStatus
}
var file_RogueAreaStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueAreaStatus_proto_init() }
func file_RogueAreaStatus_proto_init() {
	if File_RogueAreaStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueAreaStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueAreaStatus_proto_goTypes,
		DependencyIndexes: file_RogueAreaStatus_proto_depIdxs,
		EnumInfos:         file_RogueAreaStatus_proto_enumTypes,
	}.Build()
	File_RogueAreaStatus_proto = out.File
	file_RogueAreaStatus_proto_rawDesc = nil
	file_RogueAreaStatus_proto_goTypes = nil
	file_RogueAreaStatus_proto_depIdxs = nil
}
