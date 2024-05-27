// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DailyActiveLevel.proto

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

type DailyActiveLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level                 uint32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	WorldLevel            uint32 `protobuf:"varint,5,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	HasTakenDailyActivity bool   `protobuf:"varint,4,opt,name=has_taken_daily_activity,json=hasTakenDailyActivity,proto3" json:"has_taken_daily_activity,omitempty"`
	DailyActivePoint      uint32 `protobuf:"varint,1,opt,name=daily_active_point,json=dailyActivePoint,proto3" json:"daily_active_point,omitempty"`
}

func (x *DailyActiveLevel) Reset() {
	*x = DailyActiveLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DailyActiveLevel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyActiveLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyActiveLevel) ProtoMessage() {}

func (x *DailyActiveLevel) ProtoReflect() protoreflect.Message {
	mi := &file_DailyActiveLevel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyActiveLevel.ProtoReflect.Descriptor instead.
func (*DailyActiveLevel) Descriptor() ([]byte, []int) {
	return file_DailyActiveLevel_proto_rawDescGZIP(), []int{0}
}

func (x *DailyActiveLevel) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *DailyActiveLevel) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *DailyActiveLevel) GetHasTakenDailyActivity() bool {
	if x != nil {
		return x.HasTakenDailyActivity
	}
	return false
}

func (x *DailyActiveLevel) GetDailyActivePoint() uint32 {
	if x != nil {
		return x.DailyActivePoint
	}
	return 0
}

var File_DailyActiveLevel_proto protoreflect.FileDescriptor

var file_DailyActiveLevel_proto_rawDesc = []byte{
	0x0a, 0x16, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x10, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x37, 0x0a, 0x18, 0x68, 0x61, 0x73, 0x5f, 0x74, 0x61, 0x6b, 0x65,
	0x6e, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x68, 0x61, 0x73, 0x54, 0x61, 0x6b, 0x65, 0x6e,
	0x44, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a,
	0x12, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x64, 0x61, 0x69, 0x6c, 0x79,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DailyActiveLevel_proto_rawDescOnce sync.Once
	file_DailyActiveLevel_proto_rawDescData = file_DailyActiveLevel_proto_rawDesc
)

func file_DailyActiveLevel_proto_rawDescGZIP() []byte {
	file_DailyActiveLevel_proto_rawDescOnce.Do(func() {
		file_DailyActiveLevel_proto_rawDescData = protoimpl.X.CompressGZIP(file_DailyActiveLevel_proto_rawDescData)
	})
	return file_DailyActiveLevel_proto_rawDescData
}

var file_DailyActiveLevel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DailyActiveLevel_proto_goTypes = []interface{}{
	(*DailyActiveLevel)(nil), // 0: DailyActiveLevel
}
var file_DailyActiveLevel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DailyActiveLevel_proto_init() }
func file_DailyActiveLevel_proto_init() {
	if File_DailyActiveLevel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DailyActiveLevel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyActiveLevel); i {
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
			RawDescriptor: file_DailyActiveLevel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DailyActiveLevel_proto_goTypes,
		DependencyIndexes: file_DailyActiveLevel_proto_depIdxs,
		MessageInfos:      file_DailyActiveLevel_proto_msgTypes,
	}.Build()
	File_DailyActiveLevel_proto = out.File
	file_DailyActiveLevel_proto_rawDesc = nil
	file_DailyActiveLevel_proto_goTypes = nil
	file_DailyActiveLevel_proto_depIdxs = nil
}
