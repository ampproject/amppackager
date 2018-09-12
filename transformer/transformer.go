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

// Package transform invokes the golang HTML parser, executes the
// individual transfomers (unless overridden), and prints the output
// to the provided string.
package transformer

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ampproject/amppackager/transformer/printer"
	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

// Transformer functions must be added here in order to be passed in from
// the command line or invoked from other languages. Please keep alphabetical.
// Case-insensitive lookup.
//
// NOTE: The string mapping is necessary as a language cross-over to
// allow explicit transformer invocation (via the CUSTOM config).
var transformerFunctionMap = map[string]func(*transformers.Engine){
	"ampboilerplate":        transformers.AMPBoilerplateTransformer,
	"ampruntimecss":         transformers.AMPRuntimeCSSTransformer,
	"linktag":               transformers.LinkTagTransformer,
	"metatag":               transformers.MetaTagTransformer,
	"reorderhead":           transformers.ReorderHeadTransformer,
	"serversiderendering":   transformers.ServerSideRenderingTransformer,
	"transformedidentifier": transformers.TransformedIdentifierTransformer,
	"url":                   transformers.URLTransformer,
}

// The map of config to the list of transformers, in the order in
// which they should be executed.
var configMap = map[rpb.Request_TransformersConfig][]func(*transformers.Engine){
	rpb.Request_DEFAULT: {
		transformers.MetaTagTransformer,
		transformers.LinkTagTransformer,
		transformers.URLTransformer,
		transformers.AMPBoilerplateTransformer,
		transformers.ServerSideRenderingTransformer,
		// AmpRuntimeCssTransformer must run after ServerSideRenderingTransformer
		transformers.AMPRuntimeCSSTransformer,
		transformers.TransformedIdentifierTransformer,
		// ReorderHeadTransformer should run after all transformers that modify the
		// <head>, as they may do so without preserving the proper order.
		transformers.ReorderHeadTransformer,
	},
	rpb.Request_NONE: {},
	rpb.Request_VALIDATION: {
		transformers.ReorderHeadTransformer,
	},
	rpb.Request_CUSTOM: {},
}

// Override for tests.
var runTransform = func(e *transformers.Engine) {
	e.Transform()
}

// Process will parse the given request, which contains the HTML to
// transform, applying the requested list of transformers, and return the
// transformed HTML, or an error.
// If the requested list of transformers is empty, apply the default.
func Process(r *rpb.Request) (string, error) {
	doc, err := html.Parse(strings.NewReader(r.Html))
	if err != nil {
		return "", err
	}

	fns := configMap[r.Config]
	if r.Config == rpb.Request_CUSTOM {
		for _, val := range r.Transformers {
			fn, ok := transformerFunctionMap[strings.ToLower(val)]
			if !ok {
				return "", fmt.Errorf("transformer doesn't exist: %s", val)
			}
			fns = append(fns, fn)
		}
	}
	u, err := url.Parse(r.DocumentUrl)
	if err != nil {
		return "", err
	}
	e := transformers.Engine{doc, u, fns, r}
	runTransform(&e)
	var o strings.Builder
	err = printer.Print(&o, e.Doc)
	if err != nil {
		return "", err
	}
	return o.String(), nil
}
