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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/ampproject/amppackager/packager/util"
	"github.com/julienschmidt/httprouter"
)

// A cert (with its issuer chain) for testing.
var Certs = func() []*x509.Certificate {
	certPem, _ := ioutil.ReadFile("../../testdata/b1/fullchain.cert")
	certs, _ := signedexchange.ParseCertificates(certPem)
	return certs
}()

// Its corresponding private key.
var Key = func() crypto.PrivateKey {
	keyPem, _ := ioutil.ReadFile("../../testdata/b1/server.privkey")
	// This call to ParsePrivateKey() is needed by util_test.go.
	key, _ := util.ParsePrivateKey(keyPem)
	return key
}()

// A variant of http.Handler that's required by httprouter.
type AlmostHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request, httprouter.Params)
}

func Get(t *testing.T, handler AlmostHandler, target string) *http.Response {
	return GetP(t, handler, target, httprouter.Params{})
}

func GetH(t *testing.T, handler AlmostHandler, target string, headers http.Header) *http.Response {
	return GetHP(t, handler, target, headers, httprouter.Params{})
}

func GetP(t *testing.T, handler AlmostHandler, target string, params httprouter.Params) *http.Response {
	return GetHP(t, handler, target, http.Header{}, params)
}

func GetHP(t *testing.T, handler AlmostHandler, target string, headers http.Header, params httprouter.Params) *http.Response {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("", target, nil)
	for name, values := range headers {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}
	handler.ServeHTTP(rec, req, params)
	return rec.Result()
}
