package bunny

import (
	"context"
	"fmt"
)

// SetEdgeRuleEnabledOptions represents the message that is sent to Add/Update Edge Rule endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addedgerule
type SetEdgeRuleEnabledOptions struct {
	// ID must be set to the PullZone ID for that the EdgeRule should be enabled.
	ID    *int64 `json:"Id,omitempty"`
	Value *bool  `json:"Value,omitempty"`
}

// SetEdgeRuleEnabled enables or disables an Edge Rule of a Pull Zone.
// The edgeRuleGUID field is called edgeRuleID in the API message and
// documentation. It is the same then the GUID field in the EdgeRule message.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addedgerule
func (s *PullZoneService) SetEdgeRuleEnabled(ctx context.Context, pullZoneID int64, edgeRuleGUID string, opts *SetEdgeRuleEnabledOptions) error {
	if opts != nil {
		if opts.ID == nil {
			s.client.logf("SetEdgeRuleEnabled: ID field is unset in SetEdgeRuleEnabledOptions")
		} else if *opts.ID != pullZoneID {
			s.client.logf("SetEdgeRuleEnabled: mismatched pullZoneID %d and SetEdgeRuleEnabledOptions.ID %d were passed, values should be equal", pullZoneID, *opts.ID)
		}
	}

	path := fmt.Sprintf("pullzone/%d/edgerules/%s/setEdgeRuleEnabled", pullZoneID, edgeRuleGUID)
	return resourcePost(ctx, s.client, path, opts)
}
