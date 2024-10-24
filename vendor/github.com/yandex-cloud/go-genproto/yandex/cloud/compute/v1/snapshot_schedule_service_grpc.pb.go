// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/compute/v1/snapshot_schedule_service.proto

package compute

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
	SnapshotScheduleService_Get_FullMethodName                  = "/yandex.cloud.compute.v1.SnapshotScheduleService/Get"
	SnapshotScheduleService_List_FullMethodName                 = "/yandex.cloud.compute.v1.SnapshotScheduleService/List"
	SnapshotScheduleService_Create_FullMethodName               = "/yandex.cloud.compute.v1.SnapshotScheduleService/Create"
	SnapshotScheduleService_Update_FullMethodName               = "/yandex.cloud.compute.v1.SnapshotScheduleService/Update"
	SnapshotScheduleService_Delete_FullMethodName               = "/yandex.cloud.compute.v1.SnapshotScheduleService/Delete"
	SnapshotScheduleService_UpdateDisks_FullMethodName          = "/yandex.cloud.compute.v1.SnapshotScheduleService/UpdateDisks"
	SnapshotScheduleService_Disable_FullMethodName              = "/yandex.cloud.compute.v1.SnapshotScheduleService/Disable"
	SnapshotScheduleService_Enable_FullMethodName               = "/yandex.cloud.compute.v1.SnapshotScheduleService/Enable"
	SnapshotScheduleService_ListOperations_FullMethodName       = "/yandex.cloud.compute.v1.SnapshotScheduleService/ListOperations"
	SnapshotScheduleService_ListSnapshots_FullMethodName        = "/yandex.cloud.compute.v1.SnapshotScheduleService/ListSnapshots"
	SnapshotScheduleService_ListDisks_FullMethodName            = "/yandex.cloud.compute.v1.SnapshotScheduleService/ListDisks"
	SnapshotScheduleService_ListAccessBindings_FullMethodName   = "/yandex.cloud.compute.v1.SnapshotScheduleService/ListAccessBindings"
	SnapshotScheduleService_SetAccessBindings_FullMethodName    = "/yandex.cloud.compute.v1.SnapshotScheduleService/SetAccessBindings"
	SnapshotScheduleService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.compute.v1.SnapshotScheduleService/UpdateAccessBindings"
)

// SnapshotScheduleServiceClient is the client API for SnapshotScheduleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnapshotScheduleServiceClient interface {
	// Returns the specified snapshot schedule.
	//
	// To get the list of available snapshot schedules, make a [List] request.
	Get(ctx context.Context, in *GetSnapshotScheduleRequest, opts ...grpc.CallOption) (*SnapshotSchedule, error)
	// Retrieves the list of snapshot schedules in the specified folder.
	List(ctx context.Context, in *ListSnapshotSchedulesRequest, opts ...grpc.CallOption) (*ListSnapshotSchedulesResponse, error)
	// Creates a snapshot schedule in the specified folder.
	Create(ctx context.Context, in *CreateSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified snapshot schedule.
	//
	// The schedule is updated only after all snapshot creations and deletions triggered by the schedule are completed.
	Update(ctx context.Context, in *UpdateSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified snapshot schedule.
	//
	// Deleting a snapshot schedule removes its data permanently and is irreversible. However, deleting a schedule
	// does not delete any snapshots created by the schedule. You must delete snapshots separately.
	//
	// The schedule is deleted only after all snapshot creations and deletions triggered by the schedule are completed.
	Delete(ctx context.Context, in *DeleteSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the list of disks attached to the specified schedule.
	//
	// The schedule is updated only after all snapshot creations and deletions triggered by the schedule are completed.
	UpdateDisks(ctx context.Context, in *UpdateSnapshotScheduleDisksRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Disables the specified snapshot schedule.
	//
	// The [SnapshotSchedule.status] is changed to `INACTIVE`: the schedule is interrupted, snapshots won't be created
	// or deleted.
	//
	// The schedule is disabled only after all snapshot creations and deletions triggered by the schedule are completed.
	Disable(ctx context.Context, in *DisableSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Enables the specified snapshot schedule.
	//
	// The [SnapshotSchedule.status] is changed to `ACTIVE`: new disk snapshots will be created, old ones deleted
	// (if [SnapshotSchedule.retention_policy] is specified).
	Enable(ctx context.Context, in *EnableSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified snapshot schedule.
	ListOperations(ctx context.Context, in *ListSnapshotScheduleOperationsRequest, opts ...grpc.CallOption) (*ListSnapshotScheduleOperationsResponse, error)
	// Retrieves the list of snapshots created by the specified snapshot schedule.
	ListSnapshots(ctx context.Context, in *ListSnapshotScheduleSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotScheduleSnapshotsResponse, error)
	// Retrieves the list of disks attached to the specified snapshot schedule.
	ListDisks(ctx context.Context, in *ListSnapshotScheduleDisksRequest, opts ...grpc.CallOption) (*ListSnapshotScheduleDisksResponse, error)
	// Lists access bindings for the snapshot schedule.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the snapshot schedule.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the snapshot schedule.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type snapshotScheduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSnapshotScheduleServiceClient(cc grpc.ClientConnInterface) SnapshotScheduleServiceClient {
	return &snapshotScheduleServiceClient{cc}
}

func (c *snapshotScheduleServiceClient) Get(ctx context.Context, in *GetSnapshotScheduleRequest, opts ...grpc.CallOption) (*SnapshotSchedule, error) {
	out := new(SnapshotSchedule)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) List(ctx context.Context, in *ListSnapshotSchedulesRequest, opts ...grpc.CallOption) (*ListSnapshotSchedulesResponse, error) {
	out := new(ListSnapshotSchedulesResponse)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) Create(ctx context.Context, in *CreateSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) Update(ctx context.Context, in *UpdateSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) Delete(ctx context.Context, in *DeleteSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) UpdateDisks(ctx context.Context, in *UpdateSnapshotScheduleDisksRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_UpdateDisks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) Disable(ctx context.Context, in *DisableSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) Enable(ctx context.Context, in *EnableSnapshotScheduleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) ListOperations(ctx context.Context, in *ListSnapshotScheduleOperationsRequest, opts ...grpc.CallOption) (*ListSnapshotScheduleOperationsResponse, error) {
	out := new(ListSnapshotScheduleOperationsResponse)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) ListSnapshots(ctx context.Context, in *ListSnapshotScheduleSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotScheduleSnapshotsResponse, error) {
	out := new(ListSnapshotScheduleSnapshotsResponse)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_ListSnapshots_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) ListDisks(ctx context.Context, in *ListSnapshotScheduleDisksRequest, opts ...grpc.CallOption) (*ListSnapshotScheduleDisksResponse, error) {
	out := new(ListSnapshotScheduleDisksResponse)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_ListDisks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_ListAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_SetAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotScheduleServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotScheduleService_UpdateAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnapshotScheduleServiceServer is the server API for SnapshotScheduleService service.
// All implementations should embed UnimplementedSnapshotScheduleServiceServer
// for forward compatibility
type SnapshotScheduleServiceServer interface {
	// Returns the specified snapshot schedule.
	//
	// To get the list of available snapshot schedules, make a [List] request.
	Get(context.Context, *GetSnapshotScheduleRequest) (*SnapshotSchedule, error)
	// Retrieves the list of snapshot schedules in the specified folder.
	List(context.Context, *ListSnapshotSchedulesRequest) (*ListSnapshotSchedulesResponse, error)
	// Creates a snapshot schedule in the specified folder.
	Create(context.Context, *CreateSnapshotScheduleRequest) (*operation.Operation, error)
	// Updates the specified snapshot schedule.
	//
	// The schedule is updated only after all snapshot creations and deletions triggered by the schedule are completed.
	Update(context.Context, *UpdateSnapshotScheduleRequest) (*operation.Operation, error)
	// Deletes the specified snapshot schedule.
	//
	// Deleting a snapshot schedule removes its data permanently and is irreversible. However, deleting a schedule
	// does not delete any snapshots created by the schedule. You must delete snapshots separately.
	//
	// The schedule is deleted only after all snapshot creations and deletions triggered by the schedule are completed.
	Delete(context.Context, *DeleteSnapshotScheduleRequest) (*operation.Operation, error)
	// Updates the list of disks attached to the specified schedule.
	//
	// The schedule is updated only after all snapshot creations and deletions triggered by the schedule are completed.
	UpdateDisks(context.Context, *UpdateSnapshotScheduleDisksRequest) (*operation.Operation, error)
	// Disables the specified snapshot schedule.
	//
	// The [SnapshotSchedule.status] is changed to `INACTIVE`: the schedule is interrupted, snapshots won't be created
	// or deleted.
	//
	// The schedule is disabled only after all snapshot creations and deletions triggered by the schedule are completed.
	Disable(context.Context, *DisableSnapshotScheduleRequest) (*operation.Operation, error)
	// Enables the specified snapshot schedule.
	//
	// The [SnapshotSchedule.status] is changed to `ACTIVE`: new disk snapshots will be created, old ones deleted
	// (if [SnapshotSchedule.retention_policy] is specified).
	Enable(context.Context, *EnableSnapshotScheduleRequest) (*operation.Operation, error)
	// Lists operations for the specified snapshot schedule.
	ListOperations(context.Context, *ListSnapshotScheduleOperationsRequest) (*ListSnapshotScheduleOperationsResponse, error)
	// Retrieves the list of snapshots created by the specified snapshot schedule.
	ListSnapshots(context.Context, *ListSnapshotScheduleSnapshotsRequest) (*ListSnapshotScheduleSnapshotsResponse, error)
	// Retrieves the list of disks attached to the specified snapshot schedule.
	ListDisks(context.Context, *ListSnapshotScheduleDisksRequest) (*ListSnapshotScheduleDisksResponse, error)
	// Lists access bindings for the snapshot schedule.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the snapshot schedule.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the snapshot schedule.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedSnapshotScheduleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSnapshotScheduleServiceServer struct {
}

func (UnimplementedSnapshotScheduleServiceServer) Get(context.Context, *GetSnapshotScheduleRequest) (*SnapshotSchedule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) List(context.Context, *ListSnapshotSchedulesRequest) (*ListSnapshotSchedulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) Create(context.Context, *CreateSnapshotScheduleRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) Update(context.Context, *UpdateSnapshotScheduleRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) Delete(context.Context, *DeleteSnapshotScheduleRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) UpdateDisks(context.Context, *UpdateSnapshotScheduleDisksRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDisks not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) Disable(context.Context, *DisableSnapshotScheduleRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) Enable(context.Context, *EnableSnapshotScheduleRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) ListOperations(context.Context, *ListSnapshotScheduleOperationsRequest) (*ListSnapshotScheduleOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) ListSnapshots(context.Context, *ListSnapshotScheduleSnapshotsRequest) (*ListSnapshotScheduleSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshots not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) ListDisks(context.Context, *ListSnapshotScheduleDisksRequest) (*ListSnapshotScheduleDisksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDisks not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedSnapshotScheduleServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}

// UnsafeSnapshotScheduleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnapshotScheduleServiceServer will
// result in compilation errors.
type UnsafeSnapshotScheduleServiceServer interface {
	mustEmbedUnimplementedSnapshotScheduleServiceServer()
}

func RegisterSnapshotScheduleServiceServer(s grpc.ServiceRegistrar, srv SnapshotScheduleServiceServer) {
	s.RegisterService(&SnapshotScheduleService_ServiceDesc, srv)
}

func _SnapshotScheduleService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnapshotScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).Get(ctx, req.(*GetSnapshotScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotSchedulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).List(ctx, req.(*ListSnapshotSchedulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).Create(ctx, req.(*CreateSnapshotScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnapshotScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).Update(ctx, req.(*UpdateSnapshotScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnapshotScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).Delete(ctx, req.(*DeleteSnapshotScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_UpdateDisks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnapshotScheduleDisksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).UpdateDisks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_UpdateDisks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).UpdateDisks(ctx, req.(*UpdateSnapshotScheduleDisksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableSnapshotScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).Disable(ctx, req.(*DisableSnapshotScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableSnapshotScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).Enable(ctx, req.(*EnableSnapshotScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotScheduleOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).ListOperations(ctx, req.(*ListSnapshotScheduleOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_ListSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotScheduleSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).ListSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_ListSnapshots_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).ListSnapshots(ctx, req.(*ListSnapshotScheduleSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_ListDisks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotScheduleDisksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).ListDisks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_ListDisks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).ListDisks(ctx, req.(*ListSnapshotScheduleDisksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotScheduleService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotScheduleServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotScheduleService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotScheduleServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SnapshotScheduleService_ServiceDesc is the grpc.ServiceDesc for SnapshotScheduleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SnapshotScheduleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.compute.v1.SnapshotScheduleService",
	HandlerType: (*SnapshotScheduleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SnapshotScheduleService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SnapshotScheduleService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SnapshotScheduleService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SnapshotScheduleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SnapshotScheduleService_Delete_Handler,
		},
		{
			MethodName: "UpdateDisks",
			Handler:    _SnapshotScheduleService_UpdateDisks_Handler,
		},
		{
			MethodName: "Disable",
			Handler:    _SnapshotScheduleService_Disable_Handler,
		},
		{
			MethodName: "Enable",
			Handler:    _SnapshotScheduleService_Enable_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _SnapshotScheduleService_ListOperations_Handler,
		},
		{
			MethodName: "ListSnapshots",
			Handler:    _SnapshotScheduleService_ListSnapshots_Handler,
		},
		{
			MethodName: "ListDisks",
			Handler:    _SnapshotScheduleService_ListDisks_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _SnapshotScheduleService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _SnapshotScheduleService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _SnapshotScheduleService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/compute/v1/snapshot_schedule_service.proto",
}
