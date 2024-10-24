// Code generated by sdkgen. DO NOT EDIT.

// nolint
package billing

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	billing "github.com/yandex-cloud/go-genproto/yandex/cloud/billing/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// BillingAccountServiceClient is a billing.BillingAccountServiceClient with
// lazy GRPC connection initialization.
type BillingAccountServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// BindBillableObject implements billing.BillingAccountServiceClient
func (c *BillingAccountServiceClient) BindBillableObject(ctx context.Context, in *billing.BindBillableObjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return billing.NewBillingAccountServiceClient(conn).BindBillableObject(ctx, in, opts...)
}

// Get implements billing.BillingAccountServiceClient
func (c *BillingAccountServiceClient) Get(ctx context.Context, in *billing.GetBillingAccountRequest, opts ...grpc.CallOption) (*billing.BillingAccount, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return billing.NewBillingAccountServiceClient(conn).Get(ctx, in, opts...)
}

// List implements billing.BillingAccountServiceClient
func (c *BillingAccountServiceClient) List(ctx context.Context, in *billing.ListBillingAccountsRequest, opts ...grpc.CallOption) (*billing.ListBillingAccountsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return billing.NewBillingAccountServiceClient(conn).List(ctx, in, opts...)
}

type BillingAccountIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *BillingAccountServiceClient
	request *billing.ListBillingAccountsRequest

	items []*billing.BillingAccount
}

func (c *BillingAccountServiceClient) BillingAccountIterator(ctx context.Context, req *billing.ListBillingAccountsRequest, opts ...grpc.CallOption) *BillingAccountIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &BillingAccountIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *BillingAccountIterator) Next() bool {
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

	it.items = response.BillingAccounts
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *BillingAccountIterator) Take(size int64) ([]*billing.BillingAccount, error) {
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

	var result []*billing.BillingAccount

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *BillingAccountIterator) TakeAll() ([]*billing.BillingAccount, error) {
	return it.Take(0)
}

func (it *BillingAccountIterator) Value() *billing.BillingAccount {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *BillingAccountIterator) Error() error {
	return it.err
}

// ListAccessBindings implements billing.BillingAccountServiceClient
func (c *BillingAccountServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return billing.NewBillingAccountServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type BillingAccountAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *BillingAccountServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *BillingAccountServiceClient) BillingAccountAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *BillingAccountAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &BillingAccountAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *BillingAccountAccessBindingsIterator) Next() bool {
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

func (it *BillingAccountAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
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

func (it *BillingAccountAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *BillingAccountAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *BillingAccountAccessBindingsIterator) Error() error {
	return it.err
}

// ListBillableObjectBindings implements billing.BillingAccountServiceClient
func (c *BillingAccountServiceClient) ListBillableObjectBindings(ctx context.Context, in *billing.ListBillableObjectBindingsRequest, opts ...grpc.CallOption) (*billing.ListBillableObjectBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return billing.NewBillingAccountServiceClient(conn).ListBillableObjectBindings(ctx, in, opts...)
}

type BillingAccountBillableObjectBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *BillingAccountServiceClient
	request *billing.ListBillableObjectBindingsRequest

	items []*billing.BillableObjectBinding
}

func (c *BillingAccountServiceClient) BillingAccountBillableObjectBindingsIterator(ctx context.Context, req *billing.ListBillableObjectBindingsRequest, opts ...grpc.CallOption) *BillingAccountBillableObjectBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &BillingAccountBillableObjectBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *BillingAccountBillableObjectBindingsIterator) Next() bool {
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

	response, err := it.client.ListBillableObjectBindings(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.BillableObjectBindings
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *BillingAccountBillableObjectBindingsIterator) Take(size int64) ([]*billing.BillableObjectBinding, error) {
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

	var result []*billing.BillableObjectBinding

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *BillingAccountBillableObjectBindingsIterator) TakeAll() ([]*billing.BillableObjectBinding, error) {
	return it.Take(0)
}

func (it *BillingAccountBillableObjectBindingsIterator) Value() *billing.BillableObjectBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *BillingAccountBillableObjectBindingsIterator) Error() error {
	return it.err
}

// UpdateAccessBindings implements billing.BillingAccountServiceClient
func (c *BillingAccountServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return billing.NewBillingAccountServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}
