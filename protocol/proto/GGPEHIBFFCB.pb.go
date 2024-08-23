// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GGPEHIBFFCB.proto

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

type GGPEHIBFFCB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LNAIONDBECA uint32           `protobuf:"varint,4,opt,name=LNAIONDBECA,proto3" json:"LNAIONDBECA,omitempty"`
	Status      RogueBoothStatus `protobuf:"varint,3,opt,name=status,proto3,enum=RogueBoothStatus" json:"status,omitempty"`
	HKJLPJBINNM uint32           `protobuf:"varint,5,opt,name=HKJLPJBINNM,proto3" json:"HKJLPJBINNM,omitempty"`
}

func (x *GGPEHIBFFCB) Reset() {
	*x = GGPEHIBFFCB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GGPEHIBFFCB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GGPEHIBFFCB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GGPEHIBFFCB) ProtoMessage() {}

func (x *GGPEHIBFFCB) ProtoReflect() protoreflect.Message {
	mi := &file_GGPEHIBFFCB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GGPEHIBFFCB.ProtoReflect.Descriptor instead.
func (*GGPEHIBFFCB) Descriptor() ([]byte, []int) {
	return file_GGPEHIBFFCB_proto_rawDescGZIP(), []int{0}
}

func (x *GGPEHIBFFCB) GetLNAIONDBECA() uint32 {
	if x != nil {
		return x.LNAIONDBECA
	}
	return 0
}

func (x *GGPEHIBFFCB) GetStatus() RogueBoothStatus {
	if x != nil {
		return x.Status
	}
	return RogueBoothStatus_ROGUE_BOOTH_NONE
}

func (x *GGPEHIBFFCB) GetHKJLPJBINNM() uint32 {
	if x != nil {
		return x.HKJLPJBINNM
	}
	return 0
}

var File_GGPEHIBFFCB_proto protoreflect.FileDescriptor

var file_GGPEHIBFFCB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x47, 0x50, 0x45, 0x48, 0x49, 0x42, 0x46, 0x46, 0x43, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6f, 0x74, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x0b, 0x47,
	0x47, 0x50, 0x45, 0x48, 0x49, 0x42, 0x46, 0x46, 0x43, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4e,
	0x41, 0x49, 0x4f, 0x4e, 0x44, 0x42, 0x45, 0x43, 0x41, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x4c, 0x4e, 0x41, 0x49, 0x4f, 0x4e, 0x44, 0x42, 0x45, 0x43, 0x41, 0x12, 0x29, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6f, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4b, 0x4a, 0x4c, 0x50,
	0x4a, 0x42, 0x49, 0x4e, 0x4e, 0x4d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4b,
	0x4a, 0x4c, 0x50, 0x4a, 0x42, 0x49, 0x4e, 0x4e, 0x4d, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_GGPEHIBFFCB_proto_rawDescOnce sync.Once
	file_GGPEHIBFFCB_proto_rawDescData = file_GGPEHIBFFCB_proto_rawDesc
)

func file_GGPEHIBFFCB_proto_rawDescGZIP() []byte {
	file_GGPEHIBFFCB_proto_rawDescOnce.Do(func() {
		file_GGPEHIBFFCB_proto_rawDescData = protoimpl.X.CompressGZIP(file_GGPEHIBFFCB_proto_rawDescData)
	})
	return file_GGPEHIBFFCB_proto_rawDescData
}

var file_GGPEHIBFFCB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GGPEHIBFFCB_proto_goTypes = []interface{}{
	(*GGPEHIBFFCB)(nil),   // 0: GGPEHIBFFCB
	(RogueBoothStatus)(0), // 1: RogueBoothStatus
}
var file_GGPEHIBFFCB_proto_depIdxs = []int32{
	1, // 0: GGPEHIBFFCB.status:type_name -> RogueBoothStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GGPEHIBFFCB_proto_init() }
func file_GGPEHIBFFCB_proto_init() {
	if File_GGPEHIBFFCB_proto != nil {
		return
	}
	file_RogueBoothStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GGPEHIBFFCB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GGPEHIBFFCB); i {
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
			RawDescriptor: file_GGPEHIBFFCB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GGPEHIBFFCB_proto_goTypes,
		DependencyIndexes: file_GGPEHIBFFCB_proto_depIdxs,
		MessageInfos:      file_GGPEHIBFFCB_proto_msgTypes,
	}.Build()
	File_GGPEHIBFFCB_proto = out.File
	file_GGPEHIBFFCB_proto_rawDesc = nil
	file_GGPEHIBFFCB_proto_goTypes = nil
	file_GGPEHIBFFCB_proto_depIdxs = nil
}
