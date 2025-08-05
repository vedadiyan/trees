package tree

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	links := Links[string]{
		{"7", "4"},
		{"4", "5"},
		{"5", "8"},
		{"8", "9"},
		{"4", "6"},
		{"5", "10"},
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
