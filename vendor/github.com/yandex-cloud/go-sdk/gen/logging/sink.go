// Code generated by sdkgen. DO NOT EDIT.

// nolint
package logging

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	logging "github.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// SinkServiceClient is a logging.SinkServiceClient with
// lazy GRPC connection initialization.
type SinkServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements logging.SinkServiceClient
func (c *SinkServiceClient) Create(ctx context.Context, in *logging.CreateSinkRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return logging.NewSinkServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements logging.SinkServiceClient
func (c *SinkServiceClient) Delete(ctx context.Context, in *logging.DeleteSinkRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return logging.NewSinkServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements logging.SinkServiceClient
func (c *SinkServiceClient) Get(ctx context.Context, in *logging.GetSinkRequest, opts ...grpc.CallOption) (*logging.Sink, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return logging.NewSinkServiceClient(conn).Get(ctx, in, opts...)
}

// List implements logging.SinkServiceClient
func (c *SinkServiceClient) List(ctx context.Context, in *logging.ListSinksRequest, opts ...grpc.CallOption) (*logging.ListSinksResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return logging.NewSinkServiceClient(conn).List(ctx, in, opts...)
}

type SinkIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SinkServiceClient
	request *logging.ListSinksRequest

	items []*logging.Sink
}

func (c *SinkServiceClient) SinkIterator(ctx context.Context, req *logging.ListSinksRequest, opts ...grpc.CallOption) *SinkIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SinkIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SinkIterator) Next() bool {
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

	it.items = response.Sinks
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *SinkIterator) Take(size int64) ([]*logging.Sink, error) {
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

	var result []*logging.Sink

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *SinkIterator) TakeAll() ([]*logging.Sink, error) {
	return it.Take(0)
}

func (it *SinkIterator) Value() *logging.Sink {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SinkIterator) Error() error {
	return it.err
}

// ListAccessBindings implements logging.SinkServiceClient
func (c *SinkServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return logging.NewSinkServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type SinkAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SinkServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *SinkServiceClient) SinkAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *SinkAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SinkAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SinkAccessBindingsIterator) Next() bool {
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

func (it *SinkAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
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

func (it *SinkAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *SinkAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SinkAccessBindingsIterator) Error() error {
	return it.err
}

// ListOperations implements logging.SinkServiceClient
func (c *SinkServiceClient) ListOperations(ctx context.Context, in *logging.ListSinkOperationsRequest, opts ...grpc.CallOption) (*logging.ListSinkOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return logging.NewSinkServiceClient(conn).ListOperations(ctx, in, opts...)
}

type SinkOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SinkServiceClient
	request *logging.ListSinkOperationsRequest

	items []*operation.Operation
}

func (c *SinkServiceClient) SinkOperationsIterator(ctx context.Context, req *logging.ListSinkOperationsRequest, opts ...grpc.CallOption) *SinkOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SinkOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SinkOperationsIterator) Next() bool {
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

func (it *SinkOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

func (it *SinkOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *SinkOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SinkOperationsIterator) Error() error {
	return it.err
}

// SetAccessBindings implements logging.SinkServiceClient
func (c *SinkServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return logging.NewSinkServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Update implements logging.SinkServiceClient
func (c *SinkServiceClient) Update(ctx context.Context, in *logging.UpdateSinkRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return logging.NewSinkServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements logging.SinkServiceClient
func (c *SinkServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return logging.NewSinkServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}