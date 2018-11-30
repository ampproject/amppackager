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

// Utilities related to handling of URLs in AMP.

package amphtml

import (
	"crypto/sha256"
	"encoding/base32"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/net/idna"
)

// ToAbsoluteURL returns a URL string suitable for the AMP cache's
// view of the given input URL. The resulting "absolute" URL, can be
// one of two possibilies:
//  - an absolute URL pointing to the same coordinates as the {in, base} tuple
//  - the exact text passed into in if the input was malformed,
//    a data: URL, or if we are inside a mustache template. The runtime must
//    fix-up URLs inside mustache templates on the client, as only the runtime
//    knows how to expand the templates.
//
// base is derived from the <base> tag and document URL for the origin document.
//
// in is the original href value. It may be absolute, host-relative,
// path-relative, or fragment-relative. It could be a data: URL, it could
// be empty, it could be grotesquely malformed. It came from the internet.
// If relative, it is relative to base.
func ToAbsoluteURL(base *url.URL, in string) string {
	return convertToPortableOrAbsoluteURL(base, in, true)
}

// ToPortableURL is similar to ToAbsoluteURL() except that it
// preserves fragment-relative URLs when url points to the same document as base.
func ToPortableURL(base *url.URL, in string) string {
	return convertToPortableOrAbsoluteURL(base, in, false)
}

// SubresourceURL stores information about a subresource.
type SubresourceURL struct {
	// the URL string for the subresource. It could be relative or absolute.
	URLString string
	// If the subresource is an image, an optional width to convert the image to.
	DesiredWidth int
	// Optional descriptor (used for image candidates), representing width or pixel density. This is only
	// used if DesiredWidth is not set.
	descriptor string
}

// CacheURL represents an AMP Cache URL
type CacheURL struct {
	Subdomain  string // publisher's subdomain within the cache. e.g. "example-com"
	descriptor string // Optional descriptor (used for image candidates), representing width or pixel density.
	*url.URL
}

// OriginDomain returns the scheme and host name, ignoring any path info.
func (c *CacheURL) OriginDomain() string {
	return "https://" + c.Subdomain + "." + AMPCacheHostName
}

// String reassembles the URL into a URL string
func (c *CacheURL) String() string {
	s := c.URL.String()
	if len(c.descriptor) > 0 {
		s = s + " " + c.descriptor
	}
	return s
}

// ToCacheURL returns an AMP Cache URL structure for the given subresource, or an error if the input
// could not be parsed.
func (r *SubresourceURL) ToCacheURL() (*CacheURL, error) {
	origURL, err := url.Parse(r.URLString)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing URL")
	}
	c := CacheURL{URL: origURL}
	path := "/i"
	if r.DesiredWidth > 0 {
		wStr := strconv.Itoa(r.DesiredWidth)
		path = "/ii/w" + wStr
		c.descriptor = wStr + "w"
	} else {
		c.descriptor = r.descriptor
	}
	switch c.Scheme {
	case "https":
		// Add the secure infix and drop the scheme.
		path = path + "/s" + r.URLString[len("https:/"):]
	case "http":
		// Drop the scheme
		path = path + r.URLString[len("http:/"):]
	default:
		// unsupported scheme
		return &c, nil
	}
	c.Scheme = "https"
	c.Subdomain = ToCacheURLSubdomain(c.Hostname())
	c.Host = c.Subdomain + "." + AMPCacheHostName
	c.Path = path
	return &c, nil
}

// ToCacheURLSubdomain converts an origin domain name to a dot-free human-readable string,
// that can be used in combination with an AMP Cache domain to identify the publisher's
// subdomain within that cache. If problems are encountered, fallback to a one-way hash.
//
// 1. Converts the origin domain from IDN (Punycode) to UTF-8.
// 2. Replaces every "-" (dash) with "--"(2 dashes).
// 3. Replaces every "." (dot) with a "-" (dash).
// 4. Converts back to IDN (Punycode).
//
// For example, if the origin is www.example.com, its cache prefix will be www-example-com.
// On Google's AMP Cache, this will be prepended to the Google cache domain resulting in
// www-example-com.cdn.ampproject.org .
// See https://developers.google.com/amp/cache/overview for more info
func ToCacheURLSubdomain(originHost string) string {
	p := idna.New(idna.MapForLookup(), idna.VerifyDNSLength(true), idna.Transitional(true), idna.BidiRule())
	unicode, err := p.ToUnicode(originHost)
	if err != nil {
		return fallbackCacheURLSubdomain(originHost)
	}
	var sb strings.Builder
	for _, rune := range unicode {
		switch rune {
		case '.':
			sb.WriteRune('-')
		case '-':
			sb.WriteString("--")
		default:
			sb.WriteRune(rune)
		}
	}
	if result, err := p.ToASCII(sb.String()); err == nil && strings.ContainsRune(sb.String(), '-') {
		return result
	}
	return fallbackCacheURLSubdomain(originHost)
}


func fallbackCacheURLSubdomain(originHost string) string {
	sha := sha256.New()
	sha.Write([]byte(originHost))
	result := base32.StdEncoding.EncodeToString(sha.Sum(nil))
	// Remove the last four chars are always "====" which are not legal in a domain name.
	return strings.ToLower(result[0:52])
}


func convertToPortableOrAbsoluteURL(base *url.URL, in string, absolute bool) string {
	if base == nil {
		base, _ = url.Parse("")
	}
	orig := in
	in = strings.TrimSpace(in)
	if in == "" {
		return orig
	}

	// For b/27292423:
	// In general, if the origin doc was fetched on http:// and has a relative
	// URL to a resource, we must assume that the resource may only be
	// available on http. However: if the subresource has a protocol-relative
	// path (beginning with '//') this is a clear statement that either
	// HTTP or HTTPS can work. Special-case the protocol-relative case.
	if strings.HasPrefix(in, "//") {
		return "https:" + in
	}
	u, err := base.Parse(in)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		return in
	}

	uVal := u.String()
	if absolute {
		return uVal
	}

	switch uVal {
	case base.String(), base.String() + "#" + u.Fragment:
		// Keep links to page-local fragments relative.
		// Otherwise, we'll turn "#dogs" into "https://origin.com/page.html#dogs"
		// and send the user away when we could have kept them on the page they
		// already loaded for a better experience.
		//
		// This also handles the case where base == in, and neither has
		// a fragment. In which case we emit '#' which links to the top of the page.
		return "#" + u.Fragment
	}
	return uVal
}

