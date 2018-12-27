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

package amphtml

import (
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"github.com/pkg/errors"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

// Common AMP string constants.
const (
	AMPAudio = "amp-audio"

	AMPBoilerplate = "amp-boilerplate"

	AMPBoilerplateCSS = "body{-webkit-animation:-amp-start 8s steps(1,end) 0s 1 normal both;-moz-animation:-amp-start 8s steps(1,end) 0s 1 normal both;-ms-animation:-amp-start 8s steps(1,end) 0s 1 normal both;animation:-amp-start 8s steps(1,end) 0s 1 normal both}@-webkit-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-moz-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-ms-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-o-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}"

	AMPBoilerplateNoscriptCSS = "body{-webkit-animation:none;-moz-animation:none;-ms-animation:none;animation:none}"

	AMPCacheHostName = "cdn.ampproject.org"

	AMPCacheSchemeAndHost = "https://cdn.ampproject.org"

	AMPCacheRootURL = "https://cdn.ampproject.org/"

	AMP4AdsBoilerplate = "amp4ads-boilerplate"

	AMP4AdsAndAMP4EMailBoilerplateCSS = "body{visibility:hidden}"

	AMP4EmailBoilerplate = "amp4email-boilerplate"

	AMPCustom = "amp-custom"

	AMPCustomElement = "custom-element"

	AMPCustomTemplate = "custom-template"

	AMPDynamicCSSClasses = "amp-dynamic-css-classes"

	AMPExperiment = "amp-experiment"

	AMPRuntime = "amp-runtime"

	AMPStory = "amp-story"

	IAMPHTMLLayout = "i-amphtml-layout"
)

// IsAMPCustomElement returns true if the node is an AMP custom element.
func IsAMPCustomElement(n *html.Node) bool {
	return n.Type == html.ElementNode && strings.HasPrefix(n.Data, "amp-")
}

// IsScriptAMPExtension returns true if the node is a script tag with either attribute `custom-element` or `custom-template` present.
func IsScriptAMPExtension(n *html.Node) bool {
	if n.DataAtom != atom.Script {
		return false
	}
	return htmlnode.HasAttribute(n, "", AMPCustomElement) || htmlnode.HasAttribute(n, "", AMPCustomTemplate)
}

// IsScriptAMPRuntime returns true if the node is of the form <script async src=https://cdn.ampproject.org...v0.js></script>
func IsScriptAMPRuntime(n *html.Node) bool {
	if n.DataAtom != atom.Script {
		return false
	}
	if v, ok := htmlnode.GetAttributeVal(n, "", "src"); ok {
		return htmlnode.HasAttribute(n, "", "async") &&
			!htmlnode.HasAttribute(n, "", AMPCustomElement) &&
			!htmlnode.HasAttribute(n, "", AMPCustomTemplate) &&
			strings.HasPrefix(v, AMPCacheRootURL) &&
			(strings.HasSuffix(v, "/v0.js") ||
				strings.HasSuffix(v, "/amp4ads-v0.js"))
	}
	return false
}

// IsScriptAMPViewer returns true if the node is of the form <script async src=https://cdn.ampproject.org/v0/amp-viewer-integration-...js></script>
func IsScriptAMPViewer(n *html.Node) bool {
	if n.DataAtom != atom.Script {
		return false
	}
	a, ok := htmlnode.FindAttribute(n, "", "src")
	return ok &&
		!htmlnode.HasAttribute(n, "", AMPCustomTemplate) &&
		strings.HasPrefix(a.Val,
			AMPCacheSchemeAndHost+"/v0/amp-viewer-integration-") &&
		strings.HasSuffix(a.Val, ".js") &&
		htmlnode.HasAttribute(n, "", "async") &&
		!htmlnode.HasAttribute(n, "", AMPCustomElement)
}

// IsScriptRenderDelaying returns true if the node has one of these values for attribute 'custom-element': amp-dynamic-css-classes, amp-experiment, amp-story.
func IsScriptRenderDelaying(n *html.Node) bool {
	if n.DataAtom != atom.Script {
		return false
	}
	if v, ok := htmlnode.GetAttributeVal(n, "", AMPCustomElement); ok {
		// TODO(b/77581738): Remove amp-story from this list.
		return (v == AMPDynamicCSSClasses ||
			v == AMPExperiment ||
			v == AMPStory)
	}
	return false
}

// DOM encapsulates the various HTML nodes a transformer may need access to.
type DOM struct {
	RootNode *html.Node
	HTMLNode *html.Node
	HeadNode *html.Node
	BodyNode *html.Node
}

// NewDOM constructs and returns a pointer to a DOM struct by finding
// the HTML nodes relevant to an AMP Document or an error if there was
// a problem.
// TODO(alin04): I don't think this can EVER return an error. The golang
// parser creates all these nodes if they're missing.
func NewDOM(n *html.Node) (*DOM, error) {
	var ok bool
	d := &DOM{RootNode: n}
	if d.HTMLNode, ok = htmlnode.FindNode(n, atom.Html); !ok {
		return d, errors.New("missing <html> node")
	}
	if d.HeadNode, ok = htmlnode.FindNode(d.HTMLNode, atom.Head); !ok {
		return d, errors.New("missing <head> node")
	}
	if d.BodyNode, ok = htmlnode.FindNode(d.HTMLNode, atom.Body); !ok {
		return d, errors.New("missing <body> node")
	}
	return d, nil
}
