package tree

import (
	"reflect"
	"testing"
)

func TestSort_Success(t *testing.T) {
	links := Links[string]{
		{"a", "b"},
		{"b", "c"},
		{"a", "d"},
		{"d", "e"},
	}

	trees, err := Sort(links)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	root := trees[0]
	if root.Src != "a" || root.Dest != "b" {
		t.Fatalf("unexpected root: %+v", root)
	}

	expectTreeStructure(t, root,
		[]*SortedTree[string]{
			{
				Link: Link[string]{Src: "b", Dest: "c"},
			},
		},
	)
}

func TestSort_NoSource(t *testing.T) {
	links := Links[string]{
		{"a", "b"},
		{"b", "a"}, // cycle
	}

	_, err := Sort(links)
	if err == nil {
		t.Fatal("expected error due to cycle, got nil")
	}
}

func Test_findSources(t *testing.T) {
	links := Links[string]{
		{"a", "b"},
		{"b", "c"},
	}

	sources := links.findSources()
	expected := Links[string]{{"a", "b"}}
	if !reflect.DeepEqual(sources, expected) {
		t.Errorf("expected %v, got %v", expected, sources)
	}
}

func Test_next(t *testing.T) {
	links := Links[string]{
		{"a", "b"},
		{"b", "c"},
		{"b", "d"},
	}
	next := links.next(Link[string]{Src: "a", Dest: "b"})
	expected := Links[string]{
		{"b", "c"},
		{"b", "d"},
	}
	if !reflect.DeepEqual(next, expected) {
		t.Errorf("expected %v, got %v", expected, next)
	}
}

func Test_filter(t *testing.T) {
	links := Links[string]{
		{"a", "b"},
		{"b", "c"},
	}
	filtered := links.filter(Link[string]{Src: "a", Dest: "b"})
	expected := Links[string]{
		{"b", "c"},
	}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

func expectTreeStructure(t *testing.T, node *SortedTree[string], expected []*SortedTree[string]) {
	if len(node.Descendants) != len(expected) {
		t.Fatalf("expected %d children, got %d", len(expected), len(node.Descendants))
	}
	for i, child := range expected {
		if node.Descendants[i].Src != child.Src || node.Descendants[i].Dest != child.Dest {
			t.Errorf("expected %v, got %v", child, node.Descendants[i])
		}
	}
}
