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
	"context"
	"crypto/x509"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/WICG/webpackage/go/signedexchange/certurl"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/pquerna/cachecontrol"
	"golang.org/x/crypto/ocsp"
)

// The maximum representable time, per https://stackoverflow.com/questions/25065055/what-is-the-maximum-time-time-in-go.
var infiniteFuture = time.Unix(1<<63-62135596801, 999999999)

// The OCSP code below aims to satisfy items 1-8 of the "sleevi" doc:
// https://gist.github.com/sleevi/5efe9ef98961ecfb4da8
// Item #9 should be added. Item #10 is N/A for SXGs.

// 1MB is the maximum used by github.com/xenolf/lego/acmev2 in GetOCSPForCert.
// Alternatively, here's a random documented example of a 20K limit:
// https://www.ibm.com/support/knowledgecenter/en/SSPREK_9.0.0/com.ibm.isam.doc/wrp_stza_ref/reference/ref_ocsp_max_size.html
const maxOCSPResponseBytes = 1024 * 1024

// How often to check if OCSP stapling needs updating.
const ocspCheckInterval = 1 * time.Hour

var fakeOCSPServer string
var fakeOCSPExpiry *time.Time

type CertCache struct {
	// TODO(twifkak): Support multiple cert chains (for different domains, for different roots).
	certName string
	certs    []*x509.Certificate
	// Lock for upgrading a read-lock to a write-lock atomically. This ensures
	// in-process transactional semantics. (This isn't strictly necessary
	// currently, given the updater runs in a single goroutine.)
	// TODO(twifkak): Extract an ocsp struct and a transactional in-memory thingie.
	ocspMuMu        sync.Mutex
	ocspMu          sync.RWMutex // Lock over ocsp and ocspUpdateAfter.
	ocsp            []byte
	ocspUpdateAfter time.Time
	// TODO(twifkak): Implement a registry of Updateable instances which can be configured in the toml.
	ocspFile Updateable
	client   http.Client
}

func NewCertCache(certs []*x509.Certificate, ocspCache string, stop chan struct{}) (*CertCache, error) {
	this := new(CertCache)
	this.certName = CertName(certs[0])
	this.certs = certs
	// sleevi #1, sleevi #6:
	this.ocspFile = LocalFile{path: ocspCache}
	this.client = http.Client{Timeout: 60 * time.Second}
	// Prime the OCSP disk and memory cache, so we can start serving immediately.
	this.ocspUpdateAfter = infiniteFuture // Default, in case initial maybeUpdateOCSP successfully loads from disk.
	err := this.maybeUpdateOCSP()
	if err != nil {
		return nil, errors.Wrap(err, "initializing CertCache")
	}
	// Update OCSP in the background (sleevi #3, sleevi #7).
	go this.maintainOCSP(stop)
	return this, nil
}

func (this *CertCache) ServeHTTP(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if params.ByName("certName") == this.certName {
		// https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-3.3
		// This content-type is not standard, but included to reduce
		// the chance that faulty user agents employ content sniffing.
		resp.Header().Set("Content-Type", "application/tls-certificate-chain")
		resp.Header().Set("Cache-Control", "public, max-age=604800")
		resp.Header().Set("ETag", "\""+this.certName+"\"")
		// TODO(twifkak): Specify real SCT blob.
		this.ocspMu.RLock()
		cbor, err := certurl.CreateCertChainCBOR(this.certs, this.ocsp, []byte{})
		this.ocspMu.RUnlock()
		if err != nil {
			NewHTTPError(http.StatusInternalServerError, "Error build cert chain: ", err).LogAndRespond(resp)
			return
		}
		http.ServeContent(resp, req, "", time.Time{}, bytes.NewReader(cbor))
	} else {
		http.NotFound(resp, req)
	}
}

// If we've been unable to fetch a fresh OCSP response before expiry of the old
// one, or, at server start-up, if we're unable to fetch a valid OCSP request at
// all (either from disk or network), then return false. This signals to the
// packager that it should not try to package anything; just proxy the content
// unsigned. (sleevi #8)
func (this *CertCache) IsHealthy() bool {
	this.ocspMu.RLock()
	defer this.ocspMu.RUnlock()
	return this.isHealthy(this.ocsp)
}

func (this *CertCache) isHealthy(ocspResp []byte) bool {
	if ocspResp == nil {
		log.Println("OCSP response not yet fetched.")
		return false
	}
	issuer := this.findIssuer()
	if issuer == nil {
		log.Println("Cannot find issuer certificate in CertFile.")
		return false
	}
	resp, err := ocsp.ParseResponseForCert(ocspResp, this.certs[0], issuer)
	if err != nil {
		log.Println("Error parsing OCSP response:", err)
		return false
	}
	if resp.NextUpdate.Before(time.Now()) {
		log.Println("Cached OCSP is stale, NextUpdate:", resp.NextUpdate)
		return false
	}
	return true
}

func (this *CertCache) maybeUpdateOCSP() error {
	this.ocspMu.RLock()
	var ocspUpdateAfter time.Time
	ocsp, err := this.ocspFile.Read(context.Background(), this.shouldUpdateOCSP, func(orig []byte) []byte {
		return this.updateOCSP(orig, &ocspUpdateAfter)
	})
	if err != nil {
		return errors.Wrap(err, "Updating OCSP cache")
	}
	if len(ocsp) == 0 {
		return errors.New("Missing OCSP response.")
	}
	if !this.isHealthy(ocsp) {
		return errors.New("OCSP failed health check.")
	}
	// Convert read lock into write lock before updating ocsp and ocspUpdateAfter.
	this.ocspMuMu.Lock()
	this.ocspMu.RUnlock()
	this.ocspMu.Lock()
	this.ocspMuMu.Unlock()
	this.ocsp = ocsp
	if !ocspUpdateAfter.Equal(time.Time{}) {
		// updateOCSP was called, and therefore a new HTTP cache expiry was set.
		// TODO(twifkak): Write this to disk, so any replica can pick it up.
		this.ocspUpdateAfter = ocspUpdateAfter
	}
	this.ocspMu.Unlock()
	return nil
}

// Checks for OCSP updates every hour. Never terminates.
func (this *CertCache) maintainOCSP(stop chan struct{}) {
	// Only make one request per ocspCheckInterval, to minimize the impact
	// on OCSP servers that are buckling under load (sleevi #5).
	ticker := time.NewTicker(ocspCheckInterval)

	for {
		select {
		case <-ticker.C:
			err := this.maybeUpdateOCSP()
			if err != nil {
				log.Println("Warning: OCSP update failed. Cached response may expire:", err)
			}
		case <-stop:
			ticker.Stop()
			return
		}
	}
}


func (this *CertCache) shouldUpdateOCSP(bytes []byte) bool {
	if len(bytes) == 0 {
		// TODO(twifkak): Use a logging framework with support for debug-only statements.
		log.Println("Updating OCSP; none cached yet.")
		return true
	}
	issuer := this.findIssuer()
	if issuer == nil {
		log.Println("Cannot find issuer certificate in CertFile.")
		// This is a permanent error; do not attempt OCSP update.
		return false
	}
	resp, err := ocsp.ParseResponseForCert(bytes, this.certs[0], issuer)
	if err != nil {
		log.Println("Error parsing OCSP:", err)
		return true
	}
	// sleevi #3:
	midpoint := resp.ThisUpdate.Add(resp.NextUpdate.Sub(resp.ThisUpdate)/2)
	if time.Now().After(midpoint) {
		// TODO(twifkak): Use a logging framework with support for debug-only statements.
		log.Println("Updating OCSP; after midpoint: ", midpoint)
		return true
	}
	// Allow cache-control headers to indicate an earlier update time, per
	// https://tools.ietf.org/html/rfc5019#section-6.1 (sleevi #4).
	this.ocspMu.RLock()
	defer this.ocspMu.RUnlock()
	if time.Now().After(this.ocspUpdateAfter) {
		// TODO(twifkak): Use a logging framework with support for debug-only statements.
		log.Println("Updating OCSP; expired by HTTP cache headers: ", this.ocspUpdateAfter)
		return true
	}
	return false
}

func (this *CertCache) findIssuer() *x509.Certificate {
	issuerName := this.certs[0].Issuer
	for _, cert := range this.certs {
		// The subject name is guaranteed to match the issuer name per
		// https://tools.ietf.org/html/rfc3280#section-4.1.2.4 and
		// #section-4.1.2.6. (The latter guarantees that the subject
		// name will be in the subject field and not the subjectAltName
		// field for CAs.)
		//
		// However, the definition of "match" is more complicated. The
		// general "Name matching" algorithm is defined in
		// https://www.itu.int/rec/T-REC-X.501-201610-I/en. However,
		// RFC3280 defines a subset, and pkix.Name.String() defines an
		// ad hoc canonical serialization (as opposed to
		// https://tools.ietf.org/html/rfc1779 which has many forms),
		// such that comparing the two strings should be sufficient.
		if cert.Subject.String() == issuerName.String() {
			return cert
		}
	}
	return nil
}

func (this *CertCache) updateOCSP(orig []byte, ocspUpdateAfter *time.Time) []byte {
	issuer := this.findIssuer()
	if issuer == nil {
		log.Println("Cannot find issuer certificate in CertFile.")
		return orig
	}

	// The default SHA1 hash function is mandated by the Lightweight OCSP
	// Profile, https://tools.ietf.org/html/rfc5019 2.1.1 (sleevi #4).
	req, err := ocsp.CreateRequest(this.certs[0], issuer, nil)
	if err != nil {
		log.Println("Error creating OCSP request:", err)
		return orig
	}

	var ocspServer string
	if fakeOCSPServer != "" {
		ocspServer = fakeOCSPServer
	} else if len(this.certs[0].OCSPServer) < 1 {
		log.Println("Cert missing OCSPServer.")
		return orig
	} else {
		// This is a URI, per https://tools.ietf.org/html/rfc5280#section-4.2.2.1.
		ocspServer = this.certs[0].OCSPServer[0]
	}

	// Conform to the Lightweight OCSP Profile, by preferring GET over POST
	// if the request is small enough (sleevi #4).
	// https://tools.ietf.org/html/rfc2560#appendix-A.1.1 describes how the
	// URL should be formed.
	// https://tools.ietf.org/html/rfc5019#section-5 shows an example where
	// the base64 encoding includes '/' and '=' (and therefore should be
	// StdEncoding).
	getURL := ocspServer + "/" + url.PathEscape(base64.StdEncoding.EncodeToString(req))
	var httpReq *http.Request
	if len(getURL) <= 255 {
		httpReq, err = http.NewRequest("GET", getURL, nil)
		if err != nil {
			log.Println("Error creating OCSP response:", err)
			return orig
		}
	} else {
		httpReq, err = http.NewRequest("POST", ocspServer, bytes.NewReader(req))
		if err != nil {
			log.Println("Error creating OCSP response:", err)
			return orig
		}
		httpReq.Header.Set("Content-Type", "application/ocsp-request")
	}

	httpResp, err := this.client.Do(httpReq)
	if err != nil {
		log.Println("Error issuing OCSP request:", err)
		return orig
	}
	if httpResp.Body != nil {
		defer httpResp.Body.Close()
	}

	// If cache-control headers indicate a response that is not ever
	// cacheable, then ignore them. Otherwise, allow them to indicate an
	// expiry earlier than we'd usually follow.
	if fakeOCSPExpiry != nil {
		*ocspUpdateAfter = *fakeOCSPExpiry
	} else {
		reasons, expiry, err := cachecontrol.CachableResponse(httpReq, httpResp, cachecontrol.Options{PrivateCache: true})
		if len(reasons) > 0 || err != nil {
			*ocspUpdateAfter = infiniteFuture
		} else {
			*ocspUpdateAfter = expiry
		}
	}

	respBytes, err := ioutil.ReadAll(io.LimitReader(httpResp.Body, 1024*1024))
	if err != nil {
		log.Println("Error reading OCSP response:", err)
		return orig
	}
	// sleevi #2 (also required by RFC5019 2.2.2):
	resp, err := ocsp.ParseResponseForCert(respBytes, this.certs[0], issuer)
	if err != nil {
		log.Println("Error parsing OCSP response:", err)
		return orig
	}
	if resp.Status != ocsp.Good {
		log.Println("Invalid OCSP status:", resp.Status)
		return orig
	}
	if resp.ThisUpdate.After(time.Now()) {
		log.Println("OCSP thisUpdate in the future:", resp.ThisUpdate)
		return orig
	}
	if resp.NextUpdate.Before(time.Now()) {
		log.Println("OCSP nextUpdate in the past:", resp.NextUpdate)
		return orig
	}
	return respBytes
}
