// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TravelBrochureUpdatePasterPosScRsp.proto

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

type TravelBrochureUpdatePasterPosScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CDIBPKMHDCO *BBJEJALFHEN `protobuf:"bytes,15,opt,name=CDIBPKMHDCO,proto3" json:"CDIBPKMHDCO,omitempty"`
	Retcode     uint32       `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *TravelBrochureUpdatePasterPosScRsp) Reset() {
	*x = TravelBrochureUpdatePasterPosScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TravelBrochureUpdatePasterPosScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelBrochureUpdatePasterPosScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelBrochureUpdatePasterPosScRsp) ProtoMessage() {}

func (x *TravelBrochureUpdatePasterPosScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TravelBrochureUpdatePasterPosScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelBrochureUpdatePasterPosScRsp.ProtoReflect.Descriptor instead.
func (*TravelBrochureUpdatePasterPosScRsp) Descriptor() ([]byte, []int) {
	return file_TravelBrochureUpdatePasterPosScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TravelBrochureUpdatePasterPosScRsp) GetCDIBPKMHDCO() *BBJEJALFHEN {
	if x != nil {
		return x.CDIBPKMHDCO
	}
	return nil
}

func (x *TravelBrochureUpdatePasterPosScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_TravelBrochureUpdatePasterPosScRsp_proto protoreflect.FileDescriptor

var file_TravelBrochureUpdatePasterPosScRsp_proto_rawDesc = []byte{
	0x0a, 0x28, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x42, 0x72, 0x6f, 0x63, 0x68, 0x75, 0x72, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x42, 0x4a, 0x45,
	0x4a, 0x41, 0x4c, 0x46, 0x48, 0x45, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a,
	0x22, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x42, 0x72, 0x6f, 0x63, 0x68, 0x75, 0x72, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x43, 0x44, 0x49, 0x42, 0x50, 0x4b, 0x4d, 0x48, 0x44,
	0x43, 0x4f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x42, 0x4a, 0x45, 0x4a,
	0x41, 0x4c, 0x46, 0x48, 0x45, 0x4e, 0x52, 0x0b, 0x43, 0x44, 0x49, 0x42, 0x50, 0x4b, 0x4d, 0x48,
	0x44, 0x43, 0x4f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_TravelBrochureUpdatePasterPosScRsp_proto_rawDescOnce sync.Once
	file_TravelBrochureUpdatePasterPosScRsp_proto_rawDescData = file_TravelBrochureUpdatePasterPosScRsp_proto_rawDesc
)

func file_TravelBrochureUpdatePasterPosScRsp_proto_rawDescGZIP() []byte {
	file_TravelBrochureUpdatePasterPosScRsp_proto_rawDescOnce.Do(func() {
		file_TravelBrochureUpdatePasterPosScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TravelBrochureUpdatePasterPosScRsp_proto_rawDescData)
	})
	return file_TravelBrochureUpdatePasterPosScRsp_proto_rawDescData
}

var file_TravelBrochureUpdatePasterPosScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TravelBrochureUpdatePasterPosScRsp_proto_goTypes = []interface{}{
	(*TravelBrochureUpdatePasterPosScRsp)(nil), // 0: TravelBrochureUpdatePasterPosScRsp
	(*BBJEJALFHEN)(nil),                        // 1: BBJEJALFHEN
}
var file_TravelBrochureUpdatePasterPosScRsp_proto_depIdxs = []int32{
	1, // 0: TravelBrochureUpdatePasterPosScRsp.CDIBPKMHDCO:type_name -> BBJEJALFHEN
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TravelBrochureUpdatePasterPosScRsp_proto_init() }
func file_TravelBrochureUpdatePasterPosScRsp_proto_init() {
	if File_TravelBrochureUpdatePasterPosScRsp_proto != nil {
		return
	}
	file_BBJEJALFHEN_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TravelBrochureUpdatePasterPosScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelBrochureUpdatePasterPosScRsp); i {
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
			RawDescriptor: file_TravelBrochureUpdatePasterPosScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TravelBrochureUpdatePasterPosScRsp_proto_goTypes,
		DependencyIndexes: file_TravelBrochureUpdatePasterPosScRsp_proto_depIdxs,
		MessageInfos:      file_TravelBrochureUpdatePasterPosScRsp_proto_msgTypes,
	}.Build()
	File_TravelBrochureUpdatePasterPosScRsp_proto = out.File
	file_TravelBrochureUpdatePasterPosScRsp_proto_rawDesc = nil
	file_TravelBrochureUpdatePasterPosScRsp_proto_goTypes = nil
	file_TravelBrochureUpdatePasterPosScRsp_proto_depIdxs = nil
}
