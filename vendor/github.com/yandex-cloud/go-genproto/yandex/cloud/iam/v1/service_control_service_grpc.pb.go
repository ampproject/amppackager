// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/iam/v1/service_control_service.proto

package iam

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
	ServiceControlService_Get_FullMethodName          = "/yandex.cloud.iam.v1.ServiceControlService/Get"
	ServiceControlService_List_FullMethodName         = "/yandex.cloud.iam.v1.ServiceControlService/List"
	ServiceControlService_Enable_FullMethodName       = "/yandex.cloud.iam.v1.ServiceControlService/Enable"
	ServiceControlService_Disable_FullMethodName      = "/yandex.cloud.iam.v1.ServiceControlService/Disable"
	ServiceControlService_ResolveAgent_FullMethodName = "/yandex.cloud.iam.v1.ServiceControlService/ResolveAgent"
)

// ServiceControlServiceClient is the client API for ServiceControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceControlServiceClient interface {
	// Returns the Service information in the specified resource container.
	//
	// To get the list of available Services, make a [List] request.
	Get(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*Service, error)
	// Retrieves the list of Service in the specified resource container.
	List(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	// Enable a service in the specified resource container.
	Enable(ctx context.Context, in *EnableServiceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Disable a service in the specified resource container.
	Disable(ctx context.Context, in *DisableServiceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Resolve agent service account for the service in the specified resource container.
	ResolveAgent(ctx context.Context, in *ResolveServiceAgentRequest, opts ...grpc.CallOption) (*ServiceAgent, error)
}

type serviceControlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceControlServiceClient(cc grpc.ClientConnInterface) ServiceControlServiceClient {
	return &serviceControlServiceClient{cc}
}

func (c *serviceControlServiceClient) Get(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, ServiceControlService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceControlServiceClient) List(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, ServiceControlService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceControlServiceClient) Enable(ctx context.Context, in *EnableServiceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ServiceControlService_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceControlServiceClient) Disable(ctx context.Context, in *DisableServiceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ServiceControlService_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceControlServiceClient) ResolveAgent(ctx context.Context, in *ResolveServiceAgentRequest, opts ...grpc.CallOption) (*ServiceAgent, error) {
	out := new(ServiceAgent)
	err := c.cc.Invoke(ctx, ServiceControlService_ResolveAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceControlServiceServer is the server API for ServiceControlService service.
// All implementations should embed UnimplementedServiceControlServiceServer
// for forward compatibility
type ServiceControlServiceServer interface {
	// Returns the Service information in the specified resource container.
	//
	// To get the list of available Services, make a [List] request.
	Get(context.Context, *GetServiceRequest) (*Service, error)
	// Retrieves the list of Service in the specified resource container.
	List(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	// Enable a service in the specified resource container.
	Enable(context.Context, *EnableServiceRequest) (*operation.Operation, error)
	// Disable a service in the specified resource container.
	Disable(context.Context, *DisableServiceRequest) (*operation.Operation, error)
	// Resolve agent service account for the service in the specified resource container.
	ResolveAgent(context.Context, *ResolveServiceAgentRequest) (*ServiceAgent, error)
}

// UnimplementedServiceControlServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServiceControlServiceServer struct {
}

func (UnimplementedServiceControlServiceServer) Get(context.Context, *GetServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedServiceControlServiceServer) List(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedServiceControlServiceServer) Enable(context.Context, *EnableServiceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedServiceControlServiceServer) Disable(context.Context, *DisableServiceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedServiceControlServiceServer) ResolveAgent(context.Context, *ResolveServiceAgentRequest) (*ServiceAgent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveAgent not implemented")
}

// UnsafeServiceControlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceControlServiceServer will
// result in compilation errors.
type UnsafeServiceControlServiceServer interface {
	mustEmbedUnimplementedServiceControlServiceServer()
}

func RegisterServiceControlServiceServer(s grpc.ServiceRegistrar, srv ServiceControlServiceServer) {
	s.RegisterService(&ServiceControlService_ServiceDesc, srv)
}

func _ServiceControlService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceControlServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceControlService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceControlServiceServer).Get(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceControlService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceControlServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceControlService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceControlServiceServer).List(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceControlService_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceControlServiceServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceControlService_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceControlServiceServer).Enable(ctx, req.(*EnableServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceControlService_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceControlServiceServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceControlService_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceControlServiceServer).Disable(ctx, req.(*DisableServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceControlService_ResolveAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveServiceAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceControlServiceServer).ResolveAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceControlService_ResolveAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceControlServiceServer).ResolveAgent(ctx, req.(*ResolveServiceAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceControlService_ServiceDesc is the grpc.ServiceDesc for ServiceControlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceControlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.iam.v1.ServiceControlService",
	HandlerType: (*ServiceControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ServiceControlService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ServiceControlService_List_Handler,
		},
		{
			MethodName: "Enable",
			Handler:    _ServiceControlService_Enable_Handler,
		},
		{
			MethodName: "Disable",
			Handler:    _ServiceControlService_Disable_Handler,
		},
		{
			MethodName: "ResolveAgent",
			Handler:    _ServiceControlService_ResolveAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/iam/v1/service_control_service.proto",
}