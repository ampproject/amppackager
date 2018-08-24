package transformer

import (
	"github.com/ampproject/amppackager/internal/pkg/amphtml"
	"github.com/ampproject/amppackager/internal/pkg/htmlnode"
)

// TransformedIdentifierTransformer identifies that transformations
// were made for a specific platform on this document.
func TransformedIdentifierTransformer(e *Engine) {
	dom, ok := amphtml.NewDOM(e.Doc)
	if !ok {
		return
	}

	htmlnode.SetAttribute(dom.HTMLNode, "", "transformed", "google")
}
