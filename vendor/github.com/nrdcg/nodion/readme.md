# Go library for accessing the Nodion DNS API

[![Build Status](https://github.com/nrdcg/nodion/workflows/Main/badge.svg?branch=master)](https://github.com/nrdcg/nodion/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/nrdcg/nodion)](https://pkg.go.dev/github.com/nrdcg/nodion)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrdcg/nodion)](https://goreportcard.com/report/github.com/nrdcg/nodion)

A [Nodion](https://www.nodion.com) API client written in Go.

nodion is a Go client library for accessing the Nodion DNS API.

## Examples

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nrdcg/nodion"
)

const apiToken = "xxx"

func main() {
	client, err := nodion.NewClient(apiToken)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	zones, err := client.CreateZone(ctx, "example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(zones)
}
```

## API Documentation

- [API docs](https://www.nodion.com/en/docs/dns/api/)
