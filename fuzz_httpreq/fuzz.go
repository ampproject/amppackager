// +build gofuzz

// Tests that random inputs from the frontend don't crash the server.
//
// The dependency is not in the vendor/ directory; get it with:
// $ go get github.com/dvyukov/go-fuzz/...
//
// To run, cd into this directory, and then:
// $ go-fuzz-build -o=fuzz.zip github.com/ampproject/amppackager/fuzz_httpreq
// $ go-fuzz -bin=fuzz.zip -workdir=examples/

package amppackager

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/pkg/errors"

	amppkg "github.com/ampproject/amppackager"
)

// A self-signed cert for testing.
var cert = func() *x509.Certificate {
	certPem, err := ioutil.ReadFile("../testdata/cert.pem")
	if err != nil {
		panic(err)
	}
	certs, err := signedexchange.ParseCertificates(certPem)
	if err != nil {
		panic(err)
	}
	return certs[0]
}()

// Its corresponding private key.
var key = func() crypto.PrivateKey {
	keyPem, err := ioutil.ReadFile("../testdata/privkey.pem")
	if err != nil {
		panic(err)
	}
	keyBlock, _ := pem.Decode(keyPem)
	key, err := signedexchange.ParsePrivateKey(keyBlock.Bytes)
	if err != nil {
		panic(err)
	}
	return key
}()

var fakeBody = []byte("They like to OPINE. Get it? (Is he fir real? Yew gotta be kidding me.)")

var packager = newPackager()

var once sync.Once

func initBackend() {
	// Mock out AMP CDN endpoint.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write(fakeBody)
	}))
	// server.Close() is never called. It is expected that it shuts down when the binary quits.
	url, _ := url.Parse(server.URL)
	amppkg.AmpCDNBase = "http://" + url.Host + "/"
}

func stringPtr(x string) *string { return &x }

func newPackager() *amppkg.Packager {
	urlSets := []amppkg.URLSet{{
		Sign:  &amppkg.URLPattern{[]string{"https"}, "", "example.com", stringPtr(".*"), []string{}, stringPtr(""), false, nil},
		Fetch: nil,
	}}
	handler, err := amppkg.NewPackager(cert, key, "https://example.com/", urlSets)
	if err != nil {
		panic(errors.WithStack(err))
	}
	return handler
}

func Fuzz(data []byte) int {
	req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(data)))
	if err != nil {
		return 0
	}

	once.Do(initBackend)
	resp := httptest.NewRecorder()
	packager.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		return 0
	}
	return 1
}
