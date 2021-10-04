package desec

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/hashicorp/go-retryablehttp"
)

const defaultBaseURL = "https://desec.io/api/v1/"

type httpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type service struct {
	client *Client
}

// ClientOptions the options of the Client.
type ClientOptions struct {
	// HTTPClient HTTP client used to communicate with the API.
	HTTPClient *http.Client

	// Maximum number of retries
	RetryMax int

	// Customer logger instance. Can be either Logger or LeveledLogger
	Logger interface{}
}

// NewDefaultClientOptions creates a new ClientOptions with default values.
func NewDefaultClientOptions() ClientOptions {
	return ClientOptions{
		HTTPClient: http.DefaultClient,
		RetryMax:   5,
		Logger:     nil,
	}
}

// Client deSEC API client.
type Client struct {
	// Base URL for API requests.
	BaseURL string

	httpClient httpDoer

	token string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the deSEC API.
	Account *AccountService
	Tokens  *TokensService
	Records *RecordsService
	Domains *DomainsService
}

// New creates a new Client.
func New(token string, opts ClientOptions) *Client {
	// https://github.com/desec-io/desec-stack/blob/main/docs/rate-limits.rst
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = opts.RetryMax
	retryClient.HTTPClient = opts.HTTPClient
	retryClient.Logger = opts.Logger

	client := &Client{
		httpClient: retryClient.StandardClient(),
		BaseURL:    defaultBaseURL,
		token:      token,
	}

	client.common.client = client

	client.Account = (*AccountService)(&client.common)
	client.Tokens = (*TokensService)(&client.common)
	client.Records = (*RecordsService)(&client.common)
	client.Domains = (*DomainsService)(&client.common)

	return client
}

func (c *Client) newRequest(ctx context.Context, method string, endpoint fmt.Stringer, reqBody interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)

	if reqBody != nil {
		err := json.NewEncoder(buf).Encode(reqBody)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint.String(), buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.token))
	}

	return req, nil
}

func (c *Client) createEndpoint(parts ...string) (*url.URL, error) {
	base, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	endpoint, err := base.Parse(path.Join(base.Path, path.Join(parts...)))
	if err != nil {
		return nil, err
	}

	endpoint.Path += "/"

	return endpoint, nil
}

func handleResponse(resp *http.Response, respData interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &APIError{
			StatusCode: resp.StatusCode,
			err:        fmt.Errorf("failed to read response body: %w", err),
		}
	}

	err = json.Unmarshal(body, respData)
	if err != nil {
		return fmt.Errorf("failed to umarshal response body: %w", err)
	}

	return nil
}

func handleError(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusNotFound:
		return readError(resp, &NotFoundError{})
	default:
		return readRawError(resp)
	}
}
