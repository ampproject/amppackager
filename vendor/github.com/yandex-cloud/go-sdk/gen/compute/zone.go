// Code generated by sdkgen. DO NOT EDIT.

//nolint
package compute

import (
	"context"

	"google.golang.org/grpc"

	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

//revive:disable

// ZoneServiceClient is a compute.ZoneServiceClient with
// lazy GRPC connection initialization.
type ZoneServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Get implements compute.ZoneServiceClient
func (c *ZoneServiceClient) Get(ctx context.Context, in *compute.GetZoneRequest, opts ...grpc.CallOption) (*compute.Zone, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewZoneServiceClient(conn).Get(ctx, in, opts...)
}

// List implements compute.ZoneServiceClient
func (c *ZoneServiceClient) List(ctx context.Context, in *compute.ListZonesRequest, opts ...grpc.CallOption) (*compute.ListZonesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewZoneServiceClient(conn).List(ctx, in, opts...)
}

type ZoneIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ZoneServiceClient
	request *compute.ListZonesRequest

	items []*compute.Zone
}

func (c *ZoneServiceClient) ZoneIterator(ctx context.Context, req *compute.ListZonesRequest, opts ...grpc.CallOption) *ZoneIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ZoneIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ZoneIterator) Next() bool {
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

	it.items = response.Zones
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ZoneIterator) Take(size int64) ([]*compute.Zone, error) {
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

	var result []*compute.Zone

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ZoneIterator) TakeAll() ([]*compute.Zone, error) {
	return it.Take(0)
}

func (it *ZoneIterator) Value() *compute.Zone {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ZoneIterator) Error() error {
	return it.err
}
