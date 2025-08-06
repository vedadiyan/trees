package tree

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	links := Links[string]{
		{Src: "A", Dest: "B"},
		{Src: "A", Dest: "C"},
		{Src: "B", Dest: "D"},
		{Src: "C", Dest: "D"},
	}

	trees, err := NewForest(links)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	json, err := json.MarshalIndent(trees, "", "\t")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	os.WriteFile("test.json", json, os.ModePerm)
}
