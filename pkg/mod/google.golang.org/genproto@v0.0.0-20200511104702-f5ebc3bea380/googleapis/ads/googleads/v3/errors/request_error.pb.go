// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/request_error.proto

package errors

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

// Enum describing possible request errors.
type RequestErrorEnum_RequestError int32

const (
	// Enum unspecified.
	RequestErrorEnum_UNSPECIFIED RequestErrorEnum_RequestError = 0
	// The received error code is not known in this version.
	RequestErrorEnum_UNKNOWN RequestErrorEnum_RequestError = 1
	// Resource name is required for this request.
	RequestErrorEnum_RESOURCE_NAME_MISSING RequestErrorEnum_RequestError = 3
	// Resource name provided is malformed.
	RequestErrorEnum_RESOURCE_NAME_MALFORMED RequestErrorEnum_RequestError = 4
	// Resource name provided is malformed.
	RequestErrorEnum_BAD_RESOURCE_ID RequestErrorEnum_RequestError = 17
	// Customer ID is invalid.
	RequestErrorEnum_INVALID_CUSTOMER_ID RequestErrorEnum_RequestError = 16
	// Mutate operation should have either create, update, or remove specified.
	RequestErrorEnum_OPERATION_REQUIRED RequestErrorEnum_RequestError = 5
	// Requested resource not found.
	RequestErrorEnum_RESOURCE_NOT_FOUND RequestErrorEnum_RequestError = 6
	// Next page token specified in user request is invalid.
	RequestErrorEnum_INVALID_PAGE_TOKEN RequestErrorEnum_RequestError = 7
	// Next page token specified in user request has expired.
	RequestErrorEnum_EXPIRED_PAGE_TOKEN RequestErrorEnum_RequestError = 8
	// Page size specified in user request is invalid.
	RequestErrorEnum_INVALID_PAGE_SIZE RequestErrorEnum_RequestError = 22
	// Required field is missing.
	RequestErrorEnum_REQUIRED_FIELD_MISSING RequestErrorEnum_RequestError = 9
	// The field cannot be modified because it's immutable. It's also possible
	// that the field can be modified using 'create' operation but not 'update'.
	RequestErrorEnum_IMMUTABLE_FIELD RequestErrorEnum_RequestError = 11
	// Received too many entries in request.
	RequestErrorEnum_TOO_MANY_MUTATE_OPERATIONS RequestErrorEnum_RequestError = 13
	// Request cannot be executed by a manager account.
	RequestErrorEnum_CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT RequestErrorEnum_RequestError = 14
	// Mutate request was attempting to modify a readonly field.
	// For instance, Budget fields can be requested for Ad Group,
	// but are read-only for adGroups:mutate.
	RequestErrorEnum_CANNOT_MODIFY_FOREIGN_FIELD RequestErrorEnum_RequestError = 15
	// Enum value is not permitted.
	RequestErrorEnum_INVALID_ENUM_VALUE RequestErrorEnum_RequestError = 18
	// The developer-token parameter is required for all requests.
	RequestErrorEnum_DEVELOPER_TOKEN_PARAMETER_MISSING RequestErrorEnum_RequestError = 19
	// The login-customer-id parameter is required for this request.
	RequestErrorEnum_LOGIN_CUSTOMER_ID_PARAMETER_MISSING RequestErrorEnum_RequestError = 20
	// page_token is set in the validate only request
	RequestErrorEnum_VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN RequestErrorEnum_RequestError = 21
	// return_summary_row cannot be enabled if request did not select any
	// metrics field.
	RequestErrorEnum_CANNOT_RETURN_SUMMARY_ROW_FOR_REQUEST_WITHOUT_METRICS RequestErrorEnum_RequestError = 29
	// return_summary_row should not be enabled for validate only requests.
	RequestErrorEnum_CANNOT_RETURN_SUMMARY_ROW_FOR_VALIDATE_ONLY_REQUESTS RequestErrorEnum_RequestError = 30
	// return_summary_row parameter value should be the same between requests
	// with page_token field set and their original request.
	RequestErrorEnum_INCONSISTENT_RETURN_SUMMARY_ROW_VALUE RequestErrorEnum_RequestError = 31
	// The total results count cannot be returned if it was not requested in the
	// original request.
	RequestErrorEnum_TOTAL_RESULTS_COUNT_NOT_ORIGINALLY_REQUESTED RequestErrorEnum_RequestError = 32
)

var RequestErrorEnum_RequestError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "RESOURCE_NAME_MISSING",
	4:  "RESOURCE_NAME_MALFORMED",
	17: "BAD_RESOURCE_ID",
	16: "INVALID_CUSTOMER_ID",
	5:  "OPERATION_REQUIRED",
	6:  "RESOURCE_NOT_FOUND",
	7:  "INVALID_PAGE_TOKEN",
	8:  "EXPIRED_PAGE_TOKEN",
	22: "INVALID_PAGE_SIZE",
	9:  "REQUIRED_FIELD_MISSING",
	11: "IMMUTABLE_FIELD",
	13: "TOO_MANY_MUTATE_OPERATIONS",
	14: "CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT",
	15: "CANNOT_MODIFY_FOREIGN_FIELD",
	18: "INVALID_ENUM_VALUE",
	19: "DEVELOPER_TOKEN_PARAMETER_MISSING",
	20: "LOGIN_CUSTOMER_ID_PARAMETER_MISSING",
	21: "VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN",
	29: "CANNOT_RETURN_SUMMARY_ROW_FOR_REQUEST_WITHOUT_METRICS",
	30: "CANNOT_RETURN_SUMMARY_ROW_FOR_VALIDATE_ONLY_REQUESTS",
	31: "INCONSISTENT_RETURN_SUMMARY_ROW_VALUE",
	32: "TOTAL_RESULTS_COUNT_NOT_ORIGINALLY_REQUESTED",
}

var RequestErrorEnum_RequestError_value = map[string]int32{
	"UNSPECIFIED":                                           0,
	"UNKNOWN":                                               1,
	"RESOURCE_NAME_MISSING":                                 3,
	"RESOURCE_NAME_MALFORMED":                               4,
	"BAD_RESOURCE_ID":                                       17,
	"INVALID_CUSTOMER_ID":                                   16,
	"OPERATION_REQUIRED":                                    5,
	"RESOURCE_NOT_FOUND":                                    6,
	"INVALID_PAGE_TOKEN":                                    7,
	"EXPIRED_PAGE_TOKEN":                                    8,
	"INVALID_PAGE_SIZE":                                     22,
	"REQUIRED_FIELD_MISSING":                                9,
	"IMMUTABLE_FIELD":                                       11,
	"TOO_MANY_MUTATE_OPERATIONS":                            13,
	"CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT":                 14,
	"CANNOT_MODIFY_FOREIGN_FIELD":                           15,
	"INVALID_ENUM_VALUE":                                    18,
	"DEVELOPER_TOKEN_PARAMETER_MISSING":                     19,
	"LOGIN_CUSTOMER_ID_PARAMETER_MISSING":                   20,
	"VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN":                  21,
	"CANNOT_RETURN_SUMMARY_ROW_FOR_REQUEST_WITHOUT_METRICS": 29,
	"CANNOT_RETURN_SUMMARY_ROW_FOR_VALIDATE_ONLY_REQUESTS":  30,
	"INCONSISTENT_RETURN_SUMMARY_ROW_VALUE":                 31,
	"TOTAL_RESULTS_COUNT_NOT_ORIGINALLY_REQUESTED":          32,
}

func (x RequestErrorEnum_RequestError) String() string {
	return proto.EnumName(RequestErrorEnum_RequestError_name, int32(x))
}

func (RequestErrorEnum_RequestError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7bc15dc4a4456f34, []int{0, 0}
}

// Container for enum describing possible request errors.
type RequestErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestErrorEnum) Reset()         { *m = RequestErrorEnum{} }
func (m *RequestErrorEnum) String() string { return proto.CompactTextString(m) }
func (*RequestErrorEnum) ProtoMessage()    {}
func (*RequestErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bc15dc4a4456f34, []int{0}
}

func (m *RequestErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestErrorEnum.Unmarshal(m, b)
}
func (m *RequestErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestErrorEnum.Marshal(b, m, deterministic)
}
func (m *RequestErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestErrorEnum.Merge(m, src)
}
func (m *RequestErrorEnum) XXX_Size() int {
	return xxx_messageInfo_RequestErrorEnum.Size(m)
}
func (m *RequestErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_RequestErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.RequestErrorEnum_RequestError", RequestErrorEnum_RequestError_name, RequestErrorEnum_RequestError_value)
	proto.RegisterType((*RequestErrorEnum)(nil), "google.ads.googleads.v3.errors.RequestErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/request_error.proto", fileDescriptor_7bc15dc4a4456f34)
}

var fileDescriptor_7bc15dc4a4456f34 = []byte{
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xe1, 0x4e, 0x13, 0x41,
	0x10, 0x16, 0xd0, 0xa2, 0x8b, 0xca, 0xb2, 0x08, 0x44, 0xd0, 0xa2, 0x55, 0xa2, 0x26, 0xe6, 0x6a,
	0xac, 0x26, 0x7a, 0xfe, 0xda, 0xde, 0x4d, 0x8f, 0x0d, 0x77, 0xbb, 0x75, 0x6f, 0xb7, 0x50, 0xd2,
	0x64, 0x53, 0x6d, 0xd3, 0x90, 0x40, 0x0f, 0x7b, 0x85, 0x87, 0xf0, 0x25, 0x4c, 0xfc, 0xe9, 0xa3,
	0xf8, 0x28, 0xc6, 0x87, 0x30, 0x7b, 0xd7, 0x1e, 0xd5, 0x20, 0xbf, 0x6e, 0x6e, 0xe6, 0xfb, 0x66,
	0xbe, 0x99, 0xec, 0x87, 0x5e, 0x0f, 0x92, 0x64, 0x70, 0xdc, 0xaf, 0x76, 0x7b, 0x69, 0x35, 0x0f,
	0x6d, 0x74, 0x5e, 0xab, 0xf6, 0x47, 0xa3, 0x64, 0x94, 0x56, 0x47, 0xfd, 0x2f, 0x67, 0xfd, 0x74,
	0x6c, 0xb2, 0x5f, 0xe7, 0x74, 0x94, 0x8c, 0x13, 0x52, 0xce, 0x81, 0x4e, 0xb7, 0x97, 0x3a, 0x05,
	0xc7, 0x39, 0xaf, 0x39, 0x39, 0x67, 0xf3, 0xc1, 0xb4, 0xe7, 0xe9, 0x51, 0xb5, 0x3b, 0x1c, 0x26,
	0xe3, 0xee, 0xf8, 0x28, 0x19, 0xa6, 0x39, 0xbb, 0xf2, 0xad, 0x84, 0xb0, 0xcc, 0xbb, 0x82, 0xc5,
	0xc3, 0xf0, 0xec, 0xa4, 0xf2, 0xb5, 0x84, 0x6e, 0xcf, 0x26, 0xc9, 0x32, 0x5a, 0xd2, 0x3c, 0x6e,
	0x82, 0xc7, 0x1a, 0x0c, 0x7c, 0x7c, 0x8d, 0x2c, 0xa1, 0x45, 0xcd, 0xf7, 0xb8, 0xd8, 0xe7, 0x78,
	0x8e, 0xdc, 0x47, 0x6b, 0x12, 0x62, 0xa1, 0xa5, 0x07, 0x86, 0xd3, 0x08, 0x4c, 0xc4, 0xe2, 0x98,
	0xf1, 0x00, 0x2f, 0x90, 0x2d, 0xb4, 0xf1, 0x4f, 0x89, 0x86, 0x0d, 0x21, 0x23, 0xf0, 0xf1, 0x75,
	0xb2, 0x8a, 0x96, 0xeb, 0xd4, 0x37, 0x05, 0x80, 0xf9, 0x78, 0x85, 0x6c, 0xa0, 0x55, 0xc6, 0x5b,
	0x34, 0x64, 0xbe, 0xf1, 0x74, 0xac, 0x44, 0x04, 0xd2, 0x16, 0x30, 0x59, 0x47, 0x44, 0x34, 0x41,
	0x52, 0xc5, 0x04, 0x37, 0x12, 0x3e, 0x6a, 0x26, 0xc1, 0xc7, 0x37, 0x6c, 0xfe, 0x62, 0x84, 0x50,
	0xa6, 0x21, 0x34, 0xf7, 0x71, 0xc9, 0xe6, 0xa7, 0x8d, 0x9a, 0x34, 0x00, 0xa3, 0xc4, 0x1e, 0x70,
	0xbc, 0x68, 0xf3, 0x70, 0xd0, 0xb4, 0xe4, 0xd9, 0xfc, 0x4d, 0xb2, 0x86, 0x56, 0xfe, 0xc2, 0xc7,
	0xec, 0x10, 0xf0, 0x3a, 0xd9, 0x44, 0xeb, 0xd3, 0x61, 0xa6, 0xc1, 0x20, 0xf4, 0x8b, 0xed, 0x6e,
	0xd9, 0x05, 0x58, 0x14, 0x69, 0x45, 0xeb, 0x21, 0xe4, 0x45, 0xbc, 0x44, 0xca, 0x68, 0x53, 0x09,
	0x61, 0x22, 0xca, 0xdb, 0xc6, 0xd6, 0x14, 0x98, 0x42, 0x77, 0x8c, 0xef, 0x90, 0x17, 0x68, 0xc7,
	0xa3, 0xdc, 0x2a, 0xad, 0x83, 0x81, 0x03, 0xf0, 0xb4, 0x02, 0xdf, 0xd4, 0xdb, 0x96, 0x41, 0x03,
	0x90, 0x86, 0x7a, 0x9e, 0xd0, 0x5c, 0xe1, 0xbb, 0x64, 0x1b, 0x6d, 0x4d, 0xa0, 0x91, 0xf0, 0x59,
	0xa3, 0x6d, 0x1a, 0x42, 0x02, 0x0b, 0xf8, 0x64, 0xd6, 0xf2, 0xec, 0x8e, 0xc0, 0x75, 0x64, 0x5a,
	0x34, 0xd4, 0x80, 0x09, 0xd9, 0x41, 0x8f, 0x7d, 0x68, 0x41, 0x68, 0x07, 0xe7, 0x0b, 0x9a, 0x26,
	0x95, 0x34, 0x02, 0x05, 0xb2, 0xd0, 0xbf, 0x4a, 0x9e, 0xa1, 0x27, 0xa1, 0x08, 0x18, 0x9f, 0xbd,
	0xf4, 0x25, 0xc0, 0x7b, 0xe4, 0x39, 0x7a, 0x9a, 0x4d, 0xc9, 0x96, 0xe1, 0x61, 0x3b, 0xbb, 0x3f,
	0xc4, 0xca, 0xec, 0xd2, 0x78, 0xf6, 0x8a, 0x6b, 0xe4, 0x3d, 0x7a, 0x3b, 0x91, 0x2c, 0x41, 0x69,
	0xc9, 0x4d, 0xac, 0xa3, 0x88, 0xca, 0xb6, 0x91, 0x62, 0xdf, 0xca, 0x2f, 0x98, 0xfb, 0x4c, 0xed,
	0x0a, 0xad, 0x4c, 0x04, 0x4a, 0x32, 0x2f, 0xc6, 0x0f, 0xc9, 0x3b, 0xf4, 0xe6, 0x6a, 0xea, 0xa5,
	0x12, 0x62, 0x5c, 0xb6, 0x27, 0x65, 0xdc, 0x13, 0x3c, 0x66, 0xb1, 0x02, 0x7e, 0x29, 0x3f, 0xbf,
	0xcc, 0x36, 0x79, 0x85, 0x5e, 0x2a, 0xa1, 0x68, 0x68, 0x5f, 0x9d, 0x0e, 0x55, 0x6c, 0xb2, 0x5b,
	0x67, 0x0f, 0x47, 0x48, 0x16, 0x30, 0x4e, 0xc3, 0x8b, 0xde, 0xe0, 0xe3, 0x47, 0xf5, 0xdf, 0x73,
	0xa8, 0xf2, 0x39, 0x39, 0x71, 0xae, 0xb6, 0x59, 0x7d, 0x65, 0xd6, 0x30, 0x4d, 0xeb, 0xad, 0xe6,
	0xdc, 0xa1, 0x3f, 0x21, 0x0d, 0x92, 0xe3, 0xee, 0x70, 0xe0, 0x24, 0xa3, 0x41, 0x75, 0xd0, 0x1f,
	0x66, 0xce, 0x9b, 0xfa, 0xfb, 0xf4, 0x28, 0xfd, 0x9f, 0xdd, 0x3f, 0xe4, 0x9f, 0xef, 0xf3, 0x0b,
	0x01, 0xa5, 0x3f, 0xe6, 0xcb, 0x41, 0xde, 0x8c, 0xf6, 0x52, 0x27, 0x0f, 0x6d, 0xd4, 0xaa, 0x39,
	0xd9, 0xc8, 0xf4, 0xe7, 0x14, 0xd0, 0xa1, 0xbd, 0xb4, 0x53, 0x00, 0x3a, 0xad, 0x5a, 0x27, 0x07,
	0xfc, 0x9a, 0xaf, 0xe4, 0x59, 0xd7, 0xa5, 0xbd, 0xd4, 0x75, 0x0b, 0x88, 0xeb, 0xb6, 0x6a, 0xae,
	0x9b, 0x83, 0x3e, 0x95, 0x32, 0x75, 0xb5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31, 0x5a, 0x4d,
	0x14, 0x8b, 0x04, 0x00, 0x00,
}