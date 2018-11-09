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
	"net/url"
	"strings"
)

// RewriteAbsoluteURL returns a URL string suitable for the AMP cache's
// view of the given input URL. The resulting "absolute" URL, can be
// one of two possibilies:
//  - an absolute URL pointing to the same coordinates as the {url, base} tuple
//  - the exact text passed into url if the input was malformed,
//    a data: URL, or if we are inside a mustache template. The runtime must
//    fix-up URLs inside mustache templates on the client, as only the runtime
//    knows how to expand the templates.
//
// base is derived from the <base> tag and document URL for the origin document.
//
// inTemplate says whether the affected node is a descendant of a <template> tag.
//
// url is the original href value. It may be absolute, host-relative,
// path-relative, or fragment-relative. It could be a data: URL, it could
// be empty, it could be grotesquely malformed. It came from the internet.
// If relative, it is relative to base.
func RewriteAbsoluteURL(base *url.URL, inTemplate bool, url string) string {
	return rewriteURL(base, inTemplate, url, true)
}

// RewritePortableURL is similar to RewriteAbsoluteURL() except that it
// preserves fragment-relative URLs when url points to the same document as base.
func RewritePortableURL(base *url.URL, inTemplate bool, url string) string {
	return rewriteURL(base, inTemplate, url, false)
}

func rewriteURL(base *url.URL, inTemplate bool, url string, absolute bool) string {
	if inTemplate {
		// For b/26741101, do not rewrite URLs within mustache templates
		return url
	}
	orig := url
	url = strings.TrimSpace(url)
	if url == "" {
		return orig
	}

	// For b/27292423:
	// In general, if the origin doc was fetched on http:// and has a relative
	// URL to a resource, we must assume that the resource may only be
	// available on http. However: if the subresource has a protocol-relative
	// path (beginning with '//') this is a clear statement that either
	// HTTP or HTTPS can work. Special-case the protocol-relative case.
	if strings.HasPrefix(url, "//") {
		return "https:" + url
	}
	u, err := base.Parse(url)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		return url
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
		// This also handles the case where base == url, and neither has
		// a fragment. In which case we emit '#' which links to the top of the page.
		return "#" + u.Fragment
	}
	return uVal
}
