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

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func newCertCache(t *testing.T) *CertCache {
	handler, err := NewCertCache(cert)
	if err != nil {
		t.Fatal(errors.WithStack(err))
	}
	return handler
}

func TestServesCertificate(t *testing.T) {
	resp := getP(t, newCertCache(t), "/amppkg/cert/sLtQsuGUOYdCsBVuMTUG_6QBAWFHu8rhEokEHQAWmto", httprouter.Params{httprouter.Param{"certName", "sLtQsuGUOYdCsBVuMTUG_6QBAWFHu8rhEokEHQAWmto"}})
	assert.Equal(t, http.StatusOK, resp.StatusCode, "incorrect status: %#v", resp)
	body, _ := ioutil.ReadAll(resp.Body)
	// Large enough to fit a cert:
	assert.Condition(t, func() bool { return len(body) >= 20 }, "body too small: %q", body)
}

func TestServes404OnMissingCertificate(t *testing.T) {
	resp := getP(t, newCertCache(t), "/amppkg/cert/lalala", httprouter.Params{httprouter.Param{"certName", "lalala"}})
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "incorrect status: %#v", resp)
	body, _ := ioutil.ReadAll(resp.Body)
	// Small enough not to fit a cert or key:
	assert.Condition(t, func() bool { return len(body) <= 20 }, "body too large: %q", body)
}
