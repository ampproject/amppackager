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
	"encoding/json"
	"fmt"
)

// validateID assumes that id is valid JSON and returns true if len(id) == 0,
// or id represents a JSON Number, String or Null.
//
// From the JSONRPC 2.0 Spec:
//
// id
//      An identifier established by the Client that MUST contain a String,
//      Number, or NULL value if included.
//
// Below is the JSON grammar for String, Number and Null from JSON.org.
//      value
//          string
//          number
//          "null"
//
//      string
//          '"' characters '"'
//
//      number
//          integer fraction exponent
//
//      integer
//          digit
//          onenine digits
//          '-' digit
//          '-' onenine digits
//
//      digits
//          digit
//          digit digits
//
//      digit
//          '0'
//          onenine
//
//      onenine
//          '1' . '9'
//
//      fraction
//          ""
//          '.' digits
//
//      exponent
//          ""
//          'E' sign digits
//          'e' sign digits
//
// Thus if we know that we are working with valid JSON, we can determine the
// JSON type by the first byte alone.
func validateID(id json.RawMessage) error {
	if len(id) == 0 {
		return fmt.Errorf(`invalid "id": empty`)
	}
	b := id[0]
	if b == 'n' || // null
		b == '"' || // string
		isNumber(b) {
		return nil
	}
	return fmt.Errorf(`invalid "id": not a number, string, or null`)
}

func isNumber(b byte) bool {
	if b == '-' || ('0' <= b && b <= '9') || // integer
		b == 'E' || b == 'e' || // exponent
		b == '.' { // fraction
		return true
	}
	return false
}

// validateParams assumes that params is valid JSON and returns true if params is
// nil, or if it represents a structured value (Array or Object), or Null.
//
// From the JSONRPC 2.0 Spec:
//
// params
//      A Structured value that holds the parameter values to be used during
//      the invocation of the method. This member MAY be omitted.
//
// Below is the JSON grammar for Object, Array and Null from JSON.org.
//      value
//          object
//          array
//          "null"
//
//      object
//          '{' ws '}'
//          '{' members '}'
//
//      array
//          '[' ws ']'
//          '[' elements ']'
//
// Thus if we know that we are working with valid JSON, we can determine the
// JSON type by the first byte alone.
func validateParams(params json.RawMessage) error {
	if len(params) == 0 {
		return fmt.Errorf(`invalid "params": empty`)
	}
	b := params[0]
	if b == 'n' || // null
		b == '[' || // array
		b == '{' { // object
		return nil
	}
	return fmt.Errorf(`invalid "params": not an object, array, or null`)
}
