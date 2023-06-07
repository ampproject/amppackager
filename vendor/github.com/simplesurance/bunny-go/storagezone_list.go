package bunny

import "context"

// StorageZones represents the response of the List Storage Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/storagezonepublic_index
type StorageZones PaginationReply[StorageZone]

// List retrieves the Storage Zones.
// If opts is nil, DefaultPaginationPerPage and DefaultPaginationPage will be used.
// if opts.Page or or opts.PerPage is < 1, the related DefaultPagination values are used.
//
// Bunny.net API docs: https://docs.bunny.net/reference/storagezonepublic_index
func (s *StorageZoneService) List(
	ctx context.Context,
	opts *PaginationOptions,
) (*StorageZones, error) {
	return resourceList[StorageZones](ctx, s.client, "/storagezone", opts)
}
