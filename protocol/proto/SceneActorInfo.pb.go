// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneActorInfo.proto

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

type SceneActorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarType   AvatarType `protobuf:"varint,8,opt,name=avatar_type,json=avatarType,proto3,enum=AvatarType" json:"avatar_type,omitempty"`
	BaseAvatarId uint32     `protobuf:"varint,14,opt,name=base_avatar_id,json=baseAvatarId,proto3" json:"base_avatar_id,omitempty"`
	MapLayer     uint32     `protobuf:"varint,6,opt,name=map_layer,json=mapLayer,proto3" json:"map_layer,omitempty"`
	Uid          uint32     `protobuf:"varint,15,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *SceneActorInfo) Reset() {
	*x = SceneActorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneActorInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneActorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneActorInfo) ProtoMessage() {}

func (x *SceneActorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneActorInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneActorInfo.ProtoReflect.Descriptor instead.
func (*SceneActorInfo) Descriptor() ([]byte, []int) {
	return file_SceneActorInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneActorInfo) GetAvatarType() AvatarType {
	if x != nil {
		return x.AvatarType
	}
	return AvatarType_AVATAR_TYPE_NONE
}

func (x *SceneActorInfo) GetBaseAvatarId() uint32 {
	if x != nil {
		return x.BaseAvatarId
	}
	return 0
}

func (x *SceneActorInfo) GetMapLayer() uint32 {
	if x != nil {
		return x.MapLayer
	}
	return 0
}

func (x *SceneActorInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_SceneActorInfo_proto protoreflect.FileDescriptor

var file_SceneActorInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0e, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x0b, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_SceneActorInfo_proto_rawDescOnce sync.Once
	file_SceneActorInfo_proto_rawDescData = file_SceneActorInfo_proto_rawDesc
)

func file_SceneActorInfo_proto_rawDescGZIP() []byte {
	file_SceneActorInfo_proto_rawDescOnce.Do(func() {
		file_SceneActorInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneActorInfo_proto_rawDescData)
	})
	return file_SceneActorInfo_proto_rawDescData
}

var file_SceneActorInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneActorInfo_proto_goTypes = []interface{}{
	(*SceneActorInfo)(nil), // 0: SceneActorInfo
	(AvatarType)(0),        // 1: AvatarType
}
var file_SceneActorInfo_proto_depIdxs = []int32{
	1, // 0: SceneActorInfo.avatar_type:type_name -> AvatarType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneActorInfo_proto_init() }
func file_SceneActorInfo_proto_init() {
	if File_SceneActorInfo_proto != nil {
		return
	}
	file_AvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneActorInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneActorInfo); i {
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
			RawDescriptor: file_SceneActorInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneActorInfo_proto_goTypes,
		DependencyIndexes: file_SceneActorInfo_proto_depIdxs,
		MessageInfos:      file_SceneActorInfo_proto_msgTypes,
	}.Build()
	File_SceneActorInfo_proto = out.File
	file_SceneActorInfo_proto_rawDesc = nil
	file_SceneActorInfo_proto_goTypes = nil
	file_SceneActorInfo_proto_depIdxs = nil
}
