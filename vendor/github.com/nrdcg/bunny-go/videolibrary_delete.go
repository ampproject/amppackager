package bunny

import (
	"context"
	"fmt"
)

// Delete removes the Video Library with the given id.
//
// Bunny.net API docs: https://docs.bunny.net/reference/videolibrarypublic_delete
func (s *VideoLibraryService) Delete(ctx context.Context, id int64) error {
	path := fmt.Sprintf("videolibrary/%d", id)
	return resourceDelete(ctx, s.client, path, nil)
}
