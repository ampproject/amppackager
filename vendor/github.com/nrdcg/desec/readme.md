# Go library for accessing the deSEC API

[![Build Status](https://github.com/nrdcg/desec/workflows/Main/badge.svg?branch=master)](https://github.com/nrdcg/desec/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/nrdcg/desec)](https://pkg.go.dev/github.com/nrdcg/desec)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrdcg/desec)](https://goreportcard.com/report/github.com/nrdcg/desec)

An deSEC API client written in Go.

desec is a Go client library for accessing the deSEC API.

## Examples

```go
package main

import (
	"context"
	"fmt"

	"github.com/nrdcg/desec"
)

func main() {
	client := desec.NewClient("token")

	newDomain, err := client.Domains.Create(context.Background(), "example.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(newDomain)
}
```

```go
package main

import (
	"context"

	"github.com/nrdcg/desec"
)

func main() {
	client := desec.NewClient("")
	registration := desec.Registration{
		Email:    "email@example.com",
		Password: "secret",
		Captcha: &desec.Captcha{
			ID:       "00010203-0405-0607-0809-0a0b0c0d0e0f",
			Solution: "12H45",
		},
	}

	err := client.Account.Register(context.Background(), registration)
	if err != nil {
		panic(err)
	}
}
```

```go
package main

import (
	"context"
	"fmt"

	"github.com/nrdcg/desec"
)

func main() {
	client := desec.NewClient("")

	_, err := client.Account.Login(context.Background(), "email@example.com", "secret")
	if err != nil {
		panic(err)
	}

	domains, err := client.Domains.GetAll(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(domains)

	err = client.Account.Logout(context.Background())
	if err != nil {
		panic(err)
	}
}
```

## API Documentation

- [API docs](https://desec.readthedocs.io/en/latest/)
- [API endpoint reference](https://desec.readthedocs.io/en/latest/endpoint-reference.html)
