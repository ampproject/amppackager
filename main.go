package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pquerna/cachecontrol"
	// TODO(twifkak): github.com/pelletier/go-toml (chosen per https://github.com/golang/dep/issues/119)
	"github.com/nyaxt/webpackage/go/signedexchange"
)

// Advised against, per
// https://jyasskin.github.io/webpackage/implementation-draft/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#rfc.section.4.1,
// and blocked in http://crrev.com/c/958945.
var statefulResponseHeaders = map[string]bool{
	"Authentication-Control":    true,
	"Authentication-Info":       true,
	"Optional-WWW-Authenticate": true,
	"Proxy-Authenticate":        true,
	"Proxy-Authentication-Info": true,
	"Sec-WebSocket-Accept":      true,
	"Set-Cookie":                true,
	"Set-Cookie2":               true,
	"SetProfile":                true,
	"WWW-Authenticate":          true,
}

// TODO(twifkak): Remove this restriction by allowing streamed responses from the signedexchange library.
const maxBodyLength = 1 << 10

// TODO(twifkak): What value should this be?
const miRecordSize = 4096

func hello(resp http.ResponseWriter, req *http.Request) {
	// TODO(twifkak): Write.
	resp.Header().Set("Content-Type", "text/plain")
	if req.URL.Path == "/" {
		// TODO(twifkak): Link or redirect to documentation.
		_, err := resp.Write([]byte("hello world"))
		if err != nil {
			// TODO(twifkak): Log request details.
			// TODO(twifkak): Is it worth logging these? Maybe just connection drops.
			log.Println("Error serving request:", err)
		}
	} else {
		http.NotFound(resp, req)
	}
}

type httpError struct {
	InternalMsg string
	StatusCode  int
}

func newHttpError(statusCode int, msg ...interface{}) *httpError {
	return &httpError{fmt.Sprint(msg), statusCode}
}

func (e httpError) Error() string { return e.InternalMsg }
func (e httpError) ExternalMsg() string {
	// TODO(twifkak): Prevent construction of httpErrors without an ExternalMsg.
	switch e.StatusCode {
	case http.StatusBadRequest:
		return "400 bad request"
	case http.StatusInternalServerError:
		return "500 internal server error"
	case http.StatusBadGateway:
		return "502 bad gateway"
	default:
		return ""
	}
}
func (e httpError) LogAndRespond(resp http.ResponseWriter) {
	log.Println(e.InternalMsg)
	http.Error(resp, e.ExternalMsg(), e.StatusCode)
}

// Iff there is an error, it logs, writes to resp, and returns an error.
func parseUrls(fetch string, sign string, resp http.ResponseWriter) (*url.URL, *httpError) {
	// TODO(twifkak): Validate fetch and sign against respective whitelists of hosts/URLs.
	// TODO(twifkak): Validate that the fetch URL matches the sign URL.
	if fetch == "" {
		return nil, newHttpError(http.StatusBadRequest, "fetch URL is unspecified")
	}
	if sign == "" {
		return nil, newHttpError(http.StatusBadRequest, "sign URL is unspecified")
	}
	signUrl, err := url.Parse(sign)
	if err != nil {
		return nil, newHttpError(http.StatusBadRequest, "Error parsing sign url:", err)
	}
	if !signUrl.IsAbs() {
		return nil, newHttpError(http.StatusBadRequest, "Sign url is relative")
	}
	return signUrl, nil
}

func fetchUrl(fetch string) (*http.Request, *http.Response, *httpError) {
	// TODO(twifkak): Strip non-printable characters + newlines
	// before logging any input data.
	log.Println("Fetching URL:", fetch)
	client := http.Client{
		// TODO(twifkak): Load-test and see if non-default
		// transport settings (e.g. max idle conns per host)
		// are better.
		// TODO(twifkak): Is a cookie-jar necessary for
		// cross-redirect cookies?
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, fetch, nil)
	if err != nil {
		return nil, nil, newHttpError(http.StatusInternalServerError, "Error building request")
	}
	// TODO(twifkak): Do we need to do anything special for HTTPS
	// URLs (e.g. include a list of roots, enable verification)?
	resp, err := client.Do(req)
	if err != nil {
		// TODO(twifkak): Is there a chance fetchResp.Body is
		// non-nil, and hence needs to be closed? The net/http
		// doc is unclear.
		return nil, nil, newHttpError(http.StatusBadGateway, "Error fetching")
	}
	return req, resp, nil
}

func validateFetch(req *http.Request, resp *http.Response) *httpError {
	if resp.StatusCode != http.StatusOK {
		return newHttpError(http.StatusBadGateway, "Non-OK fetch:", resp.StatusCode)
	}
	// Validate response is publicly-cacheable, per
	// https://tools.ietf.org/html/draft-yasskin-http-origin-signed-responses-03#section-6.1.
	nonCachableReasons, _, err := cachecontrol.CachableResponse(req, resp, cachecontrol.Options{PrivateCache: false})
	if err != nil {
		return newHttpError(http.StatusBadGateway, "Error parsing cache headers:", err)
	}
	if len(nonCachableReasons) > 0 {
		return newHttpError(http.StatusBadGateway, "Non-cacheable response:", nonCachableReasons)
	}
	return nil
}

// TODO(twifkak): Test this.
func packager(resp http.ResponseWriter, req *http.Request) {
	// TODO(twifkak): See if there are any other validations or
	// sanitizations that need adding.

	fetch := req.FormValue("fetch")
	sign := req.FormValue("sign")
	signUrl, httpErr := parseUrls(fetch, sign, resp)
	if httpErr != nil {
		httpErr.LogAndRespond(resp)
		return
	}

	fetchReq, fetchResp, httpErr := fetchUrl(fetch)
	if httpErr != nil {
		httpErr.LogAndRespond(resp)
		return
	}
	defer func() {
		if err := fetchResp.Body.Close(); err != nil {
			log.Println("Error closing fetchResp body:", err)
		}
	}()

	if httpErr := validateFetch(fetchReq, fetchResp); httpErr != nil {
		httpErr.LogAndRespond(resp)
		return
	}

	// TODO(twifkak): Should I be more restrictive and just
	// whitelist some response headers?
	for header, _ := range statefulResponseHeaders {
		fetchResp.Header.Del(header)
	}
	body, err := ioutil.ReadAll(io.LimitReader(fetchResp.Body, maxBodyLength))
	if err != nil {
		log.Println("Error reading body:", err)
		http.Error(resp, "502 bad gateway", http.StatusBadGateway)
		return
	}
	_, err = signedexchange.NewExchange(signUrl, http.Header{}, fetchResp.StatusCode, fetchResp.Header, body, miRecordSize)
	if err != nil {
		log.Println("Error building exchange:", err)
		http.Error(resp, "500 internal server error", http.StatusInternalServerError)
		return
	}
	//signer := signedexchange.Signer{
	//	Date        time.Time
	//	Expires     time.Time
	//	Certs       []*x509.Certificate
	//	CertUrl     *url.URL
	//	ValidityUrl *url.URL
	//	PrivKey     crypto.PrivateKey
	//	Rand        io.Reader
	//}
	// TODO(twifkak): Consider rewriting cache control headers.
	// TODO(twifkak): Construct CBOR thing.
}

// Exposes an HTTP server. Don't run this on the open internet, for at least two reasons:
//  - It exposes an API that allows people to sign any URL as any other URL.
//  - It is in cleartext.
func main() {
	// TODO(twifkak): Make log output configurable.
	// TODO(twifkak): Replace with my own ServeMux.
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(hello))
	mux.Handle("/package", http.HandlerFunc(packager))
	// TODO(twifkak): Add a basic logging intercept (or use a Go lib for this stuff).
	server := http.Server{
		// TODO(twifkak): Make this configurable.
		Addr: ":8080",
		// Don't use DefaultServeMux, per
		// https://blog.cloudflare.com/exposing-go-on-the-internet/.
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		// If needing to stream the response, disable WriteTimeout and
		// use TimeoutHandler instead, per
		// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/.
		WriteTimeout: 60 * time.Second,
		// Needs Go 1.8.
		IdleTimeout: 120 * time.Second,
		// TODO(twifkak): Specify ErrorLog?
	}
	// TCP keep-alive timeout on ListenAndServe is 3 minutes. To shorten,
	// follow the above Cloudflare blog.
	log.Fatal(server.ListenAndServe())
}
