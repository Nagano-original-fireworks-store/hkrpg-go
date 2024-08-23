// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneEntityMoveScRsp.proto

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

type SceneEntityMoveScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityMotionList []*EntityMotion     `protobuf:"bytes,13,rep,name=entity_motion_list,json=entityMotionList,proto3" json:"entity_motion_list,omitempty"`
	DownloadData     *ClientDownloadData `protobuf:"bytes,4,opt,name=download_data,json=downloadData,proto3" json:"download_data,omitempty"`
	Retcode          uint32              `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *SceneEntityMoveScRsp) Reset() {
	*x = SceneEntityMoveScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneEntityMoveScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneEntityMoveScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneEntityMoveScRsp) ProtoMessage() {}

func (x *SceneEntityMoveScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SceneEntityMoveScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneEntityMoveScRsp.ProtoReflect.Descriptor instead.
func (*SceneEntityMoveScRsp) Descriptor() ([]byte, []int) {
	return file_SceneEntityMoveScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SceneEntityMoveScRsp) GetEntityMotionList() []*EntityMotion {
	if x != nil {
		return x.EntityMotionList
	}
	return nil
}

func (x *SceneEntityMoveScRsp) GetDownloadData() *ClientDownloadData {
	if x != nil {
		return x.DownloadData
	}
	return nil
}

func (x *SceneEntityMoveScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_SceneEntityMoveScRsp_proto protoreflect.FileDescriptor

var file_SceneEntityMoveScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x76,
	0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x14, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x12, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneEntityMoveScRsp_proto_rawDescOnce sync.Once
	file_SceneEntityMoveScRsp_proto_rawDescData = file_SceneEntityMoveScRsp_proto_rawDesc
)

func file_SceneEntityMoveScRsp_proto_rawDescGZIP() []byte {
	file_SceneEntityMoveScRsp_proto_rawDescOnce.Do(func() {
		file_SceneEntityMoveScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneEntityMoveScRsp_proto_rawDescData)
	})
	return file_SceneEntityMoveScRsp_proto_rawDescData
}

var file_SceneEntityMoveScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneEntityMoveScRsp_proto_goTypes = []interface{}{
	(*SceneEntityMoveScRsp)(nil), // 0: SceneEntityMoveScRsp
	(*EntityMotion)(nil),         // 1: EntityMotion
	(*ClientDownloadData)(nil),   // 2: ClientDownloadData
}
var file_SceneEntityMoveScRsp_proto_depIdxs = []int32{
	1, // 0: SceneEntityMoveScRsp.entity_motion_list:type_name -> EntityMotion
	2, // 1: SceneEntityMoveScRsp.download_data:type_name -> ClientDownloadData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SceneEntityMoveScRsp_proto_init() }
func file_SceneEntityMoveScRsp_proto_init() {
	if File_SceneEntityMoveScRsp_proto != nil {
		return
	}
	file_EntityMotion_proto_init()
	file_ClientDownloadData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneEntityMoveScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneEntityMoveScRsp); i {
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
			RawDescriptor: file_SceneEntityMoveScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneEntityMoveScRsp_proto_goTypes,
		DependencyIndexes: file_SceneEntityMoveScRsp_proto_depIdxs,
		MessageInfos:      file_SceneEntityMoveScRsp_proto_msgTypes,
	}.Build()
	File_SceneEntityMoveScRsp_proto = out.File
	file_SceneEntityMoveScRsp_proto_rawDesc = nil
	file_SceneEntityMoveScRsp_proto_goTypes = nil
	file_SceneEntityMoveScRsp_proto_depIdxs = nil
}
