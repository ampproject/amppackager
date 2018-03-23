package amppackager

import (
	"fmt"
	"log"
	"net/http"
)

// HTTPError encodes an internal message to be logged and an HTTP status code
// to be used for the external error message.
type HTTPError struct {
	InternalMsg string
	StatusCode  int
}

func NewHTTPError(statusCode int, msg ...interface{}) *HTTPError {
	return &HTTPError{fmt.Sprint(msg), statusCode}
}

func (e HTTPError) Error() string { return e.InternalMsg }

func (e HTTPError) ExternalMsg() string {
	return http.StatusText(e.StatusCode)
}

func (e HTTPError) LogAndRespond(resp http.ResponseWriter) {
	log.Println(e.InternalMsg)
	http.Error(resp, e.ExternalMsg(), e.StatusCode)
}
