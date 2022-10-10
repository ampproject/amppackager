// Code generated by sdkgen. DO NOT EDIT.

//nolint
package mongodb

import (
	"context"

	"google.golang.org/grpc"

	mongodb "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// ClusterServiceClient is a mongodb.ClusterServiceClient with
// lazy GRPC connection initialization.
type ClusterServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// AddHosts implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) AddHosts(ctx context.Context, in *mongodb.AddClusterHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).AddHosts(ctx, in, opts...)
}

// AddShard implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) AddShard(ctx context.Context, in *mongodb.AddClusterShardRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).AddShard(ctx, in, opts...)
}

// Backup implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) Backup(ctx context.Context, in *mongodb.BackupClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).Backup(ctx, in, opts...)
}

// Create implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) Create(ctx context.Context, in *mongodb.CreateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) Delete(ctx context.Context, in *mongodb.DeleteClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).Delete(ctx, in, opts...)
}

// DeleteHosts implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) DeleteHosts(ctx context.Context, in *mongodb.DeleteClusterHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).DeleteHosts(ctx, in, opts...)
}

// DeleteShard implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) DeleteShard(ctx context.Context, in *mongodb.DeleteClusterShardRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).DeleteShard(ctx, in, opts...)
}

// EnableSharding implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) EnableSharding(ctx context.Context, in *mongodb.EnableClusterShardingRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).EnableSharding(ctx, in, opts...)
}

// Get implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) Get(ctx context.Context, in *mongodb.GetClusterRequest, opts ...grpc.CallOption) (*mongodb.Cluster, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).Get(ctx, in, opts...)
}

// GetShard implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) GetShard(ctx context.Context, in *mongodb.GetClusterShardRequest, opts ...grpc.CallOption) (*mongodb.Shard, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).GetShard(ctx, in, opts...)
}

// List implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) List(ctx context.Context, in *mongodb.ListClustersRequest, opts ...grpc.CallOption) (*mongodb.ListClustersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).List(ctx, in, opts...)
}

type ClusterIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ClusterServiceClient
	request *mongodb.ListClustersRequest

	items []*mongodb.Cluster
}

func (c *ClusterServiceClient) ClusterIterator(ctx context.Context, req *mongodb.ListClustersRequest, opts ...grpc.CallOption) *ClusterIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ClusterIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ClusterIterator) Next() bool {
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

	it.items = response.Clusters
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterIterator) Take(size int64) ([]*mongodb.Cluster, error) {
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

	var result []*mongodb.Cluster

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ClusterIterator) TakeAll() ([]*mongodb.Cluster, error) {
	return it.Take(0)
}

func (it *ClusterIterator) Value() *mongodb.Cluster {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterIterator) Error() error {
	return it.err
}

// ListBackups implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) ListBackups(ctx context.Context, in *mongodb.ListClusterBackupsRequest, opts ...grpc.CallOption) (*mongodb.ListClusterBackupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).ListBackups(ctx, in, opts...)
}

type ClusterBackupsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ClusterServiceClient
	request *mongodb.ListClusterBackupsRequest

	items []*mongodb.Backup
}

func (c *ClusterServiceClient) ClusterBackupsIterator(ctx context.Context, req *mongodb.ListClusterBackupsRequest, opts ...grpc.CallOption) *ClusterBackupsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ClusterBackupsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ClusterBackupsIterator) Next() bool {
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

	response, err := it.client.ListBackups(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Backups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterBackupsIterator) Take(size int64) ([]*mongodb.Backup, error) {
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

	var result []*mongodb.Backup

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ClusterBackupsIterator) TakeAll() ([]*mongodb.Backup, error) {
	return it.Take(0)
}

func (it *ClusterBackupsIterator) Value() *mongodb.Backup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterBackupsIterator) Error() error {
	return it.err
}

// ListHosts implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) ListHosts(ctx context.Context, in *mongodb.ListClusterHostsRequest, opts ...grpc.CallOption) (*mongodb.ListClusterHostsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).ListHosts(ctx, in, opts...)
}

type ClusterHostsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ClusterServiceClient
	request *mongodb.ListClusterHostsRequest

	items []*mongodb.Host
}

func (c *ClusterServiceClient) ClusterHostsIterator(ctx context.Context, req *mongodb.ListClusterHostsRequest, opts ...grpc.CallOption) *ClusterHostsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ClusterHostsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ClusterHostsIterator) Next() bool {
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

	response, err := it.client.ListHosts(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Hosts
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterHostsIterator) Take(size int64) ([]*mongodb.Host, error) {
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

	var result []*mongodb.Host

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ClusterHostsIterator) TakeAll() ([]*mongodb.Host, error) {
	return it.Take(0)
}

func (it *ClusterHostsIterator) Value() *mongodb.Host {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterHostsIterator) Error() error {
	return it.err
}

// ListLogs implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) ListLogs(ctx context.Context, in *mongodb.ListClusterLogsRequest, opts ...grpc.CallOption) (*mongodb.ListClusterLogsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).ListLogs(ctx, in, opts...)
}

type ClusterLogsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ClusterServiceClient
	request *mongodb.ListClusterLogsRequest

	items []*mongodb.LogRecord
}

func (c *ClusterServiceClient) ClusterLogsIterator(ctx context.Context, req *mongodb.ListClusterLogsRequest, opts ...grpc.CallOption) *ClusterLogsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ClusterLogsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ClusterLogsIterator) Next() bool {
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

	response, err := it.client.ListLogs(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Logs
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterLogsIterator) Take(size int64) ([]*mongodb.LogRecord, error) {
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

	var result []*mongodb.LogRecord

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ClusterLogsIterator) TakeAll() ([]*mongodb.LogRecord, error) {
	return it.Take(0)
}

func (it *ClusterLogsIterator) Value() *mongodb.LogRecord {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterLogsIterator) Error() error {
	return it.err
}

// ListOperations implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) ListOperations(ctx context.Context, in *mongodb.ListClusterOperationsRequest, opts ...grpc.CallOption) (*mongodb.ListClusterOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).ListOperations(ctx, in, opts...)
}

type ClusterOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ClusterServiceClient
	request *mongodb.ListClusterOperationsRequest

	items []*operation.Operation
}

func (c *ClusterServiceClient) ClusterOperationsIterator(ctx context.Context, req *mongodb.ListClusterOperationsRequest, opts ...grpc.CallOption) *ClusterOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ClusterOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ClusterOperationsIterator) Next() bool {
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

func (it *ClusterOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

func (it *ClusterOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *ClusterOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterOperationsIterator) Error() error {
	return it.err
}

// ListShards implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) ListShards(ctx context.Context, in *mongodb.ListClusterShardsRequest, opts ...grpc.CallOption) (*mongodb.ListClusterShardsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).ListShards(ctx, in, opts...)
}

type ClusterShardsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ClusterServiceClient
	request *mongodb.ListClusterShardsRequest

	items []*mongodb.Shard
}

func (c *ClusterServiceClient) ClusterShardsIterator(ctx context.Context, req *mongodb.ListClusterShardsRequest, opts ...grpc.CallOption) *ClusterShardsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ClusterShardsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ClusterShardsIterator) Next() bool {
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

	response, err := it.client.ListShards(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Shards
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterShardsIterator) Take(size int64) ([]*mongodb.Shard, error) {
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

	var result []*mongodb.Shard

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ClusterShardsIterator) TakeAll() ([]*mongodb.Shard, error) {
	return it.Take(0)
}

func (it *ClusterShardsIterator) Value() *mongodb.Shard {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterShardsIterator) Error() error {
	return it.err
}

// Move implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) Move(ctx context.Context, in *mongodb.MoveClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).Move(ctx, in, opts...)
}

// RescheduleMaintenance implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) RescheduleMaintenance(ctx context.Context, in *mongodb.RescheduleMaintenanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).RescheduleMaintenance(ctx, in, opts...)
}

// ResetupHosts implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) ResetupHosts(ctx context.Context, in *mongodb.ResetupHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).ResetupHosts(ctx, in, opts...)
}

// RestartHosts implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) RestartHosts(ctx context.Context, in *mongodb.RestartHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).RestartHosts(ctx, in, opts...)
}

// Restore implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) Restore(ctx context.Context, in *mongodb.RestoreClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).Restore(ctx, in, opts...)
}

// Start implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) Start(ctx context.Context, in *mongodb.StartClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).Start(ctx, in, opts...)
}

// StepdownHosts implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) StepdownHosts(ctx context.Context, in *mongodb.StepdownHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).StepdownHosts(ctx, in, opts...)
}

// Stop implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) Stop(ctx context.Context, in *mongodb.StopClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).Stop(ctx, in, opts...)
}

// StreamLogs implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) StreamLogs(ctx context.Context, in *mongodb.StreamClusterLogsRequest, opts ...grpc.CallOption) (mongodb.ClusterService_StreamLogsClient, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).StreamLogs(ctx, in, opts...)
}

// Update implements mongodb.ClusterServiceClient
func (c *ClusterServiceClient) Update(ctx context.Context, in *mongodb.UpdateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewClusterServiceClient(conn).Update(ctx, in, opts...)
}
