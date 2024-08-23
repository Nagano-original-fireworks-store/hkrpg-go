// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CIBFGBHDPAM.proto

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

type CIBFGBHDPAM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId     uint32                 `protobuf:"varint,6,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	FAGJDONPNNE uint32                 `protobuf:"varint,10,opt,name=FAGJDONPNNE,proto3" json:"FAGJDONPNNE,omitempty"`
	EOCPNIDFFDD uint32                 `protobuf:"varint,7,opt,name=EOCPNIDFFDD,proto3" json:"EOCPNIDFFDD,omitempty"`
	CGHBPPHAMCE uint32                 `protobuf:"varint,14,opt,name=CGHBPPHAMCE,proto3" json:"CGHBPPHAMCE,omitempty"`
	State       MuseumRandomEventState `protobuf:"varint,1,opt,name=state,proto3,enum=MuseumRandomEventState" json:"state,omitempty"`
	NBMOOEFBEFD []uint32               `protobuf:"varint,4,rep,packed,name=NBMOOEFBEFD,proto3" json:"NBMOOEFBEFD,omitempty"`
	FMCKENFKHFM uint32                 `protobuf:"varint,13,opt,name=FMCKENFKHFM,proto3" json:"FMCKENFKHFM,omitempty"`
}

func (x *CIBFGBHDPAM) Reset() {
	*x = CIBFGBHDPAM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIBFGBHDPAM_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIBFGBHDPAM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIBFGBHDPAM) ProtoMessage() {}

func (x *CIBFGBHDPAM) ProtoReflect() protoreflect.Message {
	mi := &file_CIBFGBHDPAM_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIBFGBHDPAM.ProtoReflect.Descriptor instead.
func (*CIBFGBHDPAM) Descriptor() ([]byte, []int) {
	return file_CIBFGBHDPAM_proto_rawDescGZIP(), []int{0}
}

func (x *CIBFGBHDPAM) GetEventId() uint32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *CIBFGBHDPAM) GetFAGJDONPNNE() uint32 {
	if x != nil {
		return x.FAGJDONPNNE
	}
	return 0
}

func (x *CIBFGBHDPAM) GetEOCPNIDFFDD() uint32 {
	if x != nil {
		return x.EOCPNIDFFDD
	}
	return 0
}

func (x *CIBFGBHDPAM) GetCGHBPPHAMCE() uint32 {
	if x != nil {
		return x.CGHBPPHAMCE
	}
	return 0
}

func (x *CIBFGBHDPAM) GetState() MuseumRandomEventState {
	if x != nil {
		return x.State
	}
	return MuseumRandomEventState_MUSEUM_RANDOM_EVENT_STATE_NONE
}

func (x *CIBFGBHDPAM) GetNBMOOEFBEFD() []uint32 {
	if x != nil {
		return x.NBMOOEFBEFD
	}
	return nil
}

func (x *CIBFGBHDPAM) GetFMCKENFKHFM() uint32 {
	if x != nil {
		return x.FMCKENFKHFM
	}
	return 0
}

var File_CIBFGBHDPAM_proto protoreflect.FileDescriptor

var file_CIBFGBHDPAM_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x49, 0x42, 0x46, 0x47, 0x42, 0x48, 0x44, 0x50, 0x41, 0x4d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x4d, 0x75, 0x73, 0x65, 0x75, 0x6d, 0x52, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x81, 0x02, 0x0a, 0x0b, 0x43, 0x49, 0x42, 0x46, 0x47, 0x42, 0x48, 0x44, 0x50, 0x41,
	0x4d, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x46, 0x41, 0x47, 0x4a, 0x44, 0x4f, 0x4e, 0x50, 0x4e, 0x4e, 0x45, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x46, 0x41, 0x47, 0x4a, 0x44, 0x4f, 0x4e, 0x50, 0x4e, 0x4e, 0x45, 0x12, 0x20,
	0x0a, 0x0b, 0x45, 0x4f, 0x43, 0x50, 0x4e, 0x49, 0x44, 0x46, 0x46, 0x44, 0x44, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x4f, 0x43, 0x50, 0x4e, 0x49, 0x44, 0x46, 0x46, 0x44, 0x44,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x47, 0x48, 0x42, 0x50, 0x50, 0x48, 0x41, 0x4d, 0x43, 0x45, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x47, 0x48, 0x42, 0x50, 0x50, 0x48, 0x41, 0x4d,
	0x43, 0x45, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x4d, 0x75, 0x73, 0x65, 0x75, 0x6d, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x42, 0x4d, 0x4f, 0x4f, 0x45, 0x46, 0x42, 0x45, 0x46, 0x44,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x42, 0x4d, 0x4f, 0x4f, 0x45, 0x46, 0x42,
	0x45, 0x46, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4d, 0x43, 0x4b, 0x45, 0x4e, 0x46, 0x4b, 0x48,
	0x46, 0x4d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x4d, 0x43, 0x4b, 0x45, 0x4e,
	0x46, 0x4b, 0x48, 0x46, 0x4d, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CIBFGBHDPAM_proto_rawDescOnce sync.Once
	file_CIBFGBHDPAM_proto_rawDescData = file_CIBFGBHDPAM_proto_rawDesc
)

func file_CIBFGBHDPAM_proto_rawDescGZIP() []byte {
	file_CIBFGBHDPAM_proto_rawDescOnce.Do(func() {
		file_CIBFGBHDPAM_proto_rawDescData = protoimpl.X.CompressGZIP(file_CIBFGBHDPAM_proto_rawDescData)
	})
	return file_CIBFGBHDPAM_proto_rawDescData
}

var file_CIBFGBHDPAM_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CIBFGBHDPAM_proto_goTypes = []interface{}{
	(*CIBFGBHDPAM)(nil),         // 0: CIBFGBHDPAM
	(MuseumRandomEventState)(0), // 1: MuseumRandomEventState
}
var file_CIBFGBHDPAM_proto_depIdxs = []int32{
	1, // 0: CIBFGBHDPAM.state:type_name -> MuseumRandomEventState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CIBFGBHDPAM_proto_init() }
func file_CIBFGBHDPAM_proto_init() {
	if File_CIBFGBHDPAM_proto != nil {
		return
	}
	file_MuseumRandomEventState_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CIBFGBHDPAM_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIBFGBHDPAM); i {
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
			RawDescriptor: file_CIBFGBHDPAM_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CIBFGBHDPAM_proto_goTypes,
		DependencyIndexes: file_CIBFGBHDPAM_proto_depIdxs,
		MessageInfos:      file_CIBFGBHDPAM_proto_msgTypes,
	}.Build()
	File_CIBFGBHDPAM_proto = out.File
	file_CIBFGBHDPAM_proto_rawDesc = nil
	file_CIBFGBHDPAM_proto_goTypes = nil
	file_CIBFGBHDPAM_proto_depIdxs = nil
}
