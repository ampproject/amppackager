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

package certcache

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
	"strconv"
	"sync"
	"time"

	"github.com/WICG/webpackage/go/signedexchange/certurl"
	"github.com/ampproject/amppackager/packager/util"
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

type CertCache struct {
	// TODO(twifkak): Support multiple cert chains (for different domains, for different roots).
	certName          string
	certs             []*x509.Certificate
	ocspUpdateAfterMu sync.RWMutex
	ocspUpdateAfter   time.Time
	// TODO(twifkak): Implement a registry of Updateable instances which can be configured in the toml.
	ocspFile Updateable
	client   http.Client

	// "Virtual methods", exposed for testing.
	// Given a certificate, returns the OCSP responder URL for that cert.
	extractOCSPServer func(*x509.Certificate) (string, error)
	// Given an HTTP request/response, returns its cache expiry.
	httpExpiry func(*http.Request, *http.Response) time.Time
}

// Must call Init() on the returned CertCache before you can use it.
func New(certs []*x509.Certificate, ocspCache string) *CertCache {
	return &CertCache{
		certName:        util.CertName(certs[0]),
		certs:           certs,
		ocspUpdateAfter: infiniteFuture, // Default, in case initial readOCSP successfully loads from disk.
		// Distributed OCSP cache to support the following sleevi requirements:
		// 1. Support for keeping a long-lived (disk) cache of OCSP responses.
		//    This should be fairly simple. Any restarting of the service
		//    shouldn't blow away previous responses that were obtained.
		// 6. Distributed or proxiable fetching
		//    ... there may be thousands of FE servers, all with the same
		//    certificate, all needing to staple an OCSP response. You don't
		//    want to have all of them hammering the OCSP server - ideally,
		//    you'd have one request, in the backend, and updating them all.
		ocspFile: &Chained{first: &InMemory{}, second: &LocalFile{path: ocspCache}},
		client:   http.Client{Timeout: 60 * time.Second},
		extractOCSPServer: func(cert *x509.Certificate) (string, error) {
			if len(cert.OCSPServer) < 1 {
				return "", errors.New("Cert missing OCSPServer.")
			}
			// This is a URI, per https://tools.ietf.org/html/rfc5280#section-4.2.2.1.
			return cert.OCSPServer[0], nil
		},
		httpExpiry: func(req *http.Request, resp *http.Response) time.Time {
			reasons, expiry, err := cachecontrol.CachableResponse(req, resp, cachecontrol.Options{PrivateCache: true})
			if len(reasons) > 0 || err != nil {
				return infiniteFuture
			} else {
				return expiry
			}
		},
	}
}

func (this *CertCache) Init(stop chan struct{}) error {
	// Prime the OCSP disk and memory cache, so we can start serving immediately.
	_, _, err := this.readOCSP()
	if err != nil {
		return errors.Wrap(err, "initializing CertCache")
	}
	// Update OCSP in the background, per sleevi requirements:
	// 3. Refreshes the response, in the background, with sufficient time before expiration.
	//    A rule of thumb would be to fetch at notBefore + (notAfter -
	//    notBefore) / 2, which is saying "start fetching halfway through
	//    the validity period". You want to be able to handle situations
	//    like the OCSP responder giving you junk, but also sufficient time
	//    to raise an alert if something has gone really wrong.
	// 7. The ability to serve old responses while fetching new responses.
	go this.maintainOCSP(stop)
	return nil
}

func (this *CertCache) createCertChainCBOR(ocsp []byte) ([]byte, error) {
	certChain := make(certurl.CertChain, len(this.certs))
	for i, cert := range this.certs {
		certChain[i] = &certurl.CertChainItem{Cert: cert}
	}
	certChain[0].OCSPResponse = ocsp

	var buf bytes.Buffer
	err := certChain.Write(&buf)
	if err != nil {
		return nil, errors.Wrap(err, "Error writing cert chain")
	}
	return buf.Bytes(), nil
}

func (this *CertCache) ocspMidpoint(bytes []byte, issuer *x509.Certificate) (time.Time, error) {
	resp, err := ocsp.ParseResponseForCert(bytes, this.certs[0], issuer)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "Parsing OCSP")
	}
	return resp.ThisUpdate.Add(resp.NextUpdate.Sub(resp.ThisUpdate) / 2), nil
}

func (this *CertCache) ServeHTTP(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if params.ByName("certName") == this.certName {
		// https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-3.3
		// This content-type is not standard, but included to reduce
		// the chance that faulty user agents employ content sniffing.
		resp.Header().Set("Content-Type", "application/cert-chain+cbor")
		// Instruct the intermediary to reload this cert-chain at the
		// OCSP midpoint, in case it cannot parse it.
		ocsp, _, err := this.readOCSP()
		if err != nil {
			util.NewHTTPError(http.StatusInternalServerError, "Error reading OCSP: ", err).LogAndRespond(resp)
			return
		}
		midpoint, err := this.ocspMidpoint(ocsp, this.findIssuer())
		if err != nil {
			util.NewHTTPError(http.StatusInternalServerError, "Error computing OCSP midpoint: ", err).LogAndRespond(resp)
			return
		}
		// int is large enough to represent 24855 days in seconds.
		expiry := int(midpoint.Sub(time.Now()).Seconds())
		if expiry < 0 {
			expiry = 0
		}
		resp.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(expiry))
		resp.Header().Set("ETag", "\""+this.certName+"\"")
		resp.Header().Set("X-Content-Type-Options", "nosniff")
		cbor, err := this.createCertChainCBOR(ocsp)
		if err != nil {
			util.NewHTTPError(http.StatusInternalServerError, "Error building cert chain: ", err).LogAndRespond(resp)
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
// unsigned. This is per sleevi requirement:
// 8. Some idea of what to do when "things go bad".
//    What happens when it's been 7 days, no new OCSP response can be obtained,
//    and the current response is about to expire?
func (this *CertCache) IsHealthy() bool {
	ocsp, _, err := this.readOCSP()
	return err != nil || this.isHealthy(ocsp)
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

// Returns the OCSP response and expiry, refreshing if necessary.
func (this *CertCache) readOCSP() ([]byte, time.Time, error) {
	var ocspUpdateAfter time.Time
	ocsp, err := this.ocspFile.Read(context.Background(), this.shouldUpdateOCSP, func(orig []byte) []byte {
		return this.fetchOCSP(orig, &ocspUpdateAfter)
	})
	if err != nil {
		return nil, time.Time{}, errors.Wrap(err, "Updating OCSP cache")
	}
	if len(ocsp) == 0 {
		return nil, time.Time{}, errors.New("Missing OCSP response.")
	}
	if !this.isHealthy(ocsp) {
		return nil, time.Time{}, errors.New("OCSP failed health check.")
	}
	this.ocspUpdateAfterMu.Lock()
	defer this.ocspUpdateAfterMu.Unlock()
	if !ocspUpdateAfter.Equal(time.Time{}) {
		// fetchOCSP was called, and therefore a new HTTP cache expiry was set.
		// TODO(twifkak): Write this to disk, so any replica can pick it up.
		this.ocspUpdateAfter = ocspUpdateAfter
	}
	return ocsp, ocspUpdateAfter, nil
}

// Checks for OCSP updates every hour. Never terminates.
func (this *CertCache) maintainOCSP(stop chan struct{}) {
	// Only make one request per ocspCheckInterval, to minimize the impact
	// on OCSP servers that are buckling under load, per sleevi requirement:
	// 5. As with any system doing background requests on a remote server,
	//    don't be a jerk and hammer the server when things are bad...
	//    sometimes servers and networks have issues. When a[n OCSP client]
	//    has trouble getting a request, hopefully it does something
	//    smarter than just retry in a busy loop, hammering the OCSP server
	//    into further oblivion.
	ticker := time.NewTicker(ocspCheckInterval)

	for {
		select {
		case <-ticker.C:
			_, _, err := this.readOCSP()
			if err != nil {
				log.Println("Warning: OCSP update failed. Cached response may expire:", err)
			}
		case <-stop:
			ticker.Stop()
			return
		}
	}
}

// Returns true if OCSP is expired (or near enough).
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
	// Compute the midpoint per sleevi #3 (see above).
	midpoint, err := this.ocspMidpoint(bytes, issuer)
	if err != nil {
		log.Println("Error computing OCSP midpoint:", err)
		return true
	}
	if time.Now().After(midpoint) {
		// TODO(twifkak): Use a logging framework with support for debug-only statements.
		log.Println("Updating OCSP; after midpoint: ", midpoint)
		return true
	}
	// Allow cache-control headers to indicate an earlier update time, per
	// https://tools.ietf.org/html/rfc5019#section-6.1, per sleevi requirement:
	// 4. ... such a system should observe the Lightweight OCSP Profile of
	//    RFC 5019. This more or less boils down to "Use GET requests whenever
	//    possible, and observe HTTP cache semantics."
	this.ocspUpdateAfterMu.RLock()
	defer this.ocspUpdateAfterMu.RUnlock()
	if time.Now().After(this.ocspUpdateAfter) {
		// TODO(twifkak): Use a logging framework with support for debug-only statements.
		log.Println("Updating OCSP; expired by HTTP cache headers: ", this.ocspUpdateAfter)
		return true
	}
	// TODO(twifkak): Use a logging framework with support for debug-only statements.
	log.Println("No update necessary.")
	return false
}

// Finds the issuer of this cert (i.e. the second from the bottom of the
// chain).
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

// Queries the OCSP responder for this cert and return the OCSP response.
func (this *CertCache) fetchOCSP(orig []byte, ocspUpdateAfter *time.Time) []byte {
	issuer := this.findIssuer()
	if issuer == nil {
		log.Println("Cannot find issuer certificate in CertFile.")
		return orig
	}

	// The default SHA1 hash function is mandated by the Lightweight OCSP
	// Profile, https://tools.ietf.org/html/rfc5019 2.1.1 (sleevi #4, see above).
	req, err := ocsp.CreateRequest(this.certs[0], issuer, nil)
	if err != nil {
		log.Println("Error creating OCSP request:", err)
		return orig
	}

	ocspServer, err := this.extractOCSPServer(this.certs[0])
	if err != nil {
		log.Println("Error extracting OCSP server:", err)
		return orig
	}

	// Conform to the Lightweight OCSP Profile, by preferring GET over POST
	// if the request is small enough (sleevi #4, see above).
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
	*ocspUpdateAfter = this.httpExpiry(httpReq, httpResp)

	respBytes, err := ioutil.ReadAll(io.LimitReader(httpResp.Body, 1024*1024))
	if err != nil {
		log.Println("Error reading OCSP response:", err)
		return orig
	}
	// Validate the response, per sleevi requirement:
	// 2. Validate the server responses to make sure it is something the client will accept.
	// and also per sleevi #4 (see above), as required by
	// https://tools.ietf.org/html/rfc5019#section-2.2.2.
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
	// OCSP duration must be <=7 days, per
	// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cross-origin-trust.
	// Serving these responses may cause UAs to reject the SXG.
	if resp.NextUpdate.Sub(resp.ThisUpdate) > time.Hour * 24 * 7 {
		log.Printf("OCSP nextUpdate %+v too far ahead of thisUpdate %+v\n", resp.NextUpdate, resp.ThisUpdate)
		return orig
	}
	return respBytes
}
