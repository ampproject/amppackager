# INWX Go API client

[![Build Status](https://travis-ci.com/nrdcg/goinwx.svg?branch=master)](https://travis-ci.com/nrdcg/goinwx)
[![GoDoc](https://godoc.org/github.com/nrdcg/goinwx?status.svg)](https://godoc.org/github.com/nrdcg/goinwx)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrdcg/goinwx)](https://goreportcard.com/report/github.com/nrdcg/goinwx)

This go library implements some parts of the official INWX XML-RPC API.

## API

```go
package main

import (
	"log"

	"github.com/nrdcg/goinwx"
)

func main() {
	client := goinwx.NewClient("username", "password", &goinwx.ClientOptions{Sandbox: true})

	_, err := client.Account.Login()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Account.Logout(); err != nil {
			log.Printf("inwx: failed to logout: %v", err)
		}
	}()

	var request = &goinwx.NameserverRecordRequest{
		Domain:  "domain.com",
		Name:    "foo.domain.com.",
		Type:    "TXT",
		Content: "aaa",
		TTL:     300,
	}

	_, err = client.Nameservers.CreateRecord(request)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Using 2FA

If it is desired to use 2FA without manual entering the TOTP every time,
you must set the parameter `otp-key` to the secret that is shown during the setup of 2FA for you INWX account.
Otherwise, you can skip `totp.GenerateCode` step and enter the verification code of the Google Authenticator app every time manually.

The `otp-key` looks something like `EELTWFL55ESIHPTJAAHBCY7LXBZARUOJ`.

```go
package main

import (
	"log"
	"time"

	"github.com/nrdcg/goinwx"
	"github.com/pquerna/otp/totp"
)

func main() {
	client := goinwx.NewClient("username", "password", &goinwx.ClientOptions{Sandbox: true})

	resp, err := client.Account.Login()
	if err != nil {
		log.Fatal(err)
	}

	if resp.TFA != "GOOGLE-AUTH" {
		log.Fatal("unsupported 2 Factor Authentication")
	}

	tan, err := totp.GenerateCode("otp-key", time.Now())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Account.Unlock(tan)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Account.Logout(); err != nil {
			log.Printf("inwx: failed to logout: %v", err)
		}
	}()

	request := &goinwx.NameserverRecordRequest{
		Domain:  "domain.com",
		Name:    "foo.domain.com.",
		Type:    "TXT",
		Content: "aaa",
		TTL:     300,
	}

	_, err = client.Nameservers.CreateRecord(request)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Supported Features

Full API documentation can be found [here](https://www.inwx.de/en/help/apidoc).

The following parts are implemented:

* Account
  * Login
  * Logout
  * Lock
  * Unlock (with mobile TAN)
* Domains
  * Check
  * Register
  * Delete
  * Info
  * GetPrices
  * List
  * Whois
  * Update
* Nameservers
  * Check
  * Create
  * Info
  * List
  * CreateRecord
  * UpdateRecord
  * DeleteRecord
  * FindRecordById
* Contacts
  * List 
  * Info
  * Create
  * Update
  * Delete

## Contributions

Your contributions are very appreciated.
