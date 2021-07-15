package desec

import (
	"fmt"
	"net/http"
	"time"
)

// ApexZone apex zone name.
// https://desec.readthedocs.io/en/latest/dns/rrsets.html#accessing-the-zone-apex
const ApexZone = "@"

// RRSet DNS Record Set.
type RRSet struct {
	Name    string     `json:"name,omitempty"`
	Domain  string     `json:"domain,omitempty"`
	SubName string     `json:"subname,omitempty"`
	Type    string     `json:"type,omitempty"`
	Records []string   `json:"records"`
	TTL     int        `json:"ttl,omitempty"`
	Created *time.Time `json:"created,omitempty"`
}

// RRSetFilter a RRsets filter.
type RRSetFilter struct {
	Type    string
	SubName string
}

// RecordsService handles communication with the records related methods of the deSEC API.
//
// https://desec.readthedocs.io/en/latest/dns/rrsets.html
type RecordsService struct {
	client *Client
}

/*
	Domains
*/

// GetAll retrieving all RRsets in a zone.
// https://desec.readthedocs.io/en/latest/dns/rrsets.html#retrieving-all-rrsets-in-a-zone
func (s *RecordsService) GetAll(domainName string, filter *RRSetFilter) ([]RRSet, error) {
	endpoint, err := s.client.createEndpoint("domains", domainName, "rrsets")
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	if filter != nil {
		query := endpoint.Query()
		query.Set("type", filter.Type)
		query.Set("subname", filter.SubName)
		endpoint.RawQuery = query.Encode()
	}

	req, err := s.client.newRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, handleError(resp)
	}

	var rrSets []RRSet
	err = handleResponse(resp, &rrSets)
	if err != nil {
		return nil, err
	}

	return rrSets, nil
}

// Create creates a new RRSet.
// https://desec.readthedocs.io/en/latest/dns/rrsets.html#creating-a-tlsa-rrset
func (s *RecordsService) Create(rrSet RRSet) (*RRSet, error) {
	endpoint, err := s.client.createEndpoint("domains", rrSet.Domain, "rrsets")
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPost, endpoint, rrSet)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusCreated {
		return nil, handleError(resp)
	}

	var newRRSet RRSet
	err = handleResponse(resp, &newRRSet)
	if err != nil {
		return nil, err
	}

	return &newRRSet, nil
}

/*
	Domains + subname + type
*/

// Get gets a RRSet.
// https://desec.readthedocs.io/en/latest/dns/rrsets.html#retrieving-a-specific-rrset
func (s *RecordsService) Get(domainName, subName, recordType string) (*RRSet, error) {
	if subName == "" {
		subName = ApexZone
	}

	endpoint, err := s.client.createEndpoint("domains", domainName, "rrsets", subName, recordType)
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, handleError(resp)
	}

	var rrSet RRSet
	err = handleResponse(resp, &rrSet)
	if err != nil {
		return nil, err
	}

	return &rrSet, nil
}

// Update updates RRSet (PATCH).
// https://desec.readthedocs.io/en/latest/dns/rrsets.html#modifying-an-rrset
func (s *RecordsService) Update(domainName, subName, recordType string, rrSet RRSet) (*RRSet, error) {
	if subName == "" {
		subName = ApexZone
	}

	endpoint, err := s.client.createEndpoint("domains", domainName, "rrsets", subName, recordType)
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPatch, endpoint, rrSet)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	// when a RRSet is deleted (empty records)
	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, handleError(resp)
	}

	var updatedRRSet RRSet
	err = handleResponse(resp, &updatedRRSet)
	if err != nil {
		return nil, err
	}

	return &updatedRRSet, nil
}

// Replace replaces a RRSet (PUT).
// https://desec.readthedocs.io/en/latest/dns/rrsets.html#modifying-an-rrset
func (s *RecordsService) Replace(domainName, subName, recordType string, rrSet RRSet) (*RRSet, error) {
	if subName == "" {
		subName = ApexZone
	}

	endpoint, err := s.client.createEndpoint("domains", domainName, "rrsets", subName, recordType)
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPut, endpoint, rrSet)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	// when a RRSet is deleted (empty records)
	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, handleError(resp)
	}

	var updatedRRSet RRSet
	err = handleResponse(resp, &updatedRRSet)
	if err != nil {
		return nil, err
	}

	return &updatedRRSet, nil
}

// Delete deletes a RRset.
// https://desec.readthedocs.io/en/latest/dns/rrsets.html#deleting-an-rrset
func (s *RecordsService) Delete(domainName, subName, recordType string) error {
	if subName == "" {
		subName = ApexZone
	}

	endpoint, err := s.client.createEndpoint("domains", domainName, "rrsets", subName, recordType)
	if err != nil {
		return fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusNoContent {
		return handleError(resp)
	}

	return nil
}
