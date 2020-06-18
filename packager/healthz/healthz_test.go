// Copyright 2019 Google LLC
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

package healthz

import (
	"crypto/x509"
	"net/http"
	"testing"

	"github.com/ampproject/amppackager/packager/mux"
	"github.com/pkg/errors"

	pkgt "github.com/ampproject/amppackager/packager/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeHealthyCertHandler struct {
}

func (this fakeHealthyCertHandler) GetLatestCert() *x509.Certificate {
	return pkgt.Certs[0]
}

func (this fakeHealthyCertHandler) IsHealthy() error {
	return nil
}

type fakeNotHealthyCertHandler struct {
}

func (this fakeNotHealthyCertHandler) GetLatestCert() *x509.Certificate {
	return pkgt.Certs[0]
}

func (this fakeNotHealthyCertHandler) IsHealthy() error {
	return errors.New("random error")
}

func TestHealthzOk(t *testing.T) {
	handler, err := New(fakeHealthyCertHandler{})
	require.NoError(t, err)
	resp := pkgt.Get(t, mux.New(nil, nil, nil, handler, nil), "/healthz")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "ok", resp)
}

func TestHealthzFail(t *testing.T) {
	handler, err := New(fakeNotHealthyCertHandler{})
	require.NoError(t, err)
	resp := pkgt.Get(t, mux.New(nil, nil, nil, handler, nil), "/healthz")
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "error", resp)
}
