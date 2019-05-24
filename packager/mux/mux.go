package mux

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ampproject/amppackager/packager/certcache"
	"github.com/ampproject/amppackager/packager/signer"
	muxp "github.com/ampproject/amppackager/packager/mux/params"
	"github.com/ampproject/amppackager/packager/util"
	"github.com/ampproject/amppackager/packager/validitymap"
)

type mux struct {
	certCache *certcache.CertCache
	signer *signer.Signer
	validityMap *validitymap.ValidityMap
}

// Implements the HTTP routing strategy used by the amppkg server. The main
// purpose for rolling our own is to be able to route the following types of
// URLs:
//   /priv/doc/https://example.com/esc%61ped%2Furl.html
// and have the signer sign the URL exactly as it is encoded in the request, in
// order to meet docs/cache_requirements.md.
//
// The default http mux had some problems which led to me investigating 3rd
// party routers. (TODO: Remember and document those problems.)
//
// I investigated two 3rd party routers capable of handling catch-all suffix
// parameters:
//   https://github.com/julienschmidt/httprouter
//   https://github.com/dimfeld/httptreemux
// Both libs unescape their catch-all parameters, making the above use-case
// impossible. The latter does so despite documenting support for unmodified
// URL escapings. This, plus a lack of feature needs, led me to believe that
// writing our own mux was the best approach.
func New(certCache *certcache.CertCache, signer *signer.Signer, validityMap *validitymap.ValidityMap) http.Handler {
	return &mux{certCache, signer, validityMap}
}

func tryTrimPrefix(s, prefix string) (string, bool) {
	sLen := len(s)
	trimmed := strings.TrimPrefix(s, prefix)
	return trimmed, len(trimmed) != sLen
}

// TODO(twifkak): Test this. Maybe by changing all the tests!
func (this *mux) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// TODO(twifkak): Early return unless method is GET or HEAD.

	params := map[string]string{}
	req = muxp.WithParams(req, params)

	// Use EscapedPath rather than RequestURI because the latter can take
	// absolute-form, per https://tools.ietf.org/html/rfc7230#section-5.3.
	//
	// Use EscapedPath rather than RawPath to verify that the request URI
	// is a valid encoding, per
	// https://wicg.github.io/webpackage/draft-yasskin-http-origin-signed-responses.html#application-signed-exchange
	// item 3.
	path := req.URL.EscapedPath()
	if suffix, ok := tryTrimPrefix(path, "/priv/doc"); ok {
		if suffix == "" {
			this.signer.ServeHTTP(resp, req)
		} else if suffix[0] == '/' {
			params["signURL"] = suffix[1:]
			this.signer.ServeHTTP(resp, req)
		} else {
			http.NotFound(resp, req)
		}
	} else if suffix, ok := tryTrimPrefix(path, util.CertURLPrefix + "/"); ok {
		unescaped, err := url.PathUnescape(suffix)
		if err != nil {
			http.Error(resp, "400 bad request", http.StatusBadRequest)
		} else {
			params["certName"] = unescaped
			this.certCache.ServeHTTP(resp, req)
		}
	} else if path == util.ValidityMapPath {
		this.validityMap.ServeHTTP(resp, req)
	} else {
		http.NotFound(resp, req)
	}
}
