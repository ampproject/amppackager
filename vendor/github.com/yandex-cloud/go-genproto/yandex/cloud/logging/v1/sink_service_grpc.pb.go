// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/logging/v1/sink_service.proto

package logging

import (
	context "context"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
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
	SinkService_Get_FullMethodName                  = "/yandex.cloud.logging.v1.SinkService/Get"
	SinkService_List_FullMethodName                 = "/yandex.cloud.logging.v1.SinkService/List"
	SinkService_Create_FullMethodName               = "/yandex.cloud.logging.v1.SinkService/Create"
	SinkService_Update_FullMethodName               = "/yandex.cloud.logging.v1.SinkService/Update"
	SinkService_Delete_FullMethodName               = "/yandex.cloud.logging.v1.SinkService/Delete"
	SinkService_ListOperations_FullMethodName       = "/yandex.cloud.logging.v1.SinkService/ListOperations"
	SinkService_ListAccessBindings_FullMethodName   = "/yandex.cloud.logging.v1.SinkService/ListAccessBindings"
	SinkService_SetAccessBindings_FullMethodName    = "/yandex.cloud.logging.v1.SinkService/SetAccessBindings"
	SinkService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.logging.v1.SinkService/UpdateAccessBindings"
)

// SinkServiceClient is the client API for SinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SinkServiceClient interface {
	// Returns the specified sink.
	//
	// To get the list of all available sinks, make a [List] request.
	Get(ctx context.Context, in *GetSinkRequest, opts ...grpc.CallOption) (*Sink, error)
	// Retrieves the list of sinks in the specified folder.
	List(ctx context.Context, in *ListSinksRequest, opts ...grpc.CallOption) (*ListSinksResponse, error)
	// Creates a sink in the specified folder.
	Create(ctx context.Context, in *CreateSinkRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified sink.
	Update(ctx context.Context, in *UpdateSinkRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified sink.
	Delete(ctx context.Context, in *DeleteSinkRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified sink.
	ListOperations(ctx context.Context, in *ListSinkOperationsRequest, opts ...grpc.CallOption) (*ListSinkOperationsResponse, error)
	// Lists existing access bindings for the specified sink.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified sink.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified sink.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type sinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSinkServiceClient(cc grpc.ClientConnInterface) SinkServiceClient {
	return &sinkServiceClient{cc}
}

func (c *sinkServiceClient) Get(ctx context.Context, in *GetSinkRequest, opts ...grpc.CallOption) (*Sink, error) {
	out := new(Sink)
	err := c.cc.Invoke(ctx, SinkService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sinkServiceClient) List(ctx context.Context, in *ListSinksRequest, opts ...grpc.CallOption) (*ListSinksResponse, error) {
	out := new(ListSinksResponse)
	err := c.cc.Invoke(ctx, SinkService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sinkServiceClient) Create(ctx context.Context, in *CreateSinkRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SinkService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sinkServiceClient) Update(ctx context.Context, in *UpdateSinkRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SinkService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sinkServiceClient) Delete(ctx context.Context, in *DeleteSinkRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SinkService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sinkServiceClient) ListOperations(ctx context.Context, in *ListSinkOperationsRequest, opts ...grpc.CallOption) (*ListSinkOperationsResponse, error) {
	out := new(ListSinkOperationsResponse)
	err := c.cc.Invoke(ctx, SinkService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sinkServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, SinkService_ListAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sinkServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SinkService_SetAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sinkServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SinkService_UpdateAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SinkServiceServer is the server API for SinkService service.
// All implementations should embed UnimplementedSinkServiceServer
// for forward compatibility
type SinkServiceServer interface {
	// Returns the specified sink.
	//
	// To get the list of all available sinks, make a [List] request.
	Get(context.Context, *GetSinkRequest) (*Sink, error)
	// Retrieves the list of sinks in the specified folder.
	List(context.Context, *ListSinksRequest) (*ListSinksResponse, error)
	// Creates a sink in the specified folder.
	Create(context.Context, *CreateSinkRequest) (*operation.Operation, error)
	// Updates the specified sink.
	Update(context.Context, *UpdateSinkRequest) (*operation.Operation, error)
	// Deletes the specified sink.
	Delete(context.Context, *DeleteSinkRequest) (*operation.Operation, error)
	// Lists operations for the specified sink.
	ListOperations(context.Context, *ListSinkOperationsRequest) (*ListSinkOperationsResponse, error)
	// Lists existing access bindings for the specified sink.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified sink.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified sink.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedSinkServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSinkServiceServer struct {
}

func (UnimplementedSinkServiceServer) Get(context.Context, *GetSinkRequest) (*Sink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSinkServiceServer) List(context.Context, *ListSinksRequest) (*ListSinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSinkServiceServer) Create(context.Context, *CreateSinkRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSinkServiceServer) Update(context.Context, *UpdateSinkRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSinkServiceServer) Delete(context.Context, *DeleteSinkRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSinkServiceServer) ListOperations(context.Context, *ListSinkOperationsRequest) (*ListSinkOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedSinkServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedSinkServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedSinkServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}

// UnsafeSinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SinkServiceServer will
// result in compilation errors.
type UnsafeSinkServiceServer interface {
	mustEmbedUnimplementedSinkServiceServer()
}

func RegisterSinkServiceServer(s grpc.ServiceRegistrar, srv SinkServiceServer) {
	s.RegisterService(&SinkService_ServiceDesc, srv)
}

func _SinkService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SinkService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkServiceServer).Get(ctx, req.(*GetSinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinkService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SinkService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkServiceServer).List(ctx, req.(*ListSinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinkService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SinkService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkServiceServer).Create(ctx, req.(*CreateSinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinkService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SinkService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkServiceServer).Update(ctx, req.(*UpdateSinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinkService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SinkService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkServiceServer).Delete(ctx, req.(*DeleteSinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinkService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSinkOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SinkService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkServiceServer).ListOperations(ctx, req.(*ListSinkOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinkService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SinkService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinkService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SinkService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinkService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SinkService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SinkService_ServiceDesc is the grpc.ServiceDesc for SinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.logging.v1.SinkService",
	HandlerType: (*SinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SinkService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SinkService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SinkService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SinkService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SinkService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _SinkService_ListOperations_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _SinkService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _SinkService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _SinkService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/logging/v1/sink_service.proto",
}
