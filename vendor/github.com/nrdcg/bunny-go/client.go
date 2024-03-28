// Package bunny provides functionality to interact with the Bunny CDN HTTP API.
package bunny

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
)

const (
	// BaseURL is the base URL of the Bunny CDN HTTP API.
	BaseURL = "https://api.bunny.net"
	// AccessKeyHeaderKey is the name of the HTTP header that contains the Bunny API key.
	AccessKeyHeaderKey = "AccessKey"
	// DefaultUserAgent is the default value of the sent HTTP User-Agent header.
	DefaultUserAgent = "bunny-go"
)

const (
	hdrContentTypeName = "content-type"
	contentTypeJSON    = "application/json"
)

// Logf is a log function signature.
type Logf func(format string, v ...interface{})

// Client is a Bunny CDN HTTP API Client.
type Client struct {
	baseURL *url.URL
	apiKey  string

	httpClient       http.Client
	httpRequestLogf  Logf
	httpResponseLogf Logf
	logf             Logf
	userAgent        string

	PullZone     *PullZoneService
	StorageZone  *StorageZoneService
	DNSZone      *DNSZoneService
	VideoLibrary *VideoLibraryService
}

var discardLogF = func(string, ...interface{}) {}

// NewClient returns a new bunny.net API client.
// The APIKey can be found in on the Account Settings page.
//
// Bunny.net API docs: https://support.bunny.net/hc/en-us/articles/360012168840-Where-do-I-find-my-API-key-
func NewClient(APIKey string, opts ...Option) *Client {
	clt := Client{
		baseURL:          mustParseURL(BaseURL),
		apiKey:           APIKey,
		httpClient:       *http.DefaultClient,
		userAgent:        DefaultUserAgent,
		httpRequestLogf:  discardLogF,
		httpResponseLogf: discardLogF,
		logf:             discardLogF,
	}

	clt.PullZone = &PullZoneService{client: &clt}
	clt.StorageZone = &StorageZoneService{client: &clt}
	clt.DNSZone = &DNSZoneService{client: &clt}
	clt.VideoLibrary = &VideoLibraryService{client: &clt}

	for _, opt := range opts {
		opt(&clt)
	}

	return &clt
}

func mustParseURL(urlStr string) *url.URL {
	res, err := url.Parse(urlStr)
	if err != nil {
		panic(fmt.Sprintf("Parsing url: %s failed: %s", urlStr, err))
	}

	return res
}

// newRequest creates an bunny.net API request.
// urlStr maybe absolute or relative, if it is relative it is joined with
// client.baseURL.
func (c *Client) newRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	url, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set(AccessKeyHeaderKey, c.apiKey)
	req.Header.Add("Accept", contentTypeJSON)
	req.Header.Set("User-Agent", c.userAgent)

	if body != nil {
		req.Header.Set(hdrContentTypeName, contentTypeJSON)
	}

	return req, nil
}

// newGetRequest creates an bunny.NET API GET request.
// params must be a struct or nil, it is encoded into a query parameter.
// The struct must contain  `url` tags of the go-querystring package.
func (c *Client) newGetRequest(urlStr string, params interface{}) (*http.Request, error) {
	if params != nil {
		queryvals, err := query.Values(params)
		if err != nil {
			return nil, err
		}
		urlStr = urlStr + "?" + queryvals.Encode()
	}

	return c.newRequest(http.MethodGet, urlStr, nil)
}

func toJSON(data interface{}) (io.Reader, error) {
	var buf io.ReadWriter

	if data == nil {
		return http.NoBody, nil
	}

	buf = &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(data); err != nil {
		return nil, err
	}

	return buf, nil
}

// newPostRequest creates a bunny.NET API POST request.
// If body is not nil, it is encoded as JSON and send as HTTP-Body.
func (c *Client) newPostRequest(urlStr string, body interface{}) (*http.Request, error) {
	buf, err := toJSON(body)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodPost, urlStr, buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// newDeleteRequest creates a bunny.NET API DELETE request.
// If body is not nil, it is encoded as JSON and send as HTTP-Body.
func (c *Client) newDeleteRequest(urlStr string, body interface{}) (*http.Request, error) {
	buf, err := toJSON(body)
	if err != nil {
		return nil, err
	}

	return c.newRequest(http.MethodDelete, urlStr, buf)
}

// newPutRequest creates a bunny.NET API PUT request.
// If body is not nil, it is encoded as JSON and sent as a HTTP-Body.
func (c *Client) newPutRequest(urlStr string, body interface{}) (*http.Request, error) {
	buf, err := toJSON(body)
	if err != nil {
		return nil, err
	}

	return c.newRequest(http.MethodPut, urlStr, buf)
}

// sendRequest sends a http Request to the bunny API.
// If the server returns a 2xx status code with an response body, the body is
// unmarshaled as JSON into result.
// If the ctx times out ctx.Error() is returned.
// If sending the response fails (http.Client.Do), the error will be returned.
// If the server returns an 401 error, an AuthenticationError error is returned.
// If the server returned an error and contains an APIError as JSON in the body,
// an APIError is returned.
// If the server returned a status code that is not 2xx an HTTPError is returned.
// If the HTTP request was successful, the response body is read and
// unmarshaled into result.
func (c *Client) sendRequest(ctx context.Context, req *http.Request, result interface{}) error {
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	logReqID := c.logRequest(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		if urlErr, ok := err.(*url.Error); ok {
			if urlErr.Timeout() && ctx.Err() != nil {
				return ctx.Err()
			}
		}

		return err
	}

	c.logResponse(resp, logReqID)

	defer resp.Body.Close() //nolint: errcheck

	if err := c.checkResp(req, resp); err != nil {
		return err
	}

	return c.unmarshalHTTPJSONBody(resp, req.URL.String(), result)
}

func ensureJSONContentType(hdr http.Header) error {
	val := hdr.Get(hdrContentTypeName)
	if val == "" {
		return fmt.Errorf("%s header is missing or empty", hdrContentTypeName)
	}

	contentType, _, err := mime.ParseMediaType(val)
	if err != nil {
		return fmt.Errorf("could not parse %s header value: %w", hdrContentTypeName, err)
	}

	if contentType != contentTypeJSON {
		return fmt.Errorf("expected %s to be %q, got: %q", hdrContentTypeName, contentTypeJSON, contentType)
	}

	return nil
}

// checkResp checks if the resp indicates that the request was successful.
// If it wasn't an error is returned.
func (c *Client) checkResp(req *http.Request, resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			// ignore connection errors causing that the body can
			// not be received
			msg = []byte(http.StatusText(http.StatusUnauthorized))
		}

		return &AuthenticationError{
			Message: string(msg),
		}

	default:
		httpErr := HTTPError{
			RequestURL: req.URL.String(),
			StatusCode: resp.StatusCode,
		}

		return c.parseHTTPRespErrBody(resp, &httpErr)
	}
}

// parseHTTPRespErrBody processes the body of an http.Response with an non 2xx
// status code.
// If the response body is empty, baseErr is returned.
// If the body could no be parsed because of an error, the occurred errors are
// added to baseErr and baseErr is returned.
// If the body contains json data it is parsed and an APIError is returned.
func (c *Client) parseHTTPRespErrBody(resp *http.Response, baseErr *HTTPError) error {
	var err error

	baseErr.RespBody, err = io.ReadAll(resp.Body)
	if err != nil {
		baseErr.Errors = append(baseErr.Errors, fmt.Errorf("reading response body failed: %w", err))
		return baseErr
	}

	if len(baseErr.RespBody) == 0 {
		return baseErr
	}

	err = ensureJSONContentType(resp.Header)
	if err != nil {
		baseErr.Errors = append(baseErr.Errors, fmt.Errorf("processing response failed: %w", err))
		return baseErr
	}

	var apiErr APIError
	if err := json.Unmarshal(baseErr.RespBody, &apiErr); err != nil {
		baseErr.Errors = append(baseErr.Errors, fmt.Errorf("could not parse body as APIError: %w", err))
		return baseErr
	}

	apiErr.HTTPError = *baseErr
	return &apiErr
}

func (c *Client) unmarshalHTTPJSONBody(resp *http.Response, reqURL string, result interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &HTTPError{
			RequestURL: reqURL,
			StatusCode: resp.StatusCode,
			Errors:     []error{fmt.Errorf("reading response body failed: %w", err)},
		}
	}

	if len(body) == 0 {
		if result != nil {
			return &HTTPError{
				RequestURL: reqURL,
				StatusCode: resp.StatusCode,
				Errors:     []error{fmt.Errorf("response has no body, expected a json %T response body", result)},
			}
		}

		return nil
	}

	if result == nil {
		c.logf("http-response contains body but none was expected")
		return nil
	}

	err = ensureJSONContentType(resp.Header)
	if err != nil {
		return &HTTPError{
			RequestURL: reqURL,
			RespBody:   body,
			StatusCode: resp.StatusCode,
			Errors:     []error{fmt.Errorf("processing response failed: %w", err)},
		}
	}

	if err := json.Unmarshal(body, result); err != nil {
		return &HTTPError{
			RequestURL: reqURL,
			RespBody:   body,
			StatusCode: resp.StatusCode,
			Errors:     []error{fmt.Errorf("could not parse body as %T: %w", result, err)},
		}
	}

	return nil
}

// logRequest dumps the http request to the http request logger and returns a
// unique request identifier. The identifier can be used when logging the
// response for the request, to make it easier to associate request and
// response log messages.
func (c *Client) logRequest(req *http.Request) string {
	if c.httpRequestLogf == nil {
		return ""
	}

	logReqID := uuid.New().String()

	// hide the access key in the dumped request
	accessKey := req.Header.Get(AccessKeyHeaderKey)
	if accessKey != "" {
		req.Header.Set(AccessKeyHeaderKey, "***hidden***")
		defer func() { req.Header.Set(AccessKeyHeaderKey, accessKey) }()
	}

	debugReq, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		c.httpRequestLogf("dumping http request (reqID: %s) failed: %s", logReqID, err)
		return logReqID
	}

	c.httpRequestLogf("sending http-request (reqID: %s): %s", logReqID, string(debugReq))

	return logReqID
}

func (c *Client) logResponse(resp *http.Response, logReqID string) {
	if c.httpResponseLogf == nil {
		return
	}

	debugResp, err := httputil.DumpResponse(resp, true)
	if err != nil {
		c.httpRequestLogf("dumping http response (reqID: %s) failed: %s", logReqID, err)
		return
	}

	c.httpRequestLogf("received http-response (reqID: %s): %s", logReqID, string(debugResp))
}
