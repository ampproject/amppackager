package main

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"time"

	"golang.org/x/crypto/ocsp"
)

// A fake OCSP responder for the SXG certificate. Abstracted from certcache so as not to give it direct access to the private key.
// Assumes that the given cert is self-signed, and therefore that the
// configured KeyFile also corresponds to the cert's issuer.
type fakeOCSPResponder struct {
	key crypto.Signer
}

// Generates a current OCSP response for the given certificate.
func (this fakeOCSPResponder) Respond(cert *x509.Certificate) ([]byte, error) {
	thisUpdate := time.Now()

	// Construct args to ocsp.CreateResponse.
	template := ocsp.Response{
		SerialNumber: cert.SerialNumber,
		Status:       ocsp.Good,
		ThisUpdate:   thisUpdate,
		NextUpdate:   thisUpdate.Add(time.Hour * 24 * 7),
		IssuerHash:   crypto.SHA256,
	}
	subjectDER, err := asn1.Marshal(pkix.Name{CommonName: "fake-responder.example"}.ToRDNSequence())
	if err != nil {
		return nil, err
	}
	responderCert := x509.Certificate{
		// This is the only field that ocsp.CreateResponse reads.
		RawSubject: subjectDER,
	}
	resp, err := ocsp.CreateResponse(cert, &responderCert, template, this.key)
	if err != nil {
		return nil, err
	}
	_, err = ocsp.ParseResponseForCert(resp, cert, cert)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
