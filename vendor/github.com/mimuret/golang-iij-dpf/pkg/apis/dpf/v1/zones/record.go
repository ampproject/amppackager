package zones

import (
	"fmt"
	"strings"

	"github.com/miekg/dns"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

type RecordState int

const TypeANAMECode uint16 = 65280

const (
	RecordStateApplied      RecordState = 0
	RecordStateToBeAdded    RecordState = 1
	RecordStateToBeDeleted  RecordState = 2
	RecordStateToBeUpdate   RecordState = 3
	RecordStateBeforeUpdate RecordState = 5
)

func (c RecordState) String() string {
	recordStateToString := map[RecordState]string{
		RecordStateApplied:      "Applied",
		RecordStateToBeAdded:    "ToBeAdded",
		RecordStateToBeDeleted:  "ToBeDeleted",
		RecordStateToBeUpdate:   "ToBeUpdate",
		RecordStateBeforeUpdate: "BeforeUpdate",
	}
	return recordStateToString[c]
}

type Type string

const (
	TypeSOA   Type = "SOA"
	TypeA     Type = "A"
	TypeAAAA  Type = "AAAA"
	TypeCAA   Type = "CAA"
	TypeCNAME Type = "CNAME"
	TypeDS    Type = "DS"
	TypeNS    Type = "NS"
	TypeMX    Type = "MX"
	TypeNAPTR Type = "NAPTR"
	TypeSRV   Type = "SRV"
	TypeTXT   Type = "TXT"
	TypeTLSA  Type = "TLSA"
	TypePTR   Type = "PTR"
	TypeSVCB  Type = "SVCB"
	TypeHTTPS Type = "HTTPS"

	TypeANAME Type = "ANAME"
)

func (c Type) String() string {
	return string(c)
}

func (c Type) Uint16() uint16 {
	if c == TypeANAME {
		return TypeANAMECode
	}
	return dns.StringToType[string(c)]
}

func Uint16ToType(t uint16) Type {
	if t == TypeANAMECode {
		return TypeANAME
	}
	return Type(dns.TypeToString[t])
}

// +k8s:deepcopy-gen=false
type KeywordsType []Type

type RecordRDATA struct {
	Value string `read:"value" create:"value" update:"value"`
}

func (c *RecordRDATA) String() string {
	return c.Value
}

type RecordRDATASlice []RecordRDATA

func (c RecordRDATASlice) String() string {
	var res []string
	for _, value := range c {
		res = append(res, value.String())
	}
	return strings.Join(res, ",")
}

var _ Spec = &Record{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type Record struct {
	AttributeMeta

	ID          string                      `read:"id"`
	Name        string                      `read:"name" create:"name"`
	TTL         types.NullablePositiveInt32 `read:"ttl"  create:"ttl" update:"ttl"`
	RRType      Type                        `read:"rrtype"  create:"rrtype"`
	RData       RecordRDATASlice            `read:"rdata"  create:"rdata" update:"rdata"`
	State       RecordState                 `read:"state"`
	Description string                      `read:"description"  create:"description" update:"description"`
	Operator    string                      `read:"operator"`
}

func (c *Record) GetName() string { return "records" }
func (c *Record) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/%s", c.GetZoneID(), c.GetName())
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/%s/%s", c.GetZoneID(), c.GetName(), c.ID)
	case api.ActionCancel:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/%s/%s/changes", c.GetZoneID(), c.GetName(), c.ID)
	}
	return "", ""
}

func (c *Record) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneID, &c.ID)
}

var _ CountableListSpec = &RecordList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type RecordList struct {
	AttributeMeta
	api.Count
	Items []Record `read:"items"`
}

func (c *RecordList) GetName() string         { return "records" }
func (c *RecordList) GetItems() interface{}   { return &c.Items }
func (c *RecordList) Len() int                { return len(c.Items) }
func (c *RecordList) Index(i int) interface{} { return c.Items[i] }
func (c *RecordList) GetMaxLimit() int32      { return 10000 }
func (c *RecordList) ClearItems()             { c.Items = []Record{} }
func (c *RecordList) AddItem(v interface{}) bool {
	if a, ok := v.(Record); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *RecordList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *RecordList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneID)
}

func (c *RecordList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}

var _ CountableListSpec = &CurrentRecordList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type CurrentRecordList struct {
	AttributeMeta
	api.Count
	Items []Record `read:"items"`
}

func (c *CurrentRecordList) GetName() string         { return "records/currents" }
func (c *CurrentRecordList) Len() int                { return len(c.Items) }
func (c *CurrentRecordList) GetItems() interface{}   { return &c.Items }
func (c *CurrentRecordList) Index(i int) interface{} { return c.Items[i] }
func (c *CurrentRecordList) GetMaxLimit() int32      { return 10000 }
func (c *CurrentRecordList) ClearItems()             { c.Items = []Record{} }
func (c *CurrentRecordList) AddItem(v interface{}) bool {
	if a, ok := v.(Record); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *CurrentRecordList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *CurrentRecordList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}

func (c *CurrentRecordList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneID)
}

func init() {
	register(&Record{}, &RecordList{})
	register(&CurrentRecordList{})
}
