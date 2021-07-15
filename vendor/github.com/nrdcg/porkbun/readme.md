# Go library for accessing the Porkdun API

[![Build Status](https://github.com/nrdcg/porkbun/workflows/Main/badge.svg?branch=master)](https://github.com/nrdcg/porkbun/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/nrdcg/porkbun)](https://pkg.go.dev/github.com/nrdcg/porkbun)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrdcg/porkbun)](https://goreportcard.com/report/github.com/nrdcg/porkbun)

An [Porkbun](https://porkbun.com) API client written in Go.

porkbun is a Go client library for accessing the Porkbun API.

## Examples

```go
package main

import (
	"context"
	"fmt"

	"github.com/nrdcg/porkbun"
)

func main() {
	client := porkbun.New("secret", "key")

	ctx := context.Background()

	yourIP, err := client.Ping(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(yourIP)
}
```

## API Documentation

- [API docs](https://porkbun.com/api/json/v3/documentation)
