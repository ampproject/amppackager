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
	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
)

// TransformedIdentifier identifies that transformations
// were made for a specific platform on this document.
func TransformedIdentifier(e *Context) error {
	dom, err := amphtml.NewDOM(e.Doc)
	if err != nil {
		return err
	}

	htmlnode.SetAttribute(dom.HTMLNode, "", "transformed", "google")
	return nil
}
