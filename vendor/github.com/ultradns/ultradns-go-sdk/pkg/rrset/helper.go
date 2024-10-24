package rrset

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/dirpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/rdpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sbpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sfpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/slbpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/tcpool"
)

func (r RRSetKey) RecordURI() string {
	r.Owner = url.PathEscape(r.Owner)
	r.Zone = url.PathEscape(r.Zone)

	if r.RecordType == "" {
		r.RecordType = "ANY"
	}

	return fmt.Sprintf("zones/%s/rrsets/%s/%s", r.Zone, r.RecordType, r.Owner)
}

func (r RRSetKey) ProbeURI() string {
	return fmt.Sprintf("%s/probes/%s", r.RecordURI(), r.ID)
}

func (r RRSetKey) ProbeListURI(query string) string {
	return fmt.Sprintf("%s/probes?q=%s", r.RecordURI(), query)
}

func (r RRSetKey) RecordID() string {
	r.Owner = helper.GetOwnerFQDN(r.Owner, r.Zone)
	r.Zone = helper.GetZoneFQDN(r.Zone)
	r.RecordType = helper.GetRecordTypeFullString(r.RecordType)

	if r.RecordType == "" {
		r.RecordType = "ANY"
	}

	return fmt.Sprintf("%s:%s:%s", r.Owner, r.Zone, r.RecordType)
}

// PID stands for Probe Id
func (r RRSetKey) PID() string {
	return fmt.Sprintf("%s:%s", r.RecordID(), r.ID)
}

func (r *RRSet) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}

	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	if val, ok := m["ownerName"].(string); ok {
		r.OwnerName = val
	}

	if val, ok := m["rrtype"].(string); ok {
		r.RRType = val
	}

	if val, ok := m["ttl"].(float64); ok {
		r.TTL = int(val)
	}

	if val, ok := m["rdata"].([]interface{}); ok {
		r.RData = make([]string, len(val))

		for i, v := range val {
			r.RData[i] = v.(string)
		}
	}

	if val, ok := m["profile"].(map[string]interface{}); ok {
		return r.setProfile(val)
	}

	return nil
}

func (r *RRSet) setProfile(m map[string]interface{}) error {

	profileJson, err := json.Marshal(m)

	if err != nil {
		return err
	}

	var context string

	if val, ok := m["@context"].(string); ok {
		context = val
	}

	profile := getPoolProfile(context)

	err = json.Unmarshal(profileJson, profile)

	if err != nil {
		return err
	}

	r.Profile = profile
	return nil
}

func getPoolProfile(context string) RawProfile {
	switch context {
	case rdpool.Schema:
		return &rdpool.Profile{}
	case sfpool.Schema:
		return &sfpool.Profile{}
	case slbpool.Schema:
		return &slbpool.Profile{}
	case sbpool.Schema:
		return &sbpool.Profile{}
	case tcpool.Schema:
		return &tcpool.Profile{}
	case dirpool.Schema:
		return &dirpool.Profile{}
	}

	return nil
}
