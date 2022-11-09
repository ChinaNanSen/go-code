// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/dates.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A date range.
type DateRange struct {
	// The start date, in yyyy-mm-dd format. This date is inclusive.
	StartDate *wrappers.StringValue `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// The end date, in yyyy-mm-dd format. This date is inclusive.
	EndDate              *wrappers.StringValue `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DateRange) Reset()         { *m = DateRange{} }
func (m *DateRange) String() string { return proto.CompactTextString(m) }
func (*DateRange) ProtoMessage()    {}
func (*DateRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_141c54e845fb6763, []int{0}
}

func (m *DateRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateRange.Unmarshal(m, b)
}
func (m *DateRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateRange.Marshal(b, m, deterministic)
}
func (m *DateRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateRange.Merge(m, src)
}
func (m *DateRange) XXX_Size() int {
	return xxx_messageInfo_DateRange.Size(m)
}
func (m *DateRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DateRange.DiscardUnknown(m)
}

var xxx_messageInfo_DateRange proto.InternalMessageInfo

func (m *DateRange) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *DateRange) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func init() {
	proto.RegisterType((*DateRange)(nil), "google.ads.googleads.v1.common.DateRange")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/dates.proto", fileDescriptor_141c54e845fb6763)
}

var fileDescriptor_141c54e845fb6763 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0x05, 0x75, 0xf1, 0xd6, 0x93, 0x8c, 0x31, 0xa4, 0x27, 0xf1, 0x90, 0x50, 0x3d,
	0x08, 0xd9, 0xa9, 0x73, 0xb0, 0xeb, 0x98, 0xd0, 0x83, 0x14, 0xe4, 0x6d, 0x89, 0xa1, 0xd0, 0xe6,
	0x95, 0x24, 0x9b, 0x67, 0xbf, 0x8a, 0x47, 0x3f, 0x8a, 0xdf, 0xc3, 0x8b, 0x9f, 0x42, 0xd2, 0xb4,
	0xbd, 0x29, 0x9e, 0xfa, 0xa7, 0xef, 0xf7, 0xff, 0xff, 0x5f, 0x1e, 0xb9, 0x51, 0x88, 0xaa, 0x96,
	0x0c, 0x84, 0x65, 0x41, 0x7a, 0x75, 0xcc, 0xd8, 0x1e, 0x9b, 0x06, 0x35, 0x13, 0xe0, 0xa4, 0xa5,
	0xad, 0x41, 0x87, 0xc9, 0x3c, 0x00, 0x14, 0x84, 0xa5, 0x23, 0x4b, 0x8f, 0x19, 0x0d, 0xec, 0xb4,
	0x9f, 0xb3, 0x8e, 0xde, 0x1d, 0x5e, 0xd8, 0xab, 0x81, 0xb6, 0x95, 0xa6, 0xf7, 0x4f, 0x67, 0x43,
	0x57, 0x5b, 0x31, 0xd0, 0x1a, 0x1d, 0xb8, 0x0a, 0x75, 0x3f, 0x4d, 0xdf, 0x22, 0x32, 0x59, 0x81,
	0x93, 0x5b, 0xd0, 0x4a, 0x26, 0x0b, 0x42, 0xac, 0x03, 0xe3, 0x9e, 0xfd, 0x02, 0x97, 0xd1, 0x55,
	0x74, 0x7d, 0x71, 0x3b, 0xeb, 0x5b, 0xe9, 0x50, 0x40, 0x1f, 0x9d, 0xa9, 0xb4, 0x2a, 0xa0, 0x3e,
	0xc8, 0xed, 0xa4, 0xe3, 0x7d, 0x42, 0x72, 0x4f, 0xce, 0xa5, 0x16, 0xc1, 0x1a, 0xff, 0xc3, 0x7a,
	0x26, 0xb5, 0xf0, 0xc6, 0xe5, 0x57, 0x44, 0xd2, 0x3d, 0x36, 0xf4, 0xef, 0x87, 0x2e, 0x89, 0x87,
	0xed, 0xc6, 0x47, 0x6d, 0xa2, 0xa7, 0x55, 0x4f, 0x2b, 0xac, 0x41, 0x2b, 0x8a, 0x46, 0x31, 0x25,
	0x75, 0x57, 0x34, 0x9c, 0xb4, 0xad, 0xec, 0x6f, 0x17, 0x5e, 0x84, 0xcf, 0x7b, 0x7c, 0xb2, 0xce,
	0xf3, 0x8f, 0x78, 0xbe, 0x0e, 0x61, 0xb9, 0xb0, 0x34, 0x48, 0xaf, 0x8a, 0x8c, 0x3e, 0x74, 0xd8,
	0xe7, 0x00, 0x94, 0xb9, 0xb0, 0xe5, 0x08, 0x94, 0x45, 0x56, 0x06, 0xe0, 0x3b, 0x4e, 0xc3, 0x5f,
	0xce, 0x73, 0x61, 0x39, 0x1f, 0x11, 0xce, 0x8b, 0x8c, 0xf3, 0x00, 0xed, 0x4e, 0xbb, 0xed, 0xee,
	0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x8e, 0xf7, 0xb8, 0xfe, 0x01, 0x00, 0x00,
}