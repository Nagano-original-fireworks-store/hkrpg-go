// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: JLLPBIJLJIH.proto

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

type JLLPBIJLJIH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FFOLOOAGDEI []*FEMNHLAPBOG    `protobuf:"bytes,11,rep,name=FFOLOOAGDEI,proto3" json:"FFOLOOAGDEI,omitempty"`
	PCKBGKDCHAB map[uint32]uint32 `protobuf:"bytes,1,rep,name=PCKBGKDCHAB,proto3" json:"PCKBGKDCHAB,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MIMOBNLKAPN map[uint32]uint32 `protobuf:"bytes,7,rep,name=MIMOBNLKAPN,proto3" json:"MIMOBNLKAPN,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ODMDDNEJFBI []*ILGPCLDJFOE    `protobuf:"bytes,2,rep,name=ODMDDNEJFBI,proto3" json:"ODMDDNEJFBI,omitempty"`
}

func (x *JLLPBIJLJIH) Reset() {
	*x = JLLPBIJLJIH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JLLPBIJLJIH_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JLLPBIJLJIH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JLLPBIJLJIH) ProtoMessage() {}

func (x *JLLPBIJLJIH) ProtoReflect() protoreflect.Message {
	mi := &file_JLLPBIJLJIH_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JLLPBIJLJIH.ProtoReflect.Descriptor instead.
func (*JLLPBIJLJIH) Descriptor() ([]byte, []int) {
	return file_JLLPBIJLJIH_proto_rawDescGZIP(), []int{0}
}

func (x *JLLPBIJLJIH) GetFFOLOOAGDEI() []*FEMNHLAPBOG {
	if x != nil {
		return x.FFOLOOAGDEI
	}
	return nil
}

func (x *JLLPBIJLJIH) GetPCKBGKDCHAB() map[uint32]uint32 {
	if x != nil {
		return x.PCKBGKDCHAB
	}
	return nil
}

func (x *JLLPBIJLJIH) GetMIMOBNLKAPN() map[uint32]uint32 {
	if x != nil {
		return x.MIMOBNLKAPN
	}
	return nil
}

func (x *JLLPBIJLJIH) GetODMDDNEJFBI() []*ILGPCLDJFOE {
	if x != nil {
		return x.ODMDDNEJFBI
	}
	return nil
}

var File_JLLPBIJLJIH_proto protoreflect.FileDescriptor

var file_JLLPBIJLJIH_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x4c, 0x4c, 0x50, 0x42, 0x49, 0x4a, 0x4c, 0x4a, 0x49, 0x48, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x45, 0x4d, 0x4e, 0x48, 0x4c, 0x41, 0x50, 0x42, 0x4f, 0x47,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x4c, 0x47, 0x50, 0x43, 0x4c, 0x44, 0x4a,
	0x46, 0x4f, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x02, 0x0a, 0x0b, 0x4a, 0x4c,
	0x4c, 0x50, 0x42, 0x49, 0x4a, 0x4c, 0x4a, 0x49, 0x48, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x46, 0x4f,
	0x4c, 0x4f, 0x4f, 0x41, 0x47, 0x44, 0x45, 0x49, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x46, 0x45, 0x4d, 0x4e, 0x48, 0x4c, 0x41, 0x50, 0x42, 0x4f, 0x47, 0x52, 0x0b, 0x46, 0x46,
	0x4f, 0x4c, 0x4f, 0x4f, 0x41, 0x47, 0x44, 0x45, 0x49, 0x12, 0x3f, 0x0a, 0x0b, 0x50, 0x43, 0x4b,
	0x42, 0x47, 0x4b, 0x44, 0x43, 0x48, 0x41, 0x42, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x4a, 0x4c, 0x4c, 0x50, 0x42, 0x49, 0x4a, 0x4c, 0x4a, 0x49, 0x48, 0x2e, 0x50, 0x43, 0x4b,
	0x42, 0x47, 0x4b, 0x44, 0x43, 0x48, 0x41, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x50,
	0x43, 0x4b, 0x42, 0x47, 0x4b, 0x44, 0x43, 0x48, 0x41, 0x42, 0x12, 0x3f, 0x0a, 0x0b, 0x4d, 0x49,
	0x4d, 0x4f, 0x42, 0x4e, 0x4c, 0x4b, 0x41, 0x50, 0x4e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x4a, 0x4c, 0x4c, 0x50, 0x42, 0x49, 0x4a, 0x4c, 0x4a, 0x49, 0x48, 0x2e, 0x4d, 0x49,
	0x4d, 0x4f, 0x42, 0x4e, 0x4c, 0x4b, 0x41, 0x50, 0x4e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x4d, 0x49, 0x4d, 0x4f, 0x42, 0x4e, 0x4c, 0x4b, 0x41, 0x50, 0x4e, 0x12, 0x2e, 0x0a, 0x0b, 0x4f,
	0x44, 0x4d, 0x44, 0x44, 0x4e, 0x45, 0x4a, 0x46, 0x42, 0x49, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x49, 0x4c, 0x47, 0x50, 0x43, 0x4c, 0x44, 0x4a, 0x46, 0x4f, 0x45, 0x52, 0x0b,
	0x4f, 0x44, 0x4d, 0x44, 0x44, 0x4e, 0x45, 0x4a, 0x46, 0x42, 0x49, 0x1a, 0x3e, 0x0a, 0x10, 0x50,
	0x43, 0x4b, 0x42, 0x47, 0x4b, 0x44, 0x43, 0x48, 0x41, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x4d,
	0x49, 0x4d, 0x4f, 0x42, 0x4e, 0x4c, 0x4b, 0x41, 0x50, 0x4e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_JLLPBIJLJIH_proto_rawDescOnce sync.Once
	file_JLLPBIJLJIH_proto_rawDescData = file_JLLPBIJLJIH_proto_rawDesc
)

func file_JLLPBIJLJIH_proto_rawDescGZIP() []byte {
	file_JLLPBIJLJIH_proto_rawDescOnce.Do(func() {
		file_JLLPBIJLJIH_proto_rawDescData = protoimpl.X.CompressGZIP(file_JLLPBIJLJIH_proto_rawDescData)
	})
	return file_JLLPBIJLJIH_proto_rawDescData
}

var file_JLLPBIJLJIH_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_JLLPBIJLJIH_proto_goTypes = []interface{}{
	(*JLLPBIJLJIH)(nil), // 0: JLLPBIJLJIH
	nil,                 // 1: JLLPBIJLJIH.PCKBGKDCHABEntry
	nil,                 // 2: JLLPBIJLJIH.MIMOBNLKAPNEntry
	(*FEMNHLAPBOG)(nil), // 3: FEMNHLAPBOG
	(*ILGPCLDJFOE)(nil), // 4: ILGPCLDJFOE
}
var file_JLLPBIJLJIH_proto_depIdxs = []int32{
	3, // 0: JLLPBIJLJIH.FFOLOOAGDEI:type_name -> FEMNHLAPBOG
	1, // 1: JLLPBIJLJIH.PCKBGKDCHAB:type_name -> JLLPBIJLJIH.PCKBGKDCHABEntry
	2, // 2: JLLPBIJLJIH.MIMOBNLKAPN:type_name -> JLLPBIJLJIH.MIMOBNLKAPNEntry
	4, // 3: JLLPBIJLJIH.ODMDDNEJFBI:type_name -> ILGPCLDJFOE
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_JLLPBIJLJIH_proto_init() }
func file_JLLPBIJLJIH_proto_init() {
	if File_JLLPBIJLJIH_proto != nil {
		return
	}
	file_FEMNHLAPBOG_proto_init()
	file_ILGPCLDJFOE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_JLLPBIJLJIH_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JLLPBIJLJIH); i {
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
			RawDescriptor: file_JLLPBIJLJIH_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JLLPBIJLJIH_proto_goTypes,
		DependencyIndexes: file_JLLPBIJLJIH_proto_depIdxs,
		MessageInfos:      file_JLLPBIJLJIH_proto_msgTypes,
	}.Build()
	File_JLLPBIJLJIH_proto = out.File
	file_JLLPBIJLJIH_proto_rawDesc = nil
	file_JLLPBIJLJIH_proto_goTypes = nil
	file_JLLPBIJLJIH_proto_depIdxs = nil
}
