// Package shor implements parsing and querying of the shor language, a simple hierarchical
// ordered representation for data.
package shor

import (
	"io"
	"strconv"
	"strings"

	"github.com/skelterjohn/gopp"
)

const grammar = `
ignore: /^#.*/
ignore: /^\s+/

Doc => {field=Key} {/} <List>
List => {field=Type} {list} {field=Kids} <<Node>>*
Node => {field=Key} <id> ':' <Value>
Node => <Value>
Value => '{' <List> '}'
Value => {field=Type} {num} {field=Val} <num>
Value => {field=Type} {bool} {field=Val} <bool>
Value => {field=Type} {str} {field=Val} <str>

num = /([-+]?\d*\.?\d+([eE][-+]?\d+)?)/
bool = /(true|false)/
str = /"((?:[^"\\]|\\.)*)"/
id = /([\pL][\pL\pN\-_]*)/
`

var (
	decFact *gopp.DecoderFactory
)

func init() {
	var err error
	decFact, err = gopp.NewDecoderFactory(grammar, "Doc")
	if err != nil {
		panic(err)
	}
	decFact.RegisterType(&Node{})
}

// Node holds a shor node and its children.
type Node struct {
	Key    string // id if keyed, empty if keyless, or "/" if root
	Val    string // string representation of value, empty for list
	Type   string // str, num, bool, or list
	Kids   []*Node
	Parent *Node // nil if root node
}

// linkNodes updates parent links for entire tree.
func (n *Node) linkNodes(par *Node) {
	n.Parent = par
	for i := range n.Kids {
		n.Kids[i].linkNodes(n)
	}
}

// String returns a tree in single-line format
func (n *Node) String() string {
	return n.Format(-1, "")
}

// Format returns a formatted string representation of a tree.
// If depth is negative, single-line format is produced.
func (n *Node) Format(depth int, indent string) string {
	isRoot := (n.Key == "/")
	isMultiline := (depth >= 0)
	ind, keySep, eltSep, listSep := "", "", " ", ""
	if isMultiline {
		ind = strings.Repeat(indent, depth)
		keySep, eltSep, listSep = " ", "\n", "\n"
	}

	s := ""
	switch n.Type {
	case "list":
		kidDepth := depth
		if !isRoot && isMultiline {
			kidDepth++
		}
		for i := range n.Kids {
			s += n.Kids[i].Format(kidDepth, indent)
			if i < len(n.Kids)-1 {
				s += eltSep
			}
		}
		if !isRoot {
			s = "{" + listSep + s + listSep + ind + "}"
		}
	case "str":
		s = strconv.Quote(n.Val)
	default:
		s = n.Val
	}
	if n.Key != "" && !isRoot {
		s = n.Key + ":" + keySep + s
	}
	if !isRoot {
		s = ind + s
	}
	return s
}

// Query returns a Query object populated with n
func (n *Node) Query() Query {
	return Query{n}
}

// Parse parses input and returns a tree of parsed shor nodes
func Parse(in io.Reader) (*Node, error) {
	dec := decFact.NewDecoder(in)
	tree := &Node{}
	err := dec.Decode(tree)
	if err != nil {
		return nil, err
	}
	tree.linkNodes(nil)
	return tree, nil
}
