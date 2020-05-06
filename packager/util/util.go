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
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"time"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/pkg/errors"
)

const CertURLPrefix = "/amppkg/cert"
const SignerURLPrefix = "/priv/doc"

// CertName returns the basename for the given cert, as served by this
// packager's cert cache. Should be stable and unique (e.g.
// content-addressing). Clients should url.PathEscape this, just in case its
// format changes to need escaping in the future.
func CertName(cert *x509.Certificate) string {
	sum := sha256.Sum256(cert.Raw)
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

const ValidityMapPath = "/amppkg/validity"
const HealthzPath = "/healthz"
const MetricsPath = "/metrics"

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

func hasCanSignHttpExchangesExtension(cert *x509.Certificate) bool {
	// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cross-origin-cert-req
	for _, ext := range cert.Extensions {
		// 0x05, 0x00 is the DER encoding of NULL.
		if ext.Id.Equal(asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 11129, 2, 1, 22}) && bytes.Equal(ext.Value, []byte{0x05, 0x00}) {
			return true
		}
	}
	return false
}

// Returns the Duration of time before cert expires with given deadline.
// Note that the certExpiryDeadline should be the expected SXG expiration time.
// Returns error if cert is already expired. This will be used to periodically check if cert
// is still within validity range.
func GetDurationToExpiry(cert *x509.Certificate, certExpiryDeadline time.Time) (time.Duration, error) {
	if cert.NotBefore.After(certExpiryDeadline) {
		return 0, errors.New("Certificate is future-dated")
	}
	if cert.NotAfter.Before(certExpiryDeadline) {
		return 0, errors.New("Certificate is expired")
	}

	return cert.NotAfter.Sub(certExpiryDeadline), nil
}

// CanSignHttpExchanges returns nil if the given certificate has the
// CanSignHttpExchanges extension, and a valid lifetime per the SXG spec;
// otherwise it returns an error. These are not the only requirements for SXGs;
// it also needs to use the right public key type, which is not checked here.
func CanSignHttpExchanges(cert *x509.Certificate) error {
	if !hasCanSignHttpExchangesExtension(cert) {
		return errors.New("Certificate is missing CanSignHttpExchanges extension")
	}
	if cert.NotBefore.AddDate(0, 0, 90).Before(cert.NotAfter) {
		return errors.New("Certificate MUST have a Validity Period no greater than 90 days")
	}
	return nil
}

// Returns nil if the certificate matches the private key and domain, else the appropriate error.
func CertificateMatches(cert *x509.Certificate, priv crypto.PrivateKey, domain string) error {
	certPubKey := cert.PublicKey.(*ecdsa.PublicKey)
	pubKey := priv.(*ecdsa.PrivateKey).PublicKey
	if certPubKey.Curve != pubKey.Curve {
		return errors.New("PublicKey.Curve not match")
	}
	if certPubKey.X.Cmp(pubKey.X) != 0 {
		return errors.New("PublicKey.X not match")
	}
	if certPubKey.Y.Cmp(pubKey.Y) != 0 {
		return errors.New("PublicKey.Y not match")
	}
	if err := cert.VerifyHostname(domain); err != nil {
		return err
	}
	return nil
}
