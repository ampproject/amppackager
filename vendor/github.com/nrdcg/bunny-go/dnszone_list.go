package bunny

import "context"

// DNSZones represents the response of the List DNS Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_index
type DNSZones PaginationReply[DNSZone]

// List retrieves the DNS Zones.
// If opts is nil, DefaultPaginationPerPage and DefaultPaginationPage will be used.
// if opts.Page or or opts.PerPage is < 1, the related DefaultPagination values are used.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_index
func (s *DNSZoneService) List(
	ctx context.Context,
	opts *PaginationOptions,
) (*DNSZones, error) {
	return resourceList[DNSZones](ctx, s.client, "/dnszone", opts)
}
