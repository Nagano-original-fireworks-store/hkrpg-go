// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerRecordInfo.proto

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

type PlayerRecordInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OOIOMMKOIAE    uint32                `protobuf:"varint,15,opt,name=OOIOMMKOIAE,proto3" json:"OOIOMMKOIAE,omitempty"`
	DKPLJPHGDMB    uint32                `protobuf:"varint,7,opt,name=DKPLJPHGDMB,proto3" json:"DKPLJPHGDMB,omitempty"`
	HNICFKBEKJO    uint32                `protobuf:"varint,14,opt,name=HNICFKBEKJO,proto3" json:"HNICFKBEKJO,omitempty"`
	INMEAGINKPN    uint32                `protobuf:"varint,8,opt,name=INMEAGINKPN,proto3" json:"INMEAGINKPN,omitempty"`
	CPIKKMDOKHA    uint32                `protobuf:"varint,3,opt,name=CPIKKMDOKHA,proto3" json:"CPIKKMDOKHA,omitempty"`
	MGDIMDLPOGA    uint32                `protobuf:"varint,4,opt,name=MGDIMDLPOGA,proto3" json:"MGDIMDLPOGA,omitempty"`
	BIOJHIBFELK    uint32                `protobuf:"varint,10,opt,name=BIOJHIBFELK,proto3" json:"BIOJHIBFELK,omitempty"`
	CollectionInfo *PlayerCollectionInfo `protobuf:"bytes,2,opt,name=collection_info,json=collectionInfo,proto3" json:"collection_info,omitempty"`
	MFMOAJONBBF    uint32                `protobuf:"varint,12,opt,name=MFMOAJONBBF,proto3" json:"MFMOAJONBBF,omitempty"`
}

func (x *PlayerRecordInfo) Reset() {
	*x = PlayerRecordInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerRecordInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRecordInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRecordInfo) ProtoMessage() {}

func (x *PlayerRecordInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerRecordInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRecordInfo.ProtoReflect.Descriptor instead.
func (*PlayerRecordInfo) Descriptor() ([]byte, []int) {
	return file_PlayerRecordInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerRecordInfo) GetOOIOMMKOIAE() uint32 {
	if x != nil {
		return x.OOIOMMKOIAE
	}
	return 0
}

func (x *PlayerRecordInfo) GetDKPLJPHGDMB() uint32 {
	if x != nil {
		return x.DKPLJPHGDMB
	}
	return 0
}

func (x *PlayerRecordInfo) GetHNICFKBEKJO() uint32 {
	if x != nil {
		return x.HNICFKBEKJO
	}
	return 0
}

func (x *PlayerRecordInfo) GetINMEAGINKPN() uint32 {
	if x != nil {
		return x.INMEAGINKPN
	}
	return 0
}

func (x *PlayerRecordInfo) GetCPIKKMDOKHA() uint32 {
	if x != nil {
		return x.CPIKKMDOKHA
	}
	return 0
}

func (x *PlayerRecordInfo) GetMGDIMDLPOGA() uint32 {
	if x != nil {
		return x.MGDIMDLPOGA
	}
	return 0
}

func (x *PlayerRecordInfo) GetBIOJHIBFELK() uint32 {
	if x != nil {
		return x.BIOJHIBFELK
	}
	return 0
}

func (x *PlayerRecordInfo) GetCollectionInfo() *PlayerCollectionInfo {
	if x != nil {
		return x.CollectionInfo
	}
	return nil
}

func (x *PlayerRecordInfo) GetMFMOAJONBBF() uint32 {
	if x != nil {
		return x.MFMOAJONBBF
	}
	return 0
}

var File_PlayerRecordInfo_proto protoreflect.FileDescriptor

var file_PlayerRecordInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x02, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4f, 0x49,
	0x4f, 0x4d, 0x4d, 0x4b, 0x4f, 0x49, 0x41, 0x45, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4f, 0x4f, 0x49, 0x4f, 0x4d, 0x4d, 0x4b, 0x4f, 0x49, 0x41, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x4b, 0x50, 0x4c, 0x4a, 0x50, 0x48, 0x47, 0x44, 0x4d, 0x42, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x44, 0x4b, 0x50, 0x4c, 0x4a, 0x50, 0x48, 0x47, 0x44, 0x4d, 0x42, 0x12, 0x20, 0x0a,
	0x0b, 0x48, 0x4e, 0x49, 0x43, 0x46, 0x4b, 0x42, 0x45, 0x4b, 0x4a, 0x4f, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4e, 0x49, 0x43, 0x46, 0x4b, 0x42, 0x45, 0x4b, 0x4a, 0x4f, 0x12,
	0x20, 0x0a, 0x0b, 0x49, 0x4e, 0x4d, 0x45, 0x41, 0x47, 0x49, 0x4e, 0x4b, 0x50, 0x4e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4e, 0x4d, 0x45, 0x41, 0x47, 0x49, 0x4e, 0x4b, 0x50,
	0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x50, 0x49, 0x4b, 0x4b, 0x4d, 0x44, 0x4f, 0x4b, 0x48, 0x41,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x50, 0x49, 0x4b, 0x4b, 0x4d, 0x44, 0x4f,
	0x4b, 0x48, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x47, 0x44, 0x49, 0x4d, 0x44, 0x4c, 0x50, 0x4f,
	0x47, 0x41, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x47, 0x44, 0x49, 0x4d, 0x44,
	0x4c, 0x50, 0x4f, 0x47, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x49, 0x4f, 0x4a, 0x48, 0x49, 0x42,
	0x46, 0x45, 0x4c, 0x4b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x49, 0x4f, 0x4a,
	0x48, 0x49, 0x42, 0x46, 0x45, 0x4c, 0x4b, 0x12, 0x3e, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x46, 0x4d, 0x4f, 0x41,
	0x4a, 0x4f, 0x4e, 0x42, 0x42, 0x46, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x46,
	0x4d, 0x4f, 0x41, 0x4a, 0x4f, 0x4e, 0x42, 0x42, 0x46, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_PlayerRecordInfo_proto_rawDescOnce sync.Once
	file_PlayerRecordInfo_proto_rawDescData = file_PlayerRecordInfo_proto_rawDesc
)

func file_PlayerRecordInfo_proto_rawDescGZIP() []byte {
	file_PlayerRecordInfo_proto_rawDescOnce.Do(func() {
		file_PlayerRecordInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerRecordInfo_proto_rawDescData)
	})
	return file_PlayerRecordInfo_proto_rawDescData
}

var file_PlayerRecordInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerRecordInfo_proto_goTypes = []interface{}{
	(*PlayerRecordInfo)(nil),     // 0: PlayerRecordInfo
	(*PlayerCollectionInfo)(nil), // 1: PlayerCollectionInfo
}
var file_PlayerRecordInfo_proto_depIdxs = []int32{
	1, // 0: PlayerRecordInfo.collection_info:type_name -> PlayerCollectionInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerRecordInfo_proto_init() }
func file_PlayerRecordInfo_proto_init() {
	if File_PlayerRecordInfo_proto != nil {
		return
	}
	file_PlayerCollectionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerRecordInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRecordInfo); i {
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
			RawDescriptor: file_PlayerRecordInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerRecordInfo_proto_goTypes,
		DependencyIndexes: file_PlayerRecordInfo_proto_depIdxs,
		MessageInfos:      file_PlayerRecordInfo_proto_msgTypes,
	}.Build()
	File_PlayerRecordInfo_proto = out.File
	file_PlayerRecordInfo_proto_rawDesc = nil
	file_PlayerRecordInfo_proto_goTypes = nil
	file_PlayerRecordInfo_proto_depIdxs = nil
}
