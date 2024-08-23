// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneCastSkillMpUpdateScNotify.proto

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

type SceneCastSkillMpUpdateScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mp           uint32 `protobuf:"varint,4,opt,name=mp,proto3" json:"mp,omitempty"`
	CastEntityId uint32 `protobuf:"varint,5,opt,name=cast_entity_id,json=castEntityId,proto3" json:"cast_entity_id,omitempty"`
}

func (x *SceneCastSkillMpUpdateScNotify) Reset() {
	*x = SceneCastSkillMpUpdateScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneCastSkillMpUpdateScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneCastSkillMpUpdateScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneCastSkillMpUpdateScNotify) ProtoMessage() {}

func (x *SceneCastSkillMpUpdateScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SceneCastSkillMpUpdateScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneCastSkillMpUpdateScNotify.ProtoReflect.Descriptor instead.
func (*SceneCastSkillMpUpdateScNotify) Descriptor() ([]byte, []int) {
	return file_SceneCastSkillMpUpdateScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SceneCastSkillMpUpdateScNotify) GetMp() uint32 {
	if x != nil {
		return x.Mp
	}
	return 0
}

func (x *SceneCastSkillMpUpdateScNotify) GetCastEntityId() uint32 {
	if x != nil {
		return x.CastEntityId
	}
	return 0
}

var File_SceneCastSkillMpUpdateScNotify_proto protoreflect.FileDescriptor

var file_SceneCastSkillMpUpdateScNotify_proto_rawDesc = []byte{
	0x0a, 0x24, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x61, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x4d, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x1e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43,
	0x61, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4d, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x6d, 0x70, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x61, 0x73, 0x74,
	0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x63, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneCastSkillMpUpdateScNotify_proto_rawDescOnce sync.Once
	file_SceneCastSkillMpUpdateScNotify_proto_rawDescData = file_SceneCastSkillMpUpdateScNotify_proto_rawDesc
)

func file_SceneCastSkillMpUpdateScNotify_proto_rawDescGZIP() []byte {
	file_SceneCastSkillMpUpdateScNotify_proto_rawDescOnce.Do(func() {
		file_SceneCastSkillMpUpdateScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneCastSkillMpUpdateScNotify_proto_rawDescData)
	})
	return file_SceneCastSkillMpUpdateScNotify_proto_rawDescData
}

var file_SceneCastSkillMpUpdateScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneCastSkillMpUpdateScNotify_proto_goTypes = []interface{}{
	(*SceneCastSkillMpUpdateScNotify)(nil), // 0: SceneCastSkillMpUpdateScNotify
}
var file_SceneCastSkillMpUpdateScNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneCastSkillMpUpdateScNotify_proto_init() }
func file_SceneCastSkillMpUpdateScNotify_proto_init() {
	if File_SceneCastSkillMpUpdateScNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneCastSkillMpUpdateScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneCastSkillMpUpdateScNotify); i {
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
			RawDescriptor: file_SceneCastSkillMpUpdateScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneCastSkillMpUpdateScNotify_proto_goTypes,
		DependencyIndexes: file_SceneCastSkillMpUpdateScNotify_proto_depIdxs,
		MessageInfos:      file_SceneCastSkillMpUpdateScNotify_proto_msgTypes,
	}.Build()
	File_SceneCastSkillMpUpdateScNotify_proto = out.File
	file_SceneCastSkillMpUpdateScNotify_proto_rawDesc = nil
	file_SceneCastSkillMpUpdateScNotify_proto_goTypes = nil
	file_SceneCastSkillMpUpdateScNotify_proto_depIdxs = nil
}
