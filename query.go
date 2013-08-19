package thor

type Query []*Node

func (q Query) KeyRange(key string, start, end int) Query {
	match := Query{}
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
	return match
}

func (q Query) Key(key string) Query {
	return q.KeyRange(key, 0, -1)
}

func (q Query) KeyN(key string, n int) Query {
	return q.KeyRange(key, n, n)
}

func (q Query) KeyFirst(key string) Query {
	return q.KeyN(key, 0)
}

func (q Query) KeyLast(key string) Query {
	return q.KeyN(key, -1)
}

func (q Query) KeyFirstN(key string, n int) Query {
	return q.KeyRange(key, 0, n-1)
}

func (q Query) KeyLastN(key string, n int) Query {
	return q.KeyRange(key, -n, -1)
}

func (q Query) Range(start, end int) Query {
	match := Query{}
	for _, node := range q {
		s, e := calcRange(start, end, len(node.Kids))
		match = append(match, node.Kids[s:e]...)
	}
	return match
}

func (q Query) All() Query {
	return q.Range(0, -1)
}

func (q Query) At(n int) Query {
	return q.Range(n, n)
}

func (q Query) First() Query {
	return q.At(0)
}

func (q Query) Last() Query {
	return q.At(-1)
}

func (q Query) FirstN(n int) Query {
	return q.Range(0, n-1)
}

func (q Query) LastN(n int) Query {
	return q.Range(-n, -1)
}

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
