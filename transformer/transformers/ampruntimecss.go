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

// AMPRuntimeCSS inlines the contents of the AMP HTML CSS RTV.
func AMPRuntimeCSS(e *Context) error {
	// If the document already has <style amp-runtime> then remove them. This
	// can happen if this transformer is run on already transformed AMP.
	for c := e.DOM.HeadNode.FirstChild; c != nil; c = htmlnode.Next(c) {
		if c.DataAtom == atom.Style &&
			htmlnode.HasAttribute(c, "", amphtml.AMPRuntime) {
			htmlnode.RemoveNode(&c)
		}
	}
	// Create <style amp-runtime> tag.
	n := htmlnode.Element("style", html.Attribute{Key: "amp-runtime"})
	// Annotate it with the AMP Runtime version that is being inlined.
	rtv := "latest"
	if e.Request.GetRtv() != "" {
		rtv = e.Request.GetRtv()
	}
	htmlnode.SetAttribute(n, "", "i-amphtml-version", rtv)
	// Place it first in the document <head>.
	e.DOM.HeadNode.InsertBefore(n, e.DOM.HeadNode.FirstChild)

	// The contents of the runtime css are available, so inline it.
	if e.Request.GetCss() != "" {
		n.AppendChild(htmlnode.Text(strings.TrimSpace(e.Request.GetCss())))
		return nil
	}

	return nil
}
