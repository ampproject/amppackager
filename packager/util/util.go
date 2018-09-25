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

package util

import (
	"bytes"
	"crypto"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/pkg/errors"
)

const CertURLPrefix = "/amppkg/cert"

// CertName returns the basename for the given cert, as served by this
// packager's cert cache. Should be stable and unique (e.g.
// content-addressing). Clients should url.PathEscape this, just in case its
// format changes to need escaping in the future.
func CertName(cert *x509.Certificate) string {
	sum := sha256.Sum256(cert.Raw)
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

const ValidityMapPath = "/amppkg/validity"

// ParsePrivateKey returns the first PEM block that looks like a private key.
func ParsePrivateKey(keyPem []byte) (crypto.PrivateKey, error) {
	var privkey crypto.PrivateKey
	for {
		var pemBlock *pem.Block
		pemBlock, keyPem = pem.Decode(keyPem)
		if pemBlock == nil {
			return nil, errors.New("invalid PEM block in private key file")
		}

		var err error
		privkey, err = signedexchange.ParsePrivateKey(pemBlock.Bytes)
		if err == nil {
			return privkey, nil
		}
		if len(keyPem) == 0 {
			// No more PEM blocks to try.
			return nil, errors.New("failed to parse private key file")
		}
		// Else try next PEM block.
	}
}

// CanSignHttpExchanges returns true if the given certificate has the
// CanSignHttpExchanges extension. This is not the only requirement for SXGs;
// it also needs to use the right public key type, which is not checked here.
func CanSignHttpExchanges(cert *x509.Certificate) bool {
	// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cross-origin-cert-req
	for _, ext := range cert.Extensions {
		// 0x05, 0x00 is the DER encoding of NULL.
		if ext.Id.Equal(asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 11129, 2, 1, 22}) && bytes.Equal(ext.Value, []byte{0x05, 0x00}) {
			return true
		}
	}
	return false
}
