// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GroupStateInfo.proto

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

type GroupStateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OEKBJNJEEDF uint32 `protobuf:"varint,2,opt,name=OEKBJNJEEDF,proto3" json:"OEKBJNJEEDF,omitempty"`
	GroupId     uint32 `protobuf:"varint,5,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	EntryId     uint32 `protobuf:"varint,4,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	GroupState  uint32 `protobuf:"varint,3,opt,name=group_state,json=groupState,proto3" json:"group_state,omitempty"`
}

func (x *GroupStateInfo) Reset() {
	*x = GroupStateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GroupStateInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupStateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupStateInfo) ProtoMessage() {}

func (x *GroupStateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GroupStateInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupStateInfo.ProtoReflect.Descriptor instead.
func (*GroupStateInfo) Descriptor() ([]byte, []int) {
	return file_GroupStateInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GroupStateInfo) GetOEKBJNJEEDF() uint32 {
	if x != nil {
		return x.OEKBJNJEEDF
	}
	return 0
}

func (x *GroupStateInfo) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupStateInfo) GetEntryId() uint32 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

func (x *GroupStateInfo) GetGroupState() uint32 {
	if x != nil {
		return x.GroupState
	}
	return 0
}

var File_GroupStateInfo_proto protoreflect.FileDescriptor

var file_GroupStateInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x0e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x45, 0x4b,
	0x42, 0x4a, 0x4e, 0x4a, 0x45, 0x45, 0x44, 0x46, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4f, 0x45, 0x4b, 0x42, 0x4a, 0x4e, 0x4a, 0x45, 0x45, 0x44, 0x46, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GroupStateInfo_proto_rawDescOnce sync.Once
	file_GroupStateInfo_proto_rawDescData = file_GroupStateInfo_proto_rawDesc
)

func file_GroupStateInfo_proto_rawDescGZIP() []byte {
	file_GroupStateInfo_proto_rawDescOnce.Do(func() {
		file_GroupStateInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GroupStateInfo_proto_rawDescData)
	})
	return file_GroupStateInfo_proto_rawDescData
}

var file_GroupStateInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GroupStateInfo_proto_goTypes = []interface{}{
	(*GroupStateInfo)(nil), // 0: GroupStateInfo
}
var file_GroupStateInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GroupStateInfo_proto_init() }
func file_GroupStateInfo_proto_init() {
	if File_GroupStateInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GroupStateInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupStateInfo); i {
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
			RawDescriptor: file_GroupStateInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GroupStateInfo_proto_goTypes,
		DependencyIndexes: file_GroupStateInfo_proto_depIdxs,
		MessageInfos:      file_GroupStateInfo_proto_msgTypes,
	}.Build()
	File_GroupStateInfo_proto = out.File
	file_GroupStateInfo_proto_rawDesc = nil
	file_GroupStateInfo_proto_goTypes = nil
	file_GroupStateInfo_proto_depIdxs = nil
}
