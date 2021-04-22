// Copyright 2018 Google LLC
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

package testing

import (
	"crypto"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/ampproject/amppackager/packager/util"
)

// A cert (with its issuer chain) for testing.
var Certs = func() []*x509.Certificate {
	certPem, _ := ioutil.ReadFile("../../testdata/b3/fullchain.cert")
	certs, _ := signedexchange.ParseCertificates(certPem)
	return certs
}()

// Its corresponding private key.
var Key = func() crypto.PrivateKey {
	keyPem, _ := ioutil.ReadFile("../../testdata/b3/server.privkey")
	// This call to ParsePrivateKey() is needed by util_test.go.
	key, _ := util.ParsePrivateKey(keyPem)
	return key
}()

// 90 days cert of amppackageexample.com and www.amppackageexample.com in SAN
var B3Certs = func() []*x509.Certificate {
	certPem, _ := ioutil.ReadFile("../../testdata/b3/fullchain.cert")
	certs, _ := signedexchange.ParseCertificates(certPem)
	return certs
}()

// Private key of B3Certs
var B3Key = func() crypto.PrivateKey {
	keyPem, _ := ioutil.ReadFile("../../testdata/b3/server.privkey")
	key, _ := util.ParsePrivateKey(keyPem)
	return key
}()

//  90 days cert of amppackageexample2.com and www.amppackageexample2.com in SAN
var B3Certs2 = func() []*x509.Certificate {
	certPem, _ := ioutil.ReadFile("../../testdata/b3/fullchain2.cert")
	certs, _ := signedexchange.ParseCertificates(certPem)
	return certs
}()

// Private key of B3Certs2
var B3Key2 = func() crypto.PrivateKey {
	keyPem, _ := ioutil.ReadFile("../../testdata/b3/server2.privkey")
	key, _ := util.ParsePrivateKey(keyPem)
	return key
}()

// 91 days cert from B3Key
var B3Certs91Days = func() []*x509.Certificate {
	certPem, _ := ioutil.ReadFile("../../testdata/b3/fullchain_91days.cert")
	certs, _ := signedexchange.ParseCertificates(certPem)
	return certs
}()

// secp521r1 private key
var B3KeyP521 = func() crypto.PrivateKey {
	keyPem, _ := ioutil.ReadFile("../../testdata/b3/server_p521.privkey")
	key, _ := util.ParsePrivateKey(keyPem)
	return key
}()

// The URL path component corresponding to the cert's sha-256.
var CertName = util.CertName(Certs[0])

// Request encapsulates all the information needed to construct a test request
// for use by the unit tests.
type Request struct {
	T       *testing.T
	Handler http.Handler
	Target  string
	Host    string
	Header  http.Header
	Body    io.Reader
}

// NewRequest returns a new test request.
func NewRequest(t *testing.T, h http.Handler, target string) *Request {
	return &Request{
		T:       t,
		Handler: h,
		Target:  target,
	}
}

// SetHeaders sets the headers for the request. Host may be empty for the default.
func (r *Request) SetHeaders(host string, header http.Header) *Request {
	r.Host = host
	r.Header = header
	return r
}

// SetBody sets the body for the request. May be nil for an empty body.
func (r *Request) SetBody(body io.Reader) *Request {
	r.Body = body
	return r
}

// Get returns the completed test request object.
func (r *Request) Do() *http.Response {
	rec := httptest.NewRecorder()
	method := ""
	if r.Header == nil {
		r.Header = http.Header{}
	}
	if r.Body != nil {
		method = "POST"
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	req := httptest.NewRequest(method, r.Target, r.Body)
	for name, values := range r.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}
	if r.Host != "" {
		req.Host = r.Host
	}
	r.Handler.ServeHTTP(rec, req)
	return rec.Result()
}

type FakeClock struct {
	SecondsSince0 time.Duration
	Delta         time.Duration
}

func NewFakeClock() *FakeClock {
	return &FakeClock{time.Now().Sub(time.Unix(0, 0)), time.Second}
}

func (this *FakeClock) Now() time.Time {
	secondsSince0 := this.SecondsSince0
	this.SecondsSince0 = secondsSince0 + this.Delta
	return time.Unix(0, 0).Add(secondsSince0)
}
