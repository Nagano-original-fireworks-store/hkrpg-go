// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetDrinkMakerDataScRsp.proto

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

type GetDrinkMakerDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exp         uint32             `protobuf:"varint,14,opt,name=exp,proto3" json:"exp,omitempty"`
	EDBBPELBEAM uint32             `protobuf:"varint,4,opt,name=EDBBPELBEAM,proto3" json:"EDBBPELBEAM,omitempty"`
	JNAPLPKLJEI *EICPBAEMNEC       `protobuf:"bytes,2,opt,name=JNAPLPKLJEI,proto3" json:"JNAPLPKLJEI,omitempty"`
	ACIGPIMKECH []uint32           `protobuf:"varint,5,rep,packed,name=ACIGPIMKECH,proto3" json:"ACIGPIMKECH,omitempty"`
	Retcode     uint32             `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	LDENIDLLCPN []*DrinkMakerGuest `protobuf:"bytes,8,rep,name=LDENIDLLCPN,proto3" json:"LDENIDLLCPN,omitempty"`
	IJFCPIKDAEK uint32             `protobuf:"varint,13,opt,name=IJFCPIKDAEK,proto3" json:"IJFCPIKDAEK,omitempty"`
	OMNPCIOLFBP uint32             `protobuf:"varint,3,opt,name=OMNPCIOLFBP,proto3" json:"OMNPCIOLFBP,omitempty"`
	ONMCCOPHDEM uint32             `protobuf:"varint,11,opt,name=ONMCCOPHDEM,proto3" json:"ONMCCOPHDEM,omitempty"`
	Level       uint32             `protobuf:"varint,10,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *GetDrinkMakerDataScRsp) Reset() {
	*x = GetDrinkMakerDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetDrinkMakerDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrinkMakerDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrinkMakerDataScRsp) ProtoMessage() {}

func (x *GetDrinkMakerDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetDrinkMakerDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrinkMakerDataScRsp.ProtoReflect.Descriptor instead.
func (*GetDrinkMakerDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetDrinkMakerDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetDrinkMakerDataScRsp) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *GetDrinkMakerDataScRsp) GetEDBBPELBEAM() uint32 {
	if x != nil {
		return x.EDBBPELBEAM
	}
	return 0
}

func (x *GetDrinkMakerDataScRsp) GetJNAPLPKLJEI() *EICPBAEMNEC {
	if x != nil {
		return x.JNAPLPKLJEI
	}
	return nil
}

func (x *GetDrinkMakerDataScRsp) GetACIGPIMKECH() []uint32 {
	if x != nil {
		return x.ACIGPIMKECH
	}
	return nil
}

func (x *GetDrinkMakerDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetDrinkMakerDataScRsp) GetLDENIDLLCPN() []*DrinkMakerGuest {
	if x != nil {
		return x.LDENIDLLCPN
	}
	return nil
}

func (x *GetDrinkMakerDataScRsp) GetIJFCPIKDAEK() uint32 {
	if x != nil {
		return x.IJFCPIKDAEK
	}
	return 0
}

func (x *GetDrinkMakerDataScRsp) GetOMNPCIOLFBP() uint32 {
	if x != nil {
		return x.OMNPCIOLFBP
	}
	return 0
}

func (x *GetDrinkMakerDataScRsp) GetONMCCOPHDEM() uint32 {
	if x != nil {
		return x.ONMCCOPHDEM
	}
	return 0
}

func (x *GetDrinkMakerDataScRsp) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_GetDrinkMakerDataScRsp_proto protoreflect.FileDescriptor

var file_GetDrinkMakerDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x6e, 0x6b, 0x4d, 0x61, 0x6b, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x44, 0x72, 0x69, 0x6e, 0x6b, 0x4d, 0x61, 0x6b, 0x65, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x49, 0x43, 0x50, 0x42, 0x41, 0x45, 0x4d, 0x4e,
	0x45, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x02, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x44, 0x72, 0x69, 0x6e, 0x6b, 0x4d, 0x61, 0x6b, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x44, 0x42, 0x42, 0x50, 0x45, 0x4c,
	0x42, 0x45, 0x41, 0x4d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x44, 0x42, 0x42,
	0x50, 0x45, 0x4c, 0x42, 0x45, 0x41, 0x4d, 0x12, 0x2e, 0x0a, 0x0b, 0x4a, 0x4e, 0x41, 0x50, 0x4c,
	0x50, 0x4b, 0x4c, 0x4a, 0x45, 0x49, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45,
	0x49, 0x43, 0x50, 0x42, 0x41, 0x45, 0x4d, 0x4e, 0x45, 0x43, 0x52, 0x0b, 0x4a, 0x4e, 0x41, 0x50,
	0x4c, 0x50, 0x4b, 0x4c, 0x4a, 0x45, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x43, 0x49, 0x47, 0x50,
	0x49, 0x4d, 0x4b, 0x45, 0x43, 0x48, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x43,
	0x49, 0x47, 0x50, 0x49, 0x4d, 0x4b, 0x45, 0x43, 0x48, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x4c, 0x44, 0x45, 0x4e, 0x49, 0x44, 0x4c, 0x4c, 0x43,
	0x50, 0x4e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x44, 0x72, 0x69, 0x6e, 0x6b,
	0x4d, 0x61, 0x6b, 0x65, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x4c, 0x44, 0x45, 0x4e,
	0x49, 0x44, 0x4c, 0x4c, 0x43, 0x50, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4a, 0x46, 0x43, 0x50,
	0x49, 0x4b, 0x44, 0x41, 0x45, 0x4b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4a,
	0x46, 0x43, 0x50, 0x49, 0x4b, 0x44, 0x41, 0x45, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4d, 0x4e,
	0x50, 0x43, 0x49, 0x4f, 0x4c, 0x46, 0x42, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4f, 0x4d, 0x4e, 0x50, 0x43, 0x49, 0x4f, 0x4c, 0x46, 0x42, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x4f,
	0x4e, 0x4d, 0x43, 0x43, 0x4f, 0x50, 0x48, 0x44, 0x45, 0x4d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4f, 0x4e, 0x4d, 0x43, 0x43, 0x4f, 0x50, 0x48, 0x44, 0x45, 0x4d, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetDrinkMakerDataScRsp_proto_rawDescOnce sync.Once
	file_GetDrinkMakerDataScRsp_proto_rawDescData = file_GetDrinkMakerDataScRsp_proto_rawDesc
)

func file_GetDrinkMakerDataScRsp_proto_rawDescGZIP() []byte {
	file_GetDrinkMakerDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetDrinkMakerDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetDrinkMakerDataScRsp_proto_rawDescData)
	})
	return file_GetDrinkMakerDataScRsp_proto_rawDescData
}

var file_GetDrinkMakerDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetDrinkMakerDataScRsp_proto_goTypes = []interface{}{
	(*GetDrinkMakerDataScRsp)(nil), // 0: GetDrinkMakerDataScRsp
	(*EICPBAEMNEC)(nil),            // 1: EICPBAEMNEC
	(*DrinkMakerGuest)(nil),        // 2: DrinkMakerGuest
}
var file_GetDrinkMakerDataScRsp_proto_depIdxs = []int32{
	1, // 0: GetDrinkMakerDataScRsp.JNAPLPKLJEI:type_name -> EICPBAEMNEC
	2, // 1: GetDrinkMakerDataScRsp.LDENIDLLCPN:type_name -> DrinkMakerGuest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetDrinkMakerDataScRsp_proto_init() }
func file_GetDrinkMakerDataScRsp_proto_init() {
	if File_GetDrinkMakerDataScRsp_proto != nil {
		return
	}
	file_DrinkMakerGuest_proto_init()
	file_EICPBAEMNEC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetDrinkMakerDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrinkMakerDataScRsp); i {
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
			RawDescriptor: file_GetDrinkMakerDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetDrinkMakerDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetDrinkMakerDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetDrinkMakerDataScRsp_proto_msgTypes,
	}.Build()
	File_GetDrinkMakerDataScRsp_proto = out.File
	file_GetDrinkMakerDataScRsp_proto_rawDesc = nil
	file_GetDrinkMakerDataScRsp_proto_goTypes = nil
	file_GetDrinkMakerDataScRsp_proto_depIdxs = nil
}
