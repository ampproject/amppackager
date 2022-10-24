// Code generated by sdkgen. DO NOT EDIT.

//nolint
package kms

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	kms "github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// SymmetricKeyServiceClient is a kms.SymmetricKeyServiceClient with
// lazy GRPC connection initialization.
type SymmetricKeyServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// CancelVersionDestruction implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) CancelVersionDestruction(ctx context.Context, in *kms.CancelSymmetricKeyVersionDestructionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).CancelVersionDestruction(ctx, in, opts...)
}

// Create implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) Create(ctx context.Context, in *kms.CreateSymmetricKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) Delete(ctx context.Context, in *kms.DeleteSymmetricKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) Get(ctx context.Context, in *kms.GetSymmetricKeyRequest, opts ...grpc.CallOption) (*kms.SymmetricKey, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).Get(ctx, in, opts...)
}

// List implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) List(ctx context.Context, in *kms.ListSymmetricKeysRequest, opts ...grpc.CallOption) (*kms.ListSymmetricKeysResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).List(ctx, in, opts...)
}

type SymmetricKeyIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SymmetricKeyServiceClient
	request *kms.ListSymmetricKeysRequest

	items []*kms.SymmetricKey
}

func (c *SymmetricKeyServiceClient) SymmetricKeyIterator(ctx context.Context, req *kms.ListSymmetricKeysRequest, opts ...grpc.CallOption) *SymmetricKeyIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SymmetricKeyIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SymmetricKeyIterator) Next() bool {
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

func (it *SymmetricKeyIterator) Take(size int64) ([]*kms.SymmetricKey, error) {
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

	var result []*kms.SymmetricKey

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *SymmetricKeyIterator) TakeAll() ([]*kms.SymmetricKey, error) {
	return it.Take(0)
}

func (it *SymmetricKeyIterator) Value() *kms.SymmetricKey {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SymmetricKeyIterator) Error() error {
	return it.err
}

// ListAccessBindings implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type SymmetricKeyAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SymmetricKeyServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *SymmetricKeyServiceClient) SymmetricKeyAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *SymmetricKeyAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SymmetricKeyAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SymmetricKeyAccessBindingsIterator) Next() bool {
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

func (it *SymmetricKeyAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
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

func (it *SymmetricKeyAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *SymmetricKeyAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SymmetricKeyAccessBindingsIterator) Error() error {
	return it.err
}

// ListOperations implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) ListOperations(ctx context.Context, in *kms.ListSymmetricKeyOperationsRequest, opts ...grpc.CallOption) (*kms.ListSymmetricKeyOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).ListOperations(ctx, in, opts...)
}

type SymmetricKeyOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SymmetricKeyServiceClient
	request *kms.ListSymmetricKeyOperationsRequest

	items []*operation.Operation
}

func (c *SymmetricKeyServiceClient) SymmetricKeyOperationsIterator(ctx context.Context, req *kms.ListSymmetricKeyOperationsRequest, opts ...grpc.CallOption) *SymmetricKeyOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SymmetricKeyOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SymmetricKeyOperationsIterator) Next() bool {
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

func (it *SymmetricKeyOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

func (it *SymmetricKeyOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *SymmetricKeyOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SymmetricKeyOperationsIterator) Error() error {
	return it.err
}

// ListVersions implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) ListVersions(ctx context.Context, in *kms.ListSymmetricKeyVersionsRequest, opts ...grpc.CallOption) (*kms.ListSymmetricKeyVersionsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).ListVersions(ctx, in, opts...)
}

type SymmetricKeyVersionsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *SymmetricKeyServiceClient
	request *kms.ListSymmetricKeyVersionsRequest

	items []*kms.SymmetricKeyVersion
}

func (c *SymmetricKeyServiceClient) SymmetricKeyVersionsIterator(ctx context.Context, req *kms.ListSymmetricKeyVersionsRequest, opts ...grpc.CallOption) *SymmetricKeyVersionsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &SymmetricKeyVersionsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *SymmetricKeyVersionsIterator) Next() bool {
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

	response, err := it.client.ListVersions(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.KeyVersions
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *SymmetricKeyVersionsIterator) Take(size int64) ([]*kms.SymmetricKeyVersion, error) {
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

	var result []*kms.SymmetricKeyVersion

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *SymmetricKeyVersionsIterator) TakeAll() ([]*kms.SymmetricKeyVersion, error) {
	return it.Take(0)
}

func (it *SymmetricKeyVersionsIterator) Value() *kms.SymmetricKeyVersion {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *SymmetricKeyVersionsIterator) Error() error {
	return it.err
}

// Rotate implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) Rotate(ctx context.Context, in *kms.RotateSymmetricKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).Rotate(ctx, in, opts...)
}

// ScheduleVersionDestruction implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) ScheduleVersionDestruction(ctx context.Context, in *kms.ScheduleSymmetricKeyVersionDestructionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).ScheduleVersionDestruction(ctx, in, opts...)
}

// SetAccessBindings implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// SetPrimaryVersion implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) SetPrimaryVersion(ctx context.Context, in *kms.SetPrimarySymmetricKeyVersionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).SetPrimaryVersion(ctx, in, opts...)
}

// Update implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) Update(ctx context.Context, in *kms.UpdateSymmetricKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements kms.SymmetricKeyServiceClient
func (c *SymmetricKeyServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kms.NewSymmetricKeyServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}