// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/loadtesting/agent/v1/agent_registration_service.proto

package agent

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
	AgentRegistrationService_Register_FullMethodName              = "/yandex.cloud.loadtesting.agent.v1.AgentRegistrationService/Register"
	AgentRegistrationService_ExternalAgentRegister_FullMethodName = "/yandex.cloud.loadtesting.agent.v1.AgentRegistrationService/ExternalAgentRegister"
)

// AgentRegistrationServiceClient is the client API for AgentRegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentRegistrationServiceClient interface {
	// Registers specified agent.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Registers external agent.
	ExternalAgentRegister(ctx context.Context, in *ExternalAgentRegisterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type agentRegistrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentRegistrationServiceClient(cc grpc.ClientConnInterface) AgentRegistrationServiceClient {
	return &agentRegistrationServiceClient{cc}
}

func (c *agentRegistrationServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, AgentRegistrationService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRegistrationServiceClient) ExternalAgentRegister(ctx context.Context, in *ExternalAgentRegisterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AgentRegistrationService_ExternalAgentRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentRegistrationServiceServer is the server API for AgentRegistrationService service.
// All implementations should embed UnimplementedAgentRegistrationServiceServer
// for forward compatibility
type AgentRegistrationServiceServer interface {
	// Registers specified agent.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Registers external agent.
	ExternalAgentRegister(context.Context, *ExternalAgentRegisterRequest) (*operation.Operation, error)
}

// UnimplementedAgentRegistrationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAgentRegistrationServiceServer struct {
}

func (UnimplementedAgentRegistrationServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAgentRegistrationServiceServer) ExternalAgentRegister(context.Context, *ExternalAgentRegisterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExternalAgentRegister not implemented")
}

// UnsafeAgentRegistrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentRegistrationServiceServer will
// result in compilation errors.
type UnsafeAgentRegistrationServiceServer interface {
	mustEmbedUnimplementedAgentRegistrationServiceServer()
}

func RegisterAgentRegistrationServiceServer(s grpc.ServiceRegistrar, srv AgentRegistrationServiceServer) {
	s.RegisterService(&AgentRegistrationService_ServiceDesc, srv)
}

func _AgentRegistrationService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRegistrationServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentRegistrationService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRegistrationServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRegistrationService_ExternalAgentRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExternalAgentRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRegistrationServiceServer).ExternalAgentRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentRegistrationService_ExternalAgentRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRegistrationServiceServer).ExternalAgentRegister(ctx, req.(*ExternalAgentRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentRegistrationService_ServiceDesc is the grpc.ServiceDesc for AgentRegistrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentRegistrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.loadtesting.agent.v1.AgentRegistrationService",
	HandlerType: (*AgentRegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AgentRegistrationService_Register_Handler,
		},
		{
			MethodName: "ExternalAgentRegister",
			Handler:    _AgentRegistrationService_ExternalAgentRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/loadtesting/agent/v1/agent_registration_service.proto",
}
