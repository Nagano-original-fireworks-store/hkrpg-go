// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueCellUpdateNotify.proto

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

type ChessRogueCellUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CellList    []*ChessRogueCell       `protobuf:"bytes,11,rep,name=cell_list,json=cellList,proto3" json:"cell_list,omitempty"`
	BoardId     uint32                  `protobuf:"varint,2,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	BOIPFKLGMNA RogueModifierSourceType `protobuf:"varint,15,opt,name=BOIPFKLGMNA,proto3,enum=RogueModifierSourceType" json:"BOIPFKLGMNA,omitempty"`
	Reason      BHAACIMOJDF             `protobuf:"varint,9,opt,name=reason,proto3,enum=BHAACIMOJDF" json:"reason,omitempty"`
}

func (x *ChessRogueCellUpdateNotify) Reset() {
	*x = ChessRogueCellUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueCellUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueCellUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueCellUpdateNotify) ProtoMessage() {}

func (x *ChessRogueCellUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueCellUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueCellUpdateNotify.ProtoReflect.Descriptor instead.
func (*ChessRogueCellUpdateNotify) Descriptor() ([]byte, []int) {
	return file_ChessRogueCellUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueCellUpdateNotify) GetCellList() []*ChessRogueCell {
	if x != nil {
		return x.CellList
	}
	return nil
}

func (x *ChessRogueCellUpdateNotify) GetBoardId() uint32 {
	if x != nil {
		return x.BoardId
	}
	return 0
}

func (x *ChessRogueCellUpdateNotify) GetBOIPFKLGMNA() RogueModifierSourceType {
	if x != nil {
		return x.BOIPFKLGMNA
	}
	return RogueModifierSourceType_ROGUE_MODIFIER_SOURCE_NONE
}

func (x *ChessRogueCellUpdateNotify) GetReason() BHAACIMOJDF {
	if x != nil {
		return x.Reason
	}
	return BHAACIMOJDF_CHESS_ROGUE_CELL_UPDATE_REASON_NONE
}

var File_ChessRogueCellUpdateNotify_proto protoreflect.FileDescriptor

var file_ChessRogueCellUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x65, 0x6c, 0x6c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x65,
	0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x48, 0x41, 0x41, 0x43, 0x49, 0x4d,
	0x4f, 0x4a, 0x44, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x1a, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2c, 0x0a, 0x09, 0x63, 0x65, 0x6c,
	0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x08, 0x63,
	0x65, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x42, 0x4f, 0x49, 0x50, 0x46, 0x4b, 0x4c, 0x47, 0x4d, 0x4e,
	0x41, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x42, 0x4f, 0x49, 0x50, 0x46, 0x4b, 0x4c, 0x47, 0x4d, 0x4e, 0x41, 0x12, 0x24,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x42, 0x48, 0x41, 0x41, 0x43, 0x49, 0x4d, 0x4f, 0x4a, 0x44, 0x46, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueCellUpdateNotify_proto_rawDescOnce sync.Once
	file_ChessRogueCellUpdateNotify_proto_rawDescData = file_ChessRogueCellUpdateNotify_proto_rawDesc
)

func file_ChessRogueCellUpdateNotify_proto_rawDescGZIP() []byte {
	file_ChessRogueCellUpdateNotify_proto_rawDescOnce.Do(func() {
		file_ChessRogueCellUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueCellUpdateNotify_proto_rawDescData)
	})
	return file_ChessRogueCellUpdateNotify_proto_rawDescData
}

var file_ChessRogueCellUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueCellUpdateNotify_proto_goTypes = []interface{}{
	(*ChessRogueCellUpdateNotify)(nil), // 0: ChessRogueCellUpdateNotify
	(*ChessRogueCell)(nil),             // 1: ChessRogueCell
	(RogueModifierSourceType)(0),       // 2: RogueModifierSourceType
	(BHAACIMOJDF)(0),                   // 3: BHAACIMOJDF
}
var file_ChessRogueCellUpdateNotify_proto_depIdxs = []int32{
	1, // 0: ChessRogueCellUpdateNotify.cell_list:type_name -> ChessRogueCell
	2, // 1: ChessRogueCellUpdateNotify.BOIPFKLGMNA:type_name -> RogueModifierSourceType
	3, // 2: ChessRogueCellUpdateNotify.reason:type_name -> BHAACIMOJDF
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ChessRogueCellUpdateNotify_proto_init() }
func file_ChessRogueCellUpdateNotify_proto_init() {
	if File_ChessRogueCellUpdateNotify_proto != nil {
		return
	}
	file_ChessRogueCell_proto_init()
	file_RogueModifierSourceType_proto_init()
	file_BHAACIMOJDF_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueCellUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueCellUpdateNotify); i {
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
			RawDescriptor: file_ChessRogueCellUpdateNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueCellUpdateNotify_proto_goTypes,
		DependencyIndexes: file_ChessRogueCellUpdateNotify_proto_depIdxs,
		MessageInfos:      file_ChessRogueCellUpdateNotify_proto_msgTypes,
	}.Build()
	File_ChessRogueCellUpdateNotify_proto = out.File
	file_ChessRogueCellUpdateNotify_proto_rawDesc = nil
	file_ChessRogueCellUpdateNotify_proto_goTypes = nil
	file_ChessRogueCellUpdateNotify_proto_depIdxs = nil
}
