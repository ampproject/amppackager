# go-lwApi
LiquidWeb API Golang client

[![GoDoc](https://godoc.org/github.com/liquidweb/go-lwApi?status.svg)](https://godoc.org/github.com/liquidweb/go-lwApi)

## Setting up Authentication
When creating an api client, it expects to be configured via a configuration struct. Here is an example of how to get an api client.

```
package main

import (
	"fmt"

	lwApi "github.com/liquidweb/go-lwApi"
)

func main() {
	config := lwApi.LWAPIConfig{
		Username: "ExampleUsername",
		Password: "ExamplePassword",
		Url:      "api.liquidweb.com",
	}
	apiClient, iErr := lwApi.New(&config)
}
```
## Importing
``` go
import (
        lwApi "github.com/liquidweb/go-lwApi"
)
```
## Calling a method
``` go
apiClient, iErr := lwApi.New(&config)
if iErr != nil {
  panic(iErr)
}
args := map[string]interface{}{
  "uniq_id": "2UPHPL",
}
got, gotErr := apiClient.Call("bleed/asset/details", args)
if gotErr != nil {
  panic(gotErr)
}
fmt.Printf("RETURNED:\n\n%+v\n\n", got)
```

As you can see, you don't need to prefix the `params` key, as that is handled in the `Call()` function for you.
