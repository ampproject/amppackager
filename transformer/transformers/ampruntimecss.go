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

const versionAttr string = "i-amphtml-version"
const prodPrefix string = "01"

// AMPRuntimeCSS inlines the contents of the AMP HTML CSS RTV.
//
// If there is no style, a new one is added. RTV and CSS are set if available.
//
// If there are one or more styles already (e.g. from Optimizer), then they are
// all deleted and replaced with one iff both RTV and CSS are available and newer.
func AMPRuntimeCSS(e *Context) error {
	rtv := e.Request.GetRtv()
	css := e.Request.GetCss()

	nodesToRemove := []*html.Node{}
	for c := e.DOM.HeadNode.FirstChild; c != nil; c = htmlnode.Next(c) {
		if c.DataAtom == atom.Style &&
			htmlnode.HasAttribute(c, "", amphtml.AMPRuntime) {
			cRTV, cHasRTV := htmlnode.GetAttributeVal(c, "", versionAttr)
			// Replace if older than the one in the packager, or if
			// the RTV is not of the prod variety.
			if !cHasRTV || !strings.HasPrefix(cRTV, prodPrefix) ||
			   (rtv != "" && css != "" && rtv > cRTV) {
				nodesToRemove = append(nodesToRemove, c)
			} else {
				// If any styles are not going to be removed,
				// then bail. We don't want to create a
				// conflict between existing and new styles.
				return nil
			}
		}
	}

	// Delete previous styles, if any.
	for _, n := range nodesToRemove {
		htmlnode.RemoveNode(&n)
	}

	// Create <style amp-runtime> tag.
	n := htmlnode.Element("style", html.Attribute{Key: amphtml.AMPRuntime})
	// Annotate it with the AMP Runtime version that is being inlined.
	if rtv == "" {
		rtv = "latest"
	}
	htmlnode.SetAttribute(n, "", versionAttr, rtv)
	// Place it first in the document <head>.
	e.DOM.HeadNode.InsertBefore(n, e.DOM.HeadNode.FirstChild)

	// If the contents of the runtime css are available, inline it.
	if css != "" {
		n.AppendChild(htmlnode.Text(strings.TrimSpace(css)))
		return nil
	}

	return nil
}
