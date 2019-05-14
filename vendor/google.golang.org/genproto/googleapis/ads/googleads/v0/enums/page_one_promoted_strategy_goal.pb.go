// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/page_one_promoted_strategy_goal.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible strategy goals.
type PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal int32

const (
	// Not specified.
	PageOnePromotedStrategyGoalEnum_UNSPECIFIED PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal = 0
	// Used for return value only. Represents value unknown in this version.
	PageOnePromotedStrategyGoalEnum_UNKNOWN PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal = 1
	// First page on google.com.
	PageOnePromotedStrategyGoalEnum_FIRST_PAGE PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal = 2
	// Top slots of the first page on google.com.
	PageOnePromotedStrategyGoalEnum_FIRST_PAGE_PROMOTED PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal = 3
)

var PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "FIRST_PAGE",
	3: "FIRST_PAGE_PROMOTED",
}
var PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"FIRST_PAGE":          2,
	"FIRST_PAGE_PROMOTED": 3,
}

func (x PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal) String() string {
	return proto.EnumName(PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal_name, int32(x))
}
func (PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_page_one_promoted_strategy_goal_dfdc26f09d26e671, []int{0, 0}
}

// Container for enum describing possible strategy goals: where impressions are
// desired to be shown on search result pages.
type PageOnePromotedStrategyGoalEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageOnePromotedStrategyGoalEnum) Reset()         { *m = PageOnePromotedStrategyGoalEnum{} }
func (m *PageOnePromotedStrategyGoalEnum) String() string { return proto.CompactTextString(m) }
func (*PageOnePromotedStrategyGoalEnum) ProtoMessage()    {}
func (*PageOnePromotedStrategyGoalEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_page_one_promoted_strategy_goal_dfdc26f09d26e671, []int{0}
}
func (m *PageOnePromotedStrategyGoalEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageOnePromotedStrategyGoalEnum.Unmarshal(m, b)
}
func (m *PageOnePromotedStrategyGoalEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageOnePromotedStrategyGoalEnum.Marshal(b, m, deterministic)
}
func (dst *PageOnePromotedStrategyGoalEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageOnePromotedStrategyGoalEnum.Merge(dst, src)
}
func (m *PageOnePromotedStrategyGoalEnum) XXX_Size() int {
	return xxx_messageInfo_PageOnePromotedStrategyGoalEnum.Size(m)
}
func (m *PageOnePromotedStrategyGoalEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PageOnePromotedStrategyGoalEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PageOnePromotedStrategyGoalEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PageOnePromotedStrategyGoalEnum)(nil), "google.ads.googleads.v0.enums.PageOnePromotedStrategyGoalEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal", PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal_name, PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/page_one_promoted_strategy_goal.proto", fileDescriptor_page_one_promoted_strategy_goal_dfdc26f09d26e671)
}

var fileDescriptor_page_one_promoted_strategy_goal_dfdc26f09d26e671 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0xfb, 0x30,
	0x1c, 0xc5, 0x7f, 0xed, 0xe0, 0x27, 0x64, 0xa0, 0xa5, 0x5e, 0x78, 0x21, 0x43, 0xb7, 0x07, 0x48,
	0x0b, 0xde, 0xc5, 0xab, 0x6e, 0xeb, 0xca, 0x10, 0xdb, 0xb0, 0x7f, 0x82, 0x14, 0x4a, 0x34, 0x21,
	0x08, 0x6d, 0xbe, 0xa5, 0xe9, 0x06, 0x3e, 0x81, 0xef, 0xe1, 0xa5, 0x8f, 0xe2, 0xa3, 0x78, 0xef,
	0xbd, 0x34, 0xdd, 0xe6, 0x95, 0xbd, 0x29, 0xa7, 0x39, 0x27, 0x1f, 0x4e, 0x0e, 0x9a, 0x48, 0x00,
	0x99, 0x0b, 0x8f, 0x71, 0xed, 0xb5, 0xb2, 0x51, 0x3b, 0xdf, 0x13, 0x6a, 0x5b, 0x68, 0xaf, 0x64,
	0x52, 0x64, 0xa0, 0x44, 0x56, 0x56, 0x50, 0x40, 0x2d, 0x78, 0xa6, 0xeb, 0x8a, 0xd5, 0x42, 0xbe,
	0x66, 0x12, 0x58, 0x8e, 0xcb, 0x0a, 0x6a, 0x70, 0x07, 0xed, 0x4d, 0xcc, 0xb8, 0xc6, 0x47, 0x08,
	0xde, 0xf9, 0xd8, 0x40, 0x46, 0x6f, 0x16, 0xba, 0xa2, 0x4c, 0x8a, 0x44, 0x09, 0xba, 0xc7, 0x2c,
	0xf7, 0x94, 0x08, 0x58, 0x1e, 0xaa, 0x6d, 0x31, 0xe2, 0xe8, 0xb2, 0x23, 0xe2, 0x9e, 0xa1, 0xfe,
	0x3a, 0x5e, 0xd2, 0x70, 0x32, 0x9f, 0xcd, 0xc3, 0xa9, 0xf3, 0xcf, 0xed, 0xa3, 0x93, 0x75, 0x7c,
	0x17, 0x27, 0x0f, 0xb1, 0x63, 0xb9, 0xa7, 0x08, 0xcd, 0xe6, 0x8b, 0xe5, 0x2a, 0xa3, 0x41, 0x14,
	0x3a, 0xb6, 0x7b, 0x81, 0xce, 0x7f, 0xff, 0x33, 0xba, 0x48, 0xee, 0x93, 0x55, 0x38, 0x75, 0x7a,
	0xe3, 0x6f, 0x0b, 0x0d, 0x9f, 0xa1, 0xc0, 0x9d, 0x7d, 0xc7, 0xd7, 0x1d, 0x4d, 0x68, 0xf3, 0x60,
	0x6a, 0x3d, 0x8e, 0xf7, 0x08, 0x09, 0x39, 0x53, 0x12, 0x43, 0x25, 0x3d, 0x29, 0x94, 0x99, 0xe3,
	0xb0, 0x63, 0xf9, 0xa2, 0xff, 0x98, 0xf5, 0xd6, 0x7c, 0xdf, 0xed, 0x5e, 0x14, 0x04, 0x1f, 0xf6,
	0x20, 0x6a, 0x51, 0x01, 0xd7, 0xb8, 0x95, 0x8d, 0xda, 0xf8, 0xb8, 0x19, 0x46, 0x7f, 0x1e, 0xfc,
	0x34, 0xe0, 0x3a, 0x3d, 0xfa, 0xe9, 0xc6, 0x4f, 0x8d, 0xff, 0x65, 0x0f, 0xdb, 0x43, 0x42, 0x02,
	0xae, 0x09, 0x39, 0x26, 0x08, 0xd9, 0xf8, 0x84, 0x98, 0xcc, 0xd3, 0x7f, 0x53, 0xec, 0xe6, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x92, 0xd5, 0x86, 0x26, 0xee, 0x01, 0x00, 0x00,
}
