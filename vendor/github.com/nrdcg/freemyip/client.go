// Package freemyip contains a client of the DNS API of freemyip.
package freemyip

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	querystring "github.com/google/go-querystring/query"
)

// RootDomain the root domain of all domains.
const RootDomain = "freemyip.com"

const defaultBaseURL = "https://freemyip.com/update"

const (
	codeError = "ERROR"
	codeOK    = "OK"
)

type query struct {
	Token   string `url:"token"`
	Domain  string `url:"domain"`
	TXT     string `url:"txt,omitempty"`
	MyIP    string `url:"myip,omitempty"`
	Delete  string `url:"delete,omitempty"`
	Verbose string `url:"verbose,omitempty"`
}

// Client an API client for freemyip.
type Client struct {
	HTTPClient *http.Client
	baseURL    *url.URL

	token   string
	verbose bool
}

// New creates a new Client.
func New(token string, verbose bool) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	return &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    baseURL,
		token:      token,
		verbose:    verbose,
	}
}

// UpdateDomain updates a domain.
//   - `domain` is the custom part of the real domain. (ex: `YOUR_DOMAIN` in `YOUR_DOMAIN.freemyip.com`)
//   - `myIP` is optional.
func (c Client) UpdateDomain(ctx context.Context, domain, myIP string) (string, error) {
	q := query{
		Token:   c.token,
		Domain:  fmt.Sprintf("%s.%s", domain, RootDomain),
		MyIP:    myIP,
		Verbose: boolToString(c.verbose),
	}

	return c.do(ctx, q)
}

// DeleteDomain deletes a domain.
//   - `domain` is the custom part of the real domain. (ex: `YOUR_DOMAIN` in `YOUR_DOMAIN.freemyip.com`)
func (c Client) DeleteDomain(ctx context.Context, domain string) (string, error) {
	q := query{
		Token:   c.token,
		Domain:  fmt.Sprintf("%s.%s", domain, RootDomain),
		Delete:  "yes",
		Verbose: boolToString(c.verbose),
	}

	return c.do(ctx, q)
}

// EditTXTRecord creates or updates a TXT record value for a domain.
//   - `domain` is the custom part of the real domain. (ex: `YOUR_DOMAIN` in `YOUR_DOMAIN.freemyip.com`)
//   - `value` is the TXT record content.
func (c Client) EditTXTRecord(ctx context.Context, domain, value string) (string, error) {
	q := query{
		Token:   c.token,
		Domain:  fmt.Sprintf("%s.%s", domain, RootDomain),
		TXT:     value,
		Verbose: boolToString(c.verbose),
	}

	return c.do(ctx, q)
}

// DeleteTXTRecord delete a TXT record for a domain.
//   - `domain` is the custom part of the real domain. (ex: `YOUR_DOMAIN` in `YOUR_DOMAIN.freemyip.com`)
//   - `value` is the TXT record content.
func (c Client) DeleteTXTRecord(ctx context.Context, domain string) (string, error) {
	q := query{
		Token:   c.token,
		Domain:  fmt.Sprintf("%s.%s", domain, RootDomain),
		TXT:     "null",
		Verbose: boolToString(c.verbose),
	}

	return c.do(ctx, q)
}

func (c Client) do(ctx context.Context, q query) (string, error) {
	endpoint, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "/"))
	if err != nil {
		return "", fmt.Errorf("URL parsing: %w", err)
	}

	values, err := querystring.Values(q)
	if err != nil {
		return "", fmt.Errorf("query parameters: %w", err)
	}

	endpoint.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return "", fmt.Errorf("creates request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do API call: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		all, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("error: %d: %s", resp.StatusCode, string(all))
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reads response body: %w", err)
	}

	body := strings.TrimSpace(string(all))

	parts := strings.SplitN(body, "\n", 2)

	switch parts[0] {
	case codeError:
		return "", errors.New(strings.Join(parts, " "))
	case codeOK:
		return body, nil
	default:
		return "", errors.New(strings.Join(parts, " "))
	}
}

func boolToString(v bool) string {
	if v {
		return "yes"
	}

	return ""
}
