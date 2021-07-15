package desec

import (
	"fmt"
	"net/http"
	"time"
)

// Domain a domain representation.
type Domain struct {
	Name       string      `json:"name,omitempty"`
	MinimumTTL int         `json:"minimum_ttl,omitempty"`
	Keys       []DomainKey `json:"keys,omitempty"`
	Created    *time.Time  `json:"created,omitempty"`
	Published  *time.Time  `json:"published,omitempty"`
}

// DomainKey a domain key representation.
type DomainKey struct {
	DNSKey  string   `json:"dnskey,omitempty"`
	DS      []string `json:"ds,omitempty"`
	Flags   int      `json:"flags,omitempty"`
	KeyType string   `json:"keytype,omitempty"`
}

// DomainsService handles communication with the domain related methods of the deSEC API.
//
// https://desec.readthedocs.io/en/latest/dns/domains.html
type DomainsService struct {
	client *Client
}

// Create creating a domain.
// https://desec.readthedocs.io/en/latest/dns/domains.html#creating-a-domain
func (s *DomainsService) Create(domainName string) (*Domain, error) {
	endpoint, err := s.client.createEndpoint("domains")
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPost, endpoint, Domain{Name: domainName})
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

	var domain Domain
	err = handleResponse(resp, &domain)
	if err != nil {
		return nil, err
	}

	return &domain, nil
}

// GetAll listing domains.
// https://desec.readthedocs.io/en/latest/dns/domains.html#listing-domains
func (s *DomainsService) GetAll() ([]Domain, error) {
	endpoint, err := s.client.createEndpoint("domains")
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

	var domains []Domain
	err = handleResponse(resp, &domains)
	if err != nil {
		return nil, err
	}

	return domains, nil
}

// Get retrieving a specific domain.
// https://desec.readthedocs.io/en/latest/dns/domains.html#retrieving-a-specific-domain
func (s *DomainsService) Get(domainName string) (*Domain, error) {
	endpoint, err := s.client.createEndpoint("domains", domainName)
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

	var domains Domain
	err = handleResponse(resp, &domains)
	if err != nil {
		return nil, err
	}

	return &domains, nil
}

// Delete deleting a domain.
// https://desec.readthedocs.io/en/latest/dns/domains.html#deleting-a-domain
func (s *DomainsService) Delete(domainName string) error {
	endpoint, err := s.client.createEndpoint("domains", domainName)
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
