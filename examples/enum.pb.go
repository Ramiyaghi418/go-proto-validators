// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/enum.proto

package validator_examples

import (
	fmt "fmt"
	_ "github.com/Ramiyaghi418/go-proto-validators"
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

type Action int32

const (
	Action_ALLOW Action = 0
	Action_DENY  Action = 1
	Action_CHILL Action = 2
)

var Action_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
	2: "CHILL",
}

var Action_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
	"CHILL": 2,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e21dd5d0da6895c, []int{0}
}

type SomeMsg struct {
	Do                   Action   `protobuf:"varint,1,opt,name=do,proto3,enum=validator.examples.Action" json:"do,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SomeMsg) Reset()         { *m = SomeMsg{} }
func (m *SomeMsg) String() string { return proto.CompactTextString(m) }
func (*SomeMsg) ProtoMessage()    {}
func (*SomeMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e21dd5d0da6895c, []int{0}
}

func (m *SomeMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SomeMsg.Unmarshal(m, b)
}
func (m *SomeMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SomeMsg.Marshal(b, m, deterministic)
}
func (m *SomeMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SomeMsg.Merge(m, src)
}
func (m *SomeMsg) XXX_Size() int {
	return xxx_messageInfo_SomeMsg.Size(m)
}
func (m *SomeMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SomeMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SomeMsg proto.InternalMessageInfo

func (m *SomeMsg) GetDo() Action {
	if m != nil {
		return m.Do
	}
	return Action_ALLOW
}

func init() {
	proto.RegisterEnum("validator.examples.Action", Action_name, Action_value)
	proto.RegisterType((*SomeMsg)(nil), "validator.examples.SomeMsg")
}

func init() { proto.RegisterFile("examples/enum.proto", fileDescriptor_5e21dd5d0da6895c) }

var fileDescriptor_5e21dd5d0da6895c = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x4f, 0xcd, 0x2b, 0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x2a, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x83, 0x49, 0x4b, 0x99, 0xa5,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xe7, 0x96, 0x67, 0x96, 0x64, 0xe7,
	0x97, 0xeb, 0xa7, 0xe7, 0xeb, 0x82, 0x35, 0xe8, 0xc2, 0xd5, 0x17, 0xeb, 0x23, 0xb4, 0x82, 0xa5,
	0x94, 0xec, 0xb8, 0xd8, 0x83, 0xf3, 0x73, 0x53, 0x7d, 0x8b, 0xd3, 0x85, 0x8c, 0xb9, 0x98, 0x52,
	0xf2, 0x25, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0xa4, 0xf4, 0x30, 0xed, 0xd0, 0x73, 0x4c, 0x2e,
	0xc9, 0xcc, 0xcf, 0x73, 0x62, 0x7f, 0x74, 0x5f, 0x9e, 0xb9, 0x83, 0x91, 0x31, 0x88, 0x29, 0x25,
	0x5f, 0x4b, 0x83, 0x8b, 0x0d, 0x22, 0x2c, 0xc4, 0xc9, 0xc5, 0xea, 0xe8, 0xe3, 0xe3, 0x1f, 0x2e,
	0xc0, 0x20, 0xc4, 0xc1, 0xc5, 0xe2, 0xe2, 0xea, 0x17, 0x29, 0xc0, 0x08, 0x12, 0x74, 0xf6, 0xf0,
	0xf4, 0xf1, 0x11, 0x60, 0x4a, 0x62, 0x03, 0x5b, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0x4f, 0x55, 0x68, 0xd3, 0x00, 0x00, 0x00,
}
