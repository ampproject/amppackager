package amppackager

import (
	"fmt"
	"log"
	"net/http"
)

// A simple type that encodes an internal message to be logged and an HTTP
// status code to be used for the external error message.
type HttpError struct {
	InternalMsg string
	StatusCode  int
}

func NewHttpError(statusCode int, msg ...interface{}) *HttpError {
	return &HttpError{fmt.Sprint(msg), statusCode}
}

func (e HttpError) Error() string { return e.InternalMsg }

func (e HttpError) ExternalMsg() string {
	switch e.StatusCode {
	case http.StatusBadRequest:
		return "400 bad request"
	case http.StatusInternalServerError:
		return "500 internal server error"
	case http.StatusBadGateway:
		return "502 bad gateway"
	default:
		return "error"
	}
}

func (e HttpError) LogAndRespond(resp http.ResponseWriter) {
	log.Println(e.InternalMsg)
	http.Error(resp, e.ExternalMsg(), e.StatusCode)
}
