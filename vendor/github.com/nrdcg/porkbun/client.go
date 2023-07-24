// Package porkbun contains a client of the DNS API of Porkdun.
package porkbun

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

const defaultBaseURL = "https://porkbun.com/api/json/v3/"

const statusSuccess = "SUCCESS"

// DefaultTTL The minimum and the default is 300 seconds.
const DefaultTTL = "300"

// Client an API client for Porkdun.
type Client struct {
	secretAPIKey string
	apiKey       string

	BaseURL    *url.URL
	HTTPClient *http.Client
}

// New creates a new Client.
func New(secretAPIKey, apiKey string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	return &Client{
		secretAPIKey: secretAPIKey,
		apiKey:       apiKey,
		BaseURL:      baseURL,
		HTTPClient:   &http.Client{Timeout: 10 * time.Second},
	}
}

// Ping tests communication with the API.
func (c *Client) Ping(ctx context.Context) (string, error) {
	endpoint, err := c.BaseURL.Parse(path.Join(c.BaseURL.Path, "ping"))
	if err != nil {
		return "", fmt.Errorf("failed to parse endpoint: %w", err)
	}

	respBody, err := c.do(ctx, endpoint, nil)
	if err != nil {
		return "", err
	}

	pingResp := pingResponse{}
	err = json.Unmarshal(respBody, &pingResp)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if pingResp.Status.Status != statusSuccess {
		return "", pingResp.Status
	}

	return pingResp.YourIP, nil
}

// CreateRecord creates a DNS record.
//
//	name (optional): The subdomain for the record being created, not including the domain itself. Leave blank to create a record on the root domain. Use * to create a wildcard record.
//	type: The type of record being created. Valid types are: A, MX, CNAME, ALIAS, TXT, NS, AAAA, SRV, TLSA, CAA
//	content: The answer content for the record.
//	ttl (optional): The time to live in seconds for the record. The minimum and the default is 300 seconds.
//	prio (optional) The priority of the record for those that support it.
func (c *Client) CreateRecord(ctx context.Context, domain string, record Record) (int, error) {
	endpoint, err := c.BaseURL.Parse(path.Join(c.BaseURL.Path, "dns", "create", domain))
	if err != nil {
		return 0, fmt.Errorf("failed to parse endpoint: %w", err)
	}

	respBody, err := c.do(ctx, endpoint, record)
	if err != nil {
		return 0, err
	}

	createResp := createResponse{}
	err = json.Unmarshal(respBody, &createResp)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if createResp.Status.Status != statusSuccess {
		return 0, createResp.Status
	}

	return createResp.ID, nil
}

// EditRecord edits a DNS record.
//
//	name (optional): The subdomain for the record being created, not including the domain itself. Leave blank to create a record on the root domain. Use * to create a wildcard record.
//	type: The type of record being created. Valid types are: A, MX, CNAME, ALIAS, TXT, NS, AAAA, SRV, TLSA, CAA
//	content: The answer content for the record.
//	ttl (optional): The time to live in seconds for the record. The minimum and the default is 300 seconds.
//	prio (optional) The priority of the record for those that support it.
func (c *Client) EditRecord(ctx context.Context, domain string, id int, record Record) error {
	endpoint, err := c.BaseURL.Parse(path.Join(c.BaseURL.Path, "dns", "edit", domain, strconv.Itoa(id)))
	if err != nil {
		return fmt.Errorf("failed to parse endpoint: %w", err)
	}

	respBody, err := c.do(ctx, endpoint, record)
	if err != nil {
		return err
	}

	statusResp := Status{}
	err = json.Unmarshal(respBody, &statusResp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if statusResp.Status != statusSuccess {
		return statusResp
	}

	return nil
}

// DeleteRecord deletes a specific DNS record.
func (c *Client) DeleteRecord(ctx context.Context, domain string, id int) error {
	endpoint, err := c.BaseURL.Parse(path.Join(c.BaseURL.Path, "dns", "delete", domain, strconv.Itoa(id)))
	if err != nil {
		return fmt.Errorf("failed to parse endpoint: %w", err)
	}

	respBody, err := c.do(ctx, endpoint, nil)
	if err != nil {
		return err
	}

	statusResp := Status{}
	err = json.Unmarshal(respBody, &statusResp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if statusResp.Status != statusSuccess {
		return statusResp
	}

	return nil
}

// RetrieveRecords retrieve all editable DNS records associated with a domain.
func (c *Client) RetrieveRecords(ctx context.Context, domain string) ([]Record, error) {
	endpoint, err := c.BaseURL.Parse(path.Join(c.BaseURL.Path, "dns", "retrieve", domain))
	if err != nil {
		return nil, fmt.Errorf("failed to parse endpoint: %w", err)
	}

	respBody, err := c.do(ctx, endpoint, nil)
	if err != nil {
		return nil, err
	}

	retrieveResp := retrieveResponse{}
	err = json.Unmarshal(respBody, &retrieveResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if retrieveResp.Status.Status != statusSuccess {
		return nil, retrieveResp.Status
	}

	return retrieveResp.Records, nil
}

func (c *Client) do(ctx context.Context, endpoint *url.URL, apiRequest interface{}) ([]byte, error) {
	request := authRequest{
		APIKey:       c.apiKey,
		SecretAPIKey: c.secretAPIKey,
		apiRequest:   apiRequest,
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &ServerError{
			StatusCode: resp.StatusCode,
			Message:    string(respBody),
		}
	}

	return respBody, nil
}
