// Copyright 2019 Google LLC
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

// Package url attempts to fix some of the issues related to Go's url.String().
// This is a starting point for discrepancies between Go's implementation and
// the spec in https://tools.ietf.org/html/rfc3986.
//
// ABNF definitions:
// https://tools.ietf.org/html/rfc3986#appendix-A
// https://tools.ietf.org/html/rfc2234#section-6.1
package url

import (
	gourl "net/url"
	"strings"
)

// urlEncoder is used to encode the entire URL string and all its components -
// host, path, query, fragment.
// TODO(b/130250234): This is just a starting point for what may need to be
// encoded. Possibly encode more things.
var urlEncoder = strings.NewReplacer(
	"<", "%3C",
	">", "%3E",
	" ", "%20",
)

// String post-processes Go's url.String() to encode characters that should be encoded.
// https://golang.org/issue/22907
// https://golang.org/issue/30844
// https://golang.org/issue/30922
func String(input *gourl.URL) string {
	// TODO(b/130234885): Handle relative URLs.
	return urlEncoder.Replace(input.String())
}

// Splits the string on the separator (if it exists), retaining the separator.
// E.g. split("foo,bar", ",") returns "foo", ",bar"
func split(s string, sep string) (string, string) {
	i := strings.Index(s, sep)
	if i < 0 {
		return s, ""
	}
	return s[:i], s[i:]
}
