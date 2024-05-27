// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerGetTokenScRsp.proto

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

type PlayerGetTokenScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode       uint32     `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Msg           string     `protobuf:"bytes,11,opt,name=msg,proto3" json:"msg,omitempty"`
	Uid           uint32     `protobuf:"varint,6,opt,name=uid,proto3" json:"uid,omitempty"`
	BlackInfo     *BlackInfo `protobuf:"bytes,13,opt,name=black_info,json=blackInfo,proto3" json:"black_info,omitempty"`
	SecretKeySeed uint64     `protobuf:"varint,7,opt,name=secret_key_seed,json=secretKeySeed,proto3" json:"secret_key_seed,omitempty"`
}

func (x *PlayerGetTokenScRsp) Reset() {
	*x = PlayerGetTokenScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerGetTokenScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerGetTokenScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerGetTokenScRsp) ProtoMessage() {}

func (x *PlayerGetTokenScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerGetTokenScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerGetTokenScRsp.ProtoReflect.Descriptor instead.
func (*PlayerGetTokenScRsp) Descriptor() ([]byte, []int) {
	return file_PlayerGetTokenScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerGetTokenScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *PlayerGetTokenScRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PlayerGetTokenScRsp) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PlayerGetTokenScRsp) GetBlackInfo() *BlackInfo {
	if x != nil {
		return x.BlackInfo
	}
	return nil
}

func (x *PlayerGetTokenScRsp) GetSecretKeySeed() uint64 {
	if x != nil {
		return x.SecretKeySeed
	}
	return 0
}

var File_PlayerGetTokenScRsp_proto protoreflect.FileDescriptor

var file_PlayerGetTokenScRsp_proto_rawDesc = []byte{
	0x0a, 0x19, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a,
	0x13, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x29, 0x0a, 0x0a, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x09, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a,
	0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x53, 0x65, 0x65, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerGetTokenScRsp_proto_rawDescOnce sync.Once
	file_PlayerGetTokenScRsp_proto_rawDescData = file_PlayerGetTokenScRsp_proto_rawDesc
)

func file_PlayerGetTokenScRsp_proto_rawDescGZIP() []byte {
	file_PlayerGetTokenScRsp_proto_rawDescOnce.Do(func() {
		file_PlayerGetTokenScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerGetTokenScRsp_proto_rawDescData)
	})
	return file_PlayerGetTokenScRsp_proto_rawDescData
}

var file_PlayerGetTokenScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerGetTokenScRsp_proto_goTypes = []interface{}{
	(*PlayerGetTokenScRsp)(nil), // 0: PlayerGetTokenScRsp
	(*BlackInfo)(nil),           // 1: BlackInfo
}
var file_PlayerGetTokenScRsp_proto_depIdxs = []int32{
	1, // 0: PlayerGetTokenScRsp.black_info:type_name -> BlackInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerGetTokenScRsp_proto_init() }
func file_PlayerGetTokenScRsp_proto_init() {
	if File_PlayerGetTokenScRsp_proto != nil {
		return
	}
	file_BlackInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerGetTokenScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerGetTokenScRsp); i {
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
			RawDescriptor: file_PlayerGetTokenScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerGetTokenScRsp_proto_goTypes,
		DependencyIndexes: file_PlayerGetTokenScRsp_proto_depIdxs,
		MessageInfos:      file_PlayerGetTokenScRsp_proto_msgTypes,
	}.Build()
	File_PlayerGetTokenScRsp_proto = out.File
	file_PlayerGetTokenScRsp_proto_rawDesc = nil
	file_PlayerGetTokenScRsp_proto_goTypes = nil
	file_PlayerGetTokenScRsp_proto_depIdxs = nil
}
