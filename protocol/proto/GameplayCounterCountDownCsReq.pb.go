// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GameplayCounterCountDownCsReq.proto

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

type GameplayCounterCountDownCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BAOKAGNFNAB uint32 `protobuf:"varint,7,opt,name=BAOKAGNFNAB,proto3" json:"BAOKAGNFNAB,omitempty"`
	UsedTimes   uint32 `protobuf:"varint,11,opt,name=used_times,json=usedTimes,proto3" json:"used_times,omitempty"`
}

func (x *GameplayCounterCountDownCsReq) Reset() {
	*x = GameplayCounterCountDownCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameplayCounterCountDownCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameplayCounterCountDownCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameplayCounterCountDownCsReq) ProtoMessage() {}

func (x *GameplayCounterCountDownCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_GameplayCounterCountDownCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameplayCounterCountDownCsReq.ProtoReflect.Descriptor instead.
func (*GameplayCounterCountDownCsReq) Descriptor() ([]byte, []int) {
	return file_GameplayCounterCountDownCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *GameplayCounterCountDownCsReq) GetBAOKAGNFNAB() uint32 {
	if x != nil {
		return x.BAOKAGNFNAB
	}
	return 0
}

func (x *GameplayCounterCountDownCsReq) GetUsedTimes() uint32 {
	if x != nil {
		return x.UsedTimes
	}
	return 0
}

var File_GameplayCounterCountDownCsReq_proto protoreflect.FileDescriptor

var file_GameplayCounterCountDownCsReq_proto_rawDesc = []byte{
	0x0a, 0x23, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x1d, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61,
	0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x77,
	0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x41, 0x4f, 0x4b, 0x41, 0x47,
	0x4e, 0x46, 0x4e, 0x41, 0x42, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x41, 0x4f,
	0x4b, 0x41, 0x47, 0x4e, 0x46, 0x4e, 0x41, 0x42, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GameplayCounterCountDownCsReq_proto_rawDescOnce sync.Once
	file_GameplayCounterCountDownCsReq_proto_rawDescData = file_GameplayCounterCountDownCsReq_proto_rawDesc
)

func file_GameplayCounterCountDownCsReq_proto_rawDescGZIP() []byte {
	file_GameplayCounterCountDownCsReq_proto_rawDescOnce.Do(func() {
		file_GameplayCounterCountDownCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GameplayCounterCountDownCsReq_proto_rawDescData)
	})
	return file_GameplayCounterCountDownCsReq_proto_rawDescData
}

var file_GameplayCounterCountDownCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GameplayCounterCountDownCsReq_proto_goTypes = []interface{}{
	(*GameplayCounterCountDownCsReq)(nil), // 0: GameplayCounterCountDownCsReq
}
var file_GameplayCounterCountDownCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GameplayCounterCountDownCsReq_proto_init() }
func file_GameplayCounterCountDownCsReq_proto_init() {
	if File_GameplayCounterCountDownCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GameplayCounterCountDownCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameplayCounterCountDownCsReq); i {
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
			RawDescriptor: file_GameplayCounterCountDownCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GameplayCounterCountDownCsReq_proto_goTypes,
		DependencyIndexes: file_GameplayCounterCountDownCsReq_proto_depIdxs,
		MessageInfos:      file_GameplayCounterCountDownCsReq_proto_msgTypes,
	}.Build()
	File_GameplayCounterCountDownCsReq_proto = out.File
	file_GameplayCounterCountDownCsReq_proto_rawDesc = nil
	file_GameplayCounterCountDownCsReq_proto_goTypes = nil
	file_GameplayCounterCountDownCsReq_proto_depIdxs = nil
}
