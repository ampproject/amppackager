package nodion

import (
	"fmt"
	"strings"
	"time"
)

// Record types.
const (
	TypeA     = "a"
	TypeAAAA  = "aaaa"
	TypeNS    = "ns"
	TypeALIAS = "alias"
	TypeCNAME = "cname"
	TypeMX    = "mx"
	TypeTXT   = "txt"
	TypePTR   = "ptr"
	TypeSRV   = "srv"
)

// ZonesResponse represents the response of the GetZones API endpoint.
type ZonesResponse struct {
	Zones []Zone `json:"dns_zones"`
}

// ZoneResponse represents the response of the CreateZone API endpoint.
type ZoneResponse struct {
	Zone Zone `json:"dns_zone"`
}

// RecordsResponse represents the response of the GetRecords API endpoint.
type RecordsResponse struct {
	Records []Record `json:"records"`
}

// RecordResponse represents the response of the CreateRecord API endpoint.
type RecordResponse struct {
	Record Record `json:"record"`
}

// DeleteResponse represents the response of the API endpoints related to deletion of zone or record.
type DeleteResponse struct {
	Deleted bool `json:"deleted"`
}

// Zone contains all the information related to a zone.
type Zone struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Records   []Record  `json:"records,omitempty"`
}

// Record contains all the information related to a DNS record.
type Record struct {
	ID         string    `json:"id,omitempty"`
	RecordType string    `json:"record_type,omitempty"` // a, aaaa, ns, alias, cname, mx, txt, ptr, srv. Case-sensitive must be in lowercase.
	Name       string    `json:"name,omitempty"`
	Content    string    `json:"content,omitempty"`
	TTL        int       `json:"ttl,omitempty"` // a number between 60 and 86400.
	ZoneID     string    `json:"zone_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

// ZonesFilter is filter criteria for zones.
type ZonesFilter struct {
	Name string `url:"name"` // must be the exact name and no FQDN
}

// RecordsFilter is filter criteria for records.
type RecordsFilter struct {
	Name       string `url:"name"`
	RecordType string `url:"record_type"`
	Content    string `url:"content"`
}

// APIError is the error returned by the server.
type APIError struct {
	StatusCode int      `json:"status"`
	Message    string   `json:"error"`
	Errors     []string `json:"errors"`
}

func (a *APIError) Error() string {
	if a.Message != "" {
		return fmt.Sprintf("status code %d: %s", a.StatusCode, a.Message)
	}

	return fmt.Sprintf("status code %d: %s", a.StatusCode, strings.Join(a.Errors, ", "))
}
