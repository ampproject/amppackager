// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/iam/v1/iam_token_service.proto

package iam

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
	IamTokenService_Create_FullMethodName                  = "/yandex.cloud.iam.v1.IamTokenService/Create"
	IamTokenService_CreateForServiceAccount_FullMethodName = "/yandex.cloud.iam.v1.IamTokenService/CreateForServiceAccount"
	IamTokenService_Revoke_FullMethodName                  = "/yandex.cloud.iam.v1.IamTokenService/Revoke"
)

// IamTokenServiceClient is the client API for IamTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IamTokenServiceClient interface {
	// Create an IAM token for the specified identity.
	Create(ctx context.Context, in *CreateIamTokenRequest, opts ...grpc.CallOption) (*CreateIamTokenResponse, error)
	// Create an IAM token for service account.
	CreateForServiceAccount(ctx context.Context, in *CreateIamTokenForServiceAccountRequest, opts ...grpc.CallOption) (*CreateIamTokenResponse, error)
	// Revoke the IAM token.
	Revoke(ctx context.Context, in *RevokeIamTokenRequest, opts ...grpc.CallOption) (*RevokeIamTokenResponse, error)
}

type iamTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIamTokenServiceClient(cc grpc.ClientConnInterface) IamTokenServiceClient {
	return &iamTokenServiceClient{cc}
}

func (c *iamTokenServiceClient) Create(ctx context.Context, in *CreateIamTokenRequest, opts ...grpc.CallOption) (*CreateIamTokenResponse, error) {
	out := new(CreateIamTokenResponse)
	err := c.cc.Invoke(ctx, IamTokenService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamTokenServiceClient) CreateForServiceAccount(ctx context.Context, in *CreateIamTokenForServiceAccountRequest, opts ...grpc.CallOption) (*CreateIamTokenResponse, error) {
	out := new(CreateIamTokenResponse)
	err := c.cc.Invoke(ctx, IamTokenService_CreateForServiceAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamTokenServiceClient) Revoke(ctx context.Context, in *RevokeIamTokenRequest, opts ...grpc.CallOption) (*RevokeIamTokenResponse, error) {
	out := new(RevokeIamTokenResponse)
	err := c.cc.Invoke(ctx, IamTokenService_Revoke_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamTokenServiceServer is the server API for IamTokenService service.
// All implementations should embed UnimplementedIamTokenServiceServer
// for forward compatibility
type IamTokenServiceServer interface {
	// Create an IAM token for the specified identity.
	Create(context.Context, *CreateIamTokenRequest) (*CreateIamTokenResponse, error)
	// Create an IAM token for service account.
	CreateForServiceAccount(context.Context, *CreateIamTokenForServiceAccountRequest) (*CreateIamTokenResponse, error)
	// Revoke the IAM token.
	Revoke(context.Context, *RevokeIamTokenRequest) (*RevokeIamTokenResponse, error)
}

// UnimplementedIamTokenServiceServer should be embedded to have forward compatible implementations.
type UnimplementedIamTokenServiceServer struct {
}

func (UnimplementedIamTokenServiceServer) Create(context.Context, *CreateIamTokenRequest) (*CreateIamTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIamTokenServiceServer) CreateForServiceAccount(context.Context, *CreateIamTokenForServiceAccountRequest) (*CreateIamTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateForServiceAccount not implemented")
}
func (UnimplementedIamTokenServiceServer) Revoke(context.Context, *RevokeIamTokenRequest) (*RevokeIamTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revoke not implemented")
}

// UnsafeIamTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IamTokenServiceServer will
// result in compilation errors.
type UnsafeIamTokenServiceServer interface {
	mustEmbedUnimplementedIamTokenServiceServer()
}

func RegisterIamTokenServiceServer(s grpc.ServiceRegistrar, srv IamTokenServiceServer) {
	s.RegisterService(&IamTokenService_ServiceDesc, srv)
}

func _IamTokenService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIamTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamTokenServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamTokenService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamTokenServiceServer).Create(ctx, req.(*CreateIamTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamTokenService_CreateForServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIamTokenForServiceAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamTokenServiceServer).CreateForServiceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamTokenService_CreateForServiceAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamTokenServiceServer).CreateForServiceAccount(ctx, req.(*CreateIamTokenForServiceAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamTokenService_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeIamTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamTokenServiceServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamTokenService_Revoke_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamTokenServiceServer).Revoke(ctx, req.(*RevokeIamTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IamTokenService_ServiceDesc is the grpc.ServiceDesc for IamTokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IamTokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.iam.v1.IamTokenService",
	HandlerType: (*IamTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _IamTokenService_Create_Handler,
		},
		{
			MethodName: "CreateForServiceAccount",
			Handler:    _IamTokenService_CreateForServiceAccount_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _IamTokenService_Revoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/iam/v1/iam_token_service.proto",
}
