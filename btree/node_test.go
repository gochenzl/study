package btree

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

func TestNode(t *testing.T) {
	a := &node{}
	a.items = make([]item, 6)

	a.appendItem(1, 1)
	a.appendItem(2, 2)
	a.appendItem(3, 3)

	if a.itemLen() != 3 {
		t.Errorf("appendItem")
	}

	if (a.items[0] != item{1, 1}) ||
		(a.items[1] != item{2, 2}) ||
		(a.items[2] != item{3, 3}) {
		t.Error("appendItem")
	}

	a.insertItem(5, 5, intCompare)
	if a.itemLen() != 4 {
		t.Errorf("insertItem")
	}

	if (a.items[0] != item{1, 1}) ||
		(a.items[1] != item{2, 2}) ||
		(a.items[2] != item{3, 3}) ||
		(a.items[3] != item{5, 5}) {
		t.Error("insertItem")
	}

	a.insertItem(0, 0, intCompare)
	if a.itemLen() != 5 {
		t.Errorf("insertItem")
	}

	if (a.items[0] != item{0, 0}) ||
		(a.items[1] != item{1, 1}) ||
		(a.items[2] != item{2, 2}) ||
		(a.items[3] != item{3, 3}) ||
		(a.items[4] != item{5, 5}) {
		t.Error("insertItem")
	}

	a.insertItem(4, 4, intCompare)
	if a.itemLen() != 6 {
		t.Errorf("insertItem")
	}

	if (a.items[0] != item{0, 0}) ||
		(a.items[1] != item{1, 1}) ||
		(a.items[2] != item{2, 2}) ||
		(a.items[3] != item{3, 3}) ||
		(a.items[4] != item{4, 4}) ||
		(a.items[5] != item{5, 5}) {
		t.Error("insertItem")
	}

	if !a.isFull() {
		t.Error("insertItem")
	}

	a.truncItem(3)
	if a.itemLen() != 3 ||
		(a.items[0] != item{0, 0}) ||
		(a.items[1] != item{1, 1}) ||
		(a.items[2] != item{2, 2}) {
		t.Error("truncItem")
	}

	a.appendItem(3, 3)

	a.delItem(1)
	if a.itemLen() != 3 ||
		(a.items[0] != item{0, 0}) ||
		(a.items[1] != item{2, 2}) ||
		(a.items[2] != item{3, 3}) {
		t.Error("delItem")
	}

	a.delItem(1)
	if a.itemLen() != 2 ||
		(a.items[0] != item{0, 0}) ||
		(a.items[1] != item{3, 3}) {
		t.Error("delItem")
	}

	a.delItem(1)
	if a.itemLen() != 1 ||
		(a.items[0] != item{0, 0}) {
		t.Error("delItem")
	}

	a.delItem(0)
	if a.itemLen() != 0 {
		t.Error("delItem")
	}
}

func TestNodeHelp(t *testing.T) {
	a := &node{}
	a.items = make([]item, 6)

	a.appendItem(1, 1)
	a.appendItem(3, 1)
	a.appendItem(5, 1)

	if i, found := findItem(a, 1, intCompare); i != 0 || !found {
		t.Error("findItem")
	}

	if i, found := findItem(a, 3, intCompare); i != 1 || !found {
		t.Error("findItem")
	}

	if i, found := findItem(a, 5, intCompare); i != 2 || !found {
		t.Error("findItem")
	}

	if i, found := findItem(a, 2, intCompare); i != 0 || found {
		t.Error("findItem")
	}

	if findPos(a, 0, intCompare) != 0 {
		t.Error("findPos")
	}

	if findPos(a, 2, intCompare) != 1 {
		t.Error("findPos")
	}

	if findPos(a, 4, intCompare) != 2 {
		t.Error("findPos")
	}

	if findPos(a, 6, intCompare) != 3 {
		t.Error("findPos")
	}
}
