// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/keyword_plan_network.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enumerates keyword plan forecastable network types.
type KeywordPlanNetworkEnum_KeywordPlanNetwork int32

const (
	// Not specified.
	KeywordPlanNetworkEnum_UNSPECIFIED KeywordPlanNetworkEnum_KeywordPlanNetwork = 0
	// The value is unknown in this version.
	KeywordPlanNetworkEnum_UNKNOWN KeywordPlanNetworkEnum_KeywordPlanNetwork = 1
	// Google Search.
	KeywordPlanNetworkEnum_GOOGLE_SEARCH KeywordPlanNetworkEnum_KeywordPlanNetwork = 2
	// Google Search + Search partners.
	KeywordPlanNetworkEnum_GOOGLE_SEARCH_AND_PARTNERS KeywordPlanNetworkEnum_KeywordPlanNetwork = 3
)

var KeywordPlanNetworkEnum_KeywordPlanNetwork_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "GOOGLE_SEARCH",
	3: "GOOGLE_SEARCH_AND_PARTNERS",
}

var KeywordPlanNetworkEnum_KeywordPlanNetwork_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"GOOGLE_SEARCH":              2,
	"GOOGLE_SEARCH_AND_PARTNERS": 3,
}

func (x KeywordPlanNetworkEnum_KeywordPlanNetwork) String() string {
	return proto.EnumName(KeywordPlanNetworkEnum_KeywordPlanNetwork_name, int32(x))
}

func (KeywordPlanNetworkEnum_KeywordPlanNetwork) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_01afaeca21e51bd7, []int{0, 0}
}

// Container for enumeration of keyword plan forecastable network types.
type KeywordPlanNetworkEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanNetworkEnum) Reset()         { *m = KeywordPlanNetworkEnum{} }
func (m *KeywordPlanNetworkEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanNetworkEnum) ProtoMessage()    {}
func (*KeywordPlanNetworkEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_01afaeca21e51bd7, []int{0}
}

func (m *KeywordPlanNetworkEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanNetworkEnum.Unmarshal(m, b)
}
func (m *KeywordPlanNetworkEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanNetworkEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanNetworkEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanNetworkEnum.Merge(m, src)
}
func (m *KeywordPlanNetworkEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanNetworkEnum.Size(m)
}
func (m *KeywordPlanNetworkEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanNetworkEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanNetworkEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork", KeywordPlanNetworkEnum_KeywordPlanNetwork_name, KeywordPlanNetworkEnum_KeywordPlanNetwork_value)
	proto.RegisterType((*KeywordPlanNetworkEnum)(nil), "google.ads.googleads.v3.enums.KeywordPlanNetworkEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/keyword_plan_network.proto", fileDescriptor_01afaeca21e51bd7)
}

var fileDescriptor_01afaeca21e51bd7 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0x76, 0x1d, 0x28, 0x64, 0x88, 0xb5, 0x07, 0x85, 0xe1, 0x84, 0xed, 0x01, 0xd2, 0x43, 0x2f,
	0x12, 0x4f, 0xd9, 0x56, 0xeb, 0x98, 0x64, 0x65, 0x73, 0x13, 0xa4, 0x50, 0xa2, 0x09, 0x61, 0xac,
	0x4b, 0x4a, 0xd3, 0x6d, 0x78, 0xf2, 0x5d, 0x3c, 0xfa, 0x28, 0x3e, 0x8a, 0xf8, 0x10, 0xd2, 0xc4,
	0x0d, 0x64, 0xe8, 0x25, 0x7c, 0xe4, 0xfb, 0xc3, 0xf7, 0xfb, 0xc0, 0x95, 0x50, 0x4a, 0x64, 0xdc,
	0xa7, 0x4c, 0xfb, 0x16, 0x56, 0x68, 0x1d, 0xf8, 0x5c, 0xae, 0x96, 0xda, 0x5f, 0xf0, 0x97, 0x8d,
	0x2a, 0x58, 0x9a, 0x67, 0x54, 0xa6, 0x92, 0x97, 0x1b, 0x55, 0x2c, 0x60, 0x5e, 0xa8, 0x52, 0x79,
	0x2d, 0x2b, 0x87, 0x94, 0x69, 0xb8, 0x73, 0xc2, 0x75, 0x00, 0x8d, 0xb3, 0x79, 0xb1, 0x0d, 0xce,
	0xe7, 0x3e, 0x95, 0x52, 0x95, 0xb4, 0x9c, 0x2b, 0xa9, 0xad, 0xb9, 0xf3, 0x0a, 0xce, 0x86, 0x36,
	0x3a, 0xce, 0xa8, 0x24, 0x36, 0x38, 0x94, 0xab, 0x65, 0x87, 0x03, 0x6f, 0x9f, 0xf1, 0x4e, 0x40,
	0x63, 0x4a, 0x26, 0x71, 0xd8, 0x1b, 0xdc, 0x0c, 0xc2, 0xbe, 0x7b, 0xe0, 0x35, 0xc0, 0xd1, 0x94,
	0x0c, 0xc9, 0xe8, 0x81, 0xb8, 0x35, 0xef, 0x14, 0x1c, 0x47, 0xa3, 0x51, 0x74, 0x17, 0xa6, 0x93,
	0x10, 0x8f, 0x7b, 0xb7, 0xae, 0xe3, 0x5d, 0x82, 0xe6, 0xaf, 0xaf, 0x14, 0x93, 0x7e, 0x1a, 0xe3,
	0xf1, 0x3d, 0x09, 0xc7, 0x13, 0xb7, 0xde, 0xfd, 0xaa, 0x81, 0xf6, 0xb3, 0x5a, 0xc2, 0x7f, 0x8f,
	0xe8, 0x9e, 0xef, 0x57, 0x89, 0xab, 0xfe, 0x71, 0xed, 0xb1, 0xfb, 0xe3, 0x14, 0x2a, 0xa3, 0x52,
	0x40, 0x55, 0x08, 0x5f, 0x70, 0x69, 0xae, 0xdb, 0x0e, 0x99, 0xcf, 0xf5, 0x1f, 0xbb, 0x5e, 0x9b,
	0xf7, 0xcd, 0xa9, 0x47, 0x18, 0xbf, 0x3b, 0xad, 0xc8, 0x46, 0x61, 0xa6, 0xa1, 0x85, 0x15, 0x9a,
	0x05, 0xb0, 0x1a, 0x44, 0x7f, 0x6c, 0xf9, 0x04, 0x33, 0x9d, 0xec, 0xf8, 0x64, 0x16, 0x24, 0x86,
	0xff, 0x74, 0xda, 0xf6, 0x13, 0x21, 0xcc, 0x34, 0x42, 0x3b, 0x05, 0x42, 0xb3, 0x00, 0x21, 0xa3,
	0x79, 0x3a, 0x34, 0xc5, 0x82, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x75, 0x65, 0x21, 0xef,
	0x01, 0x00, 0x00,
}