package bst

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func genDatas(size int) []int {
	rand.Seed(int64(time.Now().Nanosecond()))
	var flags map[int]bool = make(map[int]bool)
	datas := make([]int, size)
	for i := 0; i < size; i++ {
		var val int
		for {
			val = rand.Int()
			if _, present := flags[val]; !present {
				break
			}
		}

		datas[i] = val
		flags[val] = true
	}

	return datas
}

func TestIterator(t *testing.T) {
	testSize := 1000
	datas := genDatas(testSize)
	tree := NewUnbalanced(intCompare)
	for i := 0; i < len(datas); i++ {
		tree.Insert(datas[i], datas[i])
	}

	iterDatas := make([]int, 0, testSize)
	iter := tree.NewIterator()
	for iter.Next() {
		iterDatas = append(iterDatas, iter.Key().(int))
	}

	if len(iterDatas) != testSize {
		t.Errorf("iterator fail")
	}

	sort.Sort(sort.IntSlice(datas))
	for i := 0; i < testSize; i++ {
		if datas[i] != iterDatas[i] {
			t.Errorf("iterator fail")
		}
	}
}
