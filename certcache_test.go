package amppackager

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/pkg/errors"
)

func newCertCache(t *testing.T) *CertCache {
	handler, err := NewCertCache(cert, certPem)
	if err != nil {
		t.Fatal(errors.WithStack(err))
	}
	return handler
}

func TestServesCertificate(t *testing.T) {
	resp := get(t, newCertCache(t), "/amppkg/cert/sLtQsuGUOYdCsBVuMTUG_6QBAWFHu8rhEokEHQAWmto=")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("invalid status: %#v", resp)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) < 20 { // too small to fit a cert
		t.Errorf("invalid body: %q", body)
	}
}

func TestServes404OnMissingCertificate(t *testing.T) {
	resp := get(t, newCertCache(t), "/amppkg/cert/lalala")
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("invalid status: %#v", resp)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) > 20 { // bigger than expected, might be a cert or key
		t.Errorf("invalid body: %q", body)
	}
}
