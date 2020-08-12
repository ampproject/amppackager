package transformers

import (
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// AMPRuntimeJS rewrites the value of src in script nodes, where applicable.
// If the value is of the form "*.js", replace it with "*.sxg.js".
func AMPRuntimeJS(e *Context) error {
	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		if n.Type != html.ElementNode {
			continue
		}
		if n.DataAtom == atom.Script {
			src, ok := htmlnode.FindAttribute(n, "", "src")
			if ok && strings.HasPrefix(src.Val, amphtml.AMPCacheRootURL) && strings.HasSuffix(src.Val, ".js") {
				src.Val = strings.TrimSuffix(src.Val, ".js") + ".sxg.js"
			}
		} else {
			continue
		}
	}
	return nil
}
