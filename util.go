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
	"crypto"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/pkg/errors"
)

// CertURLPrefix must start without a slash, for PackagerBase's sake.
const CertURLPrefix = "amppkg/cert"

// CertName returns the basename for the given cert, as served by this
// packager's cert cache. Should be stable and unique (e.g.
// content-addressing). Clients should url.PathEscape this, just in case its
// format changes to need escaping in the future.
func CertName(cert *x509.Certificate) string {
	sum := sha256.Sum256(cert.Raw)
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

// ValidityMapURL must start without a slash, for PackagerBase's sake.
const ValidityMapURL = "amppkg/validity"

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
		if err == nil || len(keyPem) == 0 {
			return privkey, nil
		}
		// Else try next PEM block.
	}
	return nil, errors.New("failed to parse private key file")
}

