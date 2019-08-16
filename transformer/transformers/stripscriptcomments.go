package transformers

import (
	"strings"
)

import (
	"github.com/ampproject/amppackager/transformer/internal/htmlnode"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Strips html style comments from all the <script> elements that are child
// nodes of the given node. Example:
// <script><!-- var a = 1; --></script> transforms to empty script node.
// <script><!--[if IE]-->var a=1;<![endif]--></script> transforms to empty
// script tag.
//
// Attributes and any content outside comment are preserved.
// <script type="text/javascript"><!-- hello -->var a = 0;</script>
//            transforms to
// <script type="text/javascript">var a = 0;</script>
//
// Any comment that looks like a comment but is a string variable is preserved.
// <script>var comment = "<!-- comment -->";</script> is not modified.
func StripScriptComments(e *Context) error {
  for n := e.DOM.RootNode; n != nil; n = htmlnode.Next(n) {
    if n.Type != html.ElementNode {
      continue
    }

    if n.DataAtom == atom.Script {
      typeVal, typeOk := htmlnode.GetAttributeVal(n, "", "type")
      if typeOk && (typeVal == "application/json" ||
                    typeVal == "application/ld+json") {
        contents := n.FirstChild
	if contents == nil || len(contents.Data) == 0 {
		return nil
	}

	contents.Data = stripComments(&contents.Data)
      }
    }
  }

  return nil
}

type commentStripper struct {
	Content string
	index int
	insideQuotes bool
	quoteChar byte
}

func (r *commentStripper) readNextByte() byte {
	r.index = r.index + 1
	if r.index >= len(r.Content) {
		return 0
	}
	return r.Content[r.index]
}

func (r *commentStripper) readByteAt(i int) byte {
	if i < 0 || i >= len(r.Content) {
		return 0
	}

	return r.Content[i]
}

func (r *commentStripper) isPreviousEscapeChar() bool {
	previous := r.index - 1
	c := r.readByteAt(previous)
	if c != '\\' {
		return false
	}

	// Check if '\' is escaped by a preceding '\', consume all preceding '\'.
	for {
		previous = previous - 1
		c := r.readByteAt(previous)
		if c != '\\' {
			break
		}
	}

	return (r.index - previous) % 2 != 0
}

func (r *commentStripper) isQuoteChar(c byte) bool {
	return c == '"' || c == '\''
}

func (r *commentStripper) skipComment() bool {
	cdataStyleComment := false
	// Checks initial !-- characters.
	if r.readByteAt(r.index + 1) != '!' && r.readByteAt(r.index + 2) != '-' && r.readByteAt(r.index + 3) != '-' {
		return false
	}

	// Skips the '!--' chars.
	r.index = r.index + 3

	// Checks if comment is of [CDATA] format.
	// <!--[condition]>, we are interested in <!--[ in which case comment
	// will end with ]-->
	if r.readByteAt(r.index + 1) == '[' {
		cdataStyleComment = true
	}

	// Skip until we encounter -->
	for {
		c := r.readNextByte()
		if c == 0 {
			// Skip everything until eof.
			return true
		}

		if c == '-' {
			if cdataStyleComment {
				if r.readByteAt(r.index - 1) != ']' {
					continue
				}
			}

			if r.readByteAt(r.index + 1) == '-' && r.readByteAt(r.index + 2) == '>' {
				r.index = r.index + 2
				return true
			}
		}
	}
}

func (r *commentStripper) stripComments() string {
	var buffer strings.Builder
	for {
		c := r.readNextByte()
		if c == 0 {
			break
		}

		if r.insideQuotes {
			if r.quoteChar == c && !r.isPreviousEscapeChar() {
				r.insideQuotes = false
				r.quoteChar = 0
			}
			buffer.WriteByte(c)
			continue

		}

		if r.isQuoteChar(c) {
			r.insideQuotes = true
			r.quoteChar = c
		        buffer.WriteByte(c)
			continue
		}

		// Normal character.
		if c == '<' && r.skipComment() {
			continue
		}

		buffer.WriteByte(c)
	}

	return buffer.String()
}

func stripComments(ss *string) string {
	cs := commentStripper{Content: *ss, insideQuotes: false, index: -1}
	return cs.stripComments()
}
