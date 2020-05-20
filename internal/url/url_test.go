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

package url

import (
	gourl "net/url"
	"testing"
)

func TestString(t *testing.T) {
	tcs := []struct {
		desc, in, out string
	}{
		{
			"<>",
			"http://fo>o<.com/bl<ah?zombo=co>m&wh>at#ev>ah",
			"http://fo%3Eo%3C.com/bl%3Cah?zombo=co%3Em&wh%3Eat#ev%3Eah",
		},
		{
			"spaces encoded",
			"https://foo.com/i haz spaces?q=i haz spaces",
			"https://foo.com/i%20haz%20spaces?q=i%20haz%20spaces",
		},
		{
			"fragment with space and quote reescaped",
			"https://example.com/amp.html#fragment-\" ",
			"https://example.com/amp.html#fragment-%22%20",
		},
		{
			"slashes in query, encoded or not, preserved",
			"https://example.com/amp.html?URL=http://bar.com%2Fbaz",
			"https://example.com/amp.html?URL=http://bar.com%2Fbaz",
		},
		{
			"slashes in path, encoded or not, preserved",
			"https://example.com/foo%2Fbar/baz.html",
			"https://example.com/foo%2Fbar/baz.html",
		},
	}
	for _, tt := range tcs {
		t.Run(tt.desc, func(t *testing.T) {
			u, err := gourl.Parse(tt.in)
			if err != nil {
				t.Fatalf("Error parsing %s", tt.in)
			}
			s := String(u)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
