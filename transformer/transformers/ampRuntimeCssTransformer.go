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
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

// AMPRuntimeCSSTransformer inlines the contents of the AMP HTML CSS RTV, or
// inserts a link into the appropriately revisioned v0.css (e.g. 102992221).
func AMPRuntimeCSSTransformer(e *Engine) {
	dom, ok := amphtml.NewDOM(e.Doc)
	if !ok {
		return
	}

	// If server side rendering is active, then look for the
	// <style amp-runtime> added by ServerSideRenderingTransformer into
	// <head>.
	n, ok := findStyleAMPRuntime(dom.HeadNode)
	if !ok {
		// No Server Side Rendering.
		return
	}

	// Annotate the <style amp-runtime> tag with the version that is being
	// inlined or loaded with tag link.
	rtv := "latest"
	if e.Request.GetRtv() != "" {
		rtv = e.Request.GetRtv()
	}
	htmlnode.SetAttribute(n, "", "i-amphtml-version", rtv)

	// The contents of the runtime css are available, so inline it.
	if e.Request.GetCss() != "" {
		n.AppendChild(htmlnode.Text(strings.TrimSpace(e.Request.GetCss())))
		return
	}

	// Otherwise: add a link to the versioned v0.css.
	link := amphtml.AMPCacheSchemeAndHost
	if e.Request.GetRtv() != "" {
		link = link + "/rtv/" + e.Request.GetRtv()
	}
	link = link + "/v0.css"
	l := htmlnode.Element("link", html.Attribute{Key: "rel", Val: "stylesheet"},
		html.Attribute{Key: "href", Val: link})
	dom.HeadNode.AppendChild(l)
}

// findStyleAMPRuntime returns the <style amp-runtime> element or false
func findStyleAMPRuntime(n *html.Node) (*html.Node, bool) {
	if n.DataAtom != atom.Head {
		return nil, false
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.DataAtom == atom.Style && htmlnode.HasAttribute(c, amphtml.AMPRuntime) {
			return c, true
		}
	}
	return nil, false
}
