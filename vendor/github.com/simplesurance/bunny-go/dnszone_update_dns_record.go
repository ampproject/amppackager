package bunny

import (
	"context"
	"fmt"
)

// UpdateDNSRecord updates a DNS record in the DNS Zone.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_updaterecord
func (s *DNSZoneService) UpdateDNSRecord(ctx context.Context, dnsZoneID int64, dnsRecordID int64, opts *AddOrUpdateDNSRecordOptions) error {
	path := fmt.Sprintf("dnszone/%d/records/%d", dnsZoneID, dnsRecordID)
	return resourcePost(ctx, s.client, path, opts)
}
