// Package nodion contains a client of the DNS API of Nodion.
package nodion

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	querystring "github.com/google/go-querystring/query"
)

const defaultBaseURL = "https://api.nodion.com/v1/"

// Client the Nodion API client.
type Client struct {
	HTTPClient *http.Client
	baseURL    *url.URL
	apiToken   string
}

// NewClient creates a new Client.
func NewClient(apiToken string) (*Client, error) {
	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, err
	}

	if apiToken == "" {
		return nil, errors.New("API token is required")
	}

	return &Client{
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
		baseURL:    baseURL,
		apiToken:   apiToken,
	}, nil
}

// CreateZone To create a new DNS Zone.
// https://www.nodion.com/en/docs/dns/api/#post-dns-zone
func (c Client) CreateZone(ctx context.Context, name string) (*Zone, error) {
	endpoint := c.baseURL.JoinPath("dns_zones")

	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(Zone{Name: name})
	if err != nil {
		return nil, fmt.Errorf("encode request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	var result ZoneResponse
	err = c.do(req, &result)
	if err != nil {
		return nil, err
	}

	return &result.Zone, nil
}

// DeleteZone To delete an existing DNS Zone.
// https://www.nodion.com/en/docs/dns/api/#delete-dns-zone
func (c Client) DeleteZone(ctx context.Context, zoneID string) (bool, error) {
	endpoint := c.baseURL.JoinPath("dns_zones", zoneID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint.String(), http.NoBody)
	if err != nil {
		return false, fmt.Errorf("create request: %w", err)
	}

	var result DeleteResponse
	err = c.do(req, &result)
	if err != nil {
		return false, err
	}

	return result.Deleted, nil
}

// GetZones To list all existing DNS zones.
// https://www.nodion.com/en/docs/dns/api/#get-dns-zones
func (c Client) GetZones(ctx context.Context, filter *ZonesFilter) ([]Zone, error) {
	endpoint := c.baseURL.JoinPath("dns_zones")

	values, err := querystring.Values(filter)
	if err != nil {
		return nil, fmt.Errorf("create zones filter: %w", err)
	}

	endpoint.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	var result ZonesResponse
	err = c.do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Zones, nil
}

// CreateRecord To create a new Record for a DNS zone.
// https://www.nodion.com/en/docs/dns/api/#post-dns-record
func (c Client) CreateRecord(ctx context.Context, zoneID string, record Record) (*Record, error) {
	endpoint := c.baseURL.JoinPath("dns_zones", zoneID, "records")

	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(record)
	if err != nil {
		return nil, fmt.Errorf("encode request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	var result RecordResponse
	err = c.do(req, &result)
	if err != nil {
		return nil, err
	}

	return &result.Record, nil
}

// DeleteRecord To delete an existing Record for a DNS zone.
// https://www.nodion.com/en/docs/dns/api/#delete-dns-record
func (c Client) DeleteRecord(ctx context.Context, zoneID, recordID string) (bool, error) {
	endpoint := c.baseURL.JoinPath("dns_zones", zoneID, "records", recordID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint.String(), http.NoBody)
	if err != nil {
		return false, fmt.Errorf("create request: %w", err)
	}

	var result DeleteResponse
	err = c.do(req, &result)
	if err != nil {
		return false, err
	}

	return result.Deleted, nil
}

// GetRecords To list all existing Records of a DNS zone.
// https://www.nodion.com/en/docs/dns/api/#get-dns-records
func (c Client) GetRecords(ctx context.Context, zoneID string, filter *RecordsFilter) ([]Record, error) {
	endpoint := c.baseURL.JoinPath("dns_zones", zoneID, "records")

	values, err := querystring.Values(filter)
	if err != nil {
		return nil, fmt.Errorf("create records filter: %w", err)
	}

	endpoint.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	var result RecordsResponse
	err = c.do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Records, nil
}

func (c Client) do(req *http.Request, result any) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("API error: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode/100 != 2 {
		return readError(req.URL, resp)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}

	err = json.Unmarshal(raw, result)
	if err != nil {
		return fmt.Errorf("unmarshaling %T error [status code=%d]: %w: %s", result, resp.StatusCode, err, string(raw))
	}

	return nil
}

func readError(endpoint *url.URL, resp *http.Response) error {
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(toUnreadableBodyMessage(endpoint, content))
	}

	errAPI := &APIError{StatusCode: resp.StatusCode}

	if len(content) == 0 {
		errAPI.Errors = []string{http.StatusText(resp.StatusCode)}
		return errAPI
	}

	err = json.Unmarshal(content, errAPI)
	if err != nil {
		errAPI.Errors = []string{toUnreadableBodyMessage(endpoint, content)}
		return fmt.Errorf("unmarshaling error: %w", errAPI)
	}

	return errAPI
}

func toUnreadableBodyMessage(endpoint *url.URL, rawBody []byte) string {
	return fmt.Sprintf("the request %s received a response with a body which is an invalid format or not readable: %q", endpoint, string(rawBody))
}
