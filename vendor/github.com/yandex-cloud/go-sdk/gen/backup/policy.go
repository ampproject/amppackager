// Code generated by sdkgen. DO NOT EDIT.

// nolint
package backup

import (
	"context"

	"google.golang.org/grpc"

	backup "github.com/yandex-cloud/go-genproto/yandex/cloud/backup/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// PolicyServiceClient is a backup.PolicyServiceClient with
// lazy GRPC connection initialization.
type PolicyServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Apply implements backup.PolicyServiceClient
func (c *PolicyServiceClient) Apply(ctx context.Context, in *backup.ApplyPolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return backup.NewPolicyServiceClient(conn).Apply(ctx, in, opts...)
}

// Create implements backup.PolicyServiceClient
func (c *PolicyServiceClient) Create(ctx context.Context, in *backup.CreatePolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return backup.NewPolicyServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements backup.PolicyServiceClient
func (c *PolicyServiceClient) Delete(ctx context.Context, in *backup.DeletePolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return backup.NewPolicyServiceClient(conn).Delete(ctx, in, opts...)
}

// Execute implements backup.PolicyServiceClient
func (c *PolicyServiceClient) Execute(ctx context.Context, in *backup.ExecuteRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return backup.NewPolicyServiceClient(conn).Execute(ctx, in, opts...)
}

// Get implements backup.PolicyServiceClient
func (c *PolicyServiceClient) Get(ctx context.Context, in *backup.GetPolicyRequest, opts ...grpc.CallOption) (*backup.Policy, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return backup.NewPolicyServiceClient(conn).Get(ctx, in, opts...)
}

// List implements backup.PolicyServiceClient
func (c *PolicyServiceClient) List(ctx context.Context, in *backup.ListPoliciesRequest, opts ...grpc.CallOption) (*backup.ListPoliciesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return backup.NewPolicyServiceClient(conn).List(ctx, in, opts...)
}

type PolicyIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *PolicyServiceClient
	request *backup.ListPoliciesRequest

	items []*backup.Policy
}

func (c *PolicyServiceClient) PolicyIterator(ctx context.Context, req *backup.ListPoliciesRequest, opts ...grpc.CallOption) *PolicyIterator {
	var pageSize int64
	const defaultPageSize = 1000

	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &PolicyIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *PolicyIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Policies
	return len(it.items) > 0
}

func (it *PolicyIterator) Take(size int64) ([]*backup.Policy, error) {
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

	var result []*backup.Policy

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *PolicyIterator) TakeAll() ([]*backup.Policy, error) {
	return it.Take(0)
}

func (it *PolicyIterator) Value() *backup.Policy {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *PolicyIterator) Error() error {
	return it.err
}

// ListApplications implements backup.PolicyServiceClient
func (c *PolicyServiceClient) ListApplications(ctx context.Context, in *backup.ListApplicationsRequest, opts ...grpc.CallOption) (*backup.ListApplicationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return backup.NewPolicyServiceClient(conn).ListApplications(ctx, in, opts...)
}

type PolicyApplicationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *PolicyServiceClient
	request *backup.ListApplicationsRequest

	items []*backup.PolicyApplication
}

func (c *PolicyServiceClient) PolicyApplicationsIterator(ctx context.Context, req *backup.ListApplicationsRequest, opts ...grpc.CallOption) *PolicyApplicationsIterator {
	var pageSize int64
	const defaultPageSize = 1000

	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &PolicyApplicationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *PolicyApplicationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started {
		return false
	}
	it.started = true

	response, err := it.client.ListApplications(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Applications
	return len(it.items) > 0
}

func (it *PolicyApplicationsIterator) Take(size int64) ([]*backup.PolicyApplication, error) {
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

	var result []*backup.PolicyApplication

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *PolicyApplicationsIterator) TakeAll() ([]*backup.PolicyApplication, error) {
	return it.Take(0)
}

func (it *PolicyApplicationsIterator) Value() *backup.PolicyApplication {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *PolicyApplicationsIterator) Error() error {
	return it.err
}

// Revoke implements backup.PolicyServiceClient
func (c *PolicyServiceClient) Revoke(ctx context.Context, in *backup.RevokeRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return backup.NewPolicyServiceClient(conn).Revoke(ctx, in, opts...)
}

// Update implements backup.PolicyServiceClient
func (c *PolicyServiceClient) Update(ctx context.Context, in *backup.UpdatePolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return backup.NewPolicyServiceClient(conn).Update(ctx, in, opts...)
}
