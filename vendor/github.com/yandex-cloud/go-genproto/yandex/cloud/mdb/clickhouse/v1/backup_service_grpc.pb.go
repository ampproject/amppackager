// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/mdb/clickhouse/v1/backup_service.proto

package clickhouse

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

// BackupServiceClient is the client API for BackupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackupServiceClient interface {
	// Returns the specified ClickHouse Backup resource.
	//
	// To get the list of available ClickHouse Backup resources, make a [List] request.
	Get(ctx context.Context, in *GetBackupRequest, opts ...grpc.CallOption) (*Backup, error)
	// Retrieves the list of Backup resources available for the specified folder.
	List(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error)
}

type backupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupServiceClient(cc grpc.ClientConnInterface) BackupServiceClient {
	return &backupServiceClient{cc}
}

func (c *backupServiceClient) Get(ctx context.Context, in *GetBackupRequest, opts ...grpc.CallOption) (*Backup, error) {
	out := new(Backup)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.clickhouse.v1.BackupService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) List(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error) {
	out := new(ListBackupsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.clickhouse.v1.BackupService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServiceServer is the server API for BackupService service.
// All implementations should embed UnimplementedBackupServiceServer
// for forward compatibility
type BackupServiceServer interface {
	// Returns the specified ClickHouse Backup resource.
	//
	// To get the list of available ClickHouse Backup resources, make a [List] request.
	Get(context.Context, *GetBackupRequest) (*Backup, error)
	// Retrieves the list of Backup resources available for the specified folder.
	List(context.Context, *ListBackupsRequest) (*ListBackupsResponse, error)
}

// UnimplementedBackupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBackupServiceServer struct {
}

func (UnimplementedBackupServiceServer) Get(context.Context, *GetBackupRequest) (*Backup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBackupServiceServer) List(context.Context, *ListBackupsRequest) (*ListBackupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeBackupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackupServiceServer will
// result in compilation errors.
type UnsafeBackupServiceServer interface {
	mustEmbedUnimplementedBackupServiceServer()
}

func RegisterBackupServiceServer(s grpc.ServiceRegistrar, srv BackupServiceServer) {
	s.RegisterService(&BackupService_ServiceDesc, srv)
}

func _BackupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.clickhouse.v1.BackupService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).Get(ctx, req.(*GetBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBackupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.clickhouse.v1.BackupService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).List(ctx, req.(*ListBackupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BackupService_ServiceDesc is the grpc.ServiceDesc for BackupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BackupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.clickhouse.v1.BackupService",
	HandlerType: (*BackupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BackupService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _BackupService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/mdb/clickhouse/v1/backup_service.proto",
}
