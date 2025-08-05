package tree

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestSort(t *testing.T) {
	input := Links[int]{
		{8, 9},
		{4, 5},
		{5, 10},
		{4, 6},
		{4, 9},
		{7, 4},
		{5, 8},
	}

	sorted, err := Sort(input)
	if err != nil {
		log.Fatalln(err)
	}
	json, err := json.MarshalIndent(sorted, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile("test.json", json, os.ModePerm)
}
