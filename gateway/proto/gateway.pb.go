// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() {
	proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5)
}

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0xcd, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0x07, 0x70, 0xba, 0xb8, 0xbb, 0x1a, 0xd4, 0x62, 0x10, 0x91, 0xb2, 0xa8, 0x14, 0x15, 0xf1,
	0xd0, 0xa0, 0xde, 0xf4, 0xe4, 0x22, 0xb8, 0xe7, 0x7a, 0x11, 0x6f, 0xd3, 0x76, 0x48, 0x03, 0x69,
	0xa7, 0xa6, 0xd9, 0xca, 0x5e, 0x7d, 0x05, 0xef, 0xbe, 0x94, 0x0f, 0x20, 0x88, 0x0f, 0x62, 0x37,
	0x11, 0xf5, 0xe6, 0x69, 0x92, 0xf9, 0xf8, 0xfd, 0xd9, 0x86, 0x04, 0x8b, 0x4f, 0xb0, 0x48, 0x1a,
	0x43, 0x96, 0xf8, 0xd0, 0x95, 0x68, 0x22, 0x89, 0xa4, 0x46, 0x01, 0x8d, 0x12, 0x50, 0xd7, 0x64,
	0xc1, 0x2a, 0xaa, 0x5b, 0xbf, 0x14, 0x85, 0x1d, 0x68, 0x55, 0x80, 0x25, 0xe3, 0x1b, 0xf1, 0x11,
	0x1b, 0xa7, 0xf8, 0x38, 0xc7, 0xd6, 0xf2, 0x88, 0xad, 0xd4, 0x50, 0xe1, 0x6e, 0x70, 0x10, 0x9c,
	0xac, 0x4d, 0x47, 0x1f, 0xef, 0xfb, 0x83, 0xfb, 0x20, 0x75, 0xbd, 0x78, 0x8f, 0xad, 0xa6, 0xd8,
	0x36, 0x3d, 0x84, 0x9c, 0xff, 0xdd, 0xf3, 0xf3, 0xf3, 0xd7, 0x80, 0xad, 0x5f, 0xeb, 0xa6, 0x84,
	0x3b, 0x34, 0x9d, 0xca, 0x91, 0xdf, 0xb0, 0xe1, 0x0c, 0xb5, 0x26, 0xbe, 0xe9, 0x83, 0x92, 0xef,
	0x94, 0x28, 0xfc, 0xf9, 0x7b, 0x2e, 0x8e, 0x9e, 0xdf, 0x3e, 0x5f, 0x06, 0xdb, 0x71, 0x28, 0xba,
	0x33, 0x01, 0x4b, 0x43, 0x94, 0xcb, 0xcb, 0xcb, 0xe0, 0x94, 0xcf, 0xd8, 0xf8, 0x96, 0xa8, 0xc8,
	0x16, 0xf8, 0xbf, 0x33, 0x71, 0xce, 0x4e, 0xbc, 0xf5, 0xeb, 0x48, 0x7f, 0xdb, 0x4b, 0xd3, 0xe3,
	0x87, 0x43, 0xa9, 0x6c, 0x39, 0xcf, 0x92, 0x9c, 0x2a, 0x51, 0x29, 0x63, 0xc8, 0xb4, 0x12, 0xfb,
	0x57, 0x6e, 0xa8, 0xc0, 0x8a, 0xae, 0x1c, 0x97, 0x8d, 0x5c, 0xb9, 0xf8, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x08, 0x6d, 0xf0, 0xee, 0x5d, 0x01, 0x00, 0x00,
}
