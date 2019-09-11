// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package certfetcher

import (
	"crypto"
	"crypto/x509"
	"strconv"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/go-acme/lego/v3/certcrypto"
	"github.com/go-acme/lego/v3/certificate"
	"github.com/go-acme/lego/v3/challenge/http01"
	"github.com/go-acme/lego/v3/lego"
	"github.com/go-acme/lego/v3/registration"
	"github.com/pkg/errors"
)

type CertFetcher struct {
	// ACME Discovery URL
	AcmeDiscoURL string
	AcmeUser     MyUser
	// Domains to validate
	Domains    []string
	legoClient *lego.Client
}

// You'll need a user or account type that implements acme.User
type MyUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *MyUser) GetEmail() string {
	return u.Email
}
func (u MyUser) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *MyUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

// Initializes the cert fetcher with information it needs to fetch new certificates in the future.
func NewFetcher(email string, privateKey crypto.PrivateKey, acmeDiscoURL string,
	domains []string, acmeChallengePort int, shouldRegister bool) (*CertFetcher, error) {
	myUser := MyUser{
		Email: email,
		key:   privateKey,
	}
	config := lego.NewConfig(&myUser)

	config.CADirURL = acmeDiscoURL
	config.Certificate.KeyType = certcrypto.RSA2048

	// A client facilitates communication with the CA server.
	client, err := lego.NewClient(config)
	if err != nil {
		return nil, err
	}

	// We specify an http port of `acmeChallengePort`
	// because we aren't running as root and can't bind a listener to port 80 and 443
	// (used later when we attempt to pass challenges). Keep in mind that you still
	// need to proxy challenge traffic to port `acmeChallengePort`.
	err = client.Challenge.SetHTTP01Provider(
		http01.NewProviderServer("", strconv.Itoa(acmeChallengePort)))
	if err != nil {
		return nil, err
	}

	// Theoretically, this should always be set to false as users should have pre-registered for access
	// to the ACME CA and agreed to the TOS.
	// TODO(banaag): revisit this when trying the class out with Digicert CA.
	if !shouldRegister {
		myUser.Registration = new(registration.Resource)
	} else {
		reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
		if err != nil {
			return nil, err
		}
		myUser.Registration = reg
	}

	return &CertFetcher{
		AcmeDiscoURL: acmeDiscoURL,
		AcmeUser:     myUser,
		Domains:      domains,
		legoClient:   client,
	}, nil
}

func (f *CertFetcher) FetchNewCert() ([]*x509.Certificate, error) {
	request := certificate.ObtainRequest{
		Domains: f.Domains,
		Bundle:  true,
	}

	// Each resource comes back with the cert bytes, the bytes of the client's
	// private key, and a certificate URL.
	resource, err := f.legoClient.Certificate.Obtain(request)
	if err != nil {
		return nil, err
	}

	if resource == nil {
		return nil, errors.New("No resource returned.")
	}

	if resource.Certificate == nil {
		return nil, errors.New("No certificates were returned.")
	}

	cert, err := signedexchange.ParseCertificates(resource.Certificate)
	if err != nil {
		return nil, err
	}

	return cert, err
}
