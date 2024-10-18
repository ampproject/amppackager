package bunny

import (
	"context"
	"fmt"
)

// DeleteEdgeRule removes an Edge Rule of a Pull Zone.
// The edgeRuleGUID field is called edgeRuleID in the API message and
// documentation. It is the same then the GUID field in the EdgeRule message.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_deleteedgerule
func (s *PullZoneService) DeleteEdgeRule(ctx context.Context, pullZoneID int64, edgeRuleGUID string) error {
	path := fmt.Sprintf("pullzone/%d/edgerules/%s", pullZoneID, edgeRuleGUID)
	return resourceDelete(ctx, s.client, path, nil)
}
