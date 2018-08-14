package htmlnode

import "golang.org/x/net/html"

// Stack is a stack of nodes.
type Stack []*html.Node

// Push pushes a node onto the stack
func (s *Stack) Push(n *html.Node) {
	*s = append(*s, n)
}

// Pop pops the stack. It will panic if s is empty.
func (s *Stack) Pop() *html.Node {
	i := len(*s)
	n := (*s)[i-1]
	*s = (*s)[:i-1]
	return n
}
