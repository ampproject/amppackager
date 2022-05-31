package zones

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ Spec = &HistoryText{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type HistoryText struct {
	AttributeMeta
	History
	Text string `read:"text"`
}

func (c *HistoryText) GetName() string { return fmt.Sprintf("zone_histories/%d/text", c.ID) }
func (c *HistoryText) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionRead {
		return action.ToMethod(), fmt.Sprintf("/zones/%s/zone_histories/%d/text", c.GetZoneID(), c.ID)
	}
	return "", ""
}

func (c *HistoryText) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneID, &c.ID)
}

func init() {
	register(&HistoryText{})
}
