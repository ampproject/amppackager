# github.com/AdamSLevy/jsonrpc2/v14
[![GoDoc](https://godoc.org/github.com/AdamSLevy/jsonrpc2?status.svg)](https://godoc.org/github.com/AdamSLevy/jsonrpc2)
[![Go Report Card](https://goreportcard.com/badge/github.com/AdamSLevy/jsonrpc2)](https://goreportcard.com/report/github.com/AdamSLevy/jsonrpc2)
[![Coverage Status](https://coveralls.io/repos/github/AdamSLevy/jsonrpc2/badge.svg?branch=master)](https://coveralls.io/github/AdamSLevy/jsonrpc2?branch=master)
[![Build Status](https://travis-ci.org/AdamSLevy/jsonrpc2.svg?branch=master)](https://travis-ci.org/AdamSLevy/jsonrpc2)

Package jsonrpc2 is a complete and strictly conforming implementation of the
JSON-RPC 2.0 protocol for both clients and servers.

The full specification can be found at https://www.jsonrpc.org.

## Clients

The simplest way to make a JSON-RPC 2.0 request is to use the provided
Client.Request.
```golang
     var c jsonrpc2.Client
     params := []float64{1, 2, 3}
     var result int
     err := c.Request(nil, "http://localhost:8080", "sum", params, &result)
     if _, ok := err.(jsonrpc2.Error); ok {
     	// received Error Request
     }
     if err != nil {
     	// some JSON marshaling or network error
     }
     fmt.Printf("The sum of %v is %v.\n", params, result)
```

For clients that do not wish to use the provided Client, the Request and
Response types can be used directly.

```golang
     req, _ := json.Marshal(jsonrpc2.Request{Method: "subtract",
       	Params: []int{5, 1},
       	ID:     0,
       })
       httpRes, _ := http.Post("www.example.com", "application/json",
       	bytes.NewReader(req))
       resBytes, _ := ioutil.ReadAll(httpRes.Body)
       res := jsonrpc2.Response{Result: &MyCustomResultType{}, ID: new(int)}
       json.Unmarshal(respBytes, &res)
```

## Servers

Servers define their own MethodFuncs and associate them with a name in a
MethodMap that is passed to HTTPRequestHandler() to return a corresponding
http.HandlerFunc. See HTTPRequestHandler for more details.
```golang
     func getUser(ctx context.Context, params json.RawMessage) interface{} {
     	var u User
     	if err := json.Unmarshal(params, &u); err != nil {
     		return jsonrpc2.InvalidParams(err)
     	}
     	conn, err := mydbpkg.GetDBConn()
     	if err != nil {
     		// The handler will recover, print debug info if enabled, and
     		// return an Internal Error to the client.
     		panic(err)
     	}
     	if err := u.Select(conn); err != nil {
     		return jsonrpc2.NewError(-30000, "user not found", u.ID)
     	}
     	return u
     }

     func StartServer() {
     	methods := jsonrpc2.MethodMap{"version": versionMethod}
     	http.ListenAndServe(":8080", jsonrpc2.HTTPRequestHandler(methods))
     }
```
