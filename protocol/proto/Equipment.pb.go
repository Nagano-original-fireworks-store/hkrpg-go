// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: Equipment.proto

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

type Equipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank          uint32 `protobuf:"varint,13,opt,name=rank,proto3" json:"rank,omitempty"`
	BaseAvatarId  uint32 `protobuf:"varint,14,opt,name=base_avatar_id,json=baseAvatarId,proto3" json:"base_avatar_id,omitempty"`
	Promotion     uint32 `protobuf:"varint,5,opt,name=promotion,proto3" json:"promotion,omitempty"`
	UniqueId      uint32 `protobuf:"varint,8,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	Level         uint32 `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	IsProtected   bool   `protobuf:"varint,12,opt,name=is_protected,json=isProtected,proto3" json:"is_protected,omitempty"`
	Exp           uint32 `protobuf:"varint,7,opt,name=exp,proto3" json:"exp,omitempty"`
	DressAvatarId uint32 `protobuf:"varint,6,opt,name=dress_avatar_id,json=dressAvatarId,proto3" json:"dress_avatar_id,omitempty"`
	Tid           uint32 `protobuf:"varint,2,opt,name=tid,proto3" json:"tid,omitempty"`
}

func (x *Equipment) Reset() {
	*x = Equipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Equipment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Equipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Equipment) ProtoMessage() {}

func (x *Equipment) ProtoReflect() protoreflect.Message {
	mi := &file_Equipment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Equipment.ProtoReflect.Descriptor instead.
func (*Equipment) Descriptor() ([]byte, []int) {
	return file_Equipment_proto_rawDescGZIP(), []int{0}
}

func (x *Equipment) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Equipment) GetBaseAvatarId() uint32 {
	if x != nil {
		return x.BaseAvatarId
	}
	return 0
}

func (x *Equipment) GetPromotion() uint32 {
	if x != nil {
		return x.Promotion
	}
	return 0
}

func (x *Equipment) GetUniqueId() uint32 {
	if x != nil {
		return x.UniqueId
	}
	return 0
}

func (x *Equipment) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Equipment) GetIsProtected() bool {
	if x != nil {
		return x.IsProtected
	}
	return false
}

func (x *Equipment) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *Equipment) GetDressAvatarId() uint32 {
	if x != nil {
		return x.DressAvatarId
	}
	return 0
}

func (x *Equipment) GetTid() uint32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

var File_Equipment_proto protoreflect.FileDescriptor

var file_Equipment_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x85, 0x02, 0x0a, 0x09, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x62, 0x61, 0x73,
	0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x78, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12,
	0x26, 0x0a, 0x0f, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x64, 0x72, 0x65, 0x73, 0x73, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x69, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Equipment_proto_rawDescOnce sync.Once
	file_Equipment_proto_rawDescData = file_Equipment_proto_rawDesc
)

func file_Equipment_proto_rawDescGZIP() []byte {
	file_Equipment_proto_rawDescOnce.Do(func() {
		file_Equipment_proto_rawDescData = protoimpl.X.CompressGZIP(file_Equipment_proto_rawDescData)
	})
	return file_Equipment_proto_rawDescData
}

var file_Equipment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Equipment_proto_goTypes = []interface{}{
	(*Equipment)(nil), // 0: Equipment
}
var file_Equipment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Equipment_proto_init() }
func file_Equipment_proto_init() {
	if File_Equipment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Equipment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Equipment); i {
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
			RawDescriptor: file_Equipment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Equipment_proto_goTypes,
		DependencyIndexes: file_Equipment_proto_depIdxs,
		MessageInfos:      file_Equipment_proto_msgTypes,
	}.Build()
	File_Equipment_proto = out.File
	file_Equipment_proto_rawDesc = nil
	file_Equipment_proto_goTypes = nil
	file_Equipment_proto_depIdxs = nil
}
