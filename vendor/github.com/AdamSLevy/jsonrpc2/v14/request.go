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

// Request represents a JSON-RPC 2.0 Request or Notification object.
//
// This type is not needed to use the Client or to write MethodFuncs for the
// HTTPRequestHandler.
//
// Request is intended to be used externally to MarshalJSON for custom clients,
// and internally to UnmarshalJSON for the HTTPRequestHandler.
//
// To make a Request, you must populate ID. If ID is empty, the Request is
// treated as a Notification and does not receive a Response. See MarshalJSON
// for more details.
type Request struct {
	// Method is a string containing the name of the method to be invoked.
	Method string `json:"method"`

	// Params is a structured value that holds the parameter values to be
	// used during the invocation of the method.
	//
	// This member MAY be omitted.
	Params interface{} `json:"params,omitempty"`

	// ID is an identifier established by the Client that MUST contain a
	// String, Number, or NULL value if included.
	//
	// If it is not included it is assumed to be a notification. The value
	// SHOULD normally not be Null and Numbers SHOULD NOT contain
	// fractional parts.
	ID interface{} `json:"id,omitempty"`
}

// jRequest adds the required "jsonrpc" field and allows for detecting if the
// "method" field is omitted or null, since an empty method name is not
// explicitly prohibited in the spec, but a missing or null "method" is.
// Additionally we can distinguish between the "id" field being omitted,
// indicating a Notification, and it being null, which is still technically a
// Request.
type jRequest struct {
	// JSONRPC specifies the version of the JSON-RPC protocol. It MUST be
	// exactly "2.0".
	JSONRPC string `json:"jsonrpc"`

	// Method allows UnmarshalJSON to detect if "method" was omitted or
	// null without bothering users with a pointer.
	Method *string `json:"method"`

	// *request allows a Request to be used directly while masking its
	// Un/MarshalJSON methods.
	*request

	// ID allow UnmarshalJSON to distinguish between a missing ID,
	// indicating a Notification, or a null ID, which is technically
	// allowed for Requests, but not recommended.
	ID json.RawMessage `json:"id,omitempty"`
}

// request masks the Request Un/MarshalJSON methods to avoid recursion.
type request Request

// MarshalJSON attempts to marshal r into a valid JSON-RPC 2.0 Request or
// Notification object.
//
// If r.ID is nil, then the returned data represents a Notification.
//
// If r.ID is not nil, then the returned data represents a Request. Also, an
// `invalid "id": ...` error is returned if the r.ID does not marshal into a
// valid JSON number, string, or null. Although technically permitted, it is
// not recommended to use json.RawMessage("null") as an ID, as this is used by
// Responses when there is an error parsing "id".
//
// If r.Params is not nil, then an `invalid "params": ...` error is returned if
// it does not marshal into a valid JSON object, array, or null.
//
// An empty Method, though not recommended, is technically valid and does not
// cause an error.
func (r Request) MarshalJSON() ([]byte, error) {
	jR := jRequest{
		JSONRPC: version,
		Method:  &r.Method,
		request: (*request)(&r),
	}
	if r.ID != nil {
		id, err := json.Marshal(r.ID)
		if err != nil {
			return nil, fmt.Errorf(`invalid "id": %w`, err)
		}
		if err := validateID(id); err != nil {
			return nil, err
		}
		jR.ID = id
	}
	if r.Params != nil {
		params, err := json.Marshal(r.Params)
		if err != nil {
			return nil, fmt.Errorf(`invalid "params": %w`, err)
		}
		if err := validateParams(params); err != nil {
			return nil, err
		}
		r.Params = json.RawMessage(params)
	}
	return json.Marshal(jR)
}

// UnmarshalJSON attempts to unmarshal a JSON-RPC 2.0 Request or Notification
// into r and then validates it.
//
// To allow for precise JSON type validation and to avoid unneeded unmarshaling
// by the HTTPRequestHandler, "id" and "params" are both unmarshaled into
// json.RawMessage. After a successful call, r.ID and r.Params are set to a
// json.RawMessage that is either nil or contains the raw JSON for the
// respective field. So, if no error is returned, this is guaranteed to not
// panic:
//      id, params := r.ID.(json.RawMessage), r.Params.(json.RawMessage)
//
// If "id" is omitted, then r.ID is set to json.RawMessage(nil). If "id" is
// null, then r.ID is set to json.RawMessage("null").
//
// If "params" is null or omitted, then r.Params is set to
// json.RawMessage(nil).
//
// If any fields are unknown, an error is returned.
//
// If the "jsonrpc" field is not set to the string "2.0", an `invalid "jsonrpc"
// version: ...` error is be returned.
//
// If the "method" field is omitted or null, a `missing "method"` error is
// returned. An explicitly empty "method" string does not cause an error.
//
// If the "id" value is not a JSON number, string, or null, an `invalid "id":
// ...` error is returned.
//
// If the "params" value is not a JSON array, object, or null, an `invalid
// "params": ...` error is returned.
func (r *Request) UnmarshalJSON(data []byte) error {
	// params stores the "params" JSON if it is not omitted or null.
	var params json.RawMessage
	r.Params = &params
	jR := jRequest{request: (*request)(r)}

	d := json.NewDecoder(bytes.NewBuffer(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&jR); err != nil {
		return err
	}

	if jR.JSONRPC != version {
		return fmt.Errorf(`invalid "jsonrpc" version: %q`, jR.JSONRPC)
	}

	if jR.Method == nil {
		return fmt.Errorf(`missing "method"`)
	}
	r.Method = *jR.Method

	if jR.ID != nil {
		if err := validateID(jR.ID); err != nil {
			return err
		}
	}
	r.ID = jR.ID

	if params != nil {
		if err := validateParams(params); err != nil {
			return err
		}
	}
	r.Params = params

	return nil
}

// String returns r as a JSON object prefixed with "--> " to indicate an
// outgoing Request.
//
// If r.MarshalJSON returns an error then the error string is returned with
// some context.
func (r Request) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("%#v.MarshalJSON(): %v", r, err)
	}
	return "--> " + string(b)
}

// BatchRequest is a type that implements fmt.Stringer for a slice of Requests.
type BatchRequest []Request

// String returns br as a JSON array prefixed with "--> " to indicate an
// outgoing BatchRequest and with newlines separating the elements of br.
func (br BatchRequest) String() string {
	s := "--> [\n"
	for i, res := range br {
		s += "  " + res.String()[4:]
		if i < len(br)-1 {
			s += ","
		}
		s += "\n"
	}
	return s + "]"
}
