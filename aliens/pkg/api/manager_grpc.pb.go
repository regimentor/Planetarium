// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: protobuf/manager.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ManagerService_Create_FullMethodName = "/api.ManagerService/Create"
	ManagerService_Update_FullMethodName = "/api.ManagerService/Update"
	ManagerService_Find_FullMethodName   = "/api.ManagerService/Find"
)

// ManagerServiceClient is the client API for ManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerServiceClient interface {
	Create(ctx context.Context, in *CreateAlienRequest, opts ...grpc.CallOption) (*CreateAlienResponse, error)
	Update(ctx context.Context, in *UpdateAlienRequest, opts ...grpc.CallOption) (*UpdateAlienResponse, error)
	Find(ctx context.Context, in *FindAlienRequest, opts ...grpc.CallOption) (*FindAlienResponse, error)
}

type managerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerServiceClient(cc grpc.ClientConnInterface) ManagerServiceClient {
	return &managerServiceClient{cc}
}

func (c *managerServiceClient) Create(ctx context.Context, in *CreateAlienRequest, opts ...grpc.CallOption) (*CreateAlienResponse, error) {
	out := new(CreateAlienResponse)
	err := c.cc.Invoke(ctx, ManagerService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServiceClient) Update(ctx context.Context, in *UpdateAlienRequest, opts ...grpc.CallOption) (*UpdateAlienResponse, error) {
	out := new(UpdateAlienResponse)
	err := c.cc.Invoke(ctx, ManagerService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServiceClient) Find(ctx context.Context, in *FindAlienRequest, opts ...grpc.CallOption) (*FindAlienResponse, error) {
	out := new(FindAlienResponse)
	err := c.cc.Invoke(ctx, ManagerService_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServiceServer is the server API for ManagerService service.
// All implementations must embed UnimplementedManagerServiceServer
// for forward compatibility
type ManagerServiceServer interface {
	Create(context.Context, *CreateAlienRequest) (*CreateAlienResponse, error)
	Update(context.Context, *UpdateAlienRequest) (*UpdateAlienResponse, error)
	Find(context.Context, *FindAlienRequest) (*FindAlienResponse, error)
	mustEmbedUnimplementedManagerServiceServer()
}

// UnimplementedManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServiceServer struct {
}

func (UnimplementedManagerServiceServer) Create(context.Context, *CreateAlienRequest) (*CreateAlienResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedManagerServiceServer) Update(context.Context, *UpdateAlienRequest) (*UpdateAlienResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedManagerServiceServer) Find(context.Context, *FindAlienRequest) (*FindAlienResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedManagerServiceServer) mustEmbedUnimplementedManagerServiceServer() {}

// UnsafeManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServiceServer will
// result in compilation errors.
type UnsafeManagerServiceServer interface {
	mustEmbedUnimplementedManagerServiceServer()
}

func RegisterManagerServiceServer(s grpc.ServiceRegistrar, srv ManagerServiceServer) {
	s.RegisterService(&ManagerService_ServiceDesc, srv)
}

func _ManagerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlienRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).Create(ctx, req.(*CreateAlienRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlienRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).Update(ctx, req.(*UpdateAlienRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAlienRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).Find(ctx, req.(*FindAlienRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagerService_ServiceDesc is the grpc.ServiceDesc for ManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ManagerService",
	HandlerType: (*ManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ManagerService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ManagerService_Update_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _ManagerService_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/manager.proto",
}
