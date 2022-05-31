package core

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ api.Spec = &Zone{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type Zone struct {
	AttributeMeta
	ID               string         `read:"id"`
	CommonConfigID   int64          `read:"common_config_id"`
	ServiceCode      string         `read:"service_code"`
	State            types.State    `read:"state"`
	Favorite         types.Favorite `read:"favorite" update:"favorite"`
	Name             string         `read:"name"`
	Network          string         `read:"network"`
	Description      string         `read:"description" update:"description"`
	ZoneProxyEnabled types.Boolean  `read:"zone_proxy_enabled"`
}

func (c *Zone) GetName() string { return "zones" }
func (c *Zone) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionRead, api.ActionUpdate:
		return action.ToMethod(), fmt.Sprintf("/zones/%s", c.ID)
	}
	return "", ""
}

func (c *Zone) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ID)
}

var _ apis.CountableListSpec = &ZoneList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type ZoneList struct {
	AttributeMeta
	api.Count
	Items []Zone `read:"items"`
}

func (c *ZoneList) GetName() string         { return "zones" }
func (c *ZoneList) GetItems() interface{}   { return &c.Items }
func (c *ZoneList) Len() int                { return len(c.Items) }
func (c *ZoneList) Index(i int) interface{} { return c.Items[i] }
func (c *ZoneList) GetMaxLimit() int32      { return 10000 }
func (c *ZoneList) ClearItems()             { c.Items = []Zone{} }
func (c *ZoneList) AddItem(v interface{}) bool {
	if a, ok := v.(Zone); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *ZoneList) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionList:
		return action.ToMethod(), "/zones"
	case api.ActionCount:
		return action.ToMethod(), "/zones/count"
	}
	return "", ""
}
func (c *ZoneList) Init() {}
func (c *ZoneList) SetPathParams(args ...interface{}) error {
	return nil
}

var _ api.SearchParams = &ZoneListSearchKeywords{}

// +k8s:deepcopy-gen=false
type ZoneListSearchKeywords struct {
	api.CommonSearchParams
	FullText         api.KeywordsString   `url:"_keywords_full_text[],omitempty"`
	ServiceCode      api.KeywordsString   `url:"_keywords_service_code[],omitempty"`
	Name             api.KeywordsString   `url:"_keywords_name[],omitempty"`
	Network          api.KeywordsString   `url:"_keywords_network[],omitempty"`
	State            api.KeywordsState    `url:"_keywords_state[],omitempty"`
	Favorite         api.KeywordsFavorite `url:"_keywords_favorite[],omitempty"`
	Description      api.KeywordsString   `url:"_keywords_description[],omitempty"`
	CommonConfigID   api.KeywordsID       `url:"_keywords_common_config_id[],omitempty"`
	ZoneProxyEnabled api.KeywordsBoolean  `url:"_keywords_zone_proxy_enabled[],omitempty"`
}

func (s *ZoneListSearchKeywords) GetValues() (url.Values, error) { return query.Values(s) }

func init() {
	register(&Zone{}, &ZoneList{})
}
