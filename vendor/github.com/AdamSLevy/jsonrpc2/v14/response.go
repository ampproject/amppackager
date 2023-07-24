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

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// version is the valid version string for the "jsonrpc" field required in all
// JSON RPC 2.0 objects.
const version = "2.0"

// Response represents a JSON-RPC 2.0 Response object.
//
// This type is not needed to use the Client or to write MethodFuncs for the
// HTTPRequestHandler.
//
// Response is intended to be used externally to UnmarshalJSON for custom
// clients, and internally to MarshalJSON for the provided HTTPRequestHandler.
//
// To receive a Response, it is recommended to set Result to a pointer to a
// value that the "result" can be unmarshaled into, if known prior to
// unmarshaling. Similarly, it is recommended to set ID to a pointer to a value
// that the "id" can be unmarshaled into, which should be the same type as the
// Request ID.
type Response struct {
	// Result is REQUIRED on success. This member MUST NOT exist if there
	// was an error invoking the method. The value of this member is
	// determined by the method invoked on the Server.
	Result interface{} `json:"result,omitempty"`

	// Error is REQUIRED on error. This member MUST NOT exist if there was
	// no error triggered during invocation. The value for this member MUST
	// be an Object as defined in section 5.1.
	//
	// See Response.HasError.
	Error Error `json:"error,omitempty"`

	// ID is an identifier established by the client that MUST contain a
	// String, Number, or NULL value if included.
	//
	// Since a Request without an ID is not responded to, this member is
	// REQUIRED. It MUST be the same as the value of the id member in the
	// Request Object. If there was an error in detecting the id in the
	// Request Object (e.g. Parse error/Invalid Request), it MUST be Null.
	ID interface{} `json:"id"`
}

// jResponse adds the required "jsonrpc" field and allows for detecting if the
// "error" field is present so that the rule that a Response may not contain
// both an "error and "result" can be enforced.
type jResponse struct {
	// JSONRPC specifies the version of the JSON-RPC protocol. It MUST be
	// exactly "2.0".
	JSONRPC string `json:"jsonrpc"`

	// Error allows UnmarshalJSON to detect if "error" was omitted or set
	// to null which is invalid if "result" is included. Additionally it
	// allows MarshalJSON to explicitly omit it or include it during
	// marhsaling.
	Error *Error `json:"error,omitempty"`

	// *request allows a Request to be used directly while masking its
	// Un/MarshalJSON methods.
	*response
}

// response masks the Response Un/MarshalJSON methods to avoid recursion.
type response Response

// MarshalJSON attempts to marshal r into a valid JSON-RPC 2.0 Response.
//
// If r.HasError(), then "result" is omitted from the JSON, and if r.ID is nil,
// it is set to json.RawMessage("null"). An error is only returned if
// r.Error.Data or r.ID is not marshalable.
//
// If !r.HasError(), then if r.ID is nil, an error is returned. If r.Result is
// nil, it is populated with json.RawMessage("null").
//
// Also, an error is returned if r.Result or r.ID is not marshalable.
func (r Response) MarshalJSON() ([]byte, error) {
	jR := jResponse{
		JSONRPC:  version,
		response: (*response)(&r),
	}
	if r.HasError() {
		jR.Error = &r.Error
		r.Result = nil
		if r.ID == nil {
			r.ID = json.RawMessage("null")
		}
	} else {
		if r.ID == nil {
			return nil, fmt.Errorf("r.ID == nil && !r.HasError()")
		}
		if r.Result == nil {
			r.Result = json.RawMessage("null")
		}
	}
	return json.Marshal(jR)
}

// UnmarshalJSON attempts to unmarshal a JSON-RPC 2.0 Response into r and then
// validates it.
//
// If any fields are unknown other than the application specific fields in the
// "result" object, an error is returned.
//
// If "error" and "result" are both present or not null, a `contains both ...`
// error is returned.
//
// If the "jsonrpc" field is not set to the string "2.0", an `invalid "jsonrpc"
// version: ...` error is returned.
func (r *Response) UnmarshalJSON(data []byte) error {
	// There may be fields in the result not defined in the user provided
	// r.Result, which will cause errors with the json.Decoder below.  So
	// first unmarshal any "result" to a json.RawMessage, which will later
	// be unmarshaled into the userResult.
	userResult := r.Result
	var resultData json.RawMessage
	r.Result = &resultData
	jR := jResponse{Error: &r.Error, response: (*response)(r)}

	// Catch any unknown fields in the top level JSON RPC Response object.
	d := json.NewDecoder(bytes.NewBuffer(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&jR); err != nil {
		return err
	}

	if jR.JSONRPC != version {
		return fmt.Errorf(`invalid "jsonrpc" version: %q`, jR.JSONRPC)
	}

	if r.HasError() {
		if resultData != nil {
			return fmt.Errorf(`contains both "result" and "error"`)
		}
		return nil
	}

	// Restore the userResult and finish unmarshaling.
	r.Result = userResult
	return json.Unmarshal(resultData, &r.Result)
}

// HasError returns true is r.Error has any non-zero values.
func (r Response) HasError() bool {
	return !r.Error.IsZero()
}

// String returns r as a JSON object prefixed with "<-- " to indicate an
// incoming Response.
//
// If r.MarshalJSON returns an error then the error string is returned with
// some context.
func (r Response) String() string {
	b, err := r.MarshalJSON()
	if err != nil {
		return fmt.Sprintf("%#v.MarshalJSON(): %v", r, err)
	}
	return "<-- " + string(b)
}

// BatchResponse is a type that implements fmt.Stringer for a slice of
// Responses.
type BatchResponse []Response

// String returns br as a JSON array prefixed with "<-- " to indicate an
// incoming BatchResponse and with newlines separating the elements of br.
func (br BatchResponse) String() string {
	s := "<-- [\n"
	for i, res := range br {
		s += "  " + res.String()[4:]
		if i < len(br)-1 {
			s += ","
		}
		s += "\n"
	}
	return s + "]"
}
