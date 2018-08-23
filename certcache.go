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
	"time"

	"github.com/WICG/webpackage/go/signedexchange/certurl"
	"github.com/julienschmidt/httprouter"
)

type CertCache struct {
	// TODO(twifkak): Support multiple certs.
	certName string
	cert     *x509.Certificate
}

func NewCertCache(cert *x509.Certificate) (*CertCache, error) {
	this := new(CertCache)
	this.certName = CertName(cert)
	this.cert = cert
	return this, nil
}

func (this CertCache) ServeHTTP(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if params.ByName("certName") == this.certName {
		// https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-3.3
		// This content-type is not standard, but included to reduce
		// the chance that faulty user agents employ content sniffing.
		resp.Header().Set("Content-Type", "application/tls-certificate-chain")
		resp.Header().Set("Cache-Control", "public, max-age=604800")
		resp.Header().Set("ETag", "\""+this.certName+"\"")
		// TODO(twifkak): Specify real OCSP and SCT blobs.
		cbor, err := certurl.CreateCertChainCBOR([]*x509.Certificate{this.cert}, []byte{}, []byte{})
		if err != nil {
			NewHTTPError(http.StatusInternalServerError, "Error build cert chain: ", err).LogAndRespond(resp)
			return
		}
		http.ServeContent(resp, req, "", time.Time{}, bytes.NewReader(cbor))
	} else {
		http.NotFound(resp, req)
	}
}
