package util

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const testPath = "../../testindex.json"

func TestAddIndex(t *testing.T) {
	defer clearIndex(testPath)
	id := time.Now().UnixMicro()
	path := "./data/test.md"

	indexMap := []Index{}

	newindex := Index{ID: id, Path: path}
	indexMap = append(indexMap, newindex)

	writeJSON(testPath, indexMap)
	data := readJSON(testPath)

	for _, i2 := range data {
		if i2.ID == id && i2.Path == path {
			return
		}
	}

	t.Errorf("Index not added to test index %v", data)

}

func TestSearch(t *testing.T) {
	defer clearIndex(testPath)
	var id int64

	indexMap := []Index{}
	randomSeed := rand.Intn(1000)

	for i := 0; i < 1000; i++ {
		if i == randomSeed {
			id = time.Now().UnixMicro()
			path := "./data/test.md"
			newindex := Index{ID: id, Path: path}
			indexMap = append(indexMap, newindex)
		} else {
			fakeId := time.Now().UnixMicro()
			fakePath := "./data/fakeTest.md"
			fakeindex := Index{ID: fakeId, Path: fakePath}
			indexMap = append(indexMap, fakeindex)
		}
	}

	writeJSON(testPath, indexMap)
	index, _, err := SearchIndex(testPath, id)

	if err != nil {
		t.Errorf("could not find index: %v", index)
	}

	if index.ID != id {
		t.Errorf("incorrect ID returned: %v", index)
	}

	fmt.Printf("ID: %d Index: %v\n", id, index)

}

func BenchmarkSearch(b *testing.B) {
	// Start setup
	defer clearIndex(testPath)
	var id int64

	indexMap := []Index{}
	randomSeed := rand.Intn(1000000)

	for i := 0; i < 1000000; i++ {
		if i == randomSeed {
			id = time.Now().UnixMicro()
			path := "./data/test.md"
			newindex := Index{ID: id, Path: path}
			indexMap = append(indexMap, newindex)
		} else {
			fakeId := time.Now().UnixMicro()
			fakePath := "./data/fakeTest.md"
			fakeindex := Index{ID: fakeId, Path: fakePath}
			indexMap = append(indexMap, fakeindex)
		}
	}

	writeJSON(testPath, indexMap)
	// End setup

	b.ResetTimer()
	_, _, err := SearchIndex(testPath, id)

	if err != nil {
		b.Errorf("could not find index: %s", err)
	}

}
