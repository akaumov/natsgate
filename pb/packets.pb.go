// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/packets.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/packets.proto

It has these top-level messages:
	Request
	Response
*/
package pb

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

type Request struct {
	UserId     string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	DeviceId   string `protobuf:"bytes,2,opt,name=deviceId" json:"deviceId,omitempty"`
	Method     string `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	Host       string `protobuf:"bytes,4,opt,name=host" json:"host,omitempty"`
	RemoteAddr string `protobuf:"bytes,5,opt,name=remoteAddr" json:"remoteAddr,omitempty"`
	RequestURI string `protobuf:"bytes,6,opt,name=requestURI" json:"requestURI,omitempty"`
	Body       []byte `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Request) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Request) GetRemoteAddr() string {
	if m != nil {
		return m.RemoteAddr
	}
	return ""
}

func (m *Request) GetRequestURI() string {
	if m != nil {
		return m.RequestURI
	}
	return ""
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Status int64  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Body   []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*Response)(nil), "pb.Response")
}

func init() { proto.RegisterFile("pb/packets.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4b, 0xc7, 0x30,
	0x10, 0x86, 0x49, 0x5b, 0xdb, 0x7a, 0x38, 0x48, 0x06, 0x09, 0x0e, 0x52, 0x3a, 0x75, 0xd2, 0x41,
	0x70, 0x77, 0xec, 0x1a, 0xf0, 0x03, 0x34, 0xcd, 0x41, 0x45, 0x6a, 0x62, 0xee, 0x2a, 0xf8, 0xe1,
	0xfc, 0x6e, 0xd2, 0xeb, 0x1f, 0x7f, 0xdb, 0xbd, 0xef, 0xc3, 0x13, 0x5e, 0x02, 0xb7, 0xd1, 0x3d,
	0xc5, 0x61, 0xfc, 0x40, 0xa6, 0xc7, 0x98, 0x02, 0x07, 0x9d, 0x45, 0xd7, 0xfe, 0x2a, 0xa8, 0x2c,
	0x7e, 0x2d, 0x48, 0xac, 0xef, 0xa0, 0x5c, 0x08, 0x53, 0xef, 0x8d, 0x6a, 0x54, 0x77, 0x6d, 0xf7,
	0xa4, 0xef, 0xa1, 0xf6, 0xf8, 0xfd, 0x3e, 0x62, 0xef, 0x4d, 0x26, 0xe4, 0xcc, 0xab, 0x33, 0x23,
	0x4f, 0xc1, 0x9b, 0x7c, 0x73, 0xb6, 0xa4, 0x35, 0x14, 0x53, 0x20, 0x36, 0x85, 0xb4, 0x72, 0xeb,
	0x07, 0x80, 0x84, 0x73, 0x60, 0x7c, 0xf5, 0x3e, 0x99, 0x2b, 0x21, 0x17, 0xcd, 0xc6, 0x65, 0xca,
	0x9b, 0xed, 0x4d, 0x79, 0xf0, 0xa3, 0x59, 0xdf, 0x74, 0xc1, 0xff, 0x98, 0xaa, 0x51, 0xdd, 0x8d,
	0x95, 0xbb, 0x7d, 0x81, 0xda, 0x22, 0xc5, 0xf0, 0x49, 0xb8, 0x6e, 0x21, 0x1e, 0x78, 0x21, 0xd9,
	0x9f, 0xdb, 0x3d, 0x9d, 0x5e, 0xf6, 0xef, 0xb9, 0x52, 0xbe, 0xe0, 0xf9, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0xc3, 0xc6, 0x7b, 0x16, 0x01, 0x00, 0x00,
}
