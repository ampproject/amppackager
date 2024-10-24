// Code generated by sdkgen. DO NOT EDIT.

// nolint
package compute

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// GpuClusterServiceClient is a compute.GpuClusterServiceClient with
// lazy GRPC connection initialization.
type GpuClusterServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements compute.GpuClusterServiceClient
func (c *GpuClusterServiceClient) Create(ctx context.Context, in *compute.CreateGpuClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewGpuClusterServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements compute.GpuClusterServiceClient
func (c *GpuClusterServiceClient) Delete(ctx context.Context, in *compute.DeleteGpuClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewGpuClusterServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements compute.GpuClusterServiceClient
func (c *GpuClusterServiceClient) Get(ctx context.Context, in *compute.GetGpuClusterRequest, opts ...grpc.CallOption) (*compute.GpuCluster, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewGpuClusterServiceClient(conn).Get(ctx, in, opts...)
}

// List implements compute.GpuClusterServiceClient
func (c *GpuClusterServiceClient) List(ctx context.Context, in *compute.ListGpuClustersRequest, opts ...grpc.CallOption) (*compute.ListGpuClustersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewGpuClusterServiceClient(conn).List(ctx, in, opts...)
}

type GpuClusterIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *GpuClusterServiceClient
	request *compute.ListGpuClustersRequest

	items []*compute.GpuCluster
}

func (c *GpuClusterServiceClient) GpuClusterIterator(ctx context.Context, req *compute.ListGpuClustersRequest, opts ...grpc.CallOption) *GpuClusterIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &GpuClusterIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *GpuClusterIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.GpuClusters
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *GpuClusterIterator) Take(size int64) ([]*compute.GpuCluster, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*compute.GpuCluster

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *GpuClusterIterator) TakeAll() ([]*compute.GpuCluster, error) {
	return it.Take(0)
}

func (it *GpuClusterIterator) Value() *compute.GpuCluster {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *GpuClusterIterator) Error() error {
	return it.err
}

// ListAccessBindings implements compute.GpuClusterServiceClient
func (c *GpuClusterServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewGpuClusterServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type GpuClusterAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *GpuClusterServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *GpuClusterServiceClient) GpuClusterAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *GpuClusterAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &GpuClusterAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *GpuClusterAccessBindingsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListAccessBindings(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.AccessBindings
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *GpuClusterAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*access.AccessBinding

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *GpuClusterAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *GpuClusterAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *GpuClusterAccessBindingsIterator) Error() error {
	return it.err
}

// ListInstances implements compute.GpuClusterServiceClient
func (c *GpuClusterServiceClient) ListInstances(ctx context.Context, in *compute.ListGpuClusterInstancesRequest, opts ...grpc.CallOption) (*compute.ListGpuClusterInstancesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewGpuClusterServiceClient(conn).ListInstances(ctx, in, opts...)
}

type GpuClusterInstancesIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *GpuClusterServiceClient
	request *compute.ListGpuClusterInstancesRequest

	items []*compute.Instance
}

func (c *GpuClusterServiceClient) GpuClusterInstancesIterator(ctx context.Context, req *compute.ListGpuClusterInstancesRequest, opts ...grpc.CallOption) *GpuClusterInstancesIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &GpuClusterInstancesIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *GpuClusterInstancesIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListInstances(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Instances
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *GpuClusterInstancesIterator) Take(size int64) ([]*compute.Instance, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*compute.Instance

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *GpuClusterInstancesIterator) TakeAll() ([]*compute.Instance, error) {
	return it.Take(0)
}

func (it *GpuClusterInstancesIterator) Value() *compute.Instance {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *GpuClusterInstancesIterator) Error() error {
	return it.err
}

// ListOperations implements compute.GpuClusterServiceClient
func (c *GpuClusterServiceClient) ListOperations(ctx context.Context, in *compute.ListGpuClusterOperationsRequest, opts ...grpc.CallOption) (*compute.ListGpuClusterOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewGpuClusterServiceClient(conn).ListOperations(ctx, in, opts...)
}

type GpuClusterOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *GpuClusterServiceClient
	request *compute.ListGpuClusterOperationsRequest

	items []*operation.Operation
}

func (c *GpuClusterServiceClient) GpuClusterOperationsIterator(ctx context.Context, req *compute.ListGpuClusterOperationsRequest, opts ...grpc.CallOption) *GpuClusterOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &GpuClusterOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *GpuClusterOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *GpuClusterOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *GpuClusterOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *GpuClusterOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *GpuClusterOperationsIterator) Error() error {
	return it.err
}

// SetAccessBindings implements compute.GpuClusterServiceClient
func (c *GpuClusterServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewGpuClusterServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Update implements compute.GpuClusterServiceClient
func (c *GpuClusterServiceClient) Update(ctx context.Context, in *compute.UpdateGpuClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewGpuClusterServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements compute.GpuClusterServiceClient
func (c *GpuClusterServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewGpuClusterServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}
