// Package shor implements parsing and querying of the shor language, a simple hierarchical
// ordered representation for data.
package shor

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/skelterjohn/gopp"
)

const (
	Num = iota
	Bool
	Str
	List
)

var decFact *gopp.DecoderFactory

var grammar = fmt.Sprintf(`
ignore: /^#.*/
ignore: /^\s+/

Doc => {field=Key} {/} <List>
List => {field=Type} {%d} {field=Kids} <<Node>>*
Node => {field=Key} <id> ':' <Value>
Node => <Value>
Value => '{' <List> '}'
Value => {field=Type} {%d} {field=Content} <num>
Value => {field=Type} {%d} {field=Content} <bool>
Value => {field=Type} {%d} {field=Content} <str>

num = /([-+]?\d*\.?\d+([eE][-+]?\d+)?)/
bool = /(true|false)/
str = /"((?:[^"\\]|\\.)*)"/
id = /([\pL][\pL\pN\-_]*)/`, List, Num, Bool, Str)

func init() {
	var err error
	decFact, err = gopp.NewDecoderFactory(grammar, "Doc")
	if err != nil {
		panic(err)
	}
}

// Node holds a shor node and its children.
type Node struct {
	Key     string      // id if keyed, empty if keyless, or "/" if root
	Content string      // string representation of value, empty for list
	Value   interface{} // typed value, nil for list
	Type    int         // Str, Num, Bool, or List
	Parent  *Node       // nil if root node
	Kids    []*Node     // empty if not list
}

// process converts content to typed value and updates parent/child links
func (n *Node) process(par *Node) {
	n.Parent = par
	n.parseValue()
	for i := range n.Kids {
		n.Kids[i].process(n)
	}
}

// parseValue converts string content to typed value
func (n *Node) parseValue() {
	switch n.Type {
	case Num:
		if v, err := strconv.ParseFloat(n.Content, 64); err == nil {
			n.Value = v
		}
	case Bool:
		if v, err := strconv.ParseBool(n.Content); err == nil {
			n.Value = v
		}
	case Str:
		n.Value = n.Content
	}
}

// Get provides the typed value represented by the node.
func (n *Node) Get(val interface{}) bool {
	switch n.Type {
	case Num:
		if p, ok := val.(*float64); ok {
			if v, ok := n.Value.(float64); ok {
				*p = v
				return true
			}
		}
	case Bool:
		if p, ok := val.(*bool); ok {
			if v, ok := n.Value.(bool); ok {
				*p = v
				return true
			}
		}
	case Str:
		if p, ok := val.(*string); ok {
			if v, ok := n.Value.(string); ok {
				*p = v
				return true
			}
		}
	}
	return false
}

// Set updates a node with a new value.
func (n *Node) Set(val interface{}) bool {
	switch val.(type) {
	case float32, float64, int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr:

		n.Type = Num
	case bool:
		n.Type = Bool
	case string:
		n.Type = Str
	default:
		return false
	}
	n.Value = val
	n.Content = fmt.Sprint(val)
	n.Kids = []*Node{} // not a list, so drop kids
	return true
}

// Append adds kid at the end of the node's list of children.
func (n *Node) Append(kid *Node) {
	n.Kids = append(n.Kids, kid)
	kid.Parent = n
}

// Insert places kid in the node's list of children at the given position.
func (n *Node) Insert(kid *Node, pos int) bool {
	if pos < 0 || pos > len(n.Kids) {
		return false
	}
	n.Kids = append(n.Kids[:pos], append([]*Node{kid}, n.Kids[pos:]...)...)
	kid.Parent = n
	return true
}

// Remove removes kide from the node's list of children.
func (n *Node) Remove(kid *Node) bool {
	for i, k := range n.Kids {
		if k == kid {
			k.Parent = nil
			n.Kids = append(n.Kids[:i], n.Kids[i:]...)
			return true
		}
	}
	return false
}

// String returns a tree in single-line format.
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
	case List:
		kidDepth := depth
		if !isRoot && isMultiline {
			kidDepth++
		}
		for i := range n.Kids {
			s += n.Kids[i].Format(kidDepth, indent)
			if isMultiline || i < len(n.Kids)-1 {
				s += eltSep
			}
		}
		if !isRoot {
			s = "{" + listSep + s + ind + "}"
		}
	case Str:
		s = strconv.Quote(n.Content)
	default:
		s = n.Content
	}
	if n.Key != "" && !isRoot {
		s = n.Key + ":" + keySep + s
	}
	if !isRoot {
		s = ind + s
	}
	return s
}

// Query returns a Query object populated with n.
func (n *Node) Query() Query {
	return Query{n}
}

// Parse parses input and returns a tree of parsed shor nodes.
func Parse(in io.Reader) (*Node, error) {
	dec := decFact.NewDecoder(in)
	tree := &Node{}
	err := dec.Decode(tree)
	if err != nil {
		return nil, err
	}
	tree.process(nil)
	return tree, nil
}
