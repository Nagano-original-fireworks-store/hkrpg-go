// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AlleyShipUsedCountScNotify.proto

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

type AlleyShipUsedCountScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GEFOCNOCKGM map[uint32]uint32 `protobuf:"bytes,11,rep,name=GEFOCNOCKGM,proto3" json:"GEFOCNOCKGM,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *AlleyShipUsedCountScNotify) Reset() {
	*x = AlleyShipUsedCountScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AlleyShipUsedCountScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlleyShipUsedCountScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlleyShipUsedCountScNotify) ProtoMessage() {}

func (x *AlleyShipUsedCountScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AlleyShipUsedCountScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlleyShipUsedCountScNotify.ProtoReflect.Descriptor instead.
func (*AlleyShipUsedCountScNotify) Descriptor() ([]byte, []int) {
	return file_AlleyShipUsedCountScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AlleyShipUsedCountScNotify) GetGEFOCNOCKGM() map[uint32]uint32 {
	if x != nil {
		return x.GEFOCNOCKGM
	}
	return nil
}

var File_AlleyShipUsedCountScNotify_proto protoreflect.FileDescriptor

var file_AlleyShipUsedCountScNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x53, 0x68, 0x69, 0x70, 0x55, 0x73, 0x65, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x1a, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x53, 0x68, 0x69, 0x70,
	0x55, 0x73, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x45, 0x46, 0x4f, 0x43, 0x4e, 0x4f, 0x43, 0x4b, 0x47, 0x4d,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x53, 0x68,
	0x69, 0x70, 0x55, 0x73, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x63, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x47, 0x45, 0x46, 0x4f, 0x43, 0x4e, 0x4f, 0x43, 0x4b, 0x47, 0x4d, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x47, 0x45, 0x46, 0x4f, 0x43, 0x4e, 0x4f, 0x43, 0x4b, 0x47,
	0x4d, 0x1a, 0x3e, 0x0a, 0x10, 0x47, 0x45, 0x46, 0x4f, 0x43, 0x4e, 0x4f, 0x43, 0x4b, 0x47, 0x4d,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AlleyShipUsedCountScNotify_proto_rawDescOnce sync.Once
	file_AlleyShipUsedCountScNotify_proto_rawDescData = file_AlleyShipUsedCountScNotify_proto_rawDesc
)

func file_AlleyShipUsedCountScNotify_proto_rawDescGZIP() []byte {
	file_AlleyShipUsedCountScNotify_proto_rawDescOnce.Do(func() {
		file_AlleyShipUsedCountScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AlleyShipUsedCountScNotify_proto_rawDescData)
	})
	return file_AlleyShipUsedCountScNotify_proto_rawDescData
}

var file_AlleyShipUsedCountScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AlleyShipUsedCountScNotify_proto_goTypes = []interface{}{
	(*AlleyShipUsedCountScNotify)(nil), // 0: AlleyShipUsedCountScNotify
	nil,                                // 1: AlleyShipUsedCountScNotify.GEFOCNOCKGMEntry
}
var file_AlleyShipUsedCountScNotify_proto_depIdxs = []int32{
	1, // 0: AlleyShipUsedCountScNotify.GEFOCNOCKGM:type_name -> AlleyShipUsedCountScNotify.GEFOCNOCKGMEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AlleyShipUsedCountScNotify_proto_init() }
func file_AlleyShipUsedCountScNotify_proto_init() {
	if File_AlleyShipUsedCountScNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AlleyShipUsedCountScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlleyShipUsedCountScNotify); i {
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
			RawDescriptor: file_AlleyShipUsedCountScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AlleyShipUsedCountScNotify_proto_goTypes,
		DependencyIndexes: file_AlleyShipUsedCountScNotify_proto_depIdxs,
		MessageInfos:      file_AlleyShipUsedCountScNotify_proto_msgTypes,
	}.Build()
	File_AlleyShipUsedCountScNotify_proto = out.File
	file_AlleyShipUsedCountScNotify_proto_rawDesc = nil
	file_AlleyShipUsedCountScNotify_proto_goTypes = nil
	file_AlleyShipUsedCountScNotify_proto_depIdxs = nil
}
