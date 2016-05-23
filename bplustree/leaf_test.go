package bplustree

import "testing"

func intCompare(a, b interface{}) int {
	ia := a.(int)
	ib := b.(int)

	if ia < ib {
		return -1
	}

	if ia > ib {
		return 1
	}

	return 0
}

func checkItems(l *leaf, items []item) bool {
	if l.count != len(items) {
		return false
	}

	for i := 0; i < len(items); i++ {
		if l.items[i] != items[i] {
			return false
		}
	}

	return true
}

func TestLeaf(t *testing.T) {
	l := &leaf{}
	items := []item{item{5, 1}, {9, 1}, {12, 1}, {16, 1}, {20, 1}}

	l.items = make([]item, len(items)+5)

	for i := 0; i < len(items); i++ {
		l.append(items[i].key, items[i].val)
	}

	if !checkItems(l, items) {
		t.Error("append")
	}

	items = []item{item{5, 1}, {6, 1}, {9, 1}, {12, 1}, {16, 1}, {20, 1}}
	l.insert(6, 1, intCompare)
	if !checkItems(l, items) {
		t.Error("insert")
	}

	items = []item{item{5, 1}, {6, 1}, {7, 1}, {9, 1}, {12, 1}, {16, 1}, {20, 1}}
	l.insert(7, 1, intCompare)
	if !checkItems(l, items) {
		t.Error("insert")
	}

	items = []item{item{2, 1}, item{5, 1}, {6, 1}, {7, 1}, {9, 1}, {12, 1}, {16, 1}, {20, 1}}
	l.insert(2, 1, intCompare)
	if !checkItems(l, items) {
		t.Error("insert")
	}

	items = []item{item{2, 1}, item{5, 1}, {6, 1}, {7, 1}, {9, 1}, {12, 1}, {16, 1}, {20, 1}, {25, 1}}
	l.insert(25, 1, intCompare)
	if !checkItems(l, items) {
		t.Error("insert")
	}

	if l.full() {
		t.Error("isFull")
	}

	items = []item{item{2, 1}, item{5, 1}, {6, 1}, {7, 1}, {9, 1}, {12, 1}, {16, 1}, {20, 1}, {25, 1}, {28, 1}}
	l.insert(28, 1, intCompare)
	if !checkItems(l, items) {
		t.Error("insert")
	}

	if !l.full() {
		t.Error("isFull")
	}

	if i, found := l.find(0, intCompare); i != 0 || found {
		t.Error("find")
	}

	if i, found := l.find(8, intCompare); i != 3 || found {
		t.Error("find")
	}

	if i, found := l.find(30, intCompare); i != 9 || found {
		t.Error("find")
	}

	if i, found := l.find(2, intCompare); i != 0 || !found {
		t.Error("find")
	}

	if i, found := l.find(28, intCompare); i != 9 || !found {
		t.Error("find")
	}

	if i, found := l.find(9, intCompare); i != 4 || !found {
		t.Error("find")
	}

	newleaf := &leaf{}
	newleaf.items = make([]item, l.count)
	middleKey := l.split(newleaf)
	if !checkItems(l, []item{{2, 1}, {5, 1}, {6, 1}, {7, 1}, {9, 1}}) {
		t.Error("split")
	}
	if !checkItems(newleaf, []item{{12, 1}, {16, 1}, {20, 1}, {25, 1}, {28, 1}}) {
		t.Error("split")
	}
	if l.next != newleaf || newleaf.pre != l {
		t.Error("split")
	}
	if middleKey != 12 {
		t.Error("split")
	}

	l.del(4)
	if !checkItems(l, []item{{2, 1}, {5, 1}, {6, 1}, {7, 1}}) {
		t.Error("del")
	}

	l.del(0)
	if !checkItems(l, []item{{5, 1}, {6, 1}, {7, 1}}) {
		t.Error("del")
	}

	l.del(1)
	if !checkItems(l, []item{{5, 1}, {7, 1}}) {
		t.Error("del")
	}

}
