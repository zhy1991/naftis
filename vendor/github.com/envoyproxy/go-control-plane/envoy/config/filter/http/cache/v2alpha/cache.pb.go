// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/cache/v2alpha/cache.proto

package envoy_config_filter_http_cache_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type CacheConfig struct {
	TypedConfig          *any.Any                      `protobuf:"bytes,1,opt,name=typed_config,json=typedConfig,proto3" json:"typed_config,omitempty"`
	AllowedVaryHeaders   []*matcher.StringMatcher      `protobuf:"bytes,2,rep,name=allowed_vary_headers,json=allowedVaryHeaders,proto3" json:"allowed_vary_headers,omitempty"`
	KeyCreatorParams     *CacheConfig_KeyCreatorParams `protobuf:"bytes,3,opt,name=key_creator_params,json=keyCreatorParams,proto3" json:"key_creator_params,omitempty"`
	MaxBodyBytes         uint32                        `protobuf:"varint,4,opt,name=max_body_bytes,json=maxBodyBytes,proto3" json:"max_body_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CacheConfig) Reset()         { *m = CacheConfig{} }
func (m *CacheConfig) String() string { return proto.CompactTextString(m) }
func (*CacheConfig) ProtoMessage()    {}
func (*CacheConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8d6a0b399e44d47, []int{0}
}

func (m *CacheConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacheConfig.Unmarshal(m, b)
}
func (m *CacheConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacheConfig.Marshal(b, m, deterministic)
}
func (m *CacheConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheConfig.Merge(m, src)
}
func (m *CacheConfig) XXX_Size() int {
	return xxx_messageInfo_CacheConfig.Size(m)
}
func (m *CacheConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CacheConfig proto.InternalMessageInfo

func (m *CacheConfig) GetTypedConfig() *any.Any {
	if m != nil {
		return m.TypedConfig
	}
	return nil
}

func (m *CacheConfig) GetAllowedVaryHeaders() []*matcher.StringMatcher {
	if m != nil {
		return m.AllowedVaryHeaders
	}
	return nil
}

func (m *CacheConfig) GetKeyCreatorParams() *CacheConfig_KeyCreatorParams {
	if m != nil {
		return m.KeyCreatorParams
	}
	return nil
}

func (m *CacheConfig) GetMaxBodyBytes() uint32 {
	if m != nil {
		return m.MaxBodyBytes
	}
	return 0
}

type CacheConfig_KeyCreatorParams struct {
	ExcludeScheme           bool                           `protobuf:"varint,1,opt,name=exclude_scheme,json=excludeScheme,proto3" json:"exclude_scheme,omitempty"`
	ExcludeHost             bool                           `protobuf:"varint,2,opt,name=exclude_host,json=excludeHost,proto3" json:"exclude_host,omitempty"`
	QueryParametersIncluded []*route.QueryParameterMatcher `protobuf:"bytes,3,rep,name=query_parameters_included,json=queryParametersIncluded,proto3" json:"query_parameters_included,omitempty"`
	QueryParametersExcluded []*route.QueryParameterMatcher `protobuf:"bytes,4,rep,name=query_parameters_excluded,json=queryParametersExcluded,proto3" json:"query_parameters_excluded,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                       `json:"-"`
	XXX_unrecognized        []byte                         `json:"-"`
	XXX_sizecache           int32                          `json:"-"`
}

func (m *CacheConfig_KeyCreatorParams) Reset()         { *m = CacheConfig_KeyCreatorParams{} }
func (m *CacheConfig_KeyCreatorParams) String() string { return proto.CompactTextString(m) }
func (*CacheConfig_KeyCreatorParams) ProtoMessage()    {}
func (*CacheConfig_KeyCreatorParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8d6a0b399e44d47, []int{0, 0}
}

func (m *CacheConfig_KeyCreatorParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacheConfig_KeyCreatorParams.Unmarshal(m, b)
}
func (m *CacheConfig_KeyCreatorParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacheConfig_KeyCreatorParams.Marshal(b, m, deterministic)
}
func (m *CacheConfig_KeyCreatorParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheConfig_KeyCreatorParams.Merge(m, src)
}
func (m *CacheConfig_KeyCreatorParams) XXX_Size() int {
	return xxx_messageInfo_CacheConfig_KeyCreatorParams.Size(m)
}
func (m *CacheConfig_KeyCreatorParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheConfig_KeyCreatorParams.DiscardUnknown(m)
}

var xxx_messageInfo_CacheConfig_KeyCreatorParams proto.InternalMessageInfo

func (m *CacheConfig_KeyCreatorParams) GetExcludeScheme() bool {
	if m != nil {
		return m.ExcludeScheme
	}
	return false
}

func (m *CacheConfig_KeyCreatorParams) GetExcludeHost() bool {
	if m != nil {
		return m.ExcludeHost
	}
	return false
}

func (m *CacheConfig_KeyCreatorParams) GetQueryParametersIncluded() []*route.QueryParameterMatcher {
	if m != nil {
		return m.QueryParametersIncluded
	}
	return nil
}

func (m *CacheConfig_KeyCreatorParams) GetQueryParametersExcluded() []*route.QueryParameterMatcher {
	if m != nil {
		return m.QueryParametersExcluded
	}
	return nil
}

func init() {
	proto.RegisterType((*CacheConfig)(nil), "envoy.config.filter.http.cache.v2alpha.CacheConfig")
	proto.RegisterType((*CacheConfig_KeyCreatorParams)(nil), "envoy.config.filter.http.cache.v2alpha.CacheConfig.KeyCreatorParams")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/cache/v2alpha/cache.proto", fileDescriptor_f8d6a0b399e44d47)
}

var fileDescriptor_f8d6a0b399e44d47 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xd3, 0x1f, 0x45, 0x9b, 0xb6, 0x8a, 0xac, 0x4a, 0x75, 0x23, 0x01, 0x29, 0x02, 0x94,
	0x0a, 0xb1, 0x2b, 0xb9, 0xbc, 0x40, 0x1d, 0x90, 0x8a, 0x10, 0x52, 0x48, 0x24, 0xae, 0xd6, 0xc6,
	0x9e, 0xc4, 0x56, 0xed, 0x5d, 0x77, 0x77, 0x6d, 0xb2, 0x37, 0x6e, 0x1c, 0xb9, 0x72, 0xe6, 0xc8,
	0x23, 0xf0, 0x04, 0x5c, 0x79, 0x0f, 0x4e, 0x1c, 0x39, 0x20, 0xe4, 0xdd, 0xb5, 0x44, 0x0b, 0x87,
	0x4a, 0x5c, 0x12, 0xef, 0xcc, 0x7c, 0xdf, 0x7c, 0x33, 0xdf, 0xa0, 0x10, 0x58, 0xc3, 0x35, 0x49,
	0x38, 0x5b, 0xe5, 0x6b, 0xb2, 0xca, 0x0b, 0x05, 0x82, 0x64, 0x4a, 0x55, 0x24, 0xa1, 0x49, 0x06,
	0xa4, 0x09, 0x69, 0x51, 0x65, 0xd4, 0xbe, 0x70, 0x25, 0xb8, 0xe2, 0xfe, 0x23, 0x83, 0xc1, 0x16,
	0x83, 0x2d, 0x06, 0xb7, 0x18, 0x6c, 0xab, 0x1c, 0x66, 0x74, 0x6a, 0xb9, 0x69, 0x95, 0x93, 0x26,
	0x24, 0x82, 0xd7, 0x0a, 0xec, 0x6f, 0x9c, 0xf0, 0xb2, 0xe2, 0x0c, 0x98, 0x92, 0x96, 0x72, 0x74,
	0xcf, 0x96, 0x2a, 0x5d, 0x01, 0x29, 0xa9, 0x4a, 0x32, 0x10, 0x44, 0x2a, 0x91, 0xb3, 0xb5, 0x2b,
	0x38, 0x5e, 0x73, 0xbe, 0x2e, 0x80, 0x98, 0xd7, 0xb2, 0x5e, 0x11, 0xca, 0xb4, 0x4b, 0xdd, 0xad,
	0xd3, 0x8a, 0x12, 0xca, 0x18, 0x57, 0x54, 0xe5, 0x9c, 0x49, 0x52, 0xe6, 0x6b, 0x41, 0x95, 0x93,
	0x3b, 0xba, 0xf3, 0x57, 0x5e, 0x2a, 0xaa, 0xea, 0xae, 0xf5, 0x51, 0x43, 0x8b, 0x3c, 0xa5, 0x0a,
	0x48, 0xf7, 0x61, 0x13, 0xf7, 0xbf, 0x6f, 0xa3, 0xc1, 0xb4, 0x1d, 0x68, 0x6a, 0x06, 0xf5, 0xcf,
	0xd1, 0x5e, 0xab, 0x2f, 0x8d, 0xed, 0xe0, 0x81, 0x37, 0xf6, 0x26, 0x83, 0xf0, 0x10, 0x5b, 0x65,
	0xb8, 0x53, 0x86, 0xcf, 0x99, 0x8e, 0xfa, 0x3f, 0xa3, 0x9d, 0x4f, 0x5e, 0xaf, 0xef, 0xcd, 0x07,
	0x06, 0xe3, 0x28, 0x16, 0xe8, 0x90, 0x16, 0x05, 0x7f, 0x0b, 0x69, 0xdc, 0x50, 0xa1, 0xe3, 0x0c,
	0x68, 0x0a, 0x42, 0x06, 0xbd, 0xf1, 0xd6, 0x64, 0x10, 0x9e, 0x60, 0xbb, 0xd8, 0x16, 0x81, 0xdd,
	0x16, 0xf0, 0xc2, 0x6c, 0xe1, 0x95, 0x7d, 0xcd, 0x7d, 0x07, 0x7f, 0x43, 0x85, 0xbe, 0xb0, 0x60,
	0x5f, 0x20, 0xff, 0x12, 0x74, 0x9c, 0x08, 0xa0, 0x8a, 0x8b, 0xb8, 0xa2, 0x82, 0x96, 0x32, 0xd8,
	0x32, 0xea, 0x9e, 0xe1, 0xdb, 0x79, 0x85, 0xff, 0x18, 0x14, 0xbf, 0x04, 0x3d, 0xb5, 0x64, 0x33,
	0xc3, 0x35, 0x1f, 0x5e, 0xde, 0x88, 0xf8, 0x0f, 0xd0, 0x41, 0x49, 0x37, 0xf1, 0x92, 0xa7, 0x3a,
	0x5e, 0x6a, 0x05, 0x32, 0xd8, 0x1e, 0x7b, 0x93, 0xfd, 0xf9, 0x5e, 0x49, 0x37, 0x11, 0x4f, 0x75,
	0xd4, 0xc6, 0x46, 0x9f, 0x7b, 0x68, 0x78, 0x93, 0xcc, 0x7f, 0x88, 0x0e, 0x60, 0x93, 0x14, 0x75,
	0x0a, 0xb1, 0x4c, 0x32, 0x28, 0xc1, 0x2c, 0xb2, 0x3f, 0xdf, 0x77, 0xd1, 0x85, 0x09, 0xfa, 0x27,
	0x68, 0xaf, 0x2b, 0xcb, 0xb8, 0x54, 0x41, 0xcf, 0x14, 0x0d, 0x5c, 0xec, 0x82, 0x4b, 0xe5, 0x03,
	0x3a, 0xbe, 0xaa, 0x41, 0x68, 0x3b, 0x32, 0x28, 0x10, 0x32, 0xce, 0x99, 0xc9, 0xa7, 0xc1, 0x96,
	0x59, 0xe9, 0xa9, 0x9b, 0x9f, 0x56, 0x39, 0x6e, 0x42, 0x6c, 0xae, 0x0f, 0xbf, 0x6e, 0x41, 0xb3,
	0x0e, 0xd3, 0xad, 0xf6, 0xe8, 0xea, 0x5a, 0x58, 0xbe, 0x70, 0x4c, 0xff, 0x6c, 0xe3, 0x64, 0xa4,
	0xc1, 0xf6, 0xff, 0xb6, 0x79, 0xee, 0x98, 0xa2, 0xf7, 0xde, 0x8f, 0x8f, 0xbf, 0x3e, 0xec, 0x3c,
	0xf1, 0x1f, 0x5b, 0x2e, 0xd8, 0x28, 0x60, 0xb2, 0xbd, 0x57, 0x67, 0x9b, 0xbc, 0xe6, 0xdb, 0x99,
	0xf1, 0xed, 0xcb, 0xbb, 0xaf, 0xdf, 0x76, 0x7b, 0x7d, 0xcf, 0xfe, 0x0f, 0x3d, 0xf4, 0x34, 0xe7,
	0x56, 0x4b, 0x25, 0xf8, 0x46, 0xdf, 0xd2, 0xfd, 0x08, 0x19, 0xfb, 0x67, 0xed, 0x05, 0xcf, 0xbc,
	0xe5, 0xae, 0x39, 0xe5, 0xb3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xa0, 0x89, 0x9b, 0x1c,
	0x04, 0x00, 0x00,
}
