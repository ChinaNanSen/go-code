// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/common/bidding.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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

// Commission is an automatic bidding strategy in which the advertiser pays a
// certain portion of the conversion value.
type Commission struct {
	// Commission rate defines the portion of the conversion value that the
	// advertiser will be billed. A commission rate of x should be passed into
	// this field as (x * 1,000,000). For example, 106,000 represents a commission
	// rate of 0.106 (10.6%).
	CommissionRateMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=commission_rate_micros,json=commissionRateMicros,proto3" json:"commission_rate_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Commission) Reset()         { *m = Commission{} }
func (m *Commission) String() string { return proto.CompactTextString(m) }
func (*Commission) ProtoMessage()    {}
func (*Commission) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{0}
}

func (m *Commission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commission.Unmarshal(m, b)
}
func (m *Commission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commission.Marshal(b, m, deterministic)
}
func (m *Commission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commission.Merge(m, src)
}
func (m *Commission) XXX_Size() int {
	return xxx_messageInfo_Commission.Size(m)
}
func (m *Commission) XXX_DiscardUnknown() {
	xxx_messageInfo_Commission.DiscardUnknown(m)
}

var xxx_messageInfo_Commission proto.InternalMessageInfo

func (m *Commission) GetCommissionRateMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CommissionRateMicros
	}
	return nil
}

// An automated bidding strategy that raises bids for clicks
// that seem more likely to lead to a conversion and lowers
// them for clicks where they seem less likely.
type EnhancedCpc struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnhancedCpc) Reset()         { *m = EnhancedCpc{} }
func (m *EnhancedCpc) String() string { return proto.CompactTextString(m) }
func (*EnhancedCpc) ProtoMessage()    {}
func (*EnhancedCpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{1}
}

func (m *EnhancedCpc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnhancedCpc.Unmarshal(m, b)
}
func (m *EnhancedCpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnhancedCpc.Marshal(b, m, deterministic)
}
func (m *EnhancedCpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnhancedCpc.Merge(m, src)
}
func (m *EnhancedCpc) XXX_Size() int {
	return xxx_messageInfo_EnhancedCpc.Size(m)
}
func (m *EnhancedCpc) XXX_DiscardUnknown() {
	xxx_messageInfo_EnhancedCpc.DiscardUnknown(m)
}

var xxx_messageInfo_EnhancedCpc proto.InternalMessageInfo

// Manual click-based bidding where user pays per click.
type ManualCpc struct {
	// Whether bids are to be enhanced based on conversion optimizer data.
	EnhancedCpcEnabled   *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enhanced_cpc_enabled,json=enhancedCpcEnabled,proto3" json:"enhanced_cpc_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ManualCpc) Reset()         { *m = ManualCpc{} }
func (m *ManualCpc) String() string { return proto.CompactTextString(m) }
func (*ManualCpc) ProtoMessage()    {}
func (*ManualCpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{2}
}

func (m *ManualCpc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualCpc.Unmarshal(m, b)
}
func (m *ManualCpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualCpc.Marshal(b, m, deterministic)
}
func (m *ManualCpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualCpc.Merge(m, src)
}
func (m *ManualCpc) XXX_Size() int {
	return xxx_messageInfo_ManualCpc.Size(m)
}
func (m *ManualCpc) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualCpc.DiscardUnknown(m)
}

var xxx_messageInfo_ManualCpc proto.InternalMessageInfo

func (m *ManualCpc) GetEnhancedCpcEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.EnhancedCpcEnabled
	}
	return nil
}

// Manual impression-based bidding where user pays per thousand impressions.
type ManualCpm struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManualCpm) Reset()         { *m = ManualCpm{} }
func (m *ManualCpm) String() string { return proto.CompactTextString(m) }
func (*ManualCpm) ProtoMessage()    {}
func (*ManualCpm) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{3}
}

func (m *ManualCpm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualCpm.Unmarshal(m, b)
}
func (m *ManualCpm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualCpm.Marshal(b, m, deterministic)
}
func (m *ManualCpm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualCpm.Merge(m, src)
}
func (m *ManualCpm) XXX_Size() int {
	return xxx_messageInfo_ManualCpm.Size(m)
}
func (m *ManualCpm) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualCpm.DiscardUnknown(m)
}

var xxx_messageInfo_ManualCpm proto.InternalMessageInfo

// View based bidding where user pays per video view.
type ManualCpv struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManualCpv) Reset()         { *m = ManualCpv{} }
func (m *ManualCpv) String() string { return proto.CompactTextString(m) }
func (*ManualCpv) ProtoMessage()    {}
func (*ManualCpv) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{4}
}

func (m *ManualCpv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualCpv.Unmarshal(m, b)
}
func (m *ManualCpv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualCpv.Marshal(b, m, deterministic)
}
func (m *ManualCpv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualCpv.Merge(m, src)
}
func (m *ManualCpv) XXX_Size() int {
	return xxx_messageInfo_ManualCpv.Size(m)
}
func (m *ManualCpv) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualCpv.DiscardUnknown(m)
}

var xxx_messageInfo_ManualCpv proto.InternalMessageInfo

// An automated bidding strategy that sets bids to help get the most conversions
// for your campaign while spending your budget.
type MaximizeConversions struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaximizeConversions) Reset()         { *m = MaximizeConversions{} }
func (m *MaximizeConversions) String() string { return proto.CompactTextString(m) }
func (*MaximizeConversions) ProtoMessage()    {}
func (*MaximizeConversions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{5}
}

func (m *MaximizeConversions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaximizeConversions.Unmarshal(m, b)
}
func (m *MaximizeConversions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaximizeConversions.Marshal(b, m, deterministic)
}
func (m *MaximizeConversions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaximizeConversions.Merge(m, src)
}
func (m *MaximizeConversions) XXX_Size() int {
	return xxx_messageInfo_MaximizeConversions.Size(m)
}
func (m *MaximizeConversions) XXX_DiscardUnknown() {
	xxx_messageInfo_MaximizeConversions.DiscardUnknown(m)
}

var xxx_messageInfo_MaximizeConversions proto.InternalMessageInfo

// An automated bidding strategy which tries to maximize conversion value
// given a daily budget.
type MaximizeConversionValue struct {
	// The target return on ad spend (ROAS) option. If set, the bid strategy will
	// maximize revenue while averaging the target return on ad spend. If the
	// target ROAS is high, the bid strategy may not be able to spend the full
	// budget. If the target ROAS is not set, the bid strategy will aim to
	// achieve the highest possible ROAS for the budget.
	TargetRoas           *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=target_roas,json=targetRoas,proto3" json:"target_roas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MaximizeConversionValue) Reset()         { *m = MaximizeConversionValue{} }
func (m *MaximizeConversionValue) String() string { return proto.CompactTextString(m) }
func (*MaximizeConversionValue) ProtoMessage()    {}
func (*MaximizeConversionValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{6}
}

func (m *MaximizeConversionValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaximizeConversionValue.Unmarshal(m, b)
}
func (m *MaximizeConversionValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaximizeConversionValue.Marshal(b, m, deterministic)
}
func (m *MaximizeConversionValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaximizeConversionValue.Merge(m, src)
}
func (m *MaximizeConversionValue) XXX_Size() int {
	return xxx_messageInfo_MaximizeConversionValue.Size(m)
}
func (m *MaximizeConversionValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MaximizeConversionValue.DiscardUnknown(m)
}

var xxx_messageInfo_MaximizeConversionValue proto.InternalMessageInfo

func (m *MaximizeConversionValue) GetTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.TargetRoas
	}
	return nil
}

// An automated bid strategy that sets bids to help get as many conversions as
// possible at the target cost-per-acquisition (CPA) you set.
type TargetCpa struct {
	// Average CPA target.
	// This target should be greater than or equal to minimum billable unit based
	// on the currency for the account.
	TargetCpaMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=target_cpa_micros,json=targetCpaMicros,proto3" json:"target_cpa_micros,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Minimum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidFloorMicros    *wrappers.Int64Value `protobuf:"bytes,3,opt,name=cpc_bid_floor_micros,json=cpcBidFloorMicros,proto3" json:"cpc_bid_floor_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetCpa) Reset()         { *m = TargetCpa{} }
func (m *TargetCpa) String() string { return proto.CompactTextString(m) }
func (*TargetCpa) ProtoMessage()    {}
func (*TargetCpa) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{7}
}

func (m *TargetCpa) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetCpa.Unmarshal(m, b)
}
func (m *TargetCpa) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetCpa.Marshal(b, m, deterministic)
}
func (m *TargetCpa) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetCpa.Merge(m, src)
}
func (m *TargetCpa) XXX_Size() int {
	return xxx_messageInfo_TargetCpa.Size(m)
}
func (m *TargetCpa) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetCpa.DiscardUnknown(m)
}

var xxx_messageInfo_TargetCpa proto.InternalMessageInfo

func (m *TargetCpa) GetTargetCpaMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetCpaMicros
	}
	return nil
}

func (m *TargetCpa) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *TargetCpa) GetCpcBidFloorMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidFloorMicros
	}
	return nil
}

// Target CPM (cost per thousand impressions) is an automated bidding strategy
// that sets bids to optimize performance given the target CPM you set.
type TargetCpm struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetCpm) Reset()         { *m = TargetCpm{} }
func (m *TargetCpm) String() string { return proto.CompactTextString(m) }
func (*TargetCpm) ProtoMessage()    {}
func (*TargetCpm) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{8}
}

func (m *TargetCpm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetCpm.Unmarshal(m, b)
}
func (m *TargetCpm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetCpm.Marshal(b, m, deterministic)
}
func (m *TargetCpm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetCpm.Merge(m, src)
}
func (m *TargetCpm) XXX_Size() int {
	return xxx_messageInfo_TargetCpm.Size(m)
}
func (m *TargetCpm) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetCpm.DiscardUnknown(m)
}

var xxx_messageInfo_TargetCpm proto.InternalMessageInfo

// An automated bidding strategy that sets bids so that a certain percentage of
// search ads are shown at the top of the first page (or other targeted
// location).
// next tag = 4
type TargetImpressionShare struct {
	// The targeted location on the search results page.
	Location enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation `protobuf:"varint,1,opt,name=location,proto3,enum=google.ads.googleads.v3.enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation" json:"location,omitempty"`
	// The desired fraction of ads to be shown in the targeted location in micros.
	// E.g. 1% equals 10,000.
	LocationFractionMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=location_fraction_micros,json=locationFractionMicros,proto3" json:"location_fraction_micros,omitempty"`
	// The highest CPC bid the automated bidding system is permitted to specify.
	// This is a required field entered by the advertiser that sets the ceiling
	// and specified in local micros.
	CpcBidCeilingMicros  *wrappers.Int64Value `protobuf:"bytes,3,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetImpressionShare) Reset()         { *m = TargetImpressionShare{} }
func (m *TargetImpressionShare) String() string { return proto.CompactTextString(m) }
func (*TargetImpressionShare) ProtoMessage()    {}
func (*TargetImpressionShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{9}
}

func (m *TargetImpressionShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetImpressionShare.Unmarshal(m, b)
}
func (m *TargetImpressionShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetImpressionShare.Marshal(b, m, deterministic)
}
func (m *TargetImpressionShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetImpressionShare.Merge(m, src)
}
func (m *TargetImpressionShare) XXX_Size() int {
	return xxx_messageInfo_TargetImpressionShare.Size(m)
}
func (m *TargetImpressionShare) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetImpressionShare.DiscardUnknown(m)
}

var xxx_messageInfo_TargetImpressionShare proto.InternalMessageInfo

func (m *TargetImpressionShare) GetLocation() enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation {
	if m != nil {
		return m.Location
	}
	return enums.TargetImpressionShareLocationEnum_UNSPECIFIED
}

func (m *TargetImpressionShare) GetLocationFractionMicros() *wrappers.Int64Value {
	if m != nil {
		return m.LocationFractionMicros
	}
	return nil
}

func (m *TargetImpressionShare) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

// An automated bidding strategy that helps you maximize revenue while
// averaging a specific target return on ad spend (ROAS).
type TargetRoas struct {
	// Required. The desired revenue (based on conversion data) per unit of spend.
	// Value must be between 0.01 and 1000.0, inclusive.
	TargetRoas *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=target_roas,json=targetRoas,proto3" json:"target_roas,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Minimum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidFloorMicros    *wrappers.Int64Value `protobuf:"bytes,3,opt,name=cpc_bid_floor_micros,json=cpcBidFloorMicros,proto3" json:"cpc_bid_floor_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetRoas) Reset()         { *m = TargetRoas{} }
func (m *TargetRoas) String() string { return proto.CompactTextString(m) }
func (*TargetRoas) ProtoMessage()    {}
func (*TargetRoas) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{10}
}

func (m *TargetRoas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetRoas.Unmarshal(m, b)
}
func (m *TargetRoas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetRoas.Marshal(b, m, deterministic)
}
func (m *TargetRoas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetRoas.Merge(m, src)
}
func (m *TargetRoas) XXX_Size() int {
	return xxx_messageInfo_TargetRoas.Size(m)
}
func (m *TargetRoas) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetRoas.DiscardUnknown(m)
}

var xxx_messageInfo_TargetRoas proto.InternalMessageInfo

func (m *TargetRoas) GetTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.TargetRoas
	}
	return nil
}

func (m *TargetRoas) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *TargetRoas) GetCpcBidFloorMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidFloorMicros
	}
	return nil
}

// An automated bid strategy that sets your bids to help get as many clicks
// as possible within your budget.
type TargetSpend struct {
	// The spend target under which to maximize clicks.
	// A TargetSpend bidder will attempt to spend the smaller of this value
	// or the natural throttling spend amount.
	// If not specified, the budget is used as the spend target.
	TargetSpendMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=target_spend_micros,json=targetSpendMicros,proto3" json:"target_spend_micros,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros  *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetSpend) Reset()         { *m = TargetSpend{} }
func (m *TargetSpend) String() string { return proto.CompactTextString(m) }
func (*TargetSpend) ProtoMessage()    {}
func (*TargetSpend) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{11}
}

func (m *TargetSpend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetSpend.Unmarshal(m, b)
}
func (m *TargetSpend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetSpend.Marshal(b, m, deterministic)
}
func (m *TargetSpend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetSpend.Merge(m, src)
}
func (m *TargetSpend) XXX_Size() int {
	return xxx_messageInfo_TargetSpend.Size(m)
}
func (m *TargetSpend) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetSpend.DiscardUnknown(m)
}

var xxx_messageInfo_TargetSpend proto.InternalMessageInfo

func (m *TargetSpend) GetTargetSpendMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetSpendMicros
	}
	return nil
}

func (m *TargetSpend) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

// A bidding strategy where bids are a fraction of the advertised price for
// some good or service.
type PercentCpc struct {
	// Maximum bid limit that can be set by the bid strategy. This is
	// an optional field entered by the advertiser and specified in local micros.
	// Note: A zero value is interpreted in the same way as having bid_ceiling
	// undefined.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Adjusts the bid for each auction upward or downward, depending on the
	// likelihood of a conversion. Individual bids may exceed
	// cpc_bid_ceiling_micros, but the average bid amount for a campaign should
	// not.
	EnhancedCpcEnabled   *wrappers.BoolValue `protobuf:"bytes,2,opt,name=enhanced_cpc_enabled,json=enhancedCpcEnabled,proto3" json:"enhanced_cpc_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PercentCpc) Reset()         { *m = PercentCpc{} }
func (m *PercentCpc) String() string { return proto.CompactTextString(m) }
func (*PercentCpc) ProtoMessage()    {}
func (*PercentCpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8dc66e731ec0522, []int{12}
}

func (m *PercentCpc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PercentCpc.Unmarshal(m, b)
}
func (m *PercentCpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PercentCpc.Marshal(b, m, deterministic)
}
func (m *PercentCpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PercentCpc.Merge(m, src)
}
func (m *PercentCpc) XXX_Size() int {
	return xxx_messageInfo_PercentCpc.Size(m)
}
func (m *PercentCpc) XXX_DiscardUnknown() {
	xxx_messageInfo_PercentCpc.DiscardUnknown(m)
}

var xxx_messageInfo_PercentCpc proto.InternalMessageInfo

func (m *PercentCpc) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *PercentCpc) GetEnhancedCpcEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.EnhancedCpcEnabled
	}
	return nil
}

func init() {
	proto.RegisterType((*Commission)(nil), "google.ads.googleads.v3.common.Commission")
	proto.RegisterType((*EnhancedCpc)(nil), "google.ads.googleads.v3.common.EnhancedCpc")
	proto.RegisterType((*ManualCpc)(nil), "google.ads.googleads.v3.common.ManualCpc")
	proto.RegisterType((*ManualCpm)(nil), "google.ads.googleads.v3.common.ManualCpm")
	proto.RegisterType((*ManualCpv)(nil), "google.ads.googleads.v3.common.ManualCpv")
	proto.RegisterType((*MaximizeConversions)(nil), "google.ads.googleads.v3.common.MaximizeConversions")
	proto.RegisterType((*MaximizeConversionValue)(nil), "google.ads.googleads.v3.common.MaximizeConversionValue")
	proto.RegisterType((*TargetCpa)(nil), "google.ads.googleads.v3.common.TargetCpa")
	proto.RegisterType((*TargetCpm)(nil), "google.ads.googleads.v3.common.TargetCpm")
	proto.RegisterType((*TargetImpressionShare)(nil), "google.ads.googleads.v3.common.TargetImpressionShare")
	proto.RegisterType((*TargetRoas)(nil), "google.ads.googleads.v3.common.TargetRoas")
	proto.RegisterType((*TargetSpend)(nil), "google.ads.googleads.v3.common.TargetSpend")
	proto.RegisterType((*PercentCpc)(nil), "google.ads.googleads.v3.common.PercentCpc")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/common/bidding.proto", fileDescriptor_c8dc66e731ec0522)
}

var fileDescriptor_c8dc66e731ec0522 = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x6a, 0xd4, 0x4e,
	0x14, 0x27, 0x5b, 0xf8, 0xf3, 0xef, 0xac, 0x1f, 0x34, 0xfd, 0xb0, 0xd4, 0x52, 0x24, 0x57, 0x5e,
	0xc8, 0x04, 0xba, 0xe2, 0x45, 0xc4, 0x8b, 0xdd, 0xf4, 0x83, 0x62, 0x0b, 0xeb, 0xb6, 0x2e, 0x2a,
	0x8b, 0x61, 0x76, 0x32, 0x4d, 0x07, 0x92, 0x99, 0x61, 0x26, 0x59, 0xa5, 0x8f, 0xe3, 0xa5, 0xf8,
	0x08, 0x3e, 0x81, 0x2f, 0x22, 0xe8, 0xad, 0x0f, 0x20, 0x99, 0x8f, 0xac, 0xd0, 0x6e, 0x6b, 0x57,
	0x04, 0xaf, 0x72, 0x4e, 0xe6, 0x77, 0x7e, 0xbf, 0x99, 0x33, 0xe7, 0xcc, 0x01, 0x8f, 0x32, 0xce,
	0xb3, 0x9c, 0x84, 0x28, 0x55, 0xa1, 0x31, 0x6b, 0x6b, 0xd2, 0x09, 0x31, 0x2f, 0x0a, 0xce, 0xc2,
	0x31, 0x4d, 0x53, 0xca, 0x32, 0x28, 0x24, 0x2f, 0xb9, 0xbf, 0x65, 0x20, 0x10, 0xa5, 0x0a, 0x36,
	0x68, 0x38, 0xe9, 0x40, 0x83, 0xde, 0xd8, 0x99, 0xc5, 0x46, 0x58, 0x55, 0xa8, 0xb0, 0x44, 0x32,
	0x23, 0x65, 0x42, 0x0b, 0x21, 0x89, 0x52, 0x94, 0xb3, 0x44, 0x9d, 0x21, 0x49, 0x92, 0x9c, 0x63,
	0x54, 0x52, 0xce, 0x8c, 0xca, 0x86, 0x55, 0x09, 0xb5, 0x37, 0xae, 0x4e, 0xc3, 0x77, 0x12, 0x09,
	0x41, 0xa4, 0xb2, 0xeb, 0x9b, 0x4e, 0x45, 0xd0, 0x10, 0x31, 0xc6, 0x4b, 0x1d, 0x6c, 0x57, 0x83,
	0x04, 0x80, 0x98, 0x17, 0x05, 0xd5, 0x02, 0xfe, 0x0b, 0xb0, 0x86, 0x1b, 0x2f, 0x91, 0xa8, 0x24,
	0x49, 0x41, 0xb1, 0xe4, 0x6a, 0xdd, 0x7b, 0xe0, 0x3d, 0x6c, 0x6f, 0xdf, 0xb7, 0xe7, 0x80, 0x4e,
	0x0c, 0x1e, 0xb0, 0xf2, 0xc9, 0xe3, 0x21, 0xca, 0x2b, 0x32, 0x58, 0x99, 0x86, 0x0e, 0x50, 0x49,
	0x8e, 0x74, 0x60, 0x70, 0x1b, 0xb4, 0x77, 0xd9, 0x19, 0x62, 0x98, 0xa4, 0xb1, 0xc0, 0xc1, 0x6b,
	0xb0, 0x78, 0x84, 0x58, 0x85, 0xf2, 0x58, 0x60, 0xff, 0x10, 0xac, 0x10, 0xbb, 0x96, 0x60, 0x81,
	0x13, 0xc2, 0xd0, 0x38, 0x27, 0xa9, 0x15, 0xdb, 0xb8, 0x20, 0xd6, 0xe3, 0x3c, 0x37, 0x5a, 0x3e,
	0x99, 0x72, 0xee, 0x9a, 0xa8, 0xa0, 0x3d, 0xa5, 0x2e, 0x7e, 0x75, 0x26, 0xc1, 0x2a, 0x58, 0x3e,
	0x42, 0xef, 0x69, 0x41, 0xcf, 0x49, 0xcc, 0xd9, 0x84, 0xc8, 0x7a, 0x8f, 0x2a, 0x78, 0x05, 0xee,
	0x5d, 0xfc, 0xad, 0xf9, 0xfd, 0x67, 0xa0, 0x6d, 0xd3, 0x2f, 0x39, 0x72, 0xa7, 0xdf, 0xbc, 0xb0,
	0xa1, 0x1d, 0x5e, 0x8d, 0x73, 0x62, 0xb6, 0x04, 0x4c, 0xc0, 0x80, 0x23, 0x15, 0xfc, 0xf0, 0xc0,
	0xe2, 0x89, 0x76, 0x63, 0x81, 0xfc, 0x7d, 0xb0, 0x64, 0xc9, 0xb0, 0x40, 0x37, 0x48, 0xe8, 0xdd,
	0xd2, 0x51, 0x98, 0x5c, 0xfa, 0x7d, 0xb0, 0x56, 0xa7, 0x69, 0x4c, 0xd3, 0x04, 0x13, 0x9a, 0x53,
	0x96, 0x39, 0xb6, 0xd6, 0xf5, 0x6c, 0xcb, 0x58, 0xe0, 0x1e, 0x4d, 0x63, 0x13, 0x68, 0x19, 0x0f,
	0xc1, 0x8a, 0x63, 0x3c, 0xcd, 0x39, 0x97, 0x8e, 0x6f, 0xe1, 0x7a, 0xbe, 0x25, 0xc3, 0xb7, 0x57,
	0x87, 0xd9, 0xbb, 0x6e, 0x4f, 0x4f, 0x5d, 0x04, 0x9f, 0x5b, 0x60, 0xd5, 0x78, 0x07, 0x4d, 0x05,
	0x1f, 0xd7, 0x05, 0xec, 0x9f, 0x83, 0xff, 0x5d, 0x0d, 0xeb, 0x34, 0xdc, 0xd9, 0x7e, 0x0b, 0x67,
	0xb5, 0x8a, 0x6e, 0x05, 0x78, 0x29, 0xcf, 0xa1, 0xe5, 0xd8, 0x65, 0x55, 0x71, 0x35, 0x62, 0xd0,
	0xe8, 0xf9, 0x2f, 0xc1, 0xba, 0xb3, 0x93, 0x53, 0x89, 0xb0, 0x36, 0x7e, 0x3f, 0x89, 0x6b, 0x2e,
	0x78, 0xcf, 0xc6, 0x5e, 0x7b, 0x33, 0x0b, 0xf3, 0xdd, 0x4c, 0xf0, 0xdd, 0x03, 0xe0, 0xa4, 0xa9,
	0xa8, 0x3f, 0x2c, 0xc8, 0x7f, 0xbe, 0x72, 0x3e, 0x79, 0xa0, 0x6d, 0x4e, 0x7b, 0x2c, 0x08, 0x4b,
	0xfd, 0xe7, 0x60, 0xd9, 0x1e, 0x57, 0xd5, 0xfe, 0x0d, 0x9a, 0xc6, 0xb6, 0x9a, 0xa6, 0xf9, 0x5b,
	0x6d, 0x53, 0x6f, 0x17, 0xf4, 0x89, 0xc4, 0x84, 0x95, 0xf5, 0x3b, 0x36, 0x5b, 0xc0, 0x9b, 0x3f,
	0xbb, 0x97, 0xbe, 0x8c, 0xad, 0x79, 0x5e, 0xc6, 0xde, 0x57, 0x0f, 0x04, 0x98, 0x17, 0xf0, 0xea,
	0x79, 0xd4, 0xbb, 0xd5, 0x33, 0xe3, 0xab, 0x5f, 0xb3, 0xf6, 0xbd, 0x37, 0x76, 0x3e, 0xc1, 0x8c,
	0xe7, 0x88, 0x65, 0x90, 0xcb, 0x2c, 0xcc, 0x08, 0xd3, 0x9a, 0x6e, 0x5e, 0x09, 0xaa, 0x66, 0x0d,
	0xc3, 0xa7, 0xe6, 0xf3, 0xa1, 0xb5, 0xb0, 0xdf, 0xed, 0x7e, 0x6c, 0x6d, 0xed, 0x1b, 0xb2, 0x6e,
	0xaa, 0xa0, 0x31, 0x6b, 0x6b, 0xd8, 0x81, 0xb1, 0x86, 0x7d, 0x71, 0x80, 0x51, 0x37, 0x55, 0xa3,
	0x06, 0x30, 0x1a, 0x76, 0x46, 0x06, 0xf0, 0xad, 0x15, 0x98, 0xbf, 0x51, 0xd4, 0x4d, 0x55, 0x14,
	0x35, 0x90, 0x28, 0x1a, 0x76, 0xa2, 0xc8, 0x80, 0xc6, 0xff, 0xe9, 0xdd, 0x75, 0x7e, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xa4, 0x8e, 0xbe, 0xfa, 0xa9, 0x07, 0x00, 0x00,
}