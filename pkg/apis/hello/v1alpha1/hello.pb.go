// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

/*
Package v1alpha1 is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	IntroRequest
	IntroResponse
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type IntroRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *IntroRequest) Reset()                    { *m = IntroRequest{} }
func (m *IntroRequest) String() string            { return proto.CompactTextString(m) }
func (*IntroRequest) ProtoMessage()               {}
func (*IntroRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IntroRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type IntroResponse struct {
	Intro string `protobuf:"bytes,1,opt,name=intro" json:"intro,omitempty"`
}

func (m *IntroResponse) Reset()                    { *m = IntroResponse{} }
func (m *IntroResponse) String() string            { return proto.CompactTextString(m) }
func (*IntroResponse) ProtoMessage()               {}
func (*IntroResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IntroResponse) GetIntro() string {
	if m != nil {
		return m.Intro
	}
	return ""
}

func init() {
	proto.RegisterType((*IntroRequest)(nil), "appscode.hello.v1alpha1.IntroRequest")
	proto.RegisterType((*IntroResponse)(nil), "appscode.hello.v1alpha1.IntroResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hello service

type HelloClient interface {
	Intro(ctx context.Context, in *IntroRequest, opts ...grpc.CallOption) (*IntroResponse, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Intro(ctx context.Context, in *IntroRequest, opts ...grpc.CallOption) (*IntroResponse, error) {
	out := new(IntroResponse)
	err := grpc.Invoke(ctx, "/appscode.hello.v1alpha1.Hello/Intro", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	Intro(context.Context, *IntroRequest) (*IntroResponse, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Intro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Intro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.hello.v1alpha1.Hello/Intro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Intro(ctx, req.(*IntroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.hello.v1alpha1.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Intro",
			Handler:    _Hello_Intro_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4f, 0x2c, 0x28, 0x28, 0x4e, 0xce, 0x4f,
	0x49, 0xd5, 0x83, 0x88, 0x96, 0x19, 0x26, 0xe6, 0x14, 0x64, 0x24, 0x1a, 0x4a, 0xc9, 0xa4, 0xe7,
	0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96,
	0x64, 0xe6, 0xe7, 0x15, 0x43, 0xb4, 0x49, 0xc9, 0xc1, 0xb4, 0x61, 0x97, 0x57, 0x52, 0xe2, 0xe2,
	0xf1, 0xcc, 0x2b, 0x29, 0xca, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62,
	0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x54, 0xb9,
	0x78, 0xa1, 0x6a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x44, 0xb8, 0x58, 0x33, 0x41, 0x02,
	0x50, 0x55, 0x10, 0x8e, 0x51, 0x07, 0x23, 0x17, 0xab, 0x07, 0xc8, 0x6d, 0x42, 0xf5, 0x5c, 0xac,
	0x60, 0x0d, 0x42, 0xaa, 0x7a, 0x38, 0x5c, 0xad, 0x87, 0x6c, 0xa9, 0x94, 0x1a, 0x21, 0x65, 0x10,
	0x7b, 0x95, 0xd4, 0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0xa4, 0x28, 0x24, 0x0f, 0xf2, 0x4c, 0xb1, 0x3e,
	0x58, 0xad, 0x3e, 0x4c, 0xad, 0x3e, 0xd8, 0x15, 0xfa, 0x59, 0xc5, 0xf9, 0x79, 0x4e, 0x96, 0x5c,
	0xf2, 0xc9, 0xf9, 0xb9, 0x08, 0x53, 0x13, 0x0b, 0x32, 0xd1, 0x4c, 0x76, 0xe2, 0x02, 0x3b, 0x35,
	0x00, 0x14, 0x08, 0x01, 0x8c, 0x51, 0x1c, 0x30, 0xf1, 0x24, 0x36, 0x70, 0xb8, 0x18, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xb2, 0x3f, 0x9e, 0x4b, 0x7d, 0x01, 0x00, 0x00,
}
