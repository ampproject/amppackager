package amppackager

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

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

func get(t *testing.T, handler http.Handler, target string) *http.Response {
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest("", target, nil))
	return rec.Result()
}
