package bunny

import (
	"context"
	"fmt"
)

// Delete removes the Pull Zone with the given id.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_delete
func (s *PullZoneService) Delete(ctx context.Context, id int64) error {
	path := fmt.Sprintf("pullzone/%d", id)
	return resourceDelete(ctx, s.client, path, nil)
}
