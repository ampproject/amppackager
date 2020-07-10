package transformers

import (
	"regexp"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

type rewritable interface {
	// Rewrite the URLs within
	rewrite(string, *url.URL, string, nodeMap)
}

type urlRewriteContext []rewritable

// textNodeContext qualifies a text node, whose entire data value has URLs embedded
// within. This struct keeps track of where the URLs occur in the body of the text.
type textNodeContext struct {
	node    *html.Node
	offsets []amphtml.SubresourceOffset
}

// This matches strings that end with ".js".
var jsRE = func() *regexp.Regexp {
	r = regexp.MustCompile("^(.)*\.js")
	return r
}()

// This matches strings that end with ".sxg.js".
var sxgJsRE = func() *regexp.Regexp {
	r = regexp.MustCompile("^(.)*\.sxg\.js")
	return r
}()

// ScriptSrcRewrite rewrites the value of src in <script> elements, where applicable.
// If the value is of the form "*.js" (and not already "*.sxg.js"), replace it with "*.sxg.js".
func ScriptSrcRewrite(e *Context) error {
	for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
		if n.Type != html.ElementNode {
			continue
		}

		if n.DataAtom == atom.Script {
			srcVal, srcOk := htmlnode.GetAttributeVal(n, "", "src")
			if srcOk {
				matchJsRE := jsRE.MatchString(srcVal)
				matchSxgJsRE := sxgJsRE.MatchString(srcVal)
				if matchJsRE && !matchSxgJsRE {
					// Rewrite logic here
					continue
				}
			}
		} else {
			continue
		}
	}
	return nil
}