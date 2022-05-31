package zones

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ api.Spec = &ZoneApply{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type ZoneApply struct {
	AttributeMeta
	Description string `apply:"description"`
}

func (c *ZoneApply) GetName() string { return "apply" }
func (c *ZoneApply) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionApply:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/changes", c.AttributeMeta.ZoneID)
	case api.ActionCancel:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/changes", c.AttributeMeta.ZoneID)
	}
	return "", ""
}

func (c *ZoneApply) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.AttributeMeta.ZoneID)
}

func init() {
	register(&ZoneApply{})
}
