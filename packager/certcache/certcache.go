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
	"crypto"
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
	"github.com/ampproject/amppackager/packager/certfetcher"
	"github.com/ampproject/amppackager/packager/certloader"
	"github.com/ampproject/amppackager/packager/mux"
	"github.com/ampproject/amppackager/packager/util"
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

// How often to check if certs needs updating.
const certCheckInterval = 24 * time.Hour

// Max number of OCSP request tries.
// This will timeout after 1 + 2 + 4 + 8 + 10 * 6 = 75 minutes.
const maxOCSPTries = 10

// Recommended renewal duration for certs. This is duration before next cert expiry.
// 8 days is recommended duration to start requesting new certs to allow for ACME server outages.
// It's 6 days + 2 days renewal grace period.
// 6 days so that generated SXGs are valid for their full lifetime, plus 2 days in front of that to allow time for the new cert
// to be obtained.
// TODO(banaag): make 2 days renewal grace period configurable in toml.
const certRenewalInterval = 8 * 24 * time.Hour

type OCSPResponder func(*x509.Certificate) ([]byte, error)

type CertHandler interface {
	GetLatestCert() *x509.Certificate
	IsHealthy() error
}

type CertCache struct {
	// TODO(twifkak): Support multiple cert chains (for different domains, for different roots).
	certName string
	certsMu  sync.RWMutex
	certs    []*x509.Certificate
	// If certFetcher is not set, that means cert auto-renewal is not available.
	certFetcher       *certfetcher.CertFetcher
	renewedCertsMu    sync.RWMutex
	renewedCertName   string
	renewedCerts      []*x509.Certificate
	ocspUpdateAfterMu sync.RWMutex
	ocspUpdateAfter   time.Time
	stop              chan struct{}
	// TODO(twifkak): Implement a registry of Updateable instances which can be configured in the toml.
	ocspFile     Updateable
	ocspFilePath string
	client       http.Client
	// Given a certificate, returns a current OCSP response for the cert;
	// this is a fallback, called when in development mode and there is no
	// OCSP URL.
	generateOCSPResponse OCSPResponder
	// Domains to validate
	Domains     []string
	CertFile    string
	NewCertFile string
	// Is CertCache initialized to do cert renewal or OCSP refreshes?
	isInitialized bool

	// "Virtual methods", exposed for testing.
	// Given a certificate, returns the OCSP responder URL for that cert.
	extractOCSPServer func(*x509.Certificate) (string, error)
	// Given an HTTP request/response, returns its cache expiry.
	httpExpiry func(*http.Request, *http.Response) time.Time
}

// Callers need to call Init() on the returned CertCache before the cache can auto-renew certs.
// Callers can use the uninitialized CertCache for testing certificates (without doing OCSP or
// cert refreshes).
//
// TODO(banaag): per gregable@ comments:
// The long argument list makes the callsites tricky to read and easy to get wrong, especially if several of the arguments have the same type.
//
// An alternative pattern would be to create an IsInitialized() bool or similarly named function that verifies all of the required fields have
// been set. Then callers can just set fields in the struct by name and assert IsInitialized before doing anything with it.
func New(certs []*x509.Certificate, certFetcher *certfetcher.CertFetcher, domains []string,
	certFile string, newCertFile string, ocspCache string, generateOCSPResponse OCSPResponder) *CertCache {
	certName := ""
	if len(certs) > 0 && certs[0] != nil {
		certName = util.CertName(certs[0])
	}
	return &CertCache{
		certName:        certName,
		certs:           certs,
		certFetcher:     certFetcher,
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
		ocspFile:     &Chained{first: &InMemory{}, second: &LocalFile{path: ocspCache}},
		ocspFilePath: ocspCache,
		stop:         make(chan struct{}),
		generateOCSPResponse: generateOCSPResponse,
		client:       http.Client{Timeout: 60 * time.Second},
		extractOCSPServer: func(cert *x509.Certificate) (string, error) {
			if cert == nil || len(cert.OCSPServer) < 1 {
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
		Domains:       domains,
		CertFile:      certFile,
		NewCertFile:   newCertFile,
		isInitialized: false,
	}
}

func (this *CertCache) Init() error {
	this.updateCertIfNecessary()

	// Prime the OCSP disk and memory cache, so we can start serving immediately.
	_, _, err := this.readOCSP(true)
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
	go this.maintainOCSP()

	if this.certFetcher != nil {
		// Update Certs in the background.
		go this.maintainCerts()
	}

	this.isInitialized = true

	return nil
}

// Stop stops the goroutines spawned in Init, which are automatically updating the certificate and the OCSP response.
// It returns true if the call actually stops them, false if they have already been stopped.
func (this *CertCache) Stop() bool {
	select {
	// this.stop will never be used for sending a value. Thus this case matches only when it has already been closed.
	case <-this.stop:
		return false
	default:
		close(this.stop)
		return true
	}
}

// Gets the latest cert.
// Returns the current cert if the cache has not been initialized or if the certFetcher is not set (good for testing)
// If cert is invalid, it will attempt to renew.
// If cert is still valid, returns the current cert.
func (this *CertCache) GetLatestCert() *x509.Certificate {
	if !this.isInitialized || this.certFetcher == nil {
		// If certcache is not initialized or certFetcher is not set,
		// just return cert without checking if it needs auto-renewal.
		return this.getCert()
	}

	if !this.hasCert() {
		return nil
	}

	d, err := util.GetDurationToExpiry(this.getCert(), time.Now())
	if err != nil {
		// Current cert is already invalid. Check if renewal is available.
		log.Println("Current cert is expired, attempting to renew: ", err)
		this.updateCertIfNecessary()
		return this.getCert()
	}
	if d >= time.Duration(certRenewalInterval) {
		// Cert is still valid.
		return this.getCert()
	} else if d < time.Duration(certRenewalInterval) {
		// Cert is still valid, but we need to start process of requesting new cert.
		log.Println("Current cert is close to expiry threshold, attempting to renew in the background.")
		return this.getCert()
	}
	return nil
}

func (this *CertCache) createCertChainCBOR(ocsp []byte) ([]byte, error) {
	this.certsMu.RLock()
	defer this.certsMu.RUnlock()

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

func (this *CertCache) parseOCSP(bytes []byte, issuer *x509.Certificate) (*ocsp.Response, error){
	resp, err := ocsp.ParseResponseForCert(bytes, this.getCert(), issuer)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing OCSP")
	}
	return resp, nil
}

func (this *CertCache) ocspMidpoint(resp *ocsp.Response) (time.Time) {
	return resp.ThisUpdate.Add(resp.NextUpdate.Sub(resp.ThisUpdate) / 2)
}

func (this *CertCache) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	params := mux.Params(req)

	// RLock for the certName
	this.certsMu.RLock()
	defer this.certsMu.RUnlock()
	if params["certName"] == this.certName {
		// https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-3.3
		// This content-type is not standard, but included to reduce
		// the chance that faulty user agents employ content sniffing.
		resp.Header().Set("Content-Type", "application/cert-chain+cbor")
		// Instruct the intermediary to reload this cert-chain at the
		// OCSP midpoint, in case it cannot parse it.
		ocsp, _, err := this.readOCSP(false)
		if err != nil {
			util.NewHTTPError(http.StatusInternalServerError, "Error reading OCSP: ", err).LogAndRespond(resp)
			return
		}
		ocspResp, err := this.parseOCSP(ocsp, this.findIssuer())
		if err != nil {
			log.Println("Invalid OCSP:", err)
			util.NewHTTPError(http.StatusInternalServerError, "Invalid OCSP: ", err).LogAndRespond(resp)
			return
		}
		midpoint := this.ocspMidpoint(ocspResp)
		// int is large enough to represent 24855 days in seconds.
		expiry := int(midpoint.Sub(time.Now()).Seconds())
		if expiry < 0 {
			expiry = 0
		}
		resp.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(expiry))
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
func (this *CertCache) IsHealthy() error {
	ocsp, _, errorOCSP := this.readOCSP(false)
	if errorOCSP != nil {
		return errorOCSP
	}
	errorHealth := this.isHealthy(ocsp)
	if errorHealth != nil {
		return errorHealth
	}
	return nil
}

func (this *CertCache) isHealthy(ocspResp []byte) error {
	if ocspResp == nil {
		return errors.New("OCSP response not yet fetched.")
	}
	issuer := this.findIssuer()
	if issuer == nil {
		return errors.New("Cannot find issuer certificate in CertFile.")
	}
	resp, err := ocsp.ParseResponseForCert(ocspResp, this.getCert(), issuer)
	if err != nil {
		return errors.Wrap(err, "Error parsing OCSP response")
	}
	if resp.NextUpdate.Before(time.Now()) {
		return errors.Errorf("Cached OCSP is stale, NextUpdate: %v", resp.NextUpdate)
	}
	return nil
}

func (this *CertCache) readOCSPHelper(numTries int, exhaustedRetries bool) ([]byte, time.Time, error) {
	var ocspUpdateAfter time.Time

	this.certsMu.RLock()
	defer this.certsMu.RUnlock()
	ocsp, err := this.ocspFile.Read(context.Background(), this.shouldUpdateOCSP, func(orig []byte) []byte {
		return this.fetchOCSP(orig, this.certs, &ocspUpdateAfter, numTries > 0)
	})
	if err != nil {
		if exhaustedRetries {
			return nil, time.Time{}, errors.Wrap(err, "Updating OCSP cache")
		} else {
			return nil, time.Time{}, nil
		}
	}
	if len(ocsp) == 0 {
		if exhaustedRetries {
			return nil, time.Time{}, errors.New("Missing OCSP response.")
		} else {
			return nil, time.Time{}, nil
		}
	}
	if err := this.isHealthy(ocsp); err != nil {
		if exhaustedRetries {
			return nil, time.Time{}, errors.Wrap(err, "OCSP failed health check")
		} else {
			return nil, time.Time{}, nil
		}
	}

	return ocsp, ocspUpdateAfter, nil
}

// Returns the OCSP response and expiry, refreshing if necessary.
func (this *CertCache) readOCSP(allowRetries bool) ([]byte, time.Time, error) {
	var ocspUpdateAfter time.Time
	var err error
	var maxTries int

	ocsp := []byte(nil)
	waitTimeInMinutes := 1
	if !allowRetries || this.certFetcher == nil {
		// If certFetcher is nil, that means we are not auto-renewing so don't retry OCSP.
		maxTries = 1
	} else {
		maxTries = maxOCSPTries
	}

	for numTries := 0; numTries < maxTries; {
		ocsp, ocspUpdateAfter, err = this.readOCSPHelper(numTries, numTries >= maxTries - 1)
		if err != nil {
			return nil, ocspUpdateAfter, err
		}
		if !this.shouldUpdateOCSP(ocsp) {
			break;
		}
		// Wait only if are not on our last try.
		if numTries < maxTries - 1 {
			waitTimeInMinutes = waitForSpecifiedTime(waitTimeInMinutes, numTries)
		}
		numTries++
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

// Print # of retries, wait for specified time and returned updated wait time.
func waitForSpecifiedTime(waitTimeInMinutes int, numRetries int) int {
	log.Printf("Retrying OCSP server: retry #%d\n", numRetries)
	// Wait using exponential backoff.
	log.Printf("Waiting for %d minute(s)\n", waitTimeInMinutes)
	waitTimeDuration := time.Duration(waitTimeInMinutes) * time.Minute
	// For exponential backoff.
	newWaitTimeInMinutes := 2 * waitTimeInMinutes
	if newWaitTimeInMinutes > 10 {
		// Cap the wait time at 10 minutes.
		newWaitTimeInMinutes = 10
	}
	time.Sleep(waitTimeDuration)
	return newWaitTimeInMinutes
}

// Checks for OCSP updates every hour. Terminates only when stop receives
// a message.
func (this *CertCache) maintainOCSP() {
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
			_, _, err := this.readOCSP(true)
			if err != nil {
				log.Println("Warning: OCSP update failed. Cached response may expire:", err)
			}
		case <-this.stop:
			ticker.Stop()
			return
		}
	}
}

// Returns true if OCSP is expired (or near enough).
func (this *CertCache) shouldUpdateOCSP(ocsp []byte) bool {
	if len(ocsp) == 0 {
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
	ocspResp, err := this.parseOCSP(ocsp, issuer)
	if err != nil {
		// An old ocsp cache causes a parse error in case of cert renewal. Do not log it.
		if this.isInitialized {
			log.Println("Invalid OCSP:", err)
		}
		return true
	}
	// Compute the midpoint per sleevi #3 (see above).
	midpoint := this.ocspMidpoint(ocspResp)
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
	log.Println("No OCSP update necessary.")
	return false
}

// Finds the issuer of this cert (i.e. the second from the bottom of the
// chain).
func (this *CertCache) findIssuer() *x509.Certificate {
	if !this.hasCert() {
		return nil
	}
	return this.findIssuerUsingCerts(this.certs)
}

// Finds the issuer of the specified cert (i.e. the second from the bottom of the
// chain).
func (this *CertCache) findIssuerUsingCerts(certs []*x509.Certificate) *x509.Certificate {
	if certs == nil || len(certs) == 0 {
		return nil
	}
	this.certsMu.RLock()
	defer this.certsMu.RUnlock()
	issuerName := certs[0].Issuer
	for _, cert := range certs {
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
func (this *CertCache) fetchOCSP(orig []byte, certs []*x509.Certificate, ocspUpdateAfter *time.Time, isRetry bool) []byte {
	issuer := this.findIssuerUsingCerts(certs)
	if issuer == nil {
		log.Println("Cannot find issuer certificate in CertFile.")
		return orig
	}
	// The default SHA1 hash function is mandated by the Lightweight OCSP
	// Profile, https://tools.ietf.org/html/rfc5019 2.1.1 (sleevi #4, see above).
	req, err := ocsp.CreateRequest(certs[0], issuer, nil)
	if err != nil {
		log.Println("Error creating OCSP request:", err)
		return orig
	}

	ocspServer, err := this.extractOCSPServer(certs[0])
	if err != nil {
		if this.generateOCSPResponse == nil {
			log.Println("Error extracting OCSP server:", err)
			return orig
		}
		log.Println("Cert lacks OCSP URL; using fake OCSP in development mode.")
		resp, err := this.generateOCSPResponse(certs[0])
		if err != nil {
			log.Println("error generating fake OCSP response:", err)
			return orig
		}
		return resp
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
	// Logic is a fallback, due to some CAs not responding as expected to a GET.
	if len(getURL) <= 255 && !isRetry {
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

	respBytes, err := ioutil.ReadAll(io.LimitReader(httpResp.Body, maxOCSPResponseBytes))
	if err != nil {
		log.Println("Error reading OCSP response:", err)
		return orig
	}

	// Validate the response, per sleevi requirement:
	// 2. Validate the server responses to make sure it is something the client will accept.
	// and also per sleevi #4 (see above), as required by
	// https://tools.ietf.org/html/rfc5019#section-2.2.2.
	resp, err := ocsp.ParseResponseForCert(respBytes, certs[0], issuer)
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
	if resp.NextUpdate.Sub(resp.ThisUpdate) > time.Hour*24*7 {
		log.Printf("OCSP nextUpdate %+v too far ahead of thisUpdate %+v\n", resp.NextUpdate, resp.ThisUpdate)
		return orig
	}
	return respBytes
}

// Checks for cert updates every certCheckInterval hours. Terminates only when stop
// receives a message.
func (this *CertCache) maintainCerts() {
	// Only make one request per certCheckInterval, to minimize the impact
	// on servers that are buckling under load.
	ticker := time.NewTicker(certCheckInterval)

	for {
		select {
		case <-ticker.C:
			this.updateCertIfNecessary()
		case <-this.stop:
			ticker.Stop()
			return
		}
	}
}

// Returns true iff cert cache contains at least 1 cert.
func (this *CertCache) hasCert() bool {
	this.certsMu.RLock()
	defer this.certsMu.RUnlock()
	return len(this.certs) > 0 && this.certs[0] != nil
}

func (this *CertCache) getCert() *x509.Certificate {
	if !this.hasCert() {
		return nil
	}
	this.certsMu.RLock()
	defer this.certsMu.RUnlock()
	return this.certs[0]
}

// Returns true iff cert cache renewal contains at least 1 cert.
func (this *CertCache) hasRenewalCert() bool {
	this.renewedCertsMu.RLock()
	defer this.renewedCertsMu.RUnlock()
	return len(this.renewedCerts) > 0 && this.renewedCerts[0] != nil
}

func (this *CertCache) getRenewalCert() *x509.Certificate {
	this.renewedCertsMu.RLock()
	defer this.renewedCertsMu.RUnlock()
	if !this.hasRenewalCert() {
		return nil
	}
	return this.renewedCerts[0]
}

// Set current cert with mutex protection.
func (this *CertCache) setCerts(certs []*x509.Certificate) {
	this.certsMu.Lock()
	defer this.certsMu.Unlock()
	this.certs = certs
	this.certName = util.CertName(certs[0])

	err := certloader.WriteCertsToFile(this.certs, this.CertFile)
	if err != nil {
		log.Printf("Unable to write certs to file: %s", this.CertFile)
	}

	// Purge OCSP cache
	certloader.RemoveFile(this.ocspFilePath)
}

// Set new cert with mutex protection.
func (this *CertCache) setNewCerts(certs []*x509.Certificate) {
	this.renewedCertsMu.Lock()
	defer this.renewedCertsMu.Unlock()
	this.renewedCerts = certs

	if this.renewedCerts == nil {
		this.renewedCertName = ""
		err := certloader.RemoveFile(this.NewCertFile)
		if err != nil {
			log.Printf("Unable to remove file: %s", this.NewCertFile)
		}
		return
	}
	this.renewedCertName = util.CertName(certs[0])

	err := certloader.WriteCertsToFile(this.renewedCerts, this.NewCertFile)
	if err != nil {
		log.Printf("Unable to write certs to file: %s", this.NewCertFile)
	}
}

// Update the cert in the cache if necessary.
func (this *CertCache) updateCertIfNecessary() {
	log.Println("Updating cert if necessary")
	if this.certFetcher == nil {
		// Don't request new certs from CA if certFetcher is not set. This means this instance of the amppackager
		// is not in autorenewcert mode. Just make an attempt at reading the cert saved on disk to see if
		// another amppackager instance that is in autorenewcert mode actually updated it with a valid cert.
		log.Println("Certfetcher is not set, skipping cert updates. Checking cert on disk if updated.")
		this.reloadCertIfExpired()
		return
	}
	d := time.Duration(0)
	err := errors.New("")
	if this.hasCert() {
		d, err = util.GetDurationToExpiry(this.getCert(), time.Now())
	}
	if err != nil {
		this.renewedCertsMu.Lock()
		defer this.renewedCertsMu.Unlock()

		// Current cert is already invalid, check if we have a pending renewal cert.
		if this.renewedCerts != nil {
			// If renewedCerts is set, copy that over to certs
			// and set renewedCerts to nil.
			this.setCerts(this.renewedCerts)
			this.setNewCerts(nil)
			return
		}
		// Current cert is already invalid. Try refreshing.
		log.Println("Warning current cert is expired, attempting to renew: ", err)
		certs, err := this.certFetcher.FetchNewCert()
		if err != nil {
			log.Println("Error trying to fetch new certificates from CA: ", err)
			return
		}
		this.setCerts(certs)
		return
	}
	if d >= time.Duration(certRenewalInterval) {
		// Cert is still valid, don't do anything.
	} else if d < time.Duration(certRenewalInterval) {
		this.renewedCertsMu.Lock()
		defer this.renewedCertsMu.Unlock()

		// Check if we already have a renewal cert waiting, fetch a new cert if not.
		if this.renewedCerts == nil {
			// Cert is still valid, but we need to start process of requesting new cert.
			log.Println("Warning: Current cert crossed threshold for renewal, attempting to renew.")
			certs, err := this.certFetcher.FetchNewCert()
			if err != nil {
				log.Println("Error trying to fetch new certificates from CA: ", err)
				return
			}
			this.setNewCerts(certs)
		} else {
			// TODO(banaag) from twifkak comments:
			// Note that this logic works, albeit might fail to fetch OCSP the first try, but will succeed 24 hours later.
			//
			// I realize it's difficult to use readOCSP here, since it's hard-coded to use this.certs and friends, rather than
			// this.renewedCerts and friends. That makes me think two things:
			//
			// The extraction of the retry logic from readOCSP would be useful.
			// We should bundle certName, certs, certsMu, ocspFile, ocspFilePath, ocspUpdateAfter, and ocspUpdateAfterMu into
			// a new struct type, and then have two copies of that in certcache - one for current certs and one for new certs.
			var ocspUpdateAfter time.Time

			ocsp, _, errorOCSP := this.readOCSP(true)
			if errorOCSP != nil {
				newOCSP := this.fetchOCSP(ocsp, this.renewedCerts, &ocspUpdateAfter, false)
				// Check if newOCSP != ocsp and that there are no errors, health-wise with new ocsp.
				if !bytes.Equal(newOCSP, ocsp) && this.isHealthy(newOCSP) == nil {
					// We were able to fetch new OCSP with renewal cert, time to switch to new certs.
					this.setCerts(this.renewedCerts)
					this.setNewCerts(nil)
				}
			}
		}
	}
}

func (this *CertCache) doesCertNeedReloading() bool {
	if !this.hasCert() { return true }
	d, err := util.GetDurationToExpiry(this.getCert(), time.Now())
	return err != nil || d < certRenewalInterval
}

func (this *CertCache) reloadCertIfExpired() {
	if !this.doesCertNeedReloading() {
		return
	}

	// If we get to here, the cert was either expired, or it's time to renew.
	// We always validate the certs here.  If we are in development mode and the certs don't validate,
	// it doesn't matter because the old certs won't be overridden (and the old certs are probably invalid, too).
	certs, err := certloader.LoadAndValidateCertsFromFile(this.CertFile, true)
	if err != nil {
		log.Println(errors.Wrap(err, "Can't load cert file"))
		certs = nil
	}
	if certs != nil {
		this.setCerts(certs)
	}

	newCerts, err := certloader.LoadAndValidateCertsFromFile(this.NewCertFile, true)
	if err != nil {
		log.Println(errors.Wrap(err, "Can't load new cert file"))
		newCerts = nil
	}
	if newCerts != nil {
		this.setNewCerts(newCerts)
	}
}

// Creates cert cache by loading certs and keys from disk, doing validation
// and populating the cert cache with current set of certificate related information.
// If development mode is true, prints a warning for certs that can't sign HTTP exchanges.
func PopulateCertCache(config *util.Config, key crypto.PrivateKey, generateOCSPResponse OCSPResponder,
	developmentMode bool, autoRenewCert bool) (*CertCache, error) {

	if config.CertFile == "" {
		return nil, errors.New("Missing cert file path in config.")
	}

	if autoRenewCert && config.NewCertFile == "" {
		return nil, errors.New("Missing new cert file path in config.")
	}

	certs, err := certloader.LoadCertsFromFile(config, developmentMode)
	if err != nil {
		log.Println(errors.Wrap(err, "Can't load cert file"))
		certs = nil
	}
	domain := ""
	for _, urlSet := range config.URLSet {
		domain = urlSet.Sign.Domain
		if certs != nil {
			if err := util.CertificateMatches(certs[0], key, domain); err != nil {
				return nil, errors.Wrapf(err, "checking %s", config.CertFile)
			}
		}
	}

	certFetcher, err := certloader.CreateCertFetcher(config, key, domain, developmentMode, autoRenewCert)
	if err != nil {
		return nil, errors.Wrap(err, "creating cert fetcher from config")
	}
	certCache := New(certs, certFetcher, []string{domain}, config.CertFile, config.NewCertFile, config.OCSPCache, generateOCSPResponse)

	return certCache, nil
}
