// Code generated by sdkgen. DO NOT EDIT.

//nolint
package resourcemanager

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	resourcemanager "github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

//revive:disable

// CloudServiceClient is a resourcemanager.CloudServiceClient with
// lazy GRPC connection initialization.
type CloudServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements resourcemanager.CloudServiceClient
func (c *CloudServiceClient) Create(ctx context.Context, in *resourcemanager.CreateCloudRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewCloudServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements resourcemanager.CloudServiceClient
func (c *CloudServiceClient) Delete(ctx context.Context, in *resourcemanager.DeleteCloudRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewCloudServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements resourcemanager.CloudServiceClient
func (c *CloudServiceClient) Get(ctx context.Context, in *resourcemanager.GetCloudRequest, opts ...grpc.CallOption) (*resourcemanager.Cloud, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewCloudServiceClient(conn).Get(ctx, in, opts...)
}

// List implements resourcemanager.CloudServiceClient
func (c *CloudServiceClient) List(ctx context.Context, in *resourcemanager.ListCloudsRequest, opts ...grpc.CallOption) (*resourcemanager.ListCloudsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewCloudServiceClient(conn).List(ctx, in, opts...)
}

type CloudIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *CloudServiceClient
	request *resourcemanager.ListCloudsRequest

	items []*resourcemanager.Cloud
}

func (c *CloudServiceClient) CloudIterator(ctx context.Context, req *resourcemanager.ListCloudsRequest, opts ...grpc.CallOption) *CloudIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &CloudIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *CloudIterator) Next() bool {
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

	it.items = response.Clouds
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *CloudIterator) Take(size int64) ([]*resourcemanager.Cloud, error) {
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

	var result []*resourcemanager.Cloud

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *CloudIterator) TakeAll() ([]*resourcemanager.Cloud, error) {
	return it.Take(0)
}

func (it *CloudIterator) Value() *resourcemanager.Cloud {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *CloudIterator) Error() error {
	return it.err
}

// ListAccessBindings implements resourcemanager.CloudServiceClient
func (c *CloudServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewCloudServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type CloudAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *CloudServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *CloudServiceClient) CloudAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *CloudAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &CloudAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *CloudAccessBindingsIterator) Next() bool {
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

func (it *CloudAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
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

func (it *CloudAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *CloudAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *CloudAccessBindingsIterator) Error() error {
	return it.err
}

// ListOperations implements resourcemanager.CloudServiceClient
func (c *CloudServiceClient) ListOperations(ctx context.Context, in *resourcemanager.ListCloudOperationsRequest, opts ...grpc.CallOption) (*resourcemanager.ListCloudOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewCloudServiceClient(conn).ListOperations(ctx, in, opts...)
}

type CloudOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *CloudServiceClient
	request *resourcemanager.ListCloudOperationsRequest

	items []*operation.Operation
}

func (c *CloudServiceClient) CloudOperationsIterator(ctx context.Context, req *resourcemanager.ListCloudOperationsRequest, opts ...grpc.CallOption) *CloudOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &CloudOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *CloudOperationsIterator) Next() bool {
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

func (it *CloudOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

func (it *CloudOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *CloudOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *CloudOperationsIterator) Error() error {
	return it.err
}

// SetAccessBindings implements resourcemanager.CloudServiceClient
func (c *CloudServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewCloudServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Update implements resourcemanager.CloudServiceClient
func (c *CloudServiceClient) Update(ctx context.Context, in *resourcemanager.UpdateCloudRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewCloudServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements resourcemanager.CloudServiceClient
func (c *CloudServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewCloudServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}
