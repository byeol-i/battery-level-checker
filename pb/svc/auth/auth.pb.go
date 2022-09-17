// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/svc/auth/auth.proto

package auth

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

type AuthTokenReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthTokenReq) Reset()         { *m = AuthTokenReq{} }
func (m *AuthTokenReq) String() string { return proto.CompactTextString(m) }
func (*AuthTokenReq) ProtoMessage()    {}
func (*AuthTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6110729df4464085, []int{0}
}

func (m *AuthTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthTokenReq.Unmarshal(m, b)
}
func (m *AuthTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthTokenReq.Marshal(b, m, deterministic)
}
func (m *AuthTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthTokenReq.Merge(m, src)
}
func (m *AuthTokenReq) XXX_Size() int {
	return xxx_messageInfo_AuthTokenReq.Size(m)
}
func (m *AuthTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_AuthTokenReq proto.InternalMessageInfo

func (m *AuthTokenReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthTokenRes struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthTokenRes) Reset()         { *m = AuthTokenRes{} }
func (m *AuthTokenRes) String() string { return proto.CompactTextString(m) }
func (*AuthTokenRes) ProtoMessage()    {}
func (*AuthTokenRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6110729df4464085, []int{1}
}

func (m *AuthTokenRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthTokenRes.Unmarshal(m, b)
}
func (m *AuthTokenRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthTokenRes.Marshal(b, m, deterministic)
}
func (m *AuthTokenRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthTokenRes.Merge(m, src)
}
func (m *AuthTokenRes) XXX_Size() int {
	return xxx_messageInfo_AuthTokenRes.Size(m)
}
func (m *AuthTokenRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthTokenRes.DiscardUnknown(m)
}

var xxx_messageInfo_AuthTokenRes proto.InternalMessageInfo

func (m *AuthTokenRes) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *AuthTokenRes) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthTokenReq)(nil), "pb.svc.auth.AuthTokenReq")
	proto.RegisterType((*AuthTokenRes)(nil), "pb.svc.auth.AuthTokenRes")
}

func init() { proto.RegisterFile("pb/svc/auth/auth.proto", fileDescriptor_6110729df4464085) }

var fileDescriptor_6110729df4464085 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x48, 0xd2, 0x2f,
	0x2e, 0x4b, 0xd6, 0x4f, 0x2c, 0x2d, 0xc9, 0x00, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0xdc, 0x05, 0x49, 0x7a, 0xc5, 0x65, 0xc9, 0x7a, 0x20, 0x21, 0x25, 0x15, 0x2e, 0x1e, 0xc7, 0xd2,
	0x92, 0x8c, 0x90, 0xfc, 0xec, 0xd4, 0xbc, 0xa0, 0xd4, 0x42, 0x21, 0x11, 0x2e, 0xd6, 0x12, 0x10,
	0x5b, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x51, 0xb2, 0x42, 0x51, 0x55, 0x2c, 0x24,
	0xc6, 0xc5, 0x56, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x02, 0x55, 0x06, 0xe5, 0x09, 0x09, 0x71, 0xb1,
	0x24, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0xd9, 0x46, 0x9e, 0x5c,
	0x2c, 0x20, 0xbd, 0x42, 0x8e, 0x5c, 0x9c, 0x70, 0x33, 0x84, 0x24, 0xf5, 0x90, 0x1c, 0xa1, 0x87,
	0xec, 0x02, 0x29, 0x9c, 0x52, 0xc5, 0x4e, 0x66, 0x51, 0x26, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x49, 0x95, 0xa9, 0xf9, 0x39, 0xba, 0x99, 0xfa, 0x49, 0x89, 0x25,
	0x25, 0xa9, 0x45, 0x95, 0xba, 0x39, 0xa9, 0x65, 0xa9, 0x39, 0xba, 0xc9, 0x19, 0xa9, 0xc9, 0xd9,
	0xa9, 0x45, 0xfa, 0x48, 0x9e, 0x4f, 0x62, 0x03, 0x7b, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x85, 0x65, 0x24, 0x89, 0x12, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	AuthToken(ctx context.Context, in *AuthTokenReq, opts ...grpc.CallOption) (*AuthTokenRes, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) AuthToken(ctx context.Context, in *AuthTokenReq, opts ...grpc.CallOption) (*AuthTokenRes, error) {
	out := new(AuthTokenRes)
	err := c.cc.Invoke(ctx, "/pb.svc.auth.Auth/AuthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	AuthToken(context.Context, *AuthTokenReq) (*AuthTokenRes, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) AuthToken(ctx context.Context, req *AuthTokenReq) (*AuthTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthToken not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_AuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.svc.auth.Auth/AuthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthToken(ctx, req.(*AuthTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.svc.auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthToken",
			Handler:    _Auth_AuthToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/svc/auth/auth.proto",
}
