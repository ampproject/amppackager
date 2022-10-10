# Go library for accessing the Aurora DNS API

[![Build Status](https://github.com/nrdcg/auroradns/workflows/Main/badge.svg?branch=master)](https://github.com/nrdcg/auroradns/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/nrdcg/auroradns)](https://pkg.go.dev/github.com/nrdcg/auroradns)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrdcg/auroradns)](https://goreportcard.com/report/github.com/nrdcg/auroradns)

An Aurora DNS API client written in Go.

auroradns is a Go client library for accessing the Aurora DNS API.

## Available API methods

Zones:
- create
- delete
- list

Records:
- create
- delete
- list

## Example

```go
tr, _ := auroradns.NewTokenTransport("apiKey", "secret")
client, _ := auroradns.NewClient(tr.Client())

zones, _, _ := client.GetZones()

fmt.Println(zones)
```

## API Documentation

- [API docs](https://libcloud.readthedocs.io/en/latest/dns/drivers/auroradns.html#api-docs)
