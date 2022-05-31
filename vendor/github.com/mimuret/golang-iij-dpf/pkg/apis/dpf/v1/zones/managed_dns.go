package zones

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ ListSpec = &ManagedDnsList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type ManagedDnsList struct {
	AttributeMeta
	Items []string `read:"items"`
}

func (c *ManagedDnsList) GetName() string         { return "managed_dns_servers" }
func (c *ManagedDnsList) GetItems() interface{}   { return &c.Items }
func (c *ManagedDnsList) Len() int                { return len(c.Items) }
func (c *ManagedDnsList) Index(i int) interface{} { return c.Items[i] }

func (c *ManagedDnsList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *ManagedDnsList) Init() {}
func (c *ManagedDnsList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneID)
}

func init() {
	register(&ManagedDnsList{})
}
