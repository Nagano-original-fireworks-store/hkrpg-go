// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerReturnTakePointRewardCsReq.proto

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

type PlayerReturnTakePointRewardCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FJFHEPECIMK uint32 `protobuf:"varint,7,opt,name=FJFHEPECIMK,proto3" json:"FJFHEPECIMK,omitempty"`
	ALHCIGCHBFJ uint32 `protobuf:"varint,13,opt,name=ALHCIGCHBFJ,proto3" json:"ALHCIGCHBFJ,omitempty"`
}

func (x *PlayerReturnTakePointRewardCsReq) Reset() {
	*x = PlayerReturnTakePointRewardCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerReturnTakePointRewardCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerReturnTakePointRewardCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerReturnTakePointRewardCsReq) ProtoMessage() {}

func (x *PlayerReturnTakePointRewardCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerReturnTakePointRewardCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerReturnTakePointRewardCsReq.ProtoReflect.Descriptor instead.
func (*PlayerReturnTakePointRewardCsReq) Descriptor() ([]byte, []int) {
	return file_PlayerReturnTakePointRewardCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerReturnTakePointRewardCsReq) GetFJFHEPECIMK() uint32 {
	if x != nil {
		return x.FJFHEPECIMK
	}
	return 0
}

func (x *PlayerReturnTakePointRewardCsReq) GetALHCIGCHBFJ() uint32 {
	if x != nil {
		return x.ALHCIGCHBFJ
	}
	return 0
}

var File_PlayerReturnTakePointRewardCsReq_proto protoreflect.FileDescriptor

var file_PlayerReturnTakePointRewardCsReq_proto_rawDesc = []byte{
	0x0a, 0x26, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54, 0x61,
	0x6b, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x20, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b,
	0x46, 0x4a, 0x46, 0x48, 0x45, 0x50, 0x45, 0x43, 0x49, 0x4d, 0x4b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x46, 0x4a, 0x46, 0x48, 0x45, 0x50, 0x45, 0x43, 0x49, 0x4d, 0x4b, 0x12, 0x20,
	0x0a, 0x0b, 0x41, 0x4c, 0x48, 0x43, 0x49, 0x47, 0x43, 0x48, 0x42, 0x46, 0x4a, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4c, 0x48, 0x43, 0x49, 0x47, 0x43, 0x48, 0x42, 0x46, 0x4a,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerReturnTakePointRewardCsReq_proto_rawDescOnce sync.Once
	file_PlayerReturnTakePointRewardCsReq_proto_rawDescData = file_PlayerReturnTakePointRewardCsReq_proto_rawDesc
)

func file_PlayerReturnTakePointRewardCsReq_proto_rawDescGZIP() []byte {
	file_PlayerReturnTakePointRewardCsReq_proto_rawDescOnce.Do(func() {
		file_PlayerReturnTakePointRewardCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerReturnTakePointRewardCsReq_proto_rawDescData)
	})
	return file_PlayerReturnTakePointRewardCsReq_proto_rawDescData
}

var file_PlayerReturnTakePointRewardCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerReturnTakePointRewardCsReq_proto_goTypes = []interface{}{
	(*PlayerReturnTakePointRewardCsReq)(nil), // 0: PlayerReturnTakePointRewardCsReq
}
var file_PlayerReturnTakePointRewardCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerReturnTakePointRewardCsReq_proto_init() }
func file_PlayerReturnTakePointRewardCsReq_proto_init() {
	if File_PlayerReturnTakePointRewardCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerReturnTakePointRewardCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerReturnTakePointRewardCsReq); i {
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
			RawDescriptor: file_PlayerReturnTakePointRewardCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerReturnTakePointRewardCsReq_proto_goTypes,
		DependencyIndexes: file_PlayerReturnTakePointRewardCsReq_proto_depIdxs,
		MessageInfos:      file_PlayerReturnTakePointRewardCsReq_proto_msgTypes,
	}.Build()
	File_PlayerReturnTakePointRewardCsReq_proto = out.File
	file_PlayerReturnTakePointRewardCsReq_proto_rawDesc = nil
	file_PlayerReturnTakePointRewardCsReq_proto_goTypes = nil
	file_PlayerReturnTakePointRewardCsReq_proto_depIdxs = nil
}
