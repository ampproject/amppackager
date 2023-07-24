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

// Package jsonrpc2 is a complete and strictly conforming implementation of the
// JSON-RPC 2.0 protocol for both clients and servers.
//
// The full specification can be found at https://www.jsonrpc.org.
//
// Clients
//
// The simplest way to make a JSON-RPC 2.0 request is to use the provided
// Client.Request.
//
//      var c jsonrpc2.Client
//      params := []float64{1, 2, 3}
//      var result int
//      err := c.Request(nil, "http://localhost:8080", "sum", params, &result)
//      if _, ok := err.(jsonrpc2.Error); ok {
//      	// received Error Request
//      }
//      if err != nil {
//      	// some JSON marshaling or network error
//      }
//      fmt.Printf("The sum of %v is %v.\n", params, result)
//
// For clients that do not wish to use the provided Client, the Request and
// Response types can be used directly.
//
//      req, _ := json.Marshal(jsonrpc2.Request{Method: "subtract",
//		Params: []int{5, 1},
//		ID:     0,
//	})
//	httpRes, _ := http.Post("www.example.com", "application/json",
//		bytes.NewReader(req))
//	resBytes, _ := ioutil.ReadAll(httpRes.Body)
//	res := jsonrpc2.Response{Result: &MyCustomResultType{}, ID: new(int)}
//	json.Unmarshal(respBytes, &res)
//
// Servers
//
// Servers define their own MethodFuncs and associate them with a name in a
// MethodMap that is passed to HTTPRequestHandler() to return a corresponding
// http.HandlerFunc. See HTTPRequestHandler for more details.
//
//      func getUser(ctx context.Context, params json.RawMessage) interface{} {
//      	var u User
//      	if err := json.Unmarshal(params, &u); err != nil {
//      		return jsonrpc2.InvalidParams(err)
//      	}
//      	conn, err := mydbpkg.GetDBConn()
//      	if err != nil {
//      		// The handler will recover, print debug info if enabled, and
//      		// return an Internal Error to the client.
//      		panic(err)
//      	}
//      	if err := u.Select(conn); err != nil {
//      		return jsonrpc2.NewError(-30000, "user not found", u.ID)
//      	}
//      	return u
//      }
//
//      func StartServer() {
//      	methods := jsonrpc2.MethodMap{"version": versionMethod}
//      	http.ListenAndServe(":8080", jsonrpc2.HTTPRequestHandler(methods,
//                      log.New(os.Stderr, "", 0)))
//      }
package jsonrpc2
