package bunny

import (
	"context"
	"fmt"
)

// DeleteDNSRecord removes a DNS Record of a DNS Zone.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_deleterecord
func (s *DNSZoneService) DeleteDNSRecord(ctx context.Context, dnsZoneID int64, dnsRecordID int64) error {
	path := fmt.Sprintf("dnszone/%d/records/%d", dnsZoneID, dnsRecordID)
	return resourceDelete(ctx, s.client, path, nil)
}
