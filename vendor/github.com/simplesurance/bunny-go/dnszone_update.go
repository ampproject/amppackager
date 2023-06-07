package bunny

import (
	"context"
	"fmt"
)

// DNSZoneUpdateOptions represents the request parameters for the Update DNS
// Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_update
type DNSZoneUpdateOptions struct {
	CustomNameserversEnabled      *bool   `json:"CustomNameserversEnabled,omitempty"`
	Nameserver1                   *string `json:"Nameserver1,omitempty"`
	Nameserver2                   *string `json:"Nameserver2,omitempty"`
	SoaEmail                      *string `json:"SoaEmail,omitempty"`
	LoggingEnabled                *bool   `json:"LoggingEnabled,omitempty"`
	LoggingIPAnonymizationEnabled *bool   `json:"LoggingIPAnonymizationEnabled,omitempty"`
	LogAnonymizationType          *int    `json:"LogAnonymizationType,omitempty"`
}

// Update changes the configuration the DNS Zone with the given ID.
// The updated DNS Zone is returned.
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_update
func (s *DNSZoneService) Update(ctx context.Context, id int64, opts *DNSZoneUpdateOptions) (*DNSZone, error) {
	path := fmt.Sprintf("dnszone/%d", id)
	return resourcePostWithResponse[DNSZone](
		ctx,
		s.client,
		path,
		opts,
	)
}
