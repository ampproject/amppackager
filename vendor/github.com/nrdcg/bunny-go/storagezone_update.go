package bunny

import (
	"context"
	"fmt"
)

// StorageZoneUpdateOptions represents the request parameters for the Update Storage
// Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_updatepullzone
type StorageZoneUpdateOptions struct {
	// NOTE: the naming in the Bunny API for this property is inconsistent.
	// In the update call its `ReplicationZones` but everywhere else its
	// referred to as `ReplicationRegions`.
	ReplicationRegions []string `json:"ReplicationZones,omitempty"`
	OriginURL          *string  `json:"OriginUrl,omitempty"`
	Custom404FilePath  *string  `json:"Custom404FilePath,omitempty"`
	Rewrite404To200    *bool    `json:"Rewrite404To200,omitempty"`
}

// Update changes the configuration the Storage-Zone with the given ID.
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_updatepullzone
func (s *StorageZoneService) Update(ctx context.Context, id int64, opts *StorageZoneUpdateOptions) error {
	path := fmt.Sprintf("storagezone/%d", id)
	return resourcePost(ctx, s.client, path, opts)
}
