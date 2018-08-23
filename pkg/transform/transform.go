// Package transform invokes the golang HTML parser, executes the
// individual transfomers (unless overridden), and prints the output
// to the provided string.
package transform

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ampproject/amppackager/pkg/printer"
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
	"LinkTagTransformer":               transformer.LinkTagTransformer,
	"MetaTagTransformer":               transformer.MetaTagTransformer,
	"ReorderHeadTransformer":           transformer.ReorderHeadTransformer,
	"ServerSideRenderingTransformer":   transformer.ServerSideRenderingTransformer,
	"TransformedIdentifierTransformer": transformer.TransformedIdentifierTransformer,
	"URLTransformer":                   transformer.URLTransformer,
}

// The transformers to execute, in the order in which to execute them.
var transformers = []string{
	"URLTransformer",
	"AMPBoilerplateTransformer",
	"LinkTagTransformer",
	"MetaTagTransformer",
	"ServerSideRenderingTransformer",
	"TransformedIdentifierTransformer",
	"ReorderHeadTransformer",
}

// Process will parse the given HTML byte array, applying all the
// transformers and return the transformed HTML, or an error.
// TODO(b/112356610): Clean up these args into a proto.
func Process(data, docURL string) (string, error) {
	return ProcessSome(data, docURL, transformers)
}

// Process will parse the given HTML byte array, and execute the named
// transformers, returning the transformed HTML, or an error.
// TODO(b/112356610): Clean up these args into a proto.
func ProcessSome(data, docURL string, transformers []string) (string, error) {
	doc, err := html.Parse(strings.NewReader(data))
	if err != nil {
		return "", err
	}

	fns := []func(*transformer.Engine){}
	for _, val := range transformers {
		fn, ok := transformerFunctionMap[val]
		if !ok {
			return "", fmt.Errorf("transformer doesn't exist: %s", val)
		}
		fns = append(fns, fn)
	}
	u, err := url.Parse(docURL)
	if err != nil {
		return "", err
	}
	e := transformer.Engine{doc, u, fns}
	e.Transform()
	var o strings.Builder
	err = printer.Print(&o, e.Doc)
	if err != nil {
		return "", err
	}
	return o.String(), nil
}
