// Package amphtml contains common constants and utilies for working
// with AMP HTML.
package amphtml

import (
	"strings"

	"github.com/ampproject/amppackager/internal/pkg/htmlnode"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Common AMP string constants.
const (
	AMPAudio = "amp-audio"

	AMPBoilerplate = "amp-boilerplate"

	AMPBoilerplateCSS = "body{-webkit-animation:-amp-start 8s steps(1,end) 0s 1 normal both;-moz-animation:-amp-start 8s steps(1,end) 0s 1 normal both;-ms-animation:-amp-start 8s steps(1,end) 0s 1 normal both;animation:-amp-start 8s steps(1,end) 0s 1 normal both}@-webkit-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-moz-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-ms-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-o-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}"

	AMPBoilerplateNoscriptCSS = `body{-webkit-animation:none;-moz-animation:none;-ms-animation:none;animation:none}`

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

// IsScriptAMPRuntime returns true if the node is of the form <script async src=https://cdn.ampproject.org...v0.js></script>
func IsScriptAMPRuntime(n *html.Node) bool {
	if n.DataAtom != atom.Script {
		return false
	}
	if v, ok := htmlnode.GetAttributeVal(n, "src"); ok {
		return htmlnode.HasAttribute(n, "async") &&
			!htmlnode.HasAttribute(n, AMPCustomElement) &&
			!htmlnode.HasAttribute(n, AMPCustomTemplate) &&
			strings.HasPrefix(v, "https://cdn.ampproject.org/") &&
			(strings.HasSuffix(v, "/v0.js") ||
				strings.HasSuffix(v, "/amp4ads-v0.js"))
	}
	return false
}

// IsScriptRenderDelaying returns true if the node has one of these values for attribute 'custom-element': amp-dynamic-css-classes, amp-experiment, amp-story.
func IsScriptRenderDelaying(n *html.Node) bool {
	if n.DataAtom != atom.Script {
		return false
	}
	if v, ok := htmlnode.GetAttributeVal(n, AMPCustomElement); ok {
		// TODO(b/77581738): Remove amp-story from this list.
		return (v == AMPDynamicCSSClasses ||
			v == AMPExperiment ||
			v == AMPStory)
	}
	return false
}

// DOM encapsulates the various HTML nodes a transformer may need access to.
type DOM struct {
	HTMLNode *html.Node
	HeadNode *html.Node
	BodyNode *html.Node
}

// NewDOM constructs and returns a pointer to a DOM struct by finding
// the HTML nodes relevant to an AMP Document or ok=false if there is an error.
func NewDOM(n *html.Node) (*DOM, bool) {
	var ok bool
	d := new(DOM)
	d.HTMLNode, ok = htmlnode.FindNode(n, atom.Html)
	if !ok {
		return d, false
	}
	d.HeadNode, ok = htmlnode.FindNode(n, atom.Head)
	if !ok {
		return d, false
	}
	d.BodyNode, ok = htmlnode.FindNode(n, atom.Body)
	if !ok {
		return d, false
	}
	return d, true
}
