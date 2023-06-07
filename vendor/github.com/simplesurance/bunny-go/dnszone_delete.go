package bunny

import (
	"context"
	"fmt"
)

// Delete removes the DNS Zone with the given id.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_delete
func (s *DNSZoneService) Delete(ctx context.Context, id int64) error {
	path := fmt.Sprintf("dnszone/%d", id)
	return resourceDelete(ctx, s.client, path, nil)
}
