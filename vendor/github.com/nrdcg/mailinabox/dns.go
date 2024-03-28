package mailinabox

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Record Represents a DNS record.
type Record struct {
	Name        string `json:"qname,omitempty"`
	Type        string `json:"rtype,omitempty"`
	Value       string `json:"value,omitempty"`
	Explanation string `json:"explanation,omitempty"`
}

// Zone Represents a DNS zone.
type Zone struct {
	Zone    string
	Records []Record
}

// Zones a slice of Zones.
// Use a custom unmarshalling method.
type Zones []Zone

// UnmarshalJSON customs unmarshalling.
func (z *Zones) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var a []json.RawMessage
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	var all []Zone
	for _, message := range a {
		var b []json.RawMessage
		if err := json.Unmarshal(message, &b); err != nil {
			return err
		}

		zone := Zone{}

		if err := json.Unmarshal(b[0], &zone.Zone); err != nil {
			return err
		}

		if len(b) <= 1 {
			all = append(all, zone)
			continue
		}

		if err := json.Unmarshal(b[1], &zone.Records); err != nil {
			return err
		}

		all = append(all, zone)
	}

	*z = all

	return nil
}

// Nameserver Represents DNS nameservers.
type Nameserver struct {
	Hostnames []string `json:"hostnames"`
}

// DNSService DNS API.
// https://mailinabox.email/api-docs.html#tag/DNS
type DNSService service

// GetSecondaryNameserver Returns a list of nameserver hostnames.
// https://mailinabox.email/api-docs.html#operation/getDnsSecondaryNameserver
func (s *DNSService) GetSecondaryNameserver(ctx context.Context) ([]string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "dns", "secondary-nameserver")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var results Nameserver

	err = s.client.doJSON(req, &results)
	if err != nil {
		return nil, err
	}

	return results.Hostnames, nil
}

// AddSecondaryNameserver Adds one or more secondary nameservers.
// https://mailinabox.email/api-docs.html#operation/addDnsSecondaryNameserver
func (s *DNSService) AddSecondaryNameserver(ctx context.Context, hostnames []string) (string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "dns", "secondary-nameserver")

	data := url.Values{}
	data.Set("hostnames", strings.Join(hostnames, ","))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// GetZones Returns an array of all managed top-level domains.
// https://mailinabox.email/api-docs.html#operation/getDnsZones
func (s *DNSService) GetZones(ctx context.Context) ([]string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "dns", "zones")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var results []string

	err = s.client.doJSON(req, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetZoneFile Returns a DNS zone file for a hostname.
// https://mailinabox.email/api-docs.html#operation/getDnsZonefile
func (s *DNSService) GetZoneFile(ctx context.Context, zone string) (string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "dns", "zonefile", zone)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	var results string

	err = s.client.doJSON(req, &results)
	if err != nil {
		return "", err
	}

	return results, nil
}

// UpdateDNS Updates the DNS. Involves creating zone files and restarting `nsd`.
// https://mailinabox.email/api-docs.html#operation/updateDns
func (s *DNSService) UpdateDNS(ctx context.Context, force bool) (string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "dns", "update")

	data := url.Values{}
	data.Set("force", boolToIntStr(force))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// GetAllRecords Returns all custom DNS records.
// https://mailinabox.email/api-docs.html#operation/getDnsCustomRecords
func (s *DNSService) GetAllRecords(ctx context.Context) ([]Record, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "dns", "custom")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var results []Record

	err = s.client.doJSON(req, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetRecords Returns all custom records for the specified query name and type.
// https://mailinabox.email/api-docs.html#operation/getDnsCustomRecordsForQNameAndType
func (s *DNSService) GetRecords(ctx context.Context, name, rType string) ([]Record, error) {
	if name == "" || rType == "" {
		return nil, errors.New("qname and rtype are required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "dns", "custom", name, rType)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var results []Record

	err = s.client.doJSON(req, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// AddRecord Adds a custom DNS record for the specified query name and type.
// https://mailinabox.email/api-docs.html#operation/addDnsCustomRecord
func (s *DNSService) AddRecord(ctx context.Context, record Record) (string, error) {
	if record.Name == "" || record.Type == "" {
		return "", errors.New("qname and rtype are required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "dns", "custom", record.Name, record.Type)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(record.Value))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// UpdateRecord Updates an existing DNS custom record value for the specified qname and type.
// https://mailinabox.email/api-docs.html#operation/updateDnsCustomRecord
func (s *DNSService) UpdateRecord(ctx context.Context, record Record, value string) (string, error) {
	if record.Name == "" || record.Type == "" {
		return "", errors.New("qname and rtype are required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "dns", "custom", record.Name, record.Type)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, endpoint.String(), strings.NewReader(value))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// RemoveRecord Removes a DNS custom record for the specified domain, type & value.
// https://mailinabox.email/api-docs.html#operation/removeDnsCustomRecord
func (s *DNSService) RemoveRecord(ctx context.Context, record Record) (string, error) {
	if record.Name == "" || record.Type == "" {
		return "", errors.New("qname and rtype are required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "dns", "custom", record.Name, record.Type)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint.String(), strings.NewReader(record.Value))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// GetARecords Returns all custom A records for the specified query name.
// https://mailinabox.email/api-docs.html#operation/getDnsCustomARecordsForQName
func (s *DNSService) GetARecords(ctx context.Context, name string) ([]Record, error) {
	if name == "" {
		return nil, errors.New("qname is required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "dns", "custom", name)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var results []Record

	err = s.client.doJSON(req, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// AddARecord Adds a custom DNS A record for the specified query name.
// https://mailinabox.email/api-docs.html#operation/addDnsCustomARecord
func (s *DNSService) AddARecord(ctx context.Context, name, value string) (string, error) {
	if name == "" {
		return "", errors.New("qname is required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "dns", "custom", name)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(value))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// UpdateARecord Updates an existing DNS custom A record value for the specified qname.
// https://mailinabox.email/api-docs.html#operation/updateDnsCustomARecord
func (s *DNSService) UpdateARecord(ctx context.Context, name, value string) (string, error) {
	if name == "" {
		return "", errors.New("qname is required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "dns", "custom", name)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, endpoint.String(), strings.NewReader(value))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// RemoveARecord Removes a DNS custom A record for the specified domain & value.
// https://mailinabox.email/api-docs.html#operation/removeDnsCustomARecord
func (s *DNSService) RemoveARecord(ctx context.Context, name, value string) (string, error) {
	if name == "" {
		return "", errors.New("qname is required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "dns", "custom", name)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint.String(), strings.NewReader(value))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// GetDump Returns all DNS records.
// https://mailinabox.email/api-docs.html#operation/getDnsDump
func (s *DNSService) GetDump(ctx context.Context) ([]Zone, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "dns", "dump")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var results Zones

	err = s.client.doJSON(req, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
