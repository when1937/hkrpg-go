// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ServerAnnounceNotify.proto

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

type ServerAnnounceNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnnounceDataList []*ServerAnnounceNotify_AnnounceData `protobuf:"bytes,2,rep,name=announce_data_list,json=announceDataList,proto3" json:"announce_data_list,omitempty"`
}

func (x *ServerAnnounceNotify) Reset() {
	*x = ServerAnnounceNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ServerAnnounceNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerAnnounceNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerAnnounceNotify) ProtoMessage() {}

func (x *ServerAnnounceNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ServerAnnounceNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerAnnounceNotify.ProtoReflect.Descriptor instead.
func (*ServerAnnounceNotify) Descriptor() ([]byte, []int) {
	return file_ServerAnnounceNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ServerAnnounceNotify) GetAnnounceDataList() []*ServerAnnounceNotify_AnnounceData {
	if x != nil {
		return x.AnnounceDataList
	}
	return nil
}

type ServerAnnounceNotify_AnnounceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeginTime                        int64  `protobuf:"varint,1,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	CenterSystemFrequency            uint32 `protobuf:"varint,7,opt,name=center_system_frequency,json=centerSystemFrequency,proto3" json:"center_system_frequency,omitempty"`
	ConfigId                         uint32 `protobuf:"varint,6,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	CountDownText                    string `protobuf:"bytes,5,opt,name=count_down_text,json=countDownText,proto3" json:"count_down_text,omitempty"`
	DungeonConfirmText               string `protobuf:"bytes,14,opt,name=dungeon_confirm_text,json=dungeonConfirmText,proto3" json:"dungeon_confirm_text,omitempty"`
	IsCenterSystemLast_5EveryMinutes bool   `protobuf:"varint,11,opt,name=is_center_system_last_5_every_minutes,json=isCenterSystemLast5EveryMinutes,proto3" json:"is_center_system_last_5_every_minutes,omitempty"`
	CenterSystemText                 string `protobuf:"bytes,8,opt,name=center_system_text,json=centerSystemText,proto3" json:"center_system_text,omitempty"`
	EndTime                          int64  `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	CountDownFrequency               uint32 `protobuf:"varint,9,opt,name=count_down_frequency,json=countDownFrequency,proto3" json:"count_down_frequency,omitempty"`
}

func (x *ServerAnnounceNotify_AnnounceData) Reset() {
	*x = ServerAnnounceNotify_AnnounceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ServerAnnounceNotify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerAnnounceNotify_AnnounceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerAnnounceNotify_AnnounceData) ProtoMessage() {}

func (x *ServerAnnounceNotify_AnnounceData) ProtoReflect() protoreflect.Message {
	mi := &file_ServerAnnounceNotify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerAnnounceNotify_AnnounceData.ProtoReflect.Descriptor instead.
func (*ServerAnnounceNotify_AnnounceData) Descriptor() ([]byte, []int) {
	return file_ServerAnnounceNotify_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ServerAnnounceNotify_AnnounceData) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *ServerAnnounceNotify_AnnounceData) GetCenterSystemFrequency() uint32 {
	if x != nil {
		return x.CenterSystemFrequency
	}
	return 0
}

func (x *ServerAnnounceNotify_AnnounceData) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *ServerAnnounceNotify_AnnounceData) GetCountDownText() string {
	if x != nil {
		return x.CountDownText
	}
	return ""
}

func (x *ServerAnnounceNotify_AnnounceData) GetDungeonConfirmText() string {
	if x != nil {
		return x.DungeonConfirmText
	}
	return ""
}

func (x *ServerAnnounceNotify_AnnounceData) GetIsCenterSystemLast_5EveryMinutes() bool {
	if x != nil {
		return x.IsCenterSystemLast_5EveryMinutes
	}
	return false
}

func (x *ServerAnnounceNotify_AnnounceData) GetCenterSystemText() string {
	if x != nil {
		return x.CenterSystemText
	}
	return ""
}

func (x *ServerAnnounceNotify_AnnounceData) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ServerAnnounceNotify_AnnounceData) GetCountDownFrequency() uint32 {
	if x != nil {
		return x.CountDownFrequency
	}
	return 0
}

var File_ServerAnnounceNotify_proto protoreflect.FileDescriptor

var file_ServerAnnounceNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x04, 0x0a,
	0x14, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x50, 0x0a, 0x12, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x10, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0xa7, 0x03, 0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x12, 0x4e, 0x0a, 0x25, 0x69, 0x73, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x35, 0x5f, 0x65, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1f, 0x69, 0x73, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x4c, 0x61, 0x73, 0x74, 0x35, 0x45, 0x76, 0x65, 0x72, 0x79, 0x4d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x14, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x66, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ServerAnnounceNotify_proto_rawDescOnce sync.Once
	file_ServerAnnounceNotify_proto_rawDescData = file_ServerAnnounceNotify_proto_rawDesc
)

func file_ServerAnnounceNotify_proto_rawDescGZIP() []byte {
	file_ServerAnnounceNotify_proto_rawDescOnce.Do(func() {
		file_ServerAnnounceNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ServerAnnounceNotify_proto_rawDescData)
	})
	return file_ServerAnnounceNotify_proto_rawDescData
}

var file_ServerAnnounceNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ServerAnnounceNotify_proto_goTypes = []interface{}{
	(*ServerAnnounceNotify)(nil),              // 0: ServerAnnounceNotify
	(*ServerAnnounceNotify_AnnounceData)(nil), // 1: ServerAnnounceNotify.AnnounceData
}
var file_ServerAnnounceNotify_proto_depIdxs = []int32{
	1, // 0: ServerAnnounceNotify.announce_data_list:type_name -> ServerAnnounceNotify.AnnounceData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ServerAnnounceNotify_proto_init() }
func file_ServerAnnounceNotify_proto_init() {
	if File_ServerAnnounceNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ServerAnnounceNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerAnnounceNotify); i {
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
		file_ServerAnnounceNotify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerAnnounceNotify_AnnounceData); i {
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
			RawDescriptor: file_ServerAnnounceNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ServerAnnounceNotify_proto_goTypes,
		DependencyIndexes: file_ServerAnnounceNotify_proto_depIdxs,
		MessageInfos:      file_ServerAnnounceNotify_proto_msgTypes,
	}.Build()
	File_ServerAnnounceNotify_proto = out.File
	file_ServerAnnounceNotify_proto_rawDesc = nil
	file_ServerAnnounceNotify_proto_goTypes = nil
	file_ServerAnnounceNotify_proto_depIdxs = nil
}