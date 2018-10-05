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
	"strings"

	"github.com/ampproject/amppackager/transformer/printer"
	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/ampproject/amppackager/transformer/transformers"
	"github.com/pkg/errors"
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
	"transformedidentifier": transformers.TransformedIdentifier,
	"url":                   transformers.URL,
}

// The map of config to the list of transformers, in the order in
// which they should be executed.
var configMap = map[rpb.Request_TransformersConfig][]func(*transformers.Context) error{
	rpb.Request_DEFAULT: {
		// NodeCleanup should be first.
		transformers.NodeCleanup,
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
	rpb.Request_NONE: {
		// TODO(alin04): Despite config being NONE, we still
		// must run NodeCleanup for comparison against cpp parser/lexer.
		// Once cpp parser is fully obsoleted, this can be removed.
		transformers.NodeCleanup,
	},
	rpb.Request_VALIDATION: {
		transformers.ReorderHead,
	},
	rpb.Request_CUSTOM: {},
}

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

// Process will parse the given request, which contains the HTML to
// transform, applying the requested list of transformers, and return the
// transformed HTML, or an error.
// If the requested list of transformers is empty, apply the default.
func Process(r *rpb.Request) (string, error) {
	doc, err := html.Parse(strings.NewReader(r.Html))
	if err != nil {
		return "", errors.Wrap(err, "Error parsing input HTML")
	}

	fns := configMap[r.Config]
	if r.Config == rpb.Request_CUSTOM {
		for _, val := range r.Transformers {
			fn, ok := transformerFunctionMap[strings.ToLower(val)]
			if !ok {
				return "", errors.Errorf("transformer doesn't exist: %s", val)
			}
			fns = append(fns, fn)
		}
	}
	u, err := url.Parse(r.DocumentUrl)
	if err != nil {
		return "", err
	}
	c := transformers.Context{doc, u, r}
	if err := runTransformers(&c, fns); err != nil {
		return "", err
	}
	var o strings.Builder
	err = printer.Print(&o, c.Doc)
	if err != nil {
		return "", err
	}
	return o.String(), nil
}
