package bunny

import (
	"context"
	"fmt"
)

// AddOrUpdateEdgeRuleOptions is the message that is sent to the
// Add/Update Edge Rule API Endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addedgerule
type AddOrUpdateEdgeRuleOptions struct {
	// GUID must only be set when updating an Edge Rule. When creating an
	// Edge Rule it must be unset. The API Endpoint will generate a GUID.
	GUID                *string            `json:"Guid,omitempty"`
	ActionType          *int               `json:"ActionType,omitempty"`
	ActionParameter1    *string            `json:"ActionParameter1,omitempty"`
	ActionParameter2    *string            `json:"ActionParameter2,omitempty"`
	Triggers            []*EdgeRuleTrigger `json:"Triggers,omitempty"`
	TriggerMatchingType *int               `json:"TriggerMatchingType,omitempty"`
	Description         *string            `json:"Description,omitempty"`
	Enabled             *bool              `json:"Enabled,omitempty"`
}

// AddOrUpdateEdgeRule adds or updates an Edge Rule of a Pull Zone.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addedgerule
func (s *PullZoneService) AddOrUpdateEdgeRule(ctx context.Context, pullZoneID int64, opts *AddOrUpdateEdgeRuleOptions) error {
	path := fmt.Sprintf("pullzone/%d/edgerules/addOrUpdate", pullZoneID)
	return resourcePost(ctx, s.client, path, opts)
}
