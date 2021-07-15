# Go library for accessing the deSEC API

[![GoDoc](https://godoc.org/github.com/nrdcg/desec?status.svg)](https://godoc.org/github.com/nrdcg/desec)
[![Build Status](https://travis-ci.com/nrdcg/desec.svg?branch=master)](https://travis-ci.com/nrdcg/desec)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrdcg/desec)](https://goreportcard.com/report/github.com/nrdcg/desec)

An deSEC API client written in Go.

desec is a Go client library for accessing the deSEC API.

## Examples

```go
package main

import (
	"fmt"

	"github.com/nrdcg/desec"
)

func main() {
	client := desec.NewClient("token")

	newDomain, err := client.Domains.Create("example.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(newDomain)
}
```

```go
package main

import (
	"fmt"

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

	err := client.Account.Register(registration)
	if err != nil {
		panic(err)
	}
}
```

```go
package main

import (
	"fmt"

	"github.com/nrdcg/desec"
)

func main() {
	client := desec.NewClient("")

	_, err := client.Account.Login("email@example.com", "secret")
	if err != nil {
		panic(err)
	}

	domains, err := client.Domains.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(domains)

	err = client.Account.Logout()
	if err != nil {
		panic(err)
	}
}
```

## API Documentation

- [API docs](https://desec.readthedocs.io/en/latest/)
- [API endpoint reference](https://desec.readthedocs.io/en/latest/endpoint-reference.html)
