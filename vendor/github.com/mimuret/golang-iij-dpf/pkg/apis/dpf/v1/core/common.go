package core

import (
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/schema"
)

const groupName = "core.api.dns-platform.jp/v1"

func register(items ...apis.Spec) {
	schema.NewRegister(groupName).Add(items...)
}

type AttributeMeta struct{}

func (s *AttributeMeta) GetGroup() string { return groupName }
