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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// HTTPRequestHandler returns an http.HandlerFunc for the given methods.
//
// The returned http.HandlerFunc efficiently handles any conforming single or
// batch requests or notifications, accurately catches all defined protocol
// errors, calls the appropriate MethodFuncs, recovers from any panics or
// invalid return values, and returns an Internal Error or Response with the
// correct ID, if not a Notification.
//
// See MethodFunc for more details.
//
// It is not safe to modify methods while the returned http.HandlerFunc is in
// use.
//
// This will panic if a method name beginning with "rpc." is used. See
// MethodMap for more details.
//
// The handler will use lgr to log any errors and debug information, if
// DebugMethodFunc is true. If lgr is nil, the default Logger from the log
// package is used.
func HTTPRequestHandler(methods MethodMap, lgr Logger) http.HandlerFunc {
	for name := range methods {
		if strings.HasPrefix(name, "rpc.") {
			panic(fmt.Errorf("invalid method name: %v", name))
		}
	}
	if lgr == nil {
		lgr = log.New(os.Stderr, "", log.LstdFlags)
	}

	return func(w http.ResponseWriter, req *http.Request) {
		res := handle(methods, req, lgr)
		if req.Context().Err() != nil || res == nil {
			return
		}
		// We should never have a JSON encoding related error because
		// MethodFunc.call() already Marshaled any user provided Data
		// or Result, and everything else is marshalable.
		//
		// However an error can be returned related to w.Write, which
		// there is nothing we can do about, so we just log it here.
		enc := json.NewEncoder(w)
		if err := enc.Encode(res); err != nil {
			lgr.Printf("req.Body.Write(): %v", err)
		}
	}
}

// handle an http.Request for the given methods.
func handle(methods MethodMap, req *http.Request, lgr Logger) interface{} {
	// Read all bytes of HTTP request body.
	reqBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return Response{Error: errorInternal(err.Error())}
	}

	// Ensure valid JSON so it can be assumed going forward.
	if !json.Valid(reqBytes) {
		return Response{Error: errorParse(nil)}
	}

	// Attempt to unmarshal into a slice to detect a batch request. Use
	// []json.RawMessage so that each Request can be unmarshaled
	// individually.
	batch := true
	rawReqs := make([]json.RawMessage, 1)
	if json.Unmarshal(reqBytes, &rawReqs) != nil {
		// Since the JSON is valid, this Unmarshal error indicates that
		// this is just a single request.
		batch = false
		rawReqs[0] = json.RawMessage(reqBytes)
	}

	// Catch empty batch requests.
	if len(rawReqs) == 0 {
		return Response{Error: errorInvalidRequest("empty batch request")}
	}

	// Process each Request, omitting any returned Response that is empty.
	responses := make(BatchResponse, 0, len(rawReqs))
	for _, rawReq := range rawReqs {
		if req.Context().Err() != nil {
			return nil
		}
		res := processRequest(req.Context(), methods, rawReq, lgr)
		if res == (Response{}) {
			// Don't respond to Notifications.
			continue
		}
		responses = append(responses, res)
	}

	// Send nothing if there are no responses.
	if len(responses) == 0 {
		return nil
	}

	// Return the BatchResponse if this was a batch request.
	if batch {
		return responses
	}

	// Return a single Response.
	return responses[0]
}

// processRequest unmarshals and processes a single Request stored in rawReq
// using the methods defined in methods. If res is zero valued, then the
// Request was a Notification and should not be responded to.
func processRequest(ctx context.Context,
	methods MethodMap, rawReq json.RawMessage, lgr Logger) (res Response) {

	// Unmarshal into req with an error on any unknown fields.
	var req Request
	if err := json.Unmarshal(rawReq, &req); err != nil {
		// At this point we know that this was valid JSON, so this is a
		// not a ParseError, but something about the Request object did
		// not conform to spec, so we return an invalidRequest Error.
		//
		// At this point we have no way to know if this was a Request
		// or Notification, so we must respond with "id" set to null.
		return Response{Error: errorInvalidRequest(err.Error())}
	}

	// Use a type assertion to get req.ID and req.Params as
	// json.RawMessage. See Request.UnmarshalJSON for details about why
	// this type assertion is safe here.
	id, params := req.ID.(json.RawMessage), req.Params.(json.RawMessage)

	// Never respond to Notifications, even if an Error occurs. For
	// Requests, always use the Request ID in the Response.
	defer func() {
		if id == nil {
			res = Response{}
			return
		}
		res.ID = id
	}()

	// Look up the requested method and call it if found.
	method, ok := methods[req.Method]
	if !ok {
		return Response{Error: errorMethodNotFound(req.Method)}
	}
	res = method.call(ctx, req.Method, params, lgr)

	// Log the method name if debugging is enabled and the method had an
	// internal error.
	if DebugMethodFunc && res.HasError() && res.Error.Code == ErrorCodeInternal {
		lgr.Printf("Method: %#v\n\n", req.Method)
	}

	return res
}
