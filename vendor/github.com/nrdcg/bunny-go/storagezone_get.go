package bunny

import (
	"context"
	"fmt"
)

// StorageZone represents the response of the the List and Get Storage Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/storagezonepublic_index2 https://docs.bunny.net/reference/storagezonepublic_index
type StorageZone struct {
	ID *int64 `json:"Id,omitempty"`

	UserID             *string     `json:"UserId,omitempty"`
	Name               *string     `json:"Name,omitempty"`
	Password           *string     `json:"Password,omitempty"`
	DateModified       *string     `json:"DateModified,omitempty"`
	Deleted            *bool       `json:"Deleted,omitempty"`
	StorageUsed        *int64      `json:"StorageUsed,omitempty"`
	FilesStored        *int64      `json:"FilesStored,omitempty"`
	Region             *string     `json:"Region,omitempty"`
	ReplicationRegions []string    `json:"ReplicationRegions,omitempty"`
	PullZones          []*PullZone `json:"PullZones,omitempty"`
	ReadOnlyPassword   *string     `json:"ReadOnlyPassword,omitempty"`
}

// Get retrieves the Storage Zone with the given id.
//
// Bunny.net API docs: https://docs.bunny.net/reference/storagezonepublic_index2
func (s *StorageZoneService) Get(ctx context.Context, id int64) (*StorageZone, error) {
	path := fmt.Sprintf("storagezone/%d", id)
	return resourceGet[StorageZone](ctx, s.client, path, nil)
}
