package core

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

type Plan int

const (
	PlanBasic   Plan = 1
	PlanPremium Plan = 2
)

func (c Plan) String() string {
	planToString := map[Plan]string{
		PlanBasic:   "basic",
		PlanPremium: "premium",
	}

	return planToString[c]
}

// +k8s:deepcopy-gen=false
type KeywordsPlan []Plan

func (c KeywordsPlan) EncodeValues(key string, v *url.Values) error {
	for _, plan := range c {
		v.Add(key, strconv.Itoa(int(plan)))
	}
	return nil
}

var _ apis.Spec = &Contract{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type Contract struct {
	AttributeMeta
	ID          string         `read:"id"`
	ServiceCode string         `read:"service_code"`
	State       types.State    `read:"state"`
	Favorite    types.Favorite `read:"favorite" update:"favorite"`
	Plan        Plan           `read:"plan"`
	Description string         `read:"description" update:"description"`
}

func (c *Contract) GetName() string { return "contracts" }
func (c *Contract) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionRead, api.ActionUpdate:
		return action.ToMethod(), fmt.Sprintf("/contracts/%s", c.ID)
	}
	return "", ""
}

func (c *Contract) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ID)
}

var _ apis.CountableListSpec = &ContractList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type ContractList struct {
	AttributeMeta
	api.Count
	Items []Contract `read:"items"`
}

func (c *ContractList) GetName() string         { return "contracts" }
func (c *ContractList) GetItems() interface{}   { return &c.Items }
func (c *ContractList) Len() int                { return len(c.Items) }
func (c *ContractList) Index(i int) interface{} { return c.Items[i] }
func (c *ContractList) GetMaxLimit() int32      { return 10000 }
func (c *ContractList) ClearItems()             { c.Items = []Contract{} }
func (c *ContractList) AddItem(v interface{}) bool {
	if a, ok := v.(Contract); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *ContractList) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionList:
		return action.ToMethod(), "/contracts"
	case api.ActionCount:
		return action.ToMethod(), "/contracts/count"
	}
	return "", ""
}

func (c *ContractList) SetPathParams(args ...interface{}) error {
	return nil
}

func (c *ContractList) Init() {}

var _ api.SearchParams = &ContractListSearchKeywords{}

// +k8s:deepcopy-gen=false
type ContractListSearchKeywords struct {
	api.CommonSearchParams
	FullText    api.KeywordsString   `url:"_keywords_full_text[],omitempty"`
	ServiceCode api.KeywordsString   `url:"_keywords_service_code[],omitempty"`
	Plan        KeywordsPlan         `url:"_keywords_plan[],omitempty"`
	State       api.KeywordsState    `url:"_keywords_state[],omitempty"`
	Favorite    api.KeywordsFavorite `url:"_keywords_favorite[],omitempty"`
	Description api.KeywordsString   `url:"_keywords_description[],omitempty"`
}

func (s *ContractListSearchKeywords) GetValues() (url.Values, error) { return query.Values(s) }

func init() {
	register(&Contract{}, &ContractList{})
}
