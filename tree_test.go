package tree

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	links := Links[string]{
		{"A", "B"},
		{"A", "C"},
	}

	trees, err := Sort(links)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	json, err := json.MarshalIndent(trees, "", "\t")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	os.WriteFile("test.json", json, os.ModePerm)
}
