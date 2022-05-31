package api

import (
	"encoding/json"
	"strings"
)

type ResponseCommon struct {
	RequestID string `read:"request_id"`
}

type RawResponse struct {
	ResponseCommon `read:",inline"`
	Result         json.RawMessage `read:"result,omitempty"`
	Results        json.RawMessage `read:"results,omitempty"`
}

const (
	ErrorTypeParamaterError  string = "ParameterError"
	ErrorTypeNotFound        string = "NotFound"
	ErrorTypeTooManyRequests string = "TooManyRequests"
	ErrorTypeSystemError     string = "SystemError"
	ErrorTypeGatewayTimeout  string = "GatewayTimeout"
)

type BadResponse struct {
	ResponseCommon `read:",inline"`
	StatusCode     int          `read:"-"`
	ErrorType      string       `read:"error_type"`
	ErrorMessage   string       `read:"error_message"`
	ErrorDetails   ErrorDetails `read:"error_details"`
}

func (r *BadResponse) Error() string {
	if r.ErrorType == ErrorTypeParamaterError {
		if r.IsAuthError() {
			return "Auth error"
		}
	}
	errorDetail := ""
	if len(r.ErrorDetails) > 0 {
		errorDetail = " Detail: " + r.ErrorDetails.Error()
	}
	return "ErrorType: " + r.ErrorType + " Message: " + r.ErrorMessage + errorDetail
}

func (r *BadResponse) IsStatusCode(code int) bool {
	return r.StatusCode == code
}

func (r *BadResponse) IsErrType(name string) bool {
	return r.ErrorType == name
}

func (r *BadResponse) IsErrMsg(msg string) bool {
	return r.ErrorMessage == msg
}

func (r *BadResponse) IsErrorCode(code string) (bool, string) {
	for _, errDetail := range r.ErrorDetails {
		if errDetail.Code == code {
			return true, errDetail.Attribute
		}
	}
	return false, ""
}

func (r *BadResponse) IsErrorCodeAttribute(code string, attribute string) bool {
	for _, errDetail := range r.ErrorDetails {
		if errDetail.Code == code && errDetail.Attribute == attribute {
			return true
		}
	}
	return false
}

func (r *BadResponse) IsAuthError() bool {
	if r.IsStatusCode(400) &&
		r.IsErrType(ErrorTypeParamaterError) &&
		r.IsErrorCodeAttribute("invalid", "access_token") {
		return true
	}
	return false
}

func (r *BadResponse) IsRequestFormatError() bool {
	if r.IsStatusCode(400) &&
		r.IsErrType(ErrorTypeParamaterError) &&
		r.IsErrMsg("JSON parse error occurred.") {
		return true
	}
	return false
}

func (r *BadResponse) IsParameterError() bool {
	if r.IsStatusCode(400) &&
		r.IsErrType(ErrorTypeParamaterError) &&
		!r.IsAuthError() &&
		!r.IsRequestFormatError() {
		return true
	}

	return false
}

func (r *BadResponse) IsNotFound() bool {
	return r.IsStatusCode(404)
}

func (r *BadResponse) IsTooManyRequests() bool {
	return r.IsStatusCode(429)
}

func (r *BadResponse) IsSystemError() bool {
	return r.IsStatusCode(500)
}

func (r *BadResponse) IsGatewayTimeout() bool {
	return r.IsStatusCode(504)
}

func (r *BadResponse) IsInvalidSchema() bool {
	return r.IsParameterError() && r.IsErrorCodeAttribute("invalid", "schema")
}

type ErrorDetails []ErrorDetail

func (e ErrorDetails) Error() string {
	res := []string{}
	for _, detail := range e {
		res = append(res, detail.Error())
	}
	return strings.Join(res, ", ")
}

type ErrorDetail struct {
	Code      string `read:"code"`
	Attribute string `read:"attribute"`
}

func (e ErrorDetail) Error() string {
	return e.Code + "=" + e.Attribute
}

type Count struct {
	Count int32 `read:"count" json:"-"`
}

func (c *Count) SetCount(v int32) { c.Count = v }
func (c *Count) GetCount() int32  { return c.Count }

type AsyncResponse struct {
	ResponseCommon `read:",inline"`
	JobsUrl        string `read:"jobs_url"`
}
