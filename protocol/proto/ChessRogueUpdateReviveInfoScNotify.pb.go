// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueUpdateReviveInfoScNotify.proto

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

type ChessRogueUpdateReviveInfoScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviveInfo *RogueAvatarReviveCost `protobuf:"bytes,10,opt,name=revive_info,json=reviveInfo,proto3" json:"revive_info,omitempty"`
}

func (x *ChessRogueUpdateReviveInfoScNotify) Reset() {
	*x = ChessRogueUpdateReviveInfoScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueUpdateReviveInfoScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueUpdateReviveInfoScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueUpdateReviveInfoScNotify) ProtoMessage() {}

func (x *ChessRogueUpdateReviveInfoScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueUpdateReviveInfoScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueUpdateReviveInfoScNotify.ProtoReflect.Descriptor instead.
func (*ChessRogueUpdateReviveInfoScNotify) Descriptor() ([]byte, []int) {
	return file_ChessRogueUpdateReviveInfoScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueUpdateReviveInfoScNotify) GetReviveInfo() *RogueAvatarReviveCost {
	if x != nil {
		return x.ReviveInfo
	}
	return nil
}

var File_ChessRogueUpdateReviveInfoScNotify_proto protoreflect.FileDescriptor

var file_ChessRogueUpdateReviveInfoScNotify_proto_rawDesc = []byte{
	0x0a, 0x28, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x76, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x22, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x76,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x37, 0x0a,
	0x0b, 0x72, 0x65, 0x76, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x52, 0x65, 0x76, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69,
	0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueUpdateReviveInfoScNotify_proto_rawDescOnce sync.Once
	file_ChessRogueUpdateReviveInfoScNotify_proto_rawDescData = file_ChessRogueUpdateReviveInfoScNotify_proto_rawDesc
)

func file_ChessRogueUpdateReviveInfoScNotify_proto_rawDescGZIP() []byte {
	file_ChessRogueUpdateReviveInfoScNotify_proto_rawDescOnce.Do(func() {
		file_ChessRogueUpdateReviveInfoScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueUpdateReviveInfoScNotify_proto_rawDescData)
	})
	return file_ChessRogueUpdateReviveInfoScNotify_proto_rawDescData
}

var file_ChessRogueUpdateReviveInfoScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueUpdateReviveInfoScNotify_proto_goTypes = []interface{}{
	(*ChessRogueUpdateReviveInfoScNotify)(nil), // 0: ChessRogueUpdateReviveInfoScNotify
	(*RogueAvatarReviveCost)(nil),              // 1: RogueAvatarReviveCost
}
var file_ChessRogueUpdateReviveInfoScNotify_proto_depIdxs = []int32{
	1, // 0: ChessRogueUpdateReviveInfoScNotify.revive_info:type_name -> RogueAvatarReviveCost
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChessRogueUpdateReviveInfoScNotify_proto_init() }
func file_ChessRogueUpdateReviveInfoScNotify_proto_init() {
	if File_ChessRogueUpdateReviveInfoScNotify_proto != nil {
		return
	}
	file_RogueAvatarReviveCost_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueUpdateReviveInfoScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueUpdateReviveInfoScNotify); i {
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
			RawDescriptor: file_ChessRogueUpdateReviveInfoScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueUpdateReviveInfoScNotify_proto_goTypes,
		DependencyIndexes: file_ChessRogueUpdateReviveInfoScNotify_proto_depIdxs,
		MessageInfos:      file_ChessRogueUpdateReviveInfoScNotify_proto_msgTypes,
	}.Build()
	File_ChessRogueUpdateReviveInfoScNotify_proto = out.File
	file_ChessRogueUpdateReviveInfoScNotify_proto_rawDesc = nil
	file_ChessRogueUpdateReviveInfoScNotify_proto_goTypes = nil
	file_ChessRogueUpdateReviveInfoScNotify_proto_depIdxs = nil
}
