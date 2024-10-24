package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Types of domain name.
const (
	A     RecordType = "A"
	AAAA  RecordType = "AAAA"
	ALIAS RecordType = "ALIAS"
	CAA   RecordType = "CAA"
	CNAME RecordType = "CNAME"
	MX    RecordType = "MX"
	NS    RecordType = "NS"
	SOA   RecordType = "SOA"
	SRV   RecordType = "SRV"
	SSHFP RecordType = "SSHFP"
	TXT   RecordType = "TXT"
)

type (
	// RecordType contains record types supported by target DNS api.
	RecordType string

	// RRSet is list of records grouped by their name and type.
	RRSet struct {
		ID        string       `json:"id"`
		ZoneID    string       `json:"zone_id"`
		Name      string       `json:"name"`
		TTL       int          `json:"ttl"`
		Type      RecordType   `json:"type"`
		Comment   string       `json:"comment"`
		ManagedBy string       `json:"managed_by"`
		Records   []RecordItem `json:"records"`
	}

	// RecordItem represents single record from RRSet.
	RecordItem struct {
		Content  string `json:"content"`
		Disabled bool   `json:"disabled"`
	}

	rrsetCreationForm struct {
		Name      string       `json:"name"`
		TTL       int          `json:"ttl"`
		Type      RecordType   `json:"type"`
		Records   []RecordItem `json:"records"`
		Comment   string       `json:"comment,omitempty"`
		ManagedBy string       `json:"managed_by,omitempty"`
	}

	rrsetUpdateForm struct {
		TTL       int          `json:"ttl"`
		Records   []RecordItem `json:"records"`
		Comment   string       `json:"comment,omitempty"`
		ManagedBy string       `json:"managed_by,omitempty"`
	}
)

func (s *RRSet) CreationForm() (io.Reader, error) {
	form := rrsetCreationForm{
		Name:      s.Name,
		TTL:       s.TTL,
		Type:      s.Type,
		Records:   s.Records,
		Comment:   s.Comment,
		ManagedBy: s.ManagedBy,
	}
	body, err := json.Marshal(form)

	return bytes.NewReader(body), err
}

func (s *RRSet) UpdateForm() (io.Reader, error) {
	form := rrsetUpdateForm{
		TTL:       s.TTL,
		Records:   s.Records,
		Comment:   s.Comment,
		ManagedBy: s.ManagedBy,
	}
	body, err := json.Marshal(form)

	return bytes.NewReader(body), err
}

// CreateRRSet request to create a new rrset for the zone.
func (c *Client) CreateRRSet(ctx context.Context, zoneID string, rrset Creatable) (*RRSet, error) {
	form, err := rrset.CreationForm()
	if err != nil {
		return nil, fmt.Errorf("rrset creation form: %w", err)
	}
	r, e := c.prepareRequest(
		ctx, http.MethodPost, fmt.Sprintf(rrsetPath, zoneID), form, nil, nil,
	)

	return processRequest[RRSet](c.httpClient, r, e)
}

// DeleteRRSet request to delete the rrset from zone by zoneID and rrsetID.
func (c *Client) DeleteRRSet(ctx context.Context, zoneID, rrsetID string) error {
	r, e := c.prepareRequest(
		ctx, http.MethodDelete, fmt.Sprintf(singleRRSetPath, zoneID, rrsetID), nil, nil, nil,
	)
	_, err := processRequest[RRSet](c.httpClient, r, e)

	return err
}

// GetRRSet returns a single rrset from zone by zoneID and rrsetID.
func (c *Client) GetRRSet(ctx context.Context, zoneID, rrsetID string) (*RRSet, error) {
	r, e := c.prepareRequest(
		ctx, http.MethodGet, fmt.Sprintf(singleRRSetPath, zoneID, rrsetID), nil, nil, nil,
	)

	return processRequest[RRSet](c.httpClient, r, e)
}

// ListRRSets returns a list of rrsets by zoneID and options.
func (c *Client) ListRRSets(ctx context.Context, zoneID string, options *map[string]string) (Listable[RRSet], error) {
	r, e := c.prepareRequest(
		ctx, http.MethodGet, fmt.Sprintf(rrsetPath, zoneID), nil, options, nil,
	)

	return processRequest[List[RRSet]](c.httpClient, r, e)
}

// UpdateRRSet request to update the rrset for zone by zoneID and rrsetID.
func (c *Client) UpdateRRSet(ctx context.Context, zoneID, rrsetID string, rrset Updatable) error {
	form, err := rrset.UpdateForm()
	if err != nil {
		return fmt.Errorf("rrset update form: %w", err)
	}
	r, e := c.prepareRequest(
		ctx, http.MethodPatch, fmt.Sprintf(singleRRSetPath, zoneID, rrsetID), form, nil, nil,
	)
	_, err = processRequest[RRSet](c.httpClient, r, e)

	return err
}
