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
	"math"
	"net/url"
	"regexp"
	"strconv"
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
	"absoluteurl":           transformers.AbsoluteURL,
	"ampboilerplate":        transformers.AMPBoilerplate,
	"ampruntimecss":         transformers.AMPRuntimeCSS,
	"linktag":               transformers.LinkTag,
	"nodecleanup":           transformers.NodeCleanup,
	"preloadimage":          transformers.PreloadImage,
	"reorderhead":           transformers.ReorderHead,
	"serversiderendering":   transformers.ServerSideRendering,
	"stripjs":               transformers.StripJS,
	"stripscriptcomments":   transformers.StripScriptComments,
	"transformedidentifier": transformers.TransformedIdentifier,
	"unusedextensions":      transformers.UnusedExtensions,
	"urlrewrite":            transformers.URLRewrite,
}

// The map of config to the list of transformers, in the order in
// which they should be executed.
var configMap = map[rpb.Request_TransformersConfig][]func(*transformers.Context) error{
	rpb.Request_DEFAULT: {
		// NodeCleanup should be first.
		transformers.NodeCleanup,
		transformers.StripJS,
		transformers.StripScriptComments,
		transformers.LinkTag,
		transformers.AbsoluteURL,
		transformers.AMPBoilerplate,
		transformers.UnusedExtensions,
		transformers.ServerSideRendering,
		transformers.AMPRuntimeCSS,
		transformers.TransformedIdentifier,
		transformers.URLRewrite,
		transformers.PreloadImage,
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

// firstSrcsetSourceRE captures the first source URL from a srcset. Caller must
// remove trailing comma, if present.
var firstSrcsetSourceRE = regexp.MustCompile(`^[\s,]*([^\s]+)`)

// The allowed AMP formats, and their serialization as an html "amp4" attribute.
var ampFormatSuffixes = map[rpb.Request_HtmlFormat]string{
	rpb.Request_AMP:       "",
	rpb.Request_AMP4ADS:   "ads",
	rpb.Request_AMP4EMAIL: "email",
}

// The keys from ampFormatSuffixes.
var ampFormats = func() []rpb.Request_HtmlFormat {
	ret := []rpb.Request_HtmlFormat{}
	for v := range ampFormatSuffixes {
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

// setDOM parses the input HTML and sets c.DOM to the parsed DOM struct.
func setDOM(c *transformers.Context, s string) error {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return errors.Wrap(err, "Error parsing input HTML")
	}

	dom, err := amphtml.NewDOM(doc)
	if err != nil {
		return err
	}
	c.DOM = dom
	return nil
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

// setBaseURL derives the absolute base URL, and sets it on c.BaseURL. The value
// is derived using the <base> href in the DOM, if it exists. If the href is
// relative, it is parsed in the context of the document URL.
// This must run after DocumentURL is set on the context.
func setBaseURL(c *transformers.Context) {
	if n, ok := htmlnode.FindNode(c.DOM.HeadNode, atom.Base); ok {
		if v, ok := htmlnode.GetAttributeVal(n, "", "href"); ok {
			if u, err := c.DocumentURL.Parse(v); err == nil {
				c.BaseURL = u
				return
			}
		}
	}
	c.BaseURL = c.DocumentURL
}

// extractPreloads returns a list of absolute URLs of the resources to preload,
// in the order to preload them. It depends on transformers.ReorderHead having
// run.
func extractPreloads(dom *amphtml.DOM) []*rpb.Metadata_Preload {
	// If you add additional preloads here, verify that they can not be
	// unintentionally author supplied.
	preloads := []*rpb.Metadata_Preload{}
	// If there are any modules present, then we will skip preloads for non-module scripts.
	hasModule := false
	for current := dom.HeadNode.FirstChild; current != nil; current = current.NextSibling {
		if current.DataAtom == atom.Script {
			if scriptType, _ := htmlnode.GetAttributeVal(current, "", "type"); scriptType == "module" {
				hasModule = true
				break
			}
		}
	}
	for current := dom.HeadNode.FirstChild; current != nil; current = current.NextSibling {
		switch current.DataAtom {
		case atom.Script:
			if src, ok := htmlnode.GetAttributeVal(current, "", "src"); ok {
				scriptType, _ := htmlnode.GetAttributeVal(current, "", "type")
				if hasModule && scriptType != "module" {
					continue
				}
				preload := &rpb.Metadata_Preload{Url: src, As: "script"}
				if crossorigin, ok := htmlnode.GetAttributeVal(current, "", "crossorigin"); ok && crossorigin != "" {
					preload.Attributes = append(preload.Attributes, &rpb.Metadata_Preload_Attribute{Key: "crossorigin", Val: crossorigin})
				}
				if scriptType == "module" {
					preload.Module = true
				}
				preloads = append(preloads, preload)
			}
		case atom.Link:
			if rel, ok := htmlnode.GetAttributeVal(current, "", "rel"); ok {
				if strings.EqualFold(rel, "stylesheet") {
					if href, ok := htmlnode.GetAttributeVal(current, "", "href"); ok {
						preloads = append(preloads, &rpb.Metadata_Preload{Url: href, As: "style"})
					}
				} else if strings.EqualFold(rel, "preload") {
					if as, ok := htmlnode.GetAttributeVal(current, "", "as"); ok && strings.EqualFold(as, "image") {
						href, _ := htmlnode.GetAttributeVal(current, "", "href")
						// It's valid for a <link> to not have a href, but when we generate a Link header it must have one.
						if href == "" {
							imagesrcset, ok := htmlnode.GetAttributeVal(current, "", "imagesrcset")
							if !ok {
								continue
							}
							// The href doesn't really matter here. Browsers will ignore it and instead prioritize whichever source is selected from imagesrcset. However, the Link header *must* have it.
							// Stub this by just finding the first source URL possible.
							firstSource := firstSrcsetSourceRE.FindStringSubmatch(imagesrcset)[1]
							if firstSource == "" {
								continue
							}
							href = strings.TrimSuffix(firstSource, ",")
						}
						preload := &rpb.Metadata_Preload{Url: href, As: "image"}
						for _, attr := range current.Attr {
							if strings.EqualFold(attr.Key, "rel") || strings.EqualFold(attr.Key, "as") || strings.EqualFold(attr.Key, "href") {
								continue
							}
							preload.Attributes = append(preload.Attributes, &rpb.Metadata_Preload_Attribute{
								Key: attr.Key,
								Val: attr.Val,
							})
						}
						preloads = append(preloads, preload)
						htmlnode.RemoveNode(&current)
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

// defaultMaxAgeSeconds is the max-age to apply when there is an inline
// amp-script without an explicit max-age. This is 1 day, to parallel the
// security precautions put in place around service workers:
// https://dev.chromium.org/Home/chromium-security/security-faq/service-worker-security-faq#TOC-Do-Service-Workers-live-forever-
const defaultMaxAgeSeconds int32 = 86400 // number of seconds in a day

// maxMaxAgeSeconds is the max duration of an SXG, per
// https://wicg.github.io/webpackage/draft-yasskin-http-origin-signed-responses.html#signature-validity.
const maxMaxAgeSeconds int32 = 7 * 86400

// computeMaxAgeSeconds returns the suggested max-age based on the presence of
// any inline <amp-script> tags on the page; callers should min() the return
// value against any other they constraints they have (e.g. the max allowed
// duration of an SXG).
func computeMaxAgeSeconds(dom *amphtml.DOM) int32 {
	var maxAge int32 = math.MaxInt32
	for node := dom.RootNode; node != nil; node = htmlnode.Next(node) {
		// The html parser downcases tag and attribute names, so we needn't.
		if node.Type == html.ElementNode && node.Data == "amp-script" && htmlnode.HasAttribute(node, "", "script") {
			nodeMaxAge := defaultMaxAgeSeconds
			if value, ok := htmlnode.GetAttributeVal(node, "", "max-age"); ok {
				if num, err := strconv.ParseInt(value, 10, 32); err == nil {
					if num < 0 {
						num = 0
					}
					nodeMaxAge = int32(num)
				}
			}
			if nodeMaxAge < maxAge {
				maxAge = nodeMaxAge
			}
		}
	}
	if maxAge > maxMaxAgeSeconds {
		maxAge = maxMaxAgeSeconds
	}
	return maxAge
}

// Process will parse the given request, which contains the HTML to
// transform, applying the requested list of transformers, and return the
// transformed HTML and list of resources to preload (absolute URLs), or an
// error.
//
// If the requested list of transformers is empty, apply the default.
func Process(r *rpb.Request) (string, *rpb.Metadata, error) {
	context := &transformers.Context{}

	if err := validateUTF8ForHTML(r.Html); err != nil {
		return "", nil, err
	}

	if err := setDOM(context, r.Html); err != nil {
		return "", nil, err
	}

	if err := requireAMPAttribute(context.DOM, r.AllowedFormats); err != nil {
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

	documentURL, err := url.Parse(r.DocumentUrl)
	if err != nil {
		return "", nil, err
	}
	context.DocumentURL = documentURL

	context.Version = r.Version
	if r.Version == 0 {
		version, err := SelectVersion(nil)
		if err != nil {
			return "", nil, err
		}
		context.Version = version
	}

	// This must run AFTER DocumentURL is parsed.
	setBaseURL(context)

	if err := runTransformers(context, fns); err != nil {
		return "", nil, err
	}
	// extractPreloads is an implicit transformer, and must run before printer.
	preloads := extractPreloads(context.DOM)
	var o strings.Builder
	if err := printer.Print(&o, context.DOM.RootNode); err != nil {
		return "", nil, err
	}
	metadata := rpb.Metadata{
		Preloads:   preloads,
		MaxAgeSecs: computeMaxAgeSeconds(context.DOM),
	}
	return o.String(), &metadata, nil
}
