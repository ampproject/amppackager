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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
)

// DebugMethodFunc controls whether additional debug information is printed to
// stdout in the event of an InternalError when a MethodFunc is called.
//
// This can be helpful when troubleshooting panics or Internal Errors from a
// MethodFunc.
var DebugMethodFunc = false

// MethodMap associates method names with MethodFuncs and is passed to
// HTTPRequestHandler to generate a corresponding http.HandlerFunc.
//
// Method names that begin with the word rpc followed by a period character
// (U+002E or ASCII 46) are reserved for rpc-internal methods and extensions
// and MUST NOT be used for anything else. If such a method name is detected
// this will panic. No such internal rpc methods are yet defined in this
// implementation.
type MethodMap map[string]MethodFunc

// MethodFunc is the function signature used for RPC methods.
//
// MethodFuncs are invoked by the HTTPRequestHandler when a valid Request is
// received. MethodFuncs do not need to concern themselves with the details of
// JSON-RPC 2.0 outside of the "params" field, as all parsing and validation is
// handled by the handler.
//
// The handler will call a MethodFunc with ctx set to the corresponding
// http.Request.Context() and params set to the JSON data from the "params"
// field of the Request. If "params" was omitted or null, params will be nil.
// Otherwise, params is guaranteed to be valid JSON that represents a JSON
// Object or Array.
//
// A MethodFunc is responsible for application specific parsing of params.  A
// MethodFunc should return an ErrorInvalidParams if there is any issue parsing
// expected parameters.
//
// To return a success Response to the client a MethodFunc must return a
// non-error value, that will not cause an error when passed to json.Marshal,
// to be used as the Response.Result. Any marshaling error will cause a panic
// and an Internal Error will be returned to the client.
//
// To return an Error Response to the client, a MethodFunc must return a valid
// Error. A valid Error must use ErrorCodeInvalidParams or any ErrorCode
// outside of the reserved range, and the Error.Data must not cause an error
// when passed to json.Marshal. If the Error is not valid, a panic will occur
// and an Internal Error will be returned to the client.
//
// If a MethodFunc panics or returns any other error, an Internal Error is
// returned to the client. If the returned error is anything other than
// context.Canceled or context.DeadlineExceeded, a panic will occur.
//
// For additional debug output from a MethodFunc regarding the cause of an
// Internal Error, set DebugMethodFunc to true. Information about the method
// call and a stack trace will be printed on panics.
type MethodFunc func(ctx context.Context, params json.RawMessage) interface{}

// call is used to safely call a method from within an http.HandlerFunc. call
// wraps the actual invocation of the method so that it can recover from panics
// and validate and sanitize the returned Response. If the method panics or
// returns an invalid Response, an Internal Error is returned.
func (method MethodFunc) call(ctx context.Context,
	name string, params json.RawMessage, lgr Logger) (res Response) {

	var result interface{}
	defer func() {
		if err := recover(); err != nil {
			res.Error = errorInternal(nil)
			if DebugMethodFunc {
				//res.Data = err
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				lgr.Printf("jsonrpc2: panic running method %q: %v\n",
					name, err)
				lgr.Printf("jsonrpc2: Params: %v\n", string(params))
				lgr.Printf("jsonrpc2: Return: %#v\n", result)
				lgr.Println(string(buf))
			}
		}
	}()
	result = method(ctx, params)
	if err, ok := result.(error); ok {
		var methodErr Error
		if errors.As(err, &methodErr) {
			// InvalidParamsCode is the only reserved ErrorCode
			// MethodFuncs are allowed to return.
			if methodErr.Code == ErrorCodeInvalidParams {
				if methodErr.Message == "" {
					// Ensure the correct message is used if none is supplied.
					methodErr.Message = ErrorMessageInvalidParams
				}
			} else if methodErr.Code.IsReserved() {
				panic(fmt.Errorf("invalid use of %v", methodErr.Code))
			}
			if methodErr.Data != nil {
				// MethodFuncs could return something that
				// cannot be marshaled. Catch that here.
				data, err := json.Marshal(methodErr.Data)
				if err != nil {
					panic(fmt.Errorf("json.Marshal(Error.Data): %w", err))
				}
				methodErr.Data = json.RawMessage(data)
			}
			res.Error = methodErr
			return
		}

		// MethodFuncs should not normally return a generic error
		// unless they are returning an error that is, or that wraps,
		// context.Canceled or context.DeadlineExceeded.
		//
		// If the http.Request.Context() is canceled then this will
		// never get returned anyway.
		if errors.Is(err, context.Canceled) ||
			errors.Is(err, context.DeadlineExceeded) {
			res.Error = errorInternal(err)
			return
		}

		// Otherwise, if a MethodFunc intends to return an error to the
		// client it must use the Error type, so this is a program
		// integrity error that should be reported as a panic.
		panic(fmt.Errorf("unexpected error: %w", err))
	}

	// MethodFuncs could return something that cannot be marshaled. Catch
	// that here.
	data, err := json.Marshal(result)
	if err != nil {
		panic(fmt.Errorf("json.Marshal(result): %w", err))
	}
	res.Result = json.RawMessage(data)
	return
}
