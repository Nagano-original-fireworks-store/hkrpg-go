// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EnterChessRogueAeonRoomCsReq.proto

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

type EnterChessRogueAeonRoomCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnterChessRogueAeonRoomCsReq) Reset() {
	*x = EnterChessRogueAeonRoomCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EnterChessRogueAeonRoomCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterChessRogueAeonRoomCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterChessRogueAeonRoomCsReq) ProtoMessage() {}

func (x *EnterChessRogueAeonRoomCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_EnterChessRogueAeonRoomCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterChessRogueAeonRoomCsReq.ProtoReflect.Descriptor instead.
func (*EnterChessRogueAeonRoomCsReq) Descriptor() ([]byte, []int) {
	return file_EnterChessRogueAeonRoomCsReq_proto_rawDescGZIP(), []int{0}
}

var File_EnterChessRogueAeonRoomCsReq_proto protoreflect.FileDescriptor

var file_EnterChessRogueAeonRoomCsReq_proto_rawDesc = []byte{
	0x0a, 0x22, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x41, 0x65, 0x6f, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x1c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EnterChessRogueAeonRoomCsReq_proto_rawDescOnce sync.Once
	file_EnterChessRogueAeonRoomCsReq_proto_rawDescData = file_EnterChessRogueAeonRoomCsReq_proto_rawDesc
)

func file_EnterChessRogueAeonRoomCsReq_proto_rawDescGZIP() []byte {
	file_EnterChessRogueAeonRoomCsReq_proto_rawDescOnce.Do(func() {
		file_EnterChessRogueAeonRoomCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnterChessRogueAeonRoomCsReq_proto_rawDescData)
	})
	return file_EnterChessRogueAeonRoomCsReq_proto_rawDescData
}

var file_EnterChessRogueAeonRoomCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EnterChessRogueAeonRoomCsReq_proto_goTypes = []interface{}{
	(*EnterChessRogueAeonRoomCsReq)(nil), // 0: EnterChessRogueAeonRoomCsReq
}
var file_EnterChessRogueAeonRoomCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EnterChessRogueAeonRoomCsReq_proto_init() }
func file_EnterChessRogueAeonRoomCsReq_proto_init() {
	if File_EnterChessRogueAeonRoomCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EnterChessRogueAeonRoomCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterChessRogueAeonRoomCsReq); i {
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
			RawDescriptor: file_EnterChessRogueAeonRoomCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnterChessRogueAeonRoomCsReq_proto_goTypes,
		DependencyIndexes: file_EnterChessRogueAeonRoomCsReq_proto_depIdxs,
		MessageInfos:      file_EnterChessRogueAeonRoomCsReq_proto_msgTypes,
	}.Build()
	File_EnterChessRogueAeonRoomCsReq_proto = out.File
	file_EnterChessRogueAeonRoomCsReq_proto_rawDesc = nil
	file_EnterChessRogueAeonRoomCsReq_proto_goTypes = nil
	file_EnterChessRogueAeonRoomCsReq_proto_depIdxs = nil
}
