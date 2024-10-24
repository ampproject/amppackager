# Go library for accessing the Mail-in-a-Box API

[![Build Status](https://github.com/nrdcg/mailinabox/actions/workflows/main.yml/badge.svg)](https://github.com/nrdcg/mailinabox/actions/workflows/main.yml)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/nrdcg/mailinabox)](https://pkg.go.dev/github.com/nrdcg/mailinabox)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrdcg/mailinabox)](https://goreportcard.com/report/github.com/nrdcg/mailinabox)

A Mail-in-a-Box API client written in Go.

`mailinabox` is a Go client library for accessing the Mail-in-a-Box API.

## Examples

```go
package main

import (
	"context"
	"fmt"

	"github.com/nrdcg/mailinabox"
)

func main() {
	client, err := mailinabox.NewClient("https://example.com", "user@example.com", "secret")
	if err != nil {
		panic(err)
	}

	record := mailinabox.Record{
		Name:  "example.com",
		Type:  "A",
		Value: "10.0.0.1",
	}

	resp, err := client.DNS.AddRecord(context.Background(), record)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
```

```go
package mailinabox_test

import (
	"context"
	"fmt"

	"github.com/nrdcg/mailinabox"
)

func main() {
	client, err := mailinabox.NewClient("https://example.com", "user@example.com", "secret")
	if err != nil {
		panic(err)
	}

	resp, err := client.System.Reboot(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
```


## API Documentation

- [API docs](https://mailinabox.email/api-docs.html)

## Supported APIs

- [x] User API
- [x] Mail API
- [x] DNS API
- [ ] SSL API
- [ ] Web API
- [ ] MFA API
- [x] System API
