package zones

import (
	"fmt"
	"net/http"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/schema"
)

const groupName = "zones.api.dns-platform.jp/v1"

func register(items ...apis.Spec) {
	schema.NewRegister(groupName).Add(items...)
}

type Spec interface {
	apis.Spec
	SetZoneID(string)
	GetZoneID() string
}

type ChildSpec interface {
	Spec
	GetID() int64
	SetID(int64)
}

type ListSpec interface {
	api.ListSpec
	Spec
}

type CountableListSpec interface {
	api.CountableListSpec
	Spec
}

type AttributeMeta struct {
	ZoneID string `read:"-"`
}

// for ctl
func (s *AttributeMeta) GetGroup() string    { return groupName }
func (s *AttributeMeta) SetZoneID(id string) { s.ZoneID = id }
func (s *AttributeMeta) GetZoneID() string   { return s.ZoneID }

func GetPathMethodForListSpec(action api.Action, s ListSpec) (string, string) {
	switch action {
	case api.ActionList:
		return http.MethodGet, fmt.Sprintf("/zones/%s/%s", s.GetZoneID(), s.GetName())
	case api.ActionCount:
		if _, ok := s.(api.CountableListSpec); ok {
			return http.MethodGet, fmt.Sprintf("/zones/%s/%s/count", s.GetZoneID(), s.GetName())
		}
	}
	return "", ""
}

func GetReadPathMethodForSpec(action api.Action, s Spec) (string, string) {
	if action == api.ActionRead {
		return http.MethodGet, fmt.Sprintf("/zones/%s/%s", s.GetZoneID(), s.GetName())
	}
	return "", ""
}
