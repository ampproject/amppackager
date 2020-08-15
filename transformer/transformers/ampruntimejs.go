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
	for n := e.DOM.HeadNode.FirstChild; n != nil; n = n.NextSibling {
		if n.Type == html.ElementNode && n.DataAtom == atom.Script {
			src, ok := htmlnode.FindAttribute(n, "", "src")
			if ok && strings.HasPrefix(src.Val, amphtml.AMPCacheRootURL) {
				u, uerr := url.Parse(src.Val)
				if uerr != nil {
					continue
				}
				query, queryerr := url.ParseQuery(u.RawQuery)
				if queryerr != nil {
					continue
				}
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
