// Code generated by sdkgen. DO NOT EDIT.

// nolint
package datasphere

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	datasphere "github.com/yandex-cloud/go-genproto/yandex/cloud/datasphere/v2"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// CommunityServiceClient is a datasphere.CommunityServiceClient with
// lazy GRPC connection initialization.
type CommunityServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// AddResource implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) AddResource(ctx context.Context, in *datasphere.AddCommunityResourceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).AddResource(ctx, in, opts...)
}

// Create implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) Create(ctx context.Context, in *datasphere.CreateCommunityRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) Delete(ctx context.Context, in *datasphere.DeleteCommunityRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) Get(ctx context.Context, in *datasphere.GetCommunityRequest, opts ...grpc.CallOption) (*datasphere.Community, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).Get(ctx, in, opts...)
}

// GetRestrictions implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) GetRestrictions(ctx context.Context, in *datasphere.GetCommunityRestrictionsRequest, opts ...grpc.CallOption) (*datasphere.RestrictionsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).GetRestrictions(ctx, in, opts...)
}

// GetRestrictionsMeta implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) GetRestrictionsMeta(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*datasphere.GetRestrictionsMetaResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).GetRestrictionsMeta(ctx, in, opts...)
}

// List implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) List(ctx context.Context, in *datasphere.ListCommunitiesRequest, opts ...grpc.CallOption) (*datasphere.ListCommunitiesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).List(ctx, in, opts...)
}

type CommunityIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *CommunityServiceClient
	request *datasphere.ListCommunitiesRequest

	items []*datasphere.Community
}

func (c *CommunityServiceClient) CommunityIterator(ctx context.Context, req *datasphere.ListCommunitiesRequest, opts ...grpc.CallOption) *CommunityIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &CommunityIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *CommunityIterator) Next() bool {
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

	it.items = response.Communities
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *CommunityIterator) Take(size int64) ([]*datasphere.Community, error) {
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

	var result []*datasphere.Community

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *CommunityIterator) TakeAll() ([]*datasphere.Community, error) {
	return it.Take(0)
}

func (it *CommunityIterator) Value() *datasphere.Community {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *CommunityIterator) Error() error {
	return it.err
}

// ListAccessBindings implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type CommunityAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *CommunityServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *CommunityServiceClient) CommunityAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *CommunityAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &CommunityAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *CommunityAccessBindingsIterator) Next() bool {
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

func (it *CommunityAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
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

func (it *CommunityAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *CommunityAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *CommunityAccessBindingsIterator) Error() error {
	return it.err
}

// RemoveResource implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) RemoveResource(ctx context.Context, in *datasphere.RemoveCommunityResourceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).RemoveResource(ctx, in, opts...)
}

// SetAccessBindings implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// SetRestrictions implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) SetRestrictions(ctx context.Context, in *datasphere.SetCommunityRestrictionsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).SetRestrictions(ctx, in, opts...)
}

// Update implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) Update(ctx context.Context, in *datasphere.UpdateCommunityRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements datasphere.CommunityServiceClient
func (c *CommunityServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewCommunityServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}