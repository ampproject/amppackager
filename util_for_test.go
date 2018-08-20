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

package amppackager

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/nyaxt/webpackage/go/signedexchange"
)

// A self-signed cert for testing.
var certPem = func() []byte {
	ret, _ := ioutil.ReadFile("testdata/cert.pem")
	return ret
}()

// The same cert, parsed.
var cert = func() *x509.Certificate {
	certs, _ := signedexchange.ParseCertificates(certPem)
	return certs[0]
}()

// Its corresponding private key.
var keyPem = func() []byte {
	ret, _ := ioutil.ReadFile("testdata/privkey.pem")
	return ret
}()

// The same key, parsed.
var key = func() crypto.PrivateKey {
	keyBlock, _ := pem.Decode(keyPem)
	key, _ := signedexchange.ParsePrivateKey(keyBlock.Bytes)
	return key
}()

// A variant of http.Handler that's required by httprouter.
type AlmostHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request, httprouter.Params)
}

func get(t *testing.T, handler AlmostHandler, target string) *http.Response {
	return getP(t, handler, target, httprouter.Params{})
}

func getP(t *testing.T, handler AlmostHandler, target string, params httprouter.Params) *http.Response {
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest("", target, nil), params)
	return rec.Result()
}
