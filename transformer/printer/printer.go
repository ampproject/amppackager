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

/*
Package printer emits the given html.Node as HTML text to an io.Writer.

This is forked from golang's render.go due to customizations needed for AMP HTML.
*/
package printer

import (
	"bufio"
	"io"
	"sort"
	"strings"

	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"github.com/pkg/errors"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

type writer interface {
	io.Writer
	io.ByteWriter
	WriteString(string) (int, error)
}

// Print emits the given Node to the given Writer.
// - Comments are skipped and not emitted.
// - Unnecessary quotes are dropped for attribute values.
func Print(w io.Writer, n *html.Node) error {
	if x, ok := w.(writer); ok {
		return render(x, n)
	}
	buf := bufio.NewWriter(w)
	if err := render(buf, n); err != nil {
		return err
	}
	return buf.Flush()
}

// isFirstNode returns true if n is the first of its siblings that is rendered.
func isFirstNode(n *html.Node) bool {
	for n = n.PrevSibling; n != nil; n = n.PrevSibling {
		// Comments are not rendered.
		if n.Type != html.CommentNode {
			return false
		}
	}
	return true
}

func render(w writer, n *html.Node) error {
	// Render non-element nodes; these are the easy cases.
	switch n.Type {
	case html.ErrorNode:
		return errors.New("html: cannot render an ErrorNode node")
	case html.TextNode:
		if n.Data == "" {
			return nil
		}
		if (n.Parent.DataAtom == atom.Pre || n.Parent.DataAtom == atom.Textarea) && isFirstNode(n) && (n.Data[0] == '\r' || n.Data[0] == '\n') {
			// Emit a line feed that will be summarily dropped
			// if re-parsed again. This is needed for idempotency,
			// when there are multiple newlines at the start of a
			// <pre> tag.
			// Here's the spec that says the leading newline is stripped:
			// https://html.spec.whatwg.org/multipage/grouping-content.html#the-pre-element
			// Here's where the go parser drops it:
			// https://github.com/golang/net/blob/26e67e76b6c3f6ce91f7c52def5af501b4e0f3a2/html/parse.go#L779
			n.Data = "\n" + n.Data
		}
		// TODO(b/78471903): Minimize extraneous whitespace.
		if _, err := w.WriteString(html.EscapeString(n.Data)); err != nil {
			return err
		}
		return nil
	case html.DocumentNode:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if err := render(w, c); err != nil {
				return err
			}
		}
		return nil
	case html.ElementNode:
		return renderElementNode(w, n)
	case html.CommentNode:
		// Comments are skipped.
		return nil
	case html.DoctypeNode:
		if _, err := w.WriteString("<!doctype "); err != nil {
			return err
		}
		if _, err := w.WriteString(n.Data); err != nil {
			return err
		}
		// Any additional attrs are ignored and not emitted. This is
		// redundant because the transformer already strips doctype attrs.
		// Belts and suspenders.
		return w.WriteByte('>')
	default:
		return errors.New("html: unknown node type")
	}
}

// isHTMLType does the encoding attribute check for isHTMLIntegrationPoint.
func isHTMLType(val *string) bool {
	// strings.EqualFold performs UTF-8 case-folding, but since the strings
	// on the right are ASCII-only, this should be equivalent to the
	// spec-requested ASCII case-folding.
	return val != nil && (strings.EqualFold(*val, "text/html") || strings.EqualFold(*val, "application/xhtml+xml"))
}

// isSVGForeignObject does the SVG check for isHTMLIntegrationPoint.
func isSVGForeignObject(a atom.Atom) bool {
	return a == atom.ForeignObject || a == atom.Desc || a == atom.Title
}

// isHTMLIntegrationPoint returns true iff the node is an HTML integration
// point, according to
// https://html.spec.whatwg.org/multipage/parsing.html#html-integration-point.
// This should be equivalent to htmlIntegrationPoint in
// https://github.com/golang/net/blob/master/html/foreign.go.
func isHTMLIntegrationPoint(n *html.Node) bool {
	return n.Type == html.ElementNode && ((n.Namespace == "math" && n.DataAtom == atom.AnnotationXml && isHTMLType(htmlnode.GetAttributeValOrNil(n, "", "encoding"))) ||
		(n.Namespace == "svg" && isSVGForeignObject(n.DataAtom)))
}

// isForeignElement returns true iff the node is a foreign element. Specified
// in https://html.spec.whatwg.org/multipage/parsing.html; search for "insert a
// foreign element".
func isForeignElement(n *html.Node) bool {
	return n.Type == html.ElementNode && (n.DataAtom == atom.Svg || n.DataAtom == atom.Math)
}

func renderElementNode(w writer, n *html.Node) error {
	if n.Type != html.ElementNode {
		// The provided node must be an ElementNode.
		return errors.Errorf("html: node %+v is not an element node", n)
	}
	// Render the <xxx> opening tag.
	if err := w.WriteByte('<'); err != nil {
		return err
	}
	if _, err := w.WriteString(strings.ToLower(n.Data)); err != nil {
		return err
	}
	// Sort attributes so <foo y x:y=bar> will emit as <foo x:y=bar y>
	sort.SliceStable(n.Attr, func(i, j int) bool {
		return strings.Compare(sortKey(n.Attr[i]), sortKey(n.Attr[j])) < 0
	})
	for _, a := range n.Attr {
		if err := renderElementAttr(w, a); err != nil {
			return err
		}
	}
	if voidElements[n.DataAtom] {
		if n.FirstChild != nil {
			return errors.Errorf("html: void element <%s> has child nodes", n.Data)
		}
		inForeignContent := false
		// Look for the innermost node that defines which parsing mode the parser will be in.
		for p := n; p != nil; p = p.Parent {
			if isHTMLIntegrationPoint(p) {
				break
			} else if isForeignElement(p) {
				inForeignContent = true
				break
			}
		}
		var str string
		if inForeignContent {
			// Space is required before solidus, so that it isn't interpreted as part of an attribute value.
			str = " />"
		} else {
			str = ">"
		}
		_, err := w.WriteString(str)
		return err
	}
	if err := w.WriteByte('>'); err != nil {
		return err
	}

	// Render any child nodes.
	switch n.Data {
	// The original golang renderer emits raw HTML for 8 tags (see
	// https://github.com/golang/net/blob/master/html/render.go#L196
	//
	// This printer only emits raw HTML for 4 tags, ignoring
	// noembed, noframes, plaintext, and xmp, which are unsupported
	// by the AMP validator (see
	// https://github.com/ampproject/amphtml/blob/master/validator/validator-main.protoascii
	// The result is that for the 4 ignored tags, their textnodes will
	// be escaped.
	case "iframe", "noscript", "script", "style":
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			// Emit the raw text elements, unescaped.
			if c.Type == html.TextNode {
				if _, err := w.WriteString(c.Data); err != nil {
					return err
				}
			} else {
				if err := render(w, c); err != nil {
					return err
				}
			}
		}
	default:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if err := render(w, c); err != nil {
				return err
			}
		}
	}

	// Render the </xxx> closing tag.
	if _, err := w.WriteString("</"); err != nil {
		return err
	}
	if _, err := w.WriteString(strings.ToLower(n.Data)); err != nil {
		return err
	}
	err := w.WriteByte('>')
	return err
}

// sortKey generates a sorting key from the given Attribute from its
// namespace and key. It ignores the value.
func sortKey(a html.Attribute) string {
	if a.Namespace == "" {
		return a.Key
	}
	return a.Namespace + ":" + a.Key
}

func renderElementAttr(w writer, a html.Attribute) error {
	if err := w.WriteByte(' '); err != nil {
		return err
	}
	if a.Namespace != "" {
		if _, err := w.WriteString(a.Namespace); err != nil {
			return err
		}
		if err := w.WriteByte(':'); err != nil {
			return err
		}
	}
	if _, err := w.WriteString(strings.ToLower(a.Key)); err != nil {
		return err
	}
	// If attribute has no values, only output the attribute name.
	if a.Val != "" {
		if _, err := w.WriteString(`=`); err != nil {
			return err
		}
		if err := writeQuoted(w, html.EscapeString(a.Val)); err != nil {
			return err
		}
	}
	return nil
}

// writeQuoted writes s to w surrounded (optionally) by quotes.
//
// From http://www.w3.org/TR/1999/REC-html401-19991224/html40.txt
// "By default, SGML requires that all attribute values be delimited
// using either double quotation marks (ASCII decimal 34) or single
// quotation marks (ASCII decimal 39)... The attribute value may
// only contain letters (a-z [0x61-0x7A] and A-Z [0x41-0x5A]),
// digits (0-9 [0x30-0x39), hyphens (ASCII decimal 45 [0x2D]),
// periods (ASCII decimal 46 [0x2E]), underscores (ASCII decimal 95
// [0x5F]), and colons (ASCII decimal 58 [0x3A]). We recommend using
// quotation marks even when it is possible to eliminate them."
//
// In order to reduce the output size of the quoted string, a relaxed
// rule is used instead: all ASCII characters which are printable but not
// in the set { 0x20(space), 0x22("), 0x27('), 0x3E(>), 0x60(`) } are
// treated as not needing quotes.
//
// The relaxed rule is verified acceptable by FF 1.5/2/3, IE 6/7/8,
// Safari, Chrome, and Opera.
func writeQuoted(w writer, s string) error {
	q := pickQuote(s)
	if _, err := w.WriteString(q); err != nil {
		return err
	}
	if _, err := w.WriteString(s); err != nil {
		return err
	}
	_, err := w.WriteString(q)
	return err
}

const needsQuotes = " \t\r\n\f\"'=<>`"

// pickQuote determines if the given string can be unquoted, or needs single
// or double quotes. It uses rules based on the spec
// https://html.spec.whatwg.org/multipage/syntax.html#attributes-2
func pickQuote(s string) string {
	if len(s) == 0 {
		return "\""
	}

	if strings.ContainsAny(s, needsQuotes) {
		// The string is already escaped, so default to using double quotes.
		return "\""
	}
	return ""
}

// Void elements only have a start tag and no content.
var voidElements = map[atom.Atom]bool{
	atom.Area:    true,
	atom.Base:    true,
	atom.Br:      true,
	atom.Col:     true,
	atom.Command: true,
	atom.Embed:   true,
	atom.Hr:      true,
	atom.Img:     true,
	atom.Input:   true,
	atom.Keygen:  true,
	atom.Link:    true,
	atom.Meta:    true,
	atom.Param:   true,
	atom.Source:  true,
	atom.Track:   true,
	atom.Wbr:     true,
}
