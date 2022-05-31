package zones

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

type RecordDiff struct {
	New *Record `read:"new"`
	Old *Record `read:"old"`
}

var _ CountableListSpec = &RecordDiffList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type RecordDiffList struct {
	AttributeMeta
	api.Count
	Items []RecordDiff `read:"items"`
}

func (c *RecordDiffList) GetName() string         { return "records/diffs" }
func (c *RecordDiffList) GetItems() interface{}   { return &c.Items }
func (c *RecordDiffList) Len() int                { return len(c.Items) }
func (c *RecordDiffList) Index(i int) interface{} { return c.Items[i] }
func (c *RecordDiffList) GetMaxLimit() int32      { return 10000 }
func (c *RecordDiffList) ClearItems()             { c.Items = []RecordDiff{} }
func (c *RecordDiffList) AddItem(v interface{}) bool {
	if a, ok := v.(RecordDiff); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *RecordDiffList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *RecordDiffList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneID)
}

func (c *RecordDiffList) Init() {
	for i := range c.Items {
		if c.Items[i].New != nil {
			c.Items[i].New.AttributeMeta = c.AttributeMeta
		}
		if c.Items[i].Old != nil {
			c.Items[i].Old.AttributeMeta = c.AttributeMeta
		}
	}
}

var _ api.SearchParams = &RecordListSearchKeywords{}

// +k8s:deepcopy-gen=false

type RecordListSearchKeywords struct {
	api.CommonSearchParams
	FullText    api.KeywordsString `url:"_keywords_full_text[],omitempty"`
	Name        api.KeywordsString `url:"_keywords_name[],omitempty"`
	TTL         []int32            `url:"_keywords_ttl[],omitempty"`
	RRType      KeywordsType       `url:"_keywords_rrtype[],omitempty"`
	RData       api.KeywordsString `url:"_keywords_rdata[],omitempty"`
	Description api.KeywordsString `url:"_keywords_description[],omitempty"`
	Operator    api.KeywordsString `url:"_keywords_operator[],omitempty"`
}

func (s *RecordListSearchKeywords) GetValues() (url.Values, error) { return query.Values(s) }

func init() {
	register(&RecordDiffList{})
}
