// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/core/v2alpha/health_check_event.proto

package envoy_data_core_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HealthCheckFailureType int32

const (
	HealthCheckFailureType_ACTIVE  HealthCheckFailureType = 0
	HealthCheckFailureType_PASSIVE HealthCheckFailureType = 1
	HealthCheckFailureType_NETWORK HealthCheckFailureType = 2
)

var HealthCheckFailureType_name = map[int32]string{
	0: "ACTIVE",
	1: "PASSIVE",
	2: "NETWORK",
}

var HealthCheckFailureType_value = map[string]int32{
	"ACTIVE":  0,
	"PASSIVE": 1,
	"NETWORK": 2,
}

func (x HealthCheckFailureType) String() string {
	return proto.EnumName(HealthCheckFailureType_name, int32(x))
}

func (HealthCheckFailureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e866c90440508830, []int{0}
}

type HealthCheckerType int32

const (
	HealthCheckerType_HTTP  HealthCheckerType = 0
	HealthCheckerType_TCP   HealthCheckerType = 1
	HealthCheckerType_GRPC  HealthCheckerType = 2
	HealthCheckerType_REDIS HealthCheckerType = 3
)

var HealthCheckerType_name = map[int32]string{
	0: "HTTP",
	1: "TCP",
	2: "GRPC",
	3: "REDIS",
}

var HealthCheckerType_value = map[string]int32{
	"HTTP":  0,
	"TCP":   1,
	"GRPC":  2,
	"REDIS": 3,
}

func (x HealthCheckerType) String() string {
	return proto.EnumName(HealthCheckerType_name, int32(x))
}

func (HealthCheckerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e866c90440508830, []int{1}
}

type HealthCheckEvent struct {
	HealthCheckerType HealthCheckerType `protobuf:"varint,1,opt,name=health_checker_type,json=healthCheckerType,proto3,enum=envoy.data.core.v2alpha.HealthCheckerType" json:"health_checker_type,omitempty"`
	Host              *core.Address     `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	ClusterName       string            `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*HealthCheckEvent_EjectUnhealthyEvent
	//	*HealthCheckEvent_AddHealthyEvent
	//	*HealthCheckEvent_HealthCheckFailureEvent
	//	*HealthCheckEvent_DegradedHealthyHost
	//	*HealthCheckEvent_NoLongerDegradedHost
	Event                isHealthCheckEvent_Event `protobuf_oneof:"event"`
	Timestamp            *timestamp.Timestamp     `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *HealthCheckEvent) Reset()         { *m = HealthCheckEvent{} }
func (m *HealthCheckEvent) String() string { return proto.CompactTextString(m) }
func (*HealthCheckEvent) ProtoMessage()    {}
func (*HealthCheckEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e866c90440508830, []int{0}
}

func (m *HealthCheckEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckEvent.Unmarshal(m, b)
}
func (m *HealthCheckEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckEvent.Marshal(b, m, deterministic)
}
func (m *HealthCheckEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckEvent.Merge(m, src)
}
func (m *HealthCheckEvent) XXX_Size() int {
	return xxx_messageInfo_HealthCheckEvent.Size(m)
}
func (m *HealthCheckEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckEvent.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckEvent proto.InternalMessageInfo

func (m *HealthCheckEvent) GetHealthCheckerType() HealthCheckerType {
	if m != nil {
		return m.HealthCheckerType
	}
	return HealthCheckerType_HTTP
}

func (m *HealthCheckEvent) GetHost() *core.Address {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *HealthCheckEvent) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

type isHealthCheckEvent_Event interface {
	isHealthCheckEvent_Event()
}

type HealthCheckEvent_EjectUnhealthyEvent struct {
	EjectUnhealthyEvent *HealthCheckEjectUnhealthy `protobuf:"bytes,4,opt,name=eject_unhealthy_event,json=ejectUnhealthyEvent,proto3,oneof"`
}

type HealthCheckEvent_AddHealthyEvent struct {
	AddHealthyEvent *HealthCheckAddHealthy `protobuf:"bytes,5,opt,name=add_healthy_event,json=addHealthyEvent,proto3,oneof"`
}

type HealthCheckEvent_HealthCheckFailureEvent struct {
	HealthCheckFailureEvent *HealthCheckFailure `protobuf:"bytes,7,opt,name=health_check_failure_event,json=healthCheckFailureEvent,proto3,oneof"`
}

type HealthCheckEvent_DegradedHealthyHost struct {
	DegradedHealthyHost *DegradedHealthyHost `protobuf:"bytes,8,opt,name=degraded_healthy_host,json=degradedHealthyHost,proto3,oneof"`
}

type HealthCheckEvent_NoLongerDegradedHost struct {
	NoLongerDegradedHost *NoLongerDegradedHost `protobuf:"bytes,9,opt,name=no_longer_degraded_host,json=noLongerDegradedHost,proto3,oneof"`
}

func (*HealthCheckEvent_EjectUnhealthyEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_AddHealthyEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_HealthCheckFailureEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_DegradedHealthyHost) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_NoLongerDegradedHost) isHealthCheckEvent_Event() {}

func (m *HealthCheckEvent) GetEvent() isHealthCheckEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *HealthCheckEvent) GetEjectUnhealthyEvent() *HealthCheckEjectUnhealthy {
	if x, ok := m.GetEvent().(*HealthCheckEvent_EjectUnhealthyEvent); ok {
		return x.EjectUnhealthyEvent
	}
	return nil
}

func (m *HealthCheckEvent) GetAddHealthyEvent() *HealthCheckAddHealthy {
	if x, ok := m.GetEvent().(*HealthCheckEvent_AddHealthyEvent); ok {
		return x.AddHealthyEvent
	}
	return nil
}

func (m *HealthCheckEvent) GetHealthCheckFailureEvent() *HealthCheckFailure {
	if x, ok := m.GetEvent().(*HealthCheckEvent_HealthCheckFailureEvent); ok {
		return x.HealthCheckFailureEvent
	}
	return nil
}

func (m *HealthCheckEvent) GetDegradedHealthyHost() *DegradedHealthyHost {
	if x, ok := m.GetEvent().(*HealthCheckEvent_DegradedHealthyHost); ok {
		return x.DegradedHealthyHost
	}
	return nil
}

func (m *HealthCheckEvent) GetNoLongerDegradedHost() *NoLongerDegradedHost {
	if x, ok := m.GetEvent().(*HealthCheckEvent_NoLongerDegradedHost); ok {
		return x.NoLongerDegradedHost
	}
	return nil
}

func (m *HealthCheckEvent) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheckEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheckEvent_EjectUnhealthyEvent)(nil),
		(*HealthCheckEvent_AddHealthyEvent)(nil),
		(*HealthCheckEvent_HealthCheckFailureEvent)(nil),
		(*HealthCheckEvent_DegradedHealthyHost)(nil),
		(*HealthCheckEvent_NoLongerDegradedHost)(nil),
	}
}

type HealthCheckEjectUnhealthy struct {
	FailureType          HealthCheckFailureType `protobuf:"varint,1,opt,name=failure_type,json=failureType,proto3,enum=envoy.data.core.v2alpha.HealthCheckFailureType" json:"failure_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheckEjectUnhealthy) Reset()         { *m = HealthCheckEjectUnhealthy{} }
func (m *HealthCheckEjectUnhealthy) String() string { return proto.CompactTextString(m) }
func (*HealthCheckEjectUnhealthy) ProtoMessage()    {}
func (*HealthCheckEjectUnhealthy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e866c90440508830, []int{1}
}

func (m *HealthCheckEjectUnhealthy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckEjectUnhealthy.Unmarshal(m, b)
}
func (m *HealthCheckEjectUnhealthy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckEjectUnhealthy.Marshal(b, m, deterministic)
}
func (m *HealthCheckEjectUnhealthy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckEjectUnhealthy.Merge(m, src)
}
func (m *HealthCheckEjectUnhealthy) XXX_Size() int {
	return xxx_messageInfo_HealthCheckEjectUnhealthy.Size(m)
}
func (m *HealthCheckEjectUnhealthy) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckEjectUnhealthy.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckEjectUnhealthy proto.InternalMessageInfo

func (m *HealthCheckEjectUnhealthy) GetFailureType() HealthCheckFailureType {
	if m != nil {
		return m.FailureType
	}
	return HealthCheckFailureType_ACTIVE
}

type HealthCheckAddHealthy struct {
	FirstCheck           bool     `protobuf:"varint,1,opt,name=first_check,json=firstCheck,proto3" json:"first_check,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckAddHealthy) Reset()         { *m = HealthCheckAddHealthy{} }
func (m *HealthCheckAddHealthy) String() string { return proto.CompactTextString(m) }
func (*HealthCheckAddHealthy) ProtoMessage()    {}
func (*HealthCheckAddHealthy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e866c90440508830, []int{2}
}

func (m *HealthCheckAddHealthy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckAddHealthy.Unmarshal(m, b)
}
func (m *HealthCheckAddHealthy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckAddHealthy.Marshal(b, m, deterministic)
}
func (m *HealthCheckAddHealthy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckAddHealthy.Merge(m, src)
}
func (m *HealthCheckAddHealthy) XXX_Size() int {
	return xxx_messageInfo_HealthCheckAddHealthy.Size(m)
}
func (m *HealthCheckAddHealthy) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckAddHealthy.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckAddHealthy proto.InternalMessageInfo

func (m *HealthCheckAddHealthy) GetFirstCheck() bool {
	if m != nil {
		return m.FirstCheck
	}
	return false
}

type HealthCheckFailure struct {
	FailureType          HealthCheckFailureType `protobuf:"varint,1,opt,name=failure_type,json=failureType,proto3,enum=envoy.data.core.v2alpha.HealthCheckFailureType" json:"failure_type,omitempty"`
	FirstCheck           bool                   `protobuf:"varint,2,opt,name=first_check,json=firstCheck,proto3" json:"first_check,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheckFailure) Reset()         { *m = HealthCheckFailure{} }
func (m *HealthCheckFailure) String() string { return proto.CompactTextString(m) }
func (*HealthCheckFailure) ProtoMessage()    {}
func (*HealthCheckFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_e866c90440508830, []int{3}
}

func (m *HealthCheckFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckFailure.Unmarshal(m, b)
}
func (m *HealthCheckFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckFailure.Marshal(b, m, deterministic)
}
func (m *HealthCheckFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckFailure.Merge(m, src)
}
func (m *HealthCheckFailure) XXX_Size() int {
	return xxx_messageInfo_HealthCheckFailure.Size(m)
}
func (m *HealthCheckFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckFailure.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckFailure proto.InternalMessageInfo

func (m *HealthCheckFailure) GetFailureType() HealthCheckFailureType {
	if m != nil {
		return m.FailureType
	}
	return HealthCheckFailureType_ACTIVE
}

func (m *HealthCheckFailure) GetFirstCheck() bool {
	if m != nil {
		return m.FirstCheck
	}
	return false
}

type DegradedHealthyHost struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DegradedHealthyHost) Reset()         { *m = DegradedHealthyHost{} }
func (m *DegradedHealthyHost) String() string { return proto.CompactTextString(m) }
func (*DegradedHealthyHost) ProtoMessage()    {}
func (*DegradedHealthyHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_e866c90440508830, []int{4}
}

func (m *DegradedHealthyHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DegradedHealthyHost.Unmarshal(m, b)
}
func (m *DegradedHealthyHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DegradedHealthyHost.Marshal(b, m, deterministic)
}
func (m *DegradedHealthyHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DegradedHealthyHost.Merge(m, src)
}
func (m *DegradedHealthyHost) XXX_Size() int {
	return xxx_messageInfo_DegradedHealthyHost.Size(m)
}
func (m *DegradedHealthyHost) XXX_DiscardUnknown() {
	xxx_messageInfo_DegradedHealthyHost.DiscardUnknown(m)
}

var xxx_messageInfo_DegradedHealthyHost proto.InternalMessageInfo

type NoLongerDegradedHost struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoLongerDegradedHost) Reset()         { *m = NoLongerDegradedHost{} }
func (m *NoLongerDegradedHost) String() string { return proto.CompactTextString(m) }
func (*NoLongerDegradedHost) ProtoMessage()    {}
func (*NoLongerDegradedHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_e866c90440508830, []int{5}
}

func (m *NoLongerDegradedHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoLongerDegradedHost.Unmarshal(m, b)
}
func (m *NoLongerDegradedHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoLongerDegradedHost.Marshal(b, m, deterministic)
}
func (m *NoLongerDegradedHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoLongerDegradedHost.Merge(m, src)
}
func (m *NoLongerDegradedHost) XXX_Size() int {
	return xxx_messageInfo_NoLongerDegradedHost.Size(m)
}
func (m *NoLongerDegradedHost) XXX_DiscardUnknown() {
	xxx_messageInfo_NoLongerDegradedHost.DiscardUnknown(m)
}

var xxx_messageInfo_NoLongerDegradedHost proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.data.core.v2alpha.HealthCheckFailureType", HealthCheckFailureType_name, HealthCheckFailureType_value)
	proto.RegisterEnum("envoy.data.core.v2alpha.HealthCheckerType", HealthCheckerType_name, HealthCheckerType_value)
	proto.RegisterType((*HealthCheckEvent)(nil), "envoy.data.core.v2alpha.HealthCheckEvent")
	proto.RegisterType((*HealthCheckEjectUnhealthy)(nil), "envoy.data.core.v2alpha.HealthCheckEjectUnhealthy")
	proto.RegisterType((*HealthCheckAddHealthy)(nil), "envoy.data.core.v2alpha.HealthCheckAddHealthy")
	proto.RegisterType((*HealthCheckFailure)(nil), "envoy.data.core.v2alpha.HealthCheckFailure")
	proto.RegisterType((*DegradedHealthyHost)(nil), "envoy.data.core.v2alpha.DegradedHealthyHost")
	proto.RegisterType((*NoLongerDegradedHost)(nil), "envoy.data.core.v2alpha.NoLongerDegradedHost")
}

func init() {
	proto.RegisterFile("envoy/data/core/v2alpha/health_check_event.proto", fileDescriptor_e866c90440508830)
}

var fileDescriptor_e866c90440508830 = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcd, 0x6a, 0xdb, 0x4a,
	0x14, 0xc7, 0x2d, 0xf9, 0xfb, 0x38, 0xdc, 0xab, 0x4c, 0xe2, 0xd8, 0xd7, 0x70, 0x89, 0x31, 0x5c,
	0x08, 0xbe, 0xad, 0x54, 0xdc, 0x4d, 0xa0, 0x50, 0xb0, 0x1c, 0xb7, 0x0e, 0x2d, 0xa9, 0x51, 0xd4,
	0x76, 0x13, 0x10, 0x13, 0x6b, 0x6c, 0x29, 0x95, 0x35, 0x42, 0x1a, 0x9b, 0x7a, 0x57, 0xfa, 0x0a,
	0xdd, 0xf6, 0x49, 0xfa, 0x04, 0xdd, 0xf6, 0x6d, 0x4a, 0x56, 0x45, 0x33, 0x92, 0x9d, 0xf8, 0x03,
	0xd2, 0x45, 0x77, 0x9a, 0x33, 0xf3, 0x3f, 0xbf, 0xf3, 0x29, 0x78, 0x42, 0xfc, 0x39, 0x5d, 0x68,
	0x36, 0x66, 0x58, 0x1b, 0xd1, 0x90, 0x68, 0xf3, 0x0e, 0xf6, 0x02, 0x07, 0x6b, 0x0e, 0xc1, 0x1e,
	0x73, 0xac, 0x91, 0x43, 0x46, 0x1f, 0x2c, 0x32, 0x27, 0x3e, 0x53, 0x83, 0x90, 0x32, 0x8a, 0x6a,
	0x5c, 0xa1, 0xc6, 0x0a, 0x35, 0x56, 0xa8, 0x89, 0xa2, 0x71, 0x2c, 0x5c, 0xe1, 0xc0, 0xd5, 0xe6,
	0x1d, 0xe1, 0x0c, 0xdb, 0x76, 0x48, 0xa2, 0x48, 0x28, 0x1b, 0xc7, 0x13, 0x4a, 0x27, 0x1e, 0xd1,
	0xf8, 0xe9, 0x7a, 0x36, 0xd6, 0x98, 0x3b, 0x25, 0x11, 0xc3, 0xd3, 0x20, 0x79, 0xf0, 0xef, 0xcc,
	0x0e, 0xb0, 0x86, 0x7d, 0x9f, 0x32, 0xcc, 0x5c, 0xea, 0x47, 0x5a, 0xc4, 0x30, 0x9b, 0xa5, 0xfa,
	0xda, 0x1c, 0x7b, 0xae, 0x8d, 0x19, 0xd1, 0xd2, 0x0f, 0x71, 0xd1, 0xfa, 0x5a, 0x00, 0x65, 0xc0,
	0xe3, 0xed, 0xc5, 0xe1, 0xf6, 0xe3, 0x68, 0x91, 0x0d, 0x07, 0x77, 0x73, 0x20, 0xa1, 0xc5, 0x16,
	0x01, 0xa9, 0x4b, 0x4d, 0xe9, 0xe4, 0xaf, 0x4e, 0x5b, 0xdd, 0x91, 0x85, 0x7a, 0xc7, 0x0f, 0x09,
	0xcd, 0x45, 0x40, 0xf4, 0xd2, 0xad, 0x9e, 0xff, 0x2c, 0xc9, 0x8a, 0x64, 0xec, 0x3b, 0xeb, 0x97,
	0x48, 0x85, 0x9c, 0x43, 0x23, 0x56, 0x97, 0x9b, 0xd2, 0x49, 0xa5, 0xd3, 0x48, 0xdc, 0xe2, 0xc0,
	0x55, 0xe7, 0x1d, 0xe1, 0xb8, 0x2b, 0x6a, 0x60, 0xf0, 0x77, 0xa8, 0x0d, 0x7b, 0x23, 0x6f, 0x16,
	0x31, 0x12, 0x5a, 0x3e, 0x9e, 0x92, 0x7a, 0xb6, 0x29, 0x9d, 0x94, 0xf5, 0xe2, 0xad, 0x9e, 0x0b,
	0xe5, 0xa6, 0x64, 0x54, 0x92, 0xcb, 0x0b, 0x3c, 0x25, 0xc8, 0x81, 0x2a, 0xb9, 0x21, 0x23, 0x66,
	0xcd, 0x7c, 0x01, 0x5e, 0x88, 0x46, 0xd4, 0x73, 0x1c, 0xd6, 0x79, 0x48, 0x0e, 0xfd, 0xd8, 0xc1,
	0xdb, 0x54, 0x3f, 0xc8, 0x18, 0x07, 0xe4, 0x9e, 0x45, 0xd4, 0xea, 0x0a, 0xf6, 0xb1, 0x6d, 0x5b,
	0xf7, 0x29, 0x79, 0x4e, 0x51, 0x1f, 0x42, 0xe9, 0xda, 0xf6, 0x60, 0x49, 0xf8, 0x1b, 0x2f, 0x4f,
	0xc2, 0xfb, 0x0d, 0x34, 0xee, 0x4d, 0xd3, 0x18, 0xbb, 0xde, 0x2c, 0x24, 0x09, 0xa6, 0xc8, 0x31,
	0xff, 0x3f, 0x04, 0xf3, 0x42, 0x08, 0x07, 0x19, 0xa3, 0xe6, 0x6c, 0x58, 0x05, 0xeb, 0x1a, 0xaa,
	0x36, 0x99, 0x84, 0xd8, 0x26, 0xab, 0x74, 0x78, 0x83, 0x4a, 0x1c, 0xf3, 0x68, 0x27, 0xe6, 0x2c,
	0x51, 0xa5, 0x79, 0xd0, 0x88, 0xc5, 0xd5, 0xb2, 0x37, 0xcd, 0x68, 0x0c, 0x35, 0x9f, 0x5a, 0x1e,
	0xf5, 0x27, 0x24, 0xb4, 0x56, 0xb4, 0x98, 0x52, 0xe6, 0x94, 0xc7, 0x3b, 0x29, 0x17, 0xf4, 0x35,
	0x97, 0x2d, 0x69, 0x02, 0x73, 0xe8, 0x6f, 0xb1, 0xa3, 0x53, 0x28, 0x2f, 0x37, 0xa4, 0x5e, 0x48,
	0x06, 0x4c, 0xec, 0x90, 0x9a, 0xee, 0x90, 0x6a, 0xa6, 0x2f, 0x8c, 0xd5, 0x63, 0x7d, 0x0f, 0xf2,
	0xbc, 0xb8, 0x28, 0xfb, 0x53, 0x97, 0x5a, 0x0b, 0xf8, 0x67, 0xe7, 0x44, 0xa0, 0x2b, 0xd8, 0x4b,
	0xfb, 0x71, 0x67, 0x3f, 0xb4, 0xdf, 0x68, 0xc7, 0xda, 0x92, 0x54, 0xc6, 0x2b, 0x73, 0xeb, 0x14,
	0xaa, 0x5b, 0xc7, 0x04, 0x1d, 0x43, 0x65, 0xec, 0x86, 0x11, 0x13, 0x23, 0xc1, 0xa9, 0x25, 0x03,
	0xb8, 0x89, 0x3f, 0x6d, 0x7d, 0x91, 0x00, 0x6d, 0xb2, 0xfe, 0x6c, 0xb8, 0xeb, 0x51, 0xc9, 0x1b,
	0x51, 0x55, 0xe1, 0x60, 0xcb, 0xa0, 0xb4, 0x8e, 0xe0, 0x70, 0x5b, 0x67, 0xdb, 0xcf, 0xe1, 0x68,
	0x7b, 0x00, 0x08, 0xa0, 0xd0, 0xed, 0x99, 0xe7, 0xef, 0xfa, 0x4a, 0x06, 0x55, 0xa0, 0x38, 0xec,
	0x5e, 0x5e, 0xc6, 0x07, 0x29, 0x3e, 0x5c, 0xf4, 0xcd, 0xf7, 0x6f, 0x8c, 0x57, 0x8a, 0xdc, 0x7e,
	0x06, 0xfb, 0x1b, 0xff, 0x23, 0x54, 0x82, 0xdc, 0xc0, 0x34, 0x87, 0x4a, 0x06, 0x15, 0x21, 0x6b,
	0xf6, 0x86, 0x8a, 0x14, 0x9b, 0x5e, 0x1a, 0xc3, 0x9e, 0x22, 0xa3, 0x32, 0xe4, 0x8d, 0xfe, 0xd9,
	0xf9, 0xa5, 0x92, 0xd5, 0x07, 0xdf, 0x3e, 0x7d, 0xff, 0x51, 0x90, 0x15, 0x09, 0xfe, 0x73, 0xa9,
	0x28, 0x50, 0x10, 0xd2, 0x8f, 0x8b, 0x5d, 0xb5, 0xd2, 0xab, 0xeb, 0xff, 0xd0, 0x61, 0x3c, 0x64,
	0x43, 0xe9, 0xba, 0xc0, 0xa7, 0xed, 0xe9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x63, 0xcb,
	0xd1, 0x2d, 0x06, 0x00, 0x00,
}
