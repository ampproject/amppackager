// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/datasphere/v2/docker_image_service.proto

package datasphere

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
	DockerImageService_Activate_FullMethodName = "/yandex.cloud.datasphere.v2.DockerImageService/Activate"
)

// DockerImageServiceClient is the client API for DockerImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DockerImageServiceClient interface {
	// Activates shared docker image in project
	Activate(ctx context.Context, in *ActivateDockerImageRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type dockerImageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDockerImageServiceClient(cc grpc.ClientConnInterface) DockerImageServiceClient {
	return &dockerImageServiceClient{cc}
}

func (c *dockerImageServiceClient) Activate(ctx context.Context, in *ActivateDockerImageRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DockerImageService_Activate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockerImageServiceServer is the server API for DockerImageService service.
// All implementations should embed UnimplementedDockerImageServiceServer
// for forward compatibility
type DockerImageServiceServer interface {
	// Activates shared docker image in project
	Activate(context.Context, *ActivateDockerImageRequest) (*operation.Operation, error)
}

// UnimplementedDockerImageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDockerImageServiceServer struct {
}

func (UnimplementedDockerImageServiceServer) Activate(context.Context, *ActivateDockerImageRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}

// UnsafeDockerImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DockerImageServiceServer will
// result in compilation errors.
type UnsafeDockerImageServiceServer interface {
	mustEmbedUnimplementedDockerImageServiceServer()
}

func RegisterDockerImageServiceServer(s grpc.ServiceRegistrar, srv DockerImageServiceServer) {
	s.RegisterService(&DockerImageService_ServiceDesc, srv)
}

func _DockerImageService_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateDockerImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerImageServiceServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DockerImageService_Activate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerImageServiceServer).Activate(ctx, req.(*ActivateDockerImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DockerImageService_ServiceDesc is the grpc.ServiceDesc for DockerImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DockerImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.datasphere.v2.DockerImageService",
	HandlerType: (*DockerImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activate",
			Handler:    _DockerImageService_Activate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/datasphere/v2/docker_image_service.proto",
}