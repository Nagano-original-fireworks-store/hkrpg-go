// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetCrossInfoScRsp.proto

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

type GetCrossInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     uint32        `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ILGJLGKPDON uint64        `protobuf:"varint,10,opt,name=ILGJLGKPDON,proto3" json:"ILGJLGKPDON,omitempty"`
	MHCFIEHGNCE FightGameMode `protobuf:"varint,11,opt,name=MHCFIEHGNCE,proto3,enum=FightGameMode" json:"MHCFIEHGNCE,omitempty"`
	RoomId      uint64        `protobuf:"varint,8,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *GetCrossInfoScRsp) Reset() {
	*x = GetCrossInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetCrossInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrossInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrossInfoScRsp) ProtoMessage() {}

func (x *GetCrossInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetCrossInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrossInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetCrossInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetCrossInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetCrossInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetCrossInfoScRsp) GetILGJLGKPDON() uint64 {
	if x != nil {
		return x.ILGJLGKPDON
	}
	return 0
}

func (x *GetCrossInfoScRsp) GetMHCFIEHGNCE() FightGameMode {
	if x != nil {
		return x.MHCFIEHGNCE
	}
	return FightGameMode_FIGHT_GAME_MODE_NONE
}

func (x *GetCrossInfoScRsp) GetRoomId() uint64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

var File_GetCrossInfoScRsp_proto protoreflect.FileDescriptor

var file_GetCrossInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x46, 0x69, 0x67, 0x68, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a,
	0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x49, 0x4c, 0x47, 0x4a, 0x4c, 0x47, 0x4b, 0x50, 0x44, 0x4f, 0x4e, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x49, 0x4c, 0x47, 0x4a, 0x4c, 0x47, 0x4b, 0x50, 0x44, 0x4f, 0x4e,
	0x12, 0x30, 0x0a, 0x0b, 0x4d, 0x48, 0x43, 0x46, 0x49, 0x45, 0x48, 0x47, 0x4e, 0x43, 0x45, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x46, 0x69, 0x67, 0x68, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x4d, 0x48, 0x43, 0x46, 0x49, 0x45, 0x48, 0x47, 0x4e,
	0x43, 0x45, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetCrossInfoScRsp_proto_rawDescOnce sync.Once
	file_GetCrossInfoScRsp_proto_rawDescData = file_GetCrossInfoScRsp_proto_rawDesc
)

func file_GetCrossInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetCrossInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetCrossInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetCrossInfoScRsp_proto_rawDescData)
	})
	return file_GetCrossInfoScRsp_proto_rawDescData
}

var file_GetCrossInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetCrossInfoScRsp_proto_goTypes = []interface{}{
	(*GetCrossInfoScRsp)(nil), // 0: GetCrossInfoScRsp
	(FightGameMode)(0),        // 1: FightGameMode
}
var file_GetCrossInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetCrossInfoScRsp.MHCFIEHGNCE:type_name -> FightGameMode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetCrossInfoScRsp_proto_init() }
func file_GetCrossInfoScRsp_proto_init() {
	if File_GetCrossInfoScRsp_proto != nil {
		return
	}
	file_FightGameMode_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetCrossInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrossInfoScRsp); i {
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
			RawDescriptor: file_GetCrossInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetCrossInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetCrossInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetCrossInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetCrossInfoScRsp_proto = out.File
	file_GetCrossInfoScRsp_proto_rawDesc = nil
	file_GetCrossInfoScRsp_proto_goTypes = nil
	file_GetCrossInfoScRsp_proto_depIdxs = nil
}
