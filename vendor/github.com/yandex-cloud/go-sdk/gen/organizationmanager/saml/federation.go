// Code generated by sdkgen. DO NOT EDIT.

//nolint
package saml

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	organizationmanager "github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	saml "github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1/saml"
)

//revive:disable

// FederationServiceClient is a saml.FederationServiceClient with
// lazy GRPC connection initialization.
type FederationServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// AddUserAccounts implements saml.FederationServiceClient
func (c *FederationServiceClient) AddUserAccounts(ctx context.Context, in *saml.AddFederatedUserAccountsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewFederationServiceClient(conn).AddUserAccounts(ctx, in, opts...)
}

// Create implements saml.FederationServiceClient
func (c *FederationServiceClient) Create(ctx context.Context, in *saml.CreateFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewFederationServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements saml.FederationServiceClient
func (c *FederationServiceClient) Delete(ctx context.Context, in *saml.DeleteFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewFederationServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements saml.FederationServiceClient
func (c *FederationServiceClient) Get(ctx context.Context, in *saml.GetFederationRequest, opts ...grpc.CallOption) (*saml.Federation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewFederationServiceClient(conn).Get(ctx, in, opts...)
}

// List implements saml.FederationServiceClient
func (c *FederationServiceClient) List(ctx context.Context, in *saml.ListFederationsRequest, opts ...grpc.CallOption) (*saml.ListFederationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewFederationServiceClient(conn).List(ctx, in, opts...)
}

type FederationIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *FederationServiceClient
	request *saml.ListFederationsRequest

	items []*saml.Federation
}

func (c *FederationServiceClient) FederationIterator(ctx context.Context, req *saml.ListFederationsRequest, opts ...grpc.CallOption) *FederationIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &FederationIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *FederationIterator) Next() bool {
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

	it.items = response.Federations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *FederationIterator) Take(size int64) ([]*saml.Federation, error) {
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

	var result []*saml.Federation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *FederationIterator) TakeAll() ([]*saml.Federation, error) {
	return it.Take(0)
}

func (it *FederationIterator) Value() *saml.Federation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *FederationIterator) Error() error {
	return it.err
}

// ListOperations implements saml.FederationServiceClient
func (c *FederationServiceClient) ListOperations(ctx context.Context, in *saml.ListFederationOperationsRequest, opts ...grpc.CallOption) (*saml.ListFederationOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewFederationServiceClient(conn).ListOperations(ctx, in, opts...)
}

type FederationOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *FederationServiceClient
	request *saml.ListFederationOperationsRequest

	items []*operation.Operation
}

func (c *FederationServiceClient) FederationOperationsIterator(ctx context.Context, req *saml.ListFederationOperationsRequest, opts ...grpc.CallOption) *FederationOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &FederationOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *FederationOperationsIterator) Next() bool {
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

func (it *FederationOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

func (it *FederationOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *FederationOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *FederationOperationsIterator) Error() error {
	return it.err
}

// ListUserAccounts implements saml.FederationServiceClient
func (c *FederationServiceClient) ListUserAccounts(ctx context.Context, in *saml.ListFederatedUserAccountsRequest, opts ...grpc.CallOption) (*saml.ListFederatedUserAccountsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewFederationServiceClient(conn).ListUserAccounts(ctx, in, opts...)
}

type FederationUserAccountsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *FederationServiceClient
	request *saml.ListFederatedUserAccountsRequest

	items []*organizationmanager.UserAccount
}

func (c *FederationServiceClient) FederationUserAccountsIterator(ctx context.Context, req *saml.ListFederatedUserAccountsRequest, opts ...grpc.CallOption) *FederationUserAccountsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &FederationUserAccountsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *FederationUserAccountsIterator) Next() bool {
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

	response, err := it.client.ListUserAccounts(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.UserAccounts
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *FederationUserAccountsIterator) Take(size int64) ([]*organizationmanager.UserAccount, error) {
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

	var result []*organizationmanager.UserAccount

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *FederationUserAccountsIterator) TakeAll() ([]*organizationmanager.UserAccount, error) {
	return it.Take(0)
}

func (it *FederationUserAccountsIterator) Value() *organizationmanager.UserAccount {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *FederationUserAccountsIterator) Error() error {
	return it.err
}

// Update implements saml.FederationServiceClient
func (c *FederationServiceClient) Update(ctx context.Context, in *saml.UpdateFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewFederationServiceClient(conn).Update(ctx, in, opts...)
}
