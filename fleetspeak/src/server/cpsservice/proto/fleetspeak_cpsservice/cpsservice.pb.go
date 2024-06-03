// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: fleetspeak/src/server/cpsservice/proto/fleetspeak_cpsservice/cpsservice.proto

package fleetspeak_cpsservice

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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Topic   string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"` // The cloud pubsub topic to which messages will be posted
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Config) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

var File_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto protoreflect.FileDescriptor

var file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x70, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70,
	0x65, 0x61, 0x6b, 0x5f, 0x63, 0x70, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x70, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x63, 0x70, 0x73, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x38, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x42, 0x5b, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b,
	0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x70, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65,
	0x61, 0x6b, 0x5f, 0x63, 0x70, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDescOnce sync.Once
	file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDescData = file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDesc
)

func file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDescGZIP() []byte {
	file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDescOnce.Do(func() {
		file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDescData)
	})
	return file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDescData
}

var file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: fleetspeak.cpsservice.Config
}
var file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_init()
}
func file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_init() {
	if File_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_goTypes,
		DependencyIndexes: file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_depIdxs,
		MessageInfos:      file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_msgTypes,
	}.Build()
	File_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto = out.File
	file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_rawDesc = nil
	file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_goTypes = nil
	file_fleetspeak_src_server_cpsservice_proto_fleetspeak_cpsservice_cpsservice_proto_depIdxs = nil
}
