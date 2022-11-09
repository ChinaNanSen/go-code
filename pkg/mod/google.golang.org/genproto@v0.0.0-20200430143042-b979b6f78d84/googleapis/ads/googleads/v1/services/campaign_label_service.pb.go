// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_label_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [CampaignLabelService.GetCampaignLabel][google.ads.googleads.v1.services.CampaignLabelService.GetCampaignLabel].
type GetCampaignLabelRequest struct {
	// Required. The resource name of the campaign-label relationship to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignLabelRequest) Reset()         { *m = GetCampaignLabelRequest{} }
func (m *GetCampaignLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignLabelRequest) ProtoMessage()    {}
func (*GetCampaignLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27173f6ba18f2b11, []int{0}
}

func (m *GetCampaignLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignLabelRequest.Unmarshal(m, b)
}
func (m *GetCampaignLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignLabelRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignLabelRequest.Merge(m, src)
}
func (m *GetCampaignLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignLabelRequest.Size(m)
}
func (m *GetCampaignLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignLabelRequest proto.InternalMessageInfo

func (m *GetCampaignLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignLabelService.MutateCampaignLabels][google.ads.googleads.v1.services.CampaignLabelService.MutateCampaignLabels].
type MutateCampaignLabelsRequest struct {
	// Required. ID of the customer whose campaign-label relationships are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on campaign-label relationships.
	Operations []*CampaignLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignLabelsRequest) Reset()         { *m = MutateCampaignLabelsRequest{} }
func (m *MutateCampaignLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignLabelsRequest) ProtoMessage()    {}
func (*MutateCampaignLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27173f6ba18f2b11, []int{1}
}

func (m *MutateCampaignLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignLabelsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignLabelsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignLabelsRequest.Merge(m, src)
}
func (m *MutateCampaignLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignLabelsRequest.Size(m)
}
func (m *MutateCampaignLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignLabelsRequest proto.InternalMessageInfo

func (m *MutateCampaignLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignLabelsRequest) GetOperations() []*CampaignLabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on a campaign-label relationship.
type CampaignLabelOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignLabelOperation_Create
	//	*CampaignLabelOperation_Remove
	Operation            isCampaignLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CampaignLabelOperation) Reset()         { *m = CampaignLabelOperation{} }
func (m *CampaignLabelOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignLabelOperation) ProtoMessage()    {}
func (*CampaignLabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_27173f6ba18f2b11, []int{2}
}

func (m *CampaignLabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignLabelOperation.Unmarshal(m, b)
}
func (m *CampaignLabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignLabelOperation.Marshal(b, m, deterministic)
}
func (m *CampaignLabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignLabelOperation.Merge(m, src)
}
func (m *CampaignLabelOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignLabelOperation.Size(m)
}
func (m *CampaignLabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignLabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignLabelOperation proto.InternalMessageInfo

type isCampaignLabelOperation_Operation interface {
	isCampaignLabelOperation_Operation()
}

type CampaignLabelOperation_Create struct {
	Create *resources.CampaignLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignLabelOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CampaignLabelOperation_Create) isCampaignLabelOperation_Operation() {}

func (*CampaignLabelOperation_Remove) isCampaignLabelOperation_Operation() {}

func (m *CampaignLabelOperation) GetOperation() isCampaignLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignLabelOperation) GetCreate() *resources.CampaignLabel {
	if x, ok := m.GetOperation().(*CampaignLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignLabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignLabelOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignLabelOperation_Create)(nil),
		(*CampaignLabelOperation_Remove)(nil),
	}
}

// Response message for a campaign labels mutate.
type MutateCampaignLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateCampaignLabelsResponse) Reset()         { *m = MutateCampaignLabelsResponse{} }
func (m *MutateCampaignLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignLabelsResponse) ProtoMessage()    {}
func (*MutateCampaignLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27173f6ba18f2b11, []int{3}
}

func (m *MutateCampaignLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignLabelsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignLabelsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignLabelsResponse.Merge(m, src)
}
func (m *MutateCampaignLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignLabelsResponse.Size(m)
}
func (m *MutateCampaignLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignLabelsResponse proto.InternalMessageInfo

func (m *MutateCampaignLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignLabelsResponse) GetResults() []*MutateCampaignLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for a campaign label mutate.
type MutateCampaignLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignLabelResult) Reset()         { *m = MutateCampaignLabelResult{} }
func (m *MutateCampaignLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignLabelResult) ProtoMessage()    {}
func (*MutateCampaignLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_27173f6ba18f2b11, []int{4}
}

func (m *MutateCampaignLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignLabelResult.Unmarshal(m, b)
}
func (m *MutateCampaignLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignLabelResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignLabelResult.Merge(m, src)
}
func (m *MutateCampaignLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignLabelResult.Size(m)
}
func (m *MutateCampaignLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignLabelResult proto.InternalMessageInfo

func (m *MutateCampaignLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignLabelRequest)(nil), "google.ads.googleads.v1.services.GetCampaignLabelRequest")
	proto.RegisterType((*MutateCampaignLabelsRequest)(nil), "google.ads.googleads.v1.services.MutateCampaignLabelsRequest")
	proto.RegisterType((*CampaignLabelOperation)(nil), "google.ads.googleads.v1.services.CampaignLabelOperation")
	proto.RegisterType((*MutateCampaignLabelsResponse)(nil), "google.ads.googleads.v1.services.MutateCampaignLabelsResponse")
	proto.RegisterType((*MutateCampaignLabelResult)(nil), "google.ads.googleads.v1.services.MutateCampaignLabelResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_label_service.proto", fileDescriptor_27173f6ba18f2b11)
}

var fileDescriptor_27173f6ba18f2b11 = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xbe, 0x76, 0xae, 0x7a, 0x6f, 0x27, 0xed, 0xbd, 0x68, 0x28, 0x6d, 0x9a, 0x56, 0x22, 0x32,
	0x15, 0x44, 0x11, 0xb2, 0x9b, 0x54, 0x42, 0xc5, 0x55, 0x10, 0x0e, 0xa2, 0x2d, 0x08, 0x68, 0x95,
	0x8a, 0x22, 0xa1, 0x20, 0x6b, 0x6a, 0x4f, 0x8d, 0x25, 0xdb, 0x63, 0x66, 0xc6, 0x91, 0xaa, 0xaa,
	0x12, 0x62, 0xcb, 0x92, 0x37, 0x60, 0xc9, 0x3b, 0xf0, 0x00, 0x74, 0xcb, 0xae, 0xab, 0x2e, 0x58,
	0xb1, 0x40, 0x5d, 0xb3, 0x42, 0xfe, 0x99, 0x24, 0x0e, 0x89, 0x22, 0xba, 0x3b, 0x3e, 0xe7, 0x3b,
	0xdf, 0xf9, 0x1f, 0x83, 0xa6, 0x43, 0x88, 0xe3, 0x61, 0x0d, 0xd9, 0x4c, 0x4b, 0xc5, 0x58, 0xea,
	0xd6, 0x35, 0x86, 0x69, 0xd7, 0xb5, 0x30, 0xd3, 0x2c, 0xe4, 0x87, 0xc8, 0x75, 0x02, 0xd3, 0x43,
	0x07, 0xd8, 0x33, 0x33, 0xbd, 0x1a, 0x52, 0xc2, 0x09, 0xac, 0xa4, 0x3e, 0x2a, 0xb2, 0x99, 0xda,
	0x73, 0x57, 0xbb, 0x75, 0x55, 0xb8, 0x97, 0xef, 0x8c, 0x0b, 0x40, 0x31, 0x23, 0x11, 0xfd, 0x3d,
	0x42, 0xca, 0x5c, 0x5e, 0x16, 0x7e, 0xa1, 0xab, 0xa1, 0x20, 0x20, 0x1c, 0x71, 0x97, 0x04, 0x2c,
	0xb3, 0x2e, 0x0c, 0x58, 0x2d, 0xcf, 0xc5, 0x01, 0xcf, 0x0c, 0xd7, 0x07, 0x0c, 0x87, 0x2e, 0xf6,
	0x6c, 0xf3, 0x00, 0xbf, 0x46, 0x5d, 0x97, 0xd0, 0x0c, 0xb0, 0x38, 0x00, 0x10, 0x29, 0x0c, 0x91,
	0xd2, 0xd0, 0xd2, 0x18, 0x47, 0x3c, 0xca, 0xa2, 0x29, 0x01, 0x58, 0xd8, 0xc2, 0xfc, 0x41, 0x96,
	0xe6, 0x93, 0x38, 0xcb, 0x36, 0x7e, 0x13, 0x61, 0xc6, 0xe1, 0x1e, 0x98, 0x15, 0x2c, 0x66, 0x80,
	0x7c, 0x5c, 0x92, 0x2a, 0x52, 0x75, 0xba, 0xa5, 0x9e, 0x1b, 0xf2, 0x4f, 0xa3, 0x0a, 0x6e, 0xf6,
	0x9b, 0x92, 0x49, 0xa1, 0xcb, 0x54, 0x8b, 0xf8, 0x5a, 0x9e, 0x6d, 0x46, 0x90, 0x3c, 0x43, 0x3e,
	0x56, 0x2e, 0x24, 0xb0, 0xf4, 0x34, 0xe2, 0x88, 0xe3, 0x1c, 0x8a, 0x89, 0xa0, 0x2b, 0xa0, 0x68,
	0x45, 0x8c, 0x13, 0x1f, 0x53, 0xd3, 0xb5, 0xb3, 0x90, 0x85, 0x73, 0x43, 0x6e, 0x03, 0xa1, 0x7f,
	0x64, 0xc3, 0x57, 0x00, 0x90, 0x10, 0xd3, 0xb4, 0x6f, 0x25, 0xb9, 0x52, 0xa8, 0x16, 0x1b, 0xeb,
	0xea, 0xa4, 0x81, 0xa9, 0xb9, 0x90, 0x3b, 0x82, 0x20, 0xa3, 0xef, 0x13, 0xc2, 0x5b, 0xe0, 0xff,
	0x10, 0x51, 0xee, 0x22, 0xcf, 0x3c, 0x44, 0xae, 0x17, 0x51, 0x5c, 0x2a, 0x54, 0xa4, 0xea, 0xbf,
	0xed, 0xff, 0x32, 0xf5, 0x66, 0xaa, 0x85, 0x37, 0xc0, 0x6c, 0x17, 0x79, 0xae, 0x8d, 0x38, 0x36,
	0x49, 0xe0, 0x1d, 0x95, 0xfe, 0x4e, 0x60, 0x33, 0x42, 0xb9, 0x13, 0x78, 0x47, 0xca, 0x7b, 0x09,
	0xcc, 0x8f, 0x8e, 0x0c, 0x1f, 0x83, 0x29, 0x8b, 0x62, 0xc4, 0xd3, 0xde, 0x16, 0x1b, 0xab, 0x63,
	0x6b, 0xe8, 0xad, 0x54, 0xbe, 0x88, 0xed, 0xbf, 0xda, 0x19, 0x03, 0x2c, 0x81, 0x29, 0x8a, 0x7d,
	0xd2, 0xc5, 0x25, 0x39, 0x6e, 0x5a, 0x6c, 0x49, 0xbf, 0x5b, 0x45, 0x30, 0xdd, 0x2b, 0x4e, 0xf9,
	0x2c, 0x81, 0xe5, 0xd1, 0x03, 0x60, 0x21, 0x09, 0x18, 0x86, 0x9b, 0xe0, 0xda, 0x50, 0xf1, 0x26,
	0xa6, 0x94, 0xd0, 0xa4, 0x05, 0xc5, 0x06, 0x14, 0x29, 0xd2, 0xd0, 0x52, 0xf7, 0x92, 0x55, 0x6a,
	0x5f, 0xcd, 0xb7, 0xe5, 0x61, 0x0c, 0x87, 0xcf, 0xc1, 0x3f, 0x14, 0xb3, 0xc8, 0xe3, 0x62, 0x40,
	0x1b, 0x93, 0x07, 0x34, 0x22, 0xb1, 0x76, 0xc2, 0xd1, 0x16, 0x5c, 0xca, 0x7d, 0xb0, 0x38, 0x16,
	0x15, 0xcf, 0x63, 0xc4, 0xca, 0xe6, 0x57, 0xb0, 0x71, 0x51, 0x00, 0x73, 0x39, 0xe7, 0xbd, 0x34,
	0x3c, 0xfc, 0x22, 0x81, 0x2b, 0xc3, 0xc7, 0x00, 0xef, 0x4e, 0xce, 0x7a, 0xcc, 0x01, 0x95, 0xff,
	0x78, 0x9a, 0xca, 0xf6, 0x99, 0x91, 0x2f, 0xe0, 0xdd, 0xd7, 0x6f, 0x1f, 0xe4, 0x06, 0x5c, 0x8d,
	0x5f, 0x95, 0xe3, 0x9c, 0xa5, 0x29, 0xce, 0x81, 0x69, 0xb5, 0xde, 0x33, 0x93, 0x8e, 0x52, 0xab,
	0x9d, 0xc0, 0x1f, 0x12, 0x98, 0x1b, 0x35, 0x66, 0xd8, 0xbc, 0xd4, 0x14, 0xc4, 0x7d, 0x96, 0xef,
	0x5d, 0xd6, 0x3d, 0xdd, 0x2e, 0xe5, 0xc5, 0x99, 0x31, 0x3f, 0x70, 0xe0, 0xb7, 0xfb, 0x57, 0x97,
	0x94, 0xba, 0xae, 0xac, 0xc5, 0xa5, 0xf6, 0x6b, 0x3b, 0x1e, 0x00, 0x37, 0x6b, 0x27, 0x43, 0x95,
	0xea, 0x7e, 0x12, 0x4b, 0x97, 0x6a, 0xe5, 0xa5, 0x53, 0xa3, 0x34, 0xee, 0x45, 0x6a, 0xbd, 0x95,
	0xc1, 0x8a, 0x45, 0xfc, 0x89, 0xb9, 0xb7, 0x16, 0x47, 0x2d, 0xc6, 0x6e, 0xfc, 0x52, 0xee, 0x4a,
	0x2f, 0xb7, 0x33, 0x77, 0x87, 0x78, 0x28, 0x70, 0x54, 0x42, 0x1d, 0xcd, 0xc1, 0x41, 0xf2, 0x8e,
	0x6a, 0xfd, 0x80, 0xe3, 0xff, 0x37, 0x1b, 0x42, 0xf8, 0x28, 0x17, 0xb6, 0x0c, 0xe3, 0x93, 0x5c,
	0xd9, 0x4a, 0x09, 0x0d, 0x9b, 0xa9, 0xa9, 0x18, 0x4b, 0xfb, 0x75, 0x35, 0x0b, 0xcc, 0x4e, 0x05,
	0xa4, 0x63, 0xd8, 0xac, 0xd3, 0x83, 0x74, 0xf6, 0xeb, 0x1d, 0x01, 0xf9, 0x2e, 0xaf, 0xa4, 0x7a,
	0x5d, 0x37, 0x6c, 0xa6, 0xeb, 0x3d, 0x90, 0xae, 0xef, 0xd7, 0x75, 0x5d, 0xc0, 0x0e, 0xa6, 0x92,
	0x3c, 0xd7, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xdd, 0xbe, 0x7f, 0x16, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignLabelServiceClient is the client API for CampaignLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignLabelServiceClient interface {
	// Returns the requested campaign-label relationship in full detail.
	GetCampaignLabel(ctx context.Context, in *GetCampaignLabelRequest, opts ...grpc.CallOption) (*resources.CampaignLabel, error)
	// Creates and removes campaign-label relationships.
	// Operation statuses are returned.
	MutateCampaignLabels(ctx context.Context, in *MutateCampaignLabelsRequest, opts ...grpc.CallOption) (*MutateCampaignLabelsResponse, error)
}

type campaignLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignLabelServiceClient(cc grpc.ClientConnInterface) CampaignLabelServiceClient {
	return &campaignLabelServiceClient{cc}
}

func (c *campaignLabelServiceClient) GetCampaignLabel(ctx context.Context, in *GetCampaignLabelRequest, opts ...grpc.CallOption) (*resources.CampaignLabel, error) {
	out := new(resources.CampaignLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignLabelService/GetCampaignLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignLabelServiceClient) MutateCampaignLabels(ctx context.Context, in *MutateCampaignLabelsRequest, opts ...grpc.CallOption) (*MutateCampaignLabelsResponse, error) {
	out := new(MutateCampaignLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignLabelService/MutateCampaignLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignLabelServiceServer is the server API for CampaignLabelService service.
type CampaignLabelServiceServer interface {
	// Returns the requested campaign-label relationship in full detail.
	GetCampaignLabel(context.Context, *GetCampaignLabelRequest) (*resources.CampaignLabel, error)
	// Creates and removes campaign-label relationships.
	// Operation statuses are returned.
	MutateCampaignLabels(context.Context, *MutateCampaignLabelsRequest) (*MutateCampaignLabelsResponse, error)
}

// UnimplementedCampaignLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignLabelServiceServer struct {
}

func (*UnimplementedCampaignLabelServiceServer) GetCampaignLabel(ctx context.Context, req *GetCampaignLabelRequest) (*resources.CampaignLabel, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCampaignLabel not implemented")
}
func (*UnimplementedCampaignLabelServiceServer) MutateCampaignLabels(ctx context.Context, req *MutateCampaignLabelsRequest) (*MutateCampaignLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignLabels not implemented")
}

func RegisterCampaignLabelServiceServer(s *grpc.Server, srv CampaignLabelServiceServer) {
	s.RegisterService(&_CampaignLabelService_serviceDesc, srv)
}

func _CampaignLabelService_GetCampaignLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignLabelServiceServer).GetCampaignLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignLabelService/GetCampaignLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignLabelServiceServer).GetCampaignLabel(ctx, req.(*GetCampaignLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignLabelService_MutateCampaignLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignLabelServiceServer).MutateCampaignLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignLabelService/MutateCampaignLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignLabelServiceServer).MutateCampaignLabels(ctx, req.(*MutateCampaignLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignLabelService",
	HandlerType: (*CampaignLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignLabel",
			Handler:    _CampaignLabelService_GetCampaignLabel_Handler,
		},
		{
			MethodName: "MutateCampaignLabels",
			Handler:    _CampaignLabelService_MutateCampaignLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_label_service.proto",
}