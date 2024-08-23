// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChallengeList.proto

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

type ChallengeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Star        uint32 `protobuf:"varint,15,opt,name=star,proto3" json:"star,omitempty"`
	ChallengeId uint32 `protobuf:"varint,9,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	BLJLFMGLKHM bool   `protobuf:"varint,10,opt,name=BLJLFMGLKHM,proto3" json:"BLJLFMGLKHM,omitempty"`
}

func (x *ChallengeList) Reset() {
	*x = ChallengeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeList_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeList) ProtoMessage() {}

func (x *ChallengeList) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeList_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeList.ProtoReflect.Descriptor instead.
func (*ChallengeList) Descriptor() ([]byte, []int) {
	return file_ChallengeList_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeList) GetStar() uint32 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *ChallengeList) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *ChallengeList) GetBLJLFMGLKHM() bool {
	if x != nil {
		return x.BLJLFMGLKHM
	}
	return false
}

var File_ChallengeList_proto protoreflect.FileDescriptor

var file_ChallengeList_proto_rawDesc = []byte{
	0x0a, 0x13, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x72, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x74, 0x61, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x42, 0x4c, 0x4a, 0x4c, 0x46, 0x4d, 0x47, 0x4c, 0x4b, 0x48, 0x4d, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x42, 0x4c, 0x4a, 0x4c, 0x46, 0x4d, 0x47, 0x4c, 0x4b, 0x48, 0x4d, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeList_proto_rawDescOnce sync.Once
	file_ChallengeList_proto_rawDescData = file_ChallengeList_proto_rawDesc
)

func file_ChallengeList_proto_rawDescGZIP() []byte {
	file_ChallengeList_proto_rawDescOnce.Do(func() {
		file_ChallengeList_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeList_proto_rawDescData)
	})
	return file_ChallengeList_proto_rawDescData
}

var file_ChallengeList_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChallengeList_proto_goTypes = []interface{}{
	(*ChallengeList)(nil), // 0: ChallengeList
}
var file_ChallengeList_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChallengeList_proto_init() }
func file_ChallengeList_proto_init() {
	if File_ChallengeList_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChallengeList_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeList); i {
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
			RawDescriptor: file_ChallengeList_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeList_proto_goTypes,
		DependencyIndexes: file_ChallengeList_proto_depIdxs,
		MessageInfos:      file_ChallengeList_proto_msgTypes,
	}.Build()
	File_ChallengeList_proto = out.File
	file_ChallengeList_proto_rawDesc = nil
	file_ChallengeList_proto_goTypes = nil
	file_ChallengeList_proto_depIdxs = nil
}
