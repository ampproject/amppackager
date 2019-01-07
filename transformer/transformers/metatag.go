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
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
)

// MetaTag operates on the <meta> tag. It will relocate all meta tags found
// inside the body into the head.
//
// It does *not* sort the meta tags. This is done by ReorderHead.
func MetaTag(e *Context) error {
	for n := e.DOM.BodyNode; n != nil; n = htmlnode.Next(n) {
		// Skip non-meta tags.
		if n.DataAtom != atom.Meta {
			continue
		}

		// Relocate meta tags in body into head.
		removed := htmlnode.RemoveNode(&n)
		e.DOM.HeadNode.AppendChild(removed)
	}
	return nil
}
