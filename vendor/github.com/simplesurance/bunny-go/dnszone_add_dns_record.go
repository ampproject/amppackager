package bunny

import (
	"context"
	"fmt"
)

// AddOrUpdateDNSRecordOptions represents the message that is sent to the
// Add DNS Record API Endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_addrecord
type AddOrUpdateDNSRecordOptions struct {
	ID                     *int64                  `json:"Id,omitempty"`
	Type                   *int                    `json:"Type,omitempty"`
	TTL                    *int32                  `json:"Ttl,omitempty"`
	Value                  *string                 `json:"Value,omitempty"`
	Name                   *string                 `json:"Name,omitempty"`
	Weight                 *int32                  `json:"Weight,omitempty"`
	Priority               *int32                  `json:"Priority,omitempty"`
	Flags                  *int                    `json:"Flags,omitempty"`
	Tag                    *string                 `json:"Tag,omitempty"`
	Port                   *int32                  `json:"Port,omitempty"`
	PullZoneID             *int64                  `json:"PullZoneId,omitempty"`
	ScriptID               *int64                  `json:"ScriptId,omitempty"`
	Accelerated            *bool                   `json:"Accelerated,omitempty"`
	MonitorType            *int                    `json:"MonitorType,omitempty"`
	GeolocationLatitude    *float64                `json:"GeolocationLatitude,omitempty"`
	GeolocationLongitude   *float64                `json:"GeolocationLongitude,omitempty"`
	LatencyZone            *string                 `json:"LatencyZone,omitempty"`
	SmartRoutingType       *int                    `json:"SmartRoutingType,omitempty"`
	Disabled               *bool                   `json:"Disabled,omitempty"`
	EnvironmentalVariables []EnvironmentalVariable `json:"EnvironmentalVariables,omitempty"`
}

// AddDNSRecord adds a DNS record to the DNS Zone.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_addrecord
func (s *DNSZoneService) AddDNSRecord(ctx context.Context, dnsZoneID int64, opts *AddOrUpdateDNSRecordOptions) (*DNSRecord, error) {
	path := fmt.Sprintf("dnszone/%d/records", dnsZoneID)
	return resourcePutWithResponse[DNSRecord](
		ctx,
		s.client,
		path,
		opts,
	)
}
