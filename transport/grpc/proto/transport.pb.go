// Code generated by protoc-gen-go.
// source: github.com/micro/go-plugins/transport/grpc/proto/transport.proto
// DO NOT EDIT!

/*
Package go_micro_grpc_transport is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/go-plugins/transport/grpc/proto/transport.proto

It has these top-level messages:
	Message
*/
package go_micro_grpc_transport

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Header map[string]string `protobuf:"bytes,1,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body   []byte            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.grpc.transport.Message")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Transport service

type TransportClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Transport_StreamClient, error)
}

type transportClient struct {
	cc *grpc.ClientConn
}

func NewTransportClient(cc *grpc.ClientConn) TransportClient {
	return &transportClient{cc}
}

func (c *transportClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Transport_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Transport_serviceDesc.Streams[0], c.cc, "/go.micro.grpc.transport.Transport/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &transportStreamClient{stream}
	return x, nil
}

type Transport_StreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type transportStreamClient struct {
	grpc.ClientStream
}

func (x *transportStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transportStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Transport service

type TransportServer interface {
	Stream(Transport_StreamServer) error
}

func RegisterTransportServer(s *grpc.Server, srv TransportServer) {
	s.RegisterService(&_Transport_serviceDesc, srv)
}

func _Transport_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransportServer).Stream(&transportStreamServer{stream})
}

type Transport_StreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type transportStreamServer struct {
	grpc.ServerStream
}

func (x *transportStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transportStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Transport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.grpc.transport.Transport",
	HandlerType: (*TransportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Transport_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("github.com/micro/go-plugins/transport/grpc/proto/transport.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x72, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x4f, 0xcf, 0xd7,
	0x2d, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x2b, 0xd6, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x2e, 0xc8, 0x2f,
	0x2a, 0xd1, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x47, 0x08, 0xea, 0x81,
	0xf9, 0x42, 0xe2, 0xe9, 0xf9, 0x7a, 0x60, 0x9d, 0x7a, 0x20, 0x45, 0x7a, 0x70, 0x69, 0xa5, 0x79,
	0x8c, 0x5c, 0xec, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x2e, 0x5c, 0x6c, 0x19, 0xa9,
	0x89, 0x29, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x3a, 0x7a, 0x38, 0x74, 0xe9,
	0x41, 0x75, 0xe8, 0x79, 0x80, 0x95, 0xbb, 0xe6, 0x95, 0x14, 0x55, 0x06, 0x41, 0xf5, 0x0a, 0x09,
	0x71, 0xb1, 0x24, 0xe5, 0xa7, 0x54, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x52,
	0x96, 0x5c, 0xdc, 0x48, 0x4a, 0x85, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x81, 0xb6, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x60, 0x5d, 0x9c,
	0x41, 0x10, 0x8e, 0x15, 0x93, 0x05, 0xa3, 0x51, 0x3c, 0x17, 0x67, 0x08, 0xcc, 0x5e, 0xa1, 0x20,
	0x2e, 0xb6, 0xe0, 0x92, 0xa2, 0xd4, 0xc4, 0x5c, 0x21, 0x05, 0x42, 0x6e, 0x93, 0x22, 0xa8, 0x42,
	0x89, 0x41, 0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0x42, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0xcc, 0x76, 0x38, 0x65, 0x01, 0x00, 0x00,
}