// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rpcproto/rpcdef.proto

package rpc

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

type TestOne struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestOne) Reset()         { *m = TestOne{} }
func (m *TestOne) String() string { return proto.CompactTextString(m) }
func (*TestOne) ProtoMessage()    {}
func (*TestOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_233c86e64e2bb07b, []int{0}
}

func (m *TestOne) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestOne.Unmarshal(m, b)
}
func (m *TestOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestOne.Marshal(b, m, deterministic)
}
func (m *TestOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestOne.Merge(m, src)
}
func (m *TestOne) XXX_Size() int {
	return xxx_messageInfo_TestOne.Size(m)
}
func (m *TestOne) XXX_DiscardUnknown() {
	xxx_messageInfo_TestOne.DiscardUnknown(m)
}

var xxx_messageInfo_TestOne proto.InternalMessageInfo

func (m *TestOne) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type TestOneRet struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestOneRet) Reset()         { *m = TestOneRet{} }
func (m *TestOneRet) String() string { return proto.CompactTextString(m) }
func (*TestOneRet) ProtoMessage()    {}
func (*TestOneRet) Descriptor() ([]byte, []int) {
	return fileDescriptor_233c86e64e2bb07b, []int{1}
}

func (m *TestOneRet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestOneRet.Unmarshal(m, b)
}
func (m *TestOneRet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestOneRet.Marshal(b, m, deterministic)
}
func (m *TestOneRet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestOneRet.Merge(m, src)
}
func (m *TestOneRet) XXX_Size() int {
	return xxx_messageInfo_TestOneRet.Size(m)
}
func (m *TestOneRet) XXX_DiscardUnknown() {
	xxx_messageInfo_TestOneRet.DiscardUnknown(m)
}

var xxx_messageInfo_TestOneRet proto.InternalMessageInfo

func (m *TestOneRet) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type TestTwo struct {
	Data                 int32    `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestTwo) Reset()         { *m = TestTwo{} }
func (m *TestTwo) String() string { return proto.CompactTextString(m) }
func (*TestTwo) ProtoMessage()    {}
func (*TestTwo) Descriptor() ([]byte, []int) {
	return fileDescriptor_233c86e64e2bb07b, []int{2}
}

func (m *TestTwo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestTwo.Unmarshal(m, b)
}
func (m *TestTwo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestTwo.Marshal(b, m, deterministic)
}
func (m *TestTwo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestTwo.Merge(m, src)
}
func (m *TestTwo) XXX_Size() int {
	return xxx_messageInfo_TestTwo.Size(m)
}
func (m *TestTwo) XXX_DiscardUnknown() {
	xxx_messageInfo_TestTwo.DiscardUnknown(m)
}

var xxx_messageInfo_TestTwo proto.InternalMessageInfo

func (m *TestTwo) GetData() int32 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *TestTwo) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type TestTwoRet struct {
	Data                 int32    `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestTwoRet) Reset()         { *m = TestTwoRet{} }
func (m *TestTwoRet) String() string { return proto.CompactTextString(m) }
func (*TestTwoRet) ProtoMessage()    {}
func (*TestTwoRet) Descriptor() ([]byte, []int) {
	return fileDescriptor_233c86e64e2bb07b, []int{3}
}

func (m *TestTwoRet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestTwoRet.Unmarshal(m, b)
}
func (m *TestTwoRet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestTwoRet.Marshal(b, m, deterministic)
}
func (m *TestTwoRet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestTwoRet.Merge(m, src)
}
func (m *TestTwoRet) XXX_Size() int {
	return xxx_messageInfo_TestTwoRet.Size(m)
}
func (m *TestTwoRet) XXX_DiscardUnknown() {
	xxx_messageInfo_TestTwoRet.DiscardUnknown(m)
}

var xxx_messageInfo_TestTwoRet proto.InternalMessageInfo

func (m *TestTwoRet) GetData() int32 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *TestTwoRet) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type TestThree struct {
	UList                []uint64 `protobuf:"varint,1,rep,packed,name=UList,proto3" json:"UList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestThree) Reset()         { *m = TestThree{} }
func (m *TestThree) String() string { return proto.CompactTextString(m) }
func (*TestThree) ProtoMessage()    {}
func (*TestThree) Descriptor() ([]byte, []int) {
	return fileDescriptor_233c86e64e2bb07b, []int{4}
}

func (m *TestThree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestThree.Unmarshal(m, b)
}
func (m *TestThree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestThree.Marshal(b, m, deterministic)
}
func (m *TestThree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestThree.Merge(m, src)
}
func (m *TestThree) XXX_Size() int {
	return xxx_messageInfo_TestThree.Size(m)
}
func (m *TestThree) XXX_DiscardUnknown() {
	xxx_messageInfo_TestThree.DiscardUnknown(m)
}

var xxx_messageInfo_TestThree proto.InternalMessageInfo

func (m *TestThree) GetUList() []uint64 {
	if m != nil {
		return m.UList
	}
	return nil
}

func init() {
	proto.RegisterType((*TestOne)(nil), "rpc.TestOne")
	proto.RegisterType((*TestOneRet)(nil), "rpc.TestOneRet")
	proto.RegisterType((*TestTwo)(nil), "rpc.TestTwo")
	proto.RegisterType((*TestTwoRet)(nil), "rpc.TestTwoRet")
	proto.RegisterType((*TestThree)(nil), "rpc.TestThree")
}

func init() { proto.RegisterFile("proto/rpcproto/rpcdef.proto", fileDescriptor_233c86e64e2bb07b) }

var fileDescriptor_233c86e64e2bb07b = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2a, 0x48, 0x86, 0x33, 0x52, 0x52, 0xd3, 0xf4, 0xc0, 0x1c, 0x21, 0xe6, 0xa2,
	0x82, 0x64, 0x25, 0x69, 0x2e, 0xf6, 0x90, 0xd4, 0xe2, 0x12, 0xff, 0xbc, 0x54, 0x21, 0x01, 0x2e,
	0xe6, 0xdc, 0xe2, 0x74, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x49, 0x8e, 0x8b,
	0x0b, 0x2a, 0x19, 0x94, 0x5a, 0x82, 0x45, 0x5e, 0x1f, 0xa2, 0x39, 0xa4, 0x3c, 0x5f, 0x48, 0x88,
	0x8b, 0x25, 0x25, 0xb1, 0x24, 0x11, 0x2c, 0xcb, 0x1a, 0x04, 0x66, 0xc3, 0x34, 0x30, 0x21, 0x34,
	0x18, 0x41, 0x0c, 0x0c, 0x29, 0xcf, 0x07, 0x19, 0x48, 0x9c, 0x1e, 0x45, 0x2e, 0x4e, 0xb0, 0x9e,
	0x8c, 0xa2, 0xd4, 0x54, 0x21, 0x11, 0x2e, 0xd6, 0x50, 0x9f, 0xcc, 0xe2, 0x12, 0x09, 0x46, 0x05,
	0x66, 0x0d, 0x96, 0x20, 0x08, 0xc7, 0x89, 0x3d, 0x8a, 0x55, 0xcf, 0xba, 0xa8, 0x20, 0x39, 0x89,
	0x0d, 0xec, 0x33, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x12, 0x3b, 0x0a, 0xf8, 0x00,
	0x00, 0x00,
}
