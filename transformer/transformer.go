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

// Package transformer invokes the golang HTML parser, executes the
// individual transformers (unless overridden), and prints the output
// to the provided string.
package transformer

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"github.com/ampproject/amppackager/transformer/printer"
	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/ampproject/amppackager/transformer/transformers"
	"github.com/pkg/errors"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

// Transformer functions must be added here in order to be passed in from
// the command line or invoked from other languages. Please keep alphabetical.
// Case-insensitive lookup.
//
// NOTE: The string mapping is necessary as a language cross-over to
// allow explicit transformer invocation (via the CUSTOM config).
var transformerFunctionMap = map[string]func(*transformers.Context) error{
	"ampboilerplate":        transformers.AMPBoilerplate,
	"ampruntimecss":         transformers.AMPRuntimeCSS,
	"linktag":               transformers.LinkTag,
	"metatag":               transformers.MetaTag,
	"nodecleanup":           transformers.NodeCleanup,
	"reorderhead":           transformers.ReorderHead,
	"serversiderendering":   transformers.ServerSideRendering,
	"stripjs":               transformers.StripJS,
	"transformedidentifier": transformers.TransformedIdentifier,
	"url":                   transformers.URL,
}

// The map of config to the list of transformers, in the order in
// which they should be executed.
var configMap = map[rpb.Request_TransformersConfig][]func(*transformers.Context) error{
	rpb.Request_DEFAULT: {
		// NodeCleanup should be first.
		transformers.NodeCleanup,
		transformers.StripJS,
		transformers.MetaTag,
		// TODO(alin04): Reenable LinkTag once validation is done.
		// transformers.LinkTag,
		// end TODO
		transformers.URL,
		transformers.AMPBoilerplate,
		// TODO(alin04): Reenable SSR once validation is done.
		// transformers.ServerSideRendering,
		// end TODO
		// AmpRuntimeCss must run after ServerSideRendering
		// TODO(alin04): Reenable AMPRuntimeCSS and
		// TransformedIdentifier once validation is done
		// transformers.AMPRuntimeCSS,
		// transformers.TransformedIdentifier,
		// end TODO
		// ReorderHead should run after all transformers that modify the
		// <head>, as they may do so without preserving the proper order.
		transformers.ReorderHead,
	},
	rpb.Request_NONE: {},
	rpb.Request_VALIDATION: {
		// TODO(alin04): Fill this in
		transformers.ReorderHead,
	},
	rpb.Request_CUSTOM: {},
}

// The maximum number of preloads to place in the Link header. This limit
// should be enforced by AMP Caches, to protect any pages that prefetch the SXG
// from an unnecessary number of fetches.
const maxPreloads = 20

// Override for tests.
var runTransformers = func(c *transformers.Context, fns []func(*transformers.Context) error) error {
	// Invoke the configured transformers
	for _, f := range fns {
		if err := f(c); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// ampAttrRE is a regexp to match html amp attributes. Its group capture should
// be compared against ampFormatSuffixes.
var ampAttrRE = func() *regexp.Regexp {
	r := regexp.MustCompile(`\A(?:⚡|amp)(?:4(.+))?\z`)
	if len(r.SubexpNames()) != 2 {
		panic("ampAttrRE must have 1 subexpression")
	}
	return r
}()

// The allowed AMP formats, and their serialization as an html "amp4" attribute.
var ampFormatSuffixes = map[rpb.Request_HtmlFormat]string{
	rpb.Request_AMP:       "",
	rpb.Request_AMP4ADS:   "ads",
	rpb.Request_AMP4EMAIL: "email",
}

// The keys from ampFormatSuffixes.
var ampFormats = func() []rpb.Request_HtmlFormat {
	ret := []rpb.Request_HtmlFormat{}
	for v, _ := range ampFormatSuffixes {
		ret = append(ret, v)
	}
	return ret
}()

// isAllowed returns true iff the declared format, as parsed from the html amp
// attribute, corresponds to one of the allowed formats. If allowedFormats is
// empty, then any AMP format is allowed.
func isAllowed(declaredFormat string, allowedFormats []rpb.Request_HtmlFormat) bool {
	// Default to all formats.
	if len(allowedFormats) == 0 {
		allowedFormats = ampFormats
	}
	for _, allowedFormat := range allowedFormats {
		suffix, ok := ampFormatSuffixes[allowedFormat]
		if ok && declaredFormat == suffix {
			return true
		}
	}
	return false
}

// requireAMPAttribute returns an error if the <html> tag doesn't contain an
// attribute indicating that the document is AMP.
func requireAMPAttribute(dom *amphtml.DOM, allowedFormats []rpb.Request_HtmlFormat) error {
	for _, attr := range dom.HTMLNode.Attr {
		if attr.Namespace == "" {
			if match := ampAttrRE.FindStringSubmatch(attr.Key); match != nil {
				if isAllowed(match[1], allowedFormats) {
					return nil
				}
			}
		}
	}
	return errors.New("html tag is missing an AMP attribute")
}

// extractPreloads returns a list of absolute URLs of the resources to preload,
// in the order to preload them. It depends on transformers.ReorderHead having
// run.
func extractPreloads(dom *amphtml.DOM) []*rpb.Metadata_Preload {
	preloads := []*rpb.Metadata_Preload{}
	for child := dom.HeadNode.FirstChild; child != nil; child = child.NextSibling {
		switch child.DataAtom {
		case atom.Script:
			if src, ok := htmlnode.GetAttributeVal(child, "src"); ok {
				preloads = append(preloads, &rpb.Metadata_Preload{Url: src, As: "script"})
			}
		case atom.Link:
			if rel, ok := htmlnode.GetAttributeVal(child, "rel"); ok {
				if rel == "stylesheet" {
					if href, ok := htmlnode.GetAttributeVal(child, "href"); ok {
						preloads = append(preloads, &rpb.Metadata_Preload{Url: href, As: "style"})
					}
				}
			}
		}
		if len(preloads) == maxPreloads {
			break
		}
	}
	return preloads
}

// Process will parse the given request, which contains the HTML to
// transform, applying the requested list of transformers, and return the
// transformed HTML and list of resources to preload (absolute URLs), or an
// error.
//
// If the requested list of transformers is empty, apply the default.
func Process(r *rpb.Request) (string, *rpb.Metadata, error) {
	doc, err := html.Parse(strings.NewReader(r.Html))
	if err != nil {
		return "", nil, errors.Wrap(err, "Error parsing input HTML")
	}
	dom, err := amphtml.NewDOM(doc)
	if err != nil {
		return "", nil, err
	}

	if err := requireAMPAttribute(dom, r.AllowedFormats); err != nil {
		return "", nil, err
	}

	fns := configMap[r.Config]
	if r.Config == rpb.Request_CUSTOM {
		for _, val := range r.Transformers {
			fn, ok := transformerFunctionMap[strings.ToLower(val)]
			if !ok {
				return "", nil, errors.Errorf("transformer doesn't exist: %s", val)
			}
			fns = append(fns, fn)
		}
	}
	u, err := url.Parse(r.DocumentUrl)
	if err != nil {
		return "", nil, err
	}
	c := transformers.Context{dom, u, r}
	if err := runTransformers(&c, fns); err != nil {
		return "", nil, err
	}
	var o strings.Builder
	err = printer.Print(&o, c.DOM.RootNode)
	if err != nil {
		return "", nil, err
	}
	return o.String(), &rpb.Metadata{Preloads: extractPreloads(dom)}, nil
}
