// +build gofuzz

// Tests that random outputs from the backend don't crash the server.
//
// The dependency is not in the vendor/ directory; get it with:
// $ go get github.com/dvyukov/go-fuzz/...
//
// To run, cd into this directory, and then:
// $ go-fuzz-build -o=fuzz.zip github.com/ampproject/amppackager/fuzz_httpresp
// $ go-fuzz -bin=fuzz.zip -workdir=examples/ -procs=1

package amppackager

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

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

var packager = newPackager()

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
	// Mock out AMP CDN endpoint.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		conn, buf, err := w.(http.Hijacker).Hijack()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		_, err = buf.Write(data)
		if err != nil {
			panic(err)
		}
		err = buf.Flush()
		if err != nil {
			panic(err)
		}
	}))
	defer server.Close()
	url, _ := url.Parse(server.URL)
	amppkg.AmpCDNBase = "http://" + url.Host + "/"

	req := httptest.NewRequest("", "/priv/doc?sign=https%3A%2F%2Fexample.com%2Fconifer-anniston-favorite-moments.html", nil)

	resp := httptest.NewRecorder()
	packager.ServeHTTP(resp, req)
	return 1
}
