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
	"crypto/rsa"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/ocsp"
)

var caCert = func() *x509.Certificate {
	certPem, _ := ioutil.ReadFile("testdata/b1/ca.cert")
	certs, _ := signedexchange.ParseCertificates(certPem)
	return certs[0]
}()

var caKey = func() *rsa.PrivateKey {
	keyPem, _ := ioutil.ReadFile("testdata/b1/ca.privkey")
	key, _ := ParsePrivateKey(keyPem)
	return key.(*rsa.PrivateKey)
}()


type CertCacheTestSuite struct {
	suite.Suite
	ocspServer *httptest.Server
	tempDir string
	stop    chan struct{}
	handler *CertCache
}

func (this *CertCacheTestSuite) SetupSuite() {
	this.ocspServer = httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		template := ocsp.Response{
			Status: ocsp.Good,
			SerialNumber: certs[0].SerialNumber,
			ThisUpdate: time.Now(),
			NextUpdate: time.Now().Add(7 * 24 * time.Hour),
			RevokedAt: time.Now().AddDate(/*years=*/0, /*months=*/0, /*days=*/14),
			RevocationReason: ocsp.Unspecified,
		}
		body, err := ocsp.CreateResponse(caCert, caCert, template, caKey)
		require.NoError(this.T(), err, "creating fake OCSP response")
		_, err = resp.Write(body)
		require.NoError(this.T(), err, "writing fake OCSP response")
	}))
	fakeOCSPServer = this.ocspServer.URL
}

func (this *CertCacheTestSuite) TearDownSuite() {
	fakeOCSPServer = ""
	this.ocspServer.Close()
}

func (this *CertCacheTestSuite) SetupTest() {
	var err error
	this.tempDir, err = ioutil.TempDir(os.TempDir(), "certcache_test")
	require.NoError(this.T(), err, "setting up test harness")

	this.stop = make(chan struct{})

	this.handler, err = NewCertCache(certs, filepath.Join(this.tempDir, "ocsp"), this.stop)
	require.NoError(this.T(), err, "instantiating CertCache")
}

func (this *CertCacheTestSuite) TearDownTest() {
	this.stop <- struct{}{}

	err := os.RemoveAll(this.tempDir)
	if err != nil {
		log.Panic("Error removing temp dir", err)
	}
}

func (this *CertCacheTestSuite) TestServesCertificate() {
	resp := getP(this.T(), this.handler, "/amppkg/cert/k9GCZZIDzAt2X0b2czRv0c2omW5vgYNh6ZaIz_UNTRQ", httprouter.Params{httprouter.Param{"certName", "k9GCZZIDzAt2X0b2czRv0c2omW5vgYNh6ZaIz_UNTRQ"}})
	assert.Equal(this.T(), http.StatusOK, resp.StatusCode, "incorrect status: %#v", resp)
	body, _ := ioutil.ReadAll(resp.Body)
	// Large enough to fit a cert:
	assert.Condition(this.T(), func() bool { return len(body) >= 20 }, "body too small: %q", body)
}

func (this *CertCacheTestSuite) TestServes404OnMissingCertificate() {
	resp := getP(this.T(), this.handler, "/amppkg/cert/lalala", httprouter.Params{httprouter.Param{"certName", "lalala"}})
	assert.Equal(this.T(), http.StatusNotFound, resp.StatusCode, "incorrect status: %#v", resp)
	body, _ := ioutil.ReadAll(resp.Body)
	// Small enough not to fit a cert or key:
	assert.Condition(this.T(), func() bool { return len(body) <= 20 }, "body too large: %q", body)
}

func TestCertCacheSuite(t *testing.T) {
	suite.Run(t, new(CertCacheTestSuite))
}

// TODO(twifkak): Test:
// no memory or disk
// disk expired
// disk fresh
// memory expired, disk fresh
// disk expired, memory fresh
