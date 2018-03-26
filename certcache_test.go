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
