// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drand/empty.proto

package drand // import "github.com/dedis/drand/protobuf/drand"

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_empty_61fad5fd43cd12bd, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "drand.Empty")
}

func init() { proto.RegisterFile("drand/empty.proto", fileDescriptor_empty_61fad5fd43cd12bd) }

var fileDescriptor_empty_61fad5fd43cd12bd = []byte{
	// 88 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x29, 0x4a, 0xcc,
	0x4b, 0xd1, 0x4f, 0xcd, 0x2d, 0x28, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x0b, 0x29, 0xb1, 0x73, 0xb1, 0xba, 0x82, 0x44, 0x9d, 0xd4, 0xa3, 0x54, 0xd3, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x53, 0x52, 0x53, 0x32, 0x8b, 0xf5, 0x21, 0xba, 0xc0,
	0xea, 0x93, 0x4a, 0xd3, 0x20, 0xdc, 0x24, 0x36, 0x30, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x35, 0xcc, 0x6c, 0xac, 0x54, 0x00, 0x00, 0x00,
}
