// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/server/proto/fleetspeak_server/resource.proto

package fleetspeak_server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents client resource-usage data in the data-store.
// Next id: 13
type ClientResourceUsageRecord struct {
	// Name of the client service that resource usage is charged/attributed to
	// e.g 'system' for the system Fleetspeak service, or the name of a daemon
	// service as specified in its config.
	Scope            string                     `protobuf:"bytes,1,opt,name=scope" json:"scope,omitempty"`
	Pid              int64                      `protobuf:"varint,2,opt,name=pid" json:"pid,omitempty"`
	ProcessStartTime *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=process_start_time,json=processStartTime" json:"process_start_time,omitempty"`
	// When the resource-usage metrics were measured on the client.
	ClientTimestamp *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=client_timestamp,json=clientTimestamp" json:"client_timestamp,omitempty"`
	// When the resource usage record was written to the data-store.
	ServerTimestamp *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=server_timestamp,json=serverTimestamp" json:"server_timestamp,omitempty"`
	// If true, indicates that the process has terminated, and that this is
	// the final resource-usage record for that process.
	ProcessTerminated bool `protobuf:"varint,12,opt,name=process_terminated,json=processTerminated" json:"process_terminated,omitempty"`
	// CPU-usage is in millis per second.
	MeanUserCpuRate       float32 `protobuf:"fixed32,6,opt,name=mean_user_cpu_rate,json=meanUserCpuRate" json:"mean_user_cpu_rate,omitempty"`
	MaxUserCpuRate        float32 `protobuf:"fixed32,7,opt,name=max_user_cpu_rate,json=maxUserCpuRate" json:"max_user_cpu_rate,omitempty"`
	MeanSystemCpuRate     float32 `protobuf:"fixed32,8,opt,name=mean_system_cpu_rate,json=meanSystemCpuRate" json:"mean_system_cpu_rate,omitempty"`
	MaxSystemCpuRate      float32 `protobuf:"fixed32,9,opt,name=max_system_cpu_rate,json=maxSystemCpuRate" json:"max_system_cpu_rate,omitempty"`
	MeanResidentMemoryMib int32   `protobuf:"varint,10,opt,name=mean_resident_memory_mib,json=meanResidentMemoryMib" json:"mean_resident_memory_mib,omitempty"`
	MaxResidentMemoryMib  int32   `protobuf:"varint,11,opt,name=max_resident_memory_mib,json=maxResidentMemoryMib" json:"max_resident_memory_mib,omitempty"`
}

func (m *ClientResourceUsageRecord) Reset()                    { *m = ClientResourceUsageRecord{} }
func (m *ClientResourceUsageRecord) String() string            { return proto.CompactTextString(m) }
func (*ClientResourceUsageRecord) ProtoMessage()               {}
func (*ClientResourceUsageRecord) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ClientResourceUsageRecord) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *ClientResourceUsageRecord) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ClientResourceUsageRecord) GetProcessStartTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.ProcessStartTime
	}
	return nil
}

func (m *ClientResourceUsageRecord) GetClientTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.ClientTimestamp
	}
	return nil
}

func (m *ClientResourceUsageRecord) GetServerTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.ServerTimestamp
	}
	return nil
}

func (m *ClientResourceUsageRecord) GetProcessTerminated() bool {
	if m != nil {
		return m.ProcessTerminated
	}
	return false
}

func (m *ClientResourceUsageRecord) GetMeanUserCpuRate() float32 {
	if m != nil {
		return m.MeanUserCpuRate
	}
	return 0
}

func (m *ClientResourceUsageRecord) GetMaxUserCpuRate() float32 {
	if m != nil {
		return m.MaxUserCpuRate
	}
	return 0
}

func (m *ClientResourceUsageRecord) GetMeanSystemCpuRate() float32 {
	if m != nil {
		return m.MeanSystemCpuRate
	}
	return 0
}

func (m *ClientResourceUsageRecord) GetMaxSystemCpuRate() float32 {
	if m != nil {
		return m.MaxSystemCpuRate
	}
	return 0
}

func (m *ClientResourceUsageRecord) GetMeanResidentMemoryMib() int32 {
	if m != nil {
		return m.MeanResidentMemoryMib
	}
	return 0
}

func (m *ClientResourceUsageRecord) GetMaxResidentMemoryMib() int32 {
	if m != nil {
		return m.MaxResidentMemoryMib
	}
	return 0
}

func init() {
	proto.RegisterType((*ClientResourceUsageRecord)(nil), "fleetspeak.server.ClientResourceUsageRecord")
}

func init() {
	proto.RegisterFile("fleetspeak/src/server/proto/fleetspeak_server/resource.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0xed, 0x76, 0xd9, 0xf5, 0x22, 0x36, 0x31, 0x45, 0x98, 0x5e, 0x88, 0x38, 0x05,
	0xa1, 0x26, 0x12, 0x08, 0x71, 0xe1, 0x56, 0x21, 0x71, 0xe9, 0xc5, 0x6d, 0xcf, 0x96, 0x93, 0x4c,
	0xab, 0x88, 0xba, 0xb6, 0x6c, 0x07, 0xa5, 0xaf, 0xcd, 0x13, 0x20, 0xdb, 0x69, 0x53, 0xca, 0x4a,
	0xbd, 0xb5, 0xf3, 0x7f, 0xff, 0xe7, 0xc9, 0xa0, 0xef, 0x9b, 0x1d, 0x80, 0x35, 0x0a, 0xf8, 0xaf,
	0xc2, 0xe8, 0xaa, 0x30, 0xa0, 0x7f, 0x83, 0x2e, 0x94, 0x96, 0x56, 0x16, 0x43, 0xc6, 0xfa, 0xb9,
	0x06, 0x23, 0x5b, 0x5d, 0x41, 0xee, 0x01, 0x9c, 0x0c, 0x44, 0x1e, 0x88, 0xe9, 0xfb, 0xad, 0x94,
	0xdb, 0x1d, 0x04, 0x43, 0xd9, 0x6e, 0x0a, 0xdb, 0x08, 0x30, 0x96, 0x0b, 0x15, 0x3a, 0x1f, 0xfe,
	0xdc, 0xa2, 0x77, 0xf3, 0x5d, 0x03, 0x7b, 0x4b, 0x7b, 0xd9, 0xda, 0xf0, 0x2d, 0x50, 0xa8, 0xa4,
	0xae, 0xf1, 0x04, 0x8d, 0x4d, 0x25, 0x15, 0x90, 0x28, 0x8d, 0xb2, 0x07, 0x1a, 0xfe, 0xe0, 0x18,
	0x8d, 0x54, 0x53, 0x93, 0x9b, 0x34, 0xca, 0x46, 0xd4, 0xfd, 0xc4, 0x3f, 0x11, 0x56, 0x5a, 0x56,
	0x60, 0x0c, 0x33, 0x96, 0x6b, 0xcb, 0xdc, 0x33, 0x64, 0x94, 0x46, 0xd9, 0xe3, 0xe7, 0x69, 0x1e,
	0x76, 0xc8, 0x8f, 0x3b, 0xe4, 0xab, 0xe3, 0x0e, 0x34, 0xee, 0x5b, 0x4b, 0x57, 0x72, 0x63, 0xfc,
	0x03, 0xc5, 0x95, 0x5f, 0x87, 0x9d, 0x36, 0x25, 0xb7, 0x57, 0x3d, 0x4f, 0xa1, 0x73, 0x1a, 0x38,
	0x4d, 0xb8, 0xc0, 0x99, 0x66, 0x7c, 0x5d, 0x13, 0x3a, 0x83, 0x66, 0x36, 0x7c, 0x97, 0x05, 0x2d,
	0x9a, 0x3d, 0xb7, 0x50, 0x93, 0x97, 0x69, 0x94, 0xdd, 0xd3, 0xa4, 0x4f, 0x56, 0xa7, 0x00, 0x7f,
	0x42, 0x58, 0x00, 0xdf, 0xb3, 0xd6, 0x80, 0x66, 0x95, 0x6a, 0x99, 0xe6, 0x16, 0xc8, 0x5d, 0x1a,
	0x65, 0x37, 0xf4, 0xc9, 0x25, 0x6b, 0x03, 0x7a, 0xae, 0x5a, 0xca, 0x2d, 0xe0, 0x8f, 0x28, 0x11,
	0xbc, 0xbb, 0x60, 0x5f, 0x78, 0xf6, 0x95, 0xe0, 0xdd, 0x39, 0x5a, 0xa0, 0x89, 0xf7, 0x9a, 0x83,
	0xb1, 0x20, 0x06, 0xfa, 0xde, 0xd3, 0x89, 0xcb, 0x96, 0x3e, 0x3a, 0x16, 0x66, 0xe8, 0xb5, 0x73,
	0x5f, 0xf2, 0x0f, 0x9e, 0x8f, 0x05, 0xef, 0xfe, 0xc5, 0xbf, 0x21, 0xe2, 0xfd, 0x1a, 0x4c, 0x53,
	0xbb, 0xdb, 0x0b, 0x10, 0x52, 0x1f, 0x98, 0x68, 0x4a, 0x82, 0xd2, 0x28, 0x1b, 0xd3, 0x37, 0x2e,
	0xa7, 0x7d, 0xbc, 0xf0, 0xe9, 0xa2, 0x29, 0xf1, 0x57, 0xf4, 0xd6, 0xbd, 0xf3, 0x5c, 0xef, 0xd1,
	0xf7, 0x26, 0x82, 0x77, 0xff, 0xd5, 0xca, 0x3b, 0x7f, 0xfb, 0x2f, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x32, 0x32, 0xd4, 0xeb, 0xef, 0x02, 0x00, 0x00,
}
