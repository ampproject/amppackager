package zones

import (
	"fmt"
	"net/http"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ Spec = &DnssecKskRollover{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type DnssecKskRollover struct {
	AttributeMeta
}

func (c *DnssecKskRollover) GetName() string { return "dnssec/ksk_rollover" }
func (c *DnssecKskRollover) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionApply {
		return http.MethodPatch, fmt.Sprintf("/zones/%s/dnssec/ksk_rollover", c.GetZoneID())
	}
	return "", ""
}

func (c *DnssecKskRollover) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneID)
}

func init() {
	register(&DnssecKskRollover{})
}
