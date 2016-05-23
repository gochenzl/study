package bplustree

import "testing"

func checkKeys(a *node, keys []int) bool {
	if a.keyCount != len(keys) {
		return false
	}

	for i := 0; i < a.keyCount; i++ {
		if a.keys[i] != keys[i] {
			return false
		}
	}

	return true
}

func TestNode(t *testing.T) {
	a := &node{}
	a.keys = make([]interface{}, 10)
	keys := []int{5, 8, 11, 16, 20}

	for i := 0; i < len(keys); i++ {
		a.appendKey(keys[i])
	}
	if !checkKeys(a, keys) {
		t.Error("appendKey")
	}

	keys = []int{3, 5, 8, 11, 16, 20}
	if a.insertKey(3, intCompare) != 0 {
		t.Error("insertKey")
	}
	if !checkKeys(a, keys) {
		t.Error("insertKey")
	}

	keys = []int{3, 5, 6, 8, 11, 16, 20}
	if a.insertKey(6, intCompare) != 2 {
		t.Error("insertKey")
	}
	if !checkKeys(a, keys) {
		t.Error("insertKey")
	}

	keys = []int{3, 5, 6, 8, 10, 11, 16, 20}
	if a.insertKey(10, intCompare) != 4 {
		t.Error("insertKey")
	}
	if !checkKeys(a, keys) {
		t.Error("insertKey")
	}

	keys = []int{3, 5, 6, 8, 10, 11, 16, 20, 21}
	if a.insertKey(21, intCompare) != 8 {
		t.Error("insertKey")
	}
	if !checkKeys(a, keys) {
		t.Error("insertKey")
	}

	if a.full() {
		t.Error("isFull")
	}

	keys = []int{3, 5, 6, 8, 10, 11, 16, 20, 21, 25}
	if a.insertKey(25, intCompare) != 9 {

		t.Error("insertKey")
	}
	if !checkKeys(a, keys) {
		t.Error("insertKey")
	}

	if !a.full() {
		t.Error("isFull")
	}
}
