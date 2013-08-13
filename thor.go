package thor

import (
	"io"
	"strconv"

	"github.com/skelterjohn/gopp"
)

const grammar = `
ignore: /^#.*\n/
ignore: /^(?:[ \t\n])+/

Doc => {type=Node} {field=Type} {root} {field=Kids} <<Node>>*
Node => {type=Node} {field=Key} <id> ':' {field=.} <Value>
Node => {type=Node} {field=.} <Value>
Value => {type=Node} {field=Type} {table} '{' {field=Kids} <<Node>>* '}'
Value => {type=Node} {field=Type} {num} {field=Val} <num>
Value => {type=Node} {field=Type} {bool} {field=Val} <bool>
Value => {type=Node} {field=Type} {str} {field=Val} <str>

num = /([-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?)/
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
	decFact.RegisterType(Node{})
}

type Node struct {
	Key  string
	Val  string
	Type string
	Kids []Node
}

func (n Node) String() string {
	s := ""
	switch n.Type {
	case "root", "table":
		for i := range n.Kids {
			s += n.Kids[i].String()
			if i < len(n.Kids)-1 {
				s += " "
			}
		}
	case "str":
		s = strconv.Quote(n.Val)
	default:
		s = n.Val
	}
	if n.Type == "table" {
		s = "{" + s + "}"
	}
	if n.Key != "" {
		s = n.Key + ":" + s
	}
	return s
}

func Parse(in io.Reader) (*Node, error) {
	dec := decFact.NewDecoder(in)
	tree := &Node{}
	err := dec.Decode(tree)
	return tree, err
}
