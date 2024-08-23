// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SpringRecoverSingleAvatarCsReq.proto

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

type SpringRecoverSingleAvatarCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaneId        uint32     `protobuf:"varint,7,opt,name=plane_id,json=planeId,proto3" json:"plane_id,omitempty"`
	AvatarType     AvatarType `protobuf:"varint,10,opt,name=avatar_type,json=avatarType,proto3,enum=AvatarType" json:"avatar_type,omitempty"`
	FloorId        uint32     `protobuf:"varint,12,opt,name=floor_id,json=floorId,proto3" json:"floor_id,omitempty"`
	Id             uint32     `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	PropEntityId   uint32     `protobuf:"varint,13,opt,name=prop_entity_id,json=propEntityId,proto3" json:"prop_entity_id,omitempty"`
	IsFullyRecover bool       `protobuf:"varint,11,opt,name=is_fully_recover,json=isFullyRecover,proto3" json:"is_fully_recover,omitempty"`
}

func (x *SpringRecoverSingleAvatarCsReq) Reset() {
	*x = SpringRecoverSingleAvatarCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SpringRecoverSingleAvatarCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpringRecoverSingleAvatarCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpringRecoverSingleAvatarCsReq) ProtoMessage() {}

func (x *SpringRecoverSingleAvatarCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SpringRecoverSingleAvatarCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpringRecoverSingleAvatarCsReq.ProtoReflect.Descriptor instead.
func (*SpringRecoverSingleAvatarCsReq) Descriptor() ([]byte, []int) {
	return file_SpringRecoverSingleAvatarCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SpringRecoverSingleAvatarCsReq) GetPlaneId() uint32 {
	if x != nil {
		return x.PlaneId
	}
	return 0
}

func (x *SpringRecoverSingleAvatarCsReq) GetAvatarType() AvatarType {
	if x != nil {
		return x.AvatarType
	}
	return AvatarType_AVATAR_TYPE_NONE
}

func (x *SpringRecoverSingleAvatarCsReq) GetFloorId() uint32 {
	if x != nil {
		return x.FloorId
	}
	return 0
}

func (x *SpringRecoverSingleAvatarCsReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SpringRecoverSingleAvatarCsReq) GetPropEntityId() uint32 {
	if x != nil {
		return x.PropEntityId
	}
	return 0
}

func (x *SpringRecoverSingleAvatarCsReq) GetIsFullyRecover() bool {
	if x != nil {
		return x.IsFullyRecover
	}
	return false
}

var File_SpringRecoverSingleAvatarCsReq_proto protoreflect.FileDescriptor

var file_SpringRecoverSingleAvatarCsReq_proto_rawDesc = []byte{
	0x0a, 0x24, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x73, 0x52, 0x65, 0x71,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x1e, 0x53, 0x70, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x24, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x66, 0x75, 0x6c, 0x6c,
	0x79, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x69, 0x73, 0x46, 0x75, 0x6c, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SpringRecoverSingleAvatarCsReq_proto_rawDescOnce sync.Once
	file_SpringRecoverSingleAvatarCsReq_proto_rawDescData = file_SpringRecoverSingleAvatarCsReq_proto_rawDesc
)

func file_SpringRecoverSingleAvatarCsReq_proto_rawDescGZIP() []byte {
	file_SpringRecoverSingleAvatarCsReq_proto_rawDescOnce.Do(func() {
		file_SpringRecoverSingleAvatarCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SpringRecoverSingleAvatarCsReq_proto_rawDescData)
	})
	return file_SpringRecoverSingleAvatarCsReq_proto_rawDescData
}

var file_SpringRecoverSingleAvatarCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SpringRecoverSingleAvatarCsReq_proto_goTypes = []interface{}{
	(*SpringRecoverSingleAvatarCsReq)(nil), // 0: SpringRecoverSingleAvatarCsReq
	(AvatarType)(0),                        // 1: AvatarType
}
var file_SpringRecoverSingleAvatarCsReq_proto_depIdxs = []int32{
	1, // 0: SpringRecoverSingleAvatarCsReq.avatar_type:type_name -> AvatarType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SpringRecoverSingleAvatarCsReq_proto_init() }
func file_SpringRecoverSingleAvatarCsReq_proto_init() {
	if File_SpringRecoverSingleAvatarCsReq_proto != nil {
		return
	}
	file_AvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SpringRecoverSingleAvatarCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpringRecoverSingleAvatarCsReq); i {
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
			RawDescriptor: file_SpringRecoverSingleAvatarCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SpringRecoverSingleAvatarCsReq_proto_goTypes,
		DependencyIndexes: file_SpringRecoverSingleAvatarCsReq_proto_depIdxs,
		MessageInfos:      file_SpringRecoverSingleAvatarCsReq_proto_msgTypes,
	}.Build()
	File_SpringRecoverSingleAvatarCsReq_proto = out.File
	file_SpringRecoverSingleAvatarCsReq_proto_rawDesc = nil
	file_SpringRecoverSingleAvatarCsReq_proto_goTypes = nil
	file_SpringRecoverSingleAvatarCsReq_proto_depIdxs = nil
}
