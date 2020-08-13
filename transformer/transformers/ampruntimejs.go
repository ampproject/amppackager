package transformers

import (
	"net/url"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// AMPRuntimeJS rewrites the value of src in script nodes, where applicable.
// If the value is of the form "*.js", replace it with "*.js?f=sxg".
func AMPRuntimeJS(e *Context) error {
	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		if n.Type != html.ElementNode {
			continue
		}
		if n.DataAtom == atom.Script {
			src, ok := htmlnode.FindAttribute(n, "", "src")
			if ok && strings.HasPrefix(src.Val, amphtml.AMPCacheRootURL) {
				u, _ := url.Parse(src.Val)
				query, _ := url.ParseQuery(u.RawQuery)
				path := u.Path
				if strings.HasSuffix(path, ".js") {
					query.Set("f", "sxg")
					u.RawQuery = query.Encode()
					src.Val = u.String()
				}
			}
		}
	}
	return nil
}
