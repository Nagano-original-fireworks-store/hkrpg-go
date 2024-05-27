// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EnterSceneByServerScNotify.proto

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

type EnterSceneByServerScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lineup *LineupInfo            `protobuf:"bytes,11,opt,name=lineup,proto3" json:"lineup,omitempty"`
	Reason EnterSceneReasonStatus `protobuf:"varint,8,opt,name=reason,proto3,enum=EnterSceneReasonStatus" json:"reason,omitempty"`
	Scene  *SceneInfo             `protobuf:"bytes,15,opt,name=scene,proto3" json:"scene,omitempty"`
}

func (x *EnterSceneByServerScNotify) Reset() {
	*x = EnterSceneByServerScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EnterSceneByServerScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterSceneByServerScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterSceneByServerScNotify) ProtoMessage() {}

func (x *EnterSceneByServerScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EnterSceneByServerScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterSceneByServerScNotify.ProtoReflect.Descriptor instead.
func (*EnterSceneByServerScNotify) Descriptor() ([]byte, []int) {
	return file_EnterSceneByServerScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EnterSceneByServerScNotify) GetLineup() *LineupInfo {
	if x != nil {
		return x.Lineup
	}
	return nil
}

func (x *EnterSceneByServerScNotify) GetReason() EnterSceneReasonStatus {
	if x != nil {
		return x.Reason
	}
	return EnterSceneReasonStatus_ENTER_SCENE_REASON_NONE
}

func (x *EnterSceneByServerScNotify) GetScene() *SceneInfo {
	if x != nil {
		return x.Scene
	}
	return nil
}

var File_EnterSceneByServerScNotify_proto protoreflect.FileDescriptor

var file_EnterSceneByServerScNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x1a, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x23, 0x0a, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EnterSceneByServerScNotify_proto_rawDescOnce sync.Once
	file_EnterSceneByServerScNotify_proto_rawDescData = file_EnterSceneByServerScNotify_proto_rawDesc
)

func file_EnterSceneByServerScNotify_proto_rawDescGZIP() []byte {
	file_EnterSceneByServerScNotify_proto_rawDescOnce.Do(func() {
		file_EnterSceneByServerScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnterSceneByServerScNotify_proto_rawDescData)
	})
	return file_EnterSceneByServerScNotify_proto_rawDescData
}

var file_EnterSceneByServerScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EnterSceneByServerScNotify_proto_goTypes = []interface{}{
	(*EnterSceneByServerScNotify)(nil), // 0: EnterSceneByServerScNotify
	(*LineupInfo)(nil),                 // 1: LineupInfo
	(EnterSceneReasonStatus)(0),        // 2: EnterSceneReasonStatus
	(*SceneInfo)(nil),                  // 3: SceneInfo
}
var file_EnterSceneByServerScNotify_proto_depIdxs = []int32{
	1, // 0: EnterSceneByServerScNotify.lineup:type_name -> LineupInfo
	2, // 1: EnterSceneByServerScNotify.reason:type_name -> EnterSceneReasonStatus
	3, // 2: EnterSceneByServerScNotify.scene:type_name -> SceneInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_EnterSceneByServerScNotify_proto_init() }
func file_EnterSceneByServerScNotify_proto_init() {
	if File_EnterSceneByServerScNotify_proto != nil {
		return
	}
	file_EnterSceneReasonStatus_proto_init()
	file_LineupInfo_proto_init()
	file_SceneInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EnterSceneByServerScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterSceneByServerScNotify); i {
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
			RawDescriptor: file_EnterSceneByServerScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnterSceneByServerScNotify_proto_goTypes,
		DependencyIndexes: file_EnterSceneByServerScNotify_proto_depIdxs,
		MessageInfos:      file_EnterSceneByServerScNotify_proto_msgTypes,
	}.Build()
	File_EnterSceneByServerScNotify_proto = out.File
	file_EnterSceneByServerScNotify_proto_rawDesc = nil
	file_EnterSceneByServerScNotify_proto_goTypes = nil
	file_EnterSceneByServerScNotify_proto_depIdxs = nil
}
