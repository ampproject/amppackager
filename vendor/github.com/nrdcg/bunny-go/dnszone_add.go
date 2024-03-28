package bunny

import "context"

// Add creates a new DNS Zone.
// opts and the non-optional parameters in the struct must be specified for a successful request.
// On success the created DNSZone is returned.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_add
func (s *DNSZoneService) Add(ctx context.Context, opts *DNSZone) (*DNSZone, error) {
	return resourcePostWithResponse[DNSZone](
		ctx,
		s.client,
		"/dnszone",
		opts,
	)
}
