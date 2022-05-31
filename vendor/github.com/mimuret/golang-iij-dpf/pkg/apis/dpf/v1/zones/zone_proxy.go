package zones

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ Spec = &ZoneProxy{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type ZoneProxy struct {
	AttributeMeta
	Enabled types.Boolean `read:"enabled" update:"enabled"`
}

func (c *ZoneProxy) GetName() string { return "zone_proxy" }
func (c *ZoneProxy) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionRead, api.ActionUpdate:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/zone_proxy", c.GetZoneID())
	}
	return "", ""
}

func (c *ZoneProxy) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneID)
}

func init() {
	register(&ZoneProxy{})
}
