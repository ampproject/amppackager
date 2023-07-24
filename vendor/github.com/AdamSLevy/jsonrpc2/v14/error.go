// Copyright 2018 Adam S Levy
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package jsonrpc2

import "fmt"

// ErrorCode indicates the Error type that occurred.
//
// The error codes from and including -32768 to -32000 are reserved for
// pre-defined errors. Any code within this range, but not defined explicitly
// below is reserved for future use.
//
// See ErrorCodeMinReserved.
type ErrorCode int

// Official JSON-RPC 2.0 Spec Error Codes and Messages
const (
	// ErrorCodeMinReserved is the minimum reserved error code. Method
	// defined errors may be less than this value.
	ErrorCodeMinReserved ErrorCode = -32768

	// ErrorCodeParse means the received data was not valid JSON.
	ErrorCodeParse    ErrorCode = -32700
	ErrorMessageParse           = "Parse error"

	// ErrorCodeInvalidRequest means the received valid JSON is not a valid
	// Request object.
	ErrorCodeInvalidRequest    ErrorCode = -32600
	ErrorMessageInvalidRequest           = "Invalid Request"

	// ErrorCodeMethodNotFound means the requested method is not defined
	// for the HTTPRequestHandler.
	ErrorCodeMethodNotFound    ErrorCode = -32601
	ErrorMessageMethodNotFound           = "Method not found"

	// ErrorCodeInvalidParams means a method is called with invalid method
	// parameter(s).
	//
	// MethodFuncs are responsible for detecting and returning this error.
	ErrorCodeInvalidParams    ErrorCode = -32602
	ErrorMessageInvalidParams           = "Invalid params"

	// ErrorCodeInternal means an internal error occurred such as a
	// MethodFunc panic.
	ErrorCodeInternal    ErrorCode = -32603
	ErrorMessageInternal           = "Internal error"

	// ErrorCodeMaxReserved is the maximum reserved error code. Method
	// defined errors may be greater than this value.
	ErrorCodeMaxReserved ErrorCode = -32000
)

// IsReserved returns true if c is within the reserved error code range:
//      [LowestReservedErrorCode, HighestReservedErrorCode]
func (c ErrorCode) IsReserved() bool {
	return ErrorCodeMinReserved <= c && c <= ErrorCodeMaxReserved
}

func (c ErrorCode) String() string {
	if !c.IsReserved() {
		return fmt.Sprintf("ErrorCode{%v}", int(c))
	}
	msg := "reserved"
	switch c {
	case ErrorCodeParse:
		msg = ErrorMessageParse
	case ErrorCodeInvalidRequest:
		msg = ErrorMessageInvalidRequest
	case ErrorCodeMethodNotFound:
		msg = ErrorMessageMethodNotFound
	case ErrorCodeInvalidParams:
		msg = ErrorMessageInvalidParams
	case ErrorCodeInternal:
		msg = ErrorMessageInternal
	}
	return fmt.Sprintf("ErrorCode{%v:%q}", int(c), msg)
}

// Error represents a JSON-RPC 2.0 Error object, which is used in the Response
// object. MethodFuncs may return an Error or *Error to return an Error
// Response to the client.
type Error struct {
	// Code is a number that indicates the error type that occurred.
	Code ErrorCode `json:"code"`

	// Message is a short description of the error. The message SHOULD be
	// limited to a concise single sentence.
	Message string `json:"message"`

	// Data is a Primitive or Structured value that contains additional
	// information about the error. This may be omitted. The value of this
	// member is defined by the Server (e.g. detailed error information,
	// nested errors etc.).
	Data interface{} `json:"data,omitempty"`
}

// Error implements the error interface.
func (e Error) Error() string {
	s := fmt.Sprintf("jsonrpc2.Error{Code:%v, Message:%q", e.Code, e.Message)
	if e.Data != nil {
		s += fmt.Sprintf(", Data:%#v", e.Data)
	}
	return s + "}"
}

// IsZero reports whether e is zero valued.
func (e Error) IsZero() bool {
	return e == Error{}
}

// NewError returns an Error with code, msg, and data.
//
// The data must not cause an error when passed to json.Marshal, else any
// MethodFunc returning it will return an Internal Error to clients instead.
//
// If data is type error, then the Error() string is used instead, since
// otherwise the error may not json.Marshal properly.
func NewError(code ErrorCode, msg string, data interface{}) Error {
	if err, ok := data.(error); ok {
		data = err.Error()
	}
	return Error{code, msg, data}
}

// ErrorInvalidParams returns
//      NewError(ErrorCodeInvalidParams, ErrorMessageInvalidParams, data)
//
// MethodFuncs are responsible for detecting invalid parameters and returning
// this error.
func ErrorInvalidParams(data interface{}) Error {
	return NewError(ErrorCodeInvalidParams, ErrorMessageInvalidParams, data)
}

func errorInternal(data interface{}) Error {
	return NewError(ErrorCodeInternal, ErrorMessageInternal, data)
}
func errorParse(data interface{}) Error {
	return NewError(ErrorCodeParse, ErrorMessageParse, data)
}
func errorInvalidRequest(data interface{}) Error {
	return NewError(ErrorCodeInvalidRequest, ErrorMessageInvalidRequest, data)
}
func errorMethodNotFound(data interface{}) Error {
	return NewError(ErrorCodeMethodNotFound, ErrorMessageMethodNotFound, data)
}
