// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SwordTrainingSkillPowerInfo.proto

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

type SwordTrainingSkillPowerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value       uint32              `protobuf:"varint,10,opt,name=value,proto3" json:"value,omitempty"`
	PCLPEHLNDAF uint32              `protobuf:"varint,4,opt,name=PCLPEHLNDAF,proto3" json:"PCLPEHLNDAF,omitempty"`
	Status      SwordTrainingStatus `protobuf:"varint,7,opt,name=Status,proto3,enum=SwordTrainingStatus" json:"Status,omitempty"`
}

func (x *SwordTrainingSkillPowerInfo) Reset() {
	*x = SwordTrainingSkillPowerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SwordTrainingSkillPowerInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwordTrainingSkillPowerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwordTrainingSkillPowerInfo) ProtoMessage() {}

func (x *SwordTrainingSkillPowerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SwordTrainingSkillPowerInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwordTrainingSkillPowerInfo.ProtoReflect.Descriptor instead.
func (*SwordTrainingSkillPowerInfo) Descriptor() ([]byte, []int) {
	return file_SwordTrainingSkillPowerInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SwordTrainingSkillPowerInfo) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *SwordTrainingSkillPowerInfo) GetPCLPEHLNDAF() uint32 {
	if x != nil {
		return x.PCLPEHLNDAF
	}
	return 0
}

func (x *SwordTrainingSkillPowerInfo) GetStatus() SwordTrainingStatus {
	if x != nil {
		return x.Status
	}
	return SwordTrainingStatus_SWORD_TRAINING_STATUS_TYPE_NONE
}

var File_SwordTrainingSkillPowerInfo_proto protoreflect.FileDescriptor

var file_SwordTrainingSkillPowerInfo_proto_rawDesc = []byte{
	0x0a, 0x21, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83,
	0x01, 0x0a, 0x1b, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x43, 0x4c, 0x50, 0x45, 0x48, 0x4c, 0x4e,
	0x44, 0x41, 0x46, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x43, 0x4c, 0x50, 0x45,
	0x48, 0x4c, 0x4e, 0x44, 0x41, 0x46, 0x12, 0x2c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SwordTrainingSkillPowerInfo_proto_rawDescOnce sync.Once
	file_SwordTrainingSkillPowerInfo_proto_rawDescData = file_SwordTrainingSkillPowerInfo_proto_rawDesc
)

func file_SwordTrainingSkillPowerInfo_proto_rawDescGZIP() []byte {
	file_SwordTrainingSkillPowerInfo_proto_rawDescOnce.Do(func() {
		file_SwordTrainingSkillPowerInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SwordTrainingSkillPowerInfo_proto_rawDescData)
	})
	return file_SwordTrainingSkillPowerInfo_proto_rawDescData
}

var file_SwordTrainingSkillPowerInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SwordTrainingSkillPowerInfo_proto_goTypes = []interface{}{
	(*SwordTrainingSkillPowerInfo)(nil), // 0: SwordTrainingSkillPowerInfo
	(SwordTrainingStatus)(0),            // 1: SwordTrainingStatus
}
var file_SwordTrainingSkillPowerInfo_proto_depIdxs = []int32{
	1, // 0: SwordTrainingSkillPowerInfo.Status:type_name -> SwordTrainingStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SwordTrainingSkillPowerInfo_proto_init() }
func file_SwordTrainingSkillPowerInfo_proto_init() {
	if File_SwordTrainingSkillPowerInfo_proto != nil {
		return
	}
	file_SwordTrainingStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SwordTrainingSkillPowerInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwordTrainingSkillPowerInfo); i {
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
			RawDescriptor: file_SwordTrainingSkillPowerInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SwordTrainingSkillPowerInfo_proto_goTypes,
		DependencyIndexes: file_SwordTrainingSkillPowerInfo_proto_depIdxs,
		MessageInfos:      file_SwordTrainingSkillPowerInfo_proto_msgTypes,
	}.Build()
	File_SwordTrainingSkillPowerInfo_proto = out.File
	file_SwordTrainingSkillPowerInfo_proto_rawDesc = nil
	file_SwordTrainingSkillPowerInfo_proto_goTypes = nil
	file_SwordTrainingSkillPowerInfo_proto_depIdxs = nil
}
