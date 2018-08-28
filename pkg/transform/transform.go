// Package transform invokes the golang HTML parser, executes the
// individual transfomers (unless overridden), and prints the output
// to the provided string.
package transform

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ampproject/amppackager/pkg/printer"
	rpb "github.com/ampproject/amppackager/pkg/transform
	"github.com/ampproject/amppackager/pkg/transformer"
	"golang.org/x/net/html"
)

// Transformer functions must be added here in order to be passed in from
// the command line or invoked from other languages. Please keep alphabetical.
//
// NOTE: The string mapping is necessary as a cross-over to allow
// invocation from C/C++.
var transformerFunctionMap = map[string]func(*transformer.Engine){
	"AMPBoilerplateTransformer":        transformer.AMPBoilerplateTransformer,
	"AMPRuntimeCSSTransformer":         transformer.AMPRuntimeCSSTransformer,
	"LinkTagTransformer":               transformer.LinkTagTransformer,
	"MetaTagTransformer":               transformer.MetaTagTransformer,
	"ReorderHeadTransformer":           transformer.ReorderHeadTransformer,
	"ServerSideRenderingTransformer":   transformer.ServerSideRenderingTransformer,
	"TransformedIdentifierTransformer": transformer.TransformedIdentifierTransformer,
	"URLTransformer":                   transformer.URLTransformer,
}

// The default set of transformers to execute, in the order in which
// to execute them.
var defaultTransformers = []string{
	"MetaTagTransformer",
	"LinkTagTransformer",
	"URLTransformer",
	"AMPBoilerplateTransformer",
	"ServerSideRenderingTransformer",
	// AmpRuntimeCssTransformer must run after ServerSideRenderingTransformer
	"AMPRuntimeCSSTransformer",
	"TransformedIdentifierTransformer",
	// ReorderHeadTransformer should run after all transformers that modify the
	// <head>, as they may do so without preserving the proper order.
	"ReorderHeadTransformer",
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

	transformers := r.Transformers
	if len(transformers) == 0 {
		transformers = defaultTransformers
	}
	fns := []func(*transformer.Engine){}
	for _, val := range transformers {
		fn, ok := transformerFunctionMap[val]
		if !ok {
			return "", fmt.Errorf("transformer doesn't exist: %s", val)
		}
		fns = append(fns, fn)
	}
	u, err := url.Parse(r.DocumentUrl)
	if err != nil {
		return "", err
	}
	e := transformer.Engine{doc, u, fns, r}
	e.Transform()
	var o strings.Builder
	err = printer.Print(&o, e.Doc)
	if err != nil {
		return "", err
	}
	return o.String(), nil
}
