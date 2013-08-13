thor
====

Thor is a textual hierarchical ordered representation for data.

I couldn't find a language that did what I wanted, so I made a new one.  Some of the
features of thor include:

* Humans and machines can read and write it
* Easy to parse
* Hierarchical
* Order is preserved
* Repeated keys
* Simple types
* Key/value and list style collections

Types
-----

**Node**

A node consists of an optional key and a required value.  The value may be a
list or any of the basic value types.

**Key**

A Unicode letter, possibly followed by Unicode letters, numbers, and underscores

**Value**

* Boolean: true or false
* String: text surrounded by double quotes.
  Double quotes in the string must be escaped with a backslash.
* Number: Double precision floating point, without any of the Inf and NaN nonsense.

**List**

A list is an container of 0 or more nodes.  The order of the nodes is
preserved.  A list may contain multiple nodes with the same key.

Document
--------

A thor document has a special top level list of nodes that constitute the
document.  The top level list is not surrounded by curly braces like other
lists.

Commands
--------

Currently the `thor` command only parses input and dumps it back out.
