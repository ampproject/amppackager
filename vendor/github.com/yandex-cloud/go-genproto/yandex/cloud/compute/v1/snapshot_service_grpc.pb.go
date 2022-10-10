// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/compute/v1/snapshot_service.proto

package compute

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

// SnapshotServiceClient is the client API for SnapshotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnapshotServiceClient interface {
	// Returns the specified Snapshot resource.
	//
	// To get the list of available Snapshot resources, make a [List] request.
	Get(ctx context.Context, in *GetSnapshotRequest, opts ...grpc.CallOption) (*Snapshot, error)
	// Retrieves the list of Snapshot resources in the specified folder.
	List(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error)
	// Creates a snapshot of the specified disk.
	Create(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified snapshot.
	//
	// Values of omitted parameters are not changed.
	Update(ctx context.Context, in *UpdateSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified snapshot.
	//
	// Deleting a snapshot removes its data permanently and is irreversible.
	Delete(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified snapshot.
	ListOperations(ctx context.Context, in *ListSnapshotOperationsRequest, opts ...grpc.CallOption) (*ListSnapshotOperationsResponse, error)
}

type snapshotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSnapshotServiceClient(cc grpc.ClientConnInterface) SnapshotServiceClient {
	return &snapshotServiceClient{cc}
}

func (c *snapshotServiceClient) Get(ctx context.Context, in *GetSnapshotRequest, opts ...grpc.CallOption) (*Snapshot, error) {
	out := new(Snapshot)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.SnapshotService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) List(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error) {
	out := new(ListSnapshotsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.SnapshotService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) Create(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.SnapshotService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) Update(ctx context.Context, in *UpdateSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.SnapshotService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) Delete(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.SnapshotService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) ListOperations(ctx context.Context, in *ListSnapshotOperationsRequest, opts ...grpc.CallOption) (*ListSnapshotOperationsResponse, error) {
	out := new(ListSnapshotOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.SnapshotService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnapshotServiceServer is the server API for SnapshotService service.
// All implementations should embed UnimplementedSnapshotServiceServer
// for forward compatibility
type SnapshotServiceServer interface {
	// Returns the specified Snapshot resource.
	//
	// To get the list of available Snapshot resources, make a [List] request.
	Get(context.Context, *GetSnapshotRequest) (*Snapshot, error)
	// Retrieves the list of Snapshot resources in the specified folder.
	List(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error)
	// Creates a snapshot of the specified disk.
	Create(context.Context, *CreateSnapshotRequest) (*operation.Operation, error)
	// Updates the specified snapshot.
	//
	// Values of omitted parameters are not changed.
	Update(context.Context, *UpdateSnapshotRequest) (*operation.Operation, error)
	// Deletes the specified snapshot.
	//
	// Deleting a snapshot removes its data permanently and is irreversible.
	Delete(context.Context, *DeleteSnapshotRequest) (*operation.Operation, error)
	// Lists operations for the specified snapshot.
	ListOperations(context.Context, *ListSnapshotOperationsRequest) (*ListSnapshotOperationsResponse, error)
}

// UnimplementedSnapshotServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSnapshotServiceServer struct {
}

func (UnimplementedSnapshotServiceServer) Get(context.Context, *GetSnapshotRequest) (*Snapshot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSnapshotServiceServer) List(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSnapshotServiceServer) Create(context.Context, *CreateSnapshotRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSnapshotServiceServer) Update(context.Context, *UpdateSnapshotRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSnapshotServiceServer) Delete(context.Context, *DeleteSnapshotRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSnapshotServiceServer) ListOperations(context.Context, *ListSnapshotOperationsRequest) (*ListSnapshotOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}

// UnsafeSnapshotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnapshotServiceServer will
// result in compilation errors.
type UnsafeSnapshotServiceServer interface {
	mustEmbedUnimplementedSnapshotServiceServer()
}

func RegisterSnapshotServiceServer(s grpc.ServiceRegistrar, srv SnapshotServiceServer) {
	s.RegisterService(&SnapshotService_ServiceDesc, srv)
}

func _SnapshotService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.SnapshotService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).Get(ctx, req.(*GetSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.SnapshotService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).List(ctx, req.(*ListSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.SnapshotService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).Create(ctx, req.(*CreateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.SnapshotService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).Update(ctx, req.(*UpdateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.SnapshotService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).Delete(ctx, req.(*DeleteSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.SnapshotService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).ListOperations(ctx, req.(*ListSnapshotOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SnapshotService_ServiceDesc is the grpc.ServiceDesc for SnapshotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SnapshotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.compute.v1.SnapshotService",
	HandlerType: (*SnapshotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SnapshotService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SnapshotService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SnapshotService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SnapshotService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SnapshotService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _SnapshotService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/compute/v1/snapshot_service.proto",
}
