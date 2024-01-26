package util

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const testPath = "../../testindex.json"

func TestAddIndex(t *testing.T) {
	defer RemoveFile(testPath)
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

func generateRandomIndexMap(length int) ([]Index, int64) {
	var id int64

	indexMap := []Index{}
	randomSeed := rand.Intn(length)

	for i := 0; i < length; i++ {
		time.Sleep(time.Microsecond) //Generate articificial time difference (possible duplicates)
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

	return indexMap, id
}

func TestSearch(t *testing.T) {
	defer RemoveFile(testPath)
	indexMap, id := generateRandomIndexMap(1000)

	writeJSON(testPath, indexMap)
	index, _, err := SearchIndex(testPath, id)

	if err != nil {
		t.Errorf("could not find index: %v", index)
	}

	if index.ID != id {
		t.Errorf("incorrect ID returned: %v", index)
	}

}

func TestRemoveIndex(t *testing.T) {
	indexMap, id := generateRandomIndexMap(1000)
	_, i, err := performSearch(indexMap, id)
	if err != nil {
		t.Errorf("could not find id: %d", id)
	}

	indexMap = removeIndex(indexMap, i)

	_, _, err = performSearch(indexMap, id)
	if err == nil {
		t.Errorf("found id in index map: %d", id)
	}

	fmt.Println(err)
}

func BenchmarkSearch(b *testing.B) {
	// Start setup
	defer RemoveFile(testPath)
	indexMap, id := generateRandomIndexMap(100000)

	writeJSON(testPath, indexMap)
	// End setup

	b.ResetTimer()
	_, _, err := SearchIndex(testPath, id)

	if err != nil {
		b.Errorf("could not find index: %s", err)
	}

}
