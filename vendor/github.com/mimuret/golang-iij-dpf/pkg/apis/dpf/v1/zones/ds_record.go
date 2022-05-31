package zones

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ ListSpec = &DsRecordList{}

type DsRecord struct {
	RRSet     string     `read:"rrset"`
	TransitAt types.Time `read:"transited_at"`
}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type DsRecordList struct {
	AttributeMeta
	Items []DsRecord `read:"items"`
}

func (c *DsRecordList) GetName() string         { return "ds_records" }
func (c *DsRecordList) GetItems() interface{}   { return &c.Items }
func (c *DsRecordList) Len() int                { return len(c.Items) }
func (c *DsRecordList) Index(i int) interface{} { return c.Items[i] }

func (c *DsRecordList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *DsRecordList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneID)
}

func (c *DsRecordList) Init() {}

func init() {
	register(&DsRecordList{})
}
