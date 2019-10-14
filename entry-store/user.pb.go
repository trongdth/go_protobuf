// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package protos

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

type UserReq struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReq) Reset()         { *m = UserReq{} }
func (m *UserReq) String() string { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()    {}
func (*UserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReq.Unmarshal(m, b)
}
func (m *UserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReq.Marshal(b, m, deterministic)
}
func (m *UserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReq.Merge(m, src)
}
func (m *UserReq) XXX_Size() int {
	return xxx_messageInfo_UserReq.Size(m)
}
func (m *UserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserReq proto.InternalMessageInfo

func (m *UserReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateUserReq struct {
	FullName             string   `protobuf:"bytes,1,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserReq) Reset()         { *m = CreateUserReq{} }
func (m *CreateUserReq) String() string { return proto.CompactTextString(m) }
func (*CreateUserReq) ProtoMessage()    {}
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *CreateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReq.Unmarshal(m, b)
}
func (m *CreateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReq.Marshal(b, m, deterministic)
}
func (m *CreateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReq.Merge(m, src)
}
func (m *CreateUserReq) XXX_Size() int {
	return xxx_messageInfo_CreateUserReq.Size(m)
}
func (m *CreateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReq proto.InternalMessageInfo

func (m *CreateUserReq) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *CreateUserReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserRes struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRes) Reset()         { *m = UserRes{} }
func (m *UserRes) String() string { return proto.CompactTextString(m) }
func (*UserRes) ProtoMessage()    {}
func (*UserRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRes.Unmarshal(m, b)
}
func (m *UserRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRes.Marshal(b, m, deterministic)
}
func (m *UserRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRes.Merge(m, src)
}
func (m *UserRes) XXX_Size() int {
	return xxx_messageInfo_UserRes.Size(m)
}
func (m *UserRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRes.DiscardUnknown(m)
}

var xxx_messageInfo_UserRes proto.InternalMessageInfo

func (m *UserRes) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRes) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *UserRes) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRes) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*UserReq)(nil), "protos.UserReq")
	proto.RegisterType((*CreateUserReq)(nil), "protos.CreateUserReq")
	proto.RegisterType((*UserRes)(nil), "protos.UserRes")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x92, 0x5c, 0xec, 0xa1,
	0xc5, 0xa9, 0x45, 0x41, 0xa9, 0x85, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xbc, 0x41, 0x4c, 0x99, 0x29, 0x4a, 0x95, 0x5c, 0xbc, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9,
	0x30, 0x05, 0x52, 0x5c, 0x1c, 0x69, 0xa5, 0x39, 0x39, 0x7e, 0x89, 0xb9, 0xa9, 0x60, 0x65, 0x9c,
	0x41, 0x70, 0xbe, 0x90, 0x08, 0x17, 0x6b, 0x6a, 0x6e, 0x62, 0x66, 0x8e, 0x04, 0x13, 0x58, 0x02,
	0xc2, 0x01, 0xe9, 0x28, 0x48, 0x2c, 0x2e, 0x2e, 0xcf, 0x2f, 0x4a, 0x91, 0x60, 0x86, 0xe8, 0x80,
	0xf1, 0x41, 0x72, 0x20, 0xf7, 0xe4, 0x81, 0x4c, 0x63, 0x81, 0xc8, 0xc1, 0xf8, 0x4a, 0xe9, 0x30,
	0x57, 0x15, 0xa3, 0xbb, 0x0a, 0xc5, 0x11, 0x4c, 0xb8, 0x1c, 0xc1, 0x8c, 0xe6, 0x08, 0x5c, 0x16,
	0x19, 0x15, 0x40, 0x2c, 0x0a, 0x2e, 0x2a, 0x13, 0xd2, 0xe5, 0x62, 0x77, 0x4f, 0x2d, 0x01, 0xf1,
	0x84, 0xf8, 0x21, 0x81, 0x54, 0xac, 0x07, 0xf5, 0xb9, 0x14, 0x9a, 0x40, 0xb1, 0x12, 0x83, 0x90,
	0x19, 0x17, 0x17, 0x22, 0x74, 0x84, 0x44, 0x61, 0x0a, 0x50, 0x42, 0x0c, 0x8b, 0xbe, 0x24, 0x48,
	0xc0, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x8b, 0x48, 0x7b, 0x8d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserSrvClient is the client API for UserSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSrvClient interface {
	GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*UserRes, error)
}

type userSrvClient struct {
	cc *grpc.ClientConn
}

func NewUserSrvClient(cc *grpc.ClientConn) UserSrvClient {
	return &userSrvClient{cc}
}

func (c *userSrvClient) GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/protos.UserSrv/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/protos.UserSrv/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSrvServer is the server API for UserSrv service.
type UserSrvServer interface {
	GetUser(context.Context, *UserReq) (*UserRes, error)
	CreateUser(context.Context, *CreateUserReq) (*UserRes, error)
}

// UnimplementedUserSrvServer can be embedded to have forward compatible implementations.
type UnimplementedUserSrvServer struct {
}

func (*UnimplementedUserSrvServer) GetUser(ctx context.Context, req *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserSrvServer) CreateUser(ctx context.Context, req *CreateUserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUserSrvServer(s *grpc.Server, srv UserSrvServer) {
	s.RegisterService(&_UserSrv_serviceDesc, srv)
}

func _UserSrv_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSrvServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserSrv/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSrvServer).GetUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSrv_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSrvServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserSrv/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSrvServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.UserSrv",
	HandlerType: (*UserSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserSrv_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserSrv_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
