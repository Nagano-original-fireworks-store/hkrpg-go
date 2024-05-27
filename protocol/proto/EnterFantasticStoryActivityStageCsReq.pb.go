// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EnterFantasticStoryActivityStageCsReq.proto

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

type EnterFantasticStoryActivityStageCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarList  []*IEOMPJKDIGL `protobuf:"bytes,15,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
	BattleId    uint32         `protobuf:"varint,8,opt,name=battle_id,json=battleId,proto3" json:"battle_id,omitempty"`
	BuffList    []uint32       `protobuf:"varint,9,rep,packed,name=buff_list,json=buffList,proto3" json:"buff_list,omitempty"`
	IPKKGLMMANH uint32         `protobuf:"varint,14,opt,name=IPKKGLMMANH,proto3" json:"IPKKGLMMANH,omitempty"`
}

func (x *EnterFantasticStoryActivityStageCsReq) Reset() {
	*x = EnterFantasticStoryActivityStageCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EnterFantasticStoryActivityStageCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterFantasticStoryActivityStageCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterFantasticStoryActivityStageCsReq) ProtoMessage() {}

func (x *EnterFantasticStoryActivityStageCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_EnterFantasticStoryActivityStageCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterFantasticStoryActivityStageCsReq.ProtoReflect.Descriptor instead.
func (*EnterFantasticStoryActivityStageCsReq) Descriptor() ([]byte, []int) {
	return file_EnterFantasticStoryActivityStageCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *EnterFantasticStoryActivityStageCsReq) GetAvatarList() []*IEOMPJKDIGL {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

func (x *EnterFantasticStoryActivityStageCsReq) GetBattleId() uint32 {
	if x != nil {
		return x.BattleId
	}
	return 0
}

func (x *EnterFantasticStoryActivityStageCsReq) GetBuffList() []uint32 {
	if x != nil {
		return x.BuffList
	}
	return nil
}

func (x *EnterFantasticStoryActivityStageCsReq) GetIPKKGLMMANH() uint32 {
	if x != nil {
		return x.IPKKGLMMANH
	}
	return 0
}

var File_EnterFantasticStoryActivityStageCsReq_proto protoreflect.FileDescriptor

var file_EnterFantasticStoryActivityStageCsReq_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x53, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49,
	0x45, 0x4f, 0x4d, 0x50, 0x4a, 0x4b, 0x44, 0x49, 0x47, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb2, 0x01, 0x0a, 0x25, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x0b, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x49, 0x45, 0x4f, 0x4d, 0x50, 0x4a, 0x4b, 0x44, 0x49, 0x47, 0x4c, 0x52, 0x0a, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x50, 0x4b, 0x4b, 0x47, 0x4c, 0x4d, 0x4d, 0x41,
	0x4e, 0x48, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x50, 0x4b, 0x4b, 0x47, 0x4c,
	0x4d, 0x4d, 0x41, 0x4e, 0x48, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EnterFantasticStoryActivityStageCsReq_proto_rawDescOnce sync.Once
	file_EnterFantasticStoryActivityStageCsReq_proto_rawDescData = file_EnterFantasticStoryActivityStageCsReq_proto_rawDesc
)

func file_EnterFantasticStoryActivityStageCsReq_proto_rawDescGZIP() []byte {
	file_EnterFantasticStoryActivityStageCsReq_proto_rawDescOnce.Do(func() {
		file_EnterFantasticStoryActivityStageCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnterFantasticStoryActivityStageCsReq_proto_rawDescData)
	})
	return file_EnterFantasticStoryActivityStageCsReq_proto_rawDescData
}

var file_EnterFantasticStoryActivityStageCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EnterFantasticStoryActivityStageCsReq_proto_goTypes = []interface{}{
	(*EnterFantasticStoryActivityStageCsReq)(nil), // 0: EnterFantasticStoryActivityStageCsReq
	(*IEOMPJKDIGL)(nil),                           // 1: IEOMPJKDIGL
}
var file_EnterFantasticStoryActivityStageCsReq_proto_depIdxs = []int32{
	1, // 0: EnterFantasticStoryActivityStageCsReq.avatar_list:type_name -> IEOMPJKDIGL
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EnterFantasticStoryActivityStageCsReq_proto_init() }
func file_EnterFantasticStoryActivityStageCsReq_proto_init() {
	if File_EnterFantasticStoryActivityStageCsReq_proto != nil {
		return
	}
	file_IEOMPJKDIGL_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EnterFantasticStoryActivityStageCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterFantasticStoryActivityStageCsReq); i {
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
			RawDescriptor: file_EnterFantasticStoryActivityStageCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnterFantasticStoryActivityStageCsReq_proto_goTypes,
		DependencyIndexes: file_EnterFantasticStoryActivityStageCsReq_proto_depIdxs,
		MessageInfos:      file_EnterFantasticStoryActivityStageCsReq_proto_msgTypes,
	}.Build()
	File_EnterFantasticStoryActivityStageCsReq_proto = out.File
	file_EnterFantasticStoryActivityStageCsReq_proto_rawDesc = nil
	file_EnterFantasticStoryActivityStageCsReq_proto_goTypes = nil
	file_EnterFantasticStoryActivityStageCsReq_proto_depIdxs = nil
}
