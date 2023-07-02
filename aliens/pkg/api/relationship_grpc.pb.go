// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: protobuf/relationship.proto

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
	RelationService_Create_FullMethodName  = "/api.RelationService/Create"
	RelationService_Submit_FullMethodName  = "/api.RelationService/Submit"
	RelationService_Decline_FullMethodName = "/api.RelationService/Decline"
)

// RelationServiceClient is the client API for RelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationServiceClient interface {
	Create(ctx context.Context, in *CreateRelationRequestRequest, opts ...grpc.CallOption) (*CreateRelationRequestResponse, error)
	Submit(ctx context.Context, in *SubmitRelationRequest, opts ...grpc.CallOption) (*SubmitRelationResponse, error)
	Decline(ctx context.Context, in *DeclineRelationRequest, opts ...grpc.CallOption) (*DeclineRelationResponse, error)
}

type relationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationServiceClient(cc grpc.ClientConnInterface) RelationServiceClient {
	return &relationServiceClient{cc}
}

func (c *relationServiceClient) Create(ctx context.Context, in *CreateRelationRequestRequest, opts ...grpc.CallOption) (*CreateRelationRequestResponse, error) {
	out := new(CreateRelationRequestResponse)
	err := c.cc.Invoke(ctx, RelationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) Submit(ctx context.Context, in *SubmitRelationRequest, opts ...grpc.CallOption) (*SubmitRelationResponse, error) {
	out := new(SubmitRelationResponse)
	err := c.cc.Invoke(ctx, RelationService_Submit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) Decline(ctx context.Context, in *DeclineRelationRequest, opts ...grpc.CallOption) (*DeclineRelationResponse, error) {
	out := new(DeclineRelationResponse)
	err := c.cc.Invoke(ctx, RelationService_Decline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServiceServer is the server API for RelationService service.
// All implementations must embed UnimplementedRelationServiceServer
// for forward compatibility
type RelationServiceServer interface {
	Create(context.Context, *CreateRelationRequestRequest) (*CreateRelationRequestResponse, error)
	Submit(context.Context, *SubmitRelationRequest) (*SubmitRelationResponse, error)
	Decline(context.Context, *DeclineRelationRequest) (*DeclineRelationResponse, error)
	mustEmbedUnimplementedRelationServiceServer()
}

// UnimplementedRelationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRelationServiceServer struct {
}

func (UnimplementedRelationServiceServer) Create(context.Context, *CreateRelationRequestRequest) (*CreateRelationRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRelationServiceServer) Submit(context.Context, *SubmitRelationRequest) (*SubmitRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedRelationServiceServer) Decline(context.Context, *DeclineRelationRequest) (*DeclineRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decline not implemented")
}
func (UnimplementedRelationServiceServer) mustEmbedUnimplementedRelationServiceServer() {}

// UnsafeRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServiceServer will
// result in compilation errors.
type UnsafeRelationServiceServer interface {
	mustEmbedUnimplementedRelationServiceServer()
}

func RegisterRelationServiceServer(s grpc.ServiceRegistrar, srv RelationServiceServer) {
	s.RegisterService(&RelationService_ServiceDesc, srv)
}

func _RelationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelationRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).Create(ctx, req.(*CreateRelationRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_Submit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).Submit(ctx, req.(*SubmitRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_Decline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclineRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).Decline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_Decline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).Decline(ctx, req.(*DeclineRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationService_ServiceDesc is the grpc.ServiceDesc for RelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.RelationService",
	HandlerType: (*RelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RelationService_Create_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _RelationService_Submit_Handler,
		},
		{
			MethodName: "Decline",
			Handler:    _RelationService_Decline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/relationship.proto",
}