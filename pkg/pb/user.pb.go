// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	UserType             string   `protobuf:"bytes,3,opt,name=userType,proto3" json:"userType,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LoginReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReply) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

func (m *LoginReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetPermissionsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPermissionsRequest) Reset()         { *m = GetPermissionsRequest{} }
func (m *GetPermissionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPermissionsRequest) ProtoMessage()    {}
func (*GetPermissionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *GetPermissionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPermissionsRequest.Unmarshal(m, b)
}
func (m *GetPermissionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPermissionsRequest.Marshal(b, m, deterministic)
}
func (m *GetPermissionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPermissionsRequest.Merge(m, src)
}
func (m *GetPermissionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetPermissionsRequest.Size(m)
}
func (m *GetPermissionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPermissionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPermissionsRequest proto.InternalMessageInfo

type GetPermissionsReply struct {
	Menus                []*Menu  `protobuf:"bytes,2,rep,name=menus,proto3" json:"menus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPermissionsReply) Reset()         { *m = GetPermissionsReply{} }
func (m *GetPermissionsReply) String() string { return proto.CompactTextString(m) }
func (*GetPermissionsReply) ProtoMessage()    {}
func (*GetPermissionsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *GetPermissionsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPermissionsReply.Unmarshal(m, b)
}
func (m *GetPermissionsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPermissionsReply.Marshal(b, m, deterministic)
}
func (m *GetPermissionsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPermissionsReply.Merge(m, src)
}
func (m *GetPermissionsReply) XXX_Size() int {
	return xxx_messageInfo_GetPermissionsReply.Size(m)
}
func (m *GetPermissionsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPermissionsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetPermissionsReply proto.InternalMessageInfo

func (m *GetPermissionsReply) GetMenus() []*Menu {
	if m != nil {
		return m.Menus
	}
	return nil
}

type Menu struct {
	Menu                 string   `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
	Route                string   `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Children             []*Menu  `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Menu) Reset()         { *m = Menu{} }
func (m *Menu) String() string { return proto.CompactTextString(m) }
func (*Menu) ProtoMessage()    {}
func (*Menu) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *Menu) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Menu.Unmarshal(m, b)
}
func (m *Menu) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Menu.Marshal(b, m, deterministic)
}
func (m *Menu) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Menu.Merge(m, src)
}
func (m *Menu) XXX_Size() int {
	return xxx_messageInfo_Menu.Size(m)
}
func (m *Menu) XXX_DiscardUnknown() {
	xxx_messageInfo_Menu.DiscardUnknown(m)
}

var xxx_messageInfo_Menu proto.InternalMessageInfo

func (m *Menu) GetMenu() string {
	if m != nil {
		return m.Menu
	}
	return ""
}

func (m *Menu) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *Menu) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Menu) GetChildren() []*Menu {
	if m != nil {
		return m.Children
	}
	return nil
}

type CheckPermissionRequest struct {
	Route                string   `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPermissionRequest) Reset()         { *m = CheckPermissionRequest{} }
func (m *CheckPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*CheckPermissionRequest) ProtoMessage()    {}
func (*CheckPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *CheckPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPermissionRequest.Unmarshal(m, b)
}
func (m *CheckPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPermissionRequest.Marshal(b, m, deterministic)
}
func (m *CheckPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPermissionRequest.Merge(m, src)
}
func (m *CheckPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_CheckPermissionRequest.Size(m)
}
func (m *CheckPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPermissionRequest proto.InternalMessageInfo

func (m *CheckPermissionRequest) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

type CheckPermissionReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPermissionReply) Reset()         { *m = CheckPermissionReply{} }
func (m *CheckPermissionReply) String() string { return proto.CompactTextString(m) }
func (*CheckPermissionReply) ProtoMessage()    {}
func (*CheckPermissionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *CheckPermissionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPermissionReply.Unmarshal(m, b)
}
func (m *CheckPermissionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPermissionReply.Marshal(b, m, deterministic)
}
func (m *CheckPermissionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPermissionReply.Merge(m, src)
}
func (m *CheckPermissionReply) XXX_Size() int {
	return xxx_messageInfo_CheckPermissionReply.Size(m)
}
func (m *CheckPermissionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPermissionReply.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPermissionReply proto.InternalMessageInfo

func (m *CheckPermissionReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "pb.LoginReply")
	proto.RegisterType((*GetPermissionsRequest)(nil), "pb.GetPermissionsRequest")
	proto.RegisterType((*GetPermissionsReply)(nil), "pb.GetPermissionsReply")
	proto.RegisterType((*Menu)(nil), "pb.Menu")
	proto.RegisterType((*CheckPermissionRequest)(nil), "pb.CheckPermissionRequest")
	proto.RegisterType((*CheckPermissionReply)(nil), "pb.CheckPermissionReply")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xdd, 0x4a, 0xf3, 0x40,
	0x10, 0xfd, 0x92, 0xa6, 0x25, 0xdf, 0x28, 0x55, 0xd6, 0xda, 0xc6, 0x5c, 0x48, 0x59, 0xbc, 0x28,
	0x08, 0xb9, 0xa8, 0xf8, 0x04, 0x42, 0x45, 0x50, 0x90, 0xa0, 0x0f, 0x90, 0x36, 0x83, 0x4d, 0x9b,
	0x66, 0xd7, 0xdd, 0x0d, 0xd2, 0x17, 0xf4, 0xb9, 0x64, 0x36, 0xe9, 0x7f, 0xee, 0xe6, 0xcc, 0x39,
	0x39, 0x73, 0x66, 0x36, 0x00, 0xa5, 0x46, 0x15, 0x49, 0x25, 0x8c, 0x60, 0xae, 0x9c, 0xf2, 0x09,
	0x9c, 0xbf, 0x8a, 0xaf, 0xac, 0x88, 0xf1, 0xbb, 0x44, 0x6d, 0x58, 0x08, 0x3e, 0x29, 0x8a, 0x64,
	0x85, 0x81, 0x33, 0x74, 0x46, 0xff, 0xe3, 0x2d, 0x26, 0x4e, 0x26, 0x5a, 0xff, 0x08, 0x95, 0x06,
	0x6e, 0xc5, 0x6d, 0x30, 0x5f, 0x00, 0xd4, 0x3e, 0x32, 0x5f, 0xb3, 0x2e, 0xb8, 0x59, 0x6a, 0xbf,
	0x6f, 0xc5, 0x6e, 0x96, 0x1e, 0xb8, 0xba, 0xa7, 0xae, 0x54, 0x7f, 0xac, 0x25, 0x06, 0xad, 0x1d,
	0x47, 0x98, 0xf5, 0xa0, 0x6d, 0xc4, 0x12, 0x8b, 0xc0, 0xb3, 0x44, 0x05, 0xf8, 0x00, 0xae, 0x9f,
	0xd1, 0xbc, 0xa3, 0x5a, 0x65, 0x5a, 0x67, 0xa2, 0xd0, 0x75, 0x78, 0xfe, 0x08, 0x57, 0xc7, 0x04,
	0xa5, 0xb9, 0x85, 0xf6, 0x0a, 0x8b, 0x52, 0x07, 0xee, 0xb0, 0x35, 0x3a, 0x1b, 0xfb, 0x91, 0x9c,
	0x46, 0x6f, 0x58, 0x94, 0x71, 0xd5, 0xe6, 0x0b, 0xf0, 0x08, 0x32, 0x06, 0x1e, 0x35, 0xea, 0xbd,
	0x6d, 0x4d, 0x09, 0x94, 0x28, 0xcd, 0x26, 0x76, 0x05, 0x48, 0x99, 0x26, 0x26, 0xa9, 0xf3, 0xda,
	0x9a, 0xdd, 0x81, 0x3f, 0x9b, 0x67, 0x79, 0xaa, 0x6c, 0xdc, 0xc3, 0x41, 0x5b, 0x86, 0x47, 0xd0,
	0x7f, 0x9a, 0xe3, 0x6c, 0xb9, 0x0b, 0xb9, 0xb9, 0xfc, 0x76, 0x92, 0xb3, 0x37, 0x89, 0x47, 0xd0,
	0x3b, 0xd1, 0xd3, 0x4e, 0x7d, 0xe8, 0x28, 0xd4, 0x65, 0x6e, 0xac, 0xdc, 0x8f, 0x6b, 0x34, 0xfe,
	0x75, 0xc0, 0xfb, 0xd4, 0xa8, 0xd8, 0x3d, 0xb4, 0xed, 0x83, 0xb0, 0x4b, 0x4a, 0xb1, 0xff, 0xc6,
	0x61, 0x77, 0xaf, 0x23, 0xf3, 0x35, 0xff, 0xc7, 0x26, 0xd0, 0x3d, 0x3c, 0x1c, 0xbb, 0x21, 0x4d,
	0xe3, 0x95, 0xc3, 0x41, 0x13, 0x55, 0xf9, 0xbc, 0xc0, 0xc5, 0x51, 0x5a, 0x16, 0x92, 0xba, 0x79,
	0xe5, 0x30, 0x68, 0xe4, 0xac, 0xd5, 0xb4, 0x63, 0xff, 0xd1, 0x87, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xeb, 0x88, 0x68, 0xf9, 0xb1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsReply, error)
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionReply, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/pb.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsReply, error) {
	out := new(GetPermissionsReply)
	err := c.cc.Invoke(ctx, "/pb.User/GetPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionReply, error) {
	out := new(CheckPermissionReply)
	err := c.cc.Invoke(ctx, "/pb.User/CheckPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsReply, error)
	CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionReply, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServer) GetPermissions(ctx context.Context, req *GetPermissionsRequest) (*GetPermissionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}
func (*UnimplementedUserServer) CheckPermission(ctx context.Context, req *CheckPermissionRequest) (*CheckPermissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/GetPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetPermissions(ctx, req.(*GetPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/CheckPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _User_GetPermissions_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _User_CheckPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
