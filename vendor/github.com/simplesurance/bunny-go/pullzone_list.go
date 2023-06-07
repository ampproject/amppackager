package bunny

import "context"

// PullZones represents the response of the List Pull Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_index
type PullZones PaginationReply[PullZone]

// List retrieves the Pull Zones.
// If opts is nil, DefaultPaginationPerPage and DefaultPaginationPage will be used.
// if opts.Page or or opts.PerPage is < 1, the related DefaultPagination values are used.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_index
func (s *PullZoneService) List(
	ctx context.Context,
	opts *PaginationOptions,
) (*PullZones, error) {
	return resourceList[PullZones](ctx, s.client, "/pullzone", opts)
}
