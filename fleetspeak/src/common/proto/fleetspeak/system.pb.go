// The system service is present on every Fleetspeak client and server. Its
// messages implement much of the basic Fleetspeak functionality.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v4.23.4
// source: fleetspeak/src/common/proto/fleetspeak/system.proto

package fleetspeak

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A MessageAck message is sent from the client to the server to acknowledge the
// successful receipt of one or more messages. Messages from the server to the
// client may be repeated until ack'd.
type MessageAckData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageIds [][]byte `protobuf:"bytes,1,rep,name=message_ids,json=messageIds,proto3" json:"message_ids,omitempty"`
}

func (x *MessageAckData) Reset() {
	*x = MessageAckData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAckData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAckData) ProtoMessage() {}

func (x *MessageAckData) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAckData.ProtoReflect.Descriptor instead.
func (*MessageAckData) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP(), []int{0}
}

func (x *MessageAckData) GetMessageIds() [][]byte {
	if x != nil {
		return x.MessageIds
	}
	return nil
}

// A MessageError message is sent from the client to the server to indicate an
// permanent error in processing the message.
type MessageErrorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId []byte `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Error     string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MessageErrorData) Reset() {
	*x = MessageErrorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageErrorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageErrorData) ProtoMessage() {}

func (x *MessageErrorData) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageErrorData.ProtoReflect.Descriptor instead.
func (*MessageErrorData) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP(), []int{1}
}

func (x *MessageErrorData) GetMessageId() []byte {
	if x != nil {
		return x.MessageId
	}
	return nil
}

func (x *MessageErrorData) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// A ClientInfo message is sent from the client to the server on initial contact
// and after every configurate change.
type ClientInfoData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Labels hardcoded by the client build, e.g. "client:Linux",
	// "client:<build-nr>", "client:canary".
	Labels   []*Label                    `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	Services []*ClientInfoData_ServiceID `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ClientInfoData) Reset() {
	*x = ClientInfoData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfoData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfoData) ProtoMessage() {}

func (x *ClientInfoData) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfoData.ProtoReflect.Descriptor instead.
func (*ClientInfoData) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP(), []int{2}
}

func (x *ClientInfoData) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ClientInfoData) GetServices() []*ClientInfoData_ServiceID {
	if x != nil {
		return x.Services
	}
	return nil
}

// A RemoveService message is sent from the server to the client to instruct the
// client to remove an existing service.
type RemoveServiceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RemoveServiceData) Reset() {
	*x = RemoveServiceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveServiceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveServiceData) ProtoMessage() {}

func (x *RemoveServiceData) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveServiceData.ProtoReflect.Descriptor instead.
func (*RemoveServiceData) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveServiceData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A SignedClientServiceConfig wraps a ClientServiceConfig with a signature,
// making it acceptable to clients.
type SignedClientServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A serialized ClientServiceConfig, defined below.
	ServiceConfig []byte `protobuf:"bytes,1,opt,name=service_config,json=serviceConfig,proto3" json:"service_config,omitempty"`
	// An RSASSA-PSS signature of service_config, made using a deployment key.
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedClientServiceConfig) Reset() {
	*x = SignedClientServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedClientServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedClientServiceConfig) ProtoMessage() {}

func (x *SignedClientServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedClientServiceConfig.ProtoReflect.Descriptor instead.
func (*SignedClientServiceConfig) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP(), []int{4}
}

func (x *SignedClientServiceConfig) GetServiceConfig() []byte {
	if x != nil {
		return x.ServiceConfig
	}
	return nil
}

func (x *SignedClientServiceConfig) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ClientServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name that the service will be known as. Primary use is to address
	// message to the service.  the service names 'server' and 'client' are
	// reserved.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name of the factory which will be used to generate the service.
	Factory string `protobuf:"bytes,2,opt,name=factory,proto3" json:"factory,omitempty"`
	// Additional configuration information for the factory to use when setting up
	// the service. The expected type depends upon the factory.
	Config *anypb.Any `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	// The service will only be installed on clients with all of the listed
	// labels. Note that labels for the special 'client' service are checked on
	// the client. All other labels are only checked on the server.
	RequiredLabels []*Label `protobuf:"bytes,6,rep,name=required_labels,json=requiredLabels,proto3" json:"required_labels,omitempty"`
	// The time at which the service configuration was signed. Should be populated
	// by the signing tool when creating a SignedClientServiceConfig.
	SignedTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=signed_time,json=signedTime,proto3" json:"signed_time,omitempty"`
}

func (x *ClientServiceConfig) Reset() {
	*x = ClientServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientServiceConfig) ProtoMessage() {}

func (x *ClientServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientServiceConfig.ProtoReflect.Descriptor instead.
func (*ClientServiceConfig) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP(), []int{5}
}

func (x *ClientServiceConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientServiceConfig) GetFactory() string {
	if x != nil {
		return x.Factory
	}
	return ""
}

func (x *ClientServiceConfig) GetConfig() *anypb.Any {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ClientServiceConfig) GetRequiredLabels() []*Label {
	if x != nil {
		return x.RequiredLabels
	}
	return nil
}

func (x *ClientServiceConfig) GetSignedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SignedTime
	}
	return nil
}

type ClientServiceConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config []*ClientServiceConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
}

func (x *ClientServiceConfigs) Reset() {
	*x = ClientServiceConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientServiceConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientServiceConfigs) ProtoMessage() {}

func (x *ClientServiceConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientServiceConfigs.ProtoReflect.Descriptor instead.
func (*ClientServiceConfigs) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP(), []int{6}
}

func (x *ClientServiceConfigs) GetConfig() []*ClientServiceConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// A list of serial numbers of certificates which should be considered revoked.
type RevokedCertificateList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serials [][]byte `protobuf:"bytes,1,rep,name=serials,proto3" json:"serials,omitempty"`
}

func (x *RevokedCertificateList) Reset() {
	*x = RevokedCertificateList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokedCertificateList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokedCertificateList) ProtoMessage() {}

func (x *RevokedCertificateList) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokedCertificateList.ProtoReflect.Descriptor instead.
func (*RevokedCertificateList) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP(), []int{7}
}

func (x *RevokedCertificateList) GetSerials() [][]byte {
	if x != nil {
		return x.Serials
	}
	return nil
}

type DieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Force bool `protobuf:"varint,1,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *DieRequest) Reset() {
	*x = DieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DieRequest) ProtoMessage() {}

func (x *DieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DieRequest.ProtoReflect.Descriptor instead.
func (*DieRequest) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP(), []int{8}
}

func (x *DieRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type RestartServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RestartServiceRequest) Reset() {
	*x = RestartServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartServiceRequest) ProtoMessage() {}

func (x *RestartServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartServiceRequest.ProtoReflect.Descriptor instead.
func (*RestartServiceRequest) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP(), []int{9}
}

func (x *RestartServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ClientInfoData_ServiceID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the installed service.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The signature of the installed service, as provided in the AddServiceData
	// message which created the service.
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ClientInfoData_ServiceID) Reset() {
	*x = ClientInfoData_ServiceID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfoData_ServiceID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfoData_ServiceID) ProtoMessage() {}

func (x *ClientInfoData_ServiceID) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfoData_ServiceID.ProtoReflect.Descriptor instead.
func (*ClientInfoData_ServiceID) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ClientInfoData_ServiceID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientInfoData_ServiceID) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_fleetspeak_src_common_proto_fleetspeak_system_proto protoreflect.FileDescriptor

var file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDesc = []byte{
	0x0a, 0x33, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61,
	0x6b, 0x1a, 0x33, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x31, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x73, 0x22, 0x47, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xbc,
	0x01, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x29, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x44, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3d,
	0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x27, 0x0a,
	0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x60, 0x0a, 0x19, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xea, 0x01, 0x0a, 0x13, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2c,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3a, 0x0a, 0x0f,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65,
	0x61, 0x6b, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x37, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x32, 0x0a, 0x16, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x07, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x22, 0x0a, 0x0a, 0x44, 0x69,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x2b,
	0x0a, 0x15, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65,
	0x61, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescOnce sync.Once
	file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescData = file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDesc
)

func file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescGZIP() []byte {
	file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescOnce.Do(func() {
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescData)
	})
	return file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDescData
}

var file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_fleetspeak_src_common_proto_fleetspeak_system_proto_goTypes = []interface{}{
	(*MessageAckData)(nil),            // 0: fleetspeak.MessageAckData
	(*MessageErrorData)(nil),          // 1: fleetspeak.MessageErrorData
	(*ClientInfoData)(nil),            // 2: fleetspeak.ClientInfoData
	(*RemoveServiceData)(nil),         // 3: fleetspeak.RemoveServiceData
	(*SignedClientServiceConfig)(nil), // 4: fleetspeak.SignedClientServiceConfig
	(*ClientServiceConfig)(nil),       // 5: fleetspeak.ClientServiceConfig
	(*ClientServiceConfigs)(nil),      // 6: fleetspeak.ClientServiceConfigs
	(*RevokedCertificateList)(nil),    // 7: fleetspeak.RevokedCertificateList
	(*DieRequest)(nil),                // 8: fleetspeak.DieRequest
	(*RestartServiceRequest)(nil),     // 9: fleetspeak.RestartServiceRequest
	(*ClientInfoData_ServiceID)(nil),  // 10: fleetspeak.ClientInfoData.ServiceID
	(*Label)(nil),                     // 11: fleetspeak.Label
	(*anypb.Any)(nil),                 // 12: google.protobuf.Any
	(*timestamppb.Timestamp)(nil),     // 13: google.protobuf.Timestamp
}
var file_fleetspeak_src_common_proto_fleetspeak_system_proto_depIdxs = []int32{
	11, // 0: fleetspeak.ClientInfoData.labels:type_name -> fleetspeak.Label
	10, // 1: fleetspeak.ClientInfoData.services:type_name -> fleetspeak.ClientInfoData.ServiceID
	12, // 2: fleetspeak.ClientServiceConfig.config:type_name -> google.protobuf.Any
	11, // 3: fleetspeak.ClientServiceConfig.required_labels:type_name -> fleetspeak.Label
	13, // 4: fleetspeak.ClientServiceConfig.signed_time:type_name -> google.protobuf.Timestamp
	5,  // 5: fleetspeak.ClientServiceConfigs.config:type_name -> fleetspeak.ClientServiceConfig
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_fleetspeak_src_common_proto_fleetspeak_system_proto_init() }
func file_fleetspeak_src_common_proto_fleetspeak_system_proto_init() {
	if File_fleetspeak_src_common_proto_fleetspeak_system_proto != nil {
		return
	}
	file_fleetspeak_src_common_proto_fleetspeak_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAckData); i {
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
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageErrorData); i {
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
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfoData); i {
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
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveServiceData); i {
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
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedClientServiceConfig); i {
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
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientServiceConfig); i {
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
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientServiceConfigs); i {
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
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokedCertificateList); i {
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
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DieRequest); i {
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
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartServiceRequest); i {
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
		file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfoData_ServiceID); i {
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
			RawDescriptor: file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fleetspeak_src_common_proto_fleetspeak_system_proto_goTypes,
		DependencyIndexes: file_fleetspeak_src_common_proto_fleetspeak_system_proto_depIdxs,
		MessageInfos:      file_fleetspeak_src_common_proto_fleetspeak_system_proto_msgTypes,
	}.Build()
	File_fleetspeak_src_common_proto_fleetspeak_system_proto = out.File
	file_fleetspeak_src_common_proto_fleetspeak_system_proto_rawDesc = nil
	file_fleetspeak_src_common_proto_fleetspeak_system_proto_goTypes = nil
	file_fleetspeak_src_common_proto_fleetspeak_system_proto_depIdxs = nil
}
