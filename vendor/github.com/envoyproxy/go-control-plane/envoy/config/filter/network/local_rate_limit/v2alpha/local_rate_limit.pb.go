// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/local_rate_limit/v2alpha/local_rate_limit.proto

package envoy_config_filter_network_local_rate_limit_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type LocalRateLimit struct {
	StatPrefix           string                   `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	TokenBucket          *_type.TokenBucket       `protobuf:"bytes,2,opt,name=token_bucket,json=tokenBucket,proto3" json:"token_bucket,omitempty"`
	RuntimeEnabled       *core.RuntimeFeatureFlag `protobuf:"bytes,3,opt,name=runtime_enabled,json=runtimeEnabled,proto3" json:"runtime_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LocalRateLimit) Reset()         { *m = LocalRateLimit{} }
func (m *LocalRateLimit) String() string { return proto.CompactTextString(m) }
func (*LocalRateLimit) ProtoMessage()    {}
func (*LocalRateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_984b0e5e5a865836, []int{0}
}

func (m *LocalRateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalRateLimit.Unmarshal(m, b)
}
func (m *LocalRateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalRateLimit.Marshal(b, m, deterministic)
}
func (m *LocalRateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalRateLimit.Merge(m, src)
}
func (m *LocalRateLimit) XXX_Size() int {
	return xxx_messageInfo_LocalRateLimit.Size(m)
}
func (m *LocalRateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalRateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_LocalRateLimit proto.InternalMessageInfo

func (m *LocalRateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *LocalRateLimit) GetTokenBucket() *_type.TokenBucket {
	if m != nil {
		return m.TokenBucket
	}
	return nil
}

func (m *LocalRateLimit) GetRuntimeEnabled() *core.RuntimeFeatureFlag {
	if m != nil {
		return m.RuntimeEnabled
	}
	return nil
}

func init() {
	proto.RegisterType((*LocalRateLimit)(nil), "envoy.config.filter.network.local_rate_limit.v2alpha.LocalRateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/local_rate_limit/v2alpha/local_rate_limit.proto", fileDescriptor_984b0e5e5a865836)
}

var fileDescriptor_984b0e5e5a865836 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xb1, 0x8e, 0xd3, 0x40,
	0x10, 0x86, 0xb5, 0x81, 0x3b, 0xc0, 0x41, 0xc7, 0xc9, 0x14, 0x17, 0x45, 0x80, 0x22, 0x24, 0xa4,
	0x54, 0xbb, 0x52, 0x02, 0x2f, 0xb0, 0x82, 0x6b, 0x38, 0xa1, 0xc8, 0xa2, 0xb7, 0xc6, 0xce, 0x24,
	0xac, 0xb2, 0xd9, 0x5d, 0xad, 0xc7, 0x26, 0xe9, 0xa8, 0x69, 0x68, 0xa9, 0x78, 0x10, 0x9e, 0x80,
	0x12, 0x5e, 0x85, 0xf2, 0x0a, 0x84, 0xd6, 0xbb, 0x11, 0xa0, 0x50, 0x5d, 0x67, 0xf9, 0xf3, 0xff,
	0xcd, 0xf8, 0x9f, 0xec, 0x35, 0x9a, 0xce, 0xee, 0x45, 0x6d, 0xcd, 0x4a, 0xad, 0xc5, 0x4a, 0x69,
	0x42, 0x2f, 0x0c, 0xd2, 0x7b, 0xeb, 0x37, 0x42, 0xdb, 0x1a, 0x74, 0xe9, 0x81, 0xb0, 0xd4, 0x6a,
	0xab, 0x48, 0x74, 0x33, 0xd0, 0xee, 0x1d, 0x1c, 0x01, 0xee, 0xbc, 0x25, 0x9b, 0x3f, 0xef, 0x65,
	0x3c, 0xca, 0x78, 0x94, 0xf1, 0x24, 0xe3, 0x47, 0x99, 0x24, 0x1b, 0x3f, 0x8a, 0x2b, 0x80, 0x53,
	0xa2, 0x9b, 0x89, 0xda, 0x7a, 0x14, 0x15, 0x34, 0x18, 0x9d, 0xe3, 0xc7, 0x91, 0xd2, 0xde, 0xa1,
	0x20, 0xbb, 0x41, 0x53, 0x56, 0x6d, 0xbd, 0xc1, 0x34, 0x72, 0xfc, 0xa4, 0x5d, 0x3a, 0x10, 0x60,
	0x8c, 0x25, 0x20, 0x65, 0x4d, 0x23, 0xb6, 0x6a, 0x1d, 0x86, 0x1c, 0xe2, 0x47, 0xbc, 0x21, 0xa0,
	0xb6, 0x49, 0xf8, 0xa2, 0x03, 0xad, 0x96, 0x40, 0x28, 0x0e, 0x0f, 0x11, 0x3c, 0xfd, 0xce, 0xb2,
	0xb3, 0xab, 0xb0, 0x71, 0x01, 0x84, 0x57, 0x61, 0xdf, 0x7c, 0x9a, 0x0d, 0x43, 0xb6, 0x74, 0x1e,
	0x57, 0x6a, 0x37, 0x62, 0x13, 0x36, 0xbd, 0x27, 0xef, 0x5c, 0xcb, 0xdb, 0x7e, 0x30, 0x61, 0x45,
	0x16, 0xd8, 0xa2, 0x47, 0xf9, 0xcb, 0xec, 0xfe, 0xdf, 0xab, 0x8e, 0x06, 0x13, 0x36, 0x1d, 0xce,
	0x2e, 0x78, 0xac, 0x27, 0xfc, 0x0a, 0x7f, 0x1b, 0xb8, 0xec, 0xb1, 0xbc, 0x7b, 0x2d, 0x4f, 0x3e,
	0xb2, 0xc1, 0x39, 0x2b, 0x86, 0xf4, 0xe7, 0x75, 0xfe, 0x26, 0x7b, 0xe0, 0x5b, 0x43, 0x6a, 0x8b,
	0x25, 0x1a, 0xa8, 0x34, 0x2e, 0x47, 0xb7, 0x7a, 0xd1, 0xb3, 0x24, 0x02, 0xa7, 0x78, 0x37, 0xe3,
	0xa1, 0x31, 0x5e, 0xc4, 0x2f, 0x2f, 0x11, 0xa8, 0xf5, 0x78, 0xa9, 0x61, 0x5d, 0x9c, 0xa5, 0xf4,
	0xab, 0x18, 0x96, 0x5f, 0xd8, 0xcf, 0xcf, 0xbf, 0x3e, 0x9d, 0xbc, 0xc8, 0xe7, 0x31, 0x8e, 0x3b,
	0x42, 0xd3, 0x84, 0x4e, 0xd2, 0xa9, 0x9a, 0xff, 0xdc, 0x2a, 0x9d, 0x6a, 0xfe, 0xf5, 0xc3, 0xb7,
	0x1f, 0xa7, 0x83, 0x73, 0x96, 0x49, 0x65, 0xe3, 0x78, 0xe7, 0xed, 0x6e, 0xcf, 0x6f, 0x72, 0x71,
	0xf9, 0xf0, 0xdf, 0x66, 0x17, 0xa1, 0xf1, 0x05, 0xab, 0x4e, 0xfb, 0xea, 0xe7, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x01, 0xdd, 0x29, 0x5f, 0x94, 0x02, 0x00, 0x00,
}
