// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetMonopolyFriendRankingListScRsp.proto

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

type GetMonopolyFriendRankingListScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HBJCAKIIGMF *LHADAFIHEFE   `protobuf:"bytes,1,opt,name=HBJCAKIIGMF,proto3" json:"HBJCAKIIGMF,omitempty"`
	Retcode     uint32         `protobuf:"varint,6,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MBMIPKJFEPA []*LHADAFIHEFE `protobuf:"bytes,3,rep,name=MBMIPKJFEPA,proto3" json:"MBMIPKJFEPA,omitempty"`
}

func (x *GetMonopolyFriendRankingListScRsp) Reset() {
	*x = GetMonopolyFriendRankingListScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMonopolyFriendRankingListScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMonopolyFriendRankingListScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMonopolyFriendRankingListScRsp) ProtoMessage() {}

func (x *GetMonopolyFriendRankingListScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetMonopolyFriendRankingListScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMonopolyFriendRankingListScRsp.ProtoReflect.Descriptor instead.
func (*GetMonopolyFriendRankingListScRsp) Descriptor() ([]byte, []int) {
	return file_GetMonopolyFriendRankingListScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetMonopolyFriendRankingListScRsp) GetHBJCAKIIGMF() *LHADAFIHEFE {
	if x != nil {
		return x.HBJCAKIIGMF
	}
	return nil
}

func (x *GetMonopolyFriendRankingListScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMonopolyFriendRankingListScRsp) GetMBMIPKJFEPA() []*LHADAFIHEFE {
	if x != nil {
		return x.MBMIPKJFEPA
	}
	return nil
}

var File_GetMonopolyFriendRankingListScRsp_proto protoreflect.FileDescriptor

var file_GetMonopolyFriendRankingListScRsp_proto_rawDesc = []byte{
	0x0a, 0x27, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x48, 0x41, 0x44, 0x41,
	0x46, 0x49, 0x48, 0x45, 0x46, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a,
	0x21, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x48, 0x42, 0x4a, 0x43, 0x41, 0x4b, 0x49, 0x49, 0x47, 0x4d,
	0x46, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x48, 0x41, 0x44, 0x41, 0x46,
	0x49, 0x48, 0x45, 0x46, 0x45, 0x52, 0x0b, 0x48, 0x42, 0x4a, 0x43, 0x41, 0x4b, 0x49, 0x49, 0x47,
	0x4d, 0x46, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b,
	0x4d, 0x42, 0x4d, 0x49, 0x50, 0x4b, 0x4a, 0x46, 0x45, 0x50, 0x41, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x48, 0x41, 0x44, 0x41, 0x46, 0x49, 0x48, 0x45, 0x46, 0x45, 0x52,
	0x0b, 0x4d, 0x42, 0x4d, 0x49, 0x50, 0x4b, 0x4a, 0x46, 0x45, 0x50, 0x41, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetMonopolyFriendRankingListScRsp_proto_rawDescOnce sync.Once
	file_GetMonopolyFriendRankingListScRsp_proto_rawDescData = file_GetMonopolyFriendRankingListScRsp_proto_rawDesc
)

func file_GetMonopolyFriendRankingListScRsp_proto_rawDescGZIP() []byte {
	file_GetMonopolyFriendRankingListScRsp_proto_rawDescOnce.Do(func() {
		file_GetMonopolyFriendRankingListScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMonopolyFriendRankingListScRsp_proto_rawDescData)
	})
	return file_GetMonopolyFriendRankingListScRsp_proto_rawDescData
}

var file_GetMonopolyFriendRankingListScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetMonopolyFriendRankingListScRsp_proto_goTypes = []interface{}{
	(*GetMonopolyFriendRankingListScRsp)(nil), // 0: GetMonopolyFriendRankingListScRsp
	(*LHADAFIHEFE)(nil),                       // 1: LHADAFIHEFE
}
var file_GetMonopolyFriendRankingListScRsp_proto_depIdxs = []int32{
	1, // 0: GetMonopolyFriendRankingListScRsp.HBJCAKIIGMF:type_name -> LHADAFIHEFE
	1, // 1: GetMonopolyFriendRankingListScRsp.MBMIPKJFEPA:type_name -> LHADAFIHEFE
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetMonopolyFriendRankingListScRsp_proto_init() }
func file_GetMonopolyFriendRankingListScRsp_proto_init() {
	if File_GetMonopolyFriendRankingListScRsp_proto != nil {
		return
	}
	file_LHADAFIHEFE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetMonopolyFriendRankingListScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMonopolyFriendRankingListScRsp); i {
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
			RawDescriptor: file_GetMonopolyFriendRankingListScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMonopolyFriendRankingListScRsp_proto_goTypes,
		DependencyIndexes: file_GetMonopolyFriendRankingListScRsp_proto_depIdxs,
		MessageInfos:      file_GetMonopolyFriendRankingListScRsp_proto_msgTypes,
	}.Build()
	File_GetMonopolyFriendRankingListScRsp_proto = out.File
	file_GetMonopolyFriendRankingListScRsp_proto_rawDesc = nil
	file_GetMonopolyFriendRankingListScRsp_proto_goTypes = nil
	file_GetMonopolyFriendRankingListScRsp_proto_depIdxs = nil
}
