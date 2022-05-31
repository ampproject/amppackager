package zones

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
)

var _ CountableListSpec = &LogList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type LogList struct {
	AttributeMeta
	api.Count
	Items []core.Log `read:"items"`
}

func (c *LogList) GetName() string         { return "logs" }
func (c *LogList) GetItems() interface{}   { return &c.Items }
func (c *LogList) Len() int                { return len(c.Items) }
func (c *LogList) Index(i int) interface{} { return c.Items[i] }
func (c *LogList) GetMaxLimit() int32      { return 100 }
func (c *LogList) ClearItems()             { c.Items = []core.Log{} }
func (c *LogList) AddItem(v interface{}) bool {
	if a, ok := v.(core.Log); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *LogList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *LogList) Init() {}
func (c *LogList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneID)
}

func init() {
	register(&LogList{})
}
