package tree

import "fmt"

type (
	Link[T comparable] struct {
		Src  T
		Dest T
	}

	Links[T comparable] []Link[T]

	SortedTree[T comparable] struct {
		Node T
		Next []*SortedTree[T]
	}
)

func NewForest[T comparable](links Links[T]) ([]*SortedTree[T], error) {
	sources := links.findSources()
	if len(sources) == 0 {
		return nil, fmt.Errorf("no source found")
	}
	return links.followMany(sources), nil
}

func (l Links[T]) followOne(link Link[T]) *SortedTree[T] {
	filteredLinks := l.filter(link)
	var out = new(SortedTree[T])
	out.Next = make([]*SortedTree[T], 0)
	out.Node = link.Src
	nextLinks := filteredLinks.next(link)
	if len(nextLinks) != 0 {
		out.Next = append(out.Next, filteredLinks.followMany(nextLinks)...)
		return out
	}
	lastNode := new(SortedTree[T])
	lastNode.Node = link.Dest
	out.Next = append(out.Next, lastNode)
	return out
}

func (l Links[T]) followMany(links Links[T]) []*SortedTree[T] {
	out := make([]*SortedTree[T], 0)
	aC := make(map[T][]*SortedTree[T])
	for _, next := range links {
		if _, ok := aC[next.Src]; !ok {
			aC[next.Src] = make([]*SortedTree[T], 0)
		}
		aC[next.Src] = append(aC[next.Src], l.followOne(next))
	}
	for key, values := range aC {
		aN := new(SortedTree[T])
		aN.Node = key
		aN.Next = make([]*SortedTree[T], 0)
		for _, value := range values {
			aN.Next = append(aN.Next, value.Next...)
		}
		out = append(out, aN)
	}
	return out
}

func (l Links[T]) next(link Link[T]) Links[T] {
	out := make(Links[T], 0)
	for _, x := range l {
		if src, dest := link.Dest, x.Src; src == dest {
			out = append(out, x)
		}
	}
	return out
}

func (l Links[T]) filter(link Link[T]) Links[T] {
	out := make(Links[T], 0)
	for _, x := range l {
		if x.Src == link.Src && x.Dest == link.Dest {
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
