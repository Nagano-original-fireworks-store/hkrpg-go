// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FMBLGECBIBP.proto

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

type FMBLGECBIBP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KEJHMBALBGN uint32 `protobuf:"varint,12,opt,name=KEJHMBALBGN,proto3" json:"KEJHMBALBGN,omitempty"`
	OMKOINCNCCB uint32 `protobuf:"varint,9,opt,name=OMKOINCNCCB,proto3" json:"OMKOINCNCCB,omitempty"`
	HNFEIDJPLJK bool   `protobuf:"varint,14,opt,name=HNFEIDJPLJK,proto3" json:"HNFEIDJPLJK,omitempty"`
	ONFHCNGNINF uint32 `protobuf:"varint,10,opt,name=ONFHCNGNINF,proto3" json:"ONFHCNGNINF,omitempty"`
}

func (x *FMBLGECBIBP) Reset() {
	*x = FMBLGECBIBP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FMBLGECBIBP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FMBLGECBIBP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FMBLGECBIBP) ProtoMessage() {}

func (x *FMBLGECBIBP) ProtoReflect() protoreflect.Message {
	mi := &file_FMBLGECBIBP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FMBLGECBIBP.ProtoReflect.Descriptor instead.
func (*FMBLGECBIBP) Descriptor() ([]byte, []int) {
	return file_FMBLGECBIBP_proto_rawDescGZIP(), []int{0}
}

func (x *FMBLGECBIBP) GetKEJHMBALBGN() uint32 {
	if x != nil {
		return x.KEJHMBALBGN
	}
	return 0
}

func (x *FMBLGECBIBP) GetOMKOINCNCCB() uint32 {
	if x != nil {
		return x.OMKOINCNCCB
	}
	return 0
}

func (x *FMBLGECBIBP) GetHNFEIDJPLJK() bool {
	if x != nil {
		return x.HNFEIDJPLJK
	}
	return false
}

func (x *FMBLGECBIBP) GetONFHCNGNINF() uint32 {
	if x != nil {
		return x.ONFHCNGNINF
	}
	return 0
}

var File_FMBLGECBIBP_proto protoreflect.FileDescriptor

var file_FMBLGECBIBP_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x4d, 0x42, 0x4c, 0x47, 0x45, 0x43, 0x42, 0x49, 0x42, 0x50, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x0b, 0x46, 0x4d, 0x42, 0x4c, 0x47, 0x45, 0x43, 0x42,
	0x49, 0x42, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x45, 0x4a, 0x48, 0x4d, 0x42, 0x41, 0x4c, 0x42,
	0x47, 0x4e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x45, 0x4a, 0x48, 0x4d, 0x42,
	0x41, 0x4c, 0x42, 0x47, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4d, 0x4b, 0x4f, 0x49, 0x4e, 0x43,
	0x4e, 0x43, 0x43, 0x42, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4d, 0x4b, 0x4f,
	0x49, 0x4e, 0x43, 0x4e, 0x43, 0x43, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4e, 0x46, 0x45, 0x49,
	0x44, 0x4a, 0x50, 0x4c, 0x4a, 0x4b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x48, 0x4e,
	0x46, 0x45, 0x49, 0x44, 0x4a, 0x50, 0x4c, 0x4a, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4e, 0x46,
	0x48, 0x43, 0x4e, 0x47, 0x4e, 0x49, 0x4e, 0x46, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4f, 0x4e, 0x46, 0x48, 0x43, 0x4e, 0x47, 0x4e, 0x49, 0x4e, 0x46, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_FMBLGECBIBP_proto_rawDescOnce sync.Once
	file_FMBLGECBIBP_proto_rawDescData = file_FMBLGECBIBP_proto_rawDesc
)

func file_FMBLGECBIBP_proto_rawDescGZIP() []byte {
	file_FMBLGECBIBP_proto_rawDescOnce.Do(func() {
		file_FMBLGECBIBP_proto_rawDescData = protoimpl.X.CompressGZIP(file_FMBLGECBIBP_proto_rawDescData)
	})
	return file_FMBLGECBIBP_proto_rawDescData
}

var file_FMBLGECBIBP_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FMBLGECBIBP_proto_goTypes = []interface{}{
	(*FMBLGECBIBP)(nil), // 0: FMBLGECBIBP
}
var file_FMBLGECBIBP_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FMBLGECBIBP_proto_init() }
func file_FMBLGECBIBP_proto_init() {
	if File_FMBLGECBIBP_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FMBLGECBIBP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FMBLGECBIBP); i {
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
			RawDescriptor: file_FMBLGECBIBP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FMBLGECBIBP_proto_goTypes,
		DependencyIndexes: file_FMBLGECBIBP_proto_depIdxs,
		MessageInfos:      file_FMBLGECBIBP_proto_msgTypes,
	}.Build()
	File_FMBLGECBIBP_proto = out.File
	file_FMBLGECBIBP_proto_rawDesc = nil
	file_FMBLGECBIBP_proto_goTypes = nil
	file_FMBLGECBIBP_proto_depIdxs = nil
}
