package bunny

import (
	"context"
	"fmt"
)

// Delete removes the Storage Zone with the given id.
//
// Bunny.net API docs: https://docs.bunny.net/reference/storagezonepublic_delete
func (s *StorageZoneService) Delete(ctx context.Context, id int64) error {
	path := fmt.Sprintf("storagezone/%d", id)
	return resourceDelete(ctx, s.client, path, nil)
}
