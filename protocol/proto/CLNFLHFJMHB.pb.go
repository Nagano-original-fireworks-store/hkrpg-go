// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CLNFLHFJMHB.proto

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

type CLNFLHFJMHB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EGLCBNJOMPD []uint32 `protobuf:"varint,12,rep,packed,name=EGLCBNJOMPD,proto3" json:"EGLCBNJOMPD,omitempty"`
	JOCEKHKJJHL bool     `protobuf:"varint,15,opt,name=JOCEKHKJJHL,proto3" json:"JOCEKHKJJHL,omitempty"`
	BILKFNALGLA uint32   `protobuf:"varint,3,opt,name=BILKFNALGLA,proto3" json:"BILKFNALGLA,omitempty"`
	JNEEFFFOIGF uint32   `protobuf:"varint,10,opt,name=JNEEFFFOIGF,proto3" json:"JNEEFFFOIGF,omitempty"`
}

func (x *CLNFLHFJMHB) Reset() {
	*x = CLNFLHFJMHB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CLNFLHFJMHB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CLNFLHFJMHB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CLNFLHFJMHB) ProtoMessage() {}

func (x *CLNFLHFJMHB) ProtoReflect() protoreflect.Message {
	mi := &file_CLNFLHFJMHB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CLNFLHFJMHB.ProtoReflect.Descriptor instead.
func (*CLNFLHFJMHB) Descriptor() ([]byte, []int) {
	return file_CLNFLHFJMHB_proto_rawDescGZIP(), []int{0}
}

func (x *CLNFLHFJMHB) GetEGLCBNJOMPD() []uint32 {
	if x != nil {
		return x.EGLCBNJOMPD
	}
	return nil
}

func (x *CLNFLHFJMHB) GetJOCEKHKJJHL() bool {
	if x != nil {
		return x.JOCEKHKJJHL
	}
	return false
}

func (x *CLNFLHFJMHB) GetBILKFNALGLA() uint32 {
	if x != nil {
		return x.BILKFNALGLA
	}
	return 0
}

func (x *CLNFLHFJMHB) GetJNEEFFFOIGF() uint32 {
	if x != nil {
		return x.JNEEFFFOIGF
	}
	return 0
}

var File_CLNFLHFJMHB_proto protoreflect.FileDescriptor

var file_CLNFLHFJMHB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x4c, 0x4e, 0x46, 0x4c, 0x48, 0x46, 0x4a, 0x4d, 0x48, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x0b, 0x43, 0x4c, 0x4e, 0x46, 0x4c, 0x48, 0x46, 0x4a,
	0x4d, 0x48, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x47, 0x4c, 0x43, 0x42, 0x4e, 0x4a, 0x4f, 0x4d,
	0x50, 0x44, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x47, 0x4c, 0x43, 0x42, 0x4e,
	0x4a, 0x4f, 0x4d, 0x50, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4f, 0x43, 0x45, 0x4b, 0x48, 0x4b,
	0x4a, 0x4a, 0x48, 0x4c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4a, 0x4f, 0x43, 0x45,
	0x4b, 0x48, 0x4b, 0x4a, 0x4a, 0x48, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x49, 0x4c, 0x4b, 0x46,
	0x4e, 0x41, 0x4c, 0x47, 0x4c, 0x41, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x49,
	0x4c, 0x4b, 0x46, 0x4e, 0x41, 0x4c, 0x47, 0x4c, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4e, 0x45,
	0x45, 0x46, 0x46, 0x46, 0x4f, 0x49, 0x47, 0x46, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4a, 0x4e, 0x45, 0x45, 0x46, 0x46, 0x46, 0x4f, 0x49, 0x47, 0x46, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CLNFLHFJMHB_proto_rawDescOnce sync.Once
	file_CLNFLHFJMHB_proto_rawDescData = file_CLNFLHFJMHB_proto_rawDesc
)

func file_CLNFLHFJMHB_proto_rawDescGZIP() []byte {
	file_CLNFLHFJMHB_proto_rawDescOnce.Do(func() {
		file_CLNFLHFJMHB_proto_rawDescData = protoimpl.X.CompressGZIP(file_CLNFLHFJMHB_proto_rawDescData)
	})
	return file_CLNFLHFJMHB_proto_rawDescData
}

var file_CLNFLHFJMHB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CLNFLHFJMHB_proto_goTypes = []interface{}{
	(*CLNFLHFJMHB)(nil), // 0: CLNFLHFJMHB
}
var file_CLNFLHFJMHB_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CLNFLHFJMHB_proto_init() }
func file_CLNFLHFJMHB_proto_init() {
	if File_CLNFLHFJMHB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CLNFLHFJMHB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CLNFLHFJMHB); i {
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
			RawDescriptor: file_CLNFLHFJMHB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CLNFLHFJMHB_proto_goTypes,
		DependencyIndexes: file_CLNFLHFJMHB_proto_depIdxs,
		MessageInfos:      file_CLNFLHFJMHB_proto_msgTypes,
	}.Build()
	File_CLNFLHFJMHB_proto = out.File
	file_CLNFLHFJMHB_proto_rawDesc = nil
	file_CLNFLHFJMHB_proto_goTypes = nil
	file_CLNFLHFJMHB_proto_depIdxs = nil
}
