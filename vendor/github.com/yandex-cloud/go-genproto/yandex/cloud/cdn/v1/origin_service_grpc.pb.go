// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/cdn/v1/origin_service.proto

package cdn

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OriginService_Get_FullMethodName    = "/yandex.cloud.cdn.v1.OriginService/Get"
	OriginService_List_FullMethodName   = "/yandex.cloud.cdn.v1.OriginService/List"
	OriginService_Create_FullMethodName = "/yandex.cloud.cdn.v1.OriginService/Create"
	OriginService_Update_FullMethodName = "/yandex.cloud.cdn.v1.OriginService/Update"
	OriginService_Delete_FullMethodName = "/yandex.cloud.cdn.v1.OriginService/Delete"
)

// OriginServiceClient is the client API for OriginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OriginServiceClient interface {
	// Get origin in origin group.
	Get(ctx context.Context, in *GetOriginRequest, opts ...grpc.CallOption) (*Origin, error)
	// Lists origins of origin group.
	List(ctx context.Context, in *ListOriginsRequest, opts ...grpc.CallOption) (*ListOriginsResponse, error)
	// Creates origin inside origin group.
	Create(ctx context.Context, in *CreateOriginRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified origin from the origin group.
	//
	// Changes may take up to 15 minutes to apply. Afterwards, it is recommended to purge cache of the resources that
	// use the origin via a [CacheService.Purge] request.
	Update(ctx context.Context, in *UpdateOriginRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes origin from origin group.
	Delete(ctx context.Context, in *DeleteOriginRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type originServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOriginServiceClient(cc grpc.ClientConnInterface) OriginServiceClient {
	return &originServiceClient{cc}
}

func (c *originServiceClient) Get(ctx context.Context, in *GetOriginRequest, opts ...grpc.CallOption) (*Origin, error) {
	out := new(Origin)
	err := c.cc.Invoke(ctx, OriginService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originServiceClient) List(ctx context.Context, in *ListOriginsRequest, opts ...grpc.CallOption) (*ListOriginsResponse, error) {
	out := new(ListOriginsResponse)
	err := c.cc.Invoke(ctx, OriginService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originServiceClient) Create(ctx context.Context, in *CreateOriginRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, OriginService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originServiceClient) Update(ctx context.Context, in *UpdateOriginRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, OriginService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originServiceClient) Delete(ctx context.Context, in *DeleteOriginRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, OriginService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OriginServiceServer is the server API for OriginService service.
// All implementations should embed UnimplementedOriginServiceServer
// for forward compatibility
type OriginServiceServer interface {
	// Get origin in origin group.
	Get(context.Context, *GetOriginRequest) (*Origin, error)
	// Lists origins of origin group.
	List(context.Context, *ListOriginsRequest) (*ListOriginsResponse, error)
	// Creates origin inside origin group.
	Create(context.Context, *CreateOriginRequest) (*operation.Operation, error)
	// Updates the specified origin from the origin group.
	//
	// Changes may take up to 15 minutes to apply. Afterwards, it is recommended to purge cache of the resources that
	// use the origin via a [CacheService.Purge] request.
	Update(context.Context, *UpdateOriginRequest) (*operation.Operation, error)
	// Deletes origin from origin group.
	Delete(context.Context, *DeleteOriginRequest) (*operation.Operation, error)
}

// UnimplementedOriginServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOriginServiceServer struct {
}

func (UnimplementedOriginServiceServer) Get(context.Context, *GetOriginRequest) (*Origin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOriginServiceServer) List(context.Context, *ListOriginsRequest) (*ListOriginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOriginServiceServer) Create(context.Context, *CreateOriginRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOriginServiceServer) Update(context.Context, *UpdateOriginRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOriginServiceServer) Delete(context.Context, *DeleteOriginRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeOriginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OriginServiceServer will
// result in compilation errors.
type UnsafeOriginServiceServer interface {
	mustEmbedUnimplementedOriginServiceServer()
}

func RegisterOriginServiceServer(s grpc.ServiceRegistrar, srv OriginServiceServer) {
	s.RegisterService(&OriginService_ServiceDesc, srv)
}

func _OriginService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOriginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OriginService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginServiceServer).Get(ctx, req.(*GetOriginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OriginService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOriginsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OriginService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginServiceServer).List(ctx, req.(*ListOriginsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OriginService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOriginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OriginService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginServiceServer).Create(ctx, req.(*CreateOriginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OriginService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOriginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OriginService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginServiceServer).Update(ctx, req.(*UpdateOriginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OriginService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOriginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OriginService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginServiceServer).Delete(ctx, req.(*DeleteOriginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OriginService_ServiceDesc is the grpc.ServiceDesc for OriginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OriginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.cdn.v1.OriginService",
	HandlerType: (*OriginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _OriginService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _OriginService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _OriginService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OriginService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OriginService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/cdn/v1/origin_service.proto",
}
