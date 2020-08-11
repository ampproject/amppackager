package transformers

import (
	"regexp"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

// AMPRuntimeJS rewrites the value of src in script nodes, where applicable.
// If the value is of the form "*.js", replace it with "*.sxg.js".
func AMPRuntimeJS(e *Context) error {
	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		if n.Type != html.ElementNode {
			continue
		}
		if n.DataAtom == atom.Script {
			if src, ok = htmlnode.FindAttribute(n, "", "src"); ok && strings.hasPrefix(src.Val, amphtml.AMPCacheRootURL) && strings.HasSuffix(src.Val, ".js"){
				src.Val = strings.TrimSuffix(src.Val, ".js") + ".js&f=sxg"
			}
		} else {
			continue
		}
	}
	return nil
}