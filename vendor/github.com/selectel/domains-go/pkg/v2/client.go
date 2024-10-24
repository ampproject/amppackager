package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type (
	QueryParam map[string]string
	Client     struct {
		httpClient     *http.Client
		defaultHeaders http.Header
		BaseURL        string
	}
	ReturnTypes interface {
		Zone | List[Zone] | RRSet | List[RRSet]
	}
)

//nolint:exhaustruct
var _ DNSClient[Zone, RRSet] = &Client{}

func NewClient(apiURL string, httpClient *http.Client, defaultHeaders http.Header) DNSClient[Zone, RRSet] {
	return &Client{
		httpClient:     httpClient,
		defaultHeaders: defaultHeaders,
		BaseURL:        apiURL,
	}
}

// WithHeaders returns reference to a copy of the initial client
// with extra headers passed in params. Conflicting headers are replaced
// with new ones.
func (c *Client) WithHeaders(headers http.Header) DNSClient[Zone, RRSet] {
	temporaryClient := *c
	temporaryClient.defaultHeaders = temporaryClient.defaultHeaders.Clone()
	if temporaryClient.defaultHeaders == nil {
		temporaryClient.defaultHeaders = http.Header{}
	}
	for k, v := range headers {
		temporaryClient.defaultHeaders.Del(k)
		for _, value := range v {
			temporaryClient.defaultHeaders.Add(k, value)
		}
	}

	return &temporaryClient
}

// prepareRequest prepares request with default headers and additional content.
func (c *Client) prepareRequest(
	ctx context.Context,
	method, path string,
	body io.Reader,
	//nolint: unparam
	params, extraHeaders *map[string]string,
) (*http.Request, error) {
	path = fmt.Sprintf("%s%s", c.BaseURL, path)
	request, err := http.NewRequestWithContext(ctx, method, path, body)
	if err != nil {
		return nil, fmt.Errorf("prepare request: %w", err)
	}

	if c.defaultHeaders != nil {
		request.Header = c.defaultHeaders.Clone()
	}

	if extraHeaders != nil {
		for key, value := range *extraHeaders {
			request.Header.Add(key, value)
		}
	}

	addedParamsToHTTPQuery(request, params)

	return request.WithContext(ctx), nil
}

func addedParamsToHTTPQuery(request *http.Request, params *map[string]string) {
	if params == nil {
		return
	}
	urlQuery := request.URL.Query()
	for key, value := range *params {
		for _, val := range strings.Split(value, ",") {
			urlQuery.Add(key, val)
		}
	}
	request.URL.RawQuery = urlQuery.Encode()
}

func processRequest[RT ReturnTypes](client *http.Client, request *http.Request, err error) (*RT, error) {
	if err != nil {
		return nil, ErrInvalidRequestObj
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("processing request: %w", err)
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("processing response: %w", err)
	}
	resp, err := checkProccessResult[RT](response.StatusCode, body)

	return resp, err
}

func checkProccessResult[RT ReturnTypes](statusCode int, body []byte) (*RT, error) {
	switch {
	case statusCode == http.StatusNoContent && len(body) == 0:
		//nolint: nilnil
		return nil, nil
	case statusCode == http.StatusNotFound:
		return nil, ErrNotFound
	case statusCode < http.StatusBadRequest:
		var result RT
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, fmt.Errorf("processing good response: %w", err)
		}

		return &result, nil
	default:
		var result BadResponseError
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, fmt.Errorf("processing error response: %w", err)
		}
		result.Code = statusCode

		return nil, &result
	}
}
