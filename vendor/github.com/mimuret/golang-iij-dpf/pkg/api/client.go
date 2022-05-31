package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"
)

const DefaultEndpoint = "https://api.dns-platform.jp/dpf/v1"

type ClientInterface interface {
	Read(ctx context.Context, s Spec) (requestID string, err error)
	List(ctx context.Context, s ListSpec, keywords SearchParams) (requestID string, err error)
	ListAll(ctx context.Context, s CountableListSpec, keywords SearchParams) (requestID string, err error)
	Count(ctx context.Context, s CountableListSpec, keywords SearchParams) (requestID string, err error)
	Update(ctx context.Context, s Spec, body interface{}) (requestID string, err error)
	Create(ctx context.Context, s Spec, body interface{}) (requestID string, err error)
	Apply(ctx context.Context, s Spec, body interface{}) (requestID string, err error)
	Delete(ctx context.Context, s Spec) (requestID string, err error)
	Cancel(ctx context.Context, s Spec) (requestID string, err error)
	WatchRead(ctx context.Context, interval time.Duration, s Spec) error
	WatchList(ctx context.Context, interval time.Duration, s ListSpec, keyword SearchParams) error
	WatchListAll(ctx context.Context, interval time.Duration, s CountableListSpec, keyword SearchParams) error
}

var _ ClientInterface = &Client{}

type Client struct {
	Endpoint string
	Token    string
	logger   Logger

	Client       *http.Client
	LastRequest  *RequestInfo
	LastResponse *ResponseInfo
	Json         JsonApiInterface
}

type RequestInfo struct {
	Method string
	Url    string
	Body   []byte
}

type ResponseInfo struct {
	Response *http.Response
	Body     []byte
}

func NewClient(token string, endpoint string, logger Logger) *Client {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}
	if logger == nil {
		logger = NewStdLogger(os.Stderr, "dpf-client", 0, 4)
	}
	return &Client{Endpoint: endpoint, Token: token, logger: logger, Client: http.DefaultClient, Json: &JsonAPIAdapter{}}
}

func (c *Client) marshalJson(action Action, body interface{}) ([]byte, error) {
	var (
		jsonBody []byte
		err      error
	)
	switch action {
	case ActionCreate:
		jsonBody, err = c.Json.MarshalCreate(body)
	case ActionUpdate:
		jsonBody, err = c.Json.MarshalUpdate(body)
	case ActionApply:
		jsonBody, err = c.Json.MarshalApply(body)
	default:
		return nil, fmt.Errorf("not support action `%s` with body request", action)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to encode body to json: %w", err)
	}
	return jsonBody, nil
}

func (c *Client) doSetup(ctx context.Context, spec Spec, action Action, body interface{}, params SearchParams) (*http.Request, error) {
	var r io.Reader
	if action == ActionCount {
		_, ok := spec.(CountableListSpec)
		if !ok {
			return nil, fmt.Errorf("spec is not CountableListSpec")
		}
	}
	c.LastRequest = &RequestInfo{}
	c.LastResponse = nil
	// create URL
	method, path := spec.GetPathMethod(action)
	if path == "" {
		return nil, fmt.Errorf("not support action %s", action)
	}
	c.LastRequest.Method = method
	url := c.Endpoint + path
	if params != nil {
		p, err := params.GetValues()
		if err != nil {
			return nil, fmt.Errorf("failed to get search params: %w", err)
		}
		url += "?" + p.Encode()
	}
	c.LastRequest.Url = url
	c.logger.Debugf("method: %s request-url: %s", method, url)
	// make request body
	if body != nil {
		jsonBody, err := c.marshalJson(action, body)
		if err != nil {
			return nil, err
		}
		c.logger.Tracef("request-body: `%s`", string(jsonBody))
		c.LastRequest.Body = jsonBody
		r = bytes.NewBuffer(jsonBody)
	}

	// make request
	req, err := http.NewRequest(method, url, r)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	// authorized
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")

	return req.WithContext(ctx), nil
}

func (c *Client) Do(ctx context.Context, spec Spec, action Action, body interface{}, params SearchParams) (requestID string, err error) {
	req, err := c.doSetup(ctx, spec, action, body, params)
	if err != nil {
		return "", err
	}
	// request
	resp, err := c.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get http response: %w", err)
	}
	defer resp.Body.Close()
	c.LastResponse = &ResponseInfo{
		Response: resp,
	}
	// get body
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to get http response body: %w", err)
	}
	c.LastResponse.Body = bs
	c.logger.Debugf("status-code: `%d`", resp.StatusCode)
	c.logger.Tracef("response-body: `%s`", string(bs))

	// if statiscode is error, response body type is BadResponse or Plantext
	if resp.StatusCode >= 400 {
		badRequest := &BadResponse{StatusCode: resp.StatusCode}
		if err := c.Json.UnmarshalRead(bs, badRequest); err != nil {
			return "", fmt.Errorf("failed to request: status code: %d body: %s err: %w", resp.StatusCode, string(bs), err)
		}
		return badRequest.RequestID, badRequest
	}

	// parse raw response
	rawResponse := &RawResponse{}
	if err := c.Json.UnmarshalRead(bs, rawResponse); err != nil {
		// maybe not executed
		return "", fmt.Errorf("failed to parse get response: %w", err)
	}
	if req.Method == http.MethodGet {
		if err := c.doReadResponse(action, spec, bs, rawResponse); err != nil {
			return rawResponse.RequestID, err
		}
	}

	// initialize process
	if d, ok := spec.(Initializer); ok {
		d.Init()
	}

	return rawResponse.RequestID, nil
}

func (c *Client) doReadResponse(action Action, spec Spec, bs []byte, rawResponse *RawResponse) error {
	switch {
	case action == ActionCount:
		// ActionCount
		count := &Count{}
		if err := c.Json.UnmarshalRead(rawResponse.Result, count); err != nil {
			return fmt.Errorf("failed to parse count response result: %w", err)
		}
		if cl, ok := spec.(CountableListSpec); ok {
			cl.SetCount(count.Count)
		}
	case rawResponse.Result != nil:
		// ActionRead
		if err := c.Json.UnmarshalRead(rawResponse.Result, spec); err != nil {
			return fmt.Errorf("failed to parse response result: %w", err)
		}
	case rawResponse.Results != nil:
		// ActionList
		listSpec, ok := spec.(ListSpec)
		if !ok {
			return fmt.Errorf("not support ListSpec %s", spec.GetName())
		}
		items := listSpec.GetItems()
		if err := c.Json.UnmarshalRead(rawResponse.Results, items); err != nil {
			return fmt.Errorf("failed to parse list response results: %w", err)
		}
	default:
		if err := c.Json.UnmarshalRead(bs, spec); err != nil {
			return fmt.Errorf("failed to parse response result: %w", err)
		}
	}
	return nil
}

func (c *Client) Read(ctx context.Context, s Spec) (requestID string, err error) {
	return c.Do(ctx, s, ActionRead, nil, nil)
}

func (c *Client) List(ctx context.Context, s ListSpec, keywords SearchParams) (requestID string, err error) {
	return c.Do(ctx, s, ActionList, nil, keywords)
}

func (c *Client) ListAll(ctx context.Context, s CountableListSpec, keywords SearchParams) (requestID string, err error) {
	req, err := c.Count(ctx, s, keywords)
	if err != nil {
		return req, err
	}

	if keywords == nil {
		keywords = &CommonSearchParams{}
		keywords.SetLimit(s.GetMaxLimit())
	}

	count := s.GetCount()
	cList := DeepCopyCountableListSpec(s)

	for offset := int32(0); offset < count; offset += keywords.GetLimit() {
		keywords.SetOffset(offset)
		req, err = c.List(ctx, cList, keywords)
		if err != nil {
			return req, err
		}
		for i := 0; i < cList.Len(); i++ {
			s.AddItem(cList.Index(i))
		}
	}
	return req, nil
}

func (c *Client) Count(ctx context.Context, s CountableListSpec, keywords SearchParams) (requestID string, err error) {
	return c.Do(ctx, s, ActionCount, nil, keywords)
}

func (c *Client) Update(ctx context.Context, s Spec, body interface{}) (requestID string, err error) {
	if body == nil {
		body = s
	}
	return c.Do(ctx, s, ActionUpdate, body, nil)
}

func (c *Client) Create(ctx context.Context, s Spec, body interface{}) (requestID string, err error) {
	if body == nil {
		body = s
	}
	return c.Do(ctx, s, ActionCreate, body, nil)
}

func (c *Client) Apply(ctx context.Context, s Spec, body interface{}) (requestID string, err error) {
	if body == nil {
		body = s
	}
	return c.Do(ctx, s, ActionApply, body, nil)
}

func (c *Client) Delete(ctx context.Context, s Spec) (requestID string, err error) {
	return c.Do(ctx, s, ActionDelete, nil, nil)
}

func (c *Client) Cancel(ctx context.Context, s Spec) (requestID string, err error) {
	return c.Do(ctx, s, ActionCancel, nil, nil)
}

func (c *Client) watch(ctx context.Context, interval time.Duration, f func(context.Context) (keep bool, err error)) error {
	if interval < time.Second {
		return fmt.Errorf("interval must greater than equals to 1s")
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
LOOP:
	for {
		select {
		case <-ticker.C:
			loopBreak, err := f(ctx)
			if err != nil {
				return err
			}
			if loopBreak {
				break LOOP
			}
		case <-ctx.Done():
			break LOOP
		}
	}
	return ctx.Err()
}

// ctx should set Deadline or Timeout
// interval must be grater than equals to 1s
// s is Readable Spec.
func (c *Client) WatchRead(ctx context.Context, interval time.Duration, s Spec) error {
	org := DeepCopySpec(s)
	return c.watch(ctx, interval, func(cctx context.Context) (bool, error) {
		_, err := c.Read(cctx, s)
		if err != nil {
			return true, err
		}
		if reflect.DeepEqual(s, org) {
			return false, nil
		}
		return true, nil
	})
}

// ctx should set Deadline or Timeout
// interval must be grater than equals to 1s
// s is ListAble Spec.
func (c *Client) WatchList(ctx context.Context, interval time.Duration, s ListSpec, keyword SearchParams) error {
	org := DeepCopyListSpec(s)
	return c.watch(ctx, interval, func(cctx context.Context) (bool, error) {
		_, err := c.List(cctx, s, keyword)
		if err != nil {
			return true, err
		}
		if reflect.DeepEqual(s, org) {
			return false, nil
		}
		return true, nil
	})
}

// ctx should set Deadline or Timeout
// interval must be grater than equals to 1s
// s is CountableListSpec Spec.
func (c *Client) WatchListAll(ctx context.Context, interval time.Duration, s CountableListSpec, keyword SearchParams) error {
	copySpec := DeepCopyCountableListSpec(s)
	copySpec.ClearItems()
	err := c.watch(ctx, interval, func(cctx context.Context) (bool, error) {
		_, err := c.ListAll(cctx, copySpec, keyword)
		if err != nil {
			return true, err
		}
		if reflect.DeepEqual(s, copySpec) {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return err
	}
	s.ClearItems()
	for i := 0; i < copySpec.Len(); i++ {
		s.AddItem(copySpec.Index(i))
	}
	s.SetCount(int32(copySpec.Len()))
	return nil
}
