// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/containerregistry/v1/lifecycle_policy_service.proto

package containerregistry

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

// LifecyclePolicyServiceClient is the client API for LifecyclePolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LifecyclePolicyServiceClient interface {
	// Returns the specified lifecycle policy.
	//
	// To get the list of all available lifecycle policies, make a [List] request.
	Get(ctx context.Context, in *GetLifecyclePolicyRequest, opts ...grpc.CallOption) (*LifecyclePolicy, error)
	// Retrieves the list of lifecycle policies in the specified repository.
	List(ctx context.Context, in *ListLifecyclePoliciesRequest, opts ...grpc.CallOption) (*ListLifecyclePoliciesResponse, error)
	// Creates a lifecycle policy in the specified repository.
	Create(ctx context.Context, in *CreateLifecyclePolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified lifecycle policy.
	Update(ctx context.Context, in *UpdateLifecyclePolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified lifecycle policy.
	Delete(ctx context.Context, in *DeleteLifecyclePolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Creates a request of a dry run of the lifecycle policy.
	DryRun(ctx context.Context, in *DryRunLifecyclePolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the dry run result of the specified lifecycle policy.
	GetDryRunResult(ctx context.Context, in *GetDryRunLifecyclePolicyResultRequest, opts ...grpc.CallOption) (*DryRunLifecyclePolicyResult, error)
	// Retrieves the list of the dry run results.
	ListDryRunResults(ctx context.Context, in *ListDryRunLifecyclePolicyResultsRequest, opts ...grpc.CallOption) (*ListDryRunLifecyclePolicyResultsResponse, error)
	// Retrieves the list of the affected images.
	ListDryRunResultAffectedImages(ctx context.Context, in *ListDryRunLifecyclePolicyResultAffectedImagesRequest, opts ...grpc.CallOption) (*ListDryRunLifecyclePolicyResultAffectedImagesResponse, error)
}

type lifecyclePolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLifecyclePolicyServiceClient(cc grpc.ClientConnInterface) LifecyclePolicyServiceClient {
	return &lifecyclePolicyServiceClient{cc}
}

func (c *lifecyclePolicyServiceClient) Get(ctx context.Context, in *GetLifecyclePolicyRequest, opts ...grpc.CallOption) (*LifecyclePolicy, error) {
	out := new(LifecyclePolicy)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifecyclePolicyServiceClient) List(ctx context.Context, in *ListLifecyclePoliciesRequest, opts ...grpc.CallOption) (*ListLifecyclePoliciesResponse, error) {
	out := new(ListLifecyclePoliciesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifecyclePolicyServiceClient) Create(ctx context.Context, in *CreateLifecyclePolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifecyclePolicyServiceClient) Update(ctx context.Context, in *UpdateLifecyclePolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifecyclePolicyServiceClient) Delete(ctx context.Context, in *DeleteLifecyclePolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifecyclePolicyServiceClient) DryRun(ctx context.Context, in *DryRunLifecyclePolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/DryRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifecyclePolicyServiceClient) GetDryRunResult(ctx context.Context, in *GetDryRunLifecyclePolicyResultRequest, opts ...grpc.CallOption) (*DryRunLifecyclePolicyResult, error) {
	out := new(DryRunLifecyclePolicyResult)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/GetDryRunResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifecyclePolicyServiceClient) ListDryRunResults(ctx context.Context, in *ListDryRunLifecyclePolicyResultsRequest, opts ...grpc.CallOption) (*ListDryRunLifecyclePolicyResultsResponse, error) {
	out := new(ListDryRunLifecyclePolicyResultsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/ListDryRunResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifecyclePolicyServiceClient) ListDryRunResultAffectedImages(ctx context.Context, in *ListDryRunLifecyclePolicyResultAffectedImagesRequest, opts ...grpc.CallOption) (*ListDryRunLifecyclePolicyResultAffectedImagesResponse, error) {
	out := new(ListDryRunLifecyclePolicyResultAffectedImagesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/ListDryRunResultAffectedImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LifecyclePolicyServiceServer is the server API for LifecyclePolicyService service.
// All implementations should embed UnimplementedLifecyclePolicyServiceServer
// for forward compatibility
type LifecyclePolicyServiceServer interface {
	// Returns the specified lifecycle policy.
	//
	// To get the list of all available lifecycle policies, make a [List] request.
	Get(context.Context, *GetLifecyclePolicyRequest) (*LifecyclePolicy, error)
	// Retrieves the list of lifecycle policies in the specified repository.
	List(context.Context, *ListLifecyclePoliciesRequest) (*ListLifecyclePoliciesResponse, error)
	// Creates a lifecycle policy in the specified repository.
	Create(context.Context, *CreateLifecyclePolicyRequest) (*operation.Operation, error)
	// Updates the specified lifecycle policy.
	Update(context.Context, *UpdateLifecyclePolicyRequest) (*operation.Operation, error)
	// Deletes the specified lifecycle policy.
	Delete(context.Context, *DeleteLifecyclePolicyRequest) (*operation.Operation, error)
	// Creates a request of a dry run of the lifecycle policy.
	DryRun(context.Context, *DryRunLifecyclePolicyRequest) (*operation.Operation, error)
	// Returns the dry run result of the specified lifecycle policy.
	GetDryRunResult(context.Context, *GetDryRunLifecyclePolicyResultRequest) (*DryRunLifecyclePolicyResult, error)
	// Retrieves the list of the dry run results.
	ListDryRunResults(context.Context, *ListDryRunLifecyclePolicyResultsRequest) (*ListDryRunLifecyclePolicyResultsResponse, error)
	// Retrieves the list of the affected images.
	ListDryRunResultAffectedImages(context.Context, *ListDryRunLifecyclePolicyResultAffectedImagesRequest) (*ListDryRunLifecyclePolicyResultAffectedImagesResponse, error)
}

// UnimplementedLifecyclePolicyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLifecyclePolicyServiceServer struct {
}

func (UnimplementedLifecyclePolicyServiceServer) Get(context.Context, *GetLifecyclePolicyRequest) (*LifecyclePolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLifecyclePolicyServiceServer) List(context.Context, *ListLifecyclePoliciesRequest) (*ListLifecyclePoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLifecyclePolicyServiceServer) Create(context.Context, *CreateLifecyclePolicyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLifecyclePolicyServiceServer) Update(context.Context, *UpdateLifecyclePolicyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLifecyclePolicyServiceServer) Delete(context.Context, *DeleteLifecyclePolicyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLifecyclePolicyServiceServer) DryRun(context.Context, *DryRunLifecyclePolicyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DryRun not implemented")
}
func (UnimplementedLifecyclePolicyServiceServer) GetDryRunResult(context.Context, *GetDryRunLifecyclePolicyResultRequest) (*DryRunLifecyclePolicyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDryRunResult not implemented")
}
func (UnimplementedLifecyclePolicyServiceServer) ListDryRunResults(context.Context, *ListDryRunLifecyclePolicyResultsRequest) (*ListDryRunLifecyclePolicyResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDryRunResults not implemented")
}
func (UnimplementedLifecyclePolicyServiceServer) ListDryRunResultAffectedImages(context.Context, *ListDryRunLifecyclePolicyResultAffectedImagesRequest) (*ListDryRunLifecyclePolicyResultAffectedImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDryRunResultAffectedImages not implemented")
}

// UnsafeLifecyclePolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LifecyclePolicyServiceServer will
// result in compilation errors.
type UnsafeLifecyclePolicyServiceServer interface {
	mustEmbedUnimplementedLifecyclePolicyServiceServer()
}

func RegisterLifecyclePolicyServiceServer(s grpc.ServiceRegistrar, srv LifecyclePolicyServiceServer) {
	s.RegisterService(&LifecyclePolicyService_ServiceDesc, srv)
}

func _LifecyclePolicyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLifecyclePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecyclePolicyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecyclePolicyServiceServer).Get(ctx, req.(*GetLifecyclePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifecyclePolicyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLifecyclePoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecyclePolicyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecyclePolicyServiceServer).List(ctx, req.(*ListLifecyclePoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifecyclePolicyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLifecyclePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecyclePolicyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecyclePolicyServiceServer).Create(ctx, req.(*CreateLifecyclePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifecyclePolicyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLifecyclePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecyclePolicyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecyclePolicyServiceServer).Update(ctx, req.(*UpdateLifecyclePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifecyclePolicyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLifecyclePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecyclePolicyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecyclePolicyServiceServer).Delete(ctx, req.(*DeleteLifecyclePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifecyclePolicyService_DryRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DryRunLifecyclePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecyclePolicyServiceServer).DryRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/DryRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecyclePolicyServiceServer).DryRun(ctx, req.(*DryRunLifecyclePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifecyclePolicyService_GetDryRunResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDryRunLifecyclePolicyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecyclePolicyServiceServer).GetDryRunResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/GetDryRunResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecyclePolicyServiceServer).GetDryRunResult(ctx, req.(*GetDryRunLifecyclePolicyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifecyclePolicyService_ListDryRunResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDryRunLifecyclePolicyResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecyclePolicyServiceServer).ListDryRunResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/ListDryRunResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecyclePolicyServiceServer).ListDryRunResults(ctx, req.(*ListDryRunLifecyclePolicyResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifecyclePolicyService_ListDryRunResultAffectedImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDryRunLifecyclePolicyResultAffectedImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecyclePolicyServiceServer).ListDryRunResultAffectedImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.LifecyclePolicyService/ListDryRunResultAffectedImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecyclePolicyServiceServer).ListDryRunResultAffectedImages(ctx, req.(*ListDryRunLifecyclePolicyResultAffectedImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LifecyclePolicyService_ServiceDesc is the grpc.ServiceDesc for LifecyclePolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LifecyclePolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.containerregistry.v1.LifecyclePolicyService",
	HandlerType: (*LifecyclePolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _LifecyclePolicyService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _LifecyclePolicyService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LifecyclePolicyService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LifecyclePolicyService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LifecyclePolicyService_Delete_Handler,
		},
		{
			MethodName: "DryRun",
			Handler:    _LifecyclePolicyService_DryRun_Handler,
		},
		{
			MethodName: "GetDryRunResult",
			Handler:    _LifecyclePolicyService_GetDryRunResult_Handler,
		},
		{
			MethodName: "ListDryRunResults",
			Handler:    _LifecyclePolicyService_ListDryRunResults_Handler,
		},
		{
			MethodName: "ListDryRunResultAffectedImages",
			Handler:    _LifecyclePolicyService_ListDryRunResultAffectedImages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/containerregistry/v1/lifecycle_policy_service.proto",
}
