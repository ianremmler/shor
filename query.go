package shor

// Query is a slice of Node pointers that represent the set of nodes that match a query.
type Query []*Node

// Range returns a range within the set of nodes that matches the query.
//
// If the given key is "*", all nodes are considered matching.
// If an index is negative, its magnitude represents the position from the end of the set, where
// -1 is the last node, -2 is second to last, and so on.
//
// The above rules apply to all queries.
func (q Query) Range(key string, start, end int) Query {
	match := Query{}
	if key == "*" { // all nodes match
		for _, node := range q {
			s, e := calcRange(start, end, len(node.Kids))
			match = append(match, node.Kids[s:e]...)
		}
	} else {
		for _, node := range q {
			submatch := Query{}
			for _, kid := range node.Kids {
				if kid.Key == key {
					submatch = append(submatch, kid)
				}
			}
			s, e := calcRange(start, end, len(submatch))
			match = append(match, submatch[s:e]...)
		}
	}
	return match
}

// All returns all matching nodes
func (q Query) All(key string) Query {
	return q.Range(key, 0, -1)
}

// At returns the nth matching node
func (q Query) At(key string, n int) Query {
	return q.Range(key, n, n)
}

// First returns the first matching node
func (q Query) First(key string) Query {
	return q.At(key, 0)
}

// Last returns the last matching node
func (q Query) Last(key string) Query {
	return q.At(key, -1)
}

// FirstN returns the first n matching nodes
func (q Query) FirstN(key string, n int) Query {
	return q.Range(key, 0, n-1)
}

// LastN returns the last n matching nodes
func (q Query) LastN(key string, n int) Query {
	return q.Range(key, -n, -1)
}

// calcRange calculates node indices base on given start and end positons.
func calcRange(start, end, max int) (int, int) {
	if start < 0 {
		start += max
	}
	if end < 0 {
		end += max
	}
	end++
	if end <= start || start < 0 || end < 0 {
		return 0, 0
	}
	if start > max {
		start = max
	}
	if end > max {
		end = max
	}
	return start, end
}
