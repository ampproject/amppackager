package core

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ api.Spec = &LBDomain{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type LBDomain struct {
	AttributeMeta
	ID               string         `read:"id"`
	CommonConfigID   int64          `read:"common_config_id"`
	ServiceCode      string         `read:"service_code"`
	State            types.State    `read:"state"`
	Favorite         types.Favorite `read:"favorite" update:"favorite"`
	Name             string         `read:"name"`
	Description      string         `read:"description" update:"description"`
	RuleResourceName string         `read:"rule_resource_name" update:"rule_resource_name"`
}

func (c *LBDomain) GetName() string { return "lb_domains" }
func (c *LBDomain) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionRead, api.ActionUpdate:
		return action.ToMethod(), fmt.Sprintf("/lb_domains/%s", c.ID)
	}
	return "", ""
}

func (c *LBDomain) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ID)
}

var _ apis.CountableListSpec = &ZoneList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type LBDomainList struct {
	AttributeMeta
	api.Count
	Items []LBDomain `read:"items"`
}

func (c *LBDomainList) GetName() string         { return "lb_domains" }
func (c *LBDomainList) GetItems() interface{}   { return &c.Items }
func (c *LBDomainList) Len() int                { return len(c.Items) }
func (c *LBDomainList) Index(i int) interface{} { return c.Items[i] }
func (c *LBDomainList) GetMaxLimit() int32      { return 10000 }
func (c *LBDomainList) ClearItems()             { c.Items = []LBDomain{} }
func (c *LBDomainList) AddItem(v interface{}) bool {
	if a, ok := v.(LBDomain); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *LBDomainList) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionList:
		return action.ToMethod(), "/lb_domains"
	case api.ActionCount:
		return action.ToMethod(), "/lb_domains/count"
	}
	return "", ""
}
func (c *LBDomainList) Init() {}
func (c *LBDomainList) SetPathParams(args ...interface{}) error {
	return nil
}

var _ api.SearchParams = &LBDomainListSearchKeywords{}

// +k8s:deepcopy-gen=false
type LBDomainListSearchKeywords struct {
	api.CommonSearchParams
	FullText       api.KeywordsString   `url:"_keywords_full_text[],omitempty"`
	ServiceCode    api.KeywordsString   `url:"_keywords_service_code[],omitempty"`
	Name           api.KeywordsString   `url:"_keywords_name[],omitempty"`
	State          api.KeywordsState    `url:"_keywords_state[],omitempty"`
	Favorite       api.KeywordsFavorite `url:"_keywords_favorite[],omitempty"`
	Description    api.KeywordsString   `url:"_keywords_description[],omitempty"`
	CommonConfigID api.KeywordsID       `url:"_keywords_common_config_id[],omitempty"`
	Label          api.KeywordsLabels   `url:"_keywords_label[],omitempty"`
}

func (s *LBDomainListSearchKeywords) GetValues() (url.Values, error) { return query.Values(s) }

func init() {
	register(&LBDomain{}, &LBDomainList{})
}
