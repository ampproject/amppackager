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
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

// AMPBoilerplate removes <style> and <noscript> tags in <head>,
// keeping only the amp-custom style tag. It then inserts the amp-boilerplate.
func AMPBoilerplate(e *Context) error {
        // Remove <style> and <noscript> tags keeping only the amp-runtime and
        // amp-custom style tag. amp-runtime may be removed later by the
        // AMPRuntimeCSS transformer.
	for n := e.DOM.HeadNode; n != nil && n.DataAtom != atom.Body; n = htmlnode.Next(n) {
		switch n.DataAtom {
		case atom.Style:
			if !htmlnode.HasAttribute(n, "", amphtml.AMPCustom) && !htmlnode.HasAttribute(n, "", amphtml.AMPRuntime) {
				htmlnode.RemoveNode(&n)
			}
		case atom.Noscript:
			htmlnode.RemoveNode(&n)
		}
	}

	if e.Version >= 3 {
		// If the document had been modified by a Server-Side-Rendering transform
		// earlier, for example by the AMP Optimizer, and that transform
		// determined that the boilerplate was unnecessary, we don't add the
		// boilerplate back. Note this can mean that an error in that transform
		// could result in boilerplate being removed when it shouldn't be.
		if htmlnode.HasAttribute(e.DOM.HTMLNode, "", "i-amphtml-no-boilerplate") {
			return nil
		}
	}

	boilerplate, css := determineBoilerplateAndCSS(e.DOM.HTMLNode)

	styleNode := htmlnode.Element("style", html.Attribute{Key: boilerplate})
	e.DOM.HeadNode.AppendChild(styleNode)

	cssNode := htmlnode.Text(css)
	styleNode.AppendChild(cssNode)

	if boilerplate != amphtml.AMPBoilerplate {
		return nil
	}

	// Regular AMP boilerplate also includes a noscript.
	noScriptNode := htmlnode.Element("noscript")
	e.DOM.HeadNode.AppendChild(noScriptNode)

	noScriptStyle := htmlnode.Element("style", html.Attribute{Key: boilerplate})
	noScriptNode.AppendChild(noScriptStyle)

	noScriptCSS := htmlnode.Text(amphtml.AMPBoilerplateNoscriptCSS)
	noScriptStyle.AppendChild(noScriptCSS)
	return nil
}

// Returns the boilerplate style and CSS for the flavor of AMP used.
// ⚡ is \u26a1.
func determineBoilerplateAndCSS(n *html.Node) (string, string) {
	boilerplate := amphtml.AMPBoilerplate
	css := amphtml.AMPBoilerplateCSS
	for i := range n.Attr {
		switch n.Attr[i].Key {
		case "amp4ads", "⚡4ads":
			boilerplate = amphtml.AMP4AdsBoilerplate
			css = amphtml.AMP4AdsAndAMP4EMailBoilerplateCSS
		case "amp4email", "⚡4email":
			boilerplate = amphtml.AMP4EmailBoilerplate
			css = amphtml.AMP4AdsAndAMP4EMailBoilerplateCSS
		}
	}
	return boilerplate, css
}
