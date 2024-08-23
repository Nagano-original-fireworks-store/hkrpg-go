// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FantasticActivityData.proto

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

type FantasticActivityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ONEDDAAMDCN map[uint32]uint32       `protobuf:"bytes,6,rep,name=ONEDDAAMDCN,proto3" json:"ONEDDAAMDCN,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	JAMODBBMMIJ uint32                  `protobuf:"varint,2,opt,name=JAMODBBMMIJ,proto3" json:"JAMODBBMMIJ,omitempty"`
	DFLLHCPLAGJ []uint32                `protobuf:"varint,8,rep,packed,name=DFLLHCPLAGJ,proto3" json:"DFLLHCPLAGJ,omitempty"`
	FDDKCHDGAGE []uint32                `protobuf:"varint,5,rep,packed,name=FDDKCHDGAGE,proto3" json:"FDDKCHDGAGE,omitempty"`
	JDBCGDGENOC []uint32                `protobuf:"varint,14,rep,packed,name=JDBCGDGENOC,proto3" json:"JDBCGDGENOC,omitempty"`
	NELLGGMKIBP map[uint32]*CNBIFBFGMGD `protobuf:"bytes,4,rep,name=NELLGGMKIBP,proto3" json:"NELLGGMKIBP,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FHDANACBFMJ []uint32                `protobuf:"varint,13,rep,packed,name=FHDANACBFMJ,proto3" json:"FHDANACBFMJ,omitempty"`
}

func (x *FantasticActivityData) Reset() {
	*x = FantasticActivityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FantasticActivityData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FantasticActivityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FantasticActivityData) ProtoMessage() {}

func (x *FantasticActivityData) ProtoReflect() protoreflect.Message {
	mi := &file_FantasticActivityData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FantasticActivityData.ProtoReflect.Descriptor instead.
func (*FantasticActivityData) Descriptor() ([]byte, []int) {
	return file_FantasticActivityData_proto_rawDescGZIP(), []int{0}
}

func (x *FantasticActivityData) GetONEDDAAMDCN() map[uint32]uint32 {
	if x != nil {
		return x.ONEDDAAMDCN
	}
	return nil
}

func (x *FantasticActivityData) GetJAMODBBMMIJ() uint32 {
	if x != nil {
		return x.JAMODBBMMIJ
	}
	return 0
}

func (x *FantasticActivityData) GetDFLLHCPLAGJ() []uint32 {
	if x != nil {
		return x.DFLLHCPLAGJ
	}
	return nil
}

func (x *FantasticActivityData) GetFDDKCHDGAGE() []uint32 {
	if x != nil {
		return x.FDDKCHDGAGE
	}
	return nil
}

func (x *FantasticActivityData) GetJDBCGDGENOC() []uint32 {
	if x != nil {
		return x.JDBCGDGENOC
	}
	return nil
}

func (x *FantasticActivityData) GetNELLGGMKIBP() map[uint32]*CNBIFBFGMGD {
	if x != nil {
		return x.NELLGGMKIBP
	}
	return nil
}

func (x *FantasticActivityData) GetFHDANACBFMJ() []uint32 {
	if x != nil {
		return x.FHDANACBFMJ
	}
	return nil
}

var File_FantasticActivityData_proto protoreflect.FileDescriptor

var file_FantasticActivityData_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43,
	0x4e, 0x42, 0x49, 0x46, 0x42, 0x46, 0x47, 0x4d, 0x47, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe5, 0x03, 0x0a, 0x15, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x0b, 0x4f, 0x4e,
	0x45, 0x44, 0x44, 0x41, 0x41, 0x4d, 0x44, 0x43, 0x4e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x4e, 0x45, 0x44, 0x44, 0x41, 0x41, 0x4d,
	0x44, 0x43, 0x4e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x4f, 0x4e, 0x45, 0x44, 0x44, 0x41,
	0x41, 0x4d, 0x44, 0x43, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x41, 0x4d, 0x4f, 0x44, 0x42, 0x42,
	0x4d, 0x4d, 0x49, 0x4a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x41, 0x4d, 0x4f,
	0x44, 0x42, 0x42, 0x4d, 0x4d, 0x49, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x46, 0x4c, 0x4c, 0x48,
	0x43, 0x50, 0x4c, 0x41, 0x47, 0x4a, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x46,
	0x4c, 0x4c, 0x48, 0x43, 0x50, 0x4c, 0x41, 0x47, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x44, 0x44,
	0x4b, 0x43, 0x48, 0x44, 0x47, 0x41, 0x47, 0x45, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b,
	0x46, 0x44, 0x44, 0x4b, 0x43, 0x48, 0x44, 0x47, 0x41, 0x47, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x4a,
	0x44, 0x42, 0x43, 0x47, 0x44, 0x47, 0x45, 0x4e, 0x4f, 0x43, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0b, 0x4a, 0x44, 0x42, 0x43, 0x47, 0x44, 0x47, 0x45, 0x4e, 0x4f, 0x43, 0x12, 0x49, 0x0a,
	0x0b, 0x4e, 0x45, 0x4c, 0x4c, 0x47, 0x47, 0x4d, 0x4b, 0x49, 0x42, 0x50, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4e, 0x45, 0x4c, 0x4c, 0x47,
	0x47, 0x4d, 0x4b, 0x49, 0x42, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x4e, 0x45, 0x4c,
	0x4c, 0x47, 0x47, 0x4d, 0x4b, 0x49, 0x42, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x48, 0x44, 0x41,
	0x4e, 0x41, 0x43, 0x42, 0x46, 0x4d, 0x4a, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x46,
	0x48, 0x44, 0x41, 0x4e, 0x41, 0x43, 0x42, 0x46, 0x4d, 0x4a, 0x1a, 0x3e, 0x0a, 0x10, 0x4f, 0x4e,
	0x45, 0x44, 0x44, 0x41, 0x41, 0x4d, 0x44, 0x43, 0x4e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4c, 0x0a, 0x10, 0x4e, 0x45,
	0x4c, 0x4c, 0x47, 0x47, 0x4d, 0x4b, 0x49, 0x42, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x43, 0x4e, 0x42, 0x49, 0x46, 0x42, 0x46, 0x47, 0x4d, 0x47, 0x44, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FantasticActivityData_proto_rawDescOnce sync.Once
	file_FantasticActivityData_proto_rawDescData = file_FantasticActivityData_proto_rawDesc
)

func file_FantasticActivityData_proto_rawDescGZIP() []byte {
	file_FantasticActivityData_proto_rawDescOnce.Do(func() {
		file_FantasticActivityData_proto_rawDescData = protoimpl.X.CompressGZIP(file_FantasticActivityData_proto_rawDescData)
	})
	return file_FantasticActivityData_proto_rawDescData
}

var file_FantasticActivityData_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_FantasticActivityData_proto_goTypes = []interface{}{
	(*FantasticActivityData)(nil), // 0: FantasticActivityData
	nil,                           // 1: FantasticActivityData.ONEDDAAMDCNEntry
	nil,                           // 2: FantasticActivityData.NELLGGMKIBPEntry
	(*CNBIFBFGMGD)(nil),           // 3: CNBIFBFGMGD
}
var file_FantasticActivityData_proto_depIdxs = []int32{
	1, // 0: FantasticActivityData.ONEDDAAMDCN:type_name -> FantasticActivityData.ONEDDAAMDCNEntry
	2, // 1: FantasticActivityData.NELLGGMKIBP:type_name -> FantasticActivityData.NELLGGMKIBPEntry
	3, // 2: FantasticActivityData.NELLGGMKIBPEntry.value:type_name -> CNBIFBFGMGD
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_FantasticActivityData_proto_init() }
func file_FantasticActivityData_proto_init() {
	if File_FantasticActivityData_proto != nil {
		return
	}
	file_CNBIFBFGMGD_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FantasticActivityData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FantasticActivityData); i {
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
			RawDescriptor: file_FantasticActivityData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FantasticActivityData_proto_goTypes,
		DependencyIndexes: file_FantasticActivityData_proto_depIdxs,
		MessageInfos:      file_FantasticActivityData_proto_msgTypes,
	}.Build()
	File_FantasticActivityData_proto = out.File
	file_FantasticActivityData_proto_rawDesc = nil
	file_FantasticActivityData_proto_goTypes = nil
	file_FantasticActivityData_proto_depIdxs = nil
}
