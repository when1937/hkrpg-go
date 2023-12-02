// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: EquipRelic.proto

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

type EquipRelic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot          uint32 `protobuf:"varint,12,opt,name=slot,proto3" json:"slot,omitempty"`
	RelicUniqueId uint32 `protobuf:"varint,4,opt,name=relic_unique_id,json=relicUniqueId,proto3" json:"relic_unique_id,omitempty"`
}

func (x *EquipRelic) Reset() {
	*x = EquipRelic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EquipRelic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipRelic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipRelic) ProtoMessage() {}

func (x *EquipRelic) ProtoReflect() protoreflect.Message {
	mi := &file_EquipRelic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipRelic.ProtoReflect.Descriptor instead.
func (*EquipRelic) Descriptor() ([]byte, []int) {
	return file_EquipRelic_proto_rawDescGZIP(), []int{0}
}

func (x *EquipRelic) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *EquipRelic) GetRelicUniqueId() uint32 {
	if x != nil {
		return x.RelicUniqueId
	}
	return 0
}

var File_EquipRelic_proto protoreflect.FileDescriptor

var file_EquipRelic_proto_rawDesc = []byte{
	0x0a, 0x10, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0a, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x65, 0x6c, 0x69, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x73, 0x6c, 0x6f, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x5f, 0x75, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x72,
	0x65, 0x6c, 0x69, 0x63, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EquipRelic_proto_rawDescOnce sync.Once
	file_EquipRelic_proto_rawDescData = file_EquipRelic_proto_rawDesc
)

func file_EquipRelic_proto_rawDescGZIP() []byte {
	file_EquipRelic_proto_rawDescOnce.Do(func() {
		file_EquipRelic_proto_rawDescData = protoimpl.X.CompressGZIP(file_EquipRelic_proto_rawDescData)
	})
	return file_EquipRelic_proto_rawDescData
}

var file_EquipRelic_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EquipRelic_proto_goTypes = []interface{}{
	(*EquipRelic)(nil), // 0: EquipRelic
}
var file_EquipRelic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EquipRelic_proto_init() }
func file_EquipRelic_proto_init() {
	if File_EquipRelic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EquipRelic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipRelic); i {
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
			RawDescriptor: file_EquipRelic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EquipRelic_proto_goTypes,
		DependencyIndexes: file_EquipRelic_proto_depIdxs,
		MessageInfos:      file_EquipRelic_proto_msgTypes,
	}.Build()
	File_EquipRelic_proto = out.File
	file_EquipRelic_proto_rawDesc = nil
	file_EquipRelic_proto_goTypes = nil
	file_EquipRelic_proto_depIdxs = nil
}