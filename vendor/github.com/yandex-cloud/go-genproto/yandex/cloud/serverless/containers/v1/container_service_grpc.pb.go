// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/serverless/containers/v1/container_service.proto

package containers

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

// ContainerServiceClient is the client API for ContainerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerServiceClient interface {
	Get(ctx context.Context, in *GetContainerRequest, opts ...grpc.CallOption) (*Container, error)
	List(ctx context.Context, in *ListContainersRequest, opts ...grpc.CallOption) (*ListContainersResponse, error)
	Create(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	Update(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	Delete(ctx context.Context, in *DeleteContainerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	DeployRevision(ctx context.Context, in *DeployContainerRevisionRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	Rollback(ctx context.Context, in *RollbackContainerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	GetRevision(ctx context.Context, in *GetContainerRevisionRequest, opts ...grpc.CallOption) (*Revision, error)
	ListRevisions(ctx context.Context, in *ListContainersRevisionsRequest, opts ...grpc.CallOption) (*ListContainersRevisionsResponse, error)
	ListOperations(ctx context.Context, in *ListContainerOperationsRequest, opts ...grpc.CallOption) (*ListContainerOperationsResponse, error)
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type containerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerServiceClient(cc grpc.ClientConnInterface) ContainerServiceClient {
	return &containerServiceClient{cc}
}

func (c *containerServiceClient) Get(ctx context.Context, in *GetContainerRequest, opts ...grpc.CallOption) (*Container, error) {
	out := new(Container)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) List(ctx context.Context, in *ListContainersRequest, opts ...grpc.CallOption) (*ListContainersResponse, error) {
	out := new(ListContainersResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) Create(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) Update(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) Delete(ctx context.Context, in *DeleteContainerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) DeployRevision(ctx context.Context, in *DeployContainerRevisionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/DeployRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) Rollback(ctx context.Context, in *RollbackContainerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/Rollback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) GetRevision(ctx context.Context, in *GetContainerRevisionRequest, opts ...grpc.CallOption) (*Revision, error) {
	out := new(Revision)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/GetRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) ListRevisions(ctx context.Context, in *ListContainersRevisionsRequest, opts ...grpc.CallOption) (*ListContainersRevisionsResponse, error) {
	out := new(ListContainersRevisionsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/ListRevisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) ListOperations(ctx context.Context, in *ListContainerOperationsRequest, opts ...grpc.CallOption) (*ListContainerOperationsResponse, error) {
	out := new(ListContainerOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/ListAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/SetAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.containers.v1.ContainerService/UpdateAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContainerServiceServer is the server API for ContainerService service.
// All implementations should embed UnimplementedContainerServiceServer
// for forward compatibility
type ContainerServiceServer interface {
	Get(context.Context, *GetContainerRequest) (*Container, error)
	List(context.Context, *ListContainersRequest) (*ListContainersResponse, error)
	Create(context.Context, *CreateContainerRequest) (*operation.Operation, error)
	Update(context.Context, *UpdateContainerRequest) (*operation.Operation, error)
	Delete(context.Context, *DeleteContainerRequest) (*operation.Operation, error)
	DeployRevision(context.Context, *DeployContainerRevisionRequest) (*operation.Operation, error)
	Rollback(context.Context, *RollbackContainerRequest) (*operation.Operation, error)
	GetRevision(context.Context, *GetContainerRevisionRequest) (*Revision, error)
	ListRevisions(context.Context, *ListContainersRevisionsRequest) (*ListContainersRevisionsResponse, error)
	ListOperations(context.Context, *ListContainerOperationsRequest) (*ListContainerOperationsResponse, error)
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedContainerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedContainerServiceServer struct {
}

func (UnimplementedContainerServiceServer) Get(context.Context, *GetContainerRequest) (*Container, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedContainerServiceServer) List(context.Context, *ListContainersRequest) (*ListContainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedContainerServiceServer) Create(context.Context, *CreateContainerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedContainerServiceServer) Update(context.Context, *UpdateContainerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedContainerServiceServer) Delete(context.Context, *DeleteContainerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedContainerServiceServer) DeployRevision(context.Context, *DeployContainerRevisionRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployRevision not implemented")
}
func (UnimplementedContainerServiceServer) Rollback(context.Context, *RollbackContainerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rollback not implemented")
}
func (UnimplementedContainerServiceServer) GetRevision(context.Context, *GetContainerRevisionRequest) (*Revision, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRevision not implemented")
}
func (UnimplementedContainerServiceServer) ListRevisions(context.Context, *ListContainersRevisionsRequest) (*ListContainersRevisionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRevisions not implemented")
}
func (UnimplementedContainerServiceServer) ListOperations(context.Context, *ListContainerOperationsRequest) (*ListContainerOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedContainerServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedContainerServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedContainerServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}

// UnsafeContainerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerServiceServer will
// result in compilation errors.
type UnsafeContainerServiceServer interface {
	mustEmbedUnimplementedContainerServiceServer()
}

func RegisterContainerServiceServer(s grpc.ServiceRegistrar, srv ContainerServiceServer) {
	s.RegisterService(&ContainerService_ServiceDesc, srv)
}

func _ContainerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).Get(ctx, req.(*GetContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).List(ctx, req.(*ListContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).Create(ctx, req.(*CreateContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).Update(ctx, req.(*UpdateContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).Delete(ctx, req.(*DeleteContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_DeployRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployContainerRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).DeployRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/DeployRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).DeployRevision(ctx, req.(*DeployContainerRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_Rollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).Rollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/Rollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).Rollback(ctx, req.(*RollbackContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_GetRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainerRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).GetRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/GetRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).GetRevision(ctx, req.(*GetContainerRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_ListRevisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContainersRevisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).ListRevisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/ListRevisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).ListRevisions(ctx, req.(*ListContainersRevisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContainerOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).ListOperations(ctx, req.(*ListContainerOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/ListAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/SetAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.containers.v1.ContainerService/UpdateAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContainerService_ServiceDesc is the grpc.ServiceDesc for ContainerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContainerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.serverless.containers.v1.ContainerService",
	HandlerType: (*ContainerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ContainerService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ContainerService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ContainerService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ContainerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ContainerService_Delete_Handler,
		},
		{
			MethodName: "DeployRevision",
			Handler:    _ContainerService_DeployRevision_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _ContainerService_Rollback_Handler,
		},
		{
			MethodName: "GetRevision",
			Handler:    _ContainerService_GetRevision_Handler,
		},
		{
			MethodName: "ListRevisions",
			Handler:    _ContainerService_ListRevisions_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _ContainerService_ListOperations_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _ContainerService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _ContainerService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _ContainerService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/serverless/containers/v1/container_service.proto",
}
