// Package transform invokes the golang HTML parser, executes any specified
// transfomers, and prints the output to the provided string.
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
// the command line. Please keep alphabetical.
var transformerFunctionMap = map[string]func(*transformer.Engine){
	"AMPBoilerplateTransformer":      transformer.AMPBoilerplateTransformer,
	"ReorderHeadTransformer":         transformer.ReorderHeadTransformer,
	"ServerSideRenderingTransformer": transformer.ServerSideRenderingTransformer,
	"URLTransformer":                 transformer.URLTransformer,
}

// Process will parse the given HTML byte array, and execute the named
// transformers, returning the transformed HTML, or an error.
// TODO(b/112356610): Clean up these args into a proto.
func Process(data, docURL string, transformers []string) (string, error) {
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
