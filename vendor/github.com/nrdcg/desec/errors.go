package desec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// NotFound Not found error.
type NotFound struct {
	Detail string `json:"detail"`
}

func (n NotFound) Error() string {
	return n.Detail
}

// APIError error from API.
type APIError struct {
	StatusCode int
	err        error
}

func (e APIError) Error() string {
	return fmt.Sprintf("%d: %v", e.StatusCode, e.err)
}

// Unwrap unwraps error.
func (e APIError) Unwrap() error {
	return e.err
}

func readError(resp *http.Response, er error) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &APIError{
			StatusCode: resp.StatusCode,
			err:        fmt.Errorf("failed to read response body: %w", err),
		}
	}

	err = json.Unmarshal(body, er)
	if err != nil {
		return &APIError{
			StatusCode: resp.StatusCode,
			err:        fmt.Errorf("failed to unmarshall response body: %w: %s", err, string(body)),
		}
	}

	return &APIError{
		StatusCode: resp.StatusCode,
		err:        er,
	}
}

func readRawError(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &APIError{
			StatusCode: resp.StatusCode,
			err:        fmt.Errorf("failed to read response body: %w", err),
		}
	}

	return &APIError{StatusCode: resp.StatusCode, err: fmt.Errorf("body: %s", string(body))}
}
