// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueChangeyAeonDimensionNotify.proto

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

type ChessRogueChangeyAeonDimensionNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ChessRogueQueryAeon `protobuf:"bytes,6,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ChessRogueChangeyAeonDimensionNotify) Reset() {
	*x = ChessRogueChangeyAeonDimensionNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueChangeyAeonDimensionNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueChangeyAeonDimensionNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueChangeyAeonDimensionNotify) ProtoMessage() {}

func (x *ChessRogueChangeyAeonDimensionNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueChangeyAeonDimensionNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueChangeyAeonDimensionNotify.ProtoReflect.Descriptor instead.
func (*ChessRogueChangeyAeonDimensionNotify) Descriptor() ([]byte, []int) {
	return file_ChessRogueChangeyAeonDimensionNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueChangeyAeonDimensionNotify) GetInfo() *ChessRogueQueryAeon {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_ChessRogueChangeyAeonDimensionNotify_proto protoreflect.FileDescriptor

var file_ChessRogueChangeyAeonDimensionNotify_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x79, 0x41, 0x65, 0x6f, 0x6e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x43, 0x68,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x65, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x24, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x79, 0x41, 0x65, 0x6f, 0x6e,
	0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x65, 0x6f, 0x6e, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ChessRogueChangeyAeonDimensionNotify_proto_rawDescOnce sync.Once
	file_ChessRogueChangeyAeonDimensionNotify_proto_rawDescData = file_ChessRogueChangeyAeonDimensionNotify_proto_rawDesc
)

func file_ChessRogueChangeyAeonDimensionNotify_proto_rawDescGZIP() []byte {
	file_ChessRogueChangeyAeonDimensionNotify_proto_rawDescOnce.Do(func() {
		file_ChessRogueChangeyAeonDimensionNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueChangeyAeonDimensionNotify_proto_rawDescData)
	})
	return file_ChessRogueChangeyAeonDimensionNotify_proto_rawDescData
}

var file_ChessRogueChangeyAeonDimensionNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueChangeyAeonDimensionNotify_proto_goTypes = []interface{}{
	(*ChessRogueChangeyAeonDimensionNotify)(nil), // 0: ChessRogueChangeyAeonDimensionNotify
	(*ChessRogueQueryAeon)(nil),                  // 1: ChessRogueQueryAeon
}
var file_ChessRogueChangeyAeonDimensionNotify_proto_depIdxs = []int32{
	1, // 0: ChessRogueChangeyAeonDimensionNotify.info:type_name -> ChessRogueQueryAeon
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChessRogueChangeyAeonDimensionNotify_proto_init() }
func file_ChessRogueChangeyAeonDimensionNotify_proto_init() {
	if File_ChessRogueChangeyAeonDimensionNotify_proto != nil {
		return
	}
	file_ChessRogueQueryAeon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueChangeyAeonDimensionNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueChangeyAeonDimensionNotify); i {
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
			RawDescriptor: file_ChessRogueChangeyAeonDimensionNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueChangeyAeonDimensionNotify_proto_goTypes,
		DependencyIndexes: file_ChessRogueChangeyAeonDimensionNotify_proto_depIdxs,
		MessageInfos:      file_ChessRogueChangeyAeonDimensionNotify_proto_msgTypes,
	}.Build()
	File_ChessRogueChangeyAeonDimensionNotify_proto = out.File
	file_ChessRogueChangeyAeonDimensionNotify_proto_rawDesc = nil
	file_ChessRogueChangeyAeonDimensionNotify_proto_goTypes = nil
	file_ChessRogueChangeyAeonDimensionNotify_proto_depIdxs = nil
}
