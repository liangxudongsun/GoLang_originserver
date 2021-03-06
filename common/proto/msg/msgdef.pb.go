// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msgproto/msgdef.proto

package msg

import (
	fmt "fmt"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MsgType int32

const (
	MsgType_MsgNil MsgType = 0
	MsgType_MsgReq MsgType = 1
	MsgType_MsgRes MsgType = 2
)

var MsgType_name = map[int32]string{
	0: "MsgNil",
	1: "MsgReq",
	2: "MsgRes",
}

var MsgType_value = map[string]int32{
	"MsgNil": 0,
	"MsgReq": 1,
	"MsgRes": 2,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_514f98c3c98cbc65, []int{0}
}

type Req struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_514f98c3c98cbc65, []int{0}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Res struct {
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Res) Reset()         { *m = Res{} }
func (m *Res) String() string { return proto.CompactTextString(m) }
func (*Res) ProtoMessage()    {}
func (*Res) Descriptor() ([]byte, []int) {
	return fileDescriptor_514f98c3c98cbc65, []int{1}
}

func (m *Res) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Res.Unmarshal(m, b)
}
func (m *Res) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Res.Marshal(b, m, deterministic)
}
func (m *Res) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Res.Merge(m, src)
}
func (m *Res) XXX_Size() int {
	return xxx_messageInfo_Res.Size(m)
}
func (m *Res) XXX_DiscardUnknown() {
	xxx_messageInfo_Res.DiscardUnknown(m)
}

var xxx_messageInfo_Res proto.InternalMessageInfo

func (m *Res) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("msg.MsgType", MsgType_name, MsgType_value)
	proto.RegisterType((*Req)(nil), "msg.Req")
	proto.RegisterType((*Res)(nil), "msg.Res")
}

func init() { proto.RegisterFile("msgproto/msgdef.proto", fileDescriptor_514f98c3c98cbc65) }

var fileDescriptor_514f98c3c98cbc65 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x2d, 0x4e, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x2d, 0x4e, 0x4f, 0x49, 0x4d, 0xd3, 0x03, 0x73, 0x84, 0x98,
	0x73, 0x8b, 0xd3, 0x95, 0xc4, 0xb9, 0x98, 0x83, 0x52, 0x0b, 0x85, 0x04, 0xb8, 0x40, 0x3c, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x84, 0x44, 0x31, 0x4c, 0x82, 0x09, 0x2e, 0xa1, 0xa5, 0xcb,
	0xc5, 0xee, 0x5b, 0x9c, 0x1e, 0x52, 0x59, 0x90, 0x2a, 0xc4, 0xc5, 0xc5, 0xe6, 0x5b, 0x9c, 0xee,
	0x97, 0x99, 0x23, 0xc0, 0x00, 0x65, 0x07, 0xa5, 0x16, 0x0a, 0x30, 0xc2, 0xd9, 0xc5, 0x02, 0x4c,
	0x4e, 0xec, 0x51, 0xac, 0x7a, 0xd6, 0xb9, 0xc5, 0xe9, 0x49, 0x6c, 0x60, 0x5b, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0xcc, 0xf4, 0x0a, 0x8e, 0x00, 0x00, 0x00,
}
