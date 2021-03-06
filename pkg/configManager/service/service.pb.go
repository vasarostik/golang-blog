// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package config_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ConfStruct struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfStruct) Reset()         { *m = ConfStruct{} }
func (m *ConfStruct) String() string { return proto.CompactTextString(m) }
func (*ConfStruct) ProtoMessage()    {}
func (*ConfStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *ConfStruct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfStruct.Unmarshal(m, b)
}
func (m *ConfStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfStruct.Marshal(b, m, deterministic)
}
func (m *ConfStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfStruct.Merge(m, src)
}
func (m *ConfStruct) XXX_Size() int {
	return xxx_messageInfo_ConfStruct.Size(m)
}
func (m *ConfStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfStruct.DiscardUnknown(m)
}

var xxx_messageInfo_ConfStruct proto.InternalMessageInfo

func (m *ConfStruct) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
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

func init() {
	proto.RegisterType((*ConfStruct)(nil), "config_service.ConfStruct")
	proto.RegisterType((*Request)(nil), "config_service.Request")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0xce, 0xcf, 0x4b, 0xcb,
	0x4c, 0x8f, 0x87, 0x8a, 0x2a, 0x29, 0x70, 0x71, 0x39, 0xe7, 0xe7, 0xa5, 0x05, 0x97, 0x14, 0x95,
	0x26, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xa4, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0,
	0x04, 0x81, 0xd9, 0x4a, 0x9c, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x46, 0xd7,
	0x19, 0xb9, 0x58, 0x82, 0x53, 0xf3, 0x52, 0x84, 0x9c, 0xb9, 0x78, 0xdc, 0x53, 0x4b, 0x1c, 0x03,
	0x3c, 0x9d, 0xc1, 0xa6, 0x09, 0x89, 0xeb, 0xa1, 0x1a, 0xab, 0x07, 0xd5, 0x21, 0x25, 0x85, 0x2e,
	0x81, 0x64, 0x99, 0x0b, 0x17, 0xaf, 0x7b, 0x6a, 0x89, 0x7b, 0x50, 0x80, 0x33, 0xe5, 0xa6, 0xf8,
	0x39, 0x86, 0x04, 0x53, 0x60, 0x4a, 0x12, 0x1b, 0x38, 0x74, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0x9d, 0x8c, 0x91, 0x2e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SendClient is the client API for Send service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SendClient interface {
	// Send Config File
	GetAPIConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ConfStruct, error)
	GetGRPCConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ConfStruct, error)
	GetNATSConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ConfStruct, error)
}

type sendClient struct {
	cc *grpc.ClientConn
}

func NewSendClient(cc *grpc.ClientConn) SendClient {
	return &sendClient{cc}
}

func (c *sendClient) GetAPIConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ConfStruct, error) {
	out := new(ConfStruct)
	err := c.cc.Invoke(ctx, "/config_service.Send/GetAPIConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendClient) GetGRPCConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ConfStruct, error) {
	out := new(ConfStruct)
	err := c.cc.Invoke(ctx, "/config_service.Send/GetGRPCConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendClient) GetNATSConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ConfStruct, error) {
	out := new(ConfStruct)
	err := c.cc.Invoke(ctx, "/config_service.Send/GetNATSConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendServer is the server API for Send service.
type SendServer interface {
	// Send Config File
	GetAPIConfig(context.Context, *Request) (*ConfStruct, error)
	GetGRPCConfig(context.Context, *Request) (*ConfStruct, error)
	GetNATSConfig(context.Context, *Request) (*ConfStruct, error)
}

// UnimplementedSendServer can be embedded to have forward compatible implementations.
type UnimplementedSendServer struct {
}

func (*UnimplementedSendServer) GetAPIConfig(ctx context.Context, req *Request) (*ConfStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIConfig not implemented")
}
func (*UnimplementedSendServer) GetGRPCConfig(ctx context.Context, req *Request) (*ConfStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGRPCConfig not implemented")
}
func (*UnimplementedSendServer) GetNATSConfig(ctx context.Context, req *Request) (*ConfStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNATSConfig not implemented")
}

func RegisterSendServer(s *grpc.Server, srv SendServer) {
	s.RegisterService(&_Send_serviceDesc, srv)
}

func _Send_GetAPIConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendServer).GetAPIConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_service.Send/GetAPIConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendServer).GetAPIConfig(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Send_GetGRPCConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendServer).GetGRPCConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_service.Send/GetGRPCConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendServer).GetGRPCConfig(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Send_GetNATSConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendServer).GetNATSConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_service.Send/GetNATSConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendServer).GetNATSConfig(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Send_serviceDesc = grpc.ServiceDesc{
	ServiceName: "config_service.Send",
	HandlerType: (*SendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAPIConfig",
			Handler:    _Send_GetAPIConfig_Handler,
		},
		{
			MethodName: "GetGRPCConfig",
			Handler:    _Send_GetGRPCConfig_Handler,
		},
		{
			MethodName: "GetNATSConfig",
			Handler:    _Send_GetNATSConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
