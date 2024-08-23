// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: JoinLineupCsReq.proto

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

type JoinLineupCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index           uint32          `protobuf:"varint,11,opt,name=index,proto3" json:"index,omitempty"`
	BaseAvatarId    uint32          `protobuf:"varint,12,opt,name=base_avatar_id,json=baseAvatarId,proto3" json:"base_avatar_id,omitempty"`
	IsVirtual       bool            `protobuf:"varint,3,opt,name=is_virtual,json=isVirtual,proto3" json:"is_virtual,omitempty"`
	Slot            uint32          `protobuf:"varint,6,opt,name=slot,proto3" json:"slot,omitempty"`
	ExtraLineupType ExtraLineupType `protobuf:"varint,10,opt,name=extra_lineup_type,json=extraLineupType,proto3,enum=ExtraLineupType" json:"extra_lineup_type,omitempty"`
	AvatarType      AvatarType      `protobuf:"varint,7,opt,name=avatar_type,json=avatarType,proto3,enum=AvatarType" json:"avatar_type,omitempty"`
	PlaneId         uint32          `protobuf:"varint,5,opt,name=plane_id,json=planeId,proto3" json:"plane_id,omitempty"`
}

func (x *JoinLineupCsReq) Reset() {
	*x = JoinLineupCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JoinLineupCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinLineupCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinLineupCsReq) ProtoMessage() {}

func (x *JoinLineupCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_JoinLineupCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinLineupCsReq.ProtoReflect.Descriptor instead.
func (*JoinLineupCsReq) Descriptor() ([]byte, []int) {
	return file_JoinLineupCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *JoinLineupCsReq) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *JoinLineupCsReq) GetBaseAvatarId() uint32 {
	if x != nil {
		return x.BaseAvatarId
	}
	return 0
}

func (x *JoinLineupCsReq) GetIsVirtual() bool {
	if x != nil {
		return x.IsVirtual
	}
	return false
}

func (x *JoinLineupCsReq) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *JoinLineupCsReq) GetExtraLineupType() ExtraLineupType {
	if x != nil {
		return x.ExtraLineupType
	}
	return ExtraLineupType_LINEUP_NONE
}

func (x *JoinLineupCsReq) GetAvatarType() AvatarType {
	if x != nil {
		return x.AvatarType
	}
	return AvatarType_AVATAR_TYPE_NONE
}

func (x *JoinLineupCsReq) GetPlaneId() uint32 {
	if x != nil {
		return x.PlaneId
	}
	return 0
}

var File_JoinLineupCsReq_proto protoreflect.FileDescriptor

var file_JoinLineupCsReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x43, 0x73, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69,
	0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x87, 0x02, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73,
	0x6c, 0x6f, 0x74, 0x12, 0x3c, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6c, 0x69, 0x6e,
	0x65, 0x75, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_JoinLineupCsReq_proto_rawDescOnce sync.Once
	file_JoinLineupCsReq_proto_rawDescData = file_JoinLineupCsReq_proto_rawDesc
)

func file_JoinLineupCsReq_proto_rawDescGZIP() []byte {
	file_JoinLineupCsReq_proto_rawDescOnce.Do(func() {
		file_JoinLineupCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_JoinLineupCsReq_proto_rawDescData)
	})
	return file_JoinLineupCsReq_proto_rawDescData
}

var file_JoinLineupCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_JoinLineupCsReq_proto_goTypes = []interface{}{
	(*JoinLineupCsReq)(nil), // 0: JoinLineupCsReq
	(ExtraLineupType)(0),    // 1: ExtraLineupType
	(AvatarType)(0),         // 2: AvatarType
}
var file_JoinLineupCsReq_proto_depIdxs = []int32{
	1, // 0: JoinLineupCsReq.extra_lineup_type:type_name -> ExtraLineupType
	2, // 1: JoinLineupCsReq.avatar_type:type_name -> AvatarType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_JoinLineupCsReq_proto_init() }
func file_JoinLineupCsReq_proto_init() {
	if File_JoinLineupCsReq_proto != nil {
		return
	}
	file_ExtraLineupType_proto_init()
	file_AvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_JoinLineupCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinLineupCsReq); i {
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
			RawDescriptor: file_JoinLineupCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JoinLineupCsReq_proto_goTypes,
		DependencyIndexes: file_JoinLineupCsReq_proto_depIdxs,
		MessageInfos:      file_JoinLineupCsReq_proto_msgTypes,
	}.Build()
	File_JoinLineupCsReq_proto = out.File
	file_JoinLineupCsReq_proto_rawDesc = nil
	file_JoinLineupCsReq_proto_goTypes = nil
	file_JoinLineupCsReq_proto_depIdxs = nil
}
