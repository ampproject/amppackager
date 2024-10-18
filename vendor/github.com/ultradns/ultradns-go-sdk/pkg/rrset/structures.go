package rrset

import (
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

type RRSet struct {
	OwnerName string     `json:"ownerName,omitempty"`
	RRType    string     `json:"rrtype,omitempty"`
	TTL       int        `json:"ttl,omitempty"`
	RData     []string   `json:"rdata,omitempty"`
	Profile   RawProfile `json:"profile,omitempty"`
}

type RRSetKey struct {
	ID         string
	Owner      string
	Zone       string
	RecordType string
	PType      string
}

type RawProfile interface {
	SetContext()
	GetContext() string
}

type ResponseList struct {
	ZoneName   string             `json:"zoneName,omitempty"`
	QueryInfo  *helper.QueryInfo  `json:"queryInfo,omitempty"`
	ResultInfo *helper.ResultInfo `json:"resultInfo,omitempty"`
	RRSets     []*RRSet           `json:"rrSets,omitempty"`
}
