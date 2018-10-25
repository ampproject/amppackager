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
	"golang.org/x/net/html"
)

// MetaTag operates on the <meta> tag. It will relocate all meta tags found
// inside the body into the head.
//
// It does *not* sort the meta tags. This is done by ReorderHead.
func MetaTag(e *Context) error {
	var stk htmlnode.Stack
	stk.Push(e.DOM.RootNode)
	for len(stk) > 0 {
		top := stk.Pop()
		// Traverse the children in reverse order so the iteration of
		// the DOM tree traversal is in the proper sequence.
		// E.g. Given <a><b/><c/></a>, we will visit a, b, c.
		// An alternative is to traverse children in forward order and
		// utilize a queue instead.
		for c := top.LastChild; c != nil; c = c.PrevSibling {
			stk.Push(c)
		}
		metaTagTransform(top, e.DOM.HeadNode)
	}
	return nil
}

// metaTagTransform does the actual work on each node.
func metaTagTransform(n, h *html.Node) {
	// Skip non-meta tags.
	if n.DataAtom != atom.Meta {
		return
	}

	// Relocate meta tags in body into head.
	if htmlnode.IsDescendantOf(n, atom.Body) {
		n.Parent.RemoveChild(n)
		h.AppendChild(n)
	}
}
