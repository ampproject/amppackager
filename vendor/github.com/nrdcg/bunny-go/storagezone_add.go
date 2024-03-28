package bunny

import "context"

// StorageZoneAddOptions are the request parameters for the Get Storage Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/storagezonepublic_add
type StorageZoneAddOptions struct {
	// The name of the storage zone
	Name *string `json:"Name,omitempty"`
	// The ID of the storage zone that the storage zone is linked to.
	Region *string `json:"Region,omitempty"`

	// The origin URL of the storage zone where the files are fetched from (Optional)
	OriginURL *string `json:"OriginUrl,omitempty"`
	// The code of the main storage zone region (Optional)
	ReplicationRegions []string `json:"ReplicationRegions,omitempty"`
}

// Add creates a new Storage Zone.
// opts and the non-optional parameters in the struct must be specified for a successful request.
// On success the created StorageZone is returned.
//
// Bunny.net API docs: https://docs.bunny.net/reference/storagezonepublic_add
func (s *StorageZoneService) Add(ctx context.Context, opts *StorageZoneAddOptions) (*StorageZone, error) {
	return resourcePostWithResponse[StorageZone](
		ctx,
		s.client,
		"/storagezone",
		opts,
	)
}
