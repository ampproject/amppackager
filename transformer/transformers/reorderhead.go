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
	"sort"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

type headNodes struct {
	linkFavicon                   []*html.Node
	linkResourceHint              []*html.Node
	linkStylesheetBeforeAMPCustom []*html.Node
	linkStylesheetRuntimeCSS      *html.Node
	linkStylesheetAmpStoryCSS     *html.Node
	metaCharset                   *html.Node
	metaViewport                  *html.Node
	metaOther                     []*html.Node
	noscript                      *html.Node
	other                         []*html.Node
	scriptAMPRuntime              []*html.Node
	scriptAMPViewer               *html.Node
	scriptNonRenderDelaying       []*html.Node
	scriptRenderDelaying          []*html.Node
	scriptAmpStoryDvhPolyfill     *html.Node
	styleAMPBoilerplate           *html.Node
	styleAMPCustom                *html.Node
	styleAMPRuntime               *html.Node
	styleAMPStory                 *html.Node
}

// ReorderHead reorders the children of <head>. Specifically, it
// orders the <head> like so:
// (0) <meta charset> tag
// (1) <script amp-story-dvh-polyfill> inline script tag (see
// AmpStoryCssTransformer)
// (2) <style amp-runtime> (inserted by ampruntimecss.go)
// (3) <style amp-extension=amp-story> OR <link rel=stylesheet
// href=https://cdn.ampproject.org/v0/amp-story-1.0.css> (inserted by
// AmpStoryCss). Only one of these will be inserted
// by this transformer.
// (4) remaining <meta> tags (those other than <meta charset>)
// (5) AMP runtime <script> tag(s)
// (6) AMP viewer runtime .js <script> tag
// (7) <script> tags that are render delaying
// (8) <script> tags for remaining extensions
// (9) <link> tag for favicons
// (10) <link> tag for resource hints
// (11) <link rel=stylesheet> tags before <style amp-custom>
// (12) <style amp-custom>
// (13) any other tags allowed in <head>
// (14) AMP boilerplate (first style amp-boilerplate, then noscript)
func ReorderHead(e *Context) error {
	hn := new(headNodes)

	// Register each set of children we care about the order of in <head>.
	for c := e.DOM.HeadNode.FirstChild; c != nil; c = c.NextSibling {
		switch c.DataAtom {
		case atom.Link:
			registerLink(c, hn)
		case atom.Meta:
			registerMeta(c, hn)
		case atom.Noscript:
			hn.noscript = c
		case atom.Script:
			registerScript(c, hn)
		case atom.Style:
			registerStyle(c, hn)
		default:
			hn.other = append(hn.other, c)
		}
	}

	// Uniquifies custom-element and custom-template scripts such that only one
	// of each is included, preferring the first one encountered and sorting by
	// custom-element or custom-template attribute's value.
	hn.scriptRenderDelaying = uniquifyAndSortByExtensionScript(hn.scriptRenderDelaying)
	hn.scriptNonRenderDelaying = uniquifyAndSortByExtensionScript(hn.scriptNonRenderDelaying)

	// Remove children of <head>.
	htmlnode.RemoveAllChildren(e.DOM.HeadNode)

	// Append children of <head> in specific order.
	if hn.metaCharset != nil {
		e.DOM.HeadNode.AppendChild(hn.metaCharset)
	}
	// We want the `<meta name=viewport>` to be before the dvh polyfill
	// because the `<meta name=viewport>` can trigger a change on the innerHeight
	// of the document on iOS which doesn't cause a resize event. This can cause
	// the calculation in the dvh polyfill script to be incorrect.
	if hn.metaViewport != nil {
		e.DOM.HeadNode.AppendChild(hn.metaViewport)
	}
	// We want the dvh polyfill to be before the amp-story styles to prevent
	// triggering an increase to CLS score.
	if hn.scriptAmpStoryDvhPolyfill != nil {
		e.DOM.HeadNode.AppendChild(hn.scriptAmpStoryDvhPolyfill)
	}
	if hn.linkStylesheetRuntimeCSS != nil {
		e.DOM.HeadNode.AppendChild(hn.linkStylesheetRuntimeCSS)
	}
	if hn.styleAMPRuntime != nil {
		e.DOM.HeadNode.AppendChild(hn.styleAMPRuntime)
	}
	if hn.linkStylesheetAmpStoryCSS != nil {
		e.DOM.HeadNode.AppendChild(hn.linkStylesheetAmpStoryCSS)
	}
	if hn.styleAMPStory != nil {
		e.DOM.HeadNode.AppendChild(hn.styleAMPStory)
	}
	htmlnode.AppendChildren(e.DOM.HeadNode, hn.metaOther...)
	htmlnode.AppendChildren(e.DOM.HeadNode, hn.scriptAMPRuntime...)
	if hn.scriptAMPViewer != nil {
		e.DOM.HeadNode.AppendChild(hn.scriptAMPViewer)
	}
	htmlnode.AppendChildren(e.DOM.HeadNode, hn.scriptRenderDelaying...)
	htmlnode.AppendChildren(e.DOM.HeadNode, hn.scriptNonRenderDelaying...)
	htmlnode.AppendChildren(e.DOM.HeadNode, hn.linkFavicon...)
	htmlnode.AppendChildren(e.DOM.HeadNode, hn.linkResourceHint...)
	htmlnode.AppendChildren(e.DOM.HeadNode, hn.linkStylesheetBeforeAMPCustom...)
	if hn.styleAMPCustom != nil {
		e.DOM.HeadNode.AppendChild(hn.styleAMPCustom)
	}
	htmlnode.AppendChildren(e.DOM.HeadNode, hn.other...)
	if hn.styleAMPBoilerplate != nil {
		e.DOM.HeadNode.AppendChild(hn.styleAMPBoilerplate)
	}
	if hn.noscript != nil {
		e.DOM.HeadNode.AppendChild(hn.noscript)
	}
	return nil
}

// registerLink registers <link> tags to different variables depending on the attributes on the <link> tag. These are (1) resource hint <link> tags, (2) favicon <link> tags, (3) stylesheets before AMP Custom stylesheet, and (4) all other <link> tags.
func registerLink(n *html.Node, hn *headNodes) {
	if a, ok := htmlnode.FindAttribute(n, "", "rel"); ok {
		switch strings.ToLower(a.Val) {
		case "dns-prefetch preconnect":
			hn.linkResourceHint = append(hn.linkResourceHint, n)
			return
		case "icon", "icon shortcut", "shortcut icon":
			hn.linkFavicon = append(hn.linkFavicon, n)
			return
		case "stylesheet":
			// The AmpRuntimeCss inserts a stylesheet for the AMP Runtime CSS. It must remain early in the head immediately before <style amp-custom>.
			if v, ok := htmlnode.GetAttributeVal(n, "", "href"); ok &&
				strings.HasPrefix(v, amphtml.AMPCacheRootURL) &&
				strings.HasSuffix(v, "/v0.css") {
				hn.linkStylesheetRuntimeCSS = n
				return
			}
			// The AmpStoryCss inserts a stylesheet for the amp-story-1.0 CSS. It must remain early in the head immediately before <style amp-custom>.
			if v, ok := htmlnode.GetAttributeVal(n, "", "href"); ok &&
				strings.HasPrefix(v, amphtml.AMPCacheRootURL) &&
				strings.HasSuffix(v, "/amp-story-1.0.css") {
				hn.linkStylesheetAmpStoryCSS = n
				return
			}

			if hn.styleAMPCustom == nil {
				hn.linkStylesheetBeforeAMPCustom = append(hn.linkStylesheetBeforeAMPCustom, n)
				return
			}
		}
	}
	hn.other = append(hn.other, n)
}

// registerMeta registers <meta> tags to different variables depending on the attributes on the <meta> tag. These are (1) the required <meta charset> and (2) all other <meta> tags.
func registerMeta(n *html.Node, hn *headNodes) {
	if htmlnode.HasAttribute(n, "", "charset") {
		hn.metaCharset = n
		return
	}
	if v, ok := htmlnode.GetAttributeVal(n, "", "name"); ok &&
		v == "viewport" {
		hn.metaViewport = n
		return
	}
	hn.metaOther = append(hn.metaOther, n)
}

// registerScript registers <script> tags to different variables depending on the attributes on the <script> tag. These are the (1) AMP Runtime script, (2) the render delaying AMP Custom Element scripts, (3) the non-render delaying AMP Custom Element scripts, (4) amp-story-dvh-polyfill script, and (5) all other <script> tags.
func registerScript(n *html.Node, hn *headNodes) {
	if amphtml.IsScriptAMPRuntime(n) {
		hn.scriptAMPRuntime = append(hn.scriptAMPRuntime, n)
		return
	}
	if amphtml.IsScriptAMPViewer(n) {
		hn.scriptAMPViewer = n
		return
	}
	if amphtml.IsScriptAMPExtension(n) {
		if amphtml.IsScriptRenderDelaying(n) {
			hn.scriptRenderDelaying = append(hn.scriptRenderDelaying, n)
			return
		}
		hn.scriptNonRenderDelaying = append(hn.scriptNonRenderDelaying, n)
		return
	}
	if amphtml.IsAmpStoryDvhPolyfillScript(n) {
		hn.scriptAmpStoryDvhPolyfill = n
		return
	}
	hn.other = append(hn.other, n)
}

// registerStyle registers <style> tags to different variables depending on the attributes on the <style> tag. These are the (1) AMP or AMP4ADS Boilerplate, (2)the AMP Custom stylesheet, (3) the AMP Runtime stylesheet and (4) all other <style> tags.
func registerStyle(n *html.Node, hn *headNodes) {
	if htmlnode.HasAttribute(n, "", amphtml.AMPBoilerplate) || htmlnode.HasAttribute(n, "", amphtml.AMP4AdsBoilerplate) {
		hn.styleAMPBoilerplate = n
		return
	}
	if htmlnode.HasAttribute(n, "", amphtml.AMPCustom) {
		hn.styleAMPCustom = n
		return
	}
	if htmlnode.HasAttribute(n, "", amphtml.AMPRuntime) {
		hn.styleAMPRuntime = n
		return
	}
	if v, ok := htmlnode.GetAttributeVal(n, "", "amp-extension"); ok &&
		v == "amp-story" {
		hn.styleAMPStory = n
		return
	}
	hn.other = append(hn.other, n)
}

// uniquifyAndSortByExtensionScript returns the unique script definition nodes, keeping the first one encountered and sorted by unique script definition.  The unique script defiintion is defined by "name-version(.js|.mjs)".  Example (amp-ad): amp-ad-0.1.js (regular/nomodule), amp-ad-0.1.mjs (module).  If the pattern is not found then uses value of "src" attribute.  The AMP Validator prevents a mix of regular and nomodule extension.
func uniquifyAndSortByExtensionScript(n []*html.Node) []*html.Node {
	var k []string
	var u []*html.Node
	m := make(map[string]*html.Node)
	for _, s := range n {
		if ext, ok := amphtml.AMPExtensionScriptDefinition(s); ok {
			if _, ok := m[ext]; !ok {
				m[ext] = s
				k = append(k, ext)
			}
		}
	}
	sort.Strings(k)
	for _, e := range k {
		u = append(u, m[e])
	}
	return u
}
