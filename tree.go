package tree

import "fmt"

type (
	Link[T comparable] struct {
		Src  T
		Dest T
	}

	Links[T comparable] []Link[T]

	SortedTree[T comparable] struct {
		Src  T
		Dest []*SortedTree[T]
	}
)

func Sort[T comparable](links Links[T]) ([]*SortedTree[T], error) {
	sources := links.findSources()
	if len(sources) == 0 {
		return nil, fmt.Errorf("no source found")
	}
	return links.aggregate(sources), nil
}

func (l Links[T]) follow(src Link[T]) *SortedTree[T] {
	filteredLinks := l.filter(src)
	var out = new(SortedTree[T])
	out.Dest = make([]*SortedTree[T], 0)
	out.Src = src.Src
	nextLinks := filteredLinks.next(src)
	if len(nextLinks) != 0 {
		out.Dest = append(out.Dest, filteredLinks.aggregate(nextLinks)...)
		return out
	}
	lastNode := new(SortedTree[T])
	lastNode.Src = src.Dest
	out.Dest = append(out.Dest, lastNode)
	return out
}

func (l Links[T]) aggregate(nextLinks Links[T]) []*SortedTree[T] {
	c := make(map[T][]*SortedTree[T])
	dest := make([]*SortedTree[T], 0)
	for _, next := range nextLinks {
		if _, ok := c[next.Src]; !ok {
			c[next.Src] = make([]*SortedTree[T], 0)
		}
		c[next.Src] = append(c[next.Src], l.follow(next))
	}
	for key, values := range c {
		aggregatedNode := new(SortedTree[T])
		aggregatedNode.Src = key
		aggregatedNode.Dest = make([]*SortedTree[T], 0)
		for _, value := range values {
			aggregatedNode.Dest = append(aggregatedNode.Dest, value.Dest...)
		}
		dest = append(dest, aggregatedNode)
	}
	return dest
}

func (l Links[T]) next(source Link[T]) Links[T] {
	out := make(Links[T], 0)
	for _, x := range l {
		if src, dest := source.Dest, x.Src; src == dest {
			out = append(out, x)
		}
	}
	return out
}

func (l Links[T]) filter(src Link[T]) Links[T] {
	out := make(Links[T], 0)
	for _, x := range l {
		if x.Src == src.Src && x.Dest == src.Dest {
			continue
		}
		out = append(out, x)
	}
	return out
}

func (l Links[T]) findSources() Links[T] {
	out := make(Links[T], 0)

	for _, x := range l {
		if !l.isDest(x.Src) {
			out = append(out, x)
		}
	}

	return out
}

func (l Links[T]) isDest(n T) bool {
	for _, x := range l {
		if x.Dest == n {
			return true
		}
	}
	return false
}
