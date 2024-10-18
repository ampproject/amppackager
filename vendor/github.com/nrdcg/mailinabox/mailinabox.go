package mailinabox

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/nrdcg/mailinabox/errutils"
)

type service struct {
	client *Client
}

// Client the Mail-in-a-Box client.
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL

	email    string
	password string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	DNS    *DNSService
	User   *UserService
	Mail   *MailService
	System *SystemService
}

// New creates a new Client.
func New(apiURL, email, password string) (*Client, error) {
	baseURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	client := &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    baseURL,
		email:      email,
		password:   password,
	}

	client.common.client = client

	client.DNS = (*DNSService)(&client.common)
	client.User = (*UserService)(&client.common)
	client.Mail = (*MailService)(&client.common)
	client.System = (*SystemService)(&client.common)

	return client, nil
}

func (c *Client) doJSON(req *http.Request, result any) error {
	req.SetBasicAuth(c.email, c.password)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errutils.NewHTTPDoError(req, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return errutils.NewUnexpectedResponseStatusCodeError(req, resp)
	}

	if result == nil {
		return nil
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return errutils.NewReadResponseError(req, resp.StatusCode, err)
	}

	err = json.Unmarshal(raw, result)
	if err != nil {
		return errutils.NewUnmarshalError(req, resp.StatusCode, raw, err)
	}

	return nil
}

func (c *Client) doPlain(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.email, c.password)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errutils.NewHTTPDoError(req, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, errutils.NewUnexpectedResponseStatusCodeError(req, resp)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errutils.NewReadResponseError(req, resp.StatusCode, err)
	}

	return raw, nil
}

func boolToIntStr(v bool) string {
	if v {
		return "1"
	}

	return "0"
}
