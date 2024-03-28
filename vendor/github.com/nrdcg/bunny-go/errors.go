package bunny

import (
	"fmt"
	"net/http"
	"strings"
)

// HTTPError is returned by the Client when an unsuccessful HTTP response was
// returned or a response could not be processed.
// If the body of an unsuccessful HTTP response contains an APIError in the
// body, APIError is returned by the Client instead.
type HTTPError struct {
	// RequestURL is the address to which the request was sent that caused the error.
	RequestURL string
	// The HTTP response status code.
	StatusCode int
	// The raw http response body. It's nil if the response had no body or it could not be received.
	RespBody []byte
	// Errors contain errors that happened while receiving or processing the HTTP response.
	Errors []error
}

// Error returns a textual representation of the error.
func (e *HTTPError) Error() string {
	var res strings.Builder

	res.WriteString(fmt.Sprintf("http-request to %s failed: %s (%d)",
		e.RequestURL, http.StatusText(e.StatusCode), e.StatusCode,
	))

	if len(e.Errors) > 0 {
		res.WriteString(", errors: " + strings.Join(errorsToStrings(e.Errors), ", "))
	}

	return res.String()
}

func errorsToStrings(errs []error) []string {
	res := make([]string, 0, len(errs))

	for _, err := range errs {
		res = append(res, err.Error())
	}

	return res
}

// AuthenticationError represents an Unauthorized (401) HTTP error.
type AuthenticationError struct {
	Message string
}

// Error returns a textual representation of the error.
func (e *AuthenticationError) Error() string {
	return e.Message
}

// APIError represents an error that is returned by some Bunny API endpoints on
// failures.
type APIError struct {
	HTTPError
	ErrorKey string `json:"ErrorKey"`
	Field    string `json:"Field"`
	Message  string `json:"Message"`
}

// Error returns the string representation of the error.
// ErrorKey, Field and Message are omitted if they are empty.
func (e *APIError) Error() string {
	var res strings.Builder

	res.WriteString(e.HTTPError.Error())
	if e.ErrorKey != "" {
		res.WriteString(", ")
		res.WriteString(e.ErrorKey)

		if e.Field != "" {
			res.WriteString(": ")
			res.WriteString(e.Field)
		}
	} else {
		if e.Field != "" {
			res.WriteString(", ")
			res.WriteString(e.Field)
		}
	}

	if e.Message != "" {
		// Field and ErrorKey contains the same information then Message, no need to log them.
		res.WriteString(", ")
		res.WriteString(e.Message)
	}

	return res.String()
}
