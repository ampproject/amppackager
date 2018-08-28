package transformer

import (
	"github.com/ampproject/amppackager/internal/pkg/amphtml"
	"github.com/ampproject/amppackager/internal/pkg/htmlnode"
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
		n.AppendChild(htmlnode.Text(e.Request.GetCss()))
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
