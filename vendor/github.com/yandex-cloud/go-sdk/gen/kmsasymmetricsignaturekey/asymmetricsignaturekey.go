// Code generated by sdkgen. DO NOT EDIT.

// nolint
package asymmetricsignature

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	asymmetricsignature "github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1/asymmetricsignature"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// AsymmetricSignatureKeyServiceClient is a asymmetricsignature.AsymmetricSignatureKeyServiceClient with
// lazy GRPC connection initialization.
type AsymmetricSignatureKeyServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements asymmetricsignature.AsymmetricSignatureKeyServiceClient
func (c *AsymmetricSignatureKeyServiceClient) Create(ctx context.Context, in *asymmetricsignature.CreateAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return asymmetricsignature.NewAsymmetricSignatureKeyServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements asymmetricsignature.AsymmetricSignatureKeyServiceClient
func (c *AsymmetricSignatureKeyServiceClient) Delete(ctx context.Context, in *asymmetricsignature.DeleteAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return asymmetricsignature.NewAsymmetricSignatureKeyServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements asymmetricsignature.AsymmetricSignatureKeyServiceClient
func (c *AsymmetricSignatureKeyServiceClient) Get(ctx context.Context, in *asymmetricsignature.GetAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*asymmetricsignature.AsymmetricSignatureKey, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return asymmetricsignature.NewAsymmetricSignatureKeyServiceClient(conn).Get(ctx, in, opts...)
}

// List implements asymmetricsignature.AsymmetricSignatureKeyServiceClient
func (c *AsymmetricSignatureKeyServiceClient) List(ctx context.Context, in *asymmetricsignature.ListAsymmetricSignatureKeysRequest, opts ...grpc.CallOption) (*asymmetricsignature.ListAsymmetricSignatureKeysResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return asymmetricsignature.NewAsymmetricSignatureKeyServiceClient(conn).List(ctx, in, opts...)
}

type AsymmetricSignatureKeyIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *AsymmetricSignatureKeyServiceClient
	request *asymmetricsignature.ListAsymmetricSignatureKeysRequest

	items []*asymmetricsignature.AsymmetricSignatureKey
}

func (c *AsymmetricSignatureKeyServiceClient) AsymmetricSignatureKeyIterator(ctx context.Context, req *asymmetricsignature.ListAsymmetricSignatureKeysRequest, opts ...grpc.CallOption) *AsymmetricSignatureKeyIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &AsymmetricSignatureKeyIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *AsymmetricSignatureKeyIterator) Next() bool {
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

	it.items = response.Keys
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *AsymmetricSignatureKeyIterator) Take(size int64) ([]*asymmetricsignature.AsymmetricSignatureKey, error) {
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

	var result []*asymmetricsignature.AsymmetricSignatureKey

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *AsymmetricSignatureKeyIterator) TakeAll() ([]*asymmetricsignature.AsymmetricSignatureKey, error) {
	return it.Take(0)
}

func (it *AsymmetricSignatureKeyIterator) Value() *asymmetricsignature.AsymmetricSignatureKey {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *AsymmetricSignatureKeyIterator) Error() error {
	return it.err
}

// ListAccessBindings implements asymmetricsignature.AsymmetricSignatureKeyServiceClient
func (c *AsymmetricSignatureKeyServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return asymmetricsignature.NewAsymmetricSignatureKeyServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type AsymmetricSignatureKeyAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *AsymmetricSignatureKeyServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *AsymmetricSignatureKeyServiceClient) AsymmetricSignatureKeyAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *AsymmetricSignatureKeyAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &AsymmetricSignatureKeyAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *AsymmetricSignatureKeyAccessBindingsIterator) Next() bool {
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

func (it *AsymmetricSignatureKeyAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
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

func (it *AsymmetricSignatureKeyAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *AsymmetricSignatureKeyAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *AsymmetricSignatureKeyAccessBindingsIterator) Error() error {
	return it.err
}

// ListOperations implements asymmetricsignature.AsymmetricSignatureKeyServiceClient
func (c *AsymmetricSignatureKeyServiceClient) ListOperations(ctx context.Context, in *asymmetricsignature.ListAsymmetricSignatureKeyOperationsRequest, opts ...grpc.CallOption) (*asymmetricsignature.ListAsymmetricSignatureKeyOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return asymmetricsignature.NewAsymmetricSignatureKeyServiceClient(conn).ListOperations(ctx, in, opts...)
}

type AsymmetricSignatureKeyOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *AsymmetricSignatureKeyServiceClient
	request *asymmetricsignature.ListAsymmetricSignatureKeyOperationsRequest

	items []*operation.Operation
}

func (c *AsymmetricSignatureKeyServiceClient) AsymmetricSignatureKeyOperationsIterator(ctx context.Context, req *asymmetricsignature.ListAsymmetricSignatureKeyOperationsRequest, opts ...grpc.CallOption) *AsymmetricSignatureKeyOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &AsymmetricSignatureKeyOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *AsymmetricSignatureKeyOperationsIterator) Next() bool {
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

func (it *AsymmetricSignatureKeyOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

func (it *AsymmetricSignatureKeyOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *AsymmetricSignatureKeyOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *AsymmetricSignatureKeyOperationsIterator) Error() error {
	return it.err
}

// SetAccessBindings implements asymmetricsignature.AsymmetricSignatureKeyServiceClient
func (c *AsymmetricSignatureKeyServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return asymmetricsignature.NewAsymmetricSignatureKeyServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Update implements asymmetricsignature.AsymmetricSignatureKeyServiceClient
func (c *AsymmetricSignatureKeyServiceClient) Update(ctx context.Context, in *asymmetricsignature.UpdateAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return asymmetricsignature.NewAsymmetricSignatureKeyServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements asymmetricsignature.AsymmetricSignatureKeyServiceClient
func (c *AsymmetricSignatureKeyServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return asymmetricsignature.NewAsymmetricSignatureKeyServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}