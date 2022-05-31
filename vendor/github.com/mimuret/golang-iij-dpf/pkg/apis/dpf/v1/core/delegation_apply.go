package core

import (
	"net/http"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ apis.Spec = &DelegationApply{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type DelegationApply struct {
	AttributeMeta
	ZoneIDs []string `apply:"zone_ids"`
}

func (c *DelegationApply) GetName() string { return "delegations" }
func (c *DelegationApply) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionApply {
		return http.MethodPost, "/delegations"
	}
	return "", ""
}

func (c *DelegationApply) SetPathParams(args ...interface{}) error {
	return nil
}

func init() {
	register(&DelegationApply{})
}
