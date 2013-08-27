# SHOR

```
description: "Simple hierarchical ordered representation for data"

features: {
	"Humans and machines can read and write it"
	"One container type with list and key/value features"
	"Keys can be repeated"
	"Order is significant"
	"Simple type system"
	"Hierarchical"
}

# Some examples of usage.  Details below.  Yay, comments are allowed!

examples: { # End of line comments work, too.
	list: {"red" "yellow" "blue"}
	keyval: {
		type: "name tag"
		name: "Bort"
		stock-count: 0
		restock: true
	}
	mixed: {
		"naked value"
		answer: 42
		answer: "I already told you, it's 42!"
		nested: {neato:true woo:"hoo"}
	}
	matrix: {{1 0 0 0}
		 {0 1 0 0}
		 {0 0 0 0}
		 {0 0 1 0}}
}

go-get-it: "go get github.com/ianremmler/shor/..."
```

I wasn't satisfied with the languages available for representing data in a way
that is easy for machines to understand, and more importantly, that is nice for
people to work with.

[JSON](http://json.org) came close, but its top level brackets or braces,
quoted keys, unordered maps, compulsory commas (but not after the last thing,
even if you want to), and lack of comments made me a bit sad.

[YAML](http://yaml.org) is quite readable.  You can simulate an ordered
key/value list with an array where each element is a map with a single
key/value pair, but it's kludgey.  The biggest drawback of YAML is complexity -
its spec is more than 80 pages long.

[TOML](https://github.com/mojombo/toml) addresses some of these issues.  It's
more readable than JSON and simpler than YAML.  It does have its quirks,
though.  On one hand it has a first class time type, but on the other it does
not allow scientific notation for floats.  Arrays must be homogeneous.  Maps
have an INI-style syntax (which you may or may not prefer - I do not) and are
unordered.

So I decided to make my own data language - a simple hierarchical ordered
representation - SHOR for short.

## Document

A SHOR document is list of nodes.  There are no curly braces around the list as
there are for all other SHOR lists, hopefully improving readability and
conciseness.

## Types

### List

A list is an container of 0 or more nodes, surrounded by curly braces.  The
nodes are ordered, and more than one may use the same key.  Because of this,
they do not map directly to the typical unordered map in may programming
languages.  SHOR is optimized for human intelligibility, and this trade off
means implementations might be a bit more complicated, but it's worth it.

### Node

A node consists of an optional key, followed by a ':', and a required value.
The value may be a list or any of the basic value types.

### Key

A Unicode letter, possibly followed by Unicode letters, numbers,
dashes, and underscores.

### Value

* Boolean: true or false.
* String: UTF-8 text surrounded by double quotes.
  Double quotes in the string must be escaped with a backslash.
* Number: Double precision floating point, without any of the Inf or NaN
  nonsense.

## Commands

The [shor](https://github.com/ianremmler/shor/blob/master/cmd/shor/shor.go)
command parses input and dumps it back out with configurable formatting.

## Examples

See the [example](https://github.com/ianremmler/shor/tree/master/example)
directory for some files translated to SHOR from other formats, and
[example/query.go](https://github.com/ianremmler/shor/blob/master/example/query.go)
which demonstrates how to use the library to parse a document and query the
result.

## Go get it

`go get github.com/ianremmler/shor/...`
