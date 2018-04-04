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
	"bytes"
	"crypto/x509"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/WICG/webpackage/go/signedexchange/certurl"
	"github.com/pkg/errors"
)

type CertCache struct {
	// TODO(twifkak): Support multiple certs.
	certName    string
	certMessage []byte
}

func NewCertCache(cert *x509.Certificate, pemContent []byte) (*CertCache, error) {
	this := new(CertCache)
	this.certName = CertName(cert)
	// TODO(twifkak): Refactor CertificateMessageFromPEM to be based on the x509.Certificate instead.
	var err error
	this.certMessage, err = certurl.CertificateMessageFromPEM(pemContent)
	if err != nil {
		return nil, errors.Wrap(err, "extracting certificate from CertFile")
	}
	return this, nil
}

func (this CertCache) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	println("path", req.URL.Path)
	if req.URL.Path == path.Join("/", CertURLPrefix, url.PathEscape(this.certName)) {
		// https://jyasskin.github.io/webpackage/implementation-draft/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cert-chain-format
		resp.Header().Set("Content-Type", "application/octet-stream")
		resp.Header().Set("ETag", "\""+this.certName+"\"")
		// TODO(twifkak): Add cache headers.
		http.ServeContent(resp, req, "", time.Time{}, bytes.NewReader(this.certMessage))
	} else {
		http.NotFound(resp, req)
	}
}
