package btree

type item struct {
	key interface{}
	val interface{}
}

type node struct {
	parent     *node
	itemCount  int
	items      []item
	children   []*node
	childCount int
}

func (a *node) itemLen() int {
	return a.itemCount
}

func (a *node) isFull() bool {
	return a.itemCount == len(a.items)
}

func (a *node) appendItem(key, val interface{}) {
	a.items[a.itemCount] = item{key, val}
	a.itemCount++
}

func (a *node) insertItem(key, val interface{}, cmp Comparer) int {
	pos := findPos(a, key, cmp)

	a.appendItem(key, val)
	if pos < a.itemLen() {
		copy(a.items[pos+1:a.itemLen()], a.items[pos:a.itemLen()-1])
		a.items[pos] = item{key, val}
	}

	return pos
}

func (a *node) truncItem(newLen int) {
	a.itemCount = newLen
}

func (a *node) delItem(pos int) {
	len := a.itemLen()
	if pos >= len {
		return
	}

	if pos == len-1 {
		a.truncItem(len - 1)
	} else {
		copy(a.items[pos:len], a.items[pos+1:len])
		a.itemCount--
	}
}

func (a *node) appendChild(child *node) {
	a.children[a.childCount] = child
	a.childCount++
}

func search(a *node, key interface{}, cmp Comparer) (*node, int) {
	for a != nil {
		i, found := findItem(a, key, cmp)
		if found {
			return a, i
		}

		if a.children == nil {
			break
		}

		var pos int
		if cmp(key, a.items[i].key) < 0 {
			pos = i
		} else {
			pos = i + 1
		}

		a = a.children[pos]
	}

	return a, -1
}

func findItem(a *node, key interface{}, cmp Comparer) (i int, found bool) {
	l := 0
	h := a.itemLen() - 1
	for l <= h {
		i = (l + h) >> 1
		ret := cmp(key, a.items[i].key)
		if ret < 0 {
			h = i - 1
		} else if ret > 0 {
			l = i + 1
		} else {
			found = true
			return
		}
	}

	return
}

func findPos(a *node, key interface{}, cmp Comparer) int {
	i, found := findItem(a, key, cmp)
	if found {
		panic("findPos input exist key")
	}

	var pos int
	if cmp(key, a.items[i].key) < 0 {
		pos = i
	} else {
		pos = i + 1
	}

	return pos

	//	if cmp(key, a.items[0].key) < 0 {
	//		return 0
	//	}

	//	if cmp(key, a.items[a.itemLen()-1].key) > 0 {
	//		return a.itemLen()
	//	}

	//	for i := 0; i < a.itemLen()-1; i++ {
	//		if cmp(key, a.items[i].key) > 0 &&
	//			cmp(key, a.items[i+1].key) < 0 {
	//			return i + 1
	//		}
	//	}

	// don't come here
	//return -1
}

func findLargest(a *node) *node {
	for a.children != nil {
		a = a.children[a.childCount-1]
	}

	return a
}
