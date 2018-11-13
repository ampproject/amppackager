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

package transformers

import (
	"net/url"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	rpb "github.com/ampproject/amppackager/transformer/request"
)

// Context stores the root DOM Node and contextual data used for the
// transformers.
type Context struct {
	// The DOM of the parsed HTML input.
	DOM *amphtml.DOM

	// The public URL of the document, i.e. the location that should appear in the browser URL bar.
	// This is the URL-typed equivalent of Request.DocumentUrl.
	DocumentURL *url.URL

	// The base URL of the document, derived from the <base> tag, if any. If the base href is
	// relative, then it is parsed in the context of DocumentURL.
	BaseURL *url.URL

	// The version to use when transforming the DOM.
	Version int64

	// The request parameters.
	Request *rpb.Request
}
