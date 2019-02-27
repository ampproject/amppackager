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

// sameURLIgnoringFragment is a helper for AbsoluteUrlValue below.
// Returns true if |base| is the same as |u| with the exception that |u| may
// also have an additional fragment component.
func sameURLIgnoringFragment(base string, u *url.URL) bool {
	// Due to https://github.com/golang/go/issues/29603 we have an extra check
	// for the empty fragment case.
	if u.Fragment == "" {
		return base == u.String()
	}

	return base+"#"+u.Fragment == u.String()
}

// isProtocolRelative is a mostly correct parse for protocol relative inputs
// by looking for a "//" prefix after stripping any leading whitespace and
// control characters.
func isProtocolRelative(urlParam string) bool {
	urlParam = strings.TrimLeftFunc(urlParam, func(r rune) bool {
		return r <= 0x20
	})
	return strings.HasPrefix(urlParam, "//")
}

var queryEncoder = strings.NewReplacer(" ", "+")

// ToAbsoluteURL absolute-ifies |urlParam|, using |baseURL| as the base if
// |urlParam| is relative. If |urlParam| contains a fragment, this method
// will return only a fragment if it's absolute URL matches |documentURL|,
// which prevents changing an in-document navigation to a out-of-document
// navigation.
func ToAbsoluteURL(documentURL string, baseURL *url.URL,
	urlParam string) string {
	if urlParam == "" {
		return ""
	}

	refurl, err := url.Parse(urlParam)
	// TODO(gregable): Should we strip this URL instead (ie: return "").
	if err != nil {
		return urlParam
	}
	// Go URL parser is strict, but we want non-strict behavior for resolving
	// references as per the algorithm in https://tools.ietf.org/html/rfc3986#section-5.2.2
	// tl;dr If the scheme is the same, unset it so the relative URL truly is relative.
	if refurl.Scheme == baseURL.Scheme {
		refurl.Scheme = ""
	}
	// Handle relative URLs that have a different scheme but no authority. In this
	// case, use the base's authority. Note that this behavior is not
	// compliant with RFC 3986 Section 5, however, this is what the Chrome browser
	// does. See b/124445904 .
	if refurl.Scheme != "" && refurl.Host == "" {
		refurl.Host = baseURL.Host
	}
	absoluteURL := baseURL.ResolveReference(refurl)

	// TODO(gregable): We should probably assemble data: / mailto: / etc URLs,
	// which will force them to be URL encoded, but this was left to maintain
	// the old behavior for now.
	if absoluteURL.Scheme != "http" && absoluteURL.Scheme != "https" {
		return urlParam
	}

	// Check for a specific case of protocol relative URL (ex: "//foo.com/")
	// which specifies the host, but not the protocol. For b/27292423.
	// Essentially we use protocol relative as a hint that this resource will
	// be available on https even if it's resolved path was http. In this hinted
	// case, we always prefer https.
	if isProtocolRelative(urlParam) {
		absoluteURL.Scheme = "https"
	}

	// Avoid rewriting a local fragment such as "#top" to a remote absolute URL
	// of "http://example.com/#top" if it wasn't a remote URL already.
	// Note that we also try to identify empty fragments (ex: href="#").
	// net/url doesn't support these (https://github.com/golang/go/issues/29603)
	// so we try to detect them heuristically.
	if (absoluteURL.Fragment != "" || strings.HasPrefix(urlParam, "#")) &&
		sameURLIgnoringFragment(documentURL, absoluteURL) {
		return "#" + absoluteURL.Fragment
	}

	// Go's URL parser doesn't properly encode query string at parse time.
	// See https://github.com/golang/go/issues/22907 .
	// This currently only does the bare minimum:
	//  - encode space to "+"
	absoluteURL.RawQuery = queryEncoder.Replace(absoluteURL.RawQuery)
	return absoluteURL.String()
}

// SubresourceType describes the type of subresource
type SubresourceType int8

const (
	// ImageType is a subresource for an image
	ImageType SubresourceType = iota
	// OtherType is a subresource for everything besides an image.
	OtherType
)

// SubresourceOffset describes the location of a subresource URL within some text.
// For example, if the text value is ".a {background-image:url(foo.jpg)}", then
// Start === 25 and End === 32
type SubresourceOffset struct {
	SubType SubresourceType
	// The offset position denoting the start of the substring (inclusive)
	Start int
	// The offset position denoting the end of the substring (exclusive)
	End int
	// If the type is an image, an optional width to convert the image so.
	DesiredImageWidth int
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

// IsCacheURL returns true if the given string is from the AMPCache domain. This check is overly
// simplistic and does no actual verification that the URL resolves (doesn't 404), nor if the URL
// is of the correct format for the resource type (image, or otherwise).
func IsCacheURL(input string) bool {
	if u, err := url.Parse(input); err == nil {
		return strings.HasSuffix(u.Hostname(), AMPCacheHostName)
	}
	return false
}

// GetCacheURL returns an AMP Cache URL structure for the URL identified by
// the given offset (relative to 'input') or an error if the URL could not be
// parsed.
func (so *SubresourceOffset) GetCacheURL(documentURL string, base *url.URL,
	input string) (*CacheURL, error) {
	urlStr := (input)[so.Start:so.End]
	absolute := ToAbsoluteURL(documentURL, base, urlStr)
	if len(absolute) == 0 {
		return nil, errors.New("unable to convert empty URL string")
	}
	origURL, err := url.Parse(absolute)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing URL")
	}
	secureInfix := ""
	switch origURL.Scheme {
	case "https":
		// Add the secure infix
		secureInfix = "s/"
	case "http":
		// Supported
	default:
		// Unsupported scheme
		return nil, errors.New("unsupported scheme")
	}

	c := CacheURL{URL: origURL}
	// simplistic idempotent check
	if IsCacheURL(absolute) {
		c.Subdomain = strings.TrimSuffix(c.Hostname(), "."+AMPCacheHostName)
		return &c, nil
	}
	prefix := "/r/"
	if so.SubType == ImageType {
		prefix = "/i/"
		if so.DesiredImageWidth > 0 {
			wStr := strconv.Itoa(so.DesiredImageWidth)
			prefix = "/ii/w" + wStr + "/"
		}
	}
	c.Path = prefix + secureInfix + c.Hostname() + c.Path
	c.Scheme = "https"
	c.Subdomain = ToCacheURLSubdomain(c.Hostname())
	c.Host = c.Subdomain + "." + AMPCacheHostName
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
// For example, if the origin is www.example.com, this returns www-example-com.
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

