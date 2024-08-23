// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetBagScRsp.proto

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

type GetBagScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MEHHHLONJDO         []uint32           `protobuf:"varint,12,rep,packed,name=MEHHHLONJDO,proto3" json:"MEHHHLONJDO,omitempty"`
	LJKCANIHCEP         []*Material0       `protobuf:"bytes,6,rep,name=LJKCANIHCEP,proto3" json:"LJKCANIHCEP,omitempty"`
	LNALMACDFMO         uint32             `protobuf:"varint,13,opt,name=LNALMACDFMO,proto3" json:"LNALMACDFMO,omitempty"`
	MaterialList        []*Material        `protobuf:"bytes,10,rep,name=material_list,json=materialList,proto3" json:"material_list,omitempty"`
	EquipmentList       []*Equipment       `protobuf:"bytes,14,rep,name=equipment_list,json=equipmentList,proto3" json:"equipment_list,omitempty"`
	Retcode             uint32             `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	TurnFoodSwitch      []TurnFoodSwitch   `protobuf:"varint,3,rep,packed,name=turn_food_switch,json=turnFoodSwitch,proto3,enum=TurnFoodSwitch" json:"turn_food_switch,omitempty"`
	APEKBJACDDH         []*Material        `protobuf:"bytes,4,rep,name=APEKBJACDDH,proto3" json:"APEKBJACDDH,omitempty"`
	UnlockFormulaList   []uint32           `protobuf:"varint,2,rep,packed,name=unlock_formula_list,json=unlockFormulaList,proto3" json:"unlock_formula_list,omitempty"`
	RelicList           []*Relic           `protobuf:"bytes,9,rep,name=relic_list,json=relicList,proto3" json:"relic_list,omitempty"`
	PileItemList        []*PileItem        `protobuf:"bytes,5,rep,name=pile_item_list,json=pileItemList,proto3" json:"pile_item_list,omitempty"`
	NOGKOKELAKC         []*Material0       `protobuf:"bytes,1,rep,name=NOGKOKELAKC,proto3" json:"NOGKOKELAKC,omitempty"`
	WaitDelResourceList []*WaitDelResource `protobuf:"bytes,15,rep,name=wait_del_resource_list,json=waitDelResourceList,proto3" json:"wait_del_resource_list,omitempty"`
}

func (x *GetBagScRsp) Reset() {
	*x = GetBagScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetBagScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBagScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBagScRsp) ProtoMessage() {}

func (x *GetBagScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetBagScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBagScRsp.ProtoReflect.Descriptor instead.
func (*GetBagScRsp) Descriptor() ([]byte, []int) {
	return file_GetBagScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetBagScRsp) GetMEHHHLONJDO() []uint32 {
	if x != nil {
		return x.MEHHHLONJDO
	}
	return nil
}

func (x *GetBagScRsp) GetLJKCANIHCEP() []*Material0 {
	if x != nil {
		return x.LJKCANIHCEP
	}
	return nil
}

func (x *GetBagScRsp) GetLNALMACDFMO() uint32 {
	if x != nil {
		return x.LNALMACDFMO
	}
	return 0
}

func (x *GetBagScRsp) GetMaterialList() []*Material {
	if x != nil {
		return x.MaterialList
	}
	return nil
}

func (x *GetBagScRsp) GetEquipmentList() []*Equipment {
	if x != nil {
		return x.EquipmentList
	}
	return nil
}

func (x *GetBagScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetBagScRsp) GetTurnFoodSwitch() []TurnFoodSwitch {
	if x != nil {
		return x.TurnFoodSwitch
	}
	return nil
}

func (x *GetBagScRsp) GetAPEKBJACDDH() []*Material {
	if x != nil {
		return x.APEKBJACDDH
	}
	return nil
}

func (x *GetBagScRsp) GetUnlockFormulaList() []uint32 {
	if x != nil {
		return x.UnlockFormulaList
	}
	return nil
}

func (x *GetBagScRsp) GetRelicList() []*Relic {
	if x != nil {
		return x.RelicList
	}
	return nil
}

func (x *GetBagScRsp) GetPileItemList() []*PileItem {
	if x != nil {
		return x.PileItemList
	}
	return nil
}

func (x *GetBagScRsp) GetNOGKOKELAKC() []*Material0 {
	if x != nil {
		return x.NOGKOKELAKC
	}
	return nil
}

func (x *GetBagScRsp) GetWaitDelResourceList() []*WaitDelResource {
	if x != nil {
		return x.WaitDelResourceList
	}
	return nil
}

var File_GetBagScRsp_proto protoreflect.FileDescriptor

var file_GetBagScRsp_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x67, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x54, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x6f, 0x64, 0x53, 0x77, 0x69,
	0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x52, 0x65, 0x6c, 0x69, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x57, 0x61, 0x69, 0x74, 0x44, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x50, 0x69, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe1, 0x04, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x61, 0x67, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12,
	0x20, 0x0a, 0x0b, 0x4d, 0x45, 0x48, 0x48, 0x48, 0x4c, 0x4f, 0x4e, 0x4a, 0x44, 0x4f, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x45, 0x48, 0x48, 0x48, 0x4c, 0x4f, 0x4e, 0x4a, 0x44,
	0x4f, 0x12, 0x2c, 0x0a, 0x0b, 0x4c, 0x4a, 0x4b, 0x43, 0x41, 0x4e, 0x49, 0x48, 0x43, 0x45, 0x50,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x30, 0x52, 0x0b, 0x4c, 0x4a, 0x4b, 0x43, 0x41, 0x4e, 0x49, 0x48, 0x43, 0x45, 0x50, 0x12,
	0x20, 0x0a, 0x0b, 0x4c, 0x4e, 0x41, 0x4c, 0x4d, 0x41, 0x43, 0x44, 0x46, 0x4d, 0x4f, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x4e, 0x41, 0x4c, 0x4d, 0x41, 0x43, 0x44, 0x46, 0x4d,
	0x4f, 0x12, 0x2e, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x0e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x45, 0x71, 0x75, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x39,
	0x0a, 0x10, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x73, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x54, 0x75, 0x72, 0x6e, 0x46,
	0x6f, 0x6f, 0x64, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x0e, 0x74, 0x75, 0x72, 0x6e, 0x46,
	0x6f, 0x6f, 0x64, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x2b, 0x0a, 0x0b, 0x41, 0x50, 0x45,
	0x4b, 0x42, 0x4a, 0x41, 0x43, 0x44, 0x44, 0x48, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x0b, 0x41, 0x50, 0x45, 0x4b, 0x42,
	0x4a, 0x41, 0x43, 0x44, 0x44, 0x48, 0x12, 0x2e, 0x0a, 0x13, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x11, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x6d, 0x75,
	0x6c, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0a, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x52, 0x65, 0x6c,
	0x69, 0x63, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a,
	0x0e, 0x70, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x69, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x0c, 0x70, 0x69, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x0b, 0x4e, 0x4f, 0x47, 0x4b, 0x4f, 0x4b, 0x45, 0x4c, 0x41, 0x4b, 0x43, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x30, 0x52,
	0x0b, 0x4e, 0x4f, 0x47, 0x4b, 0x4f, 0x4b, 0x45, 0x4c, 0x41, 0x4b, 0x43, 0x12, 0x45, 0x0a, 0x16,
	0x77, 0x61, 0x69, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x57,
	0x61, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x13,
	0x77, 0x61, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetBagScRsp_proto_rawDescOnce sync.Once
	file_GetBagScRsp_proto_rawDescData = file_GetBagScRsp_proto_rawDesc
)

func file_GetBagScRsp_proto_rawDescGZIP() []byte {
	file_GetBagScRsp_proto_rawDescOnce.Do(func() {
		file_GetBagScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetBagScRsp_proto_rawDescData)
	})
	return file_GetBagScRsp_proto_rawDescData
}

var file_GetBagScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetBagScRsp_proto_goTypes = []interface{}{
	(*GetBagScRsp)(nil),     // 0: GetBagScRsp
	(*Material0)(nil),       // 1: Material0
	(*Material)(nil),        // 2: Material
	(*Equipment)(nil),       // 3: Equipment
	(TurnFoodSwitch)(0),     // 4: TurnFoodSwitch
	(*Relic)(nil),           // 5: Relic
	(*PileItem)(nil),        // 6: PileItem
	(*WaitDelResource)(nil), // 7: WaitDelResource
}
var file_GetBagScRsp_proto_depIdxs = []int32{
	1, // 0: GetBagScRsp.LJKCANIHCEP:type_name -> Material0
	2, // 1: GetBagScRsp.material_list:type_name -> Material
	3, // 2: GetBagScRsp.equipment_list:type_name -> Equipment
	4, // 3: GetBagScRsp.turn_food_switch:type_name -> TurnFoodSwitch
	2, // 4: GetBagScRsp.APEKBJACDDH:type_name -> Material
	5, // 5: GetBagScRsp.relic_list:type_name -> Relic
	6, // 6: GetBagScRsp.pile_item_list:type_name -> PileItem
	1, // 7: GetBagScRsp.NOGKOKELAKC:type_name -> Material0
	7, // 8: GetBagScRsp.wait_del_resource_list:type_name -> WaitDelResource
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_GetBagScRsp_proto_init() }
func file_GetBagScRsp_proto_init() {
	if File_GetBagScRsp_proto != nil {
		return
	}
	file_TurnFoodSwitch_proto_init()
	file_Relic_proto_init()
	file_Equipment_proto_init()
	file_WaitDelResource_proto_init()
	file_Material0_proto_init()
	file_Material_proto_init()
	file_PileItem_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetBagScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBagScRsp); i {
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
			RawDescriptor: file_GetBagScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetBagScRsp_proto_goTypes,
		DependencyIndexes: file_GetBagScRsp_proto_depIdxs,
		MessageInfos:      file_GetBagScRsp_proto_msgTypes,
	}.Build()
	File_GetBagScRsp_proto = out.File
	file_GetBagScRsp_proto_rawDesc = nil
	file_GetBagScRsp_proto_goTypes = nil
	file_GetBagScRsp_proto_depIdxs = nil
}
