package app

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/android-lewis/note-taker/pkg/util"
)

const testPath = "../../testindex.json"

func TestAddIndex(t *testing.T) {
	defer util.ClearIndex(testPath)
	id := time.Now().UnixMicro()
	path := "./data/test.md"

	indexMap := []util.Index{}

	newindex := util.Index{ID: id, Path: path}
	indexMap = append(indexMap, newindex)

	util.WriteJSON(testPath, indexMap)
	data := util.ReadJSON(testPath)

	for i, i2 := range data {
		fmt.Println(i)
		if i2.ID == id && i2.Path == path {
			return
		}
	}

	t.Errorf("Index not added to test index %v", data)

}

func TestSearch(t *testing.T) {
	defer util.ClearIndex(testPath)
	var id int64

	indexMap := []util.Index{}
	randomSeed := rand.Intn(1000)

	for i := 0; i < 1000; i++ {
		if i == randomSeed {
			id = time.Now().UnixMicro()
			path := "./data/test.md"
			newindex := util.Index{ID: id, Path: path}
			indexMap = append(indexMap, newindex)
		} else {
			fakeId := time.Now().UnixMicro()
			fakePath := "./data/fakeTest.md"
			fakeindex := util.Index{ID: fakeId, Path: fakePath}
			indexMap = append(indexMap, fakeindex)
		}
	}

	util.WriteJSON(testPath, indexMap)
	index, err := util.SearchIndex(testPath, id)

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
	defer util.ClearIndex(testPath)
	var id int64

	indexMap := []util.Index{}
	randomSeed := rand.Intn(1000000)

	for i := 0; i < 1000000; i++ {
		if i == randomSeed {
			id = time.Now().UnixMicro()
			path := "./data/test.md"
			newindex := util.Index{ID: id, Path: path}
			indexMap = append(indexMap, newindex)
		} else {
			fakeId := time.Now().UnixMicro()
			fakePath := "./data/fakeTest.md"
			fakeindex := util.Index{ID: fakeId, Path: fakePath}
			indexMap = append(indexMap, fakeindex)
		}
	}

	util.WriteJSON(testPath, indexMap)
	// End setup

	b.ResetTimer()
	_, err := util.SearchIndex(testPath, id)

	if err != nil {
		b.Errorf("could not find index: %s", err)
	}

}
