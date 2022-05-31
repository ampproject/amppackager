package core

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type Delegation struct {
	AttributeMeta
	ID                    string     `read:"id"`
	ServiceCode           string     `read:"service_code"`
	Name                  string     `read:"name"`
	Network               string     `read:"network"`
	Description           string     `read:"description"`
	DelegationRequestedAt types.Time `read:"delegation_requested_at"`
}

var _ apis.CountableListSpec = &DelegationList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type DelegationList struct {
	AttributeMeta
	api.Count
	Items []Delegation `read:"items"`
}

func (c *DelegationList) GetName() string         { return "delegations" }
func (c *DelegationList) GetItems() interface{}   { return &c.Items }
func (c *DelegationList) Len() int                { return len(c.Items) }
func (c *DelegationList) Index(i int) interface{} { return c.Items[i] }
func (c *DelegationList) GetMaxLimit() int32      { return 10000 }
func (c *DelegationList) ClearItems()             { c.Items = []Delegation{} }
func (c *DelegationList) AddItem(v interface{}) bool {
	if a, ok := v.(Delegation); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *DelegationList) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionList:
		return action.ToMethod(), "/delegations"
	case api.ActionCount:
		return action.ToMethod(), "/delegations/count"
	}
	return "", ""
}

func (c *DelegationList) SetPathParams(args ...interface{}) error {
	return nil
}

func (c *DelegationList) Init() {}

var _ api.SearchParams = &DelegationListSearchKeywords{}

// +k8s:deepcopy-gen=false
type DelegationListSearchKeywords struct {
	api.CommonSearchParams
	FullText    api.KeywordsString   `url:"_keywords_full_text[],omitempty"`
	ServiceCode api.KeywordsString   `url:"_keywords_service_code[],omitempty"`
	Name        api.KeywordsString   `url:"_keywords_name[],omitempty"`
	Network     api.KeywordsString   `url:"_keywords_network[],omitempty"`
	Favorite    api.KeywordsFavorite `url:"_keywords_favorite[],omitempty"`
	Description api.KeywordsString   `url:"_keywords_description[],omitempty"`
	Requested   api.KeywordsBoolean  `url:"_keywords_requested[],omitempty"`
}

func (s *DelegationListSearchKeywords) GetValues() (url.Values, error) { return query.Values(s) }

func init() {
	register(&DelegationList{})
}
