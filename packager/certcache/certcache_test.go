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

package certcache

import (
	"crypto/rsa"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/WICG/webpackage/go/signedexchange/cbor"
	pkgt "github.com/ampproject/amppackager/packager/testing"
	"github.com/ampproject/amppackager/packager/util"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/ocsp"
)

var caCert = func() *x509.Certificate {
	certPem, _ := ioutil.ReadFile("../../testdata/b1/ca.cert")
	certs, _ := signedexchange.ParseCertificates(certPem)
	return certs[0]
}()

var caKey = func() *rsa.PrivateKey {
	keyPem, _ := ioutil.ReadFile("../../testdata/b1/ca.privkey")
	key, _ := util.ParsePrivateKey(keyPem)
	return key.(*rsa.PrivateKey)
}()

func FakeOCSPResponse(thisUpdate time.Time) ([]byte, error) {
	template := ocsp.Response{
		Status:           ocsp.Good,
		SerialNumber:     pkgt.Certs[0].SerialNumber,
		ThisUpdate:       thisUpdate,
		NextUpdate:       thisUpdate.Add(7 * 24 * time.Hour),
		RevokedAt:        thisUpdate.AddDate( /*years=*/ 0 /*months=*/, 0 /*days=*/, 365),
		RevocationReason: ocsp.Unspecified,
	}
	return ocsp.CreateResponse(caCert, caCert, template, caKey)
}

type CertCacheSuite struct {
	suite.Suite
	fakeOCSP            []byte
	fakeOCSPExpiry      *time.Time
	ocspServer          *httptest.Server // "const", do not set
	ocspServerWasCalled bool
	ocspHandler         func(w http.ResponseWriter, req *http.Request)
	tempDir             string
	stop                chan struct{}
	handler             *CertCache
}

func (this *CertCacheSuite) New() (*CertCache, error) {
	// TODO(twifkak): Stop the old CertCache's goroutine.
	certCache := New(pkgt.Certs, filepath.Join(this.tempDir, "ocsp"))
	certCache.extractOCSPServer = func(*x509.Certificate) (string, error) {
		return this.ocspServer.URL, nil
	}
	defaultHttpExpiry := certCache.httpExpiry
	certCache.httpExpiry = func(req *http.Request, resp *http.Response) time.Time {
		if this.fakeOCSPExpiry != nil {
			return *this.fakeOCSPExpiry
		} else {
			return defaultHttpExpiry(req, resp)
		}
	}
	err := certCache.Init(this.stop)
	return certCache, err
}

func (this *CertCacheSuite) SetupSuite() {
	this.ocspServer = httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		this.ocspHandler(resp, req)
	}))
}

func (this *CertCacheSuite) TearDownSuite() {
	this.ocspServer.Close()
}

func (this *CertCacheSuite) SetupTest() {
	var err error
	this.fakeOCSP, err = FakeOCSPResponse(time.Now())
	this.Require().NoError(err, "creating fake OCSP response")

	this.ocspHandler = func(resp http.ResponseWriter, req *http.Request) {
		this.ocspServerWasCalled = true
		_, err := resp.Write(this.fakeOCSP)
		this.Require().NoError(err, "writing fake OCSP response")
	}

	this.tempDir, err = ioutil.TempDir(os.TempDir(), "certcache_test")
	this.Require().NoError(err, "setting up test harness")

	this.stop = make(chan struct{})

	this.handler, err = this.New()
	this.Require().NoError(err, "instantiating CertCache")
}

func (this *CertCacheSuite) TearDownTest() {
	// Reset any variables that may have been overridden in test and won't be rewritten in SetupTest.
	this.fakeOCSPExpiry = nil

	// Reverse SetupTest.
	this.stop <- struct{}{}

	err := os.RemoveAll(this.tempDir)
	if err != nil {
		log.Panic("Error removing temp dir", err)
	}
}

func (this *CertCacheSuite) ocspServerCalled(f func()) bool {
	this.ocspServerWasCalled = false
	f()
	return this.ocspServerWasCalled
}

func (this *CertCacheSuite) DecodeCBOR(r io.Reader) map[string][]byte {
	decoder := cbor.NewDecoder(r)

	// Our test cert chain has exactly two certs. First entry is a magic.
	numItems, err := decoder.DecodeArrayHeader()
	this.Require().NoError(err, "decoding array header")
	this.Require().EqualValues(3, numItems)

	magic, err := decoder.DecodeTextString()
	this.Require().NoError(err, "decoding magic")
	this.Require().Equal("📜⛓", magic)

	// Decode and return the first one.
	numKeys, err := decoder.DecodeMapHeader()
	this.Require().NoError(err, "decoding map header")
	this.Require().EqualValues(2, numKeys)

	ret := map[string][]byte{}
	for i := 0; uint64(i) < numKeys; i++ {
		key, err := decoder.DecodeTextString()
		this.Require().NoError(err, "decoding key")
		value, err := decoder.DecodeByteString()
		this.Require().NoError(err, "decoding value")
		ret[key] = value
	}
	return ret
}

func (this *CertCacheSuite) TestServesCertificate() {
	resp := pkgt.GetP(this.T(), this.handler, "/amppkg/cert/"+pkgt.CertName, httprouter.Params{httprouter.Param{"certName", pkgt.CertName}})
	this.Assert().Equal(http.StatusOK, resp.StatusCode, "incorrect status: %#v", resp)
	this.Assert().Equal("nosniff", resp.Header.Get("X-Content-Type-Options"))
	cbor := this.DecodeCBOR(resp.Body)
	this.Assert().Contains(cbor, "cert")
	this.Assert().Contains(cbor, "ocsp")
	this.Assert().NotContains(cbor, "sct")
}

func (this *CertCacheSuite) TestServes404OnMissingCertificate() {
	resp := pkgt.GetP(this.T(), this.handler, "/amppkg/cert/lalala", httprouter.Params{httprouter.Param{"certName", "lalala"}})
	this.Assert().Equal(http.StatusNotFound, resp.StatusCode, "incorrect status: %#v", resp)
	body, _ := ioutil.ReadAll(resp.Body)
	// Small enough not to fit a cert or key:
	this.Assert().Condition(func() bool { return len(body) <= 20 }, "body too large: %q", body)
}

func (this *CertCacheSuite) TestOCSP() {
	// Verify it gets included in the cert-chain+cbor payload.
	resp := pkgt.GetP(this.T(), this.handler, "/amppkg/cert/"+pkgt.CertName, httprouter.Params{httprouter.Param{"certName", pkgt.CertName}})
	this.Assert().Equal(http.StatusOK, resp.StatusCode, "incorrect status: %#v", resp)
	// 302400 is 3.5 days. max-age is slightly less because of the time between fake OCSP generation and cert-chain response.
	// TODO(twifkak): Make this less flaky, by injecting a fake clock.
	this.Assert().Equal("public, max-age=302399", resp.Header.Get("Cache-Control"))
	cbor := this.DecodeCBOR(resp.Body)
	this.Assert().Equal(this.fakeOCSP, cbor["ocsp"])
}

func (this *CertCacheSuite) TestOCSPCached() {
	// Verify it is in the memory cache:
	this.Assert().False(this.ocspServerCalled(func() {
		_, _, err := this.handler.readOCSP()
		this.Assert().NoError(err)
	}))

	// Create a new handler, to see it populates the memory cache from disk, not network:
	this.Assert().False(this.ocspServerCalled(func() {
		_, err := this.New()
		this.Require().NoError(err, "reinstantiating CertCache")
	}))
}

func (this *CertCacheSuite) TestOCSPExpiry() {
	// Prime memory and disk cache with a past-midpoint OCSP:
	err := os.Remove(filepath.Join(this.tempDir, "ocsp"))
	this.Require().NoError(err, "deleting OCSP tempfile")
	this.fakeOCSP, err = FakeOCSPResponse(time.Now().Add(-4 * 24 * time.Hour))
	this.Require().NoError(err, "creating expired OCSP response")
	this.Require().True(this.ocspServerCalled(func() {
		this.handler, err = this.New()
		this.Require().NoError(err, "reinstantiating CertCache")
	}))

	// Verify HTTP response expires immediately:
	resp := pkgt.GetP(this.T(), this.handler, "/amppkg/cert/"+pkgt.CertName, httprouter.Params{httprouter.Param{"certName", pkgt.CertName}})
	this.Assert().Equal("public, max-age=0", resp.Header.Get("Cache-Control"))

	// On update, verify network is called:
	this.Assert().True(this.ocspServerCalled(func() {
		_, _, err := this.handler.readOCSP()
		this.Assert().NoError(err)
	}))
}

func (this *CertCacheSuite) TestOCSPUpdateFromDisk() {
	// Prime memory cache with a past-midpoint OCSP:
	err := os.Remove(filepath.Join(this.tempDir, "ocsp"))
	this.Require().NoError(err, "deleting OCSP tempfile")
	this.fakeOCSP, err = FakeOCSPResponse(time.Now().Add(-4 * 24 * time.Hour))
	this.Require().NoError(err, "creating stale OCSP response")
	this.Require().True(this.ocspServerCalled(func() {
		this.handler, err = this.New()
		this.Require().NoError(err, "reinstantiating CertCache")
	}))

	// Prime disk cache with a fresh OCSP.
	freshOCSP, err := FakeOCSPResponse(time.Now())
	this.Require().NoError(err, "creating fresh OCSP response")
	err = ioutil.WriteFile(filepath.Join(this.tempDir, "ocsp"), freshOCSP, 0644)
	this.Require().NoError(err, "writing fresh OCSP response to disk")

	// On update, verify network is not called (fresh OCSP from disk is used):
	this.Assert().False(this.ocspServerCalled(func() {
		_, _, err := this.handler.readOCSP()
		this.Assert().NoError(err)
	}))
}

func (this *CertCacheSuite) TestOCSPExpiredViaHTTPHeaders() {
	// Prime memory and disk cache with a fresh OCSP but soon-to-expire HTTP headers:
	err := os.Remove(filepath.Join(this.tempDir, "ocsp"))
	this.Require().NoError(err, "deleting OCSP tempfile")
	this.fakeOCSPExpiry = new(time.Time)
	*this.fakeOCSPExpiry = time.Unix(0, 1) // Infinite past. time.Time{} is used as a sentinel value to mean no update.
	this.Require().True(this.ocspServerCalled(func() {
		this.handler, err = this.New()
		this.Require().NoError(err, "reinitializing CertCache")
	}))
	this.Require().Equal(time.Unix(0, 1), this.handler.ocspUpdateAfter)

	// Verify that, 2 seconds later, a new fetch is attempted.
	this.Assert().True(this.ocspServerCalled(func() {
		_, _, err := this.handler.readOCSP()
		this.Require().NoError(err, "updating OCSP")
	}))
}

func (this *CertCacheSuite) TestOCSPIgnoreInvalidUpdate() {
	// Prime memory and disk cache with a past-midpoint OCSP:
	err := os.Remove(filepath.Join(this.tempDir, "ocsp"))
	this.Require().NoError(err, "deleting OCSP tempfile")
	staleOCSP, err := FakeOCSPResponse(time.Now().Add(-4 * 24 * time.Hour))
	this.Require().NoError(err, "creating stale OCSP response")
	this.fakeOCSP = staleOCSP
	this.Require().True(this.ocspServerCalled(func() {
		this.handler, err = this.New()
		this.Require().NoError(err, "reinstantiating CertCache")
	}))

	// Try to update with an invalid OCSP:
	this.fakeOCSP, err = FakeOCSPResponse(time.Now().Add(-8 * 24 * time.Hour))
	this.Require().NoError(err, "creating expired OCSP response")
	this.Assert().True(this.ocspServerCalled(func() {
		_, _, err := this.handler.readOCSP()
		this.Require().NoError(err, "updating OCSP")
	}))

	// Verify that the invalid update doesn't squash the valid cache entry.
	ocsp, _, err := this.handler.readOCSP()
	this.Require().NoError(err, "reading OCSP")
	this.Assert().Equal(staleOCSP, ocsp)
}

func TestCertCacheSuite(t *testing.T) {
	suite.Run(t, new(CertCacheSuite))
}
