// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/testing/service.proto

package zipkin_testing

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

type HelloRequest struct {
	Payload              string   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_eae06f0e4375fa1a, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type HelloResponse struct {
	Payload              string            `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SpanContext          map[string]string `protobuf:"bytes,3,rep,name=span_context,json=spanContext,proto3" json:"span_context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_eae06f0e4375fa1a, []int{1}
}
func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (dst *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(dst, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *HelloResponse) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *HelloResponse) GetSpanContext() map[string]string {
	if m != nil {
		return m.SpanContext
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "zipkin.testing.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "zipkin.testing.HelloResponse")
	proto.RegisterMapType((map[string]string)(nil), "zipkin.testing.HelloResponse.MetadataEntry")
	proto.RegisterMapType((map[string]string)(nil), "zipkin.testing.HelloResponse.SpanContextEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/zipkin.testing.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zipkin.testing.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zipkin.testing.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/testing/service.proto",
}

func init() {
	proto.RegisterFile("proto/testing/service.proto", fileDescriptor_service_eae06f0e4375fa1a)
}

var fileDescriptor_service_eae06f0e4375fa1a = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x49, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0xd5, 0x03, 0x8b, 0x0a, 0xf1, 0x55, 0x65, 0x16, 0x64, 0x67, 0xe6, 0xe9, 0x41, 0x65, 0x95,
	0x34, 0xb8, 0x78, 0x3c, 0x52, 0x73, 0x72, 0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0x24, 0xb8, 0xd8, 0x0b, 0x12, 0x2b, 0x73, 0xf2, 0x13, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x60, 0x5c, 0xa5, 0x43, 0x4c, 0x5c, 0xbc, 0x50, 0xa5, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9,
	0xb8, 0xd5, 0x0a, 0xb9, 0x73, 0x71, 0xe4, 0xa6, 0x96, 0x24, 0xa6, 0x24, 0x96, 0x24, 0x4a, 0x30,
	0x29, 0x30, 0x6b, 0x70, 0x1b, 0x69, 0xeb, 0xa1, 0x5a, 0xac, 0x87, 0x62, 0x94, 0x9e, 0x2f, 0x54,
	0xb5, 0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10, 0x5c, 0xb3, 0x50, 0x20, 0x17, 0x4f, 0x71, 0x41, 0x62,
	0x5e, 0x7c, 0x72, 0x7e, 0x5e, 0x49, 0x6a, 0x45, 0x89, 0x04, 0x33, 0xd8, 0x30, 0x3d, 0xfc, 0x86,
	0x05, 0x17, 0x24, 0xe6, 0x39, 0x43, 0x34, 0x40, 0xcc, 0xe3, 0x2e, 0x46, 0x88, 0x48, 0x59, 0x73,
	0xf1, 0xa2, 0xd8, 0x26, 0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x09, 0xf5, 0x02, 0x88, 0x29, 0x24,
	0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x04, 0x16, 0x83, 0x70, 0xac, 0x98, 0x2c,
	0x18, 0xa5, 0xec, 0xb8, 0x04, 0xd0, 0x4d, 0x27, 0x45, 0xbf, 0x51, 0x08, 0x34, 0xb8, 0x83, 0x21,
	0x91, 0x22, 0xe4, 0xc2, 0xc5, 0x0a, 0xe6, 0x0b, 0xc9, 0xe0, 0xf0, 0x12, 0x38, 0x56, 0xa4, 0x64,
	0xf1, 0x7a, 0x38, 0x89, 0x0d, 0x1c, 0xb7, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x24,
	0x1b, 0x5e, 0xfa, 0x01, 0x00, 0x00,
}
