// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package greeterpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GreetRequest struct {
	OpeningGreeting *Greeting `protobuf:"bytes,1,opt,name=OpeningGreeting" json:"OpeningGreeting,omitempty"`
}

func (m *GreetRequest) Reset()                    { *m = GreetRequest{} }
func (m *GreetRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()               {}
func (*GreetRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GreetRequest) GetOpeningGreeting() *Greeting {
	if m != nil {
		return m.OpeningGreeting
	}
	return nil
}

type GreetResponse struct {
	ResponseGreeting *Greeting `protobuf:"bytes,1,opt,name=ResponseGreeting" json:"ResponseGreeting,omitempty"`
}

func (m *GreetResponse) Reset()                    { *m = GreetResponse{} }
func (m *GreetResponse) String() string            { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()               {}
func (*GreetResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GreetResponse) GetResponseGreeting() *Greeting {
	if m != nil {
		return m.ResponseGreeting
	}
	return nil
}

func init() {
	proto.RegisterType((*GreetRequest)(nil), "greeter.service.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greeter.service.GreetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := grpc.Invoke(ctx, "/greeter.service.Greeter/Greet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.service.Greeter/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greeter.service.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Greeter_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d,
	0x49, 0x2d, 0xd2, 0x83, 0x0a, 0x4b, 0x41, 0x04, 0x32, 0xf3, 0xd2, 0x8b, 0x21, 0x2a, 0x94, 0x42,
	0xb9, 0x78, 0xdc, 0x41, 0x42, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xae, 0x5c, 0xfc,
	0xfe, 0x05, 0xa9, 0x79, 0x99, 0x79, 0xe9, 0xee, 0x50, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0xd2, 0x7a, 0x30, 0xb3, 0x10, 0x46, 0xc0, 0x94, 0x04, 0xa1, 0xeb, 0x51, 0x8a, 0xe0, 0xe2,
	0x85, 0x1a, 0x5b, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xe4, 0xce, 0x25, 0x00, 0x63, 0x93, 0x62,
	0x30, 0x86, 0x26, 0xa3, 0x40, 0x2e, 0x76, 0x77, 0x88, 0x7a, 0x21, 0x37, 0x2e, 0x56, 0x30, 0x53,
	0x48, 0x56, 0x0f, 0xcd, 0x9f, 0x7a, 0xc8, 0x7e, 0x92, 0x92, 0xc3, 0x25, 0x0d, 0x31, 0xda, 0x89,
	0x3b, 0x8a, 0x13, 0xaa, 0xa0, 0x20, 0x29, 0x89, 0x0d, 0x1c, 0x2e, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x21, 0xa5, 0x64, 0x6a, 0x4a, 0x01, 0x00, 0x00,
}