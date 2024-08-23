// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GroupRefreshInfo.proto

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

type GroupRefreshInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State         uint32                    `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	RefreshType   SceneGroupRefreshType     `protobuf:"varint,7,opt,name=refresh_type,json=refreshType,proto3,enum=SceneGroupRefreshType" json:"refresh_type,omitempty"`
	RefreshEntity []*SceneEntityRefreshInfo `protobuf:"bytes,5,rep,name=refresh_entity,json=refreshEntity,proto3" json:"refresh_entity,omitempty"`
	GroupId       uint32                    `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GroupRefreshInfo) Reset() {
	*x = GroupRefreshInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GroupRefreshInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupRefreshInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRefreshInfo) ProtoMessage() {}

func (x *GroupRefreshInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GroupRefreshInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRefreshInfo.ProtoReflect.Descriptor instead.
func (*GroupRefreshInfo) Descriptor() ([]byte, []int) {
	return file_GroupRefreshInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GroupRefreshInfo) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *GroupRefreshInfo) GetRefreshType() SceneGroupRefreshType {
	if x != nil {
		return x.RefreshType
	}
	return SceneGroupRefreshType_SCENE_GROUP_REFRESH_TYPE_NONE
}

func (x *GroupRefreshInfo) GetRefreshEntity() []*SceneEntityRefreshInfo {
	if x != nil {
		return x.RefreshEntity
	}
	return nil
}

func (x *GroupRefreshInfo) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_GroupRefreshInfo_proto protoreflect.FileDescriptor

var file_GroupRefreshInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x10, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39,
	0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x0e, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GroupRefreshInfo_proto_rawDescOnce sync.Once
	file_GroupRefreshInfo_proto_rawDescData = file_GroupRefreshInfo_proto_rawDesc
)

func file_GroupRefreshInfo_proto_rawDescGZIP() []byte {
	file_GroupRefreshInfo_proto_rawDescOnce.Do(func() {
		file_GroupRefreshInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GroupRefreshInfo_proto_rawDescData)
	})
	return file_GroupRefreshInfo_proto_rawDescData
}

var file_GroupRefreshInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GroupRefreshInfo_proto_goTypes = []interface{}{
	(*GroupRefreshInfo)(nil),       // 0: GroupRefreshInfo
	(SceneGroupRefreshType)(0),     // 1: SceneGroupRefreshType
	(*SceneEntityRefreshInfo)(nil), // 2: SceneEntityRefreshInfo
}
var file_GroupRefreshInfo_proto_depIdxs = []int32{
	1, // 0: GroupRefreshInfo.refresh_type:type_name -> SceneGroupRefreshType
	2, // 1: GroupRefreshInfo.refresh_entity:type_name -> SceneEntityRefreshInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GroupRefreshInfo_proto_init() }
func file_GroupRefreshInfo_proto_init() {
	if File_GroupRefreshInfo_proto != nil {
		return
	}
	file_SceneEntityRefreshInfo_proto_init()
	file_SceneGroupRefreshType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GroupRefreshInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupRefreshInfo); i {
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
			RawDescriptor: file_GroupRefreshInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GroupRefreshInfo_proto_goTypes,
		DependencyIndexes: file_GroupRefreshInfo_proto_depIdxs,
		MessageInfos:      file_GroupRefreshInfo_proto_msgTypes,
	}.Build()
	File_GroupRefreshInfo_proto = out.File
	file_GroupRefreshInfo_proto_rawDesc = nil
	file_GroupRefreshInfo_proto_goTypes = nil
	file_GroupRefreshInfo_proto_depIdxs = nil
}
