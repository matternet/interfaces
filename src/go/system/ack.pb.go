// Code generated by protoc-gen-go. DO NOT EDIT.
// source: system/ack.proto

package system

import (
	fmt "fmt"
	_ "github.com/airmap/interfaces/src/go"
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

// Ack models an acknowledgement of a message.
type Ack struct {
	Count                uint64               `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Submitted            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=submitted,proto3" json:"submitted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_fec5d8ea51106088, []int{0}
}

func (m *Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ack.Unmarshal(m, b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
}
func (m *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(m, src)
}
func (m *Ack) XXX_Size() int {
	return xxx_messageInfo_Ack.Size(m)
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func (m *Ack) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Ack) GetSubmitted() *timestamp.Timestamp {
	if m != nil {
		return m.Submitted
	}
	return nil
}

func init() {
	proto.RegisterType((*Ack)(nil), "system.Ack")
}

func init() { proto.RegisterFile("system/ack.proto", fileDescriptor_fec5d8ea51106088) }

var fileDescriptor_fec5d8ea51106088 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xbf, 0x4b, 0xc7, 0x30,
	0x10, 0x47, 0xa9, 0x3f, 0xbe, 0x60, 0x5c, 0x4a, 0x71, 0x28, 0x5d, 0x2c, 0x4e, 0x45, 0x24, 0x07,
	0xba, 0xb8, 0xea, 0x9f, 0x50, 0x74, 0x71, 0x4b, 0xe2, 0x35, 0x86, 0x9a, 0x5c, 0xc9, 0x5d, 0x40,
	0xff, 0x7b, 0xa1, 0xa1, 0x38, 0xde, 0xe3, 0x71, 0xef, 0xa3, 0x5a, 0xfe, 0x65, 0xc1, 0x08, 0xc6,
	0xad, 0x7a, 0xcb, 0x24, 0xd4, 0x9d, 0x2a, 0x19, 0x6e, 0x3d, 0x91, 0xff, 0x46, 0xd8, 0xa9, 0x2d,
	0x0b, 0x48, 0x88, 0xc8, 0x62, 0xe2, 0x56, 0xc5, 0xa1, 0xc5, 0x1f, 0xc1, 0xc4, 0x81, 0x12, 0x57,
	0x72, 0xf7, 0xae, 0xce, 0x5f, 0xdc, 0xda, 0xdd, 0xa8, 0x4b, 0x47, 0x25, 0x49, 0xdf, 0x8c, 0xcd,
	0x74, 0x31, 0xd7, 0xa3, 0x7b, 0x56, 0x57, 0x5c, 0x6c, 0x0c, 0x22, 0xf8, 0xd9, 0x9f, 0x8d, 0xcd,
	0x74, 0xfd, 0x38, 0xe8, 0xda, 0xd0, 0x47, 0x43, 0xbf, 0x1d, 0x8d, 0xf9, 0x5f, 0x7e, 0x7d, 0xf8,
	0xb8, 0xf7, 0x41, 0xbe, 0x8a, 0xd5, 0x8e, 0x22, 0x98, 0x90, 0xa3, 0xd9, 0x20, 0x24, 0xc1, 0xbc,
	0x18, 0x87, 0x0c, 0x9c, 0x1d, 0x78, 0x82, 0xba, 0xdb, 0x9e, 0xf6, 0x67, 0x4f, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x5b, 0x96, 0x11, 0x4e, 0xda, 0x00, 0x00, 0x00,
}
