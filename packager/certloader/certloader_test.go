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

package certloader

import (
	"crypto/rsa"
	"crypto/x509"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/ampproject/amppackager/packager/util"
)

func stringPtr(s string) *string {
	return &s
}

var caCert = func() *x509.Certificate {
	certPem, _ := ioutil.ReadFile("../../testdata/b3/ca.cert")
	certs, _ := signedexchange.ParseCertificates(certPem)
	return certs[0]
}()

var caKey = func() *rsa.PrivateKey {
	keyPem, _ := ioutil.ReadFile("../../testdata/b3/ca.privkey")
	key, _ := util.ParsePrivateKey(keyPem)
	return key.(*rsa.PrivateKey)
}()

func TestLoadCertsFromFile(t *testing.T) {
	// Cert file does not exist.
	certs, err := LoadCertsFromFile(
		&util.Config{
			CertFile: "file_does_not_exist",
		},
		true)
	assert.Contains(t, err.Error(), "no such file or directory")

	// Cert file is ok for dev mode.
	certs, err = LoadCertsFromFile(
		&util.Config{
			CertFile: "../../testdata/b3/ca.cert",
		},
		true)
	assert.Equal(t, caCert, certs[0])
	assert.Nil(t, err)

	// Cert file is not ok for prod mode.
	certs, err = LoadCertsFromFile(
		&util.Config{
			CertFile: "../../testdata/b3/ca.cert",
		},
		false)
	assert.Equal(t, certs, ([]*x509.Certificate)(nil))
	assert.Equal(t, err.Error(), "Certificate is missing CanSignHttpExchanges extension")
}

func TestLoadKeyFromFile(t *testing.T) {
	// Key does not exist.
	key, err := LoadKeyFromFile(
		&util.Config{
			KeyFile: "file_does_not_exist",
		})
	assert.Contains(t, err.Error(), "no such file or directory")

	// Key is valid.
	key, err = LoadKeyFromFile(
		&util.Config{
			KeyFile: "../../testdata/b3/ca.privkey",
		})
	assert.Equal(t, caKey, key)
	assert.Nil(t, err)
}
